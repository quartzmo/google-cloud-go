// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/cloud/discoveryengine/v1alpha/chunk_service.proto

package discoveryenginepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChunkService_GetChunk_FullMethodName   = "/google.cloud.discoveryengine.v1alpha.ChunkService/GetChunk"
	ChunkService_ListChunks_FullMethodName = "/google.cloud.discoveryengine.v1alpha.ChunkService/ListChunks"
)

// ChunkServiceClient is the client API for ChunkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkServiceClient interface {
	// Gets a [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*Chunk, error)
	// Gets a list of [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s.
	ListChunks(ctx context.Context, in *ListChunksRequest, opts ...grpc.CallOption) (*ListChunksResponse, error)
}

type chunkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServiceClient(cc grpc.ClientConnInterface) ChunkServiceClient {
	return &chunkServiceClient{cc}
}

func (c *chunkServiceClient) GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, ChunkService_GetChunk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServiceClient) ListChunks(ctx context.Context, in *ListChunksRequest, opts ...grpc.CallOption) (*ListChunksResponse, error) {
	out := new(ListChunksResponse)
	err := c.cc.Invoke(ctx, ChunkService_ListChunks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServiceServer is the server API for ChunkService service.
// All implementations should embed UnimplementedChunkServiceServer
// for forward compatibility
type ChunkServiceServer interface {
	// Gets a [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetChunk(context.Context, *GetChunkRequest) (*Chunk, error)
	// Gets a list of [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s.
	ListChunks(context.Context, *ListChunksRequest) (*ListChunksResponse, error)
}

// UnimplementedChunkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChunkServiceServer struct {
}

func (UnimplementedChunkServiceServer) GetChunk(context.Context, *GetChunkRequest) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}
func (UnimplementedChunkServiceServer) ListChunks(context.Context, *ListChunksRequest) (*ListChunksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChunks not implemented")
}

// UnsafeChunkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServiceServer will
// result in compilation errors.
type UnsafeChunkServiceServer interface {
	mustEmbedUnimplementedChunkServiceServer()
}

func RegisterChunkServiceServer(s grpc.ServiceRegistrar, srv ChunkServiceServer) {
	s.RegisterService(&ChunkService_ServiceDesc, srv)
}

func _ChunkService_GetChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).GetChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChunkService_GetChunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).GetChunk(ctx, req.(*GetChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkService_ListChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).ListChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChunkService_ListChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).ListChunks(ctx, req.(*ListChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkService_ServiceDesc is the grpc.ServiceDesc for ChunkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.ChunkService",
	HandlerType: (*ChunkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChunk",
			Handler:    _ChunkService_GetChunk_Handler,
		},
		{
			MethodName: "ListChunks",
			Handler:    _ChunkService_ListChunks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/chunk_service.proto",
}
