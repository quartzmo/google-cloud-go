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
// source: google/ai/generativelanguage/v1alpha/cached_content.proto

package generativelanguagepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Content that has been preprocessed and can be used in subsequent request
// to GenerativeService.
//
// Cached content can be only used with model it was created for.
type CachedContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies when this resource will expire.
	//
	// Types that are assignable to Expiration:
	//
	//	*CachedContent_ExpireTime
	//	*CachedContent_Ttl
	Expiration isCachedContent_Expiration `protobuf_oneof:"expiration"`
	// Optional. Identifier. The resource name referring to the cached content.
	// Format: `cachedContents/{id}`
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Optional. Immutable. The user-generated meaningful display name of the
	// cached content. Maximum 128 Unicode characters.
	DisplayName *string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	// Required. Immutable. The name of the `Model` to use for cached content
	// Format: `models/{model}`
	Model *string `protobuf:"bytes,2,opt,name=model,proto3,oneof" json:"model,omitempty"`
	// Optional. Input only. Immutable. Developer set system instruction.
	// Currently text only.
	SystemInstruction *Content `protobuf:"bytes,3,opt,name=system_instruction,json=systemInstruction,proto3,oneof" json:"system_instruction,omitempty"`
	// Optional. Input only. Immutable. The content to cache.
	Contents []*Content `protobuf:"bytes,4,rep,name=contents,proto3" json:"contents,omitempty"`
	// Optional. Input only. Immutable. A list of `Tools` the model may use to
	// generate the next response
	Tools []*Tool `protobuf:"bytes,5,rep,name=tools,proto3" json:"tools,omitempty"`
	// Optional. Input only. Immutable. Tool config. This config is shared for all
	// tools.
	ToolConfig *ToolConfig `protobuf:"bytes,6,opt,name=tool_config,json=toolConfig,proto3,oneof" json:"tool_config,omitempty"`
	// Output only. Creation time of the cache entry.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. When the cache entry was last updated in UTC time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Metadata on the usage of the cached content.
	UsageMetadata *CachedContent_UsageMetadata `protobuf:"bytes,12,opt,name=usage_metadata,json=usageMetadata,proto3" json:"usage_metadata,omitempty"`
}

func (x *CachedContent) Reset() {
	*x = CachedContent{}
	mi := &file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CachedContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachedContent) ProtoMessage() {}

func (x *CachedContent) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachedContent.ProtoReflect.Descriptor instead.
func (*CachedContent) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescGZIP(), []int{0}
}

func (m *CachedContent) GetExpiration() isCachedContent_Expiration {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (x *CachedContent) GetExpireTime() *timestamppb.Timestamp {
	if x, ok := x.GetExpiration().(*CachedContent_ExpireTime); ok {
		return x.ExpireTime
	}
	return nil
}

func (x *CachedContent) GetTtl() *durationpb.Duration {
	if x, ok := x.GetExpiration().(*CachedContent_Ttl); ok {
		return x.Ttl
	}
	return nil
}

func (x *CachedContent) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CachedContent) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *CachedContent) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *CachedContent) GetSystemInstruction() *Content {
	if x != nil {
		return x.SystemInstruction
	}
	return nil
}

func (x *CachedContent) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *CachedContent) GetTools() []*Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *CachedContent) GetToolConfig() *ToolConfig {
	if x != nil {
		return x.ToolConfig
	}
	return nil
}

func (x *CachedContent) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CachedContent) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *CachedContent) GetUsageMetadata() *CachedContent_UsageMetadata {
	if x != nil {
		return x.UsageMetadata
	}
	return nil
}

type isCachedContent_Expiration interface {
	isCachedContent_Expiration()
}

type CachedContent_ExpireTime struct {
	// Timestamp in UTC of when this resource is considered expired.
	// This is *always* provided on output, regardless of what was sent
	// on input.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=expire_time,json=expireTime,proto3,oneof"`
}

type CachedContent_Ttl struct {
	// Input only. New TTL for this resource, input only.
	Ttl *durationpb.Duration `protobuf:"bytes,10,opt,name=ttl,proto3,oneof"`
}

func (*CachedContent_ExpireTime) isCachedContent_Expiration() {}

func (*CachedContent_Ttl) isCachedContent_Expiration() {}

// Metadata on the usage of the cached content.
type CachedContent_UsageMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of tokens that the cached content consumes.
	TotalTokenCount int32 `protobuf:"varint,1,opt,name=total_token_count,json=totalTokenCount,proto3" json:"total_token_count,omitempty"`
}

func (x *CachedContent_UsageMetadata) Reset() {
	*x = CachedContent_UsageMetadata{}
	mi := &file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CachedContent_UsageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachedContent_UsageMetadata) ProtoMessage() {}

func (x *CachedContent_UsageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachedContent_UsageMetadata.ProtoReflect.Descriptor instead.
func (*CachedContent_UsageMetadata) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CachedContent_UsageMetadata) GetTotalTokenCount() int32 {
	if x != nil {
		return x.TotalTokenCount
	}
	return 0
}

var File_google_ai_generativelanguage_v1alpha_cached_content_proto protoreflect.FileDescriptor

var file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x89, 0x09, 0x0a, 0x0d, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x04,
	0x48, 0x00, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x08, 0xe0, 0x41, 0x01, 0x48, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x29, 0x0a, 0x27, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x03, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x6c, 0x0a, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x09, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x04, 0x48, 0x04, 0x52,
	0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0xe0, 0x41,
	0x04, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x05, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x42, 0x09, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0xe0, 0x41,
	0x04, 0x52, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x61, 0x0a, 0x0b, 0x74, 0x6f, 0x6f, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x09, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x04, 0x48, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x6d, 0x0a, 0x0e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x68, 0xea, 0x41, 0x65,
	0x0a, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x13, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2a, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x0d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xa0,
	0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x12, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x3b, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescOnce sync.Once
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescData = file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDesc
)

func file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescGZIP() []byte {
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescOnce.Do(func() {
		file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescData)
	})
	return file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDescData
}

var file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ai_generativelanguage_v1alpha_cached_content_proto_goTypes = []any{
	(*CachedContent)(nil),               // 0: google.ai.generativelanguage.v1alpha.CachedContent
	(*CachedContent_UsageMetadata)(nil), // 1: google.ai.generativelanguage.v1alpha.CachedContent.UsageMetadata
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),         // 3: google.protobuf.Duration
	(*Content)(nil),                     // 4: google.ai.generativelanguage.v1alpha.Content
	(*Tool)(nil),                        // 5: google.ai.generativelanguage.v1alpha.Tool
	(*ToolConfig)(nil),                  // 6: google.ai.generativelanguage.v1alpha.ToolConfig
}
var file_google_ai_generativelanguage_v1alpha_cached_content_proto_depIdxs = []int32{
	2, // 0: google.ai.generativelanguage.v1alpha.CachedContent.expire_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.ai.generativelanguage.v1alpha.CachedContent.ttl:type_name -> google.protobuf.Duration
	4, // 2: google.ai.generativelanguage.v1alpha.CachedContent.system_instruction:type_name -> google.ai.generativelanguage.v1alpha.Content
	4, // 3: google.ai.generativelanguage.v1alpha.CachedContent.contents:type_name -> google.ai.generativelanguage.v1alpha.Content
	5, // 4: google.ai.generativelanguage.v1alpha.CachedContent.tools:type_name -> google.ai.generativelanguage.v1alpha.Tool
	6, // 5: google.ai.generativelanguage.v1alpha.CachedContent.tool_config:type_name -> google.ai.generativelanguage.v1alpha.ToolConfig
	2, // 6: google.ai.generativelanguage.v1alpha.CachedContent.create_time:type_name -> google.protobuf.Timestamp
	2, // 7: google.ai.generativelanguage.v1alpha.CachedContent.update_time:type_name -> google.protobuf.Timestamp
	1, // 8: google.ai.generativelanguage.v1alpha.CachedContent.usage_metadata:type_name -> google.ai.generativelanguage.v1alpha.CachedContent.UsageMetadata
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ai_generativelanguage_v1alpha_cached_content_proto_init() }
func file_google_ai_generativelanguage_v1alpha_cached_content_proto_init() {
	if File_google_ai_generativelanguage_v1alpha_cached_content_proto != nil {
		return
	}
	file_google_ai_generativelanguage_v1alpha_content_proto_init()
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes[0].OneofWrappers = []any{
		(*CachedContent_ExpireTime)(nil),
		(*CachedContent_Ttl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ai_generativelanguage_v1alpha_cached_content_proto_goTypes,
		DependencyIndexes: file_google_ai_generativelanguage_v1alpha_cached_content_proto_depIdxs,
		MessageInfos:      file_google_ai_generativelanguage_v1alpha_cached_content_proto_msgTypes,
	}.Build()
	File_google_ai_generativelanguage_v1alpha_cached_content_proto = out.File
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_rawDesc = nil
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_goTypes = nil
	file_google_ai_generativelanguage_v1alpha_cached_content_proto_depIdxs = nil
}
