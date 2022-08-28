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
// source: FurnitureMakeCancelRsp.proto

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

// CmdId: 4683
// EnetChannelId: 0
// EnetIsReliable: true
type FurnitureMakeCancelRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode           int32              `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MakeId            uint32             `protobuf:"varint,2,opt,name=make_id,json=makeId,proto3" json:"make_id,omitempty"`
	FurnitureMakeSlot *FurnitureMakeSlot `protobuf:"bytes,15,opt,name=furniture_make_slot,json=furnitureMakeSlot,proto3" json:"furniture_make_slot,omitempty"`
}

func (x *FurnitureMakeCancelRsp) Reset() {
	*x = FurnitureMakeCancelRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FurnitureMakeCancelRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FurnitureMakeCancelRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FurnitureMakeCancelRsp) ProtoMessage() {}

func (x *FurnitureMakeCancelRsp) ProtoReflect() protoreflect.Message {
	mi := &file_FurnitureMakeCancelRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FurnitureMakeCancelRsp.ProtoReflect.Descriptor instead.
func (*FurnitureMakeCancelRsp) Descriptor() ([]byte, []int) {
	return file_FurnitureMakeCancelRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FurnitureMakeCancelRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *FurnitureMakeCancelRsp) GetMakeId() uint32 {
	if x != nil {
		return x.MakeId
	}
	return 0
}

func (x *FurnitureMakeCancelRsp) GetFurnitureMakeSlot() *FurnitureMakeSlot {
	if x != nil {
		return x.FurnitureMakeSlot
	}
	return nil
}

var File_FurnitureMakeCancelRsp_proto protoreflect.FileDescriptor

var file_FurnitureMakeCancelRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x46, 0x75, 0x72, 0x6e,
	0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d,
	0x61, 0x6b, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x13, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61,
	0x6b, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x11, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FurnitureMakeCancelRsp_proto_rawDescOnce sync.Once
	file_FurnitureMakeCancelRsp_proto_rawDescData = file_FurnitureMakeCancelRsp_proto_rawDesc
)

func file_FurnitureMakeCancelRsp_proto_rawDescGZIP() []byte {
	file_FurnitureMakeCancelRsp_proto_rawDescOnce.Do(func() {
		file_FurnitureMakeCancelRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_FurnitureMakeCancelRsp_proto_rawDescData)
	})
	return file_FurnitureMakeCancelRsp_proto_rawDescData
}

var file_FurnitureMakeCancelRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FurnitureMakeCancelRsp_proto_goTypes = []interface{}{
	(*FurnitureMakeCancelRsp)(nil), // 0: FurnitureMakeCancelRsp
	(*FurnitureMakeSlot)(nil),      // 1: FurnitureMakeSlot
}
var file_FurnitureMakeCancelRsp_proto_depIdxs = []int32{
	1, // 0: FurnitureMakeCancelRsp.furniture_make_slot:type_name -> FurnitureMakeSlot
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FurnitureMakeCancelRsp_proto_init() }
func file_FurnitureMakeCancelRsp_proto_init() {
	if File_FurnitureMakeCancelRsp_proto != nil {
		return
	}
	file_FurnitureMakeSlot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FurnitureMakeCancelRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FurnitureMakeCancelRsp); i {
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
			RawDescriptor: file_FurnitureMakeCancelRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FurnitureMakeCancelRsp_proto_goTypes,
		DependencyIndexes: file_FurnitureMakeCancelRsp_proto_depIdxs,
		MessageInfos:      file_FurnitureMakeCancelRsp_proto_msgTypes,
	}.Build()
	File_FurnitureMakeCancelRsp_proto = out.File
	file_FurnitureMakeCancelRsp_proto_rawDesc = nil
	file_FurnitureMakeCancelRsp_proto_goTypes = nil
	file_FurnitureMakeCancelRsp_proto_depIdxs = nil
}
