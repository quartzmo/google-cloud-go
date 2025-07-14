# Go GAPIC Librarian

This document describes the Go GAPIC Librarian, a containerized application that serves as the Go-specific code generator within the Librarian pipeline. Its responsibility is to generate release-ready Go GAPIC client libraries based on googleapis API definitions.

This generator adheres to the container contract defined by the central Librarian tool. It is an implementation of the [Librarian CLI: Reimagined](http://goto.google.com/librarian:cli-reimagined) design.

## Overview

The Librarian pipeline delegates language-specific tasks to dedicated container images. This project provides the container image for Go. The central Librarian tool is responsible for orchestrating the overall workflow, which includes:

1.  Reading configuration from a `state.yaml` file in the target language repository.
2.  Preparing the necessary inputs (API source files, configuration).
3.  Invoking this container with a specific command and a set of mounted directories.
4.  Processing the output of this container (e.g., copying generated code, updating state).

This model replaces legacy systems that relied on complex tooling like Bazel for generation, addressing security and maintenance concerns outlined in the [Python Librarian Migration](http://goto.google.com/sdk-librarian-python) design, the principles of which apply here.

## The Container Contract

The generator is invoked with a command as a positional argument. It expects a specific set of directories to be mounted into its filesystem, which provide the context needed to perform its task.

There are three primary commands defined in the contract: `configure`, `generate`, and `build`. Currently, only `generate` is implemented.

### `generate` Command

This is the core command, responsible for the actual code generation. It is invoked by the Librarian when it needs to generate a Go client library. The container expects a specific set of inputs to be provided as mounted directories and files.

#### **Inputs for the `generate` Command**

| Path or Argument | Type | Description for Go Generator |
| :--- | :--- | :--- |
| `generate` | Positional Argument | The string `generate` is the first argument passed to the container, signaling that it must execute the primary code generation logic. |
| `/librarian/generate-request.json` | File (Read) | A JSON file containing the metadata for the specific library to be generated. This is a subset of the main `state.yaml` file and tells the generator *which* APIs to build (e.g., `google/storage/v1`). Our Go application reads this file to know which API directory to process within the `/source` mount. |
| `/source` | Directory (Read) | This mount contains a complete checkout of the `googleapis` repository. It serves as the single, critical **import path** (`-I/source`) for the `protoc` command, allowing it to find and resolve all proto definitions and their dependencies (e.g., `google/api/annotations.proto`). |
| `/output` | Directory (Write) | An empty directory that the generator writes all of its output to. This includes the generated `.pb.go` and `_gapic.go` files. The Librarian is responsible for copying the contents of this directory to the correct location in the final language repository. |
| `/input` | Directory (Read) | Contains the contents of the `.librarian/generator-input` directory from the language repository. For Go, this directory is **not** used as a proto import path. Instead, it is reserved for future use, such as holding templates for `README.md` files or scripts for post-generation code tweaks. |

### How `generate` Works

The `main.go` entrypoint is straightforward. It determines which command to run based on the first argument passed to the container.

```go
// generator/main.go

func run(ctx context.Context) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("expected at least one argument for the command, got %d", len(os.Args)-1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "generate":
		return generateCmd(ctx)
	// ... other cases
	}
}
```

The `generateCmd` function orchestrates the generation logic:

1.  **Read the Request:** It reads and unmarshals the `/librarian/generate-request.json` file. This file contains the definition of the library to be generated.

    *Example `generate-request.json` (derived from `state.yaml` by Librarian):*
    ```json
    {
      "id": "google-cloud-storage-v1",
      "version": "1.15.0",
      "apis": [
        {
          "path": "google/storage/v1",
          "service_config": "storage.yaml"
        }
      ],
      "source_paths": [
        "storage/apiv1"
      ]
    }
    ```

    *The Go structs for this JSON in `main.go`:*
    ```go
    type Library struct {
    	ID          string   `json:"id"`
    	Version     string   `json:"version"`
    	APIs        []API    `json:"apis"`
    	SourcePaths []string `json:"source_paths"`
    }

    type API struct {
    	Path          string `json:"path"`
    	ServiceConfig string `json:"service_config"`
    }
    ```

2.  **Invoke Protoc:** For each API listed in the request, it constructs and executes a `protoc` command to generate the Go client library and supporting files.

    *The `protoc` function in `main.go` assembles the arguments:*
    ```go
    func protoc(ctx context.Context, lib *Library, api *API) error {
        // ... finds all .proto files in the API's source directory ...

        // Example: gapicImportPath becomes "cloud.google.com/go/storage/apiv1"
        gapicImportPath := filepath.Join("cloud.google.com/go", lib.SourcePaths[0])

        args := []string{
            "--experimental_allow_proto3_optional",
            "--go_out=" + outputDir,
            "--go-gapic_out=" + outputDir,
            "--go-gapic_opt=go-gapic-package=" + gapicImportPath,
            // The /source mount contains the entire googleapis repository,
            // which is used as the sole import path for protoc.
            "-I=" + sourceDir,
        }
        if api.ServiceConfig != "" {
            args = append(args, "--go-gapic_opt=api-service-config="+filepath.Join(apiSourceDir, api.ServiceConfig))
        }
        args = append(args, protoFiles...)

        cmd := exec.CommandContext(ctx, "protoc", args...)
        return runCommand(cmd)
    }
    ```
    The generated files are written directly to the `/output` directory. The Librarian tool will then handle copying them to the final destination in the language repository.

## How to Build and Run

The generator is packaged as a Docker image.

### Dockerfile

The `Dockerfile` is designed to be MOSS-compliant and self-contained.

```dockerfile
# generator/Dockerfile

# Use a MOSS-recommended base image
FROM marketplace.gcr.io/google/debian12:latest

# Set and install tool versions
ENV GO_VERSION=1.23.0
ENV PROTOC_VERSION=25.7
# ... other ENVs ...
RUN apt-get update && apt-get install -y wget ...
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && ...
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && ...

# Install Go protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GO_PROTOC_PLUGIN_VERSION} && \
    go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v${GAPIC_GENERATOR_VERSION}

# Copy and build the generator source
WORKDIR /app
COPY main.go .
RUN go build -o /generator main.go

# Set the entrypoint
ENTRYPOINT ["/generator"]
```

### Local Invocation Example

While the Librarian tool will normally invoke this container, you can run it locally for development and testing. This requires simulating the directory mounts that Librarian provides.

**Setup:**

1.  **Build the Docker image:**
    ```bash
    docker build -t go-generator:latest .
    ```

2.  **Create a mock environment:**
    ```bash
    # 1. Create directories to mount
    mkdir -p /tmp/mock/librarian /tmp/mock/input /tmp/mock/output
    git clone https://github.com/googleapis/googleapis.git /tmp/mock/source

    # 2. Create the request file (using the mock input provided)
    cp ../input/configure-request.json /tmp/mock/librarian/generate-request.json
    ```

**Execution:**

This command simulates the Librarian invoking the `generate` command.

```bash
# Go GAPIC Librarian

This document describes the Go GAPIC Librarian, a containerized application that serves as the Go-specific code generator within the Librarian pipeline. Its responsibility is to generate release-ready Go GAPIC client libraries based on googleapis API definitions.

This generator adheres to the container contract defined by the central Librarian tool. It is an implementation of the [Librarian CLI: Reimagined](http://goto.google.com/librarian:cli-reimagined) design.

## Overview

The Librarian pipeline delegates language-specific tasks to dedicated container images. This project provides the container image for Go. The central Librarian tool is responsible for orchestrating the overall workflow, which includes:

1.  Reading configuration from a `state.yaml` file in the target language repository.
2.  Preparing the necessary inputs (API source files, configuration).
3.  Invoking this container with a specific command and a set of mounted directories.
4.  Processing the output of this container (e.g., copying generated code, updating state).

This model replaces legacy systems that relied on complex tooling like Bazel for generation, addressing security and maintenance concerns outlined in the [Python Librarian Migration](http://goto.google.com/sdk-librarian-python) design, the principles of which apply here.

## The Container Contract

The generator is invoked with a command as a positional argument. It expects a specific set of directories to be mounted into its filesystem, which provide the context needed to perform its task.

There are three primary commands defined in the contract: `configure`, `generate`, and `build`. Currently, only `generate` is implemented.

### `generate` Command

This is the core command, responsible for the actual code generation. It is invoked by the Librarian when it needs to generate a Go client library. The container expects a specific set of inputs to be provided as mounted directories and files.

#### **Inputs for the `generate` Command**

| Path or Argument | Type | Description for Go Generator |
| :--- | :--- | :--- |
| `generate` | Positional Argument | The string `generate` is the first argument passed to the container, signaling that it must execute the primary code generation logic. |
| `/librarian/generate-request.json` | File (Read) | A JSON file containing the metadata for the specific library to be generated. This is a subset of the main `state.yaml` file and tells the generator *which* APIs to build (e.g., `google/storage/v1`). Our Go application reads this file to know which API directory to process within the `/source` mount. |
| `/source` | Directory (Read) | This mount contains a complete checkout of the `googleapis` repository. It serves as the single, critical **import path** (`-I/source`) for the `protoc` command, allowing it to find and resolve all proto definitions and their dependencies (e.g., `google/api/annotations.proto`). |
| `/output` | Directory (Write) | An empty directory that the generator writes all of its output to. This includes the generated `.pb.go` and `_gapic.go` files. The Librarian is responsible for copying the contents of this directory to the correct location in the final language repository. |
| `/input` | Directory (Read) | Contains the contents of the `.librarian/generator-input` directory from the language repository. For Go, this directory is **not** used as a proto import path. Instead, it is reserved for future use, such as holding templates for `README.md` files or scripts for post-generation code tweaks. |

### How `generate` Works

The `main.go` entrypoint is straightforward. It determines which command to run based on the first argument passed to the container.

```go
// generator/main.go

func run(ctx context.Context) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("expected at least one argument for the command, got %d", len(os.Args)-1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "generate":
		return generateCmd(ctx)
	// ... other cases
	}
}
```

The `generateCmd` function orchestrates the generation logic:

1.  **Read the Request:** It reads and unmarshals the `/librarian/generate-request.json` file. This file contains the definition of the library to be generated.

    *Example `generate-request.json` (derived from `state.yaml` by Librarian):*
    ```json
    {
      "id": "google-cloud-storage-v1",
      "version": "1.15.0",
      "apis": [
        {
          "path": "google/storage/v1",
          "service_config": "storage.yaml"
        }
      ],
      "source_paths": [
        "storage/apiv1"
      ]
    }
    ```

    *The Go structs for this JSON in `main.go`:*
    ```go
    type Library struct {
    	ID          string   `json:"id"`
    	Version     string   `json:"version"`
    	APIs        []API    `json:"apis"`
    	SourcePaths []string `json:"source_paths"`
    }

    type API struct {
    	Path          string `json:"path"`
    	ServiceConfig string `json:"service_config"`
    }
    ```

2.  **Invoke Protoc:** For each API listed in the request, it constructs and executes a `protoc` command to generate the Go client library and supporting files.

    *The `protoc` function in `main.go` assembles the arguments:*
    ```go
    func protoc(ctx context.Context, lib *Library, api *API) error {
        // ... finds all .proto files in the API's source directory ...

        // Example: gapicImportPath becomes "cloud.google.com/go/storage/apiv1"
        gapicImportPath := filepath.Join("cloud.google.com/go", lib.SourcePaths[0])

        args := []string{
            "--experimental_allow_proto3_optional",
            "--go_out=" + outputDir,
            "--go-gapic_out=" + outputDir,
            "--go-gapic_opt=go-gapic-package=" + gapicImportPath,
            // The /source mount contains the entire googleapis repository,
            // which is used as the sole import path for protoc.
            "-I=" + sourceDir,
        }
        if api.ServiceConfig != "" {
            args = append(args, "--go-gapic_opt=api-service-config="+filepath.Join(apiSourceDir, api.ServiceConfig))
        }
        args = append(args, protoFiles...)

        cmd := exec.CommandContext(ctx, "protoc", args...)
        return runCommand(cmd)
    }
    ```
    The generated files are written directly to the `/output` directory. The Librarian tool will then handle copying them to the final destination in the language repository.

## How to Build and Run

The generator is packaged as a Docker image.

### Dockerfile

The `Dockerfile` is designed to be MOSS-compliant and self-contained.

```dockerfile
# generator/Dockerfile

# Use a MOSS-recommended base image
FROM marketplace.gcr.io/google/debian12:latest

# Set and install tool versions
ENV GO_VERSION=1.23.0
ENV PROTOC_VERSION=25.7
# ... other ENVs ...
RUN apt-get update && apt-get install -y wget ...
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && ...
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && ...

# Install Go protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GO_PROTOC_PLUGIN_VERSION} && \
    go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v${GAPIC_GENERATOR_VERSION}

# Copy and build the generator source
WORKDIR /app
COPY main.go .
RUN go build -o /generator main.go

# Set the entrypoint
ENTRYPOINT ["/generator"]
```

### Local Invocation Example

While the Librarian tool will normally invoke this container, you can run it locally for development and testing. This requires simulating the directory mounts that Librarian provides.

**Setup:**

1.  **Build the Docker image:**
    ```bash
    docker build -t go-generator:latest .
    ```

2.  **Create a mock environment:**
    ```bash
    # 1. Create directories to mount
    mkdir -p /tmp/mock/librarian /tmp/mock/input /tmp/mock/output
    git clone https://github.com/googleapis/googleapis.git /tmp/mock/source

    # 2. Create the request file (using the mock input provided)
    cp ../input/configure-request.json /tmp/mock/librarian/generate-request.json
    ```

**Execution:**

This command simulates the Librarian invoking the `generate` command.

```bash
docker run --rm \
  -v /tmp/mock/librarian:/librarian:ro \
  -v /tmp/mock/input:/input:ro \
  -v /tmp/mock/source:/source:ro \
  -v /tmp/mock/output:/output:rw \
  go-generator:latest generate
```

After the command completes, the `/tmp/mock/output` directory will contain the generated Go files, just as it would in the real pipeline.

## Post-Processing Challenges

An attempt was made to run the `post-processor` binary on the generated code in the `/output` directory. This approach is not feasible because the `post-processor` is a complex tool designed for a CI/CD environment, not for simple post-generation code formatting. It has several environmental expectations that are not met inside the generator container:

1.  **Requires `config.yaml`:** The `post-processor` requires a `config.yaml` file to load its configuration, which defines modules, import paths, and other repository-specific settings. This file is not available in the generator's context.

2.  **Expects a Full Repository Checkout:** The tool is designed to operate on a complete `google-cloud-go` repository. It expects to find a `.git` directory to run `git` commands (e.g., for analyzing commit history) and a `go.work` file to manage Go modules. The `/output` directory contains only the newly generated files for a single library and does not resemble a full repository checkout.

3.  **Includes GitHub-Specific Logic:** A significant portion of the `post-processor`'s functionality involves interacting with the GitHub API to fetch pull request details and amend their titles and bodies. This logic is irrelevant and will fail in the context of a code generation container, which has no awareness of a specific pull request.

4.  **Generating `version.go` files:** This task requires identifying which subdirectories in `/output` are client packages that need a `version.go` file.
    *   **Identifying Client Directories:** In the main `post-processor`, this is done by walking the module directory and looking for `doc.go` as an indicator of a client package. In the generator, we don't have `doc.go` yet. A viable heuristic is to walk the `/output` directory and look for subdirectories whose names match the pattern `apiv...` and contain generated `*.pb.go` files. This is less robust than the original method but should work in most cases.
    *   **Determining Module Name:** The `version.go` template requires the module's import path (e.g., `cloud.google.com/go/asset/internal`). The module name (`asset`) must be derived from the `LibrarianRequest.ID` (e.g., `google-cloud-asset-v1`). This requires parsing and string manipulation that assumes a consistent naming convention for library IDs.
    *   **Hardcoded Exclusions:** The original function contains hardcoded logic to skip certain directories (e.g., `orgpolicy/apiv1`). This logic must be preserved and might need to be made more configurable in the future.

### Recommendation

The `post-processor` in its current form is not suitable for use within the generator. A more effective approach would be to create a separate, focused tool or script for post-generation tasks. This new tool should be self-contained and operate only on the directory of generated files, performing tasks like running `goimports`, `staticcheck`, or other linters without depending on a full repository structure or external services like GitHub.

## Next Steps and Future Improvements

While the current generator provides the core functionality for the `generate` command, several key improvements are planned to make it a complete and robust solution.

### 1. Implement `configure` and `build` Commands

To fully comply with the Librarian container contract, the remaining commands must be implemented.

*   **`configure` Command:** This command is essential for onboarding new libraries. It will read a minimal request file (e.g., with just a library ID) and enrich it with Go-specific conventions, such as setting a default version (`0.1.0`) and deriving `source_paths` from the ID (e.g., `google-cloud-storage-v1` -> `storage/apiv1`). The enriched configuration will be written back to `/librarian/configure-response.json`.

*   **`build` Command:** This command serves as a critical quality gate. When invoked, it will have the entire language repository (with the newly generated code) mounted at `/repo`. Its job is to navigate to the correct Go module directory (using the `source_paths` from the build request) and run `go build ./...` and `go test ./...` to validate the generated code.

### 2. Make Configuration More Robust

The generator currently makes simplifying assumptions that should be replaced with explicit configuration.

*   **Go Module Path:** The base Go module path (e.g., `cloud.google.com/go`) is currently hardcoded. This should be passed in the `generate-request.json` to make the generator more versatile and reusable across different Go projects.

### 3. Add Support for Post-Processing

The `/input` directory is currently ignored but is intended for post-generation "tweaks."

*   **Templates:** The generator should check for template files (e.g., `README.md.tpl`) in the `/input` directory. If found, it should execute them using Go's `text/template` package with the library's metadata and write the result to the `/output` directory.
*   **Scripts:** The generator could look for a `post-generate.sh` script in `/input` and execute it after `protoc` finishes to perform custom cleanup or code modifications.

### 4. Add Formal Go Module and Tests

To improve maintainability, the generator itself should be a proper Go module with its own tests.

*   **`go.mod`:** Initialize a `go.mod` file for the generator application.
*   **Unit Tests:** Add `_test.go` files to unit test specific logic, such as the construction of `protoc` arguments.
*   **Integration Tests:** Create tests that simulate the container's environment by creating temporary directories for `/source`, `/librarian`, and `/output`, allowing for end-to-end validation of the `generateCmd` logic.
