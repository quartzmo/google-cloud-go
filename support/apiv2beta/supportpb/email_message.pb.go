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
// source: google/cloud/support/v2beta/email_message.proto

package supportpb

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

// An email associated with a support case.
type EmailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Resource name for the email message.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Time when this email message object was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The user or Google Support agent that created this email
	// message. This is inferred from the headers on the email message.
	Actor *Actor `protobuf:"bytes,3,opt,name=actor,proto3" json:"actor,omitempty"`
	// Output only. Subject of the email.
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	// Output only. Email addresses the email was sent to.
	RecipientEmailAddresses []string `protobuf:"bytes,5,rep,name=recipient_email_addresses,json=recipientEmailAddresses,proto3" json:"recipient_email_addresses,omitempty"`
	// Output only. Email addresses CCed on the email.
	CcEmailAddresses []string `protobuf:"bytes,6,rep,name=cc_email_addresses,json=ccEmailAddresses,proto3" json:"cc_email_addresses,omitempty"`
	// Output only. The full email message body. A best-effort attempt is made to
	// remove extraneous reply threads.
	BodyContent *TextContent `protobuf:"bytes,8,opt,name=body_content,json=bodyContent,proto3" json:"body_content,omitempty"`
}

func (x *EmailMessage) Reset() {
	*x = EmailMessage{}
	mi := &file_google_cloud_support_v2beta_email_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailMessage) ProtoMessage() {}

func (x *EmailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2beta_email_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailMessage.ProtoReflect.Descriptor instead.
func (*EmailMessage) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2beta_email_message_proto_rawDescGZIP(), []int{0}
}

func (x *EmailMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmailMessage) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *EmailMessage) GetActor() *Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *EmailMessage) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailMessage) GetRecipientEmailAddresses() []string {
	if x != nil {
		return x.RecipientEmailAddresses
	}
	return nil
}

func (x *EmailMessage) GetCcEmailAddresses() []string {
	if x != nil {
		return x.CcEmailAddresses
	}
	return nil
}

func (x *EmailMessage) GetBodyContent() *TextContent {
	if x != nil {
		return x.BodyContent
	}
	return nil
}

var File_google_cloud_support_v2beta_email_message_proto protoreflect.FileDescriptor

var file_google_cloud_support_v2beta_email_message_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe3, 0x04, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3f, 0x0a, 0x19, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x17, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x12, 0x63, 0x63,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x63, 0x63, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x50, 0x0a,
	0x0c, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0b, 0x62, 0x6f, 0x64, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a,
	0xd3, 0x01, 0xea, 0x41, 0xcf, 0x01, 0x0a, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x73, 0x65,
	0x7d, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x7b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x7d, 0x12,
	0x47, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x73, 0x65, 0x7d, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x7d, 0x2a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xce, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x42, 0x11, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x3b,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x56, 0x32, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56,
	0x32, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a,
	0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_support_v2beta_email_message_proto_rawDescOnce sync.Once
	file_google_cloud_support_v2beta_email_message_proto_rawDescData = file_google_cloud_support_v2beta_email_message_proto_rawDesc
)

func file_google_cloud_support_v2beta_email_message_proto_rawDescGZIP() []byte {
	file_google_cloud_support_v2beta_email_message_proto_rawDescOnce.Do(func() {
		file_google_cloud_support_v2beta_email_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_support_v2beta_email_message_proto_rawDescData)
	})
	return file_google_cloud_support_v2beta_email_message_proto_rawDescData
}

var file_google_cloud_support_v2beta_email_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_support_v2beta_email_message_proto_goTypes = []any{
	(*EmailMessage)(nil),          // 0: google.cloud.support.v2beta.EmailMessage
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Actor)(nil),                 // 2: google.cloud.support.v2beta.Actor
	(*TextContent)(nil),           // 3: google.cloud.support.v2beta.TextContent
}
var file_google_cloud_support_v2beta_email_message_proto_depIdxs = []int32{
	1, // 0: google.cloud.support.v2beta.EmailMessage.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.support.v2beta.EmailMessage.actor:type_name -> google.cloud.support.v2beta.Actor
	3, // 2: google.cloud.support.v2beta.EmailMessage.body_content:type_name -> google.cloud.support.v2beta.TextContent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_support_v2beta_email_message_proto_init() }
func file_google_cloud_support_v2beta_email_message_proto_init() {
	if File_google_cloud_support_v2beta_email_message_proto != nil {
		return
	}
	file_google_cloud_support_v2beta_actor_proto_init()
	file_google_cloud_support_v2beta_content_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_support_v2beta_email_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_support_v2beta_email_message_proto_goTypes,
		DependencyIndexes: file_google_cloud_support_v2beta_email_message_proto_depIdxs,
		MessageInfos:      file_google_cloud_support_v2beta_email_message_proto_msgTypes,
	}.Build()
	File_google_cloud_support_v2beta_email_message_proto = out.File
	file_google_cloud_support_v2beta_email_message_proto_rawDesc = nil
	file_google_cloud_support_v2beta_email_message_proto_goTypes = nil
	file_google_cloud_support_v2beta_email_message_proto_depIdxs = nil
}
