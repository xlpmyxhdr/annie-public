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
// source: PingReq.proto

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

// CmdId: 7
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTime    uint32  `protobuf:"varint,12,opt,name=client_time,json=clientTime,proto3" json:"client_time,omitempty"`
	UeTime        float32 `protobuf:"fixed32,14,opt,name=ue_time,json=ueTime,proto3" json:"ue_time,omitempty"`
	TotalTickTime float64 `protobuf:"fixed64,6,opt,name=total_tick_time,json=totalTickTime,proto3" json:"total_tick_time,omitempty"`
	ScData        []byte  `protobuf:"bytes,10,opt,name=sc_data,json=scData,proto3" json:"sc_data,omitempty"`
	Seq           uint32  `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PingReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_PingReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_PingReq_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetClientTime() uint32 {
	if x != nil {
		return x.ClientTime
	}
	return 0
}

func (x *PingReq) GetUeTime() float32 {
	if x != nil {
		return x.UeTime
	}
	return 0
}

func (x *PingReq) GetTotalTickTime() float64 {
	if x != nil {
		return x.TotalTickTime
	}
	return 0
}

func (x *PingReq) GetScData() []byte {
	if x != nil {
		return x.ScData
	}
	return nil
}

func (x *PingReq) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_PingReq_proto protoreflect.FileDescriptor

var file_PingReq_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x96, 0x01, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x75,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x73, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PingReq_proto_rawDescOnce sync.Once
	file_PingReq_proto_rawDescData = file_PingReq_proto_rawDesc
)

func file_PingReq_proto_rawDescGZIP() []byte {
	file_PingReq_proto_rawDescOnce.Do(func() {
		file_PingReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PingReq_proto_rawDescData)
	})
	return file_PingReq_proto_rawDescData
}

var file_PingReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PingReq_proto_goTypes = []interface{}{
	(*PingReq)(nil), // 0: PingReq
}
var file_PingReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PingReq_proto_init() }
func file_PingReq_proto_init() {
	if File_PingReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PingReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
			RawDescriptor: file_PingReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PingReq_proto_goTypes,
		DependencyIndexes: file_PingReq_proto_depIdxs,
		MessageInfos:      file_PingReq_proto_msgTypes,
	}.Build()
	File_PingReq_proto = out.File
	file_PingReq_proto_rawDesc = nil
	file_PingReq_proto_goTypes = nil
	file_PingReq_proto_depIdxs = nil
}
