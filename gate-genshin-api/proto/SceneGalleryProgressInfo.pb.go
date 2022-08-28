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
// source: SceneGalleryProgressInfo.proto

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

type SceneGalleryProgressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProgressStageList []uint32 `protobuf:"varint,8,rep,packed,name=progress_stage_list,json=progressStageList,proto3" json:"progress_stage_list,omitempty"`
	Key               string   `protobuf:"bytes,11,opt,name=key,proto3" json:"key,omitempty"`
	Progress          uint32   `protobuf:"varint,5,opt,name=progress,proto3" json:"progress,omitempty"`
	UiForm            uint32   `protobuf:"varint,12,opt,name=ui_form,json=uiForm,proto3" json:"ui_form,omitempty"`
}

func (x *SceneGalleryProgressInfo) Reset() {
	*x = SceneGalleryProgressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGalleryProgressInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryProgressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryProgressInfo) ProtoMessage() {}

func (x *SceneGalleryProgressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGalleryProgressInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryProgressInfo.ProtoReflect.Descriptor instead.
func (*SceneGalleryProgressInfo) Descriptor() ([]byte, []int) {
	return file_SceneGalleryProgressInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryProgressInfo) GetProgressStageList() []uint32 {
	if x != nil {
		return x.ProgressStageList
	}
	return nil
}

func (x *SceneGalleryProgressInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SceneGalleryProgressInfo) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *SceneGalleryProgressInfo) GetUiForm() uint32 {
	if x != nil {
		return x.UiForm
	}
	return 0
}

var File_SceneGalleryProgressInfo_proto protoreflect.FileDescriptor

var file_SceneGalleryProgressInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x91, 0x01, 0x0a, 0x18, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x69, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69,
	0x46, 0x6f, 0x72, 0x6d, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGalleryProgressInfo_proto_rawDescOnce sync.Once
	file_SceneGalleryProgressInfo_proto_rawDescData = file_SceneGalleryProgressInfo_proto_rawDesc
)

func file_SceneGalleryProgressInfo_proto_rawDescGZIP() []byte {
	file_SceneGalleryProgressInfo_proto_rawDescOnce.Do(func() {
		file_SceneGalleryProgressInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGalleryProgressInfo_proto_rawDescData)
	})
	return file_SceneGalleryProgressInfo_proto_rawDescData
}

var file_SceneGalleryProgressInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGalleryProgressInfo_proto_goTypes = []interface{}{
	(*SceneGalleryProgressInfo)(nil), // 0: SceneGalleryProgressInfo
}
var file_SceneGalleryProgressInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneGalleryProgressInfo_proto_init() }
func file_SceneGalleryProgressInfo_proto_init() {
	if File_SceneGalleryProgressInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneGalleryProgressInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryProgressInfo); i {
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
			RawDescriptor: file_SceneGalleryProgressInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGalleryProgressInfo_proto_goTypes,
		DependencyIndexes: file_SceneGalleryProgressInfo_proto_depIdxs,
		MessageInfos:      file_SceneGalleryProgressInfo_proto_msgTypes,
	}.Build()
	File_SceneGalleryProgressInfo_proto = out.File
	file_SceneGalleryProgressInfo_proto_rawDesc = nil
	file_SceneGalleryProgressInfo_proto_goTypes = nil
	file_SceneGalleryProgressInfo_proto_depIdxs = nil
}
