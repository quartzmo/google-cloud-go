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

// The following constants define the directory layout that this generator expects
// to be present in its container, as per the Librarian container contract.
const (
	// librarianDir is the mount point for inputs from the Librarian tool itself.
	// It contains the primary request file (e.g., generate-request.json).
	librarianDir = "/librarian"

	// inputDir is the mount point for the `.librarian/generator-input` directory
	// from the language repository. It is reserved for language-specific templates
	// or post-processing scripts and is NOT used as a proto import path.
	inputDir = "/input"

	// outputDir is the mount point for an empty directory where this generator
	// must write all its output. The Librarian tool is responsible for copying
	// the contents of this directory to the language repository.
	outputDir = "/output"

	// sourceDir is the mount point for a complete checkout of the googleapis
	// repository. It serves as the sole proto import path for protoc.
	sourceDir = "/source"
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
// The first argument to the container is always the command (e.g., "generate").
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

	// The request file tells the generator which library and APIs to generate.
	// It is prepared by the Librarian tool and mounted at /librarian.
	reqPath := filepath.Join(librarianDir, "generate-request.json")
	slog.Info("reading generate request", "path", reqPath)

	reqFile, err := os.ReadFile(reqPath)
	if err != nil {
		return fmt.Errorf("failed to read generate-request.json from %s: %w", reqPath, err)
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
	// The API's source directory is relative to the /source mount.
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

	importPath, err := gapicImportPath(api.Path)
	if err != nil {
		return err
	}

	// Construct the protoc command arguments.
	args := []string{
		"--experimental_allow_proto3_optional",
		// All generated files are written to the /output directory.
		"--go_out=" + outputDir,
		"--go-gapic_out=" + outputDir,
		"--go-gapic_opt=go-gapic-package=" + importPath,
		// The -I flag specifies the import path for protoc. All protos
		// and their dependencies must be findable from this path.
		// The /source mount contains the complete googleapis repository.
		"-I=" + sourceDir,
		"-I=" + inputDir + "/my_extra_protos", // TODO: Are there additional protos needed in Go?
	}
	if api.ServiceConfig != "" {
		args = append(args, "--go-gapic_opt=api-service-config="+filepath.Join(apiSourceDir, api.ServiceConfig))
	}

	// TODO: Other potential gapic options: Where do we source these from?
	// TODO: Is this list complete? Are there other `go_gapic_opt` options needed for gapic-generator-go?
	// "--go_gapic_opt", fmt.Sprintf("api-service-config=%s", filepath.Join(uc.serviceDir, uc.serviceYaml))
	// "--go_gapic_opt", fmt.Sprintf("grpc-service-config=%s", filepath.Join(uc.serviceDir, uc.grpcServiceConfig))
	// "--go_gapic_opt", fmt.Sprintf("transport=%s", uc.transport)
	// "--go_gapic_opt", fmt.Sprintf("release-level=%s", uc.releaseLevel)
	// "--go_gapic_opt", "metadata"
	// "--go_gapic_opt", "diregapic"
	// "--go_gapic_opt", "rest-numeric-enums"

	args = append(args, protoFiles...)

	cmd := exec.CommandContext(ctx, "protoc", args...)
	return runCommand(cmd)
}

// gapicImportPath determines the Go import path for a generated GAPIC client
// from a given API proto path.
// E.g., "google/cloud/asset/v1" -> "cloud.google.com/go/asset/apiv1", nil
func gapicImportPath(apiPath string) (string, error) {
	pathParts := strings.Split(apiPath, "/")
	if len(pathParts) < 2 {
		return "", fmt.Errorf("cannot determine service and version from api.Path: %q", apiPath)
	}
	serviceName := pathParts[len(pathParts)-2]
	serviceVersion := pathParts[len(pathParts)-1]
	goVersion := "api" + strings.TrimPrefix(serviceVersion, "v")
	// This base path should ideally be configurable, but this is a reasonable default.
	baseImportPath := "cloud.google.com/go"
	importPath := filepath.Join(baseImportPath, serviceName, goVersion)
	slog.Info("derived go import path", "path", importPath)
	return importPath, nil
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
// It is unmarshalled from the generate-request.json file.
type Library struct {
	ID      string `json:"id"`
	Version string `json:"version,omitempty"`
	APIs    []API  `json:"apis"`
}

// API corresponds to a single API definition within a library.
type API struct {
	Path          string `json:"path"`
	ServiceConfig string `json:"service_config"`
}
