# `librariangen` Implementation Plan

This document outlines the development tasks for creating the `librariangen` Go application, which will replace the legacy generation toolchain for `google-cloud-go`. The development process will follow an iterative, test-driven approach. Each feature implementation will be followed by unit testing, review, and refactoring.

## Phase 1: CLI Scaffolding and Core Logic

This phase focuses on building the command-line interface and the foundational components for parsing inputs.

*   [ ] **1. Project Setup**
    *   [ ] Create the `librariangen` directory.
    *   [ ] Create a `main.go` file.
    *   [ ] Initialize the Go module: `go mod init github.com/googleapis/google-cloud-go/internal/postprocessor/librarian/generator`.

*   [ ] **2. Implement CLI Framework**
    *   **Coding:**
        *   In `main.go`, implement the command-line argument parsing.
        *   Use the `flag` package or a library like `cobra` to define subcommands: `generate`, `configure`, `build`.
        *   Implement the required filepath flags for local development: `--source`, `--librarian`, `--output`, `--input`.
        *   Create placeholder functions for each command (`handleGenerate`, `handleConfigure`, `handleBuild`).
        *   The `main` function should dispatch to the correct handler based on the subcommand.
    *   **Testing:**
        *   Create `main_test.go`.
        *   Write unit tests to verify that the correct handler is called for each subcommand.
        *   Write unit tests to ensure flags are parsed correctly and default values are set.

*   [ ] **3. Implement `generate-request.json` Parsing**
    *   **Coding:**
        *   Create a new package `request` (`librariangen/request`).
        *   In `request/request.go`, define the Go structs that map to the structure of `/librarian/generate-request.json`.
        *   Create a function `Parse(path string) (*Request, error)` that reads the file and unmarshals the JSON into the structs.
        *   Integrate this parsing logic into the `handleGenerate` function.
    *   **Testing:**
        *   Create `request/request_test.go`.
        *   Add a `testdata` directory with a sample `generate-request.json`.
        *   Write unit tests to confirm that the JSON is parsed correctly.
        *   Test error cases: file not found, malformed JSON.

## Phase 2: `protoc` Generation

This phase implements the core logic of the generator: reading Bazel configurations and executing `protoc`.

*   [ ] **4. Implement `BUILD.bazel` Parser**
    *   **Coding:**
        *   Create a new package `bazel` (`librariangen/bazel`).
        *   In `bazel/parser.go`, implement a function `Parse(path string) (*Config, error)`.
        *   This parser will not be a full Starlark interpreter. It should be a robust line-by-line or regex-based parser focused *only* on extracting key-value pairs from the `go_gapic_library` and `go_proto_library` rules.
        *   It needs to find `importpath`, `grpc_service_config`, `service_yaml`, `transport`, `release_level`, etc.
        *   The function should be able to find and read the correct `BUILD.bazel` file based on an API path (e.g., `google/cloud/workflows/v1`).
    *   **Testing:**
        *   Create `bazel/parser_test.go`.
        *   Add a `testdata` directory with several example `BUILD.bazel` files covering different cases (e.g., `go_gapic_library` vs. `go_grpc_library`, different transports, etc.).
        *   Write comprehensive unit tests to ensure all required options are extracted correctly.
        *   Test for edge cases like missing rules or malformed files.

*   [ ] **5. Implement `protoc` Command Builder**
    *   **Coding:**
        *   Create a new package `protoc` (`librariangen/protoc`).
        *   In `protoc/command.go`, create a builder that takes the parsed `request.Request` and `bazel.Config` as input.
        *   The builder will construct the full `protoc` command as a `[]string`, including all `--go-gapic_opt` flags.
    *   **Testing:**
        *   Create `protoc/command_test.go`.
        *   Write unit tests that pass mock config data to the builder and verify that the resulting command slice is 100% correct.

*   [ ] **6. Implement `protoc` Command Execution**
    *   **Coding:**
        *   In `protoc/command.go`, implement a function `Run(cmd []string, outputDir string) error`.
        *   This function will use `os/exec` to run the `protoc` command.
        *   It must capture and log `stdout` and `stderr` for debugging purposes.
        *   It must return an error if the command fails.
    *   **Integration Testing:**
        *   This step requires `protoc` and the Go plugins to be installed and in the `PATH`.
        *   Create an integration test file `protoc/command_integration_test.go`.
        *   The test should:
            1.  Set up a temporary directory structure mimicking `/source`, `/librarian`, and `/output`.
            2.  Include minimal `.proto` and `BUILD.bazel` files in the test's `/source`.
            3.  Call the `protoc` builder and runner.
            4.  Verify that the expected `.go` files are generated in the test's `/output` directory.

## Phase 3: Post-processing

This phase implements the steps that run after `protoc` to make the generated code a complete, release-ready Go module.

*   [ ] **7. Implement Post-processor Framework**
    *   **Coding:**
        *   Create a new package `postprocessor` (`librariangen/postprocessor`).
        *   In `postprocessor/postprocessor.go`, create a main function `Process(outputDir string, version string, modulePath string) error`.
        *   This function will orchestrate calls to the other post-processing steps.

*   [ ] **8. Implement Post-processing Steps**
    *   **Coding:**
        *   For each task below, create a dedicated, well-documented function in the `postprocessor` package. Each function will use `os/exec` to run the corresponding command-line tool within the `outputDir`.
        *   [ ] Run `goimports -w .`
        *   [ ] Run `go mod init <modulePath>`
        *   [ ] Run `go mod tidy`
        *   [ ] Generate `version.go` file (using a simple Go template).
        *   [ ] Run `staticcheck ./...`
    *   **Integration Testing:**
        *   Create `postprocessor/postprocessor_integration_test.go`.
        *   This requires `go`, `goimports`, and `staticcheck` to be in the `PATH`.
        *   The test should:
            1.  Set up a temporary directory with a few dummy `.go` files.
            2.  Run the main `Process` function.
            3.  Verify that `go.mod`, `go.sum`, and `version.go` files are created and have the correct content.
            4.  Verify that the command executions succeed.

## Phase 4: Finalization and Documentation

This phase brings all the components together, creates the final Docker image, and adds documentation.

*   [ ] **9. Integrate Components in `handleGenerate`**
    *   **Coding:**
        *   Flesh out the `handleGenerate` function in `main.go`.
        *   It should now contain the full, end-to-end logic:
            1.  Parse the `generate-request.json`.
            2.  Loop through each API in the request.
            3.  For each API, parse its `BUILD.bazel` file.
            4.  Build the `protoc` command.
            5.  Execute the `protoc` command.
            6.  After the loop, call the `postprocessor.Process` function.
        *   Ensure robust error handling at each step.
    *   **Review and Refactor:**
        *   Perform a full code review of the entire application.
        *   Refactor any duplicated logic (e.g., consolidate command execution helpers).
        *   Improve comments and variable names.

*   [ ] **10. Create Dockerfile**
    *   **Coding:**
        *   Create the `generator/Dockerfile` as specified in the design document.
        *   Use a multi-stage build. The first stage installs Go, `protoc`, and the plugins. The second stage builds the `librariangen` binary. The final stage is a minimal image containing only the compiled binary and required tools.
        *   Ensure all tool versions are pinned.
        *   Set the `ENTRYPOINT` to `/librariangen`.

*   [ ] **11. Write Documentation**
    *   **Coding:**
        *   Create a `generator/README.md`.
        *   Document the purpose of the `librariangen`.
        *   Provide clear instructions on how to build the binary and the Docker image.
        *   Explain how to run the tool locally for development, including the required flags.
        *   Explain how to run the tests.
