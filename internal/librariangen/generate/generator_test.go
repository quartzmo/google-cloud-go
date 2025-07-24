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
	postProcess = func(ctx context.Context, req *request.Request, modulePath, outputDir string) error {
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
