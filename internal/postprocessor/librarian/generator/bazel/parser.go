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
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bazelbuild/buildtools/build"
)

// Config holds the configuration extracted from a BUILD.bazel file.
type Config struct {
	// The fields below are all from the API version BUILD.bazel file.
	// E.g., googleapis/google/cloud/asset/v1/BUILD.bazel
	// Note that not all fields are present in every rule usage.
	// ProtoImportPath is importpath in the go_proto_library or go_grpc_library rule.
	ProtoImportPath string
	// The remaining fields below are all from the go_gapic_library rule.
	GRPCServiceConfig string
	// GAPICImportPath is importpath in the go_gapic_library rule.
	GAPICImportPath string
	// Not typically present.
	Metadata bool
	// ReleaseLevel is typically one of "beta", "" (same as beta) or "ga".
	// If "ga", gapic-generator-go does not print a warning in the package docs.
	ReleaseLevel string
	// RESTNumericEnums is typically true.
	RESTNumericEnums bool
	// ServiceYAML is the YAML file in the API version directory in googleapis.
	// E.g., googleapis/google/cloud/asset/v1/cloudasset_v1.yaml
	ServiceYAML string
	// Transport is typically one of "grpc", "rest" or "grpc+rest".
	Transport string
	// Not typically present.
	Diregapic bool
}

// Parse reads a BUILD.bazel file from the given directory and extracts the
// relevant configuration from the go_gapic_library and go_proto_library rules.
func Parse(dir string) (*Config, error) {
	c := &Config{}
	fp := filepath.Join(dir, "BUILD.bazel")
	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("failed to read BUILD.bazel file %s: %w", fp, err)
	}
	f, err := build.Parse(fp, data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse BUILD.bazel file %s: %w", fp, err)
	}

	// GAPIC build target
	for _, rule := range f.Rules("go_gapic_library") {
		if v := rule.AttrString("grpc_service_config"); v != "" {
			c.GRPCServiceConfig = v
		}
		if v := rule.AttrString("importpath"); v != "" {
			c.GAPICImportPath = v
		}
		if v := rule.AttrLiteral("metadata"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				c.Metadata = b
			} else {
				slog.Warn("failed to parse metadata", "error", err, "input", v)
			}
		}
		if v := rule.AttrString("release_level"); v != "" {
			c.ReleaseLevel = v
		}
		if v := rule.AttrLiteral("rest_numeric_enums"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				c.RESTNumericEnums = b
			} else {
				slog.Warn("failed to parse rest_numeric_enums", "error", err, "input", v)
			}
		}
		if v := rule.AttrString("service_yaml"); v != "" {
			c.ServiceYAML = v
		}
		if v := rule.AttrString("transport"); v != "" {
			c.Transport = v
		}
		if v := rule.AttrLiteral("diregapic"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				c.Diregapic = b
			} else {
				slog.Warn("failed to parse diregapic", "error", err, "input", v)
			}
		}
	}

	// We are currently migrating go_proto_library to go_grpc_library, so check
	// both for now.
	for _, rule := range f.Rules("go_proto_library") {
		if v := rule.AttrString("importpath"); v != "" {
			c.ProtoImportPath = v
		}
	}
	for _, rule := range f.Rules("go_grpc_library") {
		if v := rule.AttrString("importpath"); v != "" {
			c.ProtoImportPath = v
		}
	}
	return c, nil
}
