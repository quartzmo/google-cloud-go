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
	"errors"
	"os"
	"path/filepath"
	"testing"

	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/request"
)

func TestPostProcess(t *testing.T) {
	tests := []struct {
		name                string
		newModule           bool
		mockProtocRun       func(ctx context.Context, args []string, dir string) error
		wantFilesCreated    []string
		wantFilesNotCreated []string
		wantErr             bool
	}{
		{
			name:      "new module success",
			newModule: true,
			mockProtocRun: func(ctx context.Context, args []string, dir string) error {
				return nil
			},
			wantFilesCreated: []string{
				// "go.mod", // Not actually created by unit tests, verify go mod init called instead.
				"CHANGES.md",
				"internal/version.go",
				"README.md",
				"v1/version.go",
				"v2/version.go",
			},
			wantErr: false,
		},
		{
			name:      "existing module success",
			newModule: false,
			mockProtocRun: func(ctx context.Context, args []string, dir string) error {
				return nil
			},
			wantFilesCreated: []string{
				"README.md",
				"v1/version.go",
				"v2/version.go",
			},
			wantFilesNotCreated: []string{
				"go.mod",
				"CHANGES.md",
				"internal/version.go",
			},
			wantErr: false,
		},
		{
			name:      "goimports fails (non-fatal)",
			newModule: false,
			mockProtocRun: func(ctx context.Context, args []string, dir string) error {
				if args[0] == "goimports" {
					return errors.New("goimports failed")
				}
				return nil
			},
			wantFilesCreated: []string{
				"README.md",
				"v1/version.go",
				"v2/version.go",
			},
			wantErr: false, // goimports error is logged but not returned
		},
		{
			name:      "go mod init fails (fatal)",
			newModule: true,
			mockProtocRun: func(ctx context.Context, args []string, dir string) error {
				if args[0] == "go" && args[1] == "mod" && args[2] == "init" {
					return errors.New("go mod init failed")
				}
				return nil
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "postprocessor-test")
			if err != nil {
				t.Fatalf("failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			protocRun = tt.mockProtocRun

			req := &request.Request{
				ID: "google-cloud-chronicle",
				APIs: []request.API{
					{Path: "google/cloud/chronicle/v1"},
					{Path: "google/cloud/chronicle/v2"},
				},
			}

			if err := PostProcess(context.Background(), req, tmpDir, tt.newModule); (err != nil) != tt.wantErr {
				t.Fatalf("PostProcess() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			for _, file := range tt.wantFilesCreated {
				if _, err := os.Stat(filepath.Join(tmpDir, file)); os.IsNotExist(err) {
					t.Errorf("file %s was not created", file)
				}
			}

			for _, file := range tt.wantFilesNotCreated {
				if _, err := os.Stat(filepath.Join(tmpDir, file)); !os.IsNotExist(err) {
					t.Errorf("file %s was created, but should not have been", file)
				}
			}
		})
	}
}
