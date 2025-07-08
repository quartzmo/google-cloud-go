# Librarian: Scrappy Onboarding Guide

**Authors**: [Cody Oss](mailto:codyoss@google.com)  
**Last Updated**: May 29, 2025

# **Objective**

Create a bare minimum \`generate\` onboarding guide for Go/Python. This will later be turned into a real guide in a separate doc.

# **Information**

## **Contracts**

### **state.yaml**

All language repos will have a single state.yaml file located in their language repo at the location \`.librarian/state.yaml\`. This file lets librarian know which libraries it is responsible for generating.

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
    # Directories to which librarian contributes code to.
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

### **Container contracts**

When librarian orchestrates the generation flow it does so by making up to three invocations to a language container image. Each of these invocations is expected to do slightly different tasks given slightly different inputs. Successful container invocation runs are expected to return a 0 exit code, any non-zero is considered an error and will break the flow. When defining a Dockerfile for your container make sure to set a binary entrypoint so the librarian can properly pass any arguments needed. The processing unit of each container invocation in this flow is always the library unit, which may contain multiple apis in some languages. The three container contracts are as follows:

#### Configure container command

The configure container is only invoked during the onboarding process of a new API. The container is expected to use all of the mounted input to produce up-to two different artifacts.

1. The container is expected to return a \`configure-response.json\`, which is derived from the \`configure-request.json\` as well as language specific knowledge. The response will be committed back to the state file by librarian. The \`configure-request.json\` is a library view of one library configuration. This uses the exact same schema as the items under \`libraries\` in the state.yaml file above.  
2. If the language needs to generate any sort of “side configuration” for libraries, the container may do so at this step. This configuration should be written to the \`input\` mount which will correspond back to the \`generator-input\` folder in the language repository.

| context | type | description |
| :---- | :---- | :---- |
| /librarian | mount (read/write) | This mount will contain exactly one file named \`configure-request.json\`. The container is expected to process this request and write a file back to this mount named \`configure-response.json\`. Both of these files use the schema of a library defined above in the state file. The container may wish to add more context to the library configuration which it expresses back to librarian via this message passing. It will then be librarians responsibility to commit these changes to the state.yaml file in the language repository. |
| /input | mount (read/write) | The exact contents of the generator-input folder, e.g. google-cloud-go/.librarian/generator-input. This folder has read/write access to allow the container to add any new language specific configuration required. |
| /source | mount (read) | This folder is mounted into the container. It contains, for example, the whole contents of [googleapis](https://github.com/googleapis/googleapis). This will be needed in order to read the service config files and likely also the BUILD.bazel files that hold a lot of configuration today. |
| command | Positional Argument  | The value will always be \`configure\` for this invocation. |

An example configure-request might look like:

```json
{
  "id": "google-cloud-storage-v1",
  "apis": [
    {
      "path": "google/storage/v1",
      "service_config": "storage.yaml"
    }
  ],
}
```

And an example configure-output might look like this:

```json
{
  "id": "google-cloud-storage-v1",
  "version": "0.0.0",
  "apis": [
    {
      "path": "google/storage/v1",
      "service_config": "storage.yaml"
    }
  ],
  "source_paths": [
    "src/google/cloud/storage",
    "test/google/cloud/storage"
  ],
  "preserve_regex": [
    "src/google/cloud/storage/generated-dir/HandWrittenFile.java"
  ],
  "remove_regex": [
    "src/google/cloud/storage/generated-dir"
  ]
}
```

#### generate container command

The generate container command is the step where the most work happens \-- the code generation. The container is expected to write output to the \`output\` mount in the same directory structure as the code should land in the language repository. Like the \`configure\` command above, it is also given read/write access to the \`input\` directory. Unlike the \`configure\` command you will only receive read access to the \`generate-request.json\` message. Again, this message uses the same schema as described above.

| context | type | description |
| :---- | :---- | :---- |
| /.librarian/generate-request.json | mount (read) | A JSON file that describes which library to generate.  |
| /input | mount (read/write) | The exact contents of the generator-input folder, e.g. google-cloud-go/.librarian/generator-input. Example contents are: \- Templates ([example](https://github.com/googleapis/google-cloud-dotnet/blob/main/generator-input/README-template.md)) \- Logic to run tweaks ([example](https://github.com/googleapis/google-cloud-dotnet/blob/main/generator-input/tweaks/Google.Cloud.AIPlatform.V1/pregeneration.sh))  This folder has read/write access to allow the container to add any new language specific configuration required. |
| /output | mount (write) | This folder is mounted into the container. It is meant to be the destination for any code generated by the container. Its output structure should match that of where the code should land in the resulting repository. For example if we are generating the [secretmanger v1](https://github.com/googleapis/google-cloud-go/tree/main/secretmanager/apiv1) client for Go, we would write files to \`/output/secretmanager/apiv1\`. |
| /source | mount (read) | This folder is mounted into the container. It contains, for example, the whole contents of [googleapis](https://github.com/googleapis/googleapis). This will be needed in order to read the service config files and likely also the BUILD.bazel files that hold a lot of configuration today. |
| command | Positional Argument  | The value will always be \`generate\`\` for this invocation. |

After this container is invoked it is librarians responsibility to copy the generated code back to the language repository and perform any merging/deleting actions as defined in the libraries state entry.

#### build container command

Lastly, the \`build\` container commands task is to build/test the generated library.

| context | type | description |
| :---- | :---- | :---- |
| /librarian | mount (read) | The exact contents of the \`.librarian\` folder in the language repository. Additionally this will contain a file name \`build-request.json\` describing the library being processed. |
| /repo | mount (read/write) | The whole language repo. The mount is read/write to make diff-testing easier. Any changes made to this directory will have no-effect on the generated code, it is a deep-copy. |
| command | Positional Argument  | The value will always be \`build\` for this invocation. |

