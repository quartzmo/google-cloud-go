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
// source: google/cloud/chronicle/v1/reference_list.proto

package chroniclepb

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
	ReferenceListService_GetReferenceList_FullMethodName    = "/google.cloud.chronicle.v1.ReferenceListService/GetReferenceList"
	ReferenceListService_ListReferenceLists_FullMethodName  = "/google.cloud.chronicle.v1.ReferenceListService/ListReferenceLists"
	ReferenceListService_CreateReferenceList_FullMethodName = "/google.cloud.chronicle.v1.ReferenceListService/CreateReferenceList"
	ReferenceListService_UpdateReferenceList_FullMethodName = "/google.cloud.chronicle.v1.ReferenceListService/UpdateReferenceList"
)

// ReferenceListServiceClient is the client API for ReferenceListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReferenceListServiceClient interface {
	// Gets a single reference list.
	GetReferenceList(ctx context.Context, in *GetReferenceListRequest, opts ...grpc.CallOption) (*ReferenceList, error)
	// Lists a collection of reference lists.
	ListReferenceLists(ctx context.Context, in *ListReferenceListsRequest, opts ...grpc.CallOption) (*ListReferenceListsResponse, error)
	// Creates a new reference list.
	CreateReferenceList(ctx context.Context, in *CreateReferenceListRequest, opts ...grpc.CallOption) (*ReferenceList, error)
	// Updates an existing reference list.
	UpdateReferenceList(ctx context.Context, in *UpdateReferenceListRequest, opts ...grpc.CallOption) (*ReferenceList, error)
}

type referenceListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReferenceListServiceClient(cc grpc.ClientConnInterface) ReferenceListServiceClient {
	return &referenceListServiceClient{cc}
}

func (c *referenceListServiceClient) GetReferenceList(ctx context.Context, in *GetReferenceListRequest, opts ...grpc.CallOption) (*ReferenceList, error) {
	out := new(ReferenceList)
	err := c.cc.Invoke(ctx, ReferenceListService_GetReferenceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceListServiceClient) ListReferenceLists(ctx context.Context, in *ListReferenceListsRequest, opts ...grpc.CallOption) (*ListReferenceListsResponse, error) {
	out := new(ListReferenceListsResponse)
	err := c.cc.Invoke(ctx, ReferenceListService_ListReferenceLists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceListServiceClient) CreateReferenceList(ctx context.Context, in *CreateReferenceListRequest, opts ...grpc.CallOption) (*ReferenceList, error) {
	out := new(ReferenceList)
	err := c.cc.Invoke(ctx, ReferenceListService_CreateReferenceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceListServiceClient) UpdateReferenceList(ctx context.Context, in *UpdateReferenceListRequest, opts ...grpc.CallOption) (*ReferenceList, error) {
	out := new(ReferenceList)
	err := c.cc.Invoke(ctx, ReferenceListService_UpdateReferenceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReferenceListServiceServer is the server API for ReferenceListService service.
// All implementations should embed UnimplementedReferenceListServiceServer
// for forward compatibility
type ReferenceListServiceServer interface {
	// Gets a single reference list.
	GetReferenceList(context.Context, *GetReferenceListRequest) (*ReferenceList, error)
	// Lists a collection of reference lists.
	ListReferenceLists(context.Context, *ListReferenceListsRequest) (*ListReferenceListsResponse, error)
	// Creates a new reference list.
	CreateReferenceList(context.Context, *CreateReferenceListRequest) (*ReferenceList, error)
	// Updates an existing reference list.
	UpdateReferenceList(context.Context, *UpdateReferenceListRequest) (*ReferenceList, error)
}

// UnimplementedReferenceListServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReferenceListServiceServer struct {
}

func (UnimplementedReferenceListServiceServer) GetReferenceList(context.Context, *GetReferenceListRequest) (*ReferenceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReferenceList not implemented")
}
func (UnimplementedReferenceListServiceServer) ListReferenceLists(context.Context, *ListReferenceListsRequest) (*ListReferenceListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReferenceLists not implemented")
}
func (UnimplementedReferenceListServiceServer) CreateReferenceList(context.Context, *CreateReferenceListRequest) (*ReferenceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReferenceList not implemented")
}
func (UnimplementedReferenceListServiceServer) UpdateReferenceList(context.Context, *UpdateReferenceListRequest) (*ReferenceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReferenceList not implemented")
}

// UnsafeReferenceListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReferenceListServiceServer will
// result in compilation errors.
type UnsafeReferenceListServiceServer interface {
	mustEmbedUnimplementedReferenceListServiceServer()
}

func RegisterReferenceListServiceServer(s grpc.ServiceRegistrar, srv ReferenceListServiceServer) {
	s.RegisterService(&ReferenceListService_ServiceDesc, srv)
}

func _ReferenceListService_GetReferenceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReferenceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceListServiceServer).GetReferenceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReferenceListService_GetReferenceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceListServiceServer).GetReferenceList(ctx, req.(*GetReferenceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceListService_ListReferenceLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReferenceListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceListServiceServer).ListReferenceLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReferenceListService_ListReferenceLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceListServiceServer).ListReferenceLists(ctx, req.(*ListReferenceListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceListService_CreateReferenceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReferenceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceListServiceServer).CreateReferenceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReferenceListService_CreateReferenceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceListServiceServer).CreateReferenceList(ctx, req.(*CreateReferenceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceListService_UpdateReferenceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReferenceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceListServiceServer).UpdateReferenceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReferenceListService_UpdateReferenceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceListServiceServer).UpdateReferenceList(ctx, req.(*UpdateReferenceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReferenceListService_ServiceDesc is the grpc.ServiceDesc for ReferenceListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReferenceListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.chronicle.v1.ReferenceListService",
	HandlerType: (*ReferenceListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReferenceList",
			Handler:    _ReferenceListService_GetReferenceList_Handler,
		},
		{
			MethodName: "ListReferenceLists",
			Handler:    _ReferenceListService_ListReferenceLists_Handler,
		},
		{
			MethodName: "CreateReferenceList",
			Handler:    _ReferenceListService_CreateReferenceList_Handler,
		},
		{
			MethodName: "UpdateReferenceList",
			Handler:    _ReferenceListService_UpdateReferenceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/chronicle/v1/reference_list.proto",
}
