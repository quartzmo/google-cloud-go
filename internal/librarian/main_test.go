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
	"context"
	"flag"
	"strings"
	"testing"
)

// testRun is a helper to invoke the run function with a given set of arguments.
// It sets up and tears down the command-line flags for each run to ensure
// test isolation, as the `flag` package uses global state.
func testRun(t *testing.T, args ...string) error {
	// Each test should have a fresh FlagSet to avoid polluting the global flags.
	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	// Register the application's flags to this new, test-specific FlagSet.
	// We capture the pointers to local variables so the `run` function can access
	// the parsed values.
	var (
		testLibrarianDir = fs.String("librarian", "/librarian", "Path to the librarian-tool input directory.")
		testInputDir     = fs.String("input", "/input", "Path to the .librarian/generator-input directory.")
		testOutputDir    = fs.String("output", "/output", "Path to the empty directory where the librariangen writes its output.")
		testSourceDir    = fs.String("source", "/source", "Path to a complete checkout of the googleapis repository.")
	)

	// Parse the test arguments.
	if err := fs.Parse(args); err != nil {
		t.Fatalf("failed to parse flags: %v", err)
	}

	// The run function expects to get its commands from flag.Args(), so we need
	// to temporarily replace the global flag set with our test one.
	originalFlags := flag.CommandLine
	flag.CommandLine = fs
	// Also swap the global flag variables to point to our test variables.
	originalLibrarianDir, originalInputDir, originalOutputDir, originalSourceDir := librarianDir, inputDir, outputDir, sourceDir
	librarianDir, inputDir, outputDir, sourceDir = testLibrarianDir, testInputDir, testOutputDir, testSourceDir

	defer func() {
		// Restore the original global state after the test.
		flag.CommandLine = originalFlags
		librarianDir, inputDir, outputDir, sourceDir = originalLibrarianDir, originalInputDir, originalOutputDir, originalSourceDir
	}()

	return handleCommand(context.Background())
}

// TestRun_Commands verifies that the run function correctly dispatches to the
// correct command handler based on the first positional argument, and that it
// handles unknown or missing commands gracefully.
func TestRun_Commands(t *testing.T) {
	testCases := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "generate command fails to read missing file",
			args:    []string{"generate"},
			wantErr: "failed to read request file",
		},
		{
			name:    "configure command succeeds (no-op)",
			args:    []string{"configure"},
			wantErr: "",
		},
		{
			name:    "build command succeeds (no-op)",
			args:    []string{"build"},
			wantErr: "",
		},
		{
			name:    "unknown command",
			args:    []string{"unknown"},
			wantErr: "unknown command: unknown",
		},
		{
			name:    "no command",
			args:    []string{},
			wantErr: "expected at least one argument",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testRun(t, tc.args...)

			if tc.wantErr == "" {
				if err != nil {
					t.Errorf("run() with args %v returned unexpected error: %v", tc.args, err)
				}
			} else {
				if err == nil {
					t.Errorf("run() with args %v expected error, got nil", tc.args)
				} else if !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("run() with args %v returned error %q, want error containing %q", tc.args, err.Error(), tc.wantErr)
				}
			}
		})
	}
}

// TestRun_Flags verifies that the flags are parsed correctly and that their
// values are used by the application logic. It does this by providing a custom
// path for the --librarian flag and checking that the subsequent file-read
// error message contains the custom path.
func TestRun_Flags(t *testing.T) {
	tempDir := t.TempDir()
	args := []string{"-librarian", tempDir, "generate"}

	err := testRun(t, args...)

	wantErr := "failed to read request file from " + tempDir
	if err == nil {
		t.Fatal("run() with custom flag expected error, got nil")
	}
	if !strings.Contains(err.Error(), wantErr) {
		t.Errorf("run() with custom flag returned error %q, want error containing %q", err.Error(), wantErr)
	}
}
