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
# set -x # Print commands and their arguments as they are executed.

LIBRARIANGEN_GO_VERSION=local
LIBRARIANGEN_LOG=librariangen.log
# Start with a clean log file.
rm -f "$LIBRARIANGEN_LOG"

# --- Dependency Checks & Version Info ---
(
echo "--- Tool Versions ---"
echo "Go: $(GOWORK=off GOTOOLCHAIN=${LIBRARIANGEN_GO_VERSION} go version)"
echo "protoc: $(protoc --version 2>&1)"
echo "protoc-gen-go: $(protoc-gen-go --version 2>&1)"
echo "protoc-gen-go_gapic: v0.53.1"
echo "---------------------"
) >> "$LIBRARIANGEN_LOG" 2>&1

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
  echo "protoc-gen-go_gapic not found in PATH. Installing..."
  (GOWORK=off GOTOOLCHAIN=${LIBRARIANGEN_GO_VERSION} go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v0.53.1)
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
GOWORK=off GOTOOLCHAIN=${LIBRARIANGEN_GO_VERSION} go build -o "$BINARY_PATH" .

# 4. Run the librariangen generate command.
echo "Running librariangen..."
PATH=$(GOWORK=off GOTOOLCHAIN=${LIBRARIANGEN_GO_VERSION} go env GOPATH)/bin:$HOME/go/bin:$PATH ./librariangen \
  --source="$SOURCE_DIR" \
  --librarian="$LIBRARIAN_DIR" \
  --output="$OUTPUT_DIR" \
  generate >> "$LIBRARIANGEN_LOG" 2>&1

# Run gofmt just like the Bazel rule:
# https://github.com/googleapis/gapic-generator-go/blob/main/rules_go_gapic/go_gapic.bzl#L34
# TODO: move this to librariangen
GOWORK=off GOTOOLCHAIN=${LIBRARIANGEN_GO_VERSION} gofmt -w -l $OUTPUT_DIR > /dev/null

# --- Verify ---

# 5. Check that the command succeeded and generated files.
echo "Verifying output..."
echo "Librariangen logs are available in: $LIBRARIANGEN_LOG"
if [ -z "$(ls -A "$OUTPUT_DIR")" ]; then
  echo "Error: Output directory is empty."
  exit 1
fi

# Use a cached version of googleapis-gen if available.
if [ ! -d "$GOOGLEAPIS_GEN_DIR" ]; then
  echo "Error: GOOGLEAPIS_GEN_DIR is not set or not a directory."
  echo "Please set it to the path of your local googleapis-gen clone."
  exit 1
fi
echo "Using cached googleapis-gen from $GOOGLEAPIS_GEN_DIR"
GEN_DIR="$GOOGLEAPIS_GEN_DIR"

# Define the API paths to verify.
APIS=(
  "workflows/apiv1"
  "workflows/apiv1beta"
  "workflows/executions/apiv1"
  "workflows/executions/apiv1beta"
)
# These are the corresponding paths in the googleapis-gen repository.
GEN_API_PATHS=(
  "google/cloud/workflows/v1"
  "google/cloud/workflows/v1beta"
  "google/cloud/workflows/executions/v1"
  "google/cloud/workflows/executions/v1beta"
)

# --- Verification using Git ---
echo "Verifying output by comparing with the googleapis-gen repository..."
echo "The script will modify files in your local googleapis-gen clone."

# Before files are copied, run git reset and git clean to clean up prior run.
(
echo "--- Git Reset Summary ---"
pushd "$GEN_DIR" > /dev/null
git reset --hard HEAD
git clean -fd
popd > /dev/null
) >> "$LIBRARIANGEN_LOG" 2>&1

# Process each API by replacing the expected files with the generated ones
for i in "${!APIS[@]}"; do
  api="${APIS[$i]}"
  gen_api_path="${GEN_API_PATHS[$i]}"

  OUTPUT_API_DIR="$OUTPUT_DIR/cloud.google.com/go/$api"
  EXPECTED_API_DIR="$GEN_DIR/$gen_api_path/cloud.google.com/go/$api"

  # 1. Remove everything from the expected directory
  rm -rf "${EXPECTED_API_DIR:?}"/*

  # 2. Copy over all files from the output directory
  # Ensure the directory exists after cleaning
  mkdir -p "$EXPECTED_API_DIR"
  cp -a "$OUTPUT_API_DIR"/. "$EXPECTED_API_DIR/"
done

# After all files are copied, run git add and git status to show changes.
# This entire section is redirected to the log file for later inspection.
(
echo "--- Git Status Summary ---"
pushd "$GEN_DIR" > /dev/null
# Stage all changes. New files, modifications, and deletions will be staged.
git add .
# Print the human-readable status
git status

# --- Diff of First Modified File ---
# Use `git diff --numstat` to find the first file with actual content changes,
# ignoring the noise from permission-only differences.
first_modified_file=$(git -C "$GEN_DIR" diff --staged --numstat | awk '$1 != "0" || $2 != "0" {print $3}' | head -n 1)

if [ -n "$first_modified_file" ]; then
  echo ""
  echo "--- Diff for first modified file: $first_modified_file ---"
  # Run git diff --staged to see the staged changes for that file in patch format.
  # We use `-c core.pager=cat` to prevent git from opening an interactive pager.
  git -C "$GEN_DIR" -c core.pager=cat diff --staged -p -- "$first_modified_file"
fi

popd > /dev/null
) >> "$LIBRARIANGEN_LOG" 2>&1

echo ""
echo "Verification complete. The status above shows the difference between the"
echo "expected generated output (goldens) and the current modified state of your googleapis-gen repository (librariangen)."
echo ""
echo -e "To reset your googleapis-gen repository:"
echo "  cd $GEN_DIR"
echo "  git reset --hard HEAD && git clean -fd"

# --- Cleanup ---
echo "Cleaning up..."
rm -rf "$TEST_DIR"
rm -f "$BINARY_PATH"

echo "Binary integration test passed successfully."
echo "Generated files are available for inspection in: $OUTPUT_DIR"
