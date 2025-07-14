# Go Librarian Migration

# Objective

To migrate the Go GAPIC client library generation process from a complex, repository-specific Bazel implementation to a standardized, container-based system orchestrated by the central Librarian tool. This will simplify maintenance, improve security, and align Go's generation process with other languages in the Google Cloud ecosystem.

# Background

The current process for generating Go client libraries is tightly coupled to the `google-cloud-go` repository's Bazel build system. This approach presents several challenges:
*   **Complexity:** The generation logic is spread across numerous Bazel rules and scripts, making it difficult to understand, maintain, and debug.
*   **Security:** It requires a full checkout of the `google-cloud-go` repository and extensive permissions within the CI environment, posing potential security risks.
*   **Inconsistency:** The process is bespoke to Go, whereas other languages are migrating to a standardized Librarian pipeline. This inconsistency increases the cognitive load for developers and the SRE team.

The Librarian pipeline, which uses language-specific containers to handle generation, has already been successfully adopted for other languages like Python. This project aims to create the Go-specific component for this pipeline, thereby solving the issues above.

# Overview

This design proposes the creation of a self-contained, containerized Go generator that conforms to the Librarian tool's container contract. The central Librarian tool will manage the overall workflow, while the Go generator will be responsible solely for generating Go-specific artifacts.

The high-level workflow is as follows:

1.  **Invocation:** The Librarian tool invokes the Go generator Docker container with a specific command (e.g., `generate`).
2.  **Inputs:** Librarian provides all necessary inputs as mounted directories:
    *   `/source`: A complete checkout of the `googleapis` repository.
    *   `/librarian`: Contains a `generate-request.json` file specifying which API to generate.
    *   `/output`: An empty directory for the generated files.
    *   `/input`: A directory for optional, user-provided templates and scripts.
3.  **Execution:** The Go generator binary runs inside the container.
    *   It parses the `generate-request.json` to determine the target API.
    *   It invokes `protoc` with the Go plugins (`protoc-gen-go`, `protoc-gen-go_gapic`), using `/source` as the single import path (`-I`).
    *   It writes all generated `.go` files directly to the `/output` directory.
4.  **Post-Processing:** After generation, a built-in post-processing step formats the code, generates a `version.go` file, and runs linters.
5.  **Output:** The Librarian tool takes the contents of the `/output` directory and copies them to the correct location in the target `google-cloud-go` repository.

This model decouples the generation logic from any specific repository's build system, creating a simple, secure, and reusable component.

# Detailed Design

The implementation consists of a Go application packaged into a Docker container.

### **Container (`generator/Dockerfile`)**

The generator is built using a MOSS-compliant `debian:12` base image. The Dockerfile is responsible for:
*   Installing specific, pinned versions of Go, `protoc`, and other required tools.
*   Installing the necessary Go protoc plugins: `google.golang.org/protobuf/cmd/protoc-gen-go` and `github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic`.
*   Copying the generator's Go source code into the image.
*   Building the Go source into a single executable binary (`/generator`).
*   Setting the `ENTRYPOINT` to this `/generator` binary.

### **Generator Binary (`generator/main.go`)**

The Go application serves as the container's entrypoint. It is a simple command-line application that dispatches logic based on the first argument passed to the container.

```go
// The run function switches on the command provided by Librarian.
func run(ctx context.Context) error {
	cmd := os.Args[1]
	switch cmd {
	case "generate":
		return generateCmd(ctx)
	case "configure":
		// To be implemented
	case "build":
		// To be implemented
	}
}
```

### **`generate` Command**

This is the core command. Its logic is orchestrated in `generateCmd`:
1.  **Parse Request:** It reads and unmarshals the `/librarian/generate-request.json` file into Go structs. This file provides the API path (e.g., `google/cloud/asset/v1`) and other metadata.
2.  **Load Configuration:** For each API, it parses the corresponding `BUILD.bazel` file located in the `/source` directory (e.g., `/source/google/cloud/asset/v1/BUILD.bazel`). This file is the source of truth for `protoc` options such as the transport layer (`grpc+rest`), service YAML path, and release level. This avoids having to duplicate this configuration.
3.  **Execute `protoc`:** It constructs and executes the `protoc` command. The key arguments are:
    *   `--go_out=/output` and `--go-gapic_out=/output`: Directs all generated files to the output directory.
    *   `-I=/source`: Sets the `googleapis` checkout as the sole import path, simplifying dependency resolution.
    *   `--go-gapic_opt=...`: Passes the options extracted from the `BUILD.bazel` file to the Go GAPIC generator.

### **Post-Processing (`generator/postprocessor.go`)**

A critical design decision was to forgo the existing, complex `post-processor` in favor of a new, lightweight implementation directly within the generator. The old tool was unsuitable as it depends on a full `google-cloud-go` repository checkout and makes calls to the GitHub API.

The new, self-contained post-processor runs entirely within the `/output` directory and performs the following essential tasks:
1.  **`goimports`:** Runs `goimports -w .` to format the generated code and fix imports.
2.  **Module Initialization:** Runs `go mod init` and `go mod tidy` to create a valid `go.mod` file for the newly generated code. This is a prerequisite for running other Go tools.
3.  **`version.go` Generation:** A `version.go` file is generated for each new module, containing the library version provided in the request.
4.  **Linting:** Runs `staticcheck ./...` to catch potential issues in the generated code.

### **Testing**

The generator will be a formal Go module with its own `go.mod` file. The plan includes:
*   **Unit Tests:** For specific logic, such as parsing `BUILD.bazel` files and constructing `protoc` arguments.
*   **Integration Tests:** Tests that simulate the Librarian invocation by creating a temporary filesystem with `/source`, `/librarian`, and `/output` directories, allowing for end-to-end validation of the `generate` command.

# Alternatives considered

1.  **Status Quo (Bazel-based Generation):** The primary alternative was to continue using the existing Bazel-based system. This was rejected due to its high maintenance cost, complexity, and security concerns. It also prevents Go from aligning with the standardized, cross-language Librarian pipeline.

2.  **Adapt the Old Post-Processor:** An attempt was made to run the existing `post-processor` tool inside the container. This was deemed infeasible. The tool is not designed to be portable; it requires a full `google-cloud-go` repository context (including a `.git` directory and `go.work` file) and contains logic for interacting with GitHub pull requests, none of which are available or relevant inside the generator container. The chosen approach of a new, focused post-processor is far simpler and more robust.

# Work Estimates

The core `generate` functionality is complete. The following estimates cover the work required to make the generator feature-complete according to the Librarian contract.

*   **Week 1: Implement `configure` command.** This involves logic to read a minimal request, apply Go-specific conventions (e.g., default version, source path derivation), and write an enriched response file.
*   **Week 2: Implement `build` command.** This command will validate the generated code by running `go build ./...` and `go test ./...` within a simulated repository environment.
*   **Week 3: Add comprehensive tests.** Develop unit and integration tests to ensure the reliability of all commands and post-processing logic.
*   **Week 4: Enhance post-processing.** Add support for user-provided templates from the `/input` directory and investigate adding more valuable linters.

# Documentation plan

*   **`generator/README.md`:** This will be the primary technical document for the Go generator, detailing its architecture, commands, and how to run it locally. It will be kept up-to-date as new features are added.
*   **`docs/librarian_guide.md`:** This central Librarian document will be updated to include instructions and details for onboarding and generating Go libraries using the new system.
*   **This Design Doc:** This document will serve as the persistent, high-level design reference.

All documentation will be updated and made available before the first library is fully migrated to the new system.

# Launch plans

The migration to the new generator will be a gradual process, managed by the central Librarian tool's configuration (`state.yaml`).

*   **Visible changes:** The code generation process for Go libraries will become faster and more transparent. The underlying CI/CD jobs will be simplified significantly.
*   **Impact on production:** There will be no direct impact on production services. The generated code will be functionally identical to the code produced by the old system.
*   **New servers:** No new infrastructure is required; this system leverages the existing Librarian pipeline.
*   **Supportability:** Long-term support will be greatly simplified. The generator is a standard Go application, making it easier for any Go developer to contribute, as opposed to the specialized Bazel knowledge required by the old system.
*   **Timeline:** The rollout will proceed on a library-by-library basis over the course of a few weeks, starting with a few pilot libraries to validate the process before expanding to all Go clients.

# Risks

*   **Configuration Edge Cases:** The logic for parsing `BUILD.bazel` files in `googleapis` might not account for all possible configurations or edge cases.
    *   **Mitigation:** Test the generator against a diverse set of existing Go libraries before the migration. Add robust error handling for unexpected `BUILD.bazel` structures.
*   **Post-Processing Discrepancies:** The new, streamlined post-processor might not perform all the subtle modifications that the old system did, potentially leading to minor differences in generated code.
    *   **Mitigation:** Before migrating each library, perform a `diff` between the output of the old and new systems. Any significant differences will be investigated and the post-processor will be adjusted as needed.
*   **Tool Versioning:** The generator's `Dockerfile` pins versions for `protoc` and its Go plugins. These dependencies can become stale.
    *   **Mitigation:** Implement automated dependency scanning (e.g., RenovateBot) for the `Dockerfile` to create pull requests whenever new tool versions are released.