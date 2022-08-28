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
// source: MapMarkPoint.proto

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

type MapMarkPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId   uint32           `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	Name      string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pos       *Vector          `protobuf:"bytes,3,opt,name=pos,proto3" json:"pos,omitempty"`
	PointType MapMarkPointType `protobuf:"varint,4,opt,name=point_type,json=pointType,proto3,enum=MapMarkPointType" json:"point_type,omitempty"`
	MonsterId uint32           `protobuf:"varint,5,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
	FromType  MapMarkFromType  `protobuf:"varint,6,opt,name=from_type,json=fromType,proto3,enum=MapMarkFromType" json:"from_type,omitempty"`
	QuestId   uint32           `protobuf:"varint,7,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
}

func (x *MapMarkPoint) Reset() {
	*x = MapMarkPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MapMarkPoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMarkPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMarkPoint) ProtoMessage() {}

func (x *MapMarkPoint) ProtoReflect() protoreflect.Message {
	mi := &file_MapMarkPoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMarkPoint.ProtoReflect.Descriptor instead.
func (*MapMarkPoint) Descriptor() ([]byte, []int) {
	return file_MapMarkPoint_proto_rawDescGZIP(), []int{0}
}

func (x *MapMarkPoint) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *MapMarkPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MapMarkPoint) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *MapMarkPoint) GetPointType() MapMarkPointType {
	if x != nil {
		return x.PointType
	}
	return MapMarkPointType_MAP_MARK_POINT_TYPE_NPC
}

func (x *MapMarkPoint) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

func (x *MapMarkPoint) GetFromType() MapMarkFromType {
	if x != nil {
		return x.FromType
	}
	return MapMarkFromType_MAP_MARK_FROM_TYPE_NONE
}

func (x *MapMarkPoint) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

var File_MapMarkPoint_proto protoreflect.FileDescriptor

var file_MapMarkPoint_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x46, 0x72, 0x6f,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x4d, 0x61, 0x70,
	0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x30, 0x0a, 0x0a,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MapMarkPoint_proto_rawDescOnce sync.Once
	file_MapMarkPoint_proto_rawDescData = file_MapMarkPoint_proto_rawDesc
)

func file_MapMarkPoint_proto_rawDescGZIP() []byte {
	file_MapMarkPoint_proto_rawDescOnce.Do(func() {
		file_MapMarkPoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_MapMarkPoint_proto_rawDescData)
	})
	return file_MapMarkPoint_proto_rawDescData
}

var file_MapMarkPoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MapMarkPoint_proto_goTypes = []interface{}{
	(*MapMarkPoint)(nil),  // 0: MapMarkPoint
	(*Vector)(nil),        // 1: Vector
	(MapMarkPointType)(0), // 2: MapMarkPointType
	(MapMarkFromType)(0),  // 3: MapMarkFromType
}
var file_MapMarkPoint_proto_depIdxs = []int32{
	1, // 0: MapMarkPoint.pos:type_name -> Vector
	2, // 1: MapMarkPoint.point_type:type_name -> MapMarkPointType
	3, // 2: MapMarkPoint.from_type:type_name -> MapMarkFromType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MapMarkPoint_proto_init() }
func file_MapMarkPoint_proto_init() {
	if File_MapMarkPoint_proto != nil {
		return
	}
	file_MapMarkFromType_proto_init()
	file_MapMarkPointType_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MapMarkPoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMarkPoint); i {
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
			RawDescriptor: file_MapMarkPoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MapMarkPoint_proto_goTypes,
		DependencyIndexes: file_MapMarkPoint_proto_depIdxs,
		MessageInfos:      file_MapMarkPoint_proto_msgTypes,
	}.Build()
	File_MapMarkPoint_proto = out.File
	file_MapMarkPoint_proto_rawDesc = nil
	file_MapMarkPoint_proto_goTypes = nil
	file_MapMarkPoint_proto_depIdxs = nil
}
