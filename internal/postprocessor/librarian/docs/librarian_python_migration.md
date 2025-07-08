

# Python Librarian Migration

| \#begin-approvals-addon-section See [go/g3a-approvals](http://goto.google.com/g3a-approvals) for instructions on how to add reviewers. Do not edit this section manually. |
| :---: |

**Author(s):** [Anthonios Partheniou](mailto:partheniou@google.com), [Omair Naveed](mailto:omairn@google.com) | **Last Updated**: Jun 10, 2025  | **Status**: Draft   
**Self link:** [go/sdk-librarian-python](http://goto.google.com/sdk-librarian-python) | **Project Issue**: [b/416183224](http://b/416183224)  *phase)*

# Objective

The purpose of this proposal is to migrate Python Cloud SDK generation from using the [owlbot](http://go/owlbot-howto) pipeline (Est. 2021\) to the [librarian](http://go/sdk-librarian-design%20) pipeline(Est. 2025). There are several reasons for this migration which are outlined in [go/sdk-librarian-design](http://goto.google.com/sdk-librarian-design). For Python specifically, the following problems exist in the current pipeline, and should be solved:

- Scaevola only scans vulnerabilities on docker images not running containers, without an additional build step. [Bazel](#bookmark=id.spebjxi8inu8) installs dependencies at runtime in the container, and therefore there is no vulnerability checking with the current configuration, which raises a security concern. In addition, dependencies fetched by bazel are coming from PyPI rather than airlock. [See Offline Bazel](#offline-bazel).  
- There are [**multiple sources of inputs**](#bookmark=id.9lfiwunytd0g) when generating a client library which introduce toil and complexity  
- [There are multiple code bases for client library generation](#bookmark=id.spebjxi8inu8). Code lives in both gapic-generator-python and synthtool, which introduces toil and complexity.

# Background

**Bazel**

- Bazel is used to generate client libraries. There is currently no vulnerability checking on the dependencies which are pulled into the build.  If this is not solved, we may need to request an exemption from MOSS. One solution is to move off of Bazel, so that we can install trusted dependencies at time of building the language specific docker image.

**Multiple sources of inputs**

- Currently multiple sources of inputs are needed when generating a client library. The Python SDK generation process depends on .repo-metadata.json, synthtool, settings defined in BUILD.bazel, and API specific service configuration. These multiple sources of inputs introduce toil and complexity in the pipeline. One solution is to consolidate the settings into a single location.

**Multiple code bases for client library generation**

- Currently templates which are used to create Python client libraries exist in 2 code bases: [gapic-generator-python](https://github.com/googleapis/gapic-generator-python/tree/main/gapic/templates) and synthtool ([monorepo](https://github.com/googleapis/synthtool/tree/master/synthtool/gcp/templates/python_mono_repo_library) / [split repo](https://github.com/googleapis/synthtool/tree/master/synthtool/gcp/templates/python_library)). This goes against the DRY (don’t repeat yourself) principle and introduces toil and complexity (See [recent example](https://docs.google.com/document/d/1PgZsJjEAzeQvqwGRyGW-GV8yRkvp1v57L9whme-IdiU/edit?tab=t.i6r0vnqjimib#bookmark=id.emvet9xp43mb)). One solution is to consolidate the synthtool templates into gapic-generator-python.

# Overview

Python Client Library generation uses [owlbot](https://docs.google.com/document/d/1izqHpAmJOzYbI5qLBin-OsN427d8Zm2lSxVpgBrq2BU/edit?resourcekey=0-9SNCEbs074Ixji1sP2SSyw&tab=t.0#heading=h.g7ign0f7iema), and there is a requirement to move to librarian for the reasons outlined in [go/sdk-librarian-design](http://goto.google.com/sdk-librarian-design). Python Client Library generation currently includes several issues that affect security, velocity and introduce complexity. This proposal aims to address these issues as part of the migration to the librarian pipeline.  This proposal will include 3 phases:

* [Owlbot in a box](#owlbot-in-a-box-\(phase1\)) \- Onboard to the librarian CLI as quickly as possible.  
* [Augmented Service Config migration](#augmented-service-config-migration-\(phase-2\))  \- move the source of truth for configurations to avoid unnecessary work of maintaining a separate configuration file  
* “Move to an offline build” \- Move to an offline build ,following security best practices.

The reason for this phased approach is to meet the deadline for Python to onboard to 1PP which is expected in August 2025 as per [https://docs.google.com/document/d/1WzlFSVsB\_W-mNsbwcFPFZf5UI4hNT54Koh7PmH8Avxg/edit](https://docs.google.com/document/d/1WzlFSVsB_W-mNsbwcFPFZf5UI4hNT54Koh7PmH8Avxg/edit) .

# Detailed Design

## Owlbot in a box (Phase1)  {#owlbot-in-a-box-(phase1)}

This “owlbot in a box” idea was suggested by [Cody Oss](mailto:codyoss@google.com) and will be quickest way to onboard to librarian without unnecessarily increasing the scope of the initial onboarding

### Goals

* Minimal/no diff between migrating from Owlbot to librarian  
* Separate tasks so that we reduce risk of the overall project

### Non-Goals 

* Address security concerns with bazel/existing Owlbot pipeline

### Language specific docker image

The language specific docker image will follow [MOSS guidance](http://go/moss-phase3-onboarding-cap1#steps) such as a MOSS recommended base image, and pull dependencies from centrally hosted repositories. The language specific docker image will run bazel. A prototype of this image lives here, however the prototype still requires these tweaks to be MOSS compliant. 

* Use OS recommended base images  
* Pull dependencies from airlock  
* Any other MOSS requirements?

### Debugging the language specific CLI (TODO: move this to the to appendix)

Run the \`clean\` command for a given library

```shell
python3 tools/python_generator_cli/cli.py clean --repo-root='/usr/local/google/home/partheniou/git/google-cloud-python' --library-id='google-cloud-container'
```

Run the \`generate-library\` command for a given library

```shell
python3 tools/python_generator_cli/cli.py generate-library --api-root='/usr/local/google/home/partheniou/git/googleapis' --generator-input='generator-input' --output='/usr/local/google/home/partheniou/git/google-cloud-python' --library-id='google-cloud-container'
```

### Debugging the language specific docker image

Build the docker image which exists in the root of google-cloud-python

```shell
docker build -f Dockerfile.generator -t google-cloud-python-generator:latest .
```

Run commands in bash terminal on the running container

```shell
docker run --rm -it --entrypoint /bin/bash google-cloud-python-generator:latest
```

Run the \`clean\` command for a given library

```shell
docker run --rm --user $(id -u):$(id -g)  --network none \-v $HOME/git/googleapis:/apis  \-v $HOME/git/google-cloud-python:/repo_root \ google-cloud-python-generator:latest \clean \--repo-root=repo_root \--library-id='google-cloud-container'
```

Run the \`generate-library\` command for a given library

```shell
`
```

### Ensuring that the diff is 0 at time of migration

We may want to reduce the size of the diff at the time of migration from owlbot to librarian. The reason is that the noisy diff can reduce the chance of detecting a problem with the migration for a given library. Python has some manual files/post processing and we need to use care to ensure that we don’t break anything. Currently there is a battle between GAPIC generator and release-please where gapic-generator sets the library version 0.0.0 and release-please sets the library version to the correct version, based on the [.release-please-manifest.json](https://github.com/googleapis/google-cloud-python/blob/main/.release-please-manifest.json) file. At the time of onboarding a new library, we should reset the library version to 0.0.0 either manually or via the librarian configure command. The files which need to be updated are all \`gapic\_version.py\`files and snippet \`\*.json\` files which live under \`samples/generated\_samples\`.

## Augmented Service Config migration (Phase 2\) {#augmented-service-config-migration-(phase-2)}

This phase aims to consolidate the various sources of input currently used for Python SDK generation into a single location, which is an augmented service config, as outlined in [go/sdk-librarian-unified-config](http://goto.google.com/sdk-librarian-unified-config) . This involves moving configurations from sources like \`.repo-metadata.json\`, \`BUILD.bazel\` and settings that have not yet been moved to the API specific service config. The primary goal is to reduce toil and complexity caused by maintaining settings across multiple files, though it will require some effort to consolidate the existing configurations.

**Why should we migrate to the augmented service config, instead of the API specific service config?**  
There is a lot of information that will need to be backfilled in hundreds of APIs. If we can consolidate all of the information that needs to be backfilled across languages in a single file, we can reduce the number of times that we disturb API teams with these updates. 

## Move to an offline build (Phase 3\)

[https://github.com/googleapis/librarian/issues/277](https://github.com/googleapis/librarian/issues/277) (TODO: Close this bug)

Bazel currently installs dependencies at build time. There is no vulnerability checking on these dependencies, which is a concern.

The solution for phase 3 is to migrate off of Bazel. We can simply run protoc directly within the docker container. Rather than have bazel install the dependencies at build time, dependencies will be installed at the time of building the docker image. As a side benefit, rather than a hard requirement, this design does not require network access during build time.

Pro

* Dependencies will be installed from airlock, and the docker image will be onboarded to Scaevola for vulnerability monitoring.  
* Aligns with .Net Prototype, which also moved off of Bazel as part of the migration to librarian   
* No shared workspace across language

Con

* This is a change from how we are generating Python client libraries today  
* The WORKSPACE file in the googleapis repository becomes no longer the source of truth. For customers who rely on it, the dependencies there may become outdated. We would need to provide a migration path for them.

 

## Remove dependency on synthtool templates and synthtool post-processing logic

This logic will be moved to gapic-generator-python. See [GAPIC options outside of service yaml](#bookmark=id.dwa3nyfg6st6) which includes the necessary options. Also see prototype [https://github.com/googleapis/gapic-generator-python/tree/1pp-prototype-v1.25.0](https://github.com/googleapis/gapic-generator-python/tree/1pp-prototype-v1.25.0) 

## Configuring a new client library

[https://github.com/googleapis/librarian/issues/466](https://github.com/googleapis/librarian/issues/466) 

The librarian ‘configure’ command will help us onboard a new library. 

Certain packages don’t have a 1:1 mapping between the package name and the api path, as a result, we will need to maintain a file of “package name overrides” so that when the configure command is called, we will use the correct package name.

The list of libraries in \`pipeline-state.yaml\`will be empty at time of initial onboarding from owlbot to librarian. A single librarian \`configure\` command will run, which will call the language specific \`configure\` command through the language specific CLI. 

\<[Omair Naveed](mailto:omairn@google.com)to fill the rest in . The idea is that the diff between this [prototype \`pipeline-state.json\` config](https://github.com/googleapis/google-cloud-python/blob/1pp-add-pipeline-state/generator-input/pipeline-state.json) file and the  newly populated \`pipeline-state.yaml\` will be zero, ignoring the difference in file formats. Note: there are some changes coming as part of [go/sdk-librarian-state-config](http://goto.google.com/sdk-librarian-state-config) \>

## Split repository considerations

\<TBD\>

# 

# Alternatives considered

## Continue to use Bazel

[https://github.com/googleapis/librarian/issues/277](https://github.com/googleapis/librarian/issues/277) 

Pros:

* \`BUILD.bazel\` serves as the source of truth for some Python settings, so continuing with Bazel would eliminate the need to migrate this information.

Cons:

* Complex dependency resolution.  Multiple versions of dependencies can exist within WORKSPACE  
* Network access. In order to fetch bazel dependencies at time of generation, we need bazel 7\. The \`bazel fetch –all\` command is not supported in Bazel 6\. When you run the command with bazel 6.5.0, you get \`ERROR: \--all :: Unrecognized option: \--all\`. The upgrade to bazel 7 is challenging because simply updating the version is not enough. There are other changes required which would delay our timeline for onboarding to 1PP.  
* Shared workspace where updates to dependencies in one language can affect generation in another language.  
* Maintenance cost. As of writing, we are 2 major versions behind. We use bazel 6, however bazel 8 is currently the latest released version. In addition, we need to drop WORKSPACE and move to bzlmod if we want to move to bazel 9 as per [https://bazel.build/external/migration](https://bazel.build/external/migration) . These updates can be challenging, and we should evaluate the cost/reward for this migration. We may not want to continue to take on this maintenance cost.  
* If we want scaevola vulnerability scanning on the artifacts that bazel is fetching when building client libraries, we would need to trigger scaevola on a running container, which creates an extra step in the build process.

## Add apis.json to generator-input, consolidate multiple inputs

[https://github.com/googleapis/librarian/issues/276](https://github.com/googleapis/librarian/issues/276) 

Everything that the generator needs will be added to an augmented service config [go/sdk-librarian-unified-config](http://goto.google.com/sdk-librarian-unified-config).  This will include consolidated inputs from \`.repo-metadata.json\`, \`BUILD.bazel\`, as well as other settings that have yet to be moved to the service config. The reason for this consolidation is that having multiple sources of inputs would introduce complexity. Ideally this information should live in the API specific service config, but it’s not clear if we can do this in a reasonable timeline.

Pros

* A single location for all of the config settings

Cons

* Requires some work to consolidate the settings, which would subsequently be thrown away in favor of the augmented service config or service config in google3/google

See prototype [https://github.com/googleapis/google-cloud-python/tree/1pp-add-apis-json](https://github.com/googleapis/google-cloud-python/tree/1pp-add-apis-json)

Notes based on prototype

- The apis.json is an IOU of technical debt of configurations that should be moved to the API specific service config

Here are the steps to fulfill the IOU if we decide to do so:

### GAPIC Options outside of service YAML

| Missing from service yaml | What is it? | How to resolve? | Owner (best guess) |
| ----- | ----- | ----- | ----- |
| api-description | A summary of the API. [Example](https://github.com/googleapis/googleapis/blob/111b7383752255d1849a8d3b7259ed735acc4f97/google/cloud/accessapproval/v1/accessapproval_v1.yaml#L9-L10) | [See 1](#bookmark=kix.c6432r6mbu1x)  | Service Team |
| autogen-snippets | Whether to disable auto-generated samples. | [See 2](#bookmark=kix.kn9esrjs8pym) | Service Team |
| Default-proto-package(Part of prototype, still TBD) | Multiple API versions are bundled together in the same artifact and we use this information to configure a default version. | [See 3](#bookmark=kix.3v1n0muqgdn5) | Service Team |
| documentation-name | The name used by Cloud RAD for where Python reference docs will be located | [See 3](#bookmark=kix.3v1n0muqgdn5) | Cloud SDK Team |
| documentation-uri | Product documentation uri | [See 4](#bookmark=kix.ppvttkjgyrht) | Service Team |
| Gapic-version (Part of prototype, still TBD) | Version of the client library | [See 5](#bookmark=kix.7m95ms3tjbp6) | Cloud SDK Team |
| python-gapic-namespace | Namespace of the package | [See 3](#bookmark=kix.3v1n0muqgdn5) | Cloud SDK Team |
| python-gapic-name | Module name | [See 3](#bookmark=kix.3v1n0muqgdn5) | Cloud SDK Team |
| Reference-doc-includes(Part of prototype, still TBD) | Multiple API versions are bundled together in the same artifact and we use this information to stitch all the versions together in the docs | [See 3](#bookmark=kix.3v1n0muqgdn5) | Cloud SDK Team |
| release-level | SDK release level | [See 3](#bookmark=kix.3v1n0muqgdn5) | Service Team |
| rest-numeric-enums | Whether the API supports numeric enums | [See 2](#bookmark=kix.kn9esrjs8pym) | Service Team |
| retry-config | Grpc service config | [See 2](#bookmark=kix.kn9esrjs8pym) | Service Team |
| service-yaml | Name of the service yaml file | [See 2](#bookmark=kix.kn9esrjs8pym) | Service Team |
| transport | Transport supported by the API | [See 2](#bookmark=kix.kn9esrjs8pym) | Service Team |
| warehouse-package-name | PyPI package name | [See 3](#bookmark=kix.3v1n0muqgdn5) | Cloud SDK Team |
| title | Title of the API | [See 6](#bookmark=kix.ytyp8utt38c) | Service Team |

1 \- Documentation summary should be backfilled  
2 \- This can be part of CommonLanguageSettings (A cross language API specific config)  
3 \- Create a new PythonSetting (A language specific API config)  
4 \- Publishing section should be backfilled  
5 \- Nothing. This will be used to update the version of client libraries  
6 \- The API is missing a service yaml file. Eg. [https://github.com/googleapis/googleapis/tree/master/google/cloud/bigquery/logging/v1](https://github.com/googleapis/googleapis/tree/master/google/cloud/bigquery/logging/v1) 

# 

## Create initial pipeline-state.json

https://github.com/googleapis/librarian/issues/274 (TODO: Close this bug)

As part of the migration to librarian, a pipeline-state.json file needs to exist in the generator-input directory of the repository.

A script will be used once to create an initial pipeline-state.json file, rather than calling configure. After which, it is expected that subsequent language specific updates of pipeline-state.json will happen using the configure language specific CLI command. New libraries which are on-boarded will use the same lastGeneratedCommit value used at the time of onboarding all packages.

[https://github.com/googleapis/google-cloud-python/tree/1pp-add-pipeline-state](https://github.com/googleapis/google-cloud-python/tree/1pp-add-pipeline-state) 

# Work Estimates

# Documentation plan

# Launch plans

# Risks

# Appendix

## Offline Bazel {#offline-bazel}

Bazel build can run in \`offline\` mode but it requires bazel 7\. The \`bazel fetch –all\` command is not supported in Bazel 6\. When you run the command with bazel 6.5.0, you get \`ERROR: \--all :: Unrecognized option: \--all\`.

## Steps to run librarian for Python using the Prototype
