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
// source: ReportTrackingIOInfoNotify.proto

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

// CmdId: 4129
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ReportTrackingIOInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rydevicetype string `protobuf:"bytes,12,opt,name=rydevicetype,proto3" json:"rydevicetype,omitempty"`
	Deviceid     string `protobuf:"bytes,1,opt,name=deviceid,proto3" json:"deviceid,omitempty"`
	ClientTz     string `protobuf:"bytes,13,opt,name=client_tz,json=clientTz,proto3" json:"client_tz,omitempty"`
	Appid        string `protobuf:"bytes,14,opt,name=appid,proto3" json:"appid,omitempty"`
	Mac          string `protobuf:"bytes,15,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *ReportTrackingIOInfoNotify) Reset() {
	*x = ReportTrackingIOInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReportTrackingIOInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportTrackingIOInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportTrackingIOInfoNotify) ProtoMessage() {}

func (x *ReportTrackingIOInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ReportTrackingIOInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportTrackingIOInfoNotify.ProtoReflect.Descriptor instead.
func (*ReportTrackingIOInfoNotify) Descriptor() ([]byte, []int) {
	return file_ReportTrackingIOInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ReportTrackingIOInfoNotify) GetRydevicetype() string {
	if x != nil {
		return x.Rydevicetype
	}
	return ""
}

func (x *ReportTrackingIOInfoNotify) GetDeviceid() string {
	if x != nil {
		return x.Deviceid
	}
	return ""
}

func (x *ReportTrackingIOInfoNotify) GetClientTz() string {
	if x != nil {
		return x.ClientTz
	}
	return ""
}

func (x *ReportTrackingIOInfoNotify) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *ReportTrackingIOInfoNotify) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

var File_ReportTrackingIOInfoNotify_proto protoreflect.FileDescriptor

var file_ReportTrackingIOInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x4f, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x4f, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x79, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x79, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x7a, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x7a, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReportTrackingIOInfoNotify_proto_rawDescOnce sync.Once
	file_ReportTrackingIOInfoNotify_proto_rawDescData = file_ReportTrackingIOInfoNotify_proto_rawDesc
)

func file_ReportTrackingIOInfoNotify_proto_rawDescGZIP() []byte {
	file_ReportTrackingIOInfoNotify_proto_rawDescOnce.Do(func() {
		file_ReportTrackingIOInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReportTrackingIOInfoNotify_proto_rawDescData)
	})
	return file_ReportTrackingIOInfoNotify_proto_rawDescData
}

var file_ReportTrackingIOInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReportTrackingIOInfoNotify_proto_goTypes = []interface{}{
	(*ReportTrackingIOInfoNotify)(nil), // 0: ReportTrackingIOInfoNotify
}
var file_ReportTrackingIOInfoNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReportTrackingIOInfoNotify_proto_init() }
func file_ReportTrackingIOInfoNotify_proto_init() {
	if File_ReportTrackingIOInfoNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReportTrackingIOInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportTrackingIOInfoNotify); i {
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
			RawDescriptor: file_ReportTrackingIOInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReportTrackingIOInfoNotify_proto_goTypes,
		DependencyIndexes: file_ReportTrackingIOInfoNotify_proto_depIdxs,
		MessageInfos:      file_ReportTrackingIOInfoNotify_proto_msgTypes,
	}.Build()
	File_ReportTrackingIOInfoNotify_proto = out.File
	file_ReportTrackingIOInfoNotify_proto_rawDesc = nil
	file_ReportTrackingIOInfoNotify_proto_goTypes = nil
	file_ReportTrackingIOInfoNotify_proto_depIdxs = nil
}
