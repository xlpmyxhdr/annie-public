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
// source: Unk2700_HGMOIKODALP_ServerRsp.proto

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

// CmdId: 6220
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_HGMOIKODALP_ServerRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             int32  `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Unk2700_ONOOJBEABOE uint64 `protobuf:"varint,11,opt,name=Unk2700_ONOOJBEABOE,json=Unk2700ONOOJBEABOE,proto3" json:"Unk2700_ONOOJBEABOE,omitempty"`
}

func (x *Unk2700_HGMOIKODALP_ServerRsp) Reset() {
	*x = Unk2700_HGMOIKODALP_ServerRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_HGMOIKODALP_ServerRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_HGMOIKODALP_ServerRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_HGMOIKODALP_ServerRsp) ProtoMessage() {}

func (x *Unk2700_HGMOIKODALP_ServerRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_HGMOIKODALP_ServerRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_HGMOIKODALP_ServerRsp.ProtoReflect.Descriptor instead.
func (*Unk2700_HGMOIKODALP_ServerRsp) Descriptor() ([]byte, []int) {
	return file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_HGMOIKODALP_ServerRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *Unk2700_HGMOIKODALP_ServerRsp) GetUnk2700_ONOOJBEABOE() uint64 {
	if x != nil {
		return x.Unk2700_ONOOJBEABOE
	}
	return 0
}

var File_Unk2700_HGMOIKODALP_ServerRsp_proto protoreflect.FileDescriptor

var file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x47, 0x4d, 0x4f, 0x49, 0x4b,
	0x4f, 0x44, 0x41, 0x4c, 0x50, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x1d, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x48, 0x47, 0x4d, 0x4f, 0x49, 0x4b, 0x4f, 0x44, 0x41, 0x4c, 0x50, 0x5f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x4e, 0x4f, 0x4f,
	0x4a, 0x42, 0x45, 0x41, 0x42, 0x4f, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4f, 0x4e, 0x4f, 0x4f, 0x4a, 0x42, 0x45, 0x41, 0x42, 0x4f,
	0x45, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescOnce sync.Once
	file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescData = file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDesc
)

func file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescGZIP() []byte {
	file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescOnce.Do(func() {
		file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescData)
	})
	return file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDescData
}

var file_Unk2700_HGMOIKODALP_ServerRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_HGMOIKODALP_ServerRsp_proto_goTypes = []interface{}{
	(*Unk2700_HGMOIKODALP_ServerRsp)(nil), // 0: Unk2700_HGMOIKODALP_ServerRsp
}
var file_Unk2700_HGMOIKODALP_ServerRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2700_HGMOIKODALP_ServerRsp_proto_init() }
func file_Unk2700_HGMOIKODALP_ServerRsp_proto_init() {
	if File_Unk2700_HGMOIKODALP_ServerRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_HGMOIKODALP_ServerRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_HGMOIKODALP_ServerRsp); i {
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
			RawDescriptor: file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_HGMOIKODALP_ServerRsp_proto_goTypes,
		DependencyIndexes: file_Unk2700_HGMOIKODALP_ServerRsp_proto_depIdxs,
		MessageInfos:      file_Unk2700_HGMOIKODALP_ServerRsp_proto_msgTypes,
	}.Build()
	File_Unk2700_HGMOIKODALP_ServerRsp_proto = out.File
	file_Unk2700_HGMOIKODALP_ServerRsp_proto_rawDesc = nil
	file_Unk2700_HGMOIKODALP_ServerRsp_proto_goTypes = nil
	file_Unk2700_HGMOIKODALP_ServerRsp_proto_depIdxs = nil
}
