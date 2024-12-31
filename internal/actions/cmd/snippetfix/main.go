// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"cloud.google.com/go/internal/actions/logg"
)

var (
	dirFlag    = flag.String("dir", "", "the root directory to evaluate")
	modsFlag   = flag.String("modules", "", "Comma-separated list of module names to include (not paths).")
	modelFlag  = flag.String("model", "gemini-2.0-flash-exp", "Name of generative model to use.")
	modelsFlag = flag.Bool("list-models", false, "If true, ignores other args, skips normal execution and retrieves and lists available generative language models.")
	promptFlag = flag.Bool("prompt", false, "If true, skips generation and prints the prompt.")
)

func main() {
	flag.BoolVar(&logg.Quiet, "q", false, "quiet mode, minimal logging")
	flag.Parse()

	if *modelsFlag {
		listModels()
		return
	}

	rootDir, err := os.Getwd()
	if err != nil {
		logg.Fatal(err)
	}
	if *dirFlag != "" {
		rootDir = *dirFlag
	}
	logg.Printf("Root dir: %q", rootDir)
	logg.Printf("Model: %q", *modelFlag)

	modDirs, err := modDirs(rootDir)
	if err != nil {
		logg.Fatal(err)
	}

	includeMods := []string{}
	if *modsFlag != "" {
		includeMods = strings.Split(*modsFlag, ",")
		logg.Printf("modules to include: %s", *modsFlag)
	} else {
		logg.Println("snippetfix running on all modules.")
	}

	updatedSubmoduleDirs := []string{}
	if len(includeMods) > 0 {
		for _, mod := range includeMods {
			if strings.HasPrefix(mod, "internal") {
				continue
			}
			if !isSubmod(mod, modDirs) {
				logg.Printf("no module for: %s", mod)
				continue
			}
			updatedSubmoduleDirs = append(updatedSubmoduleDirs, mod)
		}
	} else {
		updatedSubmoduleDirs = modDirs
	}

	logg.Printf("processing: \n%s\n", strings.Join(updatedSubmoduleDirs, "\n"))
	for _, mod := range updatedSubmoduleDirs {
		err = processMod(rootDir, mod)
		if err != nil {
			logg.Fatal(err)
		}
	}
}

func processMod(rootDir, mod string) error {
	snippetsDir := fmt.Sprintf("%s/internal/generated/snippets/%s", rootDir, mod)
	if _, err := os.Stat(snippetsDir); errors.Is(err, os.ErrNotExist) {
		logg.Println("no snippets for: ", mod)
		return nil
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel(*modelFlag)

	rpc := "CreateCustomer"
	snippetPath := fmt.Sprintf("%s/apiv1/CloudChannelClient/%s/main.go", snippetsDir, rpc)
	snippet, err := os.ReadFile(snippetPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", snippetPath, err)
	}
	libraryPath := fmt.Sprintf("%s/%s/apiv1", rootDir, mod)
	library, err := readFilesFromGlob(libraryPath)
	if err != nil {
		return fmt.Errorf("failed to read files %s: %w", libraryPath, err)
	}
	prompt := fmt.Sprintf(promptTemplate, rpc, mod, rpc, rpc, snippet, mod, mod, library)
	if *promptFlag {
		fmt.Println(prompt)
		return nil
	}

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
	return nil
}

const promptTemplate = `
# Show me how to use %s with the Google Cloud %s API Golang library.

## Detailed instructions

Modify the following %s example inside the <code> tags, populating only required fields in the request struct. Do not populate any fields labeled output only or optional. Remove the TODO comment from inside the request struct. Add imports if necessary. Prefer imports from cloud.google.com/go to imports from google.golang.org/genproto. Otherwise, do not change the example in any way. Always use tabs to indent in Go source code, except for in comments. Preserve space characters used in comments. Do not wrap output in markdown. Do not wrap output in backticks. Do not return anything other than Go code.

## Example of %s

<code>
%s
</code>

## Populating fields in the request struct

Use the appropriate structs from the Google Cloud %s API Golang library source code included below to identify the required fields in the structs for your answer. Populate all required fields in your example with appropriate values.
Do not include any fields that are labeled output only. Do not include any fields that are optional, unless you feel they are important to the example.

## Source files for the Google Cloud %s API Golang library

<code>
%s
</code>
`

const apiKey = "AIzaSyC7nPP0bNZMc7YTAXff3RGAW8UmovRdXYw"

func listModels() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	modelIter := client.ListModels(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Models:")
	for {
		model, err := modelIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		modelName := strings.TrimPrefix(model.Name, "models/")
		fmt.Printf("  - %s\n", modelName)
	}
	fmt.Println("---")
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func readFilesFromGlob(root string) (string, error) {
	var fileContents []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		log.Println("visit: ", path)
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*.go", d.Name())
		log.Println("matched: ", matched, "*.go", d.Name())
		if err != nil {
			return fmt.Errorf("match error: %w", err)
		}

		if matched {
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}
			fileContents = append(fileContents, string(data))

		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to walk directory: %w", err)
	}

	return strings.Join(fileContents, "\n\n"), nil
}

func isSubmod(includeMod string, modDirs []string) bool {
	for _, modDir := range modDirs {
		if includeMod == modDir {
			return true
		}
	}
	return false
}

func modDirs(dir string) (submodulesDirs []string, err error) {
	c := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	c.Dir = dir
	b, err := c.Output()
	if err != nil {
		return submodulesDirs, err
	}
	// Skip the root mod
	list := strings.Split(strings.TrimSpace(string(b)), "\n")[1:]

	submodulesDirs = []string{}
	for _, modPath := range list {
		// Skip internal
		if strings.Contains(modPath, "internal") {
			continue
		}
		modPath = strings.TrimPrefix(modPath, dir+"/")
		// Skip nested
		if strings.Contains(modPath, "/") {
			continue
		}
		submodulesDirs = append(submodulesDirs, modPath)
	}

	return submodulesDirs, nil
}
