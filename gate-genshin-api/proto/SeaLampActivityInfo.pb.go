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
// source: SeaLampActivityInfo.proto

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

type SeaLampActivityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMechanicusOpen         bool                  `protobuf:"varint,14,opt,name=is_mechanicus_open,json=isMechanicusOpen,proto3" json:"is_mechanicus_open,omitempty"`
	DayIndex                 uint32                `protobuf:"varint,1,opt,name=day_index,json=dayIndex,proto3" json:"day_index,omitempty"`
	SectionInfoList          []*SeaLampSectionInfo `protobuf:"bytes,6,rep,name=section_info_list,json=sectionInfoList,proto3" json:"section_info_list,omitempty"`
	Popularity               uint32                `protobuf:"varint,10,opt,name=popularity,proto3" json:"popularity,omitempty"`
	SeaLampCoin              uint32                `protobuf:"varint,15,opt,name=sea_lamp_coin,json=seaLampCoin,proto3" json:"sea_lamp_coin,omitempty"`
	FirstDayStartTime        uint32                `protobuf:"varint,11,opt,name=first_day_start_time,json=firstDayStartTime,proto3" json:"first_day_start_time,omitempty"`
	MechanicusId             uint32                `protobuf:"varint,9,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	IsMechanicusFeatureClose bool                  `protobuf:"varint,12,opt,name=is_mechanicus_feature_close,json=isMechanicusFeatureClose,proto3" json:"is_mechanicus_feature_close,omitempty"`
	IsContentClosed          bool                  `protobuf:"varint,5,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
}

func (x *SeaLampActivityInfo) Reset() {
	*x = SeaLampActivityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SeaLampActivityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeaLampActivityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeaLampActivityInfo) ProtoMessage() {}

func (x *SeaLampActivityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SeaLampActivityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeaLampActivityInfo.ProtoReflect.Descriptor instead.
func (*SeaLampActivityInfo) Descriptor() ([]byte, []int) {
	return file_SeaLampActivityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SeaLampActivityInfo) GetIsMechanicusOpen() bool {
	if x != nil {
		return x.IsMechanicusOpen
	}
	return false
}

func (x *SeaLampActivityInfo) GetDayIndex() uint32 {
	if x != nil {
		return x.DayIndex
	}
	return 0
}

func (x *SeaLampActivityInfo) GetSectionInfoList() []*SeaLampSectionInfo {
	if x != nil {
		return x.SectionInfoList
	}
	return nil
}

func (x *SeaLampActivityInfo) GetPopularity() uint32 {
	if x != nil {
		return x.Popularity
	}
	return 0
}

func (x *SeaLampActivityInfo) GetSeaLampCoin() uint32 {
	if x != nil {
		return x.SeaLampCoin
	}
	return 0
}

func (x *SeaLampActivityInfo) GetFirstDayStartTime() uint32 {
	if x != nil {
		return x.FirstDayStartTime
	}
	return 0
}

func (x *SeaLampActivityInfo) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *SeaLampActivityInfo) GetIsMechanicusFeatureClose() bool {
	if x != nil {
		return x.IsMechanicusFeatureClose
	}
	return false
}

func (x *SeaLampActivityInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

var File_SeaLampActivityInfo_proto protoreflect.FileDescriptor

var file_SeaLampActivityInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x53, 0x65, 0x61,
	0x4c, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x03, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d,
	0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a,
	0x12, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x61, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x64, 0x61, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3f, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70,
	0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70,
	0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x61,
	0x5f, 0x6c, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x73, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x2f, 0x0a,
	0x14, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75,
	0x73, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x73, 0x4d, 0x65, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SeaLampActivityInfo_proto_rawDescOnce sync.Once
	file_SeaLampActivityInfo_proto_rawDescData = file_SeaLampActivityInfo_proto_rawDesc
)

func file_SeaLampActivityInfo_proto_rawDescGZIP() []byte {
	file_SeaLampActivityInfo_proto_rawDescOnce.Do(func() {
		file_SeaLampActivityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SeaLampActivityInfo_proto_rawDescData)
	})
	return file_SeaLampActivityInfo_proto_rawDescData
}

var file_SeaLampActivityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SeaLampActivityInfo_proto_goTypes = []interface{}{
	(*SeaLampActivityInfo)(nil), // 0: SeaLampActivityInfo
	(*SeaLampSectionInfo)(nil),  // 1: SeaLampSectionInfo
}
var file_SeaLampActivityInfo_proto_depIdxs = []int32{
	1, // 0: SeaLampActivityInfo.section_info_list:type_name -> SeaLampSectionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SeaLampActivityInfo_proto_init() }
func file_SeaLampActivityInfo_proto_init() {
	if File_SeaLampActivityInfo_proto != nil {
		return
	}
	file_SeaLampSectionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SeaLampActivityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeaLampActivityInfo); i {
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
			RawDescriptor: file_SeaLampActivityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SeaLampActivityInfo_proto_goTypes,
		DependencyIndexes: file_SeaLampActivityInfo_proto_depIdxs,
		MessageInfos:      file_SeaLampActivityInfo_proto_msgTypes,
	}.Build()
	File_SeaLampActivityInfo_proto = out.File
	file_SeaLampActivityInfo_proto_rawDesc = nil
	file_SeaLampActivityInfo_proto_goTypes = nil
	file_SeaLampActivityInfo_proto_depIdxs = nil
}
