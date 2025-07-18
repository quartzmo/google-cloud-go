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

package bazel

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name    string
		content string
		want    *Config
		wantErr bool
	}{
		{
			name: "valid build file",
			content: `
go_proto_library(
    name = "asset_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "cloud.google.com/go/asset/apiv1/assetpb",
    protos = [":asset_proto"],
)

go_gapic_library(
    name = "asset_go_gapic",
    srcs = [":asset_proto_with_info"],
    grpc_service_config = "cloudasset_grpc_service_config.json",
    importpath = "cloud.google.com/go/asset/apiv1;asset",
    metadata = True,
    release_level = "ga",
    rest_numeric_enums = True,
    service_yaml = "cloudasset_v1.yaml",
    transport = "grpc+rest",
    diregapic = True,
)
`,
			want: &Config{
				ProtoImportPath:   "cloud.google.com/go/asset/apiv1/assetpb",
				GRPCServiceConfig: "cloudasset_grpc_service_config.json",
				GAPICImportPath:   "cloud.google.com/go/asset/apiv1;asset",
				Metadata:          true,
				ReleaseLevel:      "ga",
				RESTNumericEnums:  true,
				ServiceYAML:       "cloudasset_v1.yaml",
				Transport:         "grpc+rest",
				Diregapic:         true,
			},
			wantErr: false,
		},
		{
			name:    "malformed build file",
			content: `go_proto_library(`,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			buildPath := filepath.Join(tmpDir, "BUILD.bazel")
			if err := os.WriteFile(buildPath, []byte(tc.content), 0644); err != nil {
				t.Fatalf("failed to write test file: %v", err)
			}

			got, err := Parse(tmpDir)

			if (err != nil) != tc.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestParse_FileNotFound(t *testing.T) {
	_, err := Parse("non-existent-dir")
	if err == nil {
		t.Error("Parse() expected error for non-existent file, got nil")
	}
}