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
// source: ForgeQueueManipulateType.proto

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

type ForgeQueueManipulateType int32

const (
	ForgeQueueManipulateType_FORGE_QUEUE_MANIPULATE_TYPE_RECEIVE_OUTPUT ForgeQueueManipulateType = 0
	ForgeQueueManipulateType_FORGE_QUEUE_MANIPULATE_TYPE_STOP_FORGE     ForgeQueueManipulateType = 1
)

// Enum value maps for ForgeQueueManipulateType.
var (
	ForgeQueueManipulateType_name = map[int32]string{
		0: "FORGE_QUEUE_MANIPULATE_TYPE_RECEIVE_OUTPUT",
		1: "FORGE_QUEUE_MANIPULATE_TYPE_STOP_FORGE",
	}
	ForgeQueueManipulateType_value = map[string]int32{
		"FORGE_QUEUE_MANIPULATE_TYPE_RECEIVE_OUTPUT": 0,
		"FORGE_QUEUE_MANIPULATE_TYPE_STOP_FORGE":     1,
	}
)

func (x ForgeQueueManipulateType) Enum() *ForgeQueueManipulateType {
	p := new(ForgeQueueManipulateType)
	*p = x
	return p
}

func (x ForgeQueueManipulateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ForgeQueueManipulateType) Descriptor() protoreflect.EnumDescriptor {
	return file_ForgeQueueManipulateType_proto_enumTypes[0].Descriptor()
}

func (ForgeQueueManipulateType) Type() protoreflect.EnumType {
	return &file_ForgeQueueManipulateType_proto_enumTypes[0]
}

func (x ForgeQueueManipulateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ForgeQueueManipulateType.Descriptor instead.
func (ForgeQueueManipulateType) EnumDescriptor() ([]byte, []int) {
	return file_ForgeQueueManipulateType_proto_rawDescGZIP(), []int{0}
}

var File_ForgeQueueManipulateType_proto protoreflect.FileDescriptor

var file_ForgeQueueManipulateType_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x61, 0x6e, 0x69,
	0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x76, 0x0a, 0x18, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x61,
	0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x2a,
	0x46, 0x4f, 0x52, 0x47, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x49,
	0x50, 0x55, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x45,
	0x49, 0x56, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26,
	0x46, 0x4f, 0x52, 0x47, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x49,
	0x50, 0x55, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50,
	0x5f, 0x46, 0x4f, 0x52, 0x47, 0x45, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ForgeQueueManipulateType_proto_rawDescOnce sync.Once
	file_ForgeQueueManipulateType_proto_rawDescData = file_ForgeQueueManipulateType_proto_rawDesc
)

func file_ForgeQueueManipulateType_proto_rawDescGZIP() []byte {
	file_ForgeQueueManipulateType_proto_rawDescOnce.Do(func() {
		file_ForgeQueueManipulateType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ForgeQueueManipulateType_proto_rawDescData)
	})
	return file_ForgeQueueManipulateType_proto_rawDescData
}

var file_ForgeQueueManipulateType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ForgeQueueManipulateType_proto_goTypes = []interface{}{
	(ForgeQueueManipulateType)(0), // 0: ForgeQueueManipulateType
}
var file_ForgeQueueManipulateType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ForgeQueueManipulateType_proto_init() }
func file_ForgeQueueManipulateType_proto_init() {
	if File_ForgeQueueManipulateType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ForgeQueueManipulateType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ForgeQueueManipulateType_proto_goTypes,
		DependencyIndexes: file_ForgeQueueManipulateType_proto_depIdxs,
		EnumInfos:         file_ForgeQueueManipulateType_proto_enumTypes,
	}.Build()
	File_ForgeQueueManipulateType_proto = out.File
	file_ForgeQueueManipulateType_proto_rawDesc = nil
	file_ForgeQueueManipulateType_proto_goTypes = nil
	file_ForgeQueueManipulateType_proto_depIdxs = nil
}
