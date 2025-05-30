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
// source: google/cloud/aiplatform/v1/feature_selector.proto

package aiplatformpb

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

// Matcher for Features of an EntityType by Feature ID.
type IdMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The following are accepted as `ids`:
	//
	//   - A single-element list containing only `*`, which selects all Features
	//     in the target EntityType, or
	//   - A list containing only Feature IDs, which selects only Features with
	//     those IDs in the target EntityType.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdMatcher) Reset() {
	*x = IdMatcher{}
	mi := &file_google_cloud_aiplatform_v1_feature_selector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdMatcher) ProtoMessage() {}

func (x *IdMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_selector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdMatcher.ProtoReflect.Descriptor instead.
func (*IdMatcher) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescGZIP(), []int{0}
}

func (x *IdMatcher) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Selector for Features of an EntityType.
type FeatureSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Matches Features based on ID.
	IdMatcher *IdMatcher `protobuf:"bytes,1,opt,name=id_matcher,json=idMatcher,proto3" json:"id_matcher,omitempty"`
}

func (x *FeatureSelector) Reset() {
	*x = FeatureSelector{}
	mi := &file_google_cloud_aiplatform_v1_feature_selector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureSelector) ProtoMessage() {}

func (x *FeatureSelector) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_selector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureSelector.ProtoReflect.Descriptor instead.
func (*FeatureSelector) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureSelector) GetIdMatcher() *IdMatcher {
	if x != nil {
		return x.IdMatcher
	}
	return nil
}

var File_google_cloud_aiplatform_v1_feature_selector_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_feature_selector_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x22, 0x0a, 0x09, 0x49, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x15, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x0f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x69, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0xd2, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62,
	0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescData = file_google_cloud_aiplatform_v1_feature_selector_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_feature_selector_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_feature_selector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1_feature_selector_proto_goTypes = []any{
	(*IdMatcher)(nil),       // 0: google.cloud.aiplatform.v1.IdMatcher
	(*FeatureSelector)(nil), // 1: google.cloud.aiplatform.v1.FeatureSelector
}
var file_google_cloud_aiplatform_v1_feature_selector_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1.FeatureSelector.id_matcher:type_name -> google.cloud.aiplatform.v1.IdMatcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_feature_selector_proto_init() }
func file_google_cloud_aiplatform_v1_feature_selector_proto_init() {
	if File_google_cloud_aiplatform_v1_feature_selector_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_feature_selector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_feature_selector_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_feature_selector_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_feature_selector_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_feature_selector_proto = out.File
	file_google_cloud_aiplatform_v1_feature_selector_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_feature_selector_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_feature_selector_proto_depIdxs = nil
}
