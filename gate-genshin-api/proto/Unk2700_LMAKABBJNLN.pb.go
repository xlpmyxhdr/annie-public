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
// source: Unk2700_LMAKABBJNLN.proto

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

// CmdId: 8253
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_LMAKABBJNLN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             int32                  `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Unk2700_COOFMKLNBND []*Unk2700_FGJFFMPOJON `protobuf:"bytes,11,rep,name=Unk2700_COOFMKLNBND,json=Unk2700COOFMKLNBND,proto3" json:"Unk2700_COOFMKLNBND,omitempty"`
	ScheduleId          uint32                 `protobuf:"varint,10,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *Unk2700_LMAKABBJNLN) Reset() {
	*x = Unk2700_LMAKABBJNLN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_LMAKABBJNLN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_LMAKABBJNLN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_LMAKABBJNLN) ProtoMessage() {}

func (x *Unk2700_LMAKABBJNLN) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_LMAKABBJNLN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_LMAKABBJNLN.ProtoReflect.Descriptor instead.
func (*Unk2700_LMAKABBJNLN) Descriptor() ([]byte, []int) {
	return file_Unk2700_LMAKABBJNLN_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_LMAKABBJNLN) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *Unk2700_LMAKABBJNLN) GetUnk2700_COOFMKLNBND() []*Unk2700_FGJFFMPOJON {
	if x != nil {
		return x.Unk2700_COOFMKLNBND
	}
	return nil
}

func (x *Unk2700_LMAKABBJNLN) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

var File_Unk2700_LMAKABBJNLN_proto protoreflect.FileDescriptor

var file_Unk2700_LMAKABBJNLN_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4d, 0x41, 0x4b, 0x41, 0x42,
	0x42, 0x4a, 0x4e, 0x4c, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x47, 0x4a, 0x46, 0x46, 0x4d, 0x50, 0x4f, 0x4a, 0x4f, 0x4e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x4c, 0x4d, 0x41, 0x4b, 0x41, 0x42, 0x42, 0x4a, 0x4e, 0x4c, 0x4e, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x43, 0x4f, 0x4f, 0x46, 0x4d, 0x4b, 0x4c, 0x4e, 0x42, 0x4e, 0x44, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x46, 0x47, 0x4a, 0x46, 0x46, 0x4d, 0x50, 0x4f, 0x4a, 0x4f, 0x4e, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x43, 0x4f, 0x4f, 0x46, 0x4d, 0x4b, 0x4c, 0x4e, 0x42, 0x4e, 0x44, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_LMAKABBJNLN_proto_rawDescOnce sync.Once
	file_Unk2700_LMAKABBJNLN_proto_rawDescData = file_Unk2700_LMAKABBJNLN_proto_rawDesc
)

func file_Unk2700_LMAKABBJNLN_proto_rawDescGZIP() []byte {
	file_Unk2700_LMAKABBJNLN_proto_rawDescOnce.Do(func() {
		file_Unk2700_LMAKABBJNLN_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_LMAKABBJNLN_proto_rawDescData)
	})
	return file_Unk2700_LMAKABBJNLN_proto_rawDescData
}

var file_Unk2700_LMAKABBJNLN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_LMAKABBJNLN_proto_goTypes = []interface{}{
	(*Unk2700_LMAKABBJNLN)(nil), // 0: Unk2700_LMAKABBJNLN
	(*Unk2700_FGJFFMPOJON)(nil), // 1: Unk2700_FGJFFMPOJON
}
var file_Unk2700_LMAKABBJNLN_proto_depIdxs = []int32{
	1, // 0: Unk2700_LMAKABBJNLN.Unk2700_COOFMKLNBND:type_name -> Unk2700_FGJFFMPOJON
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk2700_LMAKABBJNLN_proto_init() }
func file_Unk2700_LMAKABBJNLN_proto_init() {
	if File_Unk2700_LMAKABBJNLN_proto != nil {
		return
	}
	file_Unk2700_FGJFFMPOJON_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_LMAKABBJNLN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_LMAKABBJNLN); i {
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
			RawDescriptor: file_Unk2700_LMAKABBJNLN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_LMAKABBJNLN_proto_goTypes,
		DependencyIndexes: file_Unk2700_LMAKABBJNLN_proto_depIdxs,
		MessageInfos:      file_Unk2700_LMAKABBJNLN_proto_msgTypes,
	}.Build()
	File_Unk2700_LMAKABBJNLN_proto = out.File
	file_Unk2700_LMAKABBJNLN_proto_rawDesc = nil
	file_Unk2700_LMAKABBJNLN_proto_goTypes = nil
	file_Unk2700_LMAKABBJNLN_proto_depIdxs = nil
}
