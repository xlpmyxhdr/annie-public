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
// source: ShortAbilityHashPair.proto

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

type ShortAbilityHashPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityConfigHash int32 `protobuf:"fixed32,15,opt,name=ability_config_hash,json=abilityConfigHash,proto3" json:"ability_config_hash,omitempty"`
	AbilityNameHash   int32 `protobuf:"fixed32,1,opt,name=ability_name_hash,json=abilityNameHash,proto3" json:"ability_name_hash,omitempty"`
}

func (x *ShortAbilityHashPair) Reset() {
	*x = ShortAbilityHashPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShortAbilityHashPair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortAbilityHashPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortAbilityHashPair) ProtoMessage() {}

func (x *ShortAbilityHashPair) ProtoReflect() protoreflect.Message {
	mi := &file_ShortAbilityHashPair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortAbilityHashPair.ProtoReflect.Descriptor instead.
func (*ShortAbilityHashPair) Descriptor() ([]byte, []int) {
	return file_ShortAbilityHashPair_proto_rawDescGZIP(), []int{0}
}

func (x *ShortAbilityHashPair) GetAbilityConfigHash() int32 {
	if x != nil {
		return x.AbilityConfigHash
	}
	return 0
}

func (x *ShortAbilityHashPair) GetAbilityNameHash() int32 {
	if x != nil {
		return x.AbilityNameHash
	}
	return 0
}

var File_ShortAbilityHashPair_proto protoreflect.FileDescriptor

var file_ShortAbilityHashPair_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0f, 0x52, 0x11, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52,
	0x0f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShortAbilityHashPair_proto_rawDescOnce sync.Once
	file_ShortAbilityHashPair_proto_rawDescData = file_ShortAbilityHashPair_proto_rawDesc
)

func file_ShortAbilityHashPair_proto_rawDescGZIP() []byte {
	file_ShortAbilityHashPair_proto_rawDescOnce.Do(func() {
		file_ShortAbilityHashPair_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShortAbilityHashPair_proto_rawDescData)
	})
	return file_ShortAbilityHashPair_proto_rawDescData
}

var file_ShortAbilityHashPair_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ShortAbilityHashPair_proto_goTypes = []interface{}{
	(*ShortAbilityHashPair)(nil), // 0: ShortAbilityHashPair
}
var file_ShortAbilityHashPair_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ShortAbilityHashPair_proto_init() }
func file_ShortAbilityHashPair_proto_init() {
	if File_ShortAbilityHashPair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ShortAbilityHashPair_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortAbilityHashPair); i {
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
			RawDescriptor: file_ShortAbilityHashPair_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShortAbilityHashPair_proto_goTypes,
		DependencyIndexes: file_ShortAbilityHashPair_proto_depIdxs,
		MessageInfos:      file_ShortAbilityHashPair_proto_msgTypes,
	}.Build()
	File_ShortAbilityHashPair_proto = out.File
	file_ShortAbilityHashPair_proto_rawDesc = nil
	file_ShortAbilityHashPair_proto_goTypes = nil
	file_ShortAbilityHashPair_proto_depIdxs = nil
}
