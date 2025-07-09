// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// bazel-generator is an alternative implementation of the Go Librarian container
// that uses Bazel for the generation and build steps.
//
// This approach is maintained as an alternative strategy, as discussed in
// go/sdk-librarian-python, representing a "Phase 1: Owlbot in a box" model.
// It wraps the existing Bazel-based generation process to ensure a smooth
// transition with minimal initial code changes for teams that rely on it.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	// Mount points for the generator container, as per the contract.
	librarianDir = "/librarian"
	inputDir     = "/input"
	outputDir    = "/output"
	sourceDir    = "/source"
	repoDir      = "/repo"
)

// Library represents a single library definition from the state file.
type Library struct {
	ID                  string   `json:"id"`
	Version             string   `json:"version,omitempty"`
	LastGeneratedCommit string   `json:"last_generated_commit,omitempty"`
	APIs                []API    `json:"apis"`
	SourcePaths         []string `json:"source_paths,omitempty"`
	PreserveRegex       []string `json:"preserve_regex,omitempty"`
	RemoveRegex         []string `json:"remove_regex,omitempty"`
}

// API represents a single API definition within a library.
type API struct {
	Path          string `json:"path"`
	ServiceConfig string `json:"service_config"`
}

// main is the entrypoint for the Bazel-based Go generator container.
func main() {
	slog.Info("Bazel Go generator invoked", "args", os.Args)
	if err := run(context.Background()); err != nil {
		slog.Error("generator failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Bazel Go generator finished successfully")
}

// run executes the appropriate command based on the container's invocation arguments.
func run(ctx context.Context) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("expected at least one argument for the command, got %d", len(os.Args)-1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "generate":
		return generateCmd(ctx)
	case "configure":
		return configureCmd(ctx)
	case "build":
		return buildCmd(ctx)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}

// configureCmd handles the 'configure' step for a new library.
func configureCmd(ctx context.Context) error {
	slog.Info("configure command started")
	reqPath := filepath.Join(librarianDir, "configure-request.json")
	respPath := filepath.Join(librarianDir, "configure-response.json")

	reqFile, err := os.ReadFile(reqPath)
	if err != nil {
		return fmt.Errorf("failed to read configure-request.json from %s: %w", librarianDir, err)
	}

	var lib Library
	if err := json.Unmarshal(reqFile, &lib); err != nil {
		return fmt.Errorf("failed to unmarshal configure request: %w", err)
	}

	// Enrich the library config with Go-specific defaults.
	lib.Version = "0.1.0"
	goStylePath := strings.Replace(lib.ID, "-", "/", -1)
	lib.SourcePaths = []string{filepath.Join("internal", goStylePath)}
	lib.RemoveRegex = lib.SourcePaths

	respFile, err := json.MarshalIndent(lib, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal configure response: %w", err)
	}

	if err := os.WriteFile(respPath, respFile, 0644); err != nil {
		return fmt.Errorf("failed to write configure-response.json to %s: %w", respPath, err)
	}
	slog.Info("successfully wrote configuration", "path", respPath)
	return nil
}

// generateCmd handles the 'generate' step by invoking a Bazel-based process.
func generateCmd(ctx context.Context) error {
	slog.Info("generate command started (Bazel strategy)")
	reqPath := filepath.Join(librarianDir, "generate-request.json")

	reqFile, err := os.ReadFile(reqPath)
	if err != nil {
		return fmt.Errorf("failed to read generate-request.json from %s: %w", librarianDir, err)
	}

	var lib Library
	if err := json.Unmarshal(reqFile, &lib); err != nil {
		return fmt.Errorf("failed to unmarshal generate request: %w", err)
	}

	// In this strategy, we invoke a hypothetical generator CLI that wraps Bazel.
	// This simulates the "Owlbot in a box" approach.
	slog.Info("invoking Bazel-based generator script")
	args := []string{
		"run",
		"//tools/generator:main", // Example Bazel target for the generator
		"--",
		"generate-library",
		"--api-root=" + sourceDir,
		"--generator-input=" + inputDir,
		"--output=" + outputDir,
		"--library-id=" + lib.ID,
	}
	cmd := exec.CommandContext(ctx, "bazel", args...)
	cmd.Dir = repoDir // Bazel commands typically run from the repo root.
	return runCommand(cmd)
}

// buildCmd handles the 'build' step by invoking 'bazel test'.
func buildCmd(ctx context.Context) error {
	slog.Info("build command started (Bazel strategy)")
	reqPath := filepath.Join(librarianDir, "build-request.json")

	reqFile, err := os.ReadFile(reqPath)
	if err != nil {
		return fmt.Errorf("failed to read build-request.json from %s: %w", librarianDir, err)
	}

	var lib Library
	if err := json.Unmarshal(reqFile, &lib); err != nil {
		return fmt.Errorf("failed to unmarshal build request: %w", err)
	}

	// Determine the Bazel target from the library's source paths.
	if len(lib.SourcePaths) == 0 {
		return fmt.Errorf("cannot determine build target: source_paths is empty for library %s", lib.ID)
	}
	target := "//" + lib.SourcePaths[0] + "/..." // e.g., //internal/storage/...
	slog.Info("determined bazel target", "target", target)

	args := []string{"test", target}
	cmd := exec.CommandContext(ctx, "bazel", args...)
	cmd.Dir = repoDir
	return runCommand(cmd)
}

// runCommand executes a command and logs its output.
func runCommand(cmd *exec.Cmd) error {
	cmd.Env = os.Environ()
	slog.Info("running command", "command", strings.Join(cmd.Args, " "), "dir", cmd.Dir)

	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		slog.Info("command output", "output", string(output))
	}
	if err != nil {
		return fmt.Errorf("command failed with error: %w", err)
	}
	return nil
}