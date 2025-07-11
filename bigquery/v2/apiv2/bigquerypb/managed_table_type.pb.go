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
// source: google/cloud/bigquery/v2/managed_table_type.proto

package bigquerypb

import (
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

// The classification of managed table types that can be created.
type ManagedTableType int32

const (
	// No managed table type specified.
	ManagedTableType_MANAGED_TABLE_TYPE_UNSPECIFIED ManagedTableType = 0
	// The managed table is a native BigQuery table.
	ManagedTableType_NATIVE ManagedTableType = 1
	// The managed table is a BigLake table for Apache Iceberg in BigQuery.
	ManagedTableType_BIGLAKE ManagedTableType = 2
)

// Enum value maps for ManagedTableType.
var (
	ManagedTableType_name = map[int32]string{
		0: "MANAGED_TABLE_TYPE_UNSPECIFIED",
		1: "NATIVE",
		2: "BIGLAKE",
	}
	ManagedTableType_value = map[string]int32{
		"MANAGED_TABLE_TYPE_UNSPECIFIED": 0,
		"NATIVE":                         1,
		"BIGLAKE":                        2,
	}
)

func (x ManagedTableType) Enum() *ManagedTableType {
	p := new(ManagedTableType)
	*p = x
	return p
}

func (x ManagedTableType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManagedTableType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_v2_managed_table_type_proto_enumTypes[0].Descriptor()
}

func (ManagedTableType) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_v2_managed_table_type_proto_enumTypes[0]
}

func (x ManagedTableType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManagedTableType.Descriptor instead.
func (ManagedTableType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_bigquery_v2_managed_table_type_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_v2_managed_table_type_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x2a, 0x4f, 0x0a,
	0x10, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x44, 0x5f, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x47, 0x4c, 0x41, 0x4b, 0x45, 0x10, 0x02, 0x42, 0x74,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x42, 0x15,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x62, 0x3b, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescData = file_google_cloud_bigquery_v2_managed_table_type_proto_rawDesc
)

func file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescData)
	})
	return file_google_cloud_bigquery_v2_managed_table_type_proto_rawDescData
}

var file_google_cloud_bigquery_v2_managed_table_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_bigquery_v2_managed_table_type_proto_goTypes = []any{
	(ManagedTableType)(0), // 0: google.cloud.bigquery.v2.ManagedTableType
}
var file_google_cloud_bigquery_v2_managed_table_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_v2_managed_table_type_proto_init() }
func file_google_cloud_bigquery_v2_managed_table_type_proto_init() {
	if File_google_cloud_bigquery_v2_managed_table_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_bigquery_v2_managed_table_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_v2_managed_table_type_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_v2_managed_table_type_proto_depIdxs,
		EnumInfos:         file_google_cloud_bigquery_v2_managed_table_type_proto_enumTypes,
	}.Build()
	File_google_cloud_bigquery_v2_managed_table_type_proto = out.File
	file_google_cloud_bigquery_v2_managed_table_type_proto_rawDesc = nil
	file_google_cloud_bigquery_v2_managed_table_type_proto_goTypes = nil
	file_google_cloud_bigquery_v2_managed_table_type_proto_depIdxs = nil
}
