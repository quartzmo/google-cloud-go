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
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	manifestFile = ".release-please-manifest-submodules.json"
)

var (
	dir           = flag.String("dir", "", "the root directory to evaluate")
	changesRegexp = regexp.MustCompile(`\[([0-9]+\.[0-9]+\.[0-9]+)\]`)
	moduleRegexp  = regexp.MustCompile(`"([0-9]+\.[0-9]+\.[0-9]+)"`)
)

func main() {
	flag.Parse()
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if *dir != "" {
		rootDir = *dir
	}

	b, err := os.ReadFile(manifestFile)
	if err != nil {
		log.Fatal(err)
	}
	manifest := make(map[string]interface{})
	json.Unmarshal(b, &manifest)

	submodulesDirs, err := modDirs(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("|module|manifest|changelog|version.go|")
	fmt.Println("|------|--------|---------|----------|")
	for _, submodDir := range submodulesDirs {
		if strings.HasPrefix(submodDir, "internal") {
			continue
		}
		manifestVersion, ok := manifest[submodDir]
		if !ok {
			continue
		}
		changesVersion := lastChangesVersion(submodDir)
		moduleVersion := moduleVersion(submodDir)
		if manifestVersion != changesVersion || manifestVersion != moduleVersion {
			fmt.Printf("|%s|%s|%s|%s|\n", submodDir, manifestVersion, changesVersion, moduleVersion)
		}
	}
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
		if strings.Contains(modPath, "internal") {
			continue
		}
		modPath = strings.TrimPrefix(modPath, dir+"/")
		submodulesDirs = append(submodulesDirs, modPath)
	}

	return submodulesDirs, nil
}

func lastChangesVersion(dir string) string {
	file, err := os.Open(fmt.Sprintf("%s/CHANGES.md", dir))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		if c := changesRegexp.FindStringSubmatch(line); len(c) > 0 {
			s := strings.TrimPrefix(c[0], "[")
			s = strings.TrimSuffix(s, "]")
			return s
		}
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return "nil"
}

func moduleVersion(dir string) string {
	file, err := os.Open(fmt.Sprintf("%s/internal/version.go", dir))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		if c := moduleRegexp.FindStringSubmatch(line); len(c) > 0 {
			s := strings.TrimPrefix(c[0], "\"")
			s = strings.TrimSuffix(s, "\"")
			return s
		}
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return "nil"
}
