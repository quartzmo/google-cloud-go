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
// source: google/shopping/merchant/reports/v1beta/reports.proto

package reportspb

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
	ReportService_Search_FullMethodName = "/google.shopping.merchant.reports.v1beta.ReportService/Search"
)

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	// Retrieves a report defined by a search query. The response might contain
	// fewer rows than specified by `page_size`. Rely on `next_page_token` to
	// determine if there are more rows to be requested.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, ReportService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations should embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	// Retrieves a report defined by a search query. The response might contain
	// fewer rows than specified by `page_size`. Rely on `next_page_token` to
	// determine if there are more rows to be requested.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedReportServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.reports.v1beta.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _ReportService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/reports/v1beta/reports.proto",
}
