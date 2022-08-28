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
// source: MatchType.proto

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

type MatchType int32

const (
	MatchType_MATCH_TYPE_NONE       MatchType = 0
	MatchType_MATCH_TYPE_DUNGEON    MatchType = 1
	MatchType_MATCH_TYPE_MP_PLAY    MatchType = 2
	MatchType_MATCH_TYPE_MECHANICUS MatchType = 3
	MatchType_MATCH_TYPE_GENERAL    MatchType = 4
)

// Enum value maps for MatchType.
var (
	MatchType_name = map[int32]string{
		0: "MATCH_TYPE_NONE",
		1: "MATCH_TYPE_DUNGEON",
		2: "MATCH_TYPE_MP_PLAY",
		3: "MATCH_TYPE_MECHANICUS",
		4: "MATCH_TYPE_GENERAL",
	}
	MatchType_value = map[string]int32{
		"MATCH_TYPE_NONE":       0,
		"MATCH_TYPE_DUNGEON":    1,
		"MATCH_TYPE_MP_PLAY":    2,
		"MATCH_TYPE_MECHANICUS": 3,
		"MATCH_TYPE_GENERAL":    4,
	}
)

func (x MatchType) Enum() *MatchType {
	p := new(MatchType)
	*p = x
	return p
}

func (x MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_MatchType_proto_enumTypes[0].Descriptor()
}

func (MatchType) Type() protoreflect.EnumType {
	return &file_MatchType_proto_enumTypes[0]
}

func (x MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchType.Descriptor instead.
func (MatchType) EnumDescriptor() ([]byte, []int) {
	return file_MatchType_proto_rawDescGZIP(), []int{0}
}

var File_MatchType_proto protoreflect.FileDescriptor

var file_MatchType_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x83, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x5f, 0x50, 0x4c,
	0x41, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x43, 0x55, 0x53, 0x10, 0x03, 0x12,
	0x16, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45,
	0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MatchType_proto_rawDescOnce sync.Once
	file_MatchType_proto_rawDescData = file_MatchType_proto_rawDesc
)

func file_MatchType_proto_rawDescGZIP() []byte {
	file_MatchType_proto_rawDescOnce.Do(func() {
		file_MatchType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MatchType_proto_rawDescData)
	})
	return file_MatchType_proto_rawDescData
}

var file_MatchType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MatchType_proto_goTypes = []interface{}{
	(MatchType)(0), // 0: MatchType
}
var file_MatchType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MatchType_proto_init() }
func file_MatchType_proto_init() {
	if File_MatchType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MatchType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MatchType_proto_goTypes,
		DependencyIndexes: file_MatchType_proto_depIdxs,
		EnumInfos:         file_MatchType_proto_enumTypes,
	}.Build()
	File_MatchType_proto = out.File
	file_MatchType_proto_rawDesc = nil
	file_MatchType_proto_goTypes = nil
	file_MatchType_proto_depIdxs = nil
}
