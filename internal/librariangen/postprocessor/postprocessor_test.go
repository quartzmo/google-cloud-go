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
	"os"
	"path/filepath"
	"testing"

	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/protoc"
	"cloud.google.com/go/internal/postprocessor/librarian/librariangen/request"
)

var (
	protocRun = protoc.Run
)

func TestPostProcess(t *testing.T) {
	// Override dependencies with fakes.
	protocRun = func(ctx context.Context, args []string, dir string) error {
		return nil
	}

	tests := []struct {
		name              string
		newModule         bool
		wantFilesCreated  []string
		wantFilesNotCreated []string
	}{
		{
			name:      "new module",
			newModule: true,
			wantFilesCreated: []string{
				"go.mod",
				"CHANGES.md",
				"internal/version.go",
				"README.md",
			},
			wantFilesNotCreated: []string{},
		},
		{
			name:      "existing module",
			newModule: false,
			wantFilesCreated: []string{
				"README.md",
			},
			wantFilesNotCreated: []string{
				"go.mod",
				"CHANGES.md",
				"internal/version.go",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "postprocessor-test")
			if err != nil {
				t.Fatalf("failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			req := &request.Request{
				ID: "google-cloud-chronicle",
				APIs: []request.API{
					{Path: "google/cloud/chronicle/v1"},
				},
			}

			if err := PostProcess(context.Background(), req, tmpDir, tt.newModule); err != nil {
				t.Fatalf("PostProcess() error = %v", err)
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
