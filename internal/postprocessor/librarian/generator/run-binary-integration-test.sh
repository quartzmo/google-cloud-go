#!/bin/bash

# This script performs an integration test on the compiled librariangen binary.
# It simulates the environment that the Librarian tool would create by:
# 1. Creating a temporary directory structure for inputs and outputs.
# 2. Copying the testdata protos into the temporary source directory.
# 3. Creating a generate-request.json file.
# 4. Compiling and running the librariangen binary with flags pointing to the
#    temporary directories.
# 5. Verifying that the binary succeeds and generates the expected files.

set -e # Exit immediately if a command exits with a non-zero status.
set -x # Print commands and their arguments as they are executed.

# --- Dependency Checks ---

# Ensure that all required protoc dependencies are available in PATH.
if ! command -v "protoc" &> /dev/null; then
  echo "Error: protoc not found in PATH. Please install it."
fi
if ! command -v "protoc-gen-go" &> /dev/null; then
  echo "Error: protoc-gen-go not found in PATH. Please install it."
  echo "  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
fi
if ! command -v "protoc-gen-go-grpc" &> /dev/null; then
  echo "Error: protoc-gen-go-grpc not found in PATH. Please install it."
  echo "  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
fi
if ! command -v "protoc-gen-go_gapic" &> /dev/null; then
  echo "Error: protoc-gen-go_gapic not found in PATH. Please install it."
  echo "  go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@latest"
fi

# --- Setup ---

# Create a temporary directory for the entire test environment.
TEST_DIR=$(mktemp -d)
echo "Using temporary directory: $TEST_DIR"

# Define the directories replicating the mounts in the Docker container.
SOURCE_DIR="$TEST_DIR/source"
LIBRARIAN_DIR="$TEST_DIR/librarian"
OUTPUT_DIR="$TEST_DIR/output"
mkdir -p "$SOURCE_DIR" "$LIBRARIAN_DIR" "$OUTPUT_DIR"

# The compiled binary will be placed in the current directory.
BINARY_PATH="./librariangen"

# --- Prepare Inputs ---

# 1. Copy the testdata protos into the temporary source directory.
echo "Copying test fixtures..."
cp -r "testdata/source/google" "$SOURCE_DIR/google"
cp "testdata/librarian/generate-request.json" "$LIBRARIAN_DIR/"

# --- Execute ---

# 3. Compile the librariangen binary.
echo "Compiling librariangen..."
go build -o "$BINARY_PATH" .

# 4. Run the librariangen generate command.
echo "Running librariangen..."
PATH=$(go env GOPATH)/bin:$HOME/go/bin:$PATH ./librariangen \
  --source="$SOURCE_DIR" \
  --librarian="$LIBRARIAN_DIR" \
  --output="$OUTPUT_DIR" \
  generate

# --- Verify ---

# 5. Check that the command succeeded and generated files.
echo "Verifying output..."
if [ -z "$(ls -A "$OUTPUT_DIR")" ]; then
  echo "Error: Output directory is empty."
  exit 1
fi

# Check for a specific generated file for each API.
EXPECTED_FILES=(
  "$OUTPUT_DIR/cloud.google.com/go/workflows/apiv1/workflowspb/workflows.pb.go"
  "$OUTPUT_DIR/cloud.google.com/go/workflows/apiv1beta/workflowspb/workflows.pb.go"
  "$OUTPUT_DIR/cloud.google.com/go/workflows/executions/apiv1/executionspb/executions.pb.go"
  "$OUTPUT_DIR/cloud.google.com/go/workflows/executions/apiv1beta/executionspb/executions.pb.go"
)
for f in "${EXPECTED_FILES[@]}"; do
  if [ ! -f "$f" ]; then
    echo "Error: Expected file not found: $f"
    ls -R "$OUTPUT_DIR" # List contents for debugging.
    exit 1
  fi
done

# --- Cleanup ---

# 6. Clean up the temporary directory.
# echo "Cleaning up..."
# rm -rf "$TEST_DIR"
# rm -f "$BINARY_PATH"

echo "Binary integration test passed successfully."
echo "Generated files are available for inspection in: $OUTPUT_DIR"
