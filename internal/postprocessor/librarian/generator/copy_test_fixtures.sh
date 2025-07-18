#!/bin/bash

# This script sets up the testdata directory for the protoc integration test.
# It copies the necessary .proto files from local checkouts of the
# googleapis and protobuf repositories.

set -e # Exit immediately if a command exits with a non-zero status.
set -x # Print commands and their arguments as they are executed.

# --- Configuration ---

# The root directory of the local googleapis checkout.
GOOGLEAPIS_ROOT="/Users/chrisdsmith/oss/googleapis"

# The official repository for Google's well-known proto types.
PROTOBUF_REPO="https://github.com/protocolbuffers/protobuf.git"

# The target directory for the test fixtures.
TARGET_DIR="testdata/source"

# The directory containing the primary protos to be compiled.
MAIN_PROTO_DIR="google/cloud/workflows"

# List of core dependency proto directories/files needed from googleapis.
DEPS=(
  "google/api"
  "google/longrunning"
  "google/rpc"
)

# List of configuration files needed for the workflows API.
CONFIG_FILES=(
  "google/cloud/workflows/v1/BUILD.bazel"
  "google/cloud/workflows/v1/workflows_v1.yaml"
  "google/cloud/workflows/v1/workflows_grpc_service_config.json"
  "google/cloud/workflows/v1beta/BUILD.bazel"
  "google/cloud/workflows/v1beta/workflows_v1beta.yaml"
  "google/cloud/workflows/v1beta/workflows_grpc_service_config.json"
  "google/cloud/workflows/executions/v1/BUILD.bazel"
  "google/cloud/workflows/executions/v1/workflowexecutions_v1.yaml"
  "google/cloud/workflows/executions/v1/executions_grpc_service_config.json"
  "google/cloud/workflows/executions/v1beta/BUILD.bazel"
  "google/cloud/workflows/executions/v1beta/workflowexecutions_v1beta.yaml"
  "google/cloud/workflows/executions/v1beta/executions_grpc_service_config.json"
)

# --- Execution ---

# 0. Clean up any previous testdata.
echo "Cleaning up old testdata..."
rm -rf "$TARGET_DIR"

# 1. Copy the main proto directory recursively from the local googleapis checkout.
echo "Copying main proto directory..."
mkdir -p "$TARGET_DIR/$(dirname "$MAIN_PROTO_DIR")"
cp -r "$GOOGLEAPIS_ROOT/$MAIN_PROTO_DIR" "$TARGET_DIR/$MAIN_PROTO_DIR"

# 2. Copy all core dependency directories from the local googleapis checkout.
echo "Copying dependency files from googleapis..."
for dep in "${DEPS[@]}"; do
  DEST_PATH="$TARGET_DIR/$dep"
  mkdir -p "$(dirname "$DEST_PATH")"
  cp -r "$GOOGLEAPIS_ROOT/$dep" "$DEST_PATH"
done

# 3. Copy all configuration files.
echo "Copying configuration files..."
for f in "${CONFIG_FILES[@]}"; do
  DEST_PATH="$TARGET_DIR/$f"
  mkdir -p "$(dirname "$DEST_PATH")"
  cp "$GOOGLEAPIS_ROOT/$f" "$DEST_PATH"
done

# 4. Clone the protobuf repo to a temporary directory to get well-known types.
PROTOBUF_TMP_DIR=$(mktemp -d)
echo "Cloning protobuf repository to $PROTOBUF_TMP_DIR..."
git clone --depth 1 "$PROTOBUF_REPO" "$PROTOBUF_TMP_DIR"

# 5. Copy the well-known types into the testdata directory.
echo "Copying well-known protobuf types..."
mkdir -p "$TARGET_DIR/google"
cp -r "$PROTOBUF_TMP_DIR/src/google/protobuf" "$TARGET_DIR/google/"

# 6. Clean up the temporary protobuf clone.
echo "Cleaning up temporary protobuf clone..."
rm -rf "$PROTOBUF_TMP_DIR"

echo "Successfully copied all test fixtures."
