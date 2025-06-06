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
// source: google/cloud/aiplatform/v1beta1/example_store_service.proto

package aiplatformpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
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
	ExampleStoreService_CreateExampleStore_FullMethodName = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/CreateExampleStore"
	ExampleStoreService_GetExampleStore_FullMethodName    = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/GetExampleStore"
	ExampleStoreService_UpdateExampleStore_FullMethodName = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/UpdateExampleStore"
	ExampleStoreService_DeleteExampleStore_FullMethodName = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/DeleteExampleStore"
	ExampleStoreService_ListExampleStores_FullMethodName  = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/ListExampleStores"
	ExampleStoreService_UpsertExamples_FullMethodName     = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/UpsertExamples"
	ExampleStoreService_RemoveExamples_FullMethodName     = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/RemoveExamples"
	ExampleStoreService_SearchExamples_FullMethodName     = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/SearchExamples"
	ExampleStoreService_FetchExamples_FullMethodName      = "/google.cloud.aiplatform.v1beta1.ExampleStoreService/FetchExamples"
)

// ExampleStoreServiceClient is the client API for ExampleStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleStoreServiceClient interface {
	// Create an ExampleStore.
	CreateExampleStore(ctx context.Context, in *CreateExampleStoreRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Get an ExampleStore.
	GetExampleStore(ctx context.Context, in *GetExampleStoreRequest, opts ...grpc.CallOption) (*ExampleStore, error)
	// Update an ExampleStore.
	UpdateExampleStore(ctx context.Context, in *UpdateExampleStoreRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Delete an ExampleStore.
	DeleteExampleStore(ctx context.Context, in *DeleteExampleStoreRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// List ExampleStores in a Location.
	ListExampleStores(ctx context.Context, in *ListExampleStoresRequest, opts ...grpc.CallOption) (*ListExampleStoresResponse, error)
	// Create or update Examples in the Example Store.
	UpsertExamples(ctx context.Context, in *UpsertExamplesRequest, opts ...grpc.CallOption) (*UpsertExamplesResponse, error)
	// Remove Examples from the Example Store.
	RemoveExamples(ctx context.Context, in *RemoveExamplesRequest, opts ...grpc.CallOption) (*RemoveExamplesResponse, error)
	// Search for similar Examples for given selection criteria.
	SearchExamples(ctx context.Context, in *SearchExamplesRequest, opts ...grpc.CallOption) (*SearchExamplesResponse, error)
	// Get Examples from the Example Store.
	FetchExamples(ctx context.Context, in *FetchExamplesRequest, opts ...grpc.CallOption) (*FetchExamplesResponse, error)
}

type exampleStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleStoreServiceClient(cc grpc.ClientConnInterface) ExampleStoreServiceClient {
	return &exampleStoreServiceClient{cc}
}

func (c *exampleStoreServiceClient) CreateExampleStore(ctx context.Context, in *CreateExampleStoreRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExampleStoreService_CreateExampleStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) GetExampleStore(ctx context.Context, in *GetExampleStoreRequest, opts ...grpc.CallOption) (*ExampleStore, error) {
	out := new(ExampleStore)
	err := c.cc.Invoke(ctx, ExampleStoreService_GetExampleStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) UpdateExampleStore(ctx context.Context, in *UpdateExampleStoreRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExampleStoreService_UpdateExampleStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) DeleteExampleStore(ctx context.Context, in *DeleteExampleStoreRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExampleStoreService_DeleteExampleStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) ListExampleStores(ctx context.Context, in *ListExampleStoresRequest, opts ...grpc.CallOption) (*ListExampleStoresResponse, error) {
	out := new(ListExampleStoresResponse)
	err := c.cc.Invoke(ctx, ExampleStoreService_ListExampleStores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) UpsertExamples(ctx context.Context, in *UpsertExamplesRequest, opts ...grpc.CallOption) (*UpsertExamplesResponse, error) {
	out := new(UpsertExamplesResponse)
	err := c.cc.Invoke(ctx, ExampleStoreService_UpsertExamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) RemoveExamples(ctx context.Context, in *RemoveExamplesRequest, opts ...grpc.CallOption) (*RemoveExamplesResponse, error) {
	out := new(RemoveExamplesResponse)
	err := c.cc.Invoke(ctx, ExampleStoreService_RemoveExamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) SearchExamples(ctx context.Context, in *SearchExamplesRequest, opts ...grpc.CallOption) (*SearchExamplesResponse, error) {
	out := new(SearchExamplesResponse)
	err := c.cc.Invoke(ctx, ExampleStoreService_SearchExamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleStoreServiceClient) FetchExamples(ctx context.Context, in *FetchExamplesRequest, opts ...grpc.CallOption) (*FetchExamplesResponse, error) {
	out := new(FetchExamplesResponse)
	err := c.cc.Invoke(ctx, ExampleStoreService_FetchExamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleStoreServiceServer is the server API for ExampleStoreService service.
// All implementations should embed UnimplementedExampleStoreServiceServer
// for forward compatibility
type ExampleStoreServiceServer interface {
	// Create an ExampleStore.
	CreateExampleStore(context.Context, *CreateExampleStoreRequest) (*longrunningpb.Operation, error)
	// Get an ExampleStore.
	GetExampleStore(context.Context, *GetExampleStoreRequest) (*ExampleStore, error)
	// Update an ExampleStore.
	UpdateExampleStore(context.Context, *UpdateExampleStoreRequest) (*longrunningpb.Operation, error)
	// Delete an ExampleStore.
	DeleteExampleStore(context.Context, *DeleteExampleStoreRequest) (*longrunningpb.Operation, error)
	// List ExampleStores in a Location.
	ListExampleStores(context.Context, *ListExampleStoresRequest) (*ListExampleStoresResponse, error)
	// Create or update Examples in the Example Store.
	UpsertExamples(context.Context, *UpsertExamplesRequest) (*UpsertExamplesResponse, error)
	// Remove Examples from the Example Store.
	RemoveExamples(context.Context, *RemoveExamplesRequest) (*RemoveExamplesResponse, error)
	// Search for similar Examples for given selection criteria.
	SearchExamples(context.Context, *SearchExamplesRequest) (*SearchExamplesResponse, error)
	// Get Examples from the Example Store.
	FetchExamples(context.Context, *FetchExamplesRequest) (*FetchExamplesResponse, error)
}

// UnimplementedExampleStoreServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExampleStoreServiceServer struct {
}

func (UnimplementedExampleStoreServiceServer) CreateExampleStore(context.Context, *CreateExampleStoreRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExampleStore not implemented")
}
func (UnimplementedExampleStoreServiceServer) GetExampleStore(context.Context, *GetExampleStoreRequest) (*ExampleStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleStore not implemented")
}
func (UnimplementedExampleStoreServiceServer) UpdateExampleStore(context.Context, *UpdateExampleStoreRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExampleStore not implemented")
}
func (UnimplementedExampleStoreServiceServer) DeleteExampleStore(context.Context, *DeleteExampleStoreRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExampleStore not implemented")
}
func (UnimplementedExampleStoreServiceServer) ListExampleStores(context.Context, *ListExampleStoresRequest) (*ListExampleStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExampleStores not implemented")
}
func (UnimplementedExampleStoreServiceServer) UpsertExamples(context.Context, *UpsertExamplesRequest) (*UpsertExamplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertExamples not implemented")
}
func (UnimplementedExampleStoreServiceServer) RemoveExamples(context.Context, *RemoveExamplesRequest) (*RemoveExamplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveExamples not implemented")
}
func (UnimplementedExampleStoreServiceServer) SearchExamples(context.Context, *SearchExamplesRequest) (*SearchExamplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchExamples not implemented")
}
func (UnimplementedExampleStoreServiceServer) FetchExamples(context.Context, *FetchExamplesRequest) (*FetchExamplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchExamples not implemented")
}

// UnsafeExampleStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleStoreServiceServer will
// result in compilation errors.
type UnsafeExampleStoreServiceServer interface {
	mustEmbedUnimplementedExampleStoreServiceServer()
}

func RegisterExampleStoreServiceServer(s grpc.ServiceRegistrar, srv ExampleStoreServiceServer) {
	s.RegisterService(&ExampleStoreService_ServiceDesc, srv)
}

func _ExampleStoreService_CreateExampleStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExampleStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).CreateExampleStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_CreateExampleStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).CreateExampleStore(ctx, req.(*CreateExampleStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_GetExampleStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExampleStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).GetExampleStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_GetExampleStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).GetExampleStore(ctx, req.(*GetExampleStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_UpdateExampleStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExampleStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).UpdateExampleStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_UpdateExampleStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).UpdateExampleStore(ctx, req.(*UpdateExampleStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_DeleteExampleStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExampleStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).DeleteExampleStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_DeleteExampleStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).DeleteExampleStore(ctx, req.(*DeleteExampleStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_ListExampleStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExampleStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).ListExampleStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_ListExampleStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).ListExampleStores(ctx, req.(*ListExampleStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_UpsertExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertExamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).UpsertExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_UpsertExamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).UpsertExamples(ctx, req.(*UpsertExamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_RemoveExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveExamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).RemoveExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_RemoveExamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).RemoveExamples(ctx, req.(*RemoveExamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_SearchExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchExamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).SearchExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_SearchExamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).SearchExamples(ctx, req.(*SearchExamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleStoreService_FetchExamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchExamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleStoreServiceServer).FetchExamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleStoreService_FetchExamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleStoreServiceServer).FetchExamples(ctx, req.(*FetchExamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleStoreService_ServiceDesc is the grpc.ServiceDesc for ExampleStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1beta1.ExampleStoreService",
	HandlerType: (*ExampleStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExampleStore",
			Handler:    _ExampleStoreService_CreateExampleStore_Handler,
		},
		{
			MethodName: "GetExampleStore",
			Handler:    _ExampleStoreService_GetExampleStore_Handler,
		},
		{
			MethodName: "UpdateExampleStore",
			Handler:    _ExampleStoreService_UpdateExampleStore_Handler,
		},
		{
			MethodName: "DeleteExampleStore",
			Handler:    _ExampleStoreService_DeleteExampleStore_Handler,
		},
		{
			MethodName: "ListExampleStores",
			Handler:    _ExampleStoreService_ListExampleStores_Handler,
		},
		{
			MethodName: "UpsertExamples",
			Handler:    _ExampleStoreService_UpsertExamples_Handler,
		},
		{
			MethodName: "RemoveExamples",
			Handler:    _ExampleStoreService_RemoveExamples_Handler,
		},
		{
			MethodName: "SearchExamples",
			Handler:    _ExampleStoreService_SearchExamples_Handler,
		},
		{
			MethodName: "FetchExamples",
			Handler:    _ExampleStoreService_FetchExamples_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1beta1/example_store_service.proto",
}
