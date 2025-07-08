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
	// Mount points for the generator container.
	librarianDir = "/librarian"
	inputDir     = "/input"
	outputDir    = "/output"
	sourceDir    = "/source"
)

// main is the entrypoint for the generator container.
func main() {
	slog.Info("Go generator invoked", "args", os.Args)
	if err := run(context.Background()); err != nil {
		slog.Error("generator failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Go generator finished successfully")
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
		slog.Warn("configure command is not yet implemented")
		return nil
	case "build":
		slog.Warn("build command is not yet implemented")
		return nil
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}

// generateCmd handles the logic for the 'generate' container command.
// It reads a request file, and for each API specified, it invokes protoc
// to generate the client library.
func generateCmd(ctx context.Context) error {
	slog.Info("generate command started")

	reqPath := filepath.Join(librarianDir, "generate-request.json")
	slog.Info("reading generate request", "path", reqPath)

	reqFile, err := os.ReadFile(reqPath)
	if err != nil {
		// The user provided mock input is named configure-request.json, so we'll
		// check for that as a fallback for local testing.
		reqPath = filepath.Join(librarianDir, "configure-request.json")
		slog.Info("generate-request.json not found, trying configure-request.json", "path", reqPath)
		reqFile, err = os.ReadFile(reqPath)
		if err != nil {
			return fmt.Errorf("failed to read request file from %s: %w", librarianDir, err)
		}
	}

	var lib Library
	if err := json.Unmarshal(reqFile, &lib); err != nil {
		return fmt.Errorf("failed to unmarshal request file %s: %w", reqPath, err)
	}
	slog.Info("successfully unmarshalled request", "library_id", lib.ID)

	for _, api := range lib.APIs {
		if err := protoc(ctx, &lib, &api); err != nil {
			return fmt.Errorf("protoc failed for api %q in library %q: %w", api.Path, lib.ID, err)
		}
	}

	slog.Info("generate command finished")
	return nil
}

// protoc constructs and executes the protoc command for a given API.
func protoc(ctx context.Context, lib *Library, api *API) error {
	apiSourceDir := filepath.Join(sourceDir, api.Path)
	slog.Info("running protoc", "api_source_dir", apiSourceDir)

	// Gather all .proto files in the API's source directory.
	entries, err := os.ReadDir(apiSourceDir)
	if err != nil {
		return fmt.Errorf("failed to read API source directory %s: %w", apiSourceDir, err)
	}

	var protoFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".proto" {
			protoFiles = append(protoFiles, filepath.Join(apiSourceDir, entry.Name()))
		}
	}

	if len(protoFiles) == 0 {
		return fmt.Errorf("no .proto files found in %s", apiSourceDir)
	}
	slog.Info("found proto files", "files", protoFiles)

	// Determine the Go import path for the generated GAPIC client.
	// This is a simplification. A real implementation might get this from a
	// more reliable source in the request. We assume the first source_path
	// is the primary one for the library.
	var gapicImportPath string
	if len(lib.SourcePaths) > 0 {
		// Assuming a base module path. This should ideally be discovered or configured.
		// Example: cloud.google.com/go/storage/apiv1
		gapicImportPath = filepath.Join("cloud.google.com/go", lib.SourcePaths[0])
	} else {
		return fmt.Errorf("cannot determine go-gapic-package: source_paths is empty for library %s", lib.ID)
	}

	// Construct the protoc command arguments.
	args := []string{
		"--experimental_allow_proto3_optional",
		"--go_out=" + outputDir,
		"--go-gapic_out=" + outputDir,
		"--go-gapic_opt=go-gapic-package=" + gapicImportPath,
		// The 'input' directory from the language repo is mounted at /input.
		// It can contain shared proto definitions.
		"-I=" + sourceDir,
		"-I=" + inputDir,
	}
	if api.ServiceConfig != "" {
		args = append(args, "--go-gapic_opt=api-service-config="+filepath.Join(apiSourceDir, api.ServiceConfig))
	}

	args = append(args, protoFiles...)

	cmd := exec.CommandContext(ctx, "protoc", args...)
	return runCommand(cmd)
}

// runCommand executes a command and logs its output.
func runCommand(cmd *exec.Cmd) error {
	cmd.Env = os.Environ()
	cmd.Dir = outputDir // Run commands from the output directory.
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

// Library corresponds to a single library definition from the state file.
type Library struct {
	ID          string   `json:"id"`
	Version     string   `json:"version"`
	APIs        []API    `json:"apis"`
	SourcePaths []string `json:"source_paths"`
}

// API corresponds to a single API definition within a library.
type API struct {
	Path          string `json:"path"`
	ServiceConfig string `json:"service_config"`
}
