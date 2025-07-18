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
        *   Create placeholder functions for each command (`handleGenerate`, `handleConfigure`, `handleBuild`). The initial implementation will focus on `handleGenerate`.
        *   The `main` function should dispatch to the correct handler based on the subcommand and return a non-zero exit code on error.
    *   **Testing:**
        *   Create `main_test.go`.
        *   Write unit tests to verify that the correct handler is called for each subcommand.
        *   Write unit tests to ensure flags are parsed correctly and default values are set.

*   [ ] **3. Implement `generate-request.json` Parsing**
    *   **Coding:**
        *   Create a new package `request` (`librariangen/request`).
        *   In `request/request.go`, define the Go structs that map to the structure of `/librarian/generate-request.json`, including the `source_roots`, `preserve_regex`, and `remove_regex` fields.
        *   Create a function `Parse(path string) (*Request, error)` that reads the file and unmarshals the JSON into the structs.
        *   Integrate this parsing logic into the `handleGenerate` function.
    *   **Testing:**
        *   Create `request/request_test.go`.
        *   Add a `testdata` directory with a sample `generate-request.json` that includes all new fields.
        *   Write unit tests to confirm that the JSON is parsed correctly.
        *   Test error cases: file not found, malformed JSON.

## Phase 2: `protoc`-based Generation

This phase implements the primary logic of the generator: reading Bazel configurations and executing `protoc` directly.

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

*   [ ] **5. Implement `protoc` Command Builder and Executor**
    *   **Coding:**
        *   Create a new package `protoc` (`librariangen/protoc`).
        *   In `protoc/command.go`, create a builder that takes the parsed `request.Request` and `bazel.Config` as input and constructs the full `protoc` command as a `[]string`.
        *   Implement a function `Run(cmd []string, outputDir string) error` that uses `os/exec` to run the command, capturing output for debugging.
    *   **Integration Testing:**
        *   Create an integration test file `protoc/command_integration_test.go`.
        *   The test should set up a temporary directory structure, call the builder and runner, and verify that the expected `.go` files are generated in the correct subdirectories of `/output`.

## Phase 3: Post-processing

This phase implements the steps that run after generation to make the code a complete, release-ready Go module.

*   [ ] **6. Implement Post-processing Steps**
    *   **Coding:**
        *   Create a new package `postprocessor` (`librariangen/postprocessor`).
        *   Create a main function `Process(outputDir string, version string, modulePath string) error`.
        *   For each task below, create a dedicated, well-documented function in the `postprocessor` package. Each function will use `os/exec` to run the corresponding command-line tool within the appropriate subdirectory of `outputDir`.
        *   [ ] Run `goimports -w .`
        *   [ ] Run `go mod init <modulePath>`
        *   [ ] Run `go mod tidy` (using `-mod=vendor` if necessary for hermeticity).
        *   [ ] Generate `version.go` file (using a simple Go template).
        *   [ ] Run `staticcheck ./...`
        *   [ ] Implement snippet metadata updates (e.g., updating the `$VERSION` placeholder in `snippet_metadata.*.json` files).
    *   **Integration Testing:**
        *   Create `postprocessor/postprocessor_integration_test.go`.
        *   The test should set up a temporary directory with dummy `.go` files and a `snippet_metadata.json` file, run the `Process` function, and verify that all files are created and modified correctly.

## Phase 4: Finalization and Documentation

This phase brings all the components together, creates the final Docker image, and adds documentation.

*   [ ] **7. Integrate Components in `handleGenerate`**
    *   **Coding:**
        *   Flesh out the `handleGenerate` function in `main.go` to contain the full, end-to-end logic for the `protoc`-based approach.
        *   Ensure robust error handling at each step.
    *   **Review and Refactor:**
        *   Perform a full code review of the application.
        *   Refactor any duplicated logic and improve comments.

*   [ ] **8. Create Dockerfile**
    *   **Coding:**
        *   Create the `generator/Dockerfile` as specified in the design document.
        *   Use a multi-stage build to create a minimal final image containing the compiled binary and all required tools (`protoc`, `go`, `goimports`, `staticcheck`, etc.).
        *   Ensure the container build process is hermetic (has no network access). This may involve vendoring Go modules.
        *   Ensure all tool versions are pinned.
        *   Set the `ENTRYPOINT` to `/librariangen`.

*   [ ] **9. Write Documentation**
    *   **Coding:**
        *   Create a `generator/README.md`.
        *   Document the purpose of the `librariangen`, how to build it, how to run it locally, and how to run tests.
        *   Add a section detailing known limitations, such as the handling of global files.

## Phase 5: Testing and Validation

This phase focuses on the high-level testing required to validate the entire solution.

*   [ ] **10. Create Test and Deployment Scripts**
    *   [ ] Create a `test.sh` smoke test script to perform basic validation of the final Docker image.
    *   [ ] Create a `run-container-tests.sh` script to execute the full, containerized workflow for a suite of test libraries.
    *   [ ] Document the process for publishing the container to the `us-central1-docker.pkg.dev` repository, including any IAM requirements.

*   [ ] **11. Manual Integration Testing**
    *   [ ] Execute the manual testing plan as laid out in `design.md`:
        1.  Onboard a single canary library.
        2.  Onboard the ~90% of straightforward libraries.
        3.  Onboard the 8 "difficult" libraries.
        4.  Onboard all remaining libraries.
    *   For each step, confirm that the output produces **no diff** against the code generated by the legacy system. Document all findings.

## Phase 6: Contingency and Future Work

This phase is for the Bazel fallback strategy and documenting future work.

*   [ ] **12. Implement Bazel Fallback Strategy**
    *   **Description:** As a mitigation for APIs with complex or unusual `BUILD.bazel` files that the parser cannot handle, implement a secondary generation strategy that invokes Bazel directly.
    *   **Coding:**
        *   [ ] Add a flag or configuration to `librariangen` to trigger the Bazel-based path.
        *   [ ] Implement logic to construct and execute a `bazel build` command for the target API.
        *   [ ] Implement logic to find the generated artifacts in the `bazel-bin` directory and copy them to the correct location within the `/output` directory.
    *   **Testing:**
        *   Create integration tests specifically for the Bazel-based generation path, verifying it produces the correct output.

*   [ ] **13. Document Unresolved Issues**
    *   [ ] Investigate and document a proposed solution for structured error reporting from the container.
    *   [ ] Add placeholder implementations for the `configure` and `build` commands with clear "Not Implemented" errors. Create tracking issues for their future implementation.
