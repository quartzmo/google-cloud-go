# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This Dockerfile creates a MOSS-compliant image for the Go generator,
# which is invoked by the Librarian tool.

# Per MOSS guidance, use a recommended base image.
FROM marketplace.gcr.io/google/debian12:latest

# Set environment variables for tool versions for easy updates.
ENV GO_VERSION=1.23.0
ENV PROTOC_VERSION=25.7
ENV GO_PROTOC_PLUGIN_VERSION=1.36.6
ENV GAPIC_GENERATOR_VERSION=0.53.1

# Set up PATH for Go and other tools.
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:/usr/local/bin:$PATH

# Install essential system packages.
# Clean up apt cache to keep the image smaller.
RUN apt-get update && \
    apt-get install -y \
    build-essential \
    ca-certificates \
    curl \
    unzip \
    wget && \
    rm -rf /var/lib/apt/lists/*

# Install Go.
# TODO: Per MOSS guidance, this dependency should be pulled from an internal
# "airlock" repository. The URL will need to be updated.
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -O go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz && \
    go version

# Install protoc (Protobuf Compiler).
# TODO: Per MOSS guidance, this dependency should be pulled from an internal
# "airlock" repository. The URL will need to be updated.
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -O protoc.zip && \
    unzip protoc.zip -d /usr/local && \
    rm protoc.zip && \
    protoc --version

# Install required Go tools for protoc.
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GO_PROTOC_PLUGIN_VERSION} && \
    go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v${GAPIC_GENERATOR_VERSION}

# Copy the generator source code into the image.
WORKDIR /app
COPY main.go .

# Build the generator binary.
RUN go build -o /generator main.go

# Set the entrypoint for the container to run the compiled generator.
# The Librarian will provide commands like 'generate' as arguments.
ENTRYPOINT ["/generator"]
