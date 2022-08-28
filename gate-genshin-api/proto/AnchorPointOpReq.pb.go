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
// source: AnchorPointOpReq.proto

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

type AnchorPointOpReq_AnchorPointOpType int32

const (
	AnchorPointOpReq_ANCHOR_POINT_OP_TYPE_NONE     AnchorPointOpReq_AnchorPointOpType = 0
	AnchorPointOpReq_ANCHOR_POINT_OP_TYPE_TELEPORT AnchorPointOpReq_AnchorPointOpType = 1
	AnchorPointOpReq_ANCHOR_POINT_OP_TYPE_REMOVE   AnchorPointOpReq_AnchorPointOpType = 2
)

// Enum value maps for AnchorPointOpReq_AnchorPointOpType.
var (
	AnchorPointOpReq_AnchorPointOpType_name = map[int32]string{
		0: "ANCHOR_POINT_OP_TYPE_NONE",
		1: "ANCHOR_POINT_OP_TYPE_TELEPORT",
		2: "ANCHOR_POINT_OP_TYPE_REMOVE",
	}
	AnchorPointOpReq_AnchorPointOpType_value = map[string]int32{
		"ANCHOR_POINT_OP_TYPE_NONE":     0,
		"ANCHOR_POINT_OP_TYPE_TELEPORT": 1,
		"ANCHOR_POINT_OP_TYPE_REMOVE":   2,
	}
)

func (x AnchorPointOpReq_AnchorPointOpType) Enum() *AnchorPointOpReq_AnchorPointOpType {
	p := new(AnchorPointOpReq_AnchorPointOpType)
	*p = x
	return p
}

func (x AnchorPointOpReq_AnchorPointOpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnchorPointOpReq_AnchorPointOpType) Descriptor() protoreflect.EnumDescriptor {
	return file_AnchorPointOpReq_proto_enumTypes[0].Descriptor()
}

func (AnchorPointOpReq_AnchorPointOpType) Type() protoreflect.EnumType {
	return &file_AnchorPointOpReq_proto_enumTypes[0]
}

func (x AnchorPointOpReq_AnchorPointOpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnchorPointOpReq_AnchorPointOpType.Descriptor instead.
func (AnchorPointOpReq_AnchorPointOpType) EnumDescriptor() ([]byte, []int) {
	return file_AnchorPointOpReq_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 4257
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type AnchorPointOpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnchorPointId     uint32 `protobuf:"varint,9,opt,name=anchor_point_id,json=anchorPointId,proto3" json:"anchor_point_id,omitempty"`
	AnchorPointOpType uint32 `protobuf:"varint,12,opt,name=anchor_point_op_type,json=anchorPointOpType,proto3" json:"anchor_point_op_type,omitempty"`
}

func (x *AnchorPointOpReq) Reset() {
	*x = AnchorPointOpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AnchorPointOpReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnchorPointOpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnchorPointOpReq) ProtoMessage() {}

func (x *AnchorPointOpReq) ProtoReflect() protoreflect.Message {
	mi := &file_AnchorPointOpReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnchorPointOpReq.ProtoReflect.Descriptor instead.
func (*AnchorPointOpReq) Descriptor() ([]byte, []int) {
	return file_AnchorPointOpReq_proto_rawDescGZIP(), []int{0}
}

func (x *AnchorPointOpReq) GetAnchorPointId() uint32 {
	if x != nil {
		return x.AnchorPointId
	}
	return 0
}

func (x *AnchorPointOpReq) GetAnchorPointOpType() uint32 {
	if x != nil {
		return x.AnchorPointOpType
	}
	return 0
}

var File_AnchorPointOpReq_proto protoreflect.FileDescriptor

var file_AnchorPointOpReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x10, 0x41, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a,
	0x0f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x76, 0x0a, 0x11, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x41,
	0x4e, 0x43, 0x48, 0x4f, 0x52, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x4f, 0x50, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x4e,
	0x43, 0x48, 0x4f, 0x52, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x4f, 0x50, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x54, 0x45, 0x4c, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12, 0x1f, 0x0a,
	0x1b, 0x41, 0x4e, 0x43, 0x48, 0x4f, 0x52, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x4f, 0x50,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_AnchorPointOpReq_proto_rawDescOnce sync.Once
	file_AnchorPointOpReq_proto_rawDescData = file_AnchorPointOpReq_proto_rawDesc
)

func file_AnchorPointOpReq_proto_rawDescGZIP() []byte {
	file_AnchorPointOpReq_proto_rawDescOnce.Do(func() {
		file_AnchorPointOpReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AnchorPointOpReq_proto_rawDescData)
	})
	return file_AnchorPointOpReq_proto_rawDescData
}

var file_AnchorPointOpReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AnchorPointOpReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AnchorPointOpReq_proto_goTypes = []interface{}{
	(AnchorPointOpReq_AnchorPointOpType)(0), // 0: AnchorPointOpReq.AnchorPointOpType
	(*AnchorPointOpReq)(nil),                // 1: AnchorPointOpReq
}
var file_AnchorPointOpReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AnchorPointOpReq_proto_init() }
func file_AnchorPointOpReq_proto_init() {
	if File_AnchorPointOpReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AnchorPointOpReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnchorPointOpReq); i {
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
			RawDescriptor: file_AnchorPointOpReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AnchorPointOpReq_proto_goTypes,
		DependencyIndexes: file_AnchorPointOpReq_proto_depIdxs,
		EnumInfos:         file_AnchorPointOpReq_proto_enumTypes,
		MessageInfos:      file_AnchorPointOpReq_proto_msgTypes,
	}.Build()
	File_AnchorPointOpReq_proto = out.File
	file_AnchorPointOpReq_proto_rawDesc = nil
	file_AnchorPointOpReq_proto_goTypes = nil
	file_AnchorPointOpReq_proto_depIdxs = nil
}
