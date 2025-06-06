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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/ai/generativelanguage/v1beta2/model_service.proto

package generativelanguagepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for getting information about a specific Model.
type GetModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the model.
	//
	// This name should match a model name returned by the `ListModels` method.
	//
	// Format: `models/{model}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetModelRequest) Reset() {
	*x = GetModelRequest{}
	mi := &file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelRequest) ProtoMessage() {}

func (x *GetModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelRequest.ProtoReflect.Descriptor instead.
func (*GetModelRequest) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetModelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for listing all Models.
type ListModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of `Models` to return (per page).
	//
	// The service may return fewer models.
	// If unspecified, at most 50 models will be returned per page.
	// This method returns at most 1000 models per page, even if you pass a larger
	// page_size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListModels` call.
	//
	// Provide the `page_token` returned by one request as an argument to the next
	// request to retrieve the next page.
	//
	// When paginating, all other parameters provided to `ListModels` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListModelsRequest) Reset() {
	*x = ListModelsRequest{}
	mi := &file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsRequest) ProtoMessage() {}

func (x *ListModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsRequest.ProtoReflect.Descriptor instead.
func (*ListModelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListModelsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListModelsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response from `ListModel` containing a paginated list of Models.
type ListModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The returned Models.
	Models []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	//
	// If this field is omitted, there are no more pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListModelsResponse) Reset() {
	*x = ListModelsResponse{}
	mi := &file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsResponse) ProtoMessage() {}

func (x *ListModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsResponse.ProtoReflect.Descriptor instead.
func (*ListModelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListModelsResponse) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

func (x *ListModelsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_ai_generativelanguage_v1beta2_model_service_proto protoreflect.FileDescriptor

var file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x81,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x32, 0x80, 0x03, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x27, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xaf, 0x01,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x37, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2e, 0xda, 0x41, 0x14, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x2c, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a,
	0x24, 0xca, 0x41, 0x21, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x9f, 0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x42, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x70, 0x62, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescOnce sync.Once
	file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescData = file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDesc
)

func file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescGZIP() []byte {
	file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescOnce.Do(func() {
		file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescData)
	})
	return file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDescData
}

var file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ai_generativelanguage_v1beta2_model_service_proto_goTypes = []any{
	(*GetModelRequest)(nil),    // 0: google.ai.generativelanguage.v1beta2.GetModelRequest
	(*ListModelsRequest)(nil),  // 1: google.ai.generativelanguage.v1beta2.ListModelsRequest
	(*ListModelsResponse)(nil), // 2: google.ai.generativelanguage.v1beta2.ListModelsResponse
	(*Model)(nil),              // 3: google.ai.generativelanguage.v1beta2.Model
}
var file_google_ai_generativelanguage_v1beta2_model_service_proto_depIdxs = []int32{
	3, // 0: google.ai.generativelanguage.v1beta2.ListModelsResponse.models:type_name -> google.ai.generativelanguage.v1beta2.Model
	0, // 1: google.ai.generativelanguage.v1beta2.ModelService.GetModel:input_type -> google.ai.generativelanguage.v1beta2.GetModelRequest
	1, // 2: google.ai.generativelanguage.v1beta2.ModelService.ListModels:input_type -> google.ai.generativelanguage.v1beta2.ListModelsRequest
	3, // 3: google.ai.generativelanguage.v1beta2.ModelService.GetModel:output_type -> google.ai.generativelanguage.v1beta2.Model
	2, // 4: google.ai.generativelanguage.v1beta2.ModelService.ListModels:output_type -> google.ai.generativelanguage.v1beta2.ListModelsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ai_generativelanguage_v1beta2_model_service_proto_init() }
func file_google_ai_generativelanguage_v1beta2_model_service_proto_init() {
	if File_google_ai_generativelanguage_v1beta2_model_service_proto != nil {
		return
	}
	file_google_ai_generativelanguage_v1beta2_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ai_generativelanguage_v1beta2_model_service_proto_goTypes,
		DependencyIndexes: file_google_ai_generativelanguage_v1beta2_model_service_proto_depIdxs,
		MessageInfos:      file_google_ai_generativelanguage_v1beta2_model_service_proto_msgTypes,
	}.Build()
	File_google_ai_generativelanguage_v1beta2_model_service_proto = out.File
	file_google_ai_generativelanguage_v1beta2_model_service_proto_rawDesc = nil
	file_google_ai_generativelanguage_v1beta2_model_service_proto_goTypes = nil
	file_google_ai_generativelanguage_v1beta2_model_service_proto_depIdxs = nil
}
