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
// source: EvtEntityStartDieEndNotify.proto

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

// CmdId: 381
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type EvtEntityStartDieEndNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Immediately  bool        `protobuf:"varint,15,opt,name=immediately,proto3" json:"immediately,omitempty"`
	DieStateFlag uint32      `protobuf:"varint,12,opt,name=die_state_flag,json=dieStateFlag,proto3" json:"die_state_flag,omitempty"`
	EntityId     uint32      `protobuf:"varint,8,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ForwardType  ForwardType `protobuf:"varint,11,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
}

func (x *EvtEntityStartDieEndNotify) Reset() {
	*x = EvtEntityStartDieEndNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtEntityStartDieEndNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtEntityStartDieEndNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtEntityStartDieEndNotify) ProtoMessage() {}

func (x *EvtEntityStartDieEndNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtEntityStartDieEndNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtEntityStartDieEndNotify.ProtoReflect.Descriptor instead.
func (*EvtEntityStartDieEndNotify) Descriptor() ([]byte, []int) {
	return file_EvtEntityStartDieEndNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtEntityStartDieEndNotify) GetImmediately() bool {
	if x != nil {
		return x.Immediately
	}
	return false
}

func (x *EvtEntityStartDieEndNotify) GetDieStateFlag() uint32 {
	if x != nil {
		return x.DieStateFlag
	}
	return 0
}

func (x *EvtEntityStartDieEndNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtEntityStartDieEndNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

var File_EvtEntityStartDieEndNotify_proto protoreflect.FileDescriptor

var file_EvtEntityStartDieEndNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x45, 0x76, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x69, 0x65, 0x45, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x1a, 0x45, 0x76, 0x74, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x69, 0x65, 0x45, 0x6e, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6d, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x69, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x64, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtEntityStartDieEndNotify_proto_rawDescOnce sync.Once
	file_EvtEntityStartDieEndNotify_proto_rawDescData = file_EvtEntityStartDieEndNotify_proto_rawDesc
)

func file_EvtEntityStartDieEndNotify_proto_rawDescGZIP() []byte {
	file_EvtEntityStartDieEndNotify_proto_rawDescOnce.Do(func() {
		file_EvtEntityStartDieEndNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtEntityStartDieEndNotify_proto_rawDescData)
	})
	return file_EvtEntityStartDieEndNotify_proto_rawDescData
}

var file_EvtEntityStartDieEndNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtEntityStartDieEndNotify_proto_goTypes = []interface{}{
	(*EvtEntityStartDieEndNotify)(nil), // 0: EvtEntityStartDieEndNotify
	(ForwardType)(0),                   // 1: ForwardType
}
var file_EvtEntityStartDieEndNotify_proto_depIdxs = []int32{
	1, // 0: EvtEntityStartDieEndNotify.forward_type:type_name -> ForwardType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EvtEntityStartDieEndNotify_proto_init() }
func file_EvtEntityStartDieEndNotify_proto_init() {
	if File_EvtEntityStartDieEndNotify_proto != nil {
		return
	}
	file_ForwardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtEntityStartDieEndNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtEntityStartDieEndNotify); i {
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
			RawDescriptor: file_EvtEntityStartDieEndNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtEntityStartDieEndNotify_proto_goTypes,
		DependencyIndexes: file_EvtEntityStartDieEndNotify_proto_depIdxs,
		MessageInfos:      file_EvtEntityStartDieEndNotify_proto_msgTypes,
	}.Build()
	File_EvtEntityStartDieEndNotify_proto = out.File
	file_EvtEntityStartDieEndNotify_proto_rawDesc = nil
	file_EvtEntityStartDieEndNotify_proto_goTypes = nil
	file_EvtEntityStartDieEndNotify_proto_depIdxs = nil
}
