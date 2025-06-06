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
// source: google/cloud/apihub/v1/provisioning_service.proto

package apihubpb

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
	Provisioning_CreateApiHubInstance_FullMethodName = "/google.cloud.apihub.v1.Provisioning/CreateApiHubInstance"
	Provisioning_GetApiHubInstance_FullMethodName    = "/google.cloud.apihub.v1.Provisioning/GetApiHubInstance"
	Provisioning_LookupApiHubInstance_FullMethodName = "/google.cloud.apihub.v1.Provisioning/LookupApiHubInstance"
)

// ProvisioningClient is the client API for Provisioning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisioningClient interface {
	// Provisions instance resources for the API Hub.
	CreateApiHubInstance(ctx context.Context, in *CreateApiHubInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Gets details of a single API Hub instance.
	GetApiHubInstance(ctx context.Context, in *GetApiHubInstanceRequest, opts ...grpc.CallOption) (*ApiHubInstance, error)
	// Looks up an Api Hub instance in a given GCP project. There will always be
	// only one Api Hub instance for a GCP project across all locations.
	LookupApiHubInstance(ctx context.Context, in *LookupApiHubInstanceRequest, opts ...grpc.CallOption) (*LookupApiHubInstanceResponse, error)
}

type provisioningClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisioningClient(cc grpc.ClientConnInterface) ProvisioningClient {
	return &provisioningClient{cc}
}

func (c *provisioningClient) CreateApiHubInstance(ctx context.Context, in *CreateApiHubInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, Provisioning_CreateApiHubInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisioningClient) GetApiHubInstance(ctx context.Context, in *GetApiHubInstanceRequest, opts ...grpc.CallOption) (*ApiHubInstance, error) {
	out := new(ApiHubInstance)
	err := c.cc.Invoke(ctx, Provisioning_GetApiHubInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisioningClient) LookupApiHubInstance(ctx context.Context, in *LookupApiHubInstanceRequest, opts ...grpc.CallOption) (*LookupApiHubInstanceResponse, error) {
	out := new(LookupApiHubInstanceResponse)
	err := c.cc.Invoke(ctx, Provisioning_LookupApiHubInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisioningServer is the server API for Provisioning service.
// All implementations should embed UnimplementedProvisioningServer
// for forward compatibility
type ProvisioningServer interface {
	// Provisions instance resources for the API Hub.
	CreateApiHubInstance(context.Context, *CreateApiHubInstanceRequest) (*longrunningpb.Operation, error)
	// Gets details of a single API Hub instance.
	GetApiHubInstance(context.Context, *GetApiHubInstanceRequest) (*ApiHubInstance, error)
	// Looks up an Api Hub instance in a given GCP project. There will always be
	// only one Api Hub instance for a GCP project across all locations.
	LookupApiHubInstance(context.Context, *LookupApiHubInstanceRequest) (*LookupApiHubInstanceResponse, error)
}

// UnimplementedProvisioningServer should be embedded to have forward compatible implementations.
type UnimplementedProvisioningServer struct {
}

func (UnimplementedProvisioningServer) CreateApiHubInstance(context.Context, *CreateApiHubInstanceRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiHubInstance not implemented")
}
func (UnimplementedProvisioningServer) GetApiHubInstance(context.Context, *GetApiHubInstanceRequest) (*ApiHubInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiHubInstance not implemented")
}
func (UnimplementedProvisioningServer) LookupApiHubInstance(context.Context, *LookupApiHubInstanceRequest) (*LookupApiHubInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupApiHubInstance not implemented")
}

// UnsafeProvisioningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisioningServer will
// result in compilation errors.
type UnsafeProvisioningServer interface {
	mustEmbedUnimplementedProvisioningServer()
}

func RegisterProvisioningServer(s grpc.ServiceRegistrar, srv ProvisioningServer) {
	s.RegisterService(&Provisioning_ServiceDesc, srv)
}

func _Provisioning_CreateApiHubInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiHubInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServer).CreateApiHubInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioning_CreateApiHubInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServer).CreateApiHubInstance(ctx, req.(*CreateApiHubInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioning_GetApiHubInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiHubInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServer).GetApiHubInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioning_GetApiHubInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServer).GetApiHubInstance(ctx, req.(*GetApiHubInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioning_LookupApiHubInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupApiHubInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServer).LookupApiHubInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioning_LookupApiHubInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServer).LookupApiHubInstance(ctx, req.(*LookupApiHubInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provisioning_ServiceDesc is the grpc.ServiceDesc for Provisioning service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provisioning_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.apihub.v1.Provisioning",
	HandlerType: (*ProvisioningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApiHubInstance",
			Handler:    _Provisioning_CreateApiHubInstance_Handler,
		},
		{
			MethodName: "GetApiHubInstance",
			Handler:    _Provisioning_GetApiHubInstance_Handler,
		},
		{
			MethodName: "LookupApiHubInstance",
			Handler:    _Provisioning_LookupApiHubInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/apihub/v1/provisioning_service.proto",
}
