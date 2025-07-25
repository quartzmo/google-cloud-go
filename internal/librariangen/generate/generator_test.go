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
	"os"
	"path/filepath"
	"testing"

	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/request"
)

func TestFixPermissions(t *testing.T) {
	// Create a temporary directory for the test.
	tmpDir, err := os.MkdirTemp("", "permissions-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a subdirectory to test recursive behavior.
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	// Create test files with incorrect permissions.
	goFile1 := filepath.Join(tmpDir, "file1.go")
	goFile2 := filepath.Join(subDir, "file2.go")
	otherFile := filepath.Join(tmpDir, "other.txt")

	if err := os.WriteFile(goFile1, []byte("package main"), 0777); err != nil {
		t.Fatalf("failed to write goFile1: %v", err)
	}
	if err := os.WriteFile(goFile2, []byte("package main"), 0777); err != nil {
		t.Fatalf("failed to write goFile2: %v", err)
	}
	if err := os.WriteFile(otherFile, []byte("some text"), 0777); err != nil {
		t.Fatalf("failed to write otherFile: %v", err)
	}

	// Run the function to fix permissions.
	if err := fixPermissions(tmpDir); err != nil {
		t.Fatalf("fixPermissions() failed: %v", err)
	}

	// Check permissions of the .go files.
	for _, f := range []string{goFile1, goFile2} {
		info, err := os.Stat(f)
		if err != nil {
			t.Fatalf("failed to stat %s: %v", f, err)
		}
		if info.Mode().Perm() != 0644 {
			t.Errorf("permission of %s is %v, want 0644", f, info.Mode().Perm())
		}
	}

	// Check that the permission of the other file is unchanged.
	info, err := os.Stat(otherFile)
	if err != nil {
		t.Fatalf("failed to stat %s: %v", otherFile, err)
	}
	if info.Mode().Perm() == 0644 {
		t.Errorf("permission of %s was changed, should not have been", otherFile)
	}
}

func TestFlattenOutput(t *testing.T) {
	// Create a temporary directory for the test.
	tmpDir, err := os.MkdirTemp("", "flatten-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the nested directory structure.
	goDir := filepath.Join(tmpDir, "cloud.google.com", "go")
	if err := os.MkdirAll(goDir, 0755); err != nil {
		t.Fatalf("failed to create goDir: %v", err)
	}

	// Create a file to be moved.
	filePath := filepath.Join(goDir, "file.txt")
	if err := os.WriteFile(filePath, []byte("test"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	// Run the flatten function.
	if err := flattenOutput(tmpDir); err != nil {
		t.Fatalf("flattenOutput() failed: %v", err)
	}

	// Check that the file was moved to the top level.
	newFilePath := filepath.Join(tmpDir, "file.txt")
	if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
		t.Errorf("file was not moved to the top level")
	}

	// Check that the old directory was removed.
	if _, err := os.Stat(filepath.Join(tmpDir, "cloud.google.com")); !os.IsNotExist(err) {
		t.Errorf("old directory was not removed")
	}
}

func TestGenerate(t *testing.T) {
	// Create a temporary directory for the test.
	tmpDir, err := os.MkdirTemp("", "generator-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create fake directories.
	librarianDir := filepath.Join(tmpDir, "librarian")
	sourceDir := filepath.Join(tmpDir, "source")
	outputDir := filepath.Join(tmpDir, "output")
	if err := os.MkdirAll(librarianDir, 0755); err != nil {
		t.Fatalf("failed to create librarian dir: %v", err)
	}
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Fatalf("failed to create source dir: %v", err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("failed to create output dir: %v", err)
	}

	// Create a fake generate-request.json.
	reqFile, err := os.Create(filepath.Join(librarianDir, "generate-request.json"))
	if err != nil {
		t.Fatalf("failed to create fake request file: %v", err)
	}
	fmt.Fprintln(reqFile, `{"id": "foo", "apis": [{"path": "api/v1"}]}`)
	reqFile.Close()

	// Create a fake BUILD.bazel file.
	apiDir := filepath.Join(sourceDir, "api/v1")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatalf("failed to create api dir: %v", err)
	}
	// Create a fake .proto file.
	protoFile, err := os.Create(filepath.Join(apiDir, "fake.proto"))
	if err != nil {
		t.Fatalf("failed to create fake proto file: %v", err)
	}
	protoFile.Close()
	bazelFile, err := os.Create(filepath.Join(apiDir, "BUILD.bazel"))
	if err != nil {
		t.Fatalf("failed to create fake bazel file: %v", err)
	}
	fmt.Fprint(bazelFile, `
go_gapic_library(
    name = "v1_gapic",
    importpath = "path/to/v1;v1",
)
`)
	bazelFile.Close()

	// Override dependencies with fakes.
	postProcess = func(ctx context.Context, req *request.Request, moduleDir string, newModule bool) error {
		return nil
	}
	protocRun = func(ctx context.Context, args []string, dir string) error {
		return nil
	}
	// We can use the real bazel.Parse and request.Parse because we created the
	// necessary files.

	cfg := &Config{
		LibrarianDir: librarianDir,
		InputDir:     "fake-input",
		OutputDir:    outputDir,
		SourceDir:    sourceDir,
	}

	if err := Generate(context.Background(), cfg); err != nil {
		t.Errorf("Generate() error = %v, wantErr %v", err, false)
	}
}


func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantErr bool
	}{
		{
			name: "valid",
			cfg: &Config{
				LibrarianDir: "a",
				InputDir:     "b",
				OutputDir:    "c",
				SourceDir:    "d",
			},
			wantErr: false,
		},
		{
			name: "missing librarian dir",
			cfg: &Config{
				InputDir:  "b",
				OutputDir: "c",
				SourceDir: "d",
			},
			wantErr: true,
		},
		{
			name: "missing input dir",
			cfg: &Config{
				LibrarianDir: "a",
				OutputDir:    "c",
				SourceDir:    "d",
			},
			wantErr: true,
		},
		{
			name: "missing output dir",
			cfg: &Config{
				LibrarianDir: "a",
				InputDir:     "b",
				SourceDir:    "d",
			},
			wantErr: true,
		},
		{
			name: "missing source dir",
			cfg: &Config{
				LibrarianDir: "a",
				InputDir:     "b",
				OutputDir:    "c",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
