# Go LIbrarian Migration

*\#begin-approvals-addon-section*

**Author(s):** [Chris Smith](mailto:chrisdsmith@google.com)| **Last Updated**: Jul 10, 2025  | **Status**: Draft   
**Self link:** [go/librarian:go-migration](https://goto.google.com/librarian:go-migration)  | **Project Issue**: [b/425766855](http://b/425766855), [b/430109429](http://b/430109429)

# Objective

Onboard Go GAPIC client library generation to Librarian 0.2 per the contract outlined in [librarian CLI: Reimagined](go/librarian:cli-reimagined) and [Librarian: Onboarding Guide](go/librarian:language-onboarding).

The scope of this effort is to onboard a single straightforward, typical client, such as the Google Cloud Asset client shown in the examples below. The generated assets must be perfectly backward-compatible with the existing published assets for the client in the google-cloud-go repo, including the snippets files located outside of the client module in google-cloud-[go/internal/generated/snippets](TODO). One exception to backward compatibility is the update of shared files

# Background

The current code generation process is a multi-staged pipeline that relies on the coordination of several distinct external tools and bots:

1. **Generation:** [bazel-bot](https://github.com/googleapis/repo-automation-bots/tree/main/packages/bazel-bot) automatically [invokes](https://github.com/googleapis/repo-automation-bots/blob/main/packages/bazel-bot/docker-image/generate-googleapis-gen.sh) Bazel to generate source code (using [protoc](https://protobuf.dev/getting-started/gotutorial/) and [gapic-generator-go](https://github.com/googleapis/gapic-generator-go)) from the public API definitions repo ([googleapis](https://github.com/googleapis/googleapis)) to the private target repo ([googleapis-gen](https://github.com/googleapis/googleapis-gen)).  
2. **File placement:** After generation, [OwlBot](http://go/owlbot-howto#githubowlbotyaml) automatically copies the source code to the [google-cloud-go](https://github.com/googleapis/google-cloud-go) repo. This operation deletes existing files and then copies the generated files per the deep-remove-regex and deep-copy-regex rules declared in google-cloud-go’s single [.OwlBot.yaml](https://github.com/googleapis/google-cloud-go/blob/main/.github/.OwlBot.yaml) file. Some existing files are preserved via deep-preserve-regex rules.  
3. **Post-processing:** Finally, OwlBot runs the google-cloud-go [OwlBot post-processor](https://github.com/googleapis/google-cloud-go/tree/v0.121.3/internal/postprocessor) to add additional files and perform cleanup steps that are not currently handled in `protoc` generation. (This is typically necessary due to incomplete configuration in the [googleapis](https://github.com/googleapis/googleapis) repo, but may also be done simply for historical reasons, since protoc and [gapic-generator-go](https://github.com/googleapis/gapic-generator-go) were sometimes perceived as having a narrower role.)

The bazel-bot/OwlBot toolchain architecture has been determined to be vulnerable to security and maintenance concerns, as documented in the [1PP Engineering Plan](https://docs.google.com/document/d/1_iLARXR7ZEJALSyXUjdFWhFbSJq9GCDhu1BUUIPKLkI/edit).

The Librarian project replaces the bazel-bot/OwlBot toolchain.

# Overview

This document outlines a phased approach for migrating Go client library generation to [Librarian 0.2](http://go/librarian:cli-reimagined). The core strategy for incremental migration is to adapt existing tools and inputs to a containerized (Docker) workflow. This means that the new Librarian workflow will continue to use parts of the existing toolchain, including:

* The googleapis repo, including Protobuf definitions and Bazel configuration files (but not the Bazel tool itself).  
* protoc  
* gapic-generator-go  
* Some configuration files in google-cloud-go.

New components include:

* A Docker container will all required dependencies (Go, protoc, etc).  
* A new Go program, in the container, fulfilling the Librarian container contract.  
* An adaptation of the Go OwlBot post-processor within this program.  
* Tests covering all of the above.  
* Documentation covering all of the above.

Deferring some architectural changes, such as replacing the Bazel configuration files, allows us to focus on an accelerated delivery schedule and corresponding technical risk reduction. Work estimates assume that there will be no change to the release tooling. If there are changes to the release process which will block the initial onboarding, then we will need to increase the overall level of effort.

# Detailed Design

This design modernizes our code generation pipeline by removing dependencies on the [googleapis-gen](https://github.com/googleapis/googleapis-gen) repo and [bazel-bot](https://github.com/googleapis/repo-automation-bots/tree/main/packages/bazel-bot). This change aligns with the updated standards in the [Librarian: Scrappy Onboarding Guide](https://docs.google.com/document/d/1_A8O-UE-RKPPG3_4bK4gDQJBPlrcq28rctJSohXIjCY/edit?resourcekey=0-UsQlfcvhjNmbhtaqms52Tg&tab=t.0#heading=h.xgjl2srtytjt).

This document describes the Go GAPIC Librarian migration. It proposes a containerized (Docker) Go build tool delivering backward-compatible Go GAPIC client library generation managed by the Librarian pipeline. Its responsibility is to generate backward-compatible and release-ready Go GAPIC client libraries based on the existing [googleapis](https://source.corp.google.com/piper///depot/google3/third_party/googleapis/stable/) Protobuf API definitions and the co-located Bazel configuration files.

The central Librarian tool is responsible for orchestrating the overall workflow, which includes:

1. Reading configuration from a `state.yaml` file in the target language repository.  
2. Preparing the necessary inputs (API source files, configuration) and a set of mounted directories.  
3. Invoking the Docker container with a specific command (generate).  
4. Processing the output of this container (e.g., copying generated code, updating state).

The Docker file entrypoint is straightforward: The main function in the cloud.google.com/internal/librarian/generator module determines which command to run based on the first argument passed to the container. There are three primary commands: `configure`, `generate`, and `build`. Currently, only `generate` is to be implemented.

## The generate command

The generate command is Librarian’s core command, responsible for the actual code generation. The `generate` command handler orchestrates the generation logic:

1. **Read inputs:** It reads and unmarshals the `/librarian/generate-request.json` file. This file contains the definition of the library to be generated.

2. **Invoke protoc:** 

3. **Invoke post-processor:** The

### Read inputs

The 

##### Generate inputs

| Context | Type | Description |
| :---- | :---- | :---- |
| `/librarian/generate-request.json` | Mount (Read) | A JSON file that describes which library to generate. Its schema is a subset of the `libraries` entry in `state.yaml`. |
| `/input` | Mount (Read) | The contents of the `.librarian/generator-input` folder from the language repository. This can contain shared or language-specific files needed for generation. |
| `/output` | Mount (Write) | An empty directory where the container must write all generated code. The directory structure should match the desired structure in the final repository. |
| `/source` | Mount (Read) | A directory containing the API definitions (e.g., a clone of `googleapis/googleapis`). |
| `generate` | Positional Arg | The command that invokes this operation. |

### 

##### generate-request.json (simple)

```json
{
  "id": "google-cloud-asset",
  "apis": [
    {
      "path": "google/cloud/asset/v1",
      "service_config": "cloudasset_v1.yaml"
    }
  ]
}
```

See the complete set of protos and related BUILD.bazel configuration for the Asset API at [googleapis/google/cloud/asset](https://github.com/googleapis/googleapis/tree/master/google/cloud/asset/).

However, more complex configurations do exist, such multiple versions of the same API as well as nested APIs.

##### generate-request.json (nested APIs)

```json
{
  "id": "google-cloud-workflows",
  "apis": [
    {
      "path": "google/cloud/workflows/v1",
      "service_config": "workflows_v1.yaml"
    },
    {
      "path": "google/cloud/workflows/v1beta",
      "service_config": "workflows_v1beta.yaml"
    },
    {
      "path": "google/cloud/workflows/executions/v1",
      "service_config": "workflowexecutions_v1.yaml"
    },
    {
      "path": "google/cloud/workflows/executions/v1beta",
      "service_config": "workflowexecutions_v1beta.yaml"
    }
  ]
}

```

See the complete set of protos and related BUILD.bazel configuration for the Workflows API at [googleapis/google/cloud/workflows](https://github.com/googleapis/googleapis/tree/master/google/cloud/workflows/).

### Invoke protoc

For each API listed in the request, it constructs and executes a `protoc` command with the [GAPIC protoc plugin](https://github.com/googleapis/gapic-generator-go/tree/v0.53.1/cmd/protoc-gen-go_gapic) to generate the Go client library and supporting files. A typical `protoc` invocation would be:

```shell
# Dump of the protoc command for google/cloud/asset/v1
TODO:
	
```

The generated files are written directly to the `/output` directory. They should match the current [OwlBot](https://github.com/googleapis/repo-automation-bots/blob/main/packages/owl-bot/README.md) output in the private repo [googleapis/googleapis-gen](https://github.com/googleapis/googleapis-gen). For the limited, v1\-only, simple example google-cloud-asset, above, the output should look as follows. The generated output for more complex configurations such as google-cloud-workflows should be similar, only broader. 

```shell
cloud.google.com/
└── go
    ├── asset
    │   └── apiv1
    │       ├── asset_client_example_go123_test.go
    │       ├── asset_client_example_test.go
    │       ├── asset_client.go
    │       ├── assetpb
    │       │   ├── asset_enrichment_resourceowners_grpc.pb.go
    │       │   ├── asset_enrichment_resourceowners.pb.go
    │       │   ├── asset_service_grpc.pb.go
    │       │   ├── asset_service.pb.go
    │       │   ├── assets_grpc.pb.go
    │       │   └── assets.pb.go
    │       ├── auxiliary.go
    │       ├── auxiliary_go123.go
    │       ├── doc.go
    │       └── helpers.go
    └── internal
        └── generated
            └── snippets
                └── asset
                    └── apiv1
                        ├── Client
                        │   ├── AnalyzeIamPolicy
                        │   │   └── main.go
                        │   (And so on..) 
                        │   └── UpdateSavedQuery
                        │       └── main.go
                        └── snippet_metadata.google.cloud.asset.v1.json

35 directories, 38 files

```

## Usage

The generator is packaged as a Docker image.

### Dockerfile

The `Dockerfile` is designed to be MOSS-compliant and self-contained.

```
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

1. **Build the Docker image:**

```shell
docker build -t go-generator:latest .
```

2. **Create a mock environment:**

```shell
# 1. Create directories to mount
mkdir -p /tmp/mock/librarian /tmp/mock/input /tmp/mock/output
git clone https://github.com/googleapis/googleapis.git /tmp/mock/source

# 2. Create the request file (using the mock input provided)
cp ../input/configure-request.json /tmp/mock/librarian/generate-request.json
```

**Execution:**

This command simulates the Librarian invoking the `generate` command.

```shell
docker run --rm \
  -v /tmp/mock/librarian:/librarian:ro \
  -v /tmp/mock/input:/input:ro \
  -v /tmp/mock/source:/source:ro \
  -v /tmp/mock/output:/output:rw \
  go-generator:latest generate
```

After the command completes, the `/tmp/mock/output` directory will contain the generated Go files, just as it would in the real pipeline.

### 1\. Execution Environment

The entire generation process will be encapsulated within a Docker container.

#### 1.1 Dockerfile

* **Base Image:** We will use a [Google-approved base image](https://cloud.google.com/software-supply-chain-security/docs/base-images#google-provided_base_images) (Ubuntu 24.04) to ensure it meets security and compliance standards.  
* **Tooling:** The following essential command-line tools will be installed in the image:

```shell
protoc # 5 repositories include gencode. This will only be used for those repos
bazelisk # for running the Go GAPIC generator
wget # for fetching dependencies such as bazelisk and protoc
zip # for unzipping the tar file after running bazelisk
```

### 

#### 1.2 Go Dependencies

All Go packages will be managed via pip-compile and installed from Google’s internal “[Airlock](http://go/airlock/howto_pypi#pypi-configuration)” repository to ensure stability and control over dependencies.

* **Dependency declaration**(requirements.in):

```shell
# START dependency of gapic-generator-go

```

* **Secure Installation Method:** We will enforce hash-checking for all installed packages (ensuring that the dependencies are using [airlock](http://go/airlock/howto_pypi#pypi-configuration), rather than PyPI). The installation process within the Dockerfile will consist of two steps:  
  1. Generate a requirements.txt file with pinned versions and hashes:

```shell
pip-compile --generate-hashes requirements.in
```

  2. Install the packages using the generated file:

```shell
pip install --require-hashes -r requirements.txt
```

#### Generator input

All configuration for the new Librarian workflow will be centralized within a .librarian directory at the root of gapic-generator-go:

```shell
.librarian/
└── generator-input/
      └── library-api-paths-mapping.json
      └── client-post-processing/
         └── # (Scripts that currently reside in <repo>/scripts/client-post-processing) 


```

##### Generation Process Walkthrough

The following steps outline the generation process which takes place in the Go generator command CLI:

### 3\. Migration Considerations

#### 3.1 Handling Logic from .Owlbot.yaml

Some existing .Owlbot.yaml files contain more than just file paths; they include conditional logic for the generation process. This embedded logic must be carefully migrated to the appropriate post-processing scripts which will be located in .librarian/generator-input/client-post-processing.

#### 3.2 Handwritten Files Preservation Strategy

It is critical that handwritten code and important metadata are not overwritten during code generation. The following files and directories will be configured to be preserved.

* packages/\<package\>/tests/system/\*\*  
* packages/\<package\>/.repo-metadata.json  
* packages/\<package\>/CHANGELOG.md  
* packages/\<package\>/.OwlBot.yaml  
  * **Note:** This file will be preserved temporarily. The goal is to minimize the “diff” when migrating to Librarian, making the transition easier to review. It can be removed later.

##### Further Investigation Required:

The following paths are often symlinks, not actual files. We must verify that the Librarian tooling supports preserving symlinks.

* packages/\<package\>/docs/CHANGELOG.md  
* packages/\<package\>/docs/README.rst  
* packages/\<package\>/scripts/client-post-processing/\*\*  
  * **Note:** Preserving this symlink is not a hard requirement because we can recreate the new necessary links using scripts in the .librarian/generator-input/client-post-processing/ folder if needed.

# Alternatives considered

As mentioned in [go/sdk-librarian-python](http://goto.google.com/sdk-librarian-python), we considered completely removing the dependency on Bazel and Synthtool before migrating to the Librarian workflow. However, this alternative was rejected due to the substantial, time consuming effort and complexity it would introduce compared to the chosen approach (i.e. migrating to Librarian with existing tools and then replacing them within the new framework) due to the following reasons:

* **Large-Scale Manual Review:** Migrating all 200+ packages off of Bazel and Synthtool at once would generate massive code changes. The manual review effort required to safely validate these changes would be immense and impractical.  
* **Required Post-Processor Changes:** The existing post-processor is tightly coupled with the current tooling and decoupling it would require extra work.  
* **Synthtool Template Migration:**  Code generation templates currently managed in Synthtool would need to be migrated to the Go GAPIC generator, which is a complex project on its own.  
* **Metadata Source Relocation:** The source of truth for repo-metadata.json would need to be moved.

# Documentation plan

Internal process documentation for Python will be updated to remove references to the Owlbot CLI and Owlbot Bootstrapper which are no longer used. The documentation will in turn be updated with the new Librarian Workflow.

# Testing plan

## Manual E2E Testing

As part of Phase 1, we will be running the librarian CLI to verify that we’re able to successfully regenerate (i.e. generate and build) \~10 libraries without a diff.

* Manually update the language specific state.yaml file for 10 libraries.  
* Run the librarian CLI generate to regenerate these libraries and run nox session tests.  
* Verify that there is no diff.

As part of Phase 2, we will be running the Librarian CLI to verify that we’re able to successfully onboard new libraries i.e. (configure, generate, and build) \~180 libraries without a diff.

* Use a script to run the librarian CLI generate command  against a list of library ids and api-paths.  
* Verify that there is no diff.  
* Verify that the state.yaml file is populated and has all the necessary information for the onboarded libraries.

For the most part, this manual testing will verify that onboarding to Librarian is successful.

## Smoke Testing

We will test the Dockerfile and ensure that we’re able to build it successfully. We will run this build and ensure that all the required tools and Python dependencies are part of this image. All of this can be done via a bash script. 

Here’s an example:

```shell
#!/bin/bash
# Exit immediately if any command fails.
set -e

IMAGE_NAME="python-generator:test"

echo "--- Building Docker image for smoke test ---"
docker build -t $IMAGE_NAME .

echo "--- Running smoke tests inside the container ---"

# This command runs a series of checks inside the container.
# If any check fails, 'set -e' will cause the script to exit with an error.
docker run --rm $IMAGE_NAME /bin/bash -c "
  echo 'Checking for required tools...'
  which protoc
  which bazelisk
  which python3

  echo 'Checking for key Python dependencies...'
  pip list | grep -F 'click'
  pip list | grep -F 'google-api-core'
"
```

This script can be manually run or be part of a Github Action Workflow (for automation).

## Docker based E2E Testing

An end to end test is considered successful if:

1. The generate command orchestrates all its steps such as bazelisk build, tar, synthtool, etc.  
2. The resulting generated library passes its own test suite when the build command is run.

This approach ensures that our tests are robust and meaningful, validating that our tooling works as expected without failing on any cosmetic changes from the generator.

We will execute a script (e.g. run-[container-tests.sh](http://container-tests.sh)) that performs the following steps for each test case:

1. Build the docker test image to ensure a clean environment.  
2. Run generate command inside the container for a specific test library.  
3. After generate succeeds, execute the build command against the newly generated code in a container,

Successful completion of these e2e tests confirm that our workflow correctly produced a valid and functional library.

We will maintain a separate file for a test suite, covering any necessary edge cases.

**Note:** We can probably combine the smoke test and the e2e test into a single script but it is okay to keep these as separate sections for the purpose of this design.

## Unit Testing

Any Python code that we write as part of this project will be accompanied by unit tests such as the language container CLI tool.

# Risks

Hidden logic in Bazel rules

Updating shared files

internal/.repo-metadata-full.json

## Hidden Logic in Configuration Files

There may be subtle hidden logic in .Owlbot.yaml which may be undocumented or not easy to identify. Migrating this logic to post-processing scripts could be more time-consuming than anticipated.

### Mitigation

Before starting the full migration, allocate time to specifically perform an audit of .Owlbot.yaml to categorize the types of logic we will encounter and estimate the effort more accurately.
