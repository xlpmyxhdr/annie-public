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
// source: ResinChangeNotify.proto

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

// CmdId: 642
// EnetChannelId: 0
// EnetIsReliable: true
type ResinChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextAddTimestamp uint32 `protobuf:"varint,6,opt,name=next_add_timestamp,json=nextAddTimestamp,proto3" json:"next_add_timestamp,omitempty"`
	CurBuyCount      uint32 `protobuf:"varint,4,opt,name=cur_buy_count,json=curBuyCount,proto3" json:"cur_buy_count,omitempty"`
	CurValue         uint32 `protobuf:"varint,12,opt,name=cur_value,json=curValue,proto3" json:"cur_value,omitempty"`
}

func (x *ResinChangeNotify) Reset() {
	*x = ResinChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ResinChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResinChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResinChangeNotify) ProtoMessage() {}

func (x *ResinChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ResinChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResinChangeNotify.ProtoReflect.Descriptor instead.
func (*ResinChangeNotify) Descriptor() ([]byte, []int) {
	return file_ResinChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ResinChangeNotify) GetNextAddTimestamp() uint32 {
	if x != nil {
		return x.NextAddTimestamp
	}
	return 0
}

func (x *ResinChangeNotify) GetCurBuyCount() uint32 {
	if x != nil {
		return x.CurBuyCount
	}
	return 0
}

func (x *ResinChangeNotify) GetCurValue() uint32 {
	if x != nil {
		return x.CurValue
	}
	return 0
}

var File_ResinChangeNotify_proto protoreflect.FileDescriptor

var file_ResinChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6e, 0x65, 0x78,
	0x74, 0x41, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a,
	0x0d, 0x63, 0x75, 0x72, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ResinChangeNotify_proto_rawDescOnce sync.Once
	file_ResinChangeNotify_proto_rawDescData = file_ResinChangeNotify_proto_rawDesc
)

func file_ResinChangeNotify_proto_rawDescGZIP() []byte {
	file_ResinChangeNotify_proto_rawDescOnce.Do(func() {
		file_ResinChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ResinChangeNotify_proto_rawDescData)
	})
	return file_ResinChangeNotify_proto_rawDescData
}

var file_ResinChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ResinChangeNotify_proto_goTypes = []interface{}{
	(*ResinChangeNotify)(nil), // 0: ResinChangeNotify
}
var file_ResinChangeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ResinChangeNotify_proto_init() }
func file_ResinChangeNotify_proto_init() {
	if File_ResinChangeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ResinChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResinChangeNotify); i {
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
			RawDescriptor: file_ResinChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ResinChangeNotify_proto_goTypes,
		DependencyIndexes: file_ResinChangeNotify_proto_depIdxs,
		MessageInfos:      file_ResinChangeNotify_proto_msgTypes,
	}.Build()
	File_ResinChangeNotify_proto = out.File
	file_ResinChangeNotify_proto_rawDesc = nil
	file_ResinChangeNotify_proto_goTypes = nil
	file_ResinChangeNotify_proto_depIdxs = nil
}
