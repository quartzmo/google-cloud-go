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
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/bazel"
	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/postprocessor"
	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/request"
)

var (
	postProcess  = postprocessor.PostProcess
	bazelParse   = bazel.Parse
	protocRun    = protoc.Run
	requestParse = request.Parse
)

// Config holds the configuration for the generate command.
type Config struct {
	LibrarianDir        string
	InputDir            string
	OutputDir           string
	SourceDir           string
	EnablePostProcessor bool
}

// Validate ensures that the configuration is valid.
func (c *Config) Validate() error {
	if c.LibrarianDir == "" {
		return fmt.Errorf("librarian directory must be set")
	}
	if c.InputDir == "" {
		return fmt.Errorf("input directory must be set")
	}
	if c.OutputDir == "" {
		return fmt.Errorf("output directory must be set")
	}
	if c.SourceDir == "" {
		return fmt.Errorf("source directory must be set")
	}
	return nil
}

// Generate handles the logic for the 'generate' CLI command.
// It reads a request file, and for each API specified, it invokes protoc
// to generate the client library.
func Generate(ctx context.Context, cfg *Config) error {
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}
	modulePath, err := handleGapicgen(ctx, cfg)
	if err != nil {
		return fmt.Errorf("gapic generation failed: %w", err)
	}
	if err := fixPermissions(cfg.OutputDir); err != nil {
		return fmt.Errorf("failed to fix permissions: %w", err)
	}
	slog.Info("using module path from final API", "importpath", modulePath)

	if cfg.EnablePostProcessor {
		slog.Info("post-processor enabled")
		generateReq, err := readGenerateReq(cfg.LibrarianDir)
		if err != nil {
			return err
		}
		if err := postProcess(ctx, generateReq, modulePath, cfg.OutputDir); err != nil {
			return fmt.Errorf("post-processing failed: %w", err)
		}
	}

	slog.Info("generate command finished")
	return nil
}

// handleGapicgen handles the protoc GAPIC generation logic for the 'generate' CLI command.
// It reads a request file, and for each API specified, it invokes protoc
// to generate the client library.
func handleGapicgen(ctx context.Context, cfg *Config) (string, error) {
	slog.Info("generate command started")

	generateReq, err := readGenerateReq(cfg.LibrarianDir)
	if err != nil {
		return "", err
	}
	var bazelConfig *bazel.Config
	for _, api := range generateReq.APIs {
		apiServiceDir := filepath.Join(cfg.SourceDir, api.Path)
		slog.Info("processing api", "service_dir", apiServiceDir)
		var err error
		bazelConfig, err = bazelParse(apiServiceDir)
		if err != nil {
			return "", fmt.Errorf("failed to parse BUILD.bazel for %s: %w", apiServiceDir, err)
		}
		slog.Info("bazel config loaded", "conf", fmt.Sprintf("%+v", bazelConfig))
		args, err := protoc.Build(generateReq, &api, apiServiceDir, bazelConfig, cfg.SourceDir, cfg.OutputDir)
		if err != nil {
			return "", fmt.Errorf("failed to build protoc command for api %q in library %q: %w", api.Path, generateReq.ID, err)
		}
		if err := protocRun(ctx, args, cfg.OutputDir); err != nil {
			return "", fmt.Errorf("protoc failed for api %q in library %q: %w", api.Path, generateReq.ID, err)
		}
	}

	// We'll use the import path of the last API's BUILD.bazel to initialize the module.
	// This assumes all APIs in the request belong to the same module.
	// TODO: Ensure the root module path is used here.
	modulePath := bazelConfig.ModulePath()
	return modulePath, nil
}

// readGenerateReq reads generate-request.json from the librarian-tool input directory.
// The request file tells librariangen which library and APIs to generate.
// It is prepared by the Librarian tool and mounted at /librarian.
func readGenerateReq(librarianDir string) (*request.Request, error) {
	reqPath := filepath.Join(librarianDir, "generate-request.json")
	slog.Info("reading generate request", "path", reqPath)

	generateReq, err := requestParse(reqPath)
	if err != nil {
		return nil, err
	}
	slog.Info("successfully unmarshalled request", "library_id", generateReq.ID)
	return generateReq, nil
}

// fixPermissions recursively finds all .go files in the given directory and sets
// their permissions to 0644.
func fixPermissions(dir string) error {
	slog.Info("fixing file permissions", "dir", dir)
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			slog.Info("fixing file", "path", path)
			if err := os.Chmod(path, 0644); err != nil {
				return fmt.Errorf("failed to chmod %s: %w", path, err)
			}
		}
		return nil
	})
}
