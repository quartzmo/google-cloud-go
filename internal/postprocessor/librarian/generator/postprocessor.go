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
	_ "embed"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"

	"cloud.google.com/go/internal/postprocessor/librarian/generator/bazel"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/generator/request"
)

//go:embed _README.md.txt
var readmeTmpl string

// postProcess is the entrypoint for post-processing generated files.
// It runs formatters and other tools to ensure code quality.
func postProcess(ctx context.Context, req *request.Request, bazelConfig *bazel.Config) error {
	slog.Info("starting post-processing", "directory", outputDir)

	if err := goimports(ctx); err != nil {
		// Log a warning instead of failing, as goimports might not be critical.
		slog.Warn("goimports failed, continuing without it", "error", err)
	}

	// To run more advanced tools like staticcheck, we need a go.mod file.
	// Let's try to initialize one.
	if len(req.APIs) > 0 {
		// We'll use the import path of the first API to initialize the module.
		// This assumes all APIs in the request belong to the same module.
		// TODO: Ensure the root module path is used here.
		importPath := bazelConfig.GAPICImportPath

		if err := generateReadmeAndChanges(*outputDir, importPath, req.ID); err != nil {
			return fmt.Errorf("failed to generate README/CHANGES.md: %w", err)
		}

		if err := goModInit(ctx, importPath); err != nil {
			return fmt.Errorf("failed to run 'go mod init': %w", err)
		}

		if err := goModTidy(ctx); err != nil {
			return fmt.Errorf("failed to run 'go mod tidy': %w", err)
		}

		if err := staticcheck(ctx); err != nil {
			// Also a warning, as it might be too strict for generated code.
			slog.Warn("staticcheck failed, continuing without it", "error", err)
		}
	}

	slog.Info("post-processing finished successfully")
	return nil
}

// goimports runs the goimports tool on a directory to format Go files and
// manage imports.
func goimports(ctx context.Context) error {
	slog.Info("running goimports", "directory", *outputDir)
	// The `.` argument will make goimports process all go files in the directory
	// and its subdirectories. The -w flag writes results back to source files.
	args := []string{"goimports", "-w", "."}
	return protoc.Run(ctx, args, *outputDir)
}

// goModInit initializes a go.mod file in the given directory.
func goModInit(ctx context.Context, importPath string) error {
	slog.Info("running go mod init", "directory", *outputDir, "importPath", importPath)
	args := []string{"go", "mod", "init", importPath}
	return protoc.Run(ctx, args, *outputDir)
}

// goModTidy tidies the go.mod file, adding missing and removing unused dependencies.
func goModTidy(ctx context.Context) error {
	slog.Info("running go mod tidy", "directory", *outputDir)
	args := []string{"go", "mod", "tidy"}
	return protoc.Run(ctx, args, *outputDir)
}

// staticcheck runs the staticcheck linter on the code in a directory.
func staticcheck(ctx context.Context) error {
	slog.Info("running staticcheck", "directory", *outputDir)
	// ./... checks all packages in the current directory and subdirectories.
	args := []string{"staticcheck", "./..."}
	return protoc.Run(ctx, args, *outputDir)
}

// generateReadmeAndChanges creates a README.md and CHANGES.md file for a new module.
func generateReadmeAndChanges(path, importPath, apiName string) error {
	readmePath := filepath.Join(path, "README.md")
	slog.Info("creating file", "path", readmePath)
	readmeFile, err := os.Create(readmePath)
	if err != nil {
		return err
	}
	defer readmeFile.Close()
	t := template.Must(template.New("readme").Parse(readmeTmpl))
	readmeData := struct {
		Name       string
		ImportPath string
	}{
		Name:       apiName,
		ImportPath: importPath,
	}
	if err := t.Execute(readmeFile, readmeData); err != nil {
		return err
	}

	changesPath := filepath.Join(path, "CHANGES.md")
	slog.Info("creating file", "path", changesPath)
	changesFile, err := os.Create(changesPath)
	if err != nil {
		return err
	}
	defer changesFile.Close()
	_, err = changesFile.WriteString("# Changes\n")
	return err
}
