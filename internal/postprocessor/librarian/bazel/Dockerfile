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

# This Dockerfile creates a MOSS-compliant image for the Bazel-based
# Go generator, which is invoked by the Librarian tool.

# Per MOSS guidance, use a recommended base image.
FROM marketplace.gcr.io/google/debian12:latest

# Set environment variables for tool versions for easy updates.
ENV GO_VERSION=1.23.0
ENV BAZEL_VERSION=7.2.0

# Set up PATH for Go and other tools.
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:/usr/local/bin:$PATH

# Install essential system packages and Bazel dependencies.
# Clean up apt cache to keep the image smaller.
RUN apt-get update && \
    apt-get install -y \
    build-essential \
    ca-certificates \
    curl \
    git \
    gnupg \
    openssh-client \
    wget && \
    rm -rf /var/lib/apt/lists/*

# Install Go.
# TODO: Per MOSS guidance, this dependency should be pulled from an internal
# "airlock" repository. The URL will need to be updated.
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -O go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz && \
    go version

# Install Bazel.
# Using the official Bazel apt repository.
# TODO: Per MOSS guidance, this dependency should be pulled from an internal
# "airlock" repository if a package is available there.
RUN curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > /etc/apt/trusted.gpg.d/bazel.gpg && \
    echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list && \
    apt-get update && \
    apt-get install -y bazel=${BAZEL_VERSION} && \
    rm -rf /var/lib/apt/lists/* && \
    bazel --version

# Copy the generator source code into the image.
WORKDIR /app
COPY main.go go.mod ./

# Build the generator binary.
RUN go build -o /bazel-generator main.go

# Set the entrypoint for the container to run the compiled generator.
# The Librarian will provide commands like 'generate' as arguments.
ENTRYPOINT ["/bazel-generator"]