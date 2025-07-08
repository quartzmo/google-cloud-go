# Go Client Library Generator

This document describes the Go Client Library Generator, a containerized application that serves as the Go-specific component within the reimagined Librarian pipeline. Its primary responsibility is to generate high-quality, release-ready Go client libraries based on API definitions.

This generator adheres to a strict container contract defined by the central Librarian tool, ensuring a modular, maintainable, and secure generation process. It is a direct implementation of the principles laid out in the [Librarian CLI: Reimagined](http://goto.google.com/librarian:cli-reimagined) design.

## Overview

The Librarian pipeline delegates language-specific tasks to dedicated container images. This `generator` is the official image for Go. The central Librarian tool is responsible for orchestrating the overall workflow, which includes:

1.  Reading configuration from a `state.yaml` file in the target language repository.
2.  Preparing the necessary inputs (API source files, configuration).
3.  Invoking this container with a specific command and a set of mounted directories.
4.  Processing the output of this container (e.g., copying generated code, updating state).

This model replaces legacy systems that relied on complex tooling like Bazel for generation, addressing security and maintenance concerns outlined in the [Python Librarian Migration](http://goto.google.com/sdk-librarian-python) design, the principles of which apply here.

## The Container Contract

The generator is invoked with a command as a positional argument. It expects a specific set of directories to be mounted into its filesystem, which provide the context needed to perform its task.

There are three primary commands defined in the contract: `configure`, `generate`, and `build`. Currently, only `generate` is implemented.

### `generate` Command

This is the core command, responsible for the actual code generation.

| Context | Type | Description |
| :--- | :--- | :--- |
| `/librarian/generate-request.json` | Mount (Read) | A JSON file that describes which library to generate. Its schema is a subset of the `libraries` entry in `state.yaml`. |
| `/input` | Mount (Read) | The contents of the `.librarian/generator-input` folder from the language repository. This can contain shared or language-specific files needed for generation. |
| `/output` | Mount (Write) | An empty directory where the container must write all generated code. The directory structure should match the desired structure in the final repository. |
| `/source` | Mount (Read) | A directory containing the API definitions (e.g., a clone of `googleapis/googleapis`). |
| `generate` | Positional Arg | The command that invokes this operation. |

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
        // ... finds all .proto files ...

        // Example: gapicImportPath becomes "cloud.google.com/go/storage/apiv1"
        gapicImportPath := filepath.Join("cloud.google.com/go", lib.SourcePaths[0])

        args := []string{
            "--experimental_allow_proto3_optional",
            "--go_out=" + outputDir,
            "--go-gapic_out=" + outputDir,
            "--go-gapic_opt=go-gapic-package=" + gapicImportPath,
            // Include paths for googleapis and shared generator inputs
            "-I=" + sourceDir,
            "-I=" + inputDir,
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
