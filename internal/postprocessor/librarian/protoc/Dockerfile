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

# This Dockerfile creates a MOSS-compliant image for a Go application
# that depends on the protoc compiler.

# Per MOSS guidance, use a recommended base image.
FROM marketplace.gcr.io/google/debian12:latest

# Set environment variables for tool versions for easy updates.
ENV GO_VERSION 1.23.0
ENV PROTOC_VERSION 25.7

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

# Install common Go tools.
# These are installed into $GOPATH/bin which is on the PATH.
# `golint` is deprecated and not included. `staticcheck` is a modern replacement.
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GO_PROTOC_PLUGIN_VERSION} && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v${GAPIC_GENERATOR_VERSION}

# Set the working directory and copy the application source.
# This assumes the build context is the 'protoc' directory.
WORKDIR /app
COPY . .

# Download Go module dependencies and build the application.
RUN go mod download
RUN CGO_ENABLED=0 go build -o librarian

# Set the entrypoint for the container to run the compiled application.
ENTRYPOINT ["/app/librarian"]
