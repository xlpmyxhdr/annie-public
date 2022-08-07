// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarSkillChangeNotify.proto

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

type AvatarSkillChangeNotify_CmdId int32

const (
	AvatarSkillChangeNotify_NONE             AvatarSkillChangeNotify_CmdId = 0
	AvatarSkillChangeNotify_ENET_CHANNEL_ID  AvatarSkillChangeNotify_CmdId = 0
	AvatarSkillChangeNotify_ENET_IS_RELIABLE AvatarSkillChangeNotify_CmdId = 1
	AvatarSkillChangeNotify_CMD_ID           AvatarSkillChangeNotify_CmdId = 1030
)

// Enum value maps for AvatarSkillChangeNotify_CmdId.
var (
	AvatarSkillChangeNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1030: "CMD_ID",
	}
	AvatarSkillChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1030,
	}
)

func (x AvatarSkillChangeNotify_CmdId) Enum() *AvatarSkillChangeNotify_CmdId {
	p := new(AvatarSkillChangeNotify_CmdId)
	*p = x
	return p
}

func (x AvatarSkillChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarSkillChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarSkillChangeNotify_proto_enumTypes[0].Descriptor()
}

func (AvatarSkillChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarSkillChangeNotify_proto_enumTypes[0]
}

func (x AvatarSkillChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarSkillChangeNotify_CmdId.Descriptor instead.
func (AvatarSkillChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarSkillChangeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarSkillChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid    uint64 `protobuf:"varint,4,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	EntityId      uint32 `protobuf:"varint,15,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	SkillDepotId  uint32 `protobuf:"varint,5,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	AvatarSkillId uint32 `protobuf:"varint,2,opt,name=avatar_skill_id,json=avatarSkillId,proto3" json:"avatar_skill_id,omitempty"`
	OldLevel      uint32 `protobuf:"varint,10,opt,name=old_level,json=oldLevel,proto3" json:"old_level,omitempty"`
	CurLevel      uint32 `protobuf:"varint,7,opt,name=cur_level,json=curLevel,proto3" json:"cur_level,omitempty"`
}

func (x *AvatarSkillChangeNotify) Reset() {
	*x = AvatarSkillChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarSkillChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarSkillChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarSkillChangeNotify) ProtoMessage() {}

func (x *AvatarSkillChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarSkillChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarSkillChangeNotify.ProtoReflect.Descriptor instead.
func (*AvatarSkillChangeNotify) Descriptor() ([]byte, []int) {
	return file_AvatarSkillChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarSkillChangeNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarSkillChangeNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AvatarSkillChangeNotify) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *AvatarSkillChangeNotify) GetAvatarSkillId() uint32 {
	if x != nil {
		return x.AvatarSkillId
	}
	return 0
}

func (x *AvatarSkillChangeNotify) GetOldLevel() uint32 {
	if x != nil {
		return x.OldLevel
	}
	return 0
}

func (x *AvatarSkillChangeNotify) GetCurLevel() uint32 {
	if x != nil {
		return x.CurLevel
	}
	return 0
}

var File_AvatarSkillChangeNotify_proto protoreflect.FileDescriptor

var file_AvatarSkillChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x17, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47,
	0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44,
	0x65, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x75, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0x86, 0x08, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarSkillChangeNotify_proto_rawDescOnce sync.Once
	file_AvatarSkillChangeNotify_proto_rawDescData = file_AvatarSkillChangeNotify_proto_rawDesc
)

func file_AvatarSkillChangeNotify_proto_rawDescGZIP() []byte {
	file_AvatarSkillChangeNotify_proto_rawDescOnce.Do(func() {
		file_AvatarSkillChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarSkillChangeNotify_proto_rawDescData)
	})
	return file_AvatarSkillChangeNotify_proto_rawDescData
}

var file_AvatarSkillChangeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarSkillChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarSkillChangeNotify_proto_goTypes = []interface{}{
	(AvatarSkillChangeNotify_CmdId)(0), // 0: proto.AvatarSkillChangeNotify.CmdId
	(*AvatarSkillChangeNotify)(nil),    // 1: proto.AvatarSkillChangeNotify
}
var file_AvatarSkillChangeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarSkillChangeNotify_proto_init() }
func file_AvatarSkillChangeNotify_proto_init() {
	if File_AvatarSkillChangeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarSkillChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarSkillChangeNotify); i {
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
			RawDescriptor: file_AvatarSkillChangeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarSkillChangeNotify_proto_goTypes,
		DependencyIndexes: file_AvatarSkillChangeNotify_proto_depIdxs,
		EnumInfos:         file_AvatarSkillChangeNotify_proto_enumTypes,
		MessageInfos:      file_AvatarSkillChangeNotify_proto_msgTypes,
	}.Build()
	File_AvatarSkillChangeNotify_proto = out.File
	file_AvatarSkillChangeNotify_proto_rawDesc = nil
	file_AvatarSkillChangeNotify_proto_goTypes = nil
	file_AvatarSkillChangeNotify_proto_depIdxs = nil
}
