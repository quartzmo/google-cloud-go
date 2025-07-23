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
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"cloud.google.com/go/internal/postprocessor/librarian/generator/bazel"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/request"
)

var (
	// The following flags define the directory layout that this generator expects.
	// They default to the paths specified in the Librarian container contract
	// but can be overridden for local development.
	librarianDir = flag.String("librarian", "/librarian", "Path to the librarian-tool input directory. Contains generate-request.json.")
	inputDir     = flag.String("input", "/input", "Path to the .librarian/generator-input directory from the language repository.")
	outputDir    = flag.String("output", "/output", "Path to the empty directory where the generator writes its output.")
	sourceDir    = flag.String("source", "/source", "Path to a complete checkout of the googleapis repository.")
)

// main is the entrypoint for the generator container.
func main() {
	flag.Parse()
	slog.Info("Go generator invoked", "args", os.Args)
	if err := run(context.Background()); err != nil {
		slog.Error("generator failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Go generator finished successfully")
}

// run executes the appropriate command based on the container's invocation arguments.
// The first non-flag argument is the command (e.g., "generate").
func run(ctx context.Context) error {
	args := flag.Args()
	if len(args) < 1 {
		return fmt.Errorf("expected at least one argument for the command, got %d", len(args))
	}
	cmd := args[0]
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
	reqPath := filepath.Join(*librarianDir, "generate-request.json")
	slog.Info("reading generate request", "path", reqPath)

	generateReq, err := request.Parse(reqPath)
	if err != nil {
		return err
	}
	slog.Info("successfully unmarshalled request", "library_id", generateReq.ID)

	for _, api := range generateReq.APIs {
		apiServiceDir := filepath.Join(*sourceDir, api.Path)
		slog.Info("processing api", "service_dir", apiServiceDir)
		bazelConfig, err := bazel.Parse(apiServiceDir)
		if err != nil {
			return fmt.Errorf("failed to parse BUILD.bazel for %s: %w", apiServiceDir, err)
		}
		slog.Info("bazel config loaded", "conf", fmt.Sprintf("%+v", bazelConfig))
		args, err := protoc.Build(generateReq, &api, apiServiceDir, bazelConfig, *sourceDir, *outputDir)
		if err != nil {
			return fmt.Errorf("failed to build protoc command for api %q in library %q: %w", api.Path, generateReq.ID, err)
		}
		if err := protoc.Run(ctx, args, *outputDir); err != nil {
			return fmt.Errorf("protoc failed for api %q in library %q: %w", api.Path, generateReq.ID, err)
		}
	}

	// TODO(quartzmo): Implement post-processing.

	slog.Info("generate command finished")
	return nil
}
