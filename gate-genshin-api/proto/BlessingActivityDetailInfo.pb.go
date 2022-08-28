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
// source: BlessingActivityDetailInfo.proto

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

type BlessingActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurDayScanType   uint32            `protobuf:"varint,9,opt,name=cur_day_scan_type,json=curDayScanType,proto3" json:"cur_day_scan_type,omitempty"`
	IsContentClosed  bool              `protobuf:"varint,11,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	PicNumMap        map[uint32]uint32 `protobuf:"bytes,15,rep,name=pic_num_map,json=picNumMap,proto3" json:"pic_num_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ContentCloseTime uint32            `protobuf:"varint,2,opt,name=content_close_time,json=contentCloseTime,proto3" json:"content_close_time,omitempty"`
	CurDayScanNum    uint32            `protobuf:"varint,4,opt,name=cur_day_scan_num,json=curDayScanNum,proto3" json:"cur_day_scan_num,omitempty"`
	RedeemRewardNum  uint32            `protobuf:"varint,1,opt,name=redeem_reward_num,json=redeemRewardNum,proto3" json:"redeem_reward_num,omitempty"`
	IsActivated      bool              `protobuf:"varint,13,opt,name=is_activated,json=isActivated,proto3" json:"is_activated,omitempty"`
	NextRefreshTime  uint32            `protobuf:"varint,6,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
}

func (x *BlessingActivityDetailInfo) Reset() {
	*x = BlessingActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlessingActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlessingActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlessingActivityDetailInfo) ProtoMessage() {}

func (x *BlessingActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BlessingActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlessingActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*BlessingActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_BlessingActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BlessingActivityDetailInfo) GetCurDayScanType() uint32 {
	if x != nil {
		return x.CurDayScanType
	}
	return 0
}

func (x *BlessingActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *BlessingActivityDetailInfo) GetPicNumMap() map[uint32]uint32 {
	if x != nil {
		return x.PicNumMap
	}
	return nil
}

func (x *BlessingActivityDetailInfo) GetContentCloseTime() uint32 {
	if x != nil {
		return x.ContentCloseTime
	}
	return 0
}

func (x *BlessingActivityDetailInfo) GetCurDayScanNum() uint32 {
	if x != nil {
		return x.CurDayScanNum
	}
	return 0
}

func (x *BlessingActivityDetailInfo) GetRedeemRewardNum() uint32 {
	if x != nil {
		return x.RedeemRewardNum
	}
	return 0
}

func (x *BlessingActivityDetailInfo) GetIsActivated() bool {
	if x != nil {
		return x.IsActivated
	}
	return false
}

func (x *BlessingActivityDetailInfo) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

var File_BlessingActivityDetailInfo_proto protoreflect.FileDescriptor

var file_BlessingActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x03, 0x0a, 0x1a, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x73, 0x63, 0x61,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x75,
	0x72, 0x44, 0x61, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x5f,
	0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x69, 0x63, 0x4e, 0x75,
	0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x70, 0x69, 0x63, 0x4e, 0x75,
	0x6d, 0x4d, 0x61, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x73, 0x63,
	0x61, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x75,
	0x72, 0x44, 0x61, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x72,
	0x65, 0x64, 0x65, 0x65, 0x6d, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlessingActivityDetailInfo_proto_rawDescOnce sync.Once
	file_BlessingActivityDetailInfo_proto_rawDescData = file_BlessingActivityDetailInfo_proto_rawDesc
)

func file_BlessingActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_BlessingActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_BlessingActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlessingActivityDetailInfo_proto_rawDescData)
	})
	return file_BlessingActivityDetailInfo_proto_rawDescData
}

var file_BlessingActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_BlessingActivityDetailInfo_proto_goTypes = []interface{}{
	(*BlessingActivityDetailInfo)(nil), // 0: BlessingActivityDetailInfo
	nil,                                // 1: BlessingActivityDetailInfo.PicNumMapEntry
}
var file_BlessingActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: BlessingActivityDetailInfo.pic_num_map:type_name -> BlessingActivityDetailInfo.PicNumMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BlessingActivityDetailInfo_proto_init() }
func file_BlessingActivityDetailInfo_proto_init() {
	if File_BlessingActivityDetailInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BlessingActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlessingActivityDetailInfo); i {
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
			RawDescriptor: file_BlessingActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BlessingActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_BlessingActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_BlessingActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_BlessingActivityDetailInfo_proto = out.File
	file_BlessingActivityDetailInfo_proto_rawDesc = nil
	file_BlessingActivityDetailInfo_proto_goTypes = nil
	file_BlessingActivityDetailInfo_proto_depIdxs = nil
}
