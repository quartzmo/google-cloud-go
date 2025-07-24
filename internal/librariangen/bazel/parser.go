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
	"strings"

	"github.com/bazelbuild/buildtools/build"
)

// Config holds the configuration extracted from a BUILD.bazel file.
type Config struct {
	// The fields below are all from the API version BUILD.bazel file.
	// E.g., googleapis/google/cloud/asset/v1/BUILD.bazel
	// Note that not all fields are present in every rule usage.
	// protoImportPath is importpath in the go_proto_library or go_grpc_library rule.
	protoImportPath string
	// The remaining fields below are all from the go_gapic_library rule.
	grpcServiceConfig string
	// gapicImportPath is importpath in the go_gapic_library rule.
	gapicImportPath string
	// Not typically present.
	metadata bool
	// releaseLevel is typically one of "beta", "" (same as beta) or "ga".
	// If "ga", gapic-generator-go does not print a warning in the package docs.
	releaseLevel string
	// restNumericEnums is typically true.
	restNumericEnums bool
	// serviceYAML is the YAML file in the API version directory in googleapis.
	// E.g., googleapis/google/cloud/asset/v1/cloudasset_v1.yaml
	serviceYAML string
	// transport is typically one of "grpc", "rest" or "grpc+rest".
	transport string
	// Not typically present.
	diregapic bool
	// Set to true if a go_grpc_library rule is found in the BUILD.bazel file.
	hasGoGRPC bool
}

// GAPICImportPath returns the GAPIC import path.
func (c *Config) GAPICImportPath() string { return c.gapicImportPath }

// ModulePath returns the module path from the GAPIC import path.
// E.g., "cloud.google.com/go/chronicle/apiv1;chronicle" -> "cloud.google.com/go/chronicle/apiv1"
func (c *Config) ModulePath() string {
	if idx := strings.Index(c.gapicImportPath, ";"); idx != -1 {
		return c.gapicImportPath[:idx]
	}
	return c.gapicImportPath
}

// ServiceYAML returns the service YAML file name.
func (c *Config) ServiceYAML() string { return c.serviceYAML }

// GRPCServiceConfig returns the gRPC service config file name.
func (c *Config) GRPCServiceConfig() string { return c.grpcServiceConfig }

// Transport returns the transport type.
func (c *Config) Transport() string { return c.transport }

// ReleaseLevel returns the release level.
func (c *Config) ReleaseLevel() string { return c.releaseLevel }

// HasMetadata returns true if metadata should be generated.
func (c *Config) HasMetadata() bool { return c.metadata }

// HasDiregapic returns true if diregapic should be enabled.
func (c *Config) HasDiregapic() bool { return c.diregapic }

// HasRESTNumericEnums returns true if REST numeric enums should be enabled.
func (c *Config) HasRESTNumericEnums() bool { return c.restNumericEnums }

// HasGoGRPC returns true if a go_grpc_library rule was found.
func (c *Config) HasGoGRPC() bool { return c.hasGoGRPC }

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
			c.grpcServiceConfig = v
		}
		if v := rule.AttrString("importpath"); v != "" {
			c.gapicImportPath = v
		}
		if v := rule.AttrLiteral("metadata"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				c.metadata = b
			} else {
				slog.Warn("failed to parse metadata", "error", err, "input", v)
			}
		}
		if v := rule.AttrString("release_level"); v != "" {
			c.releaseLevel = v
		}
		if v := rule.AttrLiteral("rest_numeric_enums"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				c.restNumericEnums = b
			} else {
				slog.Warn("failed to parse rest_numeric_enums", "error", err, "input", v)
			}
		}
		if v := rule.AttrString("service_yaml"); v != "" {
			c.serviceYAML = v
		}
		if v := rule.AttrString("transport"); v != "" {
			c.transport = v
		}
		if v := rule.AttrLiteral("diregapic"); v != "" {
			if b, err := strconv.ParseBool(v); err == nil {
				c.diregapic = b
			} else {
				slog.Warn("failed to parse diregapic", "error", err, "input", v)
			}
		}
	}

	// We are currently migrating go_proto_library to go_grpc_library, so check
	// both for now.
	for _, rule := range f.Rules("go_proto_library") {
		if v := rule.AttrString("importpath"); v != "" {
			c.protoImportPath = v
		}
	}
	for _, rule := range f.Rules("go_grpc_library") {
		if v := rule.AttrString("importpath"); v != "" {
			c.protoImportPath = v
			c.hasGoGRPC = true
		}
	}
	return c, nil
}