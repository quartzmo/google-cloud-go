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

package postprocessor

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"

	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/request"
)

//go:embed _README.md.txt
var readmeTmpl string

// postProcess is the entrypoint for post-processing generated files.
// It runs formatters and other tools to ensure code quality.
func PostProcess(ctx context.Context, req *request.Request, modulePath, moduleDir string) error {
	slog.Info("starting post-processing", "directory", moduleDir)

	if err := goimports(ctx, moduleDir); err != nil {
		// Log a warning instead of failing, as goimports might not be critical.
		slog.Warn("goimports failed, continuing without it", "error", err)
	}

	// To run more advanced tools like staticcheck, we need a go.mod file.
	// Let's try to initialize one.
	if len(req.APIs) > 0 {

		if err := generateReadmeAndChanges(moduleDir, modulePath, req.ID); err != nil {
			return fmt.Errorf("failed to generate README/CHANGES.md: %w", err)
		}

		if err := goModInit(ctx, modulePath, moduleDir); err != nil {
			return fmt.Errorf("failed to run 'go mod init': %w", err)
		}

		if err := goModTidy(ctx, moduleDir); err != nil {
			return fmt.Errorf("failed to run 'go mod tidy': %w", err)
		}

		if err := staticcheck(ctx, moduleDir); err != nil {
			// Also a warning, as it might be too strict for generated code.
			slog.Warn("staticcheck failed, continuing without it", "error", err)
		}
	}

	slog.Info("post-processing finished successfully")
	return nil
}

// goimports runs the goimports tool on a directory to format Go files and
// manage imports.
func goimports(ctx context.Context, dir string) error {
	slog.Info("running goimports", "directory", dir)
	// The `.` argument will make goimports process all go files in the directory
	// and its subdirectories. The -w flag writes results back to source files.
	args := []string{"goimports", "-w", "."}
	return protoc.Run(ctx, args, dir)
}

// goModInit initializes a go.mod file in the given directory.
func goModInit(ctx context.Context, modulePath, dir string) error {
	slog.Info("running go mod init", "directory", dir, "modulePath", modulePath)
	args := []string{"go", "mod", "init", modulePath}
	return protoc.Run(ctx, args, dir)
}

// goModTidy tidies the go.mod file, adding missing and removing unused dependencies.
func goModTidy(ctx context.Context, dir string) error {
	slog.Info("running go mod tidy", "directory", dir)
	args := []string{"go", "mod", "tidy"}
	return protoc.Run(ctx, args, dir)
}

// staticcheck runs the staticcheck linter on the code in a directory.
func staticcheck(ctx context.Context, dir string) error {
	slog.Info("running staticcheck", "directory", dir)
	// ./... checks all packages in the current directory and subdirectories.
	args := []string{"staticcheck", "./..."}
	return protoc.Run(ctx, args, dir)
}

// generateReadmeAndChanges creates a README.md and CHANGES.md file for a new module.
func generateReadmeAndChanges(path, modulePath, apiName string) error {
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
		ModulePath string
	}{
		Name:       apiName,
		ModulePath: modulePath,
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
