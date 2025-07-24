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

package generate

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"path/filepath"

	"cloud.google.com/go/internal/postprocessor/librarian/generator/bazel"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/postprocessor"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/request"
)

// generateConfig holds the configuration for the generate command.
type generateConfig struct {
	librarianDir string
	inputDir     string
	outputDir    string
	sourceDir    string
}

// Validate ensures that the configuration is valid.
func (c *generateConfig) validate() error {
	if c.librarianDir == "" {
		return fmt.Errorf("librarian directory must be set")
	}
	if c.inputDir == "" {
		return fmt.Errorf("input directory must be set")
	}
	if c.outputDir == "" {
		return fmt.Errorf("output directory must be set")
	}
	if c.sourceDir == "" {
		return fmt.Errorf("source directory must be set")
	}
	return nil
}

// generate handles the logic for the 'generate' CLI command.
// It reads a request file, and for each API specified, it invokes protoc
// to generate the client library.
func Generate(ctx context.Context) error {
	cfg := generateConfigFromFlag()
	flag.Parse()
	if err := cfg.validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}
	importPath, err := handleGapicgen(ctx, cfg)
	if err != nil {
		return fmt.Errorf("post-processing failed: %w", err)
	}
	generateReq, err := readGenerateReq(cfg.librarianDir)
	if err != nil {
		return fmt.Errorf("reading generate-request.json failed: %w", err)
	}

	// TODO(quartzmo): Implement post-processing.
	if err := postprocessor.PostProcess(ctx, generateReq, importPath, cfg.outputDir); err != nil {
		return fmt.Errorf("post-processing failed: %w", err)
	}

	slog.Info("generate command finished")
	return nil
}

// NewGenerateConfig creates a new generateConfig from command-line flags.
func generateConfigFromFlag() *generateConfig {
	return &generateConfig{
		librarianDir: *flag.String("librarian", "/librarian", "Path to the librarian-tool input directory. Contains generate-request.json."),
		inputDir:     *flag.String("input", "/input", "Path to the .librarian/generator-input directory from the language repository."),
		outputDir:    *flag.String("output", "/output", "Path to the empty directory where librariangen writes its output."),
		sourceDir:    *flag.String("source", "/source", "Path to a complete checkout of the googleapis repository."),
	}
}

// handleGapicgen handles the protoc GAPIC generation logic for the 'generate' CLI command.
// It reads a request file, and for each API specified, it invokes protoc
// to generate the client library.
func handleGapicgen(ctx context.Context, cfg *generateConfig) (string, error) {
	slog.Info("generate command started")

	// The request file tells librariangen which library and APIs to generate.
	// It is prepared by the Librarian tool and mounted at /librarian.
	reqPath := filepath.Join(cfg.librarianDir, "generate-request.json")
	slog.Info("reading generate request", "path", reqPath)

	generateReq, err := request.Parse(reqPath)
	if err != nil {
		return "", err
	}
	slog.Info("successfully unmarshalled request", "library_id", generateReq.ID)

	var bazelConfig *bazel.Config
	for _, api := range generateReq.APIs {
		apiServiceDir := filepath.Join(cfg.sourceDir, api.Path)
		slog.Info("processing api", "service_dir", apiServiceDir)
		var err error
		bazelConfig, err = bazel.Parse(apiServiceDir)
		if err != nil {
			return "", fmt.Errorf("failed to parse BUILD.bazel for %s: %w", apiServiceDir, err)
		}
		slog.Info("bazel config loaded", "conf", fmt.Sprintf("%+v", bazelConfig))
		args, err := protoc.Build(generateReq, &api, apiServiceDir, bazelConfig, cfg.sourceDir, cfg.outputDir)
		if err != nil {
			return "", fmt.Errorf("failed to build protoc command for api %q in library %q: %w", api.Path, generateReq.ID, err)
		}
		if err := protoc.Run(ctx, args, cfg.outputDir); err != nil {
			return "", fmt.Errorf("protoc failed for api %q in library %q: %w", api.Path, generateReq.ID, err)
		}
	}

	// We'll use the import path of the last API's BUILD.bazel to initialize the module.
	// This assumes all APIs in the request belong to the same module.
	// TODO: Ensure the root module path is used here.
	importPath := bazelConfig.GAPICImportPath()
	return importPath, nil
}

// readGenerateReq reads generate-request.json from the librarian-tool input directory.
func readGenerateReq(librarianDir string) (*request.Request, error) {
	slog.Info("generate command started")

	// The request file tells librariangen which library and APIs to generate.
	// It is prepared by the Librarian tool and mounted at /librarian.
	reqPath := filepath.Join(librarianDir, "generate-request.json")
	slog.Info("reading generate request", "path", reqPath)

	generateReq, err := request.Parse(reqPath)
	if err != nil {
		return nil, err
	}
	slog.Info("successfully unmarshalled request", "library_id", generateReq.ID)
	return generateReq, nil
}
