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
// source: MusicRecord.proto

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

type MusicRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_MBJFOAGKKDJ []*Unk2700_AAAMOFPACEA `protobuf:"bytes,4,rep,name=Unk2700_MBJFOAGKKDJ,json=Unk2700MBJFOAGKKDJ,proto3" json:"Unk2700_MBJFOAGKKDJ,omitempty"`
	Unk2700_DFIBAIILJHN uint32                 `protobuf:"varint,13,opt,name=Unk2700_DFIBAIILJHN,json=Unk2700DFIBAIILJHN,proto3" json:"Unk2700_DFIBAIILJHN,omitempty"`
}

func (x *MusicRecord) Reset() {
	*x = MusicRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MusicRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MusicRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MusicRecord) ProtoMessage() {}

func (x *MusicRecord) ProtoReflect() protoreflect.Message {
	mi := &file_MusicRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MusicRecord.ProtoReflect.Descriptor instead.
func (*MusicRecord) Descriptor() ([]byte, []int) {
	return file_MusicRecord_proto_rawDescGZIP(), []int{0}
}

func (x *MusicRecord) GetUnk2700_MBJFOAGKKDJ() []*Unk2700_AAAMOFPACEA {
	if x != nil {
		return x.Unk2700_MBJFOAGKKDJ
	}
	return nil
}

func (x *MusicRecord) GetUnk2700_DFIBAIILJHN() uint32 {
	if x != nil {
		return x.Unk2700_DFIBAIILJHN
	}
	return 0
}

var File_MusicRecord_proto protoreflect.FileDescriptor

var file_MusicRecord_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x41, 0x41,
	0x4d, 0x4f, 0x46, 0x50, 0x41, 0x43, 0x45, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85,
	0x01, 0x0a, 0x0b, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x45,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x42, 0x4a, 0x46, 0x4f, 0x41,
	0x47, 0x4b, 0x4b, 0x44, 0x4a, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x41, 0x41, 0x4d, 0x4f, 0x46, 0x50, 0x41, 0x43, 0x45,
	0x41, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4d, 0x42, 0x4a, 0x46, 0x4f, 0x41,
	0x47, 0x4b, 0x4b, 0x44, 0x4a, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x44, 0x46, 0x49, 0x42, 0x41, 0x49, 0x49, 0x4c, 0x4a, 0x48, 0x4e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x46, 0x49, 0x42, 0x41,
	0x49, 0x49, 0x4c, 0x4a, 0x48, 0x4e, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MusicRecord_proto_rawDescOnce sync.Once
	file_MusicRecord_proto_rawDescData = file_MusicRecord_proto_rawDesc
)

func file_MusicRecord_proto_rawDescGZIP() []byte {
	file_MusicRecord_proto_rawDescOnce.Do(func() {
		file_MusicRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_MusicRecord_proto_rawDescData)
	})
	return file_MusicRecord_proto_rawDescData
}

var file_MusicRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MusicRecord_proto_goTypes = []interface{}{
	(*MusicRecord)(nil),         // 0: MusicRecord
	(*Unk2700_AAAMOFPACEA)(nil), // 1: Unk2700_AAAMOFPACEA
}
var file_MusicRecord_proto_depIdxs = []int32{
	1, // 0: MusicRecord.Unk2700_MBJFOAGKKDJ:type_name -> Unk2700_AAAMOFPACEA
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MusicRecord_proto_init() }
func file_MusicRecord_proto_init() {
	if File_MusicRecord_proto != nil {
		return
	}
	file_Unk2700_AAAMOFPACEA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MusicRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MusicRecord); i {
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
			RawDescriptor: file_MusicRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MusicRecord_proto_goTypes,
		DependencyIndexes: file_MusicRecord_proto_depIdxs,
		MessageInfos:      file_MusicRecord_proto_msgTypes,
	}.Build()
	File_MusicRecord_proto = out.File
	file_MusicRecord_proto_rawDesc = nil
	file_MusicRecord_proto_goTypes = nil
	file_MusicRecord_proto_depIdxs = nil
}
