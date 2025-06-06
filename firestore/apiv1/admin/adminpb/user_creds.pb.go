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
// source: google/firestore/admin/v1/user_creds.proto

package adminpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The state of the user creds (ENABLED or DISABLED).
type UserCreds_State int32

const (
	// The default value. Should not be used.
	UserCreds_STATE_UNSPECIFIED UserCreds_State = 0
	// The user creds are enabled.
	UserCreds_ENABLED UserCreds_State = 1
	// The user creds are disabled.
	UserCreds_DISABLED UserCreds_State = 2
)

// Enum value maps for UserCreds_State.
var (
	UserCreds_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
	}
	UserCreds_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ENABLED":           1,
		"DISABLED":          2,
	}
)

func (x UserCreds_State) Enum() *UserCreds_State {
	p := new(UserCreds_State)
	*p = x
	return p
}

func (x UserCreds_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserCreds_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_user_creds_proto_enumTypes[0].Descriptor()
}

func (UserCreds_State) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_user_creds_proto_enumTypes[0]
}

func (x UserCreds_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserCreds_State.Descriptor instead.
func (UserCreds_State) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_user_creds_proto_rawDescGZIP(), []int{0, 0}
}

// A Cloud Firestore User Creds.
type UserCreds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the UserCreds.
	// Format:
	// `projects/{project}/databases/{database}/userCreds/{user_creds}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The time the user creds were created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time the user creds were last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Whether the user creds are enabled or disabled. Defaults to
	// ENABLED on creation.
	State UserCreds_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.firestore.admin.v1.UserCreds_State" json:"state,omitempty"`
	// Output only. The plaintext server-generated password for the user creds.
	// Only populated in responses for CreateUserCreds and ResetUserPassword.
	SecurePassword string `protobuf:"bytes,5,opt,name=secure_password,json=securePassword,proto3" json:"secure_password,omitempty"`
	// Identity associated with this User Creds.
	//
	// Types that are assignable to UserCredsIdentity:
	//
	//	*UserCreds_ResourceIdentity_
	UserCredsIdentity isUserCreds_UserCredsIdentity `protobuf_oneof:"UserCredsIdentity"`
}

func (x *UserCreds) Reset() {
	*x = UserCreds{}
	mi := &file_google_firestore_admin_v1_user_creds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreds) ProtoMessage() {}

func (x *UserCreds) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_user_creds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreds.ProtoReflect.Descriptor instead.
func (*UserCreds) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_user_creds_proto_rawDescGZIP(), []int{0}
}

func (x *UserCreds) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserCreds) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *UserCreds) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *UserCreds) GetState() UserCreds_State {
	if x != nil {
		return x.State
	}
	return UserCreds_STATE_UNSPECIFIED
}

func (x *UserCreds) GetSecurePassword() string {
	if x != nil {
		return x.SecurePassword
	}
	return ""
}

func (m *UserCreds) GetUserCredsIdentity() isUserCreds_UserCredsIdentity {
	if m != nil {
		return m.UserCredsIdentity
	}
	return nil
}

func (x *UserCreds) GetResourceIdentity() *UserCreds_ResourceIdentity {
	if x, ok := x.GetUserCredsIdentity().(*UserCreds_ResourceIdentity_); ok {
		return x.ResourceIdentity
	}
	return nil
}

type isUserCreds_UserCredsIdentity interface {
	isUserCreds_UserCredsIdentity()
}

type UserCreds_ResourceIdentity_ struct {
	// Resource Identity descriptor.
	ResourceIdentity *UserCreds_ResourceIdentity `protobuf:"bytes,6,opt,name=resource_identity,json=resourceIdentity,proto3,oneof"`
}

func (*UserCreds_ResourceIdentity_) isUserCreds_UserCredsIdentity() {}

// Describes a Resource Identity principal.
type UserCreds_ResourceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Principal identifier string.
	// See: https://cloud.google.com/iam/docs/principal-identifiers
	Principal string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (x *UserCreds_ResourceIdentity) Reset() {
	*x = UserCreds_ResourceIdentity{}
	mi := &file_google_firestore_admin_v1_user_creds_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreds_ResourceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreds_ResourceIdentity) ProtoMessage() {}

func (x *UserCreds_ResourceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_user_creds_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreds_ResourceIdentity.ProtoReflect.Descriptor instead.
func (*UserCreds_ResourceIdentity) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_user_creds_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UserCreds_ResourceIdentity) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

var File_google_firestore_admin_v1_user_creds_proto protoreflect.FileDescriptor

var file_google_firestore_admin_v1_user_creds_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x05, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x64, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x35, 0x0a, 0x10, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21,
	0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x22, 0x39, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x3a, 0x7d, 0xea, 0x41,
	0x7a, 0x0a, 0x22, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x73, 0x12, 0x3e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x7d, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x73, 0x7d, 0x2a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73,
	0x32, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x42, 0xdd, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0xa2,
	0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_admin_v1_user_creds_proto_rawDescOnce sync.Once
	file_google_firestore_admin_v1_user_creds_proto_rawDescData = file_google_firestore_admin_v1_user_creds_proto_rawDesc
)

func file_google_firestore_admin_v1_user_creds_proto_rawDescGZIP() []byte {
	file_google_firestore_admin_v1_user_creds_proto_rawDescOnce.Do(func() {
		file_google_firestore_admin_v1_user_creds_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_admin_v1_user_creds_proto_rawDescData)
	})
	return file_google_firestore_admin_v1_user_creds_proto_rawDescData
}

var file_google_firestore_admin_v1_user_creds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_firestore_admin_v1_user_creds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_firestore_admin_v1_user_creds_proto_goTypes = []any{
	(UserCreds_State)(0),               // 0: google.firestore.admin.v1.UserCreds.State
	(*UserCreds)(nil),                  // 1: google.firestore.admin.v1.UserCreds
	(*UserCreds_ResourceIdentity)(nil), // 2: google.firestore.admin.v1.UserCreds.ResourceIdentity
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
}
var file_google_firestore_admin_v1_user_creds_proto_depIdxs = []int32{
	3, // 0: google.firestore.admin.v1.UserCreds.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.firestore.admin.v1.UserCreds.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.firestore.admin.v1.UserCreds.state:type_name -> google.firestore.admin.v1.UserCreds.State
	2, // 3: google.firestore.admin.v1.UserCreds.resource_identity:type_name -> google.firestore.admin.v1.UserCreds.ResourceIdentity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_firestore_admin_v1_user_creds_proto_init() }
func file_google_firestore_admin_v1_user_creds_proto_init() {
	if File_google_firestore_admin_v1_user_creds_proto != nil {
		return
	}
	file_google_firestore_admin_v1_user_creds_proto_msgTypes[0].OneofWrappers = []any{
		(*UserCreds_ResourceIdentity_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_admin_v1_user_creds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_admin_v1_user_creds_proto_goTypes,
		DependencyIndexes: file_google_firestore_admin_v1_user_creds_proto_depIdxs,
		EnumInfos:         file_google_firestore_admin_v1_user_creds_proto_enumTypes,
		MessageInfos:      file_google_firestore_admin_v1_user_creds_proto_msgTypes,
	}.Build()
	File_google_firestore_admin_v1_user_creds_proto = out.File
	file_google_firestore_admin_v1_user_creds_proto_rawDesc = nil
	file_google_firestore_admin_v1_user_creds_proto_goTypes = nil
	file_google_firestore_admin_v1_user_creds_proto_depIdxs = nil
}
