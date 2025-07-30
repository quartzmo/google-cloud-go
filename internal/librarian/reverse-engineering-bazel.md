# Reverse-Engineering Bazel for a Bazel-Free World

One of the primary goals of the `librariangen` project was to replace the legacy, Bazel-based GAPIC generation pipeline with a standalone, containerized Go application. This would improve security, simplify maintenance, and align with modern CI/CD practices. However, a significant obstacle stood in the way: the entire generation process was configured through `BUILD.bazel` files scattered throughout the `googleapis` repository.

This document explains how we overcame this challenge by reverse-engineering the Bazel configuration, allowing us to eliminate our dependency on the Bazel **tool** without immediately needing to replace the Bazel **configuration files**.

## The Key Insight: Separating the Tool from the Configuration

Initially, the problem seemed monolithic: to get rid of Bazel, we had to get rid of `BUILD.bazel` files. The breakthrough came when we made a crucial distinction:

*   **The Bazel Tool:** A complex build and test automation tool that we wanted to stop using for client generation.
*   **The `BUILD.bazel` Files:** Structured text files that, from our perspective, were simply a data source containing the configuration needed to generate a client library.

This insight allowed us to reframe the problem. We didn't need to replace the configuration format right away. We just needed to learn how to *read* it without invoking the Bazel tool. This enabled a phased approach: first, eliminate the tool dependency, and later, migrate away from the configuration format.

## Step 1: Defining the Contract with `protoc.ConfigProvider`

Before we could build a `protoc` command, we needed to know what information was required. By analyzing the `go_gapic_library` rules in various `BUILD.bazel` files, we identified the essential attributes that controlled the generator's behavior.

Consider this `go_gapic_library` rule from `testdata/source/google/cloud/workflows/v1/BUILD.bazel`:

```bazel
go_gapic_library(
    name = "workflows_go_gapic",
    srcs = [":workflows_proto"],
    grpc_service_config = "workflows_grpc_service_config.json",
    importpath = "cloud.google.com/go/workflows/apiv1;workflows",
    metadata = True,
    release_level = "ga",
    rest_numeric_enums = True,
    service_yaml = "workflows_v1.yaml",
    transport = "grpc+rest",
)
```

These attributes map directly to a configuration contract. We formalized this contract in our Go code by creating the [`protoc.ConfigProvider` interface](https://github.com/googleapis/google-cloud-go/blob/main/internal/librariangen/protoc/protoc.go#L28-L38). This interface defined exactly what data we needed to extract, effectively decoupling the `protoc` command builder from the `BUILD.bazel` parser.

Here is the `ConfigProvider` interface from [`protoc/protoc.go`](https://github.com/googleapis/google-cloud-go/blob/main/internal/librariangen/protoc/protoc.go):

```go
type ConfigProvider interface {
	GAPICImportPath() string
	ServiceYAML() string
	GRPCServiceConfig() string
	Transport() string
	ReleaseLevel() string
	HasMetadata() bool
	HasDiregapic() bool
	HasRESTNumericEnums() bool
	HasGoGRPC() bool
}
```

The correspondence is clear: `importpath` maps to `GAPICImportPath()`, `service_yaml` maps to `ServiceYAML()`, and so on.

## Step 2: Reverse-Engineering the `protoc` Command from the Bazel Rule

With the data contract defined, the next step was to figure out how the `go_gapic_library` rule used this data to construct a `protoc` command. The answer lay not in the Bazel tool itself, but in the definition of the rule.

By inspecting the source code of the `go_gapic_library` rule in the [`gapic-generator-go` repository](https://github.com/googleapis/gapic-generator-go/blob/main/rules_go_gapic/go_gapic.bzl), we were able to reverse-engineer the logic. We discovered how each attribute was translated into a `--go-gapic_opt` flag for the `protoc-gen-go_gapic` plugin.

This reverse-engineering effort resulted in our [`protoc.Build` function](https://github.com/googleapis/google-cloud-go/blob/main/internal/librariangen/protoc/protoc.go#L41). It takes a `ConfigProvider` and methodically constructs the exact `protoc` arguments that the Bazel rule would have generated.

For example, the function translates the configuration into the following arguments:

*   `config.GAPICImportPath()` becomes `--go-gapic_opt="go-gapic-package=..."`
*   `config.ServiceYAML()` becomes `--go-gapic_opt="api-service-config=..."`
*   `config.Transport()` becomes `--go-gapic_opt="transport=..."`
*   `config.HasMetadata()` becomes `--go-gapic_opt="metadata"`

This allowed us to replicate the behavior of the Bazel rule with 100% fidelity, without ever calling Bazel.

## Step 3: Implementing a Bazel-Free Parser

The final piece was to implement a parser that could read a `BUILD.bazel` file and provide the data required by the `ConfigProvider` interface.

Our first attempt used the `github.com/bazelbuild/buildtools` library. However, this introduced a new, unnecessary dependency. We quickly realized that since we were only treating the `BUILD.bazel` file as a text-based data source, we didn't need a full-fledged parser.

We replaced the library with a much simpler implementation in [`bazel/parser.go`](https://github.com/googleapis/google-cloud-go/blob/main/internal/librariangen/bazel/parser.go#L130). This new parser first isolates the `go_gapic_library(...)` block and then extracts the values for each required attribute. This approach was not only more lightweight but also eliminated the problematic dependency, reinforcing our core strategy of treating the build files as simple, structured text.

## Conclusion

By distinguishing between the Bazel tool and its configuration files, we were able to devise a pragmatic, incremental migration path. We successfully reverse-engineered the logic embedded in the `go_gapic_library` rule, allowing us to build a standalone Go application that produces bit-for-bit identical output to the legacy system. This project serves as a case study in tackling large-scale system migrations by breaking them down and focusing on interfaces, contracts, and the separation of concerns.