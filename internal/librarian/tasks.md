# `librariangen` Implementation Plan

This document outlines the development tasks for creating the `librariangen` Go application. The development process will follow an iterative, test-driven approach.

## Phase 1: CLI Scaffolding and Core Logic

This phase focuses on building the command-line interface and the foundational components for parsing inputs.

*   [x] **1. Project Setup**
*   [x] **2. Implement CLI Framework**
*   [x] **3. Implement `generate-request.json` Parsing**

## Phase 2: `protoc` Generation

This phase implements the core logic of the generator: reading Bazel configurations and executing `protoc`.

*   [x] **4. Implement `BUILD.bazel` Parser**
*   [x] **5. Implement `protoc` Command Builder and Executor**

## Phase 3: Integration

This phase integrates the `protoc` generation logic into the main CLI application.

*   [x] **6. Integrate Components in `handleGenerate`**

## Phase 4: Binary Integration Testing

This phase focuses on testing the compiled Go binary directly to ensure all components work together correctly outside of a container.

*   [x] **7. Create Binary Integration Test Script**
    *   **Coding:**
        *   Create a new script: `run-binary-integration-test.sh`.
        *   The script should:
            1.  Create a temporary directory structure to simulate the required inputs (e.g., `/tmp/test-env/source`, `/tmp/test-env/librarian`, `/tmp/test-env/output`).
            2.  Copy the `generator/protoc/testdata/source` directory into the temporary source directory.
            3.  Create a sample `generate-request.json` in the temporary librarian directory, pointing to an API to generate (e.g., `google/cloud/workflows/v1`).
            4.  Compile the `librariangen` binary using `go build`.
            5.  Execute the compiled binary with the `generate` command, using flags (`--source`, `--librarian`, `--output`) to point to the temporary directories.
            6.  Verify that the command succeeds and that `.pb.go` files are created in the temporary output directory.
            [ ] 7.  Check out the `https://github.com/googleapis/googleapis-gen` private repository to the temporary root directory.
            [ ] 8.  Copy the temporary output directory contents into the `googleapis-gen/tree/master/google/cloud/workflows/v1` directory.
            [ ] 9.  Run `git diff` and raise an error with the output if non-empty. This means that the output of `librarian` gen does not match the output of the legacy Bazel build.
            10. Clean up the temporary directories.
    *   **Testing:**
        *   Run the script locally to confirm it passes.

## Phase 5: Post-processing

This phase implements the steps that run after generation to make the code a complete, release-ready Go module.

*   [ ] **8. Implement Post-processing Steps**
    *   **Coding:**
        *   Create a new package `postprocessor` (`librariangen/postprocessor`).
        *   Create a main function `Process(outputDir string, version string, modulePath string) error`.
        *   Implement the following post-processing steps:
        *   [ ] Run `goimports -w .`
        *   [ ] Run `go mod init <modulePath>`
        *   [ ] Run `go mod tidy`
        *   [ ] Generate `version.go` file.
        *   [ ] Run `staticcheck ./...`
        *   [ ] Implement snippet metadata updates.
    *   **Integration Testing:**
        *   Create `postprocessor/postprocessor_integration_test.go` to test the `Process` function.

## Phase 6: Containerization and E2E Testing

This phase focuses on packaging the application in Docker and performing high-level, end-to-end validation.

*   [ ] **9. Create Dockerfile**
    *   **Coding:**
        *   Create the `generator/Dockerfile` as specified in the design document.
        *   Use a multi-stage build to create a minimal final image.
        *   Ensure the container build process is hermetic.
        *   Ensure all tool versions are pinned.
        *   Set the `ENTRYPOINT` to `/librariangen`.

*   [ ] **10. Create Test and Deployment Scripts**
    *   [ ] Create a `test.sh` smoke test script to perform basic validation of the final Docker image.
    *   [ ] Create a `run-container-tests.sh` script to execute the full, containerized workflow.
    *   [ ] Document the process for publishing the container.

*   [ ] **11. Manual Integration Testing**
    *   [ ] Execute the manual testing plan as laid out in `design.md` to confirm that the output of the core generation matches the legacy system.

## Phase 7: Contingency and Future Work

This phase is for the Bazel fallback strategy and documenting future work.

*   [ ] **12. Implement Bazel Fallback Strategy**
    *   **Description:** Implement a secondary generation strategy that invokes Bazel directly as a fallback for complex APIs.
    *   **Coding:**
        *   [ ] Add a flag or configuration to trigger the Bazel-based path.
        *   [ ] Implement logic to construct and execute a `bazel build` command.
        *   [ ] Implement logic to find and copy the generated artifacts.
    *   **Testing:**
        *   Create integration tests specifically for the Bazel-based generation path.

*   [ ] **13. Document Unresolved Issues**
    *   [ ] In the `README.md`, document known limitations (e.g., handling of global files).
    *   [ ] Investigate and document a proposed solution for structured error reporting.
    *   [ ] Add placeholder implementations for the `configure` and `build` commands.
