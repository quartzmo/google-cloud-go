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
// source: google/cloud/discoveryengine/v1beta/serving_config_service.proto

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
	ServingConfigService_UpdateServingConfig_FullMethodName = "/google.cloud.discoveryengine.v1beta.ServingConfigService/UpdateServingConfig"
	ServingConfigService_GetServingConfig_FullMethodName    = "/google.cloud.discoveryengine.v1beta.ServingConfigService/GetServingConfig"
	ServingConfigService_ListServingConfigs_FullMethodName  = "/google.cloud.discoveryengine.v1beta.ServingConfigService/ListServingConfigs"
)

// ServingConfigServiceClient is the client API for ServingConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServingConfigServiceClient interface {
	// Updates a ServingConfig.
	//
	// Returns a NOT_FOUND error if the ServingConfig does not exist.
	UpdateServingConfig(ctx context.Context, in *UpdateServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error)
	// Gets a ServingConfig.
	//
	// Returns a NotFound error if the ServingConfig does not exist.
	GetServingConfig(ctx context.Context, in *GetServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error)
	// Lists all ServingConfigs linked to this dataStore.
	ListServingConfigs(ctx context.Context, in *ListServingConfigsRequest, opts ...grpc.CallOption) (*ListServingConfigsResponse, error)
}

type servingConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServingConfigServiceClient(cc grpc.ClientConnInterface) ServingConfigServiceClient {
	return &servingConfigServiceClient{cc}
}

func (c *servingConfigServiceClient) UpdateServingConfig(ctx context.Context, in *UpdateServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error) {
	out := new(ServingConfig)
	err := c.cc.Invoke(ctx, ServingConfigService_UpdateServingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servingConfigServiceClient) GetServingConfig(ctx context.Context, in *GetServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error) {
	out := new(ServingConfig)
	err := c.cc.Invoke(ctx, ServingConfigService_GetServingConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servingConfigServiceClient) ListServingConfigs(ctx context.Context, in *ListServingConfigsRequest, opts ...grpc.CallOption) (*ListServingConfigsResponse, error) {
	out := new(ListServingConfigsResponse)
	err := c.cc.Invoke(ctx, ServingConfigService_ListServingConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServingConfigServiceServer is the server API for ServingConfigService service.
// All implementations should embed UnimplementedServingConfigServiceServer
// for forward compatibility
type ServingConfigServiceServer interface {
	// Updates a ServingConfig.
	//
	// Returns a NOT_FOUND error if the ServingConfig does not exist.
	UpdateServingConfig(context.Context, *UpdateServingConfigRequest) (*ServingConfig, error)
	// Gets a ServingConfig.
	//
	// Returns a NotFound error if the ServingConfig does not exist.
	GetServingConfig(context.Context, *GetServingConfigRequest) (*ServingConfig, error)
	// Lists all ServingConfigs linked to this dataStore.
	ListServingConfigs(context.Context, *ListServingConfigsRequest) (*ListServingConfigsResponse, error)
}

// UnimplementedServingConfigServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServingConfigServiceServer struct {
}

func (UnimplementedServingConfigServiceServer) UpdateServingConfig(context.Context, *UpdateServingConfigRequest) (*ServingConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServingConfig not implemented")
}
func (UnimplementedServingConfigServiceServer) GetServingConfig(context.Context, *GetServingConfigRequest) (*ServingConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServingConfig not implemented")
}
func (UnimplementedServingConfigServiceServer) ListServingConfigs(context.Context, *ListServingConfigsRequest) (*ListServingConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServingConfigs not implemented")
}

// UnsafeServingConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServingConfigServiceServer will
// result in compilation errors.
type UnsafeServingConfigServiceServer interface {
	mustEmbedUnimplementedServingConfigServiceServer()
}

func RegisterServingConfigServiceServer(s grpc.ServiceRegistrar, srv ServingConfigServiceServer) {
	s.RegisterService(&ServingConfigService_ServiceDesc, srv)
}

func _ServingConfigService_UpdateServingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingConfigServiceServer).UpdateServingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServingConfigService_UpdateServingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingConfigServiceServer).UpdateServingConfig(ctx, req.(*UpdateServingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServingConfigService_GetServingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingConfigServiceServer).GetServingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServingConfigService_GetServingConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingConfigServiceServer).GetServingConfig(ctx, req.(*GetServingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServingConfigService_ListServingConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServingConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingConfigServiceServer).ListServingConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServingConfigService_ListServingConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingConfigServiceServer).ListServingConfigs(ctx, req.(*ListServingConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServingConfigService_ServiceDesc is the grpc.ServiceDesc for ServingConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServingConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1beta.ServingConfigService",
	HandlerType: (*ServingConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateServingConfig",
			Handler:    _ServingConfigService_UpdateServingConfig_Handler,
		},
		{
			MethodName: "GetServingConfig",
			Handler:    _ServingConfigService_GetServingConfig_Handler,
		},
		{
			MethodName: "ListServingConfigs",
			Handler:    _ServingConfigService_ListServingConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1beta/serving_config_service.proto",
}
