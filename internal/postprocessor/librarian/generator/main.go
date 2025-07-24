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
	"fmt"
	"log/slog"
	"os"

	"cloud.google.com/go/internal/postprocessor/librarian/generator/generate"
)

// main is the entrypoint for the librariangen CLI.
func main() {
	slog.Info("librariangen invoked", "args", os.Args)
	if err := handleCommand(context.Background(), os.Args); err != nil {
		slog.Error("librariangen failed", "error", err)
		os.Exit(1)
	}
	slog.Info("librariangen finished successfully")
}

// handleCommand executes the appropriate command based on the CLI's invocation arguments.
// The first non-flag argument is the command (e.g., "generate").
func handleCommand(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("expected at least one argument for the command, got %d", len(args))
	}
	cmd := args[0]
	switch cmd {
	case "generate":
		return generate.Generate(ctx)
	case "configure":
		slog.Warn("configure command is not yet implemented")
		return nil
	case "build":
		slog.Warn("build command is not yet implemented")
		return nil
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}
