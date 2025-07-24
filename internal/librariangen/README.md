# Go GAPIC Generator (`librariangen`)

This directory contains the source code for `librariangen`, a containerized Go application that serves as the Go-specific code generator within the Librarian pipeline. Its responsibility is to generate release-ready Go GAPIC client libraries from `googleapis` API definitions, replacing the legacy `bazel-bot` and `OwlBot` toolchain.

## How it Works (The Container Contract)

The `librariangen` binary is designed to be run inside a Docker container orchestrated by the central Librarian tool. It adheres to a specific "container contract" by accepting commands and expecting a set of mounted directories for its inputs and outputs.

The primary command is `generate`.

### `generate` Command Workflow

1.  **Inputs:** The container is provided with several mounted directories:
    *   `/source`: A complete checkout of the `googleapis` repository. This is the primary include path (`-I`) for `protoc`.
    *   `/librarian`: Contains a `generate-request.json` file, which specifies the library and the specific API protos to be generated.
    *   `/output`: An empty directory where all generated Go files will be written.
    *   `/input`: A directory for future use (e.g., templates, scripts).

2.  **Execution (`gapicgen`):**
    *   The `librariangen` binary parses the `generate-request.json`.
    *   For each API specified in the request, it locates the corresponding `BUILD.bazel` file within the `/source` directory.
    *   It parses this `BUILD.bazel` file to extract the necessary options for the `protoc` command (e.g., import paths, transport, service config paths).
    *   It constructs and executes a `protoc` command, invoking the `protoc-gen-go` and `protoc-gen-go_gapic` plugins.

3.  **Output:** All generated files (`*.pb.go`, `*_gapic.go`, etc.) are written to the `/output` directory. The Librarian tool is then responsible for copying these files to their final destination in the `google-cloud-go` repository.

## Development & Testing

### Dependencies

To build and test `librariangen` locally, you must have the following tools installed and available in your `PATH`:

*   **Go Toolchain:** (Version `1.23.0` or higher)
*   **`protoc`:** (Version `25.7` or higher)
*   **`protoc-gen-go`:** `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
*   **`protoc-gen-go-grpc`:** `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
*   **`protoc-gen-go_gapic`:** `go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@latest`

### Building

To compile the binary:
```bash
go build -o librariangen/librariangen ./librariangen
```

### Testing

The project has a multi-layered testing strategy.

1.  **Unit Tests:** Each Go package has its own unit tests. To run all of them:
    ```bash
    go test ./librariangen/...
    ```

2.  **Binary Integration Test:** A shell script (`run-binary-integration-test.sh`) provides a full, end-to-end test of the compiled binary. This is the primary way to validate changes to the core generation logic.
    *   **Setup:** The test requires a local checkout of the `googleapis` repository. Before running the test for the first time, you must run the `copy_test_fixtures.sh` script to populate the `testdata` directory.
        ```bash
        # Run this once to set up the test data
        bash copy_test_fixtures.sh
        ```
    *   **Execution:**
        ```bash
        bash run-binary-integration-test.sh
        ```
    This script will compile the binary and run it against the realistic `workflows` API fixtures in the `testdata` directory, verifying that the correct Go files are generated.

## Future Work

This implementation currently focuses only on the core `generate` command. Future work includes:

*   **Post-processing:** Implementing a new, lightweight post-processor to run `goimports`, `go mod tidy`, and generate `version.go` files.
*   **`configure` and `build` commands:** Implementing the remaining commands from the Librarian container contract.
*   **Bazel Fallback:** Implementing a contingency plan to invoke Bazel directly for APIs with highly complex or unusual configurations.
