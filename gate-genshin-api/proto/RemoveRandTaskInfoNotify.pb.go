// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: RemoveRandTaskInfoNotify.proto

package proto

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

type RemoveRandTaskInfoNotify_FinishReason int32

const (
	RemoveRandTaskInfoNotify_FINISH_REASON_DEFAULT  RemoveRandTaskInfoNotify_FinishReason = 0
	RemoveRandTaskInfoNotify_FINISH_REASON_CLEAR    RemoveRandTaskInfoNotify_FinishReason = 1
	RemoveRandTaskInfoNotify_FINISH_REASON_DISTANCE RemoveRandTaskInfoNotify_FinishReason = 2
	RemoveRandTaskInfoNotify_FINISH_REASON_FINISH   RemoveRandTaskInfoNotify_FinishReason = 3
)

// Enum value maps for RemoveRandTaskInfoNotify_FinishReason.
var (
	RemoveRandTaskInfoNotify_FinishReason_name = map[int32]string{
		0: "FINISH_REASON_DEFAULT",
		1: "FINISH_REASON_CLEAR",
		2: "FINISH_REASON_DISTANCE",
		3: "FINISH_REASON_FINISH",
	}
	RemoveRandTaskInfoNotify_FinishReason_value = map[string]int32{
		"FINISH_REASON_DEFAULT":  0,
		"FINISH_REASON_CLEAR":    1,
		"FINISH_REASON_DISTANCE": 2,
		"FINISH_REASON_FINISH":   3,
	}
)

func (x RemoveRandTaskInfoNotify_FinishReason) Enum() *RemoveRandTaskInfoNotify_FinishReason {
	p := new(RemoveRandTaskInfoNotify_FinishReason)
	*p = x
	return p
}

func (x RemoveRandTaskInfoNotify_FinishReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoveRandTaskInfoNotify_FinishReason) Descriptor() protoreflect.EnumDescriptor {
	return file_RemoveRandTaskInfoNotify_proto_enumTypes[0].Descriptor()
}

func (RemoveRandTaskInfoNotify_FinishReason) Type() protoreflect.EnumType {
	return &file_RemoveRandTaskInfoNotify_proto_enumTypes[0]
}

func (x RemoveRandTaskInfoNotify_FinishReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemoveRandTaskInfoNotify_FinishReason.Descriptor instead.
func (RemoveRandTaskInfoNotify_FinishReason) EnumDescriptor() ([]byte, []int) {
	return file_RemoveRandTaskInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 161
// EnetChannelId: 0
// EnetIsReliable: true
type RemoveRandTaskInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSucc     bool                                  `protobuf:"varint,9,opt,name=is_succ,json=isSucc,proto3" json:"is_succ,omitempty"`
	Reason     RemoveRandTaskInfoNotify_FinishReason `protobuf:"varint,10,opt,name=reason,proto3,enum=RemoveRandTaskInfoNotify_FinishReason" json:"reason,omitempty"`
	RandTaskId uint32                                `protobuf:"varint,13,opt,name=rand_task_id,json=randTaskId,proto3" json:"rand_task_id,omitempty"`
}

func (x *RemoveRandTaskInfoNotify) Reset() {
	*x = RemoveRandTaskInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RemoveRandTaskInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRandTaskInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRandTaskInfoNotify) ProtoMessage() {}

func (x *RemoveRandTaskInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RemoveRandTaskInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRandTaskInfoNotify.ProtoReflect.Descriptor instead.
func (*RemoveRandTaskInfoNotify) Descriptor() ([]byte, []int) {
	return file_RemoveRandTaskInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RemoveRandTaskInfoNotify) GetIsSucc() bool {
	if x != nil {
		return x.IsSucc
	}
	return false
}

func (x *RemoveRandTaskInfoNotify) GetReason() RemoveRandTaskInfoNotify_FinishReason {
	if x != nil {
		return x.Reason
	}
	return RemoveRandTaskInfoNotify_FINISH_REASON_DEFAULT
}

func (x *RemoveRandTaskInfoNotify) GetRandTaskId() uint32 {
	if x != nil {
		return x.RandTaskId
	}
	return 0
}

var File_RemoveRandTaskInfoNotify_proto protoreflect.FileDescriptor

var file_RemoveRandTaskInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x02, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x61, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x61,
	0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x49,
	0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RemoveRandTaskInfoNotify_proto_rawDescOnce sync.Once
	file_RemoveRandTaskInfoNotify_proto_rawDescData = file_RemoveRandTaskInfoNotify_proto_rawDesc
)

func file_RemoveRandTaskInfoNotify_proto_rawDescGZIP() []byte {
	file_RemoveRandTaskInfoNotify_proto_rawDescOnce.Do(func() {
		file_RemoveRandTaskInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RemoveRandTaskInfoNotify_proto_rawDescData)
	})
	return file_RemoveRandTaskInfoNotify_proto_rawDescData
}

var file_RemoveRandTaskInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RemoveRandTaskInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RemoveRandTaskInfoNotify_proto_goTypes = []interface{}{
	(RemoveRandTaskInfoNotify_FinishReason)(0), // 0: RemoveRandTaskInfoNotify.FinishReason
	(*RemoveRandTaskInfoNotify)(nil),           // 1: RemoveRandTaskInfoNotify
}
var file_RemoveRandTaskInfoNotify_proto_depIdxs = []int32{
	0, // 0: RemoveRandTaskInfoNotify.reason:type_name -> RemoveRandTaskInfoNotify.FinishReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RemoveRandTaskInfoNotify_proto_init() }
func file_RemoveRandTaskInfoNotify_proto_init() {
	if File_RemoveRandTaskInfoNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RemoveRandTaskInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRandTaskInfoNotify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RemoveRandTaskInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RemoveRandTaskInfoNotify_proto_goTypes,
		DependencyIndexes: file_RemoveRandTaskInfoNotify_proto_depIdxs,
		EnumInfos:         file_RemoveRandTaskInfoNotify_proto_enumTypes,
		MessageInfos:      file_RemoveRandTaskInfoNotify_proto_msgTypes,
	}.Build()
	File_RemoveRandTaskInfoNotify_proto = out.File
	file_RemoveRandTaskInfoNotify_proto_rawDesc = nil
	file_RemoveRandTaskInfoNotify_proto_goTypes = nil
	file_RemoveRandTaskInfoNotify_proto_depIdxs = nil
}
