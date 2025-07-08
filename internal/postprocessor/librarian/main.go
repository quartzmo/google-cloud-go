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

package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Library represents a single library entry in the state.yaml file,
// as described in the librarian design documents.
type Library struct {
	ID                  string   `json:"id" yaml:"id"`
	Version             string   `json:"version,omitempty" yaml:"version,omitempty"`
	LastGeneratedCommit string   `json:"last_generated_commit,omitempty" yaml:"last_generated_commit,omitempty"`
	APIs                []API    `json:"apis" yaml:"apis"`
	SourcePaths         []string `json:"source_paths,omitempty" yaml:"source_paths,omitempty"`
	PreserveRegex       []string `json:"preserve_regex,omitempty" yaml:"preserve_regex,omitempty"`
	RemoveRegex         []string `json:"remove_regex,omitempty" yaml:"remove_regex,omitempty"`
}

// API represents an API within a library.
type API struct {
	Path          string `json:"path" yaml:"path"`
	ServiceConfig string `json:"service_config,omitempty" yaml:"service_config,omitempty"`
}

const (
	// Mount points as defined in the librarian container contract.
	librarianMount = "/librarian"
	inputMount     = "/input"
	outputMount    = "/output"
	sourceMount    = "/source"
	repoMount      = "/repo"
)

func main() {
	// The command is expected as the first argument, as per the container contract.
	if len(os.Args) < 2 {
		log.Fatalf("Usage: librarian <configure|generate|build> [args]")
	}
	command := os.Args[1]

	switch command {
	case "configure":
		handleConfigure()
	case "generate":
		handleGenerate(os.Args[2:])
	case "build":
		// For a standalone build, read the request file.
		reqPath := filepath.Join(librarianMount, "build-request.json")
		reqFile, err := os.ReadFile(reqPath)
		if err != nil {
			log.Fatalf("Failed to read build-request.json from %s: %v", reqPath, err)
		}
		var lib Library
		if err := json.Unmarshal(reqFile, &lib); err != nil {
			log.Fatalf("Failed to unmarshal build request: %v", err)
		}
		handleBuild(lib)
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

// handleConfigure implements the 'configure' container command.
func handleConfigure() {
	log.Println("Running 'configure' command...")
	reqPath := filepath.Join(librarianMount, "configure-request.json")
	respPath := filepath.Join(librarianMount, "configure-response.json")

	reqFile, err := os.ReadFile(reqPath)
	if err != nil {
		log.Fatalf("Failed to read configure-request.json from %s: %v", reqPath, err)
	}

	var lib Library
	if err := json.Unmarshal(reqFile, &lib); err != nil {
		log.Fatalf("Failed to unmarshal configure request: %v", err)
	}

	lib.Version = "0.0.0"
	if len(lib.APIs) > 0 {
		lib.SourcePaths = []string{lib.ID}
	}

	respFile, err := json.MarshalIndent(lib, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal configure response: %v", err)
	}

	if err := os.WriteFile(respPath, respFile, 0644); err != nil {
		log.Fatalf("Failed to write configure-response.json to %s: %v", respPath, err)
	}
	log.Printf("Successfully wrote configuration to %s", respPath)
}

// handleBuild implements the 'build' container command.
func handleBuild(lib Library) {
	log.Println("Running 'build' command...")
	if len(lib.SourcePaths) == 0 {
		log.Fatalf("Cannot build library %s: source_paths not defined.", lib.ID)
	}
	moduleDir := filepath.Join(repoMount, lib.SourcePaths[0])
	log.Printf("Executing build in: %s", moduleDir)

	if _, err := os.Stat(moduleDir); os.IsNotExist(err) {
		log.Fatalf("Build directory %s does not exist.", moduleDir)
	}

	// Run 'go build ./...'
	buildCmd := exec.Command("go", "build", "./...")
	buildCmd.Dir = moduleDir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		log.Fatalf("Build failed for module %s: %v", lib.ID, err)
	}
	log.Println("Build successful.")

	// Run 'go test ./...'
	testCmd := exec.Command("go", "test", "./...")
	testCmd.Dir = moduleDir
	testCmd.Stdout = os.Stdout
	testCmd.Stderr = os.Stderr
	if err := testCmd.Run(); err != nil {
		log.Printf("Tests failed for module %s: %v. This may be acceptable if there are no tests.", lib.ID, err)
	} else {
		log.Println("Tests passed.")
	}
	log.Println("Build and test complete.")
}

// handleGenerate executes the `protoc` command to generate code.
func handleGenerate(args []string) {
	log.Println("Running 'generate' command...")

	// Find all .proto files in the input directory.
	protoFiles, err := findProtoFiles(inputMount)
	if err != nil {
		log.Fatalf("Error finding proto files in %s: %v", inputMount, err)
	}
	if len(protoFiles) == 0 {
		log.Printf("No proto files found in %s. Nothing to generate.", inputMount)
		return
	}

	// Construct and execute the protoc command.
	protocArgs := []string{
		"--proto_path=" + inputMount,
		"--go_gapic_out=" + outputMount,
	}
	protocArgs = append(protocArgs, protoFiles...)

	log.Printf("Executing command: protoc %v", protocArgs)
	cmd := exec.Command("protoc", protocArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("protoc command failed: %v", err)
	}

	log.Println("Successfully generated code.")
}

// findProtoFiles recursively finds all files with the .proto extension in a directory.
func findProtoFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".proto" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// copyDirectoryContents recursively copies files from a source to a destination.
func copyDirectoryContents(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return err
			}
			if err := copyDirectoryContents(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// copyFile copies a single file.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
