// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bazelbuild/buildtools/build"
	"github.com/ghodss/yaml"
	"google.golang.org/genproto/googleapis/api/serviceconfig"
	"google.golang.org/protobuf/encoding/protojson"
)

type BazelConfig struct {
	// serviceConfig is the YAML file in the API version directory in googleapis. TODO: redundant with serviceYaml, below?
	// E.g., googleapis/google/cloud/asset/v1/cloudasset_v1.yaml
	serviceConfig *serviceconfig.Service
	serviceDir    string
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
	// serviceYaml is the YAML file in the API version directory in googleapis. TODO: redundant with serviceConfig, above?
	// E.g., googleapis/google/cloud/asset/v1/cloudasset_v1.yaml
	serviceYaml string
	// transport is typically one of "grpc", "rest" or "grpc+rest".
	transport string
	// Not typically present.
	diregapic bool
}

// NewBazelConfig loads legacy configuration from 
func NewBazelConfig(serviceDir string) (*BazelConfig, error) {
	uc := &BazelConfig{
		serviceDir: serviceDir,
	}
	if err := uc.loadBazelRules(serviceDir); err != nil {
		return nil, fmt.Errorf("failed to load bazel rules: %w", err)
	}
	if err := uc.loadServiceConfig(serviceDir); err != nil {
		return nil, fmt.Errorf("failed to load service config: %w", err)
	}

	slog.Info("unified config loaded", "conf", fmt.Sprintf("%+v", uc))
	return uc, nil
}

func (c *BazelConfig) loadBazelRules(dir string) error {
	fp := filepath.Join(dir, "BUILD.bazel")
	data, err := os.ReadFile(fp)
	if err != nil {
		return fmt.Errorf("failed to read BUILD.bazel file %s: %w", fp, err)
	}
	f, err := build.Parse(fp, data)
	if err != nil {
		return fmt.Errorf("failed to parse BUILD.bazel file %s: %w", fp, err)
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
			c.serviceYaml = v
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
		}
	}
	return nil
}

func (c *BazelConfig) loadServiceConfig(dir string) error {
	if c.serviceYaml == "" {
		return fmt.Errorf("service config is empty")
	}
	fp := filepath.Join(dir, c.serviceYaml)
	yb, err := os.ReadFile(fp)
	if err != nil {
		return fmt.Errorf("failed to read service config file %s: %w", fp, err)
	}
	jb, err := yaml.YAMLToJSON(yb)
	if err != nil {
		return fmt.Errorf("failed to convert yaml to json for %s: %w", fp, err)
	}
	sc := &serviceconfig.Service{}
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(jb, sc); err != nil {
		return fmt.Errorf("failed to unmarshal service config file %s: %w", fp, err)
	}
	c.serviceConfig = sc
	return nil
}
