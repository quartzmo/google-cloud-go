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

// python_librarian.go is an exploratory implementation of the Python Librarian container.
//
// As outlined in go/librarian:python and go/sdk-librarian-python, the Python
// Librarian is a containerized application for generating Python client libraries.
// This script acts as the main entrypoint for that container, orchestrating the
// code generation process based on the contracts defined in go/librarian:cli-reimagined.
//
// This implementation is intended for instructional purposes for new engineers.
// It specifically models the "Phase 1: Owlbot in a box" strategy from the
// migration plan. In this phase, the container wraps the existing Bazel-based
// generation process to ensure a smooth transition with minimal initial code changes.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// The following structs represent the configuration for a library that Librarian
// manages. This structure is defined in the `state.yaml` file within a language
// repository and is used for communication between the Librarian CLI and this
// generator container. See go/librarian:cli-reimagined for more details.

// Library represents a single releasable unit, like a Python package.
type Library struct {
	// A language-specific identifier, e.g., "google-cloud-storage".
	ID string `json:"id"`
	// The last released version, e.g., "1.15.0".
	Version string `json:"version,omitempty"`
	// The commit hash from the API definition repo at the last generation.
	LastGeneratedCommit string `json:"last_generated_commit,omitempty"`
	// APIs bundled in this library. A library can contain multiple APIs.
	APIs []API `json:"apis"`
	// Paths where generated source code is placed.
	SourcePaths []string `json:"source_paths,omitempty"`
	// Regex patterns for files to preserve during generation.
	PreserveRegex []string `json:"preserve_regex,omitempty"`
	// Regex patterns for files/directories to remove before copying new files.
	RemoveRegex []string `json:"remove_regex,omitempty"`
}

// API represents a single API definition, e.g., "google/storage/v1".
type API struct {
	// The path to the API definition relative to the root of the API repo.
	Path string `json:"path"`
	// The service config file name for this API.
	ServiceConfig string `json:"service_config"`
}

// main is the entrypoint for the Python Librarian container.
//
// The Librarian CLI invokes this container with a specific command as the first
// argument. This function parses that command and delegates to the appropriate
// handler. This adheres to the container contract where the container is invoked
// with commands like 'configure', 'generate', or 'build'.
func main() {
	log.Println("Python Librarian container started.")

	if len(os.Args) < 2 {
		log.Fatal("Error: Missing command. Usage: python_librarian <command>")
	}

	// The command is passed as the first argument by the Librarian CLI.
	command := os.Args[1]
	log.Printf("Executing command: %s", command)

	switch command {
	case "configure":
		if err := handleConfigure(); err != nil {
			log.Fatalf("Configuration failed: %v", err)
		}
	case "generate":
		if err := handleGenerate(); err != nil {
			log.Fatalf("Generation failed: %v", err)
		}
	case "build":
		if err := handleBuild(); err != nil {
			log.Fatalf("Build failed: %v", err)
		}
	default:
		log.Fatalf("Error: Unknown command '%s'", command)
	}

	log.Printf("Command '%s' completed successfully.", command)
}

// handleConfigure is responsible for the 'configure' step of onboarding a new library.
//
// As per the 'configure container command' contract (go/librarian:guide), this
// function reads a minimal library definition from `/librarian/configure-request.json`,
// enriches it with Python-specific details, and writes the result back to
// `/librarian/configure-response.json`.
func handleConfigure() error {
	log.Println("--- Running Configure Step ---")
	requestPath := "/librarian/configure-request.json"
	responsePath := "/librarian/configure-response.json"

	// 1. Read the request from the Librarian CLI.
	log.Printf("Reading configure request from %s", requestPath)
	data, err := ioutil.ReadFile(requestPath)
	if err != nil {
		return fmt.Errorf("failed to read configure request: %w", err)
	}

	var lib Library
	if err := json.Unmarshal(data, &lib); err != nil {
		return fmt.Errorf("failed to parse configure request: %w", err)
	}

	// 2. Enrich the library configuration with Python-specific conventions.
	// For a new library, we set a default version and determine source paths.
	// As noted in go/sdk-librarian-python, resetting the version is an important
	// step for ensuring a clean migration from the old system.
	log.Printf("Enriching configuration for library: %s", lib.ID)
	lib.Version = "0.1.0" // Starting version for a new library.
	// Example: id `google-cloud-storage` -> path `google/cloud/storage`
	pythonStylePath := strings.Replace(lib.ID, "-", "/", -1)
	lib.SourcePaths = []string{
		// Main source code path
		filepath.Join("packages", lib.ID, pythonStylePath),
		// Tests path
		filepath.Join("packages", lib.ID, "tests"),
	}
	// By default, we remove the old source paths before copying new ones.
	lib.RemoveRegex = lib.SourcePaths

	// 3. Write the enriched configuration back for the Librarian CLI to process.
	log.Printf("Writing configure response to %s", responsePath)
	outData, err := json.MarshalIndent(lib, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize configure response: %w", err)
	}

	return ioutil.WriteFile(responsePath, outData, 0644)
}

// handleGenerate is responsible for the 'generate' step.
//
// This function orchestrates the core logic of code generation. For Phase 1,
// this means invoking the existing Bazel-based build process, effectively
// putting "OwlBot in a box".
func handleGenerate() error {
	log.Println("--- Running Generate Step (Phase 1: Owlbot in a box) ---")
	requestPath := "/librarian/generate-request.json"

	// 1. Parse the generation request.
	log.Printf("Reading generate request from %s", requestPath)
	data, err := ioutil.ReadFile(requestPath)
	if err != nil {
		return fmt.Errorf("failed to read generate request: %w", err)
	}
	var lib Library
	if err := json.Unmarshal(data, &lib); err != nil {
		return fmt.Errorf("failed to parse generate request: %w", err)
	}

	// The /output directory is the destination for all generated code.
	outputDir := "/output"

	// 2. Invoke the underlying code generator.
	// In Phase 1, this is a Python script that wraps Bazel. We simulate that call here.
	// The actual script would take parameters pointing to /source (APIs), /input,
	// and /output. See go/sdk-librarian-python for CLI examples.
	log.Println("Invoking the Python generator CLI (which wraps Bazel)...")
	cmd := exec.Command("python3", "tools/python_generator_cli/cli.py",
		"generate-library",
		"--api-root=/source",
		"--generator-input=/input",
		fmt.Sprintf("--output=%s", outputDir),
		fmt.Sprintf("--library-id=%s", lib.ID),
	)
	// In a real run, we would capture and log stdout/stderr.
	// For this example, we just log the simulated command.
	log.Printf("Simulating command: %s", cmd.String())
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("python generator cli failed: %w", err)
	// }
	log.Println("Generator finished.")

	// 3. Post-processing.
	// In the "Owlbot in a box" model, the generator script handles most of this.
	// Any steps that Librarian needs to do after the container runs (like copying
	// files from /output) are handled by the Librarian CLI, not this script.
	log.Println("Post-processing is handled by the generator script and Librarian CLI.")

	return nil
}

// handleBuild is responsible for the 'build' step.
//
// This function's task is to verify that the newly generated code is valid.
// In the "Owlbot in a box" phase, this means running the Bazel tests.
func handleBuild() error {
	log.Println("--- Running Build Step (Phase 1: Owlbot in a box) ---")
	requestPath := "/librarian/build-request.json"

	log.Printf("Reading build request from %s", requestPath)
	data, err := ioutil.ReadFile(requestPath)
	if err != nil {
		return fmt.Errorf("failed to read build request: %w", err)
	}
	var lib Library
	if err := json.Unmarshal(data, &lib); err != nil {
		return fmt.Errorf("failed to parse build request: %w", err)
	}

	// The working directory for build commands is the root of the repository.
	workDir := "/repo"
	log.Printf("Working directory is %s", workDir)

	// In a real container, we would set the working directory for the command.
	// if err := os.Chdir(workDir); err != nil {
	// 	 return fmt.Errorf("failed to change directory to %s: %w", workDir, err)
	// }

	// 1. Run tests using Bazel.
	// The target would correspond to the specific library being built.
	log.Println("Running tests with Bazel...")
	target := fmt.Sprintf("//%s/...", lib.ID) // e.g. //google-cloud-storage/...
	testCmd := exec.Command("bazel", "test", target)
	testCmd.Dir = workDir
	log.Printf("Simulating command: %s", testCmd.String())
	// if err := testCmd.Run(); err != nil {
	// 	return fmt.Errorf("bazel test failed: %w", err)
	// }

	return nil
}

// applyTemplate is a helper to execute a Go template and write it to a file.
// Note: In the "Owlbot in a box" model, templating is handled inside the
// generator script itself, so this helper would not be used in Phase 1.
// It is kept here for future reference (Phase 3).
func applyTemplate(path string, tmpl string, data interface{}) error {
	t, err := template.New(filepath.Base(path)).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", path, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer f.Close()

	return t.Execute(f, data)
}

// --- Templates ---
// In a real implementation, these would be loaded from the `/input` directory.
// They are included here for demonstration purposes for a future, non-Bazel generator.

const setupPyTemplate = `import setuptools

setuptools.setup(
    name="{{.ID}}",
    version="{{.Version}}",
    author="Google LLC",
    author_email="googleapis-packages@google.com",
    description="Python client for {{.ID}}",
    packages=setuptools.find_packages(),
    python_requires=">=3.7",
)
`

const readmeTemplate = `# Python Client for {{.ID}}

This is the Python client library for the {{.ID}} API.

## Installation

\`\`\`bash
pip install {{.ID}}
\`\`\`

## Usage

(Usage examples would go here)
`
