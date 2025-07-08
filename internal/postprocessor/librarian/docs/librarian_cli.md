# librarian CLI: Reimagined

| \#begin-approvals-addon-section See [go/g3a-approvals](http://goto.google.com/g3a-approvals) for instructions on how to add reviewers. Do not edit this section manually. |
| :---: |

**Author:** [Cody Oss](mailto:codyoss@google.com) | **Last Up**Jun 27, 2025**dated**:   | **Status**: Draft   
**Self link:** [go/librarian:cli-reimagined](http://goto.google.com/librarian:cli-reimagined)

# Objective

Take the current [CLI design](http://go/1pp:cli) and distill the commands into two logical CUJs: generate and release.

# Background

In the last couple of weeks the librarian project has added many new contributors to the team. With this some new ideas have been popping up about what the [librarian CLI](https://github.com/googleapis/librarian) currently is and maybe what it could be. There have been many design docs written for what librarian currently is and these docs heavily inspired this one.

The current CLI specification, which can be seen fully in the [appendix](#heading=h.p8csn98ew8e), consists of 8 task oriented commands. These commands orchestrate a process in which they delegate language specific bits to opaque language-specific container images. Currently there are 10 of these container image commands that are invoked by the librarian, which again you can learn more about in the [appendix](#heading=h.8mcr4znlv1d9).

# Overview

This design reimagines the librarian CLI and container contracts while trying to stay true to the overall design philosophies of the original designs. Instead of 8 top-level commands this design proposes two focused on the two main experiences we deal with in the client library space: generate and release.

 Many of the inputs for these commands have also been reworked and simplified. The intention is to provide a simple human user interface for running the CLI locally while at the same time providing a limited set of knobs for more advanced use cases and automation. This design intentionally ignores some of the use cases, like running integration tests, that are covered in the original design. That is not to say these use-cases are not valid. The hope is to start with a simpler baseline and build out more commands and/or integration points in the future as the need arises.

# Detailed Design

## Configuration

All librarian state/config, along with any additional configuration, will be stored in each language repo in a folder called \`.librarian\`. Additionally, any language specific generation input will be in a sub-directory called \`generator-input'. The contents of the \`generator-input\` folder will be mounted into the delegated container(s), to be discussed further down in this document.

### state.yaml

The state file is used to track the current status of all files librarian is managing. It should **never** need to be edited by hand or by any other program than librarian. Other tools that wish to edit this file should do so by passing context to librarian about which fields need to be updated and not try to update these fields directly \-- this will be discussed more later in the design.

```textproto
# The name of the image and tag to use.
image: "gcr.io/my-special-project/language-generator:v1.2.5"

# The state of each library which is released within this repository.
libraries:
  - # The library identifier (language-specific format). api_paths configured under a     
    # given id should correspond to a releasable unit in a given language 
    id: "google-cloud-storage-v1"
    # The last version that was released, if any.
    version: "1.15.0"
    # The commit hash (within the API definition repo) at which
    # the repository was last generated.
    last_generated_commit: "a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2"
    # APIs that are bundled as a part of this library.
    apis:
      - # The API path included in this library, relative to the root
        # of the API definition repo, e.g. "google/cloud/functions/v2".
        path: "google/storage/v1"
        # The name of the service config file, relative to the path.
        service_config: "storage.yaml"
    # Directories to which librarian contributes code to. The root of these paths is the
    # lauguage repository.
    source_paths:
      - "src/google/cloud/storage"
      - "test/google/cloud/storage"
    # Directories files in the local repo to leave untouched during copy and remove.
    preserve_regex:
      - "src/google/cloud/storage/generated-dir/HandWrittenFile.java"
    # If configured, these files/dirs will be removed before generated code is copied 
    # over. A more specific `preserve_regex` takes preceidece. If not not set, defaults
    # to the `souce_paths`.
    remove_regex:
      - "src/google/cloud/storage/generated-dir"
```

### env-config.yaml {#env-config.yaml}

In contrast to the state file, the environment config file **will** be edited by humans and live in Piper, alongside Kokoro configuration. The values stored in this file will be passed to librarian via a flag. Librarian will then source any secrets if needed and pass them along to the configured container.

```textproto
# A project to source any secrets from.
secret_project: "my-project"
# Specific configuration for each individual command.
commandEnvironmentVariables:
  # Common values are passed to all container commands
  common:
  - name: SOURCELINK_REPO
    default: https://github.com/googleapis/google-cloud-dotnet
  # Config for the 'generate' container command, but can be any command
  generate:
    # The environment variables to populate for this command.
    environment_variables:
      - # The name of the environment variable (e.g. TEST_PROJECT).
        name: "GENERATOR_FLAGS"
        # The default value to specify if no other source is found for the environment
        # variable. If this is not provided and no other source is found, the environment
        # variable will not be passed to the container at all.
        default_value: "--clean --new-clients"
      - name: "API_KEY"
        # The name of the secret to be used to fetch the value of the environment
        # variable when it's not present in the host system. If this is not specified,
        # or if a Secret Manager project has not been provided to Librarian,
        # Secret Manager will not be used as a source for the environment variable.
        secret_name: "generator-api-key" # Fetched from Secret Manager
        default_value: "" # No default, will be empty if secret not found

```

### generator-input folder

This folder can contain any extra configuration needed to help generate a more formalized library. This may contain any tweaks/hacks that we want to apply to the codebase at generation time. It should be noted that it **must** be possible to generate a library without this folder existing for our API producer CUJs. Libraries generated in this manner must be able to compile/run, but may not meet our normal standard for release-ready, which is okay for this CUJ. We should only use this extra configuration to apply patches to our externally facing libraries.

### Environment Variables

Librarian will have some well-defined environment variables it will honor. All librarian specific environment variables will be prefixed with \`LIBRARIAN\_\`. Additionally librarian will also honor some kokoro specific environment variables when processing data in CI. The current list of these variables that are honored are:

| name | description |
| :---- | :---- |
| LIBRARIAN\_GITHUB\_TOKEN | Specifies the access token to use for all operations involving GitHub. Because we don't want it shown in the command line output log, we will use a file to pass this value to the container. |

## Generate command

The configure, generate, and update-apis commands will be simplified to one command: generate. This command will have the following inputs:

| generate flags | description | required |
| :---- | :---- | :---- |
| repo | Code repository where the generated code will reside. Can be a remote in the format of a remote URL such as \`https://github.com/{owner}/{repo} or a local file path like \`/path/to/repo\`. Both absolute and relative paths are supported. | true |
| image | Language specific image used to invoke code generation. If not specified, use an image configured in the \`state.yaml\`. | false |
| library | The library ID to generate (e.g. secretmanager). This corresponds to a releasable language unit \-- for Go this would be a Go module or for dotnet the name of a NuGet package. If not specified all currently managed libraries will be regenerated. | false |
| api | Relative path to the API to be configured/generated (e.g., google/cloud/functions/v2). This value must be specified when generating a library for the first time and may be specified to target a specific APIs regeneration. | false |
| api-source | location of api repository. If undefined, googleapis is the default and will be cloned to the output | false |
| output | working directory root. When this is not specified, a working directory will be created in /tmp. | false |
| push-config | If specified, will try to create a commit and pull request for the generated changes. The format should be \`”{git-email-address},{author-name}”\`. Also, when this field is specified it is expected a Github token will be provided with push access via the environment variable \`LIBRARIAN\_GITHUB\_TOKEN\`. | false |
| build | If true, will attempt to delegate a build/test job to the container. If the container returns a non-zero exit code the CLI command fails. | false |
| env-config | The environment config is a reference to a file that contains environment variables that should be made available to certain container invocations as [defined above](#env-config.yaml). This file will exist in Piper and be sourced by librarian. | false |
| host-mount | When Librarian is already running within Docker, this is a value of the form \<host-mount\>:\<local-mount\> where host-mount is a mount point from the perspective of the Docker host, and local-mount is how that same mount point appears within Docker (i.e. where Librarian will see it).When this is specified, Librarian must validate that any mount points *it* provides to language containers are within this local-mount, and replace mount sources to refer to host-mount, so that the language container "sees" the right directory. When running in Kokoro jobs, this is expected to be specified as ${KOKORO\_HOST\_ROOT\_DIR}:${KOKORO\_ROOT\_DIR} | false |
| LIBRARIAN\_GITHUB\_TOKEN | Specifies the access token to use for all operations involving GitHub. Because we don't want it shown in the command line output log, we will use a file to pass this value to the container. | false, required only if opening a PR |

By default, this command is just meant to run locally while only supplying the required flags. If the \`push-config\` flag is set the tooling will try to create a commit and pull request for the generated changes. This is meant to be both engineer and machine friendly \-- it should be easy to automate. The generated commits from librarian will forward the commits messages landed in googleapis to the language repository \-- just like owl-bot does today. For example, this [CLs public message](https://critique.corp.google.com/cl/775377248) is landed into [googleapis](https://github.com/googleapis/googleapis/commit/d2835e84647d7477511f7ae48e36d4cfe7b04a10) and then into a [language repository](https://github.com/googleapis/google-cloud-go/pull/12494). This is to facilitate the context of the commits landing in the release notes for a given library.

### Container Contracts

The current librarian design uses a container delegation model, and this design is no different. For \`generate\` this means the actual language specific code used to configure/generate/build the libraries will be codified as a Dockerfile and invoked by librarian in different ways.

#### configure container command {#configure-container-command}

It is the configure container…

| context | type | description |
| :---- | :---- | :---- |
| /librarian | mount (read/write) | This mount will contain exactly one file named \`configure-request.json\`. The container is expected to process this request and write a file back to this mount named \`configure-response.json\`. Both of these files use the schema of a library defined above in the state file. The container may wish to add more context to the library configuration which it expresses back to librarian via this message passing. It will then be librarians responsibility to commit these changes to the state.yaml file in the language repository. |
| /input | mount (read/write) | The exact contents of the generator-input folder, e.g. google-cloud-go/.librarian/generator-input. This folder has read/write access to allow the container to add any new language specific configuration required. |
| /source | mount (read) | This folder is mounted into the container. It contains, for example, the whole contents of [googleapis](https://github.com/googleapis/googleapis). This will be needed in order to read the service config files and likely also the BUILD.bazel files that hold a lot of configuration today. |
| command | Positional Argument  | The value will always be \`configure\` for this invocation. |

In order for the the container to have enough context on how and what to generate, librarian will provide the container the following context:

#### generate container command {#generate-container-command}

| context | type | description |
| :---- | :---- | :---- |
| /librarian/generate-request.json | mount (read) | A JSON file that describes which library to generate.  |
| /input | mount (read/write) | The exact contents of the generator-input folder, e.g. google-cloud-go/.librarian/generator-input. Example contents are: \- Templates ([example](https://github.com/googleapis/google-cloud-dotnet/blob/main/generator-input/README-template.md)) \- Logic to run tweaks ([example](https://github.com/googleapis/google-cloud-dotnet/blob/main/generator-input/tweaks/Google.Cloud.AIPlatform.V1/pregeneration.sh))  This folder has read/write access to allow the container to add any new language specific configuration required. |
| /output | mount (write) | This folder is mounted into the container. It is meant to be the destination for any code generated by the container. Its output structure should match that of where the code should land in the resulting repository. For example if we are generating the [secretmanger v1](https://github.com/googleapis/google-cloud-go/tree/main/secretmanager/apiv1) client for Go, we would write files to \`/output/secretmanager\`. |
| /source | mount (read) | This folder is mounted into the container. It contains, for example, the whole contents of [googleapis](https://github.com/googleapis/googleapis). This will be needed in order to read the service config files and likely also the BUILD.bazel files that hold a lot of configuration today. |
| command | Positional Argument  | The value will always be \`generate\`\` for this invocation. |

These language containers will be invoked once per-library configured with librarian.

In addition to the “generate” container command, if the \`build\` flag is specified during generation librarian will invoke the container image again in “build/test” mode. During execution, the container is expected to try to compile/unit-test/etc to make sure that the generated code is functional.

#### build container command {#build-container-command}

| context | type | description |
| :---- | :---- | :---- |
| /librarian | mount (read) | The exact contents of the \`.librarian\` folder in the language repository. Additionally this will contain a file name \`build-request.json\` describing the library being processed. |
| /repo | mount (read/write) | The whole language repo. The mount is read/write to make diff-testing easier. Any changes made to this directory will have no-effect on the generated code, it is a deep-copy. |
| command | Positional Argument  | The value will always be \`build\` for this invocation. |

### Commits and pull requests

TODO: Once we decide if we will have contextual commit or not

### Example CLI commands

#### Generate a new library

Generate a library for the first time or regenerate a specific api.

```
librarian generate -repo=googleapis/google-cloud-go -library=secretmanager -api=google/cloud/secretmanager/v1
```

#### Regenerate a single library

```
librarian generate -repo=googleapis/google-cloud-go -library=secretmanager
```

#### Regenerate all libraries

```
librarian generate -repo=googleapis/google-cloud-go
```

#### Regenerate all libraries, create commit, and open PR

This CUJ is mostly intended to be invoked by CI using bot credentials.

```
LIBRARIAN_GITHUB_TOKEN=xxx librarian generate -repo=googleapis/google-cloud-go -push-config=”codyoss@google.com,Cody Oss”
```

### 

### Example life of a \`librarian generate\` command

* Something/Someone wants to generate a client library for an API that has not been onboarded before. 

```
librarian generate -repo=googleapis/google-cloud-go -library=secretmanager -api=google/cloud/secretmanager/v1 -build -push-config="something,something"
```

* librarian gathers all of its flags and environment variables.  
* librarian clones the language repo and googleapis  
* librarian synthesizes a state.yaml for the new API  
* librarian delegates configuration of state file modification and generator-input to the language container. librarian follows the [configure contract](#configure-container-command).  
* librarian will update the state for the library based on the contents of state.json.  
* librarian delegates generation of the code to a language container specified in the state.yaml. librarian follows the [generate contract.](#generate-container-command)  
* librarian will update the state for the library based on the contents of state.json.  
* librarian will copy over the generated artifacts to the cloned language repo respecting the rules defined by the state.yaml file for the library.  
* librarian will then try to build the code because the \`build\` flag is present. librarian follows the [build contract](#build-container-command).  
* librarian will commit and push the code using the \`push-config\` provided as well as the token provided via the environment .

## Release command

TODO

Notes

* I think for compatibility sake we need to fully support release-please, initially. It has a lot of configuration options. I think we do this by containerizing release-please and feeding it all the options required from the librarian. It should be easy enough to create a shim layer for translation as long as we mount the whole repository into the container. This would mean the release-please tool holds the logic for creating the pull requests.  
  * [https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md)  
* Eventually I could see us wanting to get that logic directly into librarian for more control in one place, but release-please has a lot of custom logic per language today that we would need to figure out what to do with: [https://github.com/googleapis/release-please/tree/main/src/strategies](https://github.com/googleapis/release-please/tree/main/src/strategies)  
* Call out shared flags between this and generate.

## Large Difference from the current design

The following are a list of differences from the current design [CLI/container design](https://docs.google.com/document/d/1i2H1xpgwN_e9G2EHd2dJCw_Uk72FxbkKnkAcl3A7fmc/edit?usp=sharing):

* Configure/generate/update-apis are condensed into one command. In order to do this I do think configure is the one command and flow changes the most. With the above design as written I expect configure to no longer scan service.yamls to see if it should purpose new libraries. Onboarding is a more manual process.  
*  Maybe we eventually need a \`state edit\` command that has editing features (e.g. [go mod edit](https://go.dev/ref/mod#go-mod-edit)). This could be used to update the state file but keep it completely under librarian. The use cases:  
  * container image revision update  
* Containers no longer edit state file directly, they do so implicitly by passing state via JSON.  
* Most of the flags are the same, but some have been renamed and the language flag has been removed. The only valuable CUJ I saw for the language flag was API producers. I think we can just have good docs for what to set the repo value to instead in that workflow.  
* I have abridged the config files and added a new preserve feature to the config to make the design more compatible with owlbots offerings. Also, by doing this I think…  
* We no longer need a clean container command. Maybe I don’t fully understand why this was needed though.  
* Removed integration testing.  I feel pretty strongly that this testing should be a feature of the language repository and not librarian.  
* The current design maintains the last generated (and released) API commit per library; the new design only does this for the last generated commit.  
* The current design keeps the language container tag as part of the Librarian-managed state; the new design moves it into config (with a corresponding responsibility change, presumably)

# Testing Plan

TODO: something something, create a test github repo that we use for end-to-end conformance testing of the librarian contracts.

# Alternatives considered

The original design: [CLI/container design](https://docs.google.com/document/d/1i2H1xpgwN_e9G2EHd2dJCw_Uk72FxbkKnkAcl3A7fmc/edit?usp=sharing).

# Work Estimates

∞

# Documentation plan

For now the CLI will be self documenting with its help text. Once things have stabilized we should create a process that uploads dumps of help text to go/client-home.

# Launch plans

🚀

# Risks

* [Librarian for .NET Design](https://docs.google.com/document/d/1Fy95DOCwgSzYSMvuBhgRs_hk7Jli9JqxBHtF8zUnZ9Y/edit?usp=sharing) is already implemented and running in production. Any changes we make here will mean rework and potential feature gaps for dotnet.  
* Similar to the item above, [Python Librarian Migration (go/sdk-librarian-python)](https://docs.google.com/document/d/1aF1zVz7VMr5KuQEZW2JzRySOcXmk9IjHDb9J0ckSeVk/edit?usp=sharing) was written against the current CLI and container contract design. This is much less of a risk than dotnet as less work has been designed and done.  
* Shifting designs may affect currently communicated timelines.  
  * Specifically for our first languages being onboarded, we need to lock down our container contracts by EoD Jul 2, 2025\.

## Decisions that need to be made

* Will we keep conventional commits and commit message propagation?  
  * Separate commit per-library?  
* Do we support multiple state files, one per-library?  
* Can we include a copyright year in the state file. This would make onboarding for Rust and C++ easier.

