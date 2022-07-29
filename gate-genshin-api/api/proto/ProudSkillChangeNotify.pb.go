// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ProudSkillChangeNotify.proto

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

type ProudSkillChangeNotify_CmdId int32

const (
	ProudSkillChangeNotify_NONE             ProudSkillChangeNotify_CmdId = 0
	ProudSkillChangeNotify_ENET_CHANNEL_ID  ProudSkillChangeNotify_CmdId = 0
	ProudSkillChangeNotify_ENET_IS_RELIABLE ProudSkillChangeNotify_CmdId = 1
	ProudSkillChangeNotify_CMD_ID           ProudSkillChangeNotify_CmdId = 1012
)

// Enum value maps for ProudSkillChangeNotify_CmdId.
var (
	ProudSkillChangeNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1012: "CMD_ID",
	}
	ProudSkillChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1012,
	}
)

func (x ProudSkillChangeNotify_CmdId) Enum() *ProudSkillChangeNotify_CmdId {
	p := new(ProudSkillChangeNotify_CmdId)
	*p = x
	return p
}

func (x ProudSkillChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProudSkillChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ProudSkillChangeNotify_proto_enumTypes[0].Descriptor()
}

func (ProudSkillChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ProudSkillChangeNotify_proto_enumTypes[0]
}

func (x ProudSkillChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProudSkillChangeNotify_CmdId.Descriptor instead.
func (ProudSkillChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ProudSkillChangeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ProudSkillChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid     uint64   `protobuf:"varint,1,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	EntityId       uint32   `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	SkillDepotId   uint32   `protobuf:"varint,3,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	ProudSkillList []uint32 `protobuf:"varint,4,rep,packed,name=proud_skill_list,json=proudSkillList,proto3" json:"proud_skill_list,omitempty"`
}

func (x *ProudSkillChangeNotify) Reset() {
	*x = ProudSkillChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProudSkillChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProudSkillChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProudSkillChangeNotify) ProtoMessage() {}

func (x *ProudSkillChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ProudSkillChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProudSkillChangeNotify.ProtoReflect.Descriptor instead.
func (*ProudSkillChangeNotify) Descriptor() ([]byte, []int) {
	return file_ProudSkillChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ProudSkillChangeNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *ProudSkillChangeNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ProudSkillChangeNotify) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *ProudSkillChangeNotify) GetProudSkillList() []uint32 {
	if x != nil {
		return x.ProudSkillList
	}
	return nil
}

var File_ProudSkillChangeNotify_proto protoreflect.FileDescriptor

var file_ProudSkillChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x70,
	0x6f, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf4, 0x07, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ProudSkillChangeNotify_proto_rawDescOnce sync.Once
	file_ProudSkillChangeNotify_proto_rawDescData = file_ProudSkillChangeNotify_proto_rawDesc
)

func file_ProudSkillChangeNotify_proto_rawDescGZIP() []byte {
	file_ProudSkillChangeNotify_proto_rawDescOnce.Do(func() {
		file_ProudSkillChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProudSkillChangeNotify_proto_rawDescData)
	})
	return file_ProudSkillChangeNotify_proto_rawDescData
}

var file_ProudSkillChangeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ProudSkillChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ProudSkillChangeNotify_proto_goTypes = []interface{}{
	(ProudSkillChangeNotify_CmdId)(0), // 0: proto.ProudSkillChangeNotify.CmdId
	(*ProudSkillChangeNotify)(nil),    // 1: proto.ProudSkillChangeNotify
}
var file_ProudSkillChangeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProudSkillChangeNotify_proto_init() }
func file_ProudSkillChangeNotify_proto_init() {
	if File_ProudSkillChangeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProudSkillChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProudSkillChangeNotify); i {
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
			RawDescriptor: file_ProudSkillChangeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProudSkillChangeNotify_proto_goTypes,
		DependencyIndexes: file_ProudSkillChangeNotify_proto_depIdxs,
		EnumInfos:         file_ProudSkillChangeNotify_proto_enumTypes,
		MessageInfos:      file_ProudSkillChangeNotify_proto_msgTypes,
	}.Build()
	File_ProudSkillChangeNotify_proto = out.File
	file_ProudSkillChangeNotify_proto_rawDesc = nil
	file_ProudSkillChangeNotify_proto_goTypes = nil
	file_ProudSkillChangeNotify_proto_depIdxs = nil
}
