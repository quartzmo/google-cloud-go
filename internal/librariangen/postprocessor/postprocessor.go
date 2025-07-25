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
	"strings"

	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/request"
)

//go:embed _README.md.txt
var readmeTmpl string

// PostProcess is the entrypoint for post-processing generated files.
// It runs formatters and other tools to ensure code quality.
func PostProcess(ctx context.Context, req *request.Request, moduleDir string, newModule bool) error {
	slog.Info("starting post-processing", "directory", moduleDir, "new_module", newModule)

	if err := goimports(ctx, moduleDir); err != nil {
		slog.Warn("goimports failed, continuing without it", "error", err)
	}

	if len(req.APIs) == 0 {
		slog.Info("no APIs in request, skipping module initialization")
		return nil
	}

	// E.g. google-cloud-chronicle -> chronicle
	moduleName := strings.TrimPrefix(req.ID, "google-cloud-")
	shortModulePath := "cloud.google.com/go/" + moduleName
	// E.g. google-cloud-chronicle -> Chronicle API
	friendlyAPIName := strings.Title(strings.Replace(moduleName, "-", " ", -1)) + " API"

	if newModule {
		slog.Info("initializing new module")
		if err := goModInit(ctx, shortModulePath, moduleDir); err != nil {
			return fmt.Errorf("failed to run 'go mod init': %w", err)
		}
		if err := generateChanges(moduleDir); err != nil {
			return fmt.Errorf("failed to generate CHANGES.md: %w", err)
		}
		if err := generateInternalVersionFile(moduleDir); err != nil {
			return fmt.Errorf("failed to generate internal/version.go: %w", err)
		}
	}

	// The README should be updated on every run.
	if err := generateReadme(moduleDir, shortModulePath, friendlyAPIName); err != nil {
		return fmt.Errorf("failed to generate README.md: %w", err)
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

// generateReadme creates a README.md file for a new module.
func generateReadme(path, modulePath, apiName string) error {
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
	return t.Execute(readmeFile, readmeData)
}

// generateChanges creates a CHANGES.md file for a new module.
func generateChanges(moduleDir string) error {
	changesPath := filepath.Join(moduleDir, "CHANGES.md")
	slog.Info("creating file", "path", changesPath)
	content := "# Changes\n"
	return os.WriteFile(changesPath, []byte(content), 0644)
}

// generateInternalVersionFile creates an internal/version.go file for a new module.
func generateInternalVersionFile(moduleDir string) error {
	internalDir := filepath.Join(moduleDir, "internal")
	if err := os.MkdirAll(internalDir, 0755); err != nil {
		return err
	}
	versionPath := filepath.Join(internalDir, "version.go")
	slog.Info("creating file", "path", versionPath)
	content := `package internal

const Version = "0.0.1"
`
	return os.WriteFile(versionPath, []byte(content), 0644)
}
