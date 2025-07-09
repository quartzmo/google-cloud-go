# Bazel-based Go GAPIC Librarian


This generator is maintained as a parallel strategy to the primary [`protoc`-based generator](../generator/README.md). It represents the "Phase 1: Owlbot in a box" migration strategy discussed in the [Python Librarian Migration](http://goto.google.com/sdk-librarian-python) design. The goal is to provide a compatibility layer for teams that have an existing, complex generation process built around Bazel, allowing them to onboard to the new Librarian pipeline with minimal disruption.

## Overview

Like the primary generator, this version adheres to the standard Librarian container contract. The central Librarian tool invokes this container to perform Go-specific tasks, but the internal implementation of those tasks relies on `bazel` commands instead of direct `protoc` calls.

The workflow is as follows:
1.  The Librarian tool reads the `state.yaml` file from a language repository.
2.  It invokes this container with a command (`generate`, `build`, etc.) and mounts the required directories (`/source`, `/librarian`, etc.).
3.  This container's entrypoint (`/bazel-generator`) receives the command and executes the corresponding Bazel-driven logic.
4.  The Librarian tool processes the results from the `/output` directory.

## The Container Contract

This generator implements the same three commands as the primary generator: `configure`, `generate`, and `build`.

### `generate` Command

This command uses `bazel run` to execute a generation script defined within the repository's Bazel workspace.

*   **Invocation:** `bazel-generator generate`
*   **Inputs:** Reads `/librarian/generate-request.json`.
*   **Core Logic (`main.go`):
    ```go
    func generateCmd(ctx context.Context) error {
        // ... reads request file ...

        // Invokes a hypothetical generator CLI that wraps Bazel.
        slog.Info("invoking Bazel-based generator script")
        args := []string{
            "run",
            "//tools/generator:main", // Example Bazel target
            "--",
            "generate-library",
            "--api-root=" + sourceDir,
            "--generator-input=" + inputDir,
            "--output=" + outputDir,
            "--library-id=" + lib.ID,
        }
        cmd := exec.CommandContext(ctx, "bazel", args...)
        cmd.Dir = repoDir // Run from the repo root
        return runCommand(cmd)
    }
    ```

### `build` Command

This command uses `bazel test` to compile and test the generated code.

*   **Invocation:** `bazel-generator build`
*   **Inputs:** Reads `/librarian/build-request.json`.
*   **Core Logic (`main.go`):
    ```go
    func buildCmd(ctx context.Context) error {
        // ... reads request file ...

        // Determines the Bazel target from the library's source_paths.
        target := "//" + lib.SourcePaths[0] + "/..." // e.g., //internal/storage/...
        slog.Info("determined bazel target", "target", target)

        args := []string{"test", target}
        cmd := exec.CommandContext(ctx, "bazel", args...)
        cmd.Dir = repoDir
        return runCommand(cmd)
    }
    ```

## How to Build and Run

The generator is packaged as a Docker image that includes Go and the correct version of Bazel.

### Dockerfile

The `Dockerfile` installs all necessary dependencies and sets up the Go application as the entrypoint.

```dockerfile
# bazel/Dockerfile

# Use a MOSS-recommended base image
FROM marketplace.gcr.io/google/debian12:latest

# Install Go, Bazel, and other dependencies
# ... (see Dockerfile for details) ...

# Copy and build the generator source
WORKDIR /app
COPY main.go go.mod ./
RUN go build -o /bazel-generator main.go

# Set the entrypoint
ENTRYPOINT ["/bazel-generator"]
```

### Local Invocation Example

To test this generator locally, you must simulate the environment provided by the Librarian.

**Setup:**

1.  **Build the Docker image:**
    ```bash
    docker build -f bazel/Dockerfile -t go-bazel-generator:latest .
    ```

2.  **Create a mock environment:**
    ```bash
    # Create directories for mounts
    mkdir -p /tmp/mock/librarian /tmp/mock/input /tmp/mock/output

    # You need a mock repo with a Bazel workspace and the generator tool
    # git clone <your-repo-with-bazel-setup> /tmp/mock/repo
    # git clone https://github.com/googleapis/googleapis.git /tmp/mock/source

    # Create the request file
    cp input/configure-request.json /tmp/mock/librarian/generate-request.json
    ```

**Execution:**

This command simulates the Librarian invoking the `generate` command. Note the additional mount for `/repo`.

```bash
docker run --rm \
  -v /tmp/mock/librarian:/librarian:ro \
  -v /tmp/mock/input:/input:ro \
  -v /tmp/mock/source:/source:ro \
  -v /tmp/mock/repo:/repo:ro \
  -v /tmp/mock/output:/output:rw \
  go-bazel-generator:latest generate
```
The command will execute the Bazel target specified in `main.go`, placing the generated artifacts in the `/tmp/mock/output` directory.
