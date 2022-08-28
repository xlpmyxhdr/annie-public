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
// source: Unk3000_NJNPNJDFEOL.proto

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

// CmdId: 6112
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type Unk3000_NJNPNJDFEOL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unk3000_NJNPNJDFEOL) Reset() {
	*x = Unk3000_NJNPNJDFEOL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_NJNPNJDFEOL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_NJNPNJDFEOL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_NJNPNJDFEOL) ProtoMessage() {}

func (x *Unk3000_NJNPNJDFEOL) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_NJNPNJDFEOL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_NJNPNJDFEOL.ProtoReflect.Descriptor instead.
func (*Unk3000_NJNPNJDFEOL) Descriptor() ([]byte, []int) {
	return file_Unk3000_NJNPNJDFEOL_proto_rawDescGZIP(), []int{0}
}

var File_Unk3000_NJNPNJDFEOL_proto protoreflect.FileDescriptor

var file_Unk3000_NJNPNJDFEOL_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4e, 0x4a, 0x4e, 0x50, 0x4e, 0x4a,
	0x44, 0x46, 0x45, 0x4f, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4e, 0x4a, 0x4e, 0x50, 0x4e, 0x4a, 0x44, 0x46, 0x45,
	0x4f, 0x4c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_NJNPNJDFEOL_proto_rawDescOnce sync.Once
	file_Unk3000_NJNPNJDFEOL_proto_rawDescData = file_Unk3000_NJNPNJDFEOL_proto_rawDesc
)

func file_Unk3000_NJNPNJDFEOL_proto_rawDescGZIP() []byte {
	file_Unk3000_NJNPNJDFEOL_proto_rawDescOnce.Do(func() {
		file_Unk3000_NJNPNJDFEOL_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_NJNPNJDFEOL_proto_rawDescData)
	})
	return file_Unk3000_NJNPNJDFEOL_proto_rawDescData
}

var file_Unk3000_NJNPNJDFEOL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_NJNPNJDFEOL_proto_goTypes = []interface{}{
	(*Unk3000_NJNPNJDFEOL)(nil), // 0: Unk3000_NJNPNJDFEOL
}
var file_Unk3000_NJNPNJDFEOL_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk3000_NJNPNJDFEOL_proto_init() }
func file_Unk3000_NJNPNJDFEOL_proto_init() {
	if File_Unk3000_NJNPNJDFEOL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_NJNPNJDFEOL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_NJNPNJDFEOL); i {
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
			RawDescriptor: file_Unk3000_NJNPNJDFEOL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_NJNPNJDFEOL_proto_goTypes,
		DependencyIndexes: file_Unk3000_NJNPNJDFEOL_proto_depIdxs,
		MessageInfos:      file_Unk3000_NJNPNJDFEOL_proto_msgTypes,
	}.Build()
	File_Unk3000_NJNPNJDFEOL_proto = out.File
	file_Unk3000_NJNPNJDFEOL_proto_rawDesc = nil
	file_Unk3000_NJNPNJDFEOL_proto_goTypes = nil
	file_Unk3000_NJNPNJDFEOL_proto_depIdxs = nil
}
