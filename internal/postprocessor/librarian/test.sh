#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# This script tests the 'generate' command of the librarian container.
# It sets up a mock environment with the required directories and files,
# builds the Docker container, runs it, and checks for the expected output.

set -euo pipefail

# -- Setup --
# Create a temporary directory for test artifacts.
# The `trap` command ensures this directory is cleaned up on script exit.
TEST_DIR=$(mktemp -d)
trap 'rm -rf -- "$TEST_DIR"' EXIT
echo "🧪 Test directory created at: $TEST_DIR"

# Create mock directories for the container mounts.
# /source simulates the googleapis repository.
# /output is where the generated code will be written.
# /librarian contains the request file.
mkdir -p "$TEST_DIR/source/google/cloud/accessapproval/v1"
mkdir -p "$TEST_DIR/output"
mkdir -p "$TEST_DIR/librarian"

# Create a mock source file to be copied during generation.
touch "$TEST_DIR/source/google/cloud/accessapproval/v1/accessapproval.pb.go"
echo "✅ Mock source file created."

# Create the generate-request.json file. This file tells the container
# which library to generate. The structure is based on the librarian guide.
# We are requesting the 'accessapproval' library.
cat > "$TEST_DIR/librarian/generate-request.json" <<EOF
{
  "id": "accessapproval",
  "version": "1.7.0",
  "last_generated_commit": "a1b2c3d4e5f6",
  "apis": [
    {
      "path": "google/cloud/accessapproval/v1",
      "service_config": "accessapproval_v1.yaml"
    }
  ],
  "source_paths": [
    "accessapproval"
  ]
}
EOF
echo "✅ generate-request.json created."

# -- Build --
# Create a go.mod file needed for the build.
# This must be done in the same directory as main.go.
# The user must create this file themselves before running the test.
if [ ! -f "go.mod" ]; then
    echo " go.mod not found. Please run 'go mod init <module_name> && go mod tidy' first."
    exit 1
fi

IMAGE_NAME="librarian-test:latest"
echo "🐳 Building Docker image: $IMAGE_NAME..."
docker build -t "$IMAGE_NAME" .
echo "✅ Docker image built successfully."


# -- Run --
# Run the container, mounting the mock directories to the paths
# expected by the container contract, and passing the specified flags
# to the 'generate' command.
# NOTE: This assumes the container's entrypoint (`main.go`) has been
# updated to parse these flags for the 'generate' command.
echo "🚀 Running 'generate' command in container..."
docker run --rm \
  -v "$TEST_DIR/source:/source" \
  -v "$TEST_DIR/output:/output" \
  -v "$TEST_DIR/librarian:/librarian" \
  "$IMAGE_NAME" generate \
  --repo="https://github.com/googleapis/google-cloud-go" \
  --library="accessapproval" \
  --api="google/cloud/accessapproval/v1" \
  --api-source="/source" \
  --output="/output" \
  --build
echo "✅ Container finished execution."


# -- Verify --
# Check if the generated file exists in the output directory.
# The expected path is /output/{source_path}/{api_version_folder}/{filename}.
EXPECTED_FILE="$TEST_DIR/output/accessapproval/v1/accessapproval.pb.go"
echo "🔍 Verifying output file: $EXPECTED_FILE"
if [ -f "$EXPECTED_FILE" ]; then
  echo "✅ SUCCESS: Generated file found."
else
  echo "❌ FAILURE: Generated file not found."
  # List the contents of the output directory for debugging.
  echo "Output directory contents:"
  ls -R "$TEST_DIR/output"
  exit 1
fi
