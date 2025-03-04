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
	dir       = flag.String("dir", "", "the root directory to evaluate")
	write     = flag.Bool("w", false, "overwrite the version.go value with the manifest value")
	changesRe = regexp.MustCompile(`\[([0-9]+\.[0-9]+\.[0-9]+)\]`)
	moduleRe  = regexp.MustCompile(`"([0-9]+\.[0-9]+\.[0-9]+)"`)
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

	var count int
	fmt.Println("|num|module|manifest|version.go|")
	fmt.Println("|---|------|--------|----------|")
	for _, submodDir := range submodulesDirs {
		if strings.HasPrefix(submodDir, "internal") {
			continue
		}
		mv, ok := manifest[submodDir]
		if !ok {
			continue
		}
		manifestVersion, ok := mv.(string)
		if !ok {
			log.Fatalf("failed to cast value to string for key: %s", submodDir)
		}

		moduleVersion := moduleVersion(submodDir, manifestVersion)
		count++
		if manifestVersion != moduleVersion {
			fmt.Printf("|%d|%s|%s|%s|\n", count, submodDir, manifestVersion, moduleVersion)
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

func moduleVersion(dir, newValue string) string {
	file := fmt.Sprintf("%s/internal/version.go", dir)
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	m := moduleRe.FindStringSubmatch(string(b))
	if len(m) < 1 {
		log.Fatalf("expected to find version in version.go: %s", file)
	}
	moduleVersion := strings.TrimPrefix(m[0], "\"")
	moduleVersion = strings.TrimSuffix(moduleVersion, "\"")
	if *write && (moduleVersion != newValue) {
		b2 := moduleRe.ReplaceAll(b, []byte(fmt.Sprintf("\"%s\"", newValue)))
		if err := os.WriteFile(file, b2, 0644); err != nil {
			log.Fatalf("error writing file: %s", err)
		}
	}
	return moduleVersion
}
