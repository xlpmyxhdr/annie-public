// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PlayerLevelRewardUpdateNotify.proto

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

type PlayerLevelRewardUpdateNotify_CmdId int32

const (
	PlayerLevelRewardUpdateNotify_NONE             PlayerLevelRewardUpdateNotify_CmdId = 0
	PlayerLevelRewardUpdateNotify_ENET_CHANNEL_ID  PlayerLevelRewardUpdateNotify_CmdId = 0
	PlayerLevelRewardUpdateNotify_ENET_IS_RELIABLE PlayerLevelRewardUpdateNotify_CmdId = 1
	PlayerLevelRewardUpdateNotify_CMD_ID           PlayerLevelRewardUpdateNotify_CmdId = 152
)

// Enum value maps for PlayerLevelRewardUpdateNotify_CmdId.
var (
	PlayerLevelRewardUpdateNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		152: "CMD_ID",
	}
	PlayerLevelRewardUpdateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           152,
	}
)

func (x PlayerLevelRewardUpdateNotify_CmdId) Enum() *PlayerLevelRewardUpdateNotify_CmdId {
	p := new(PlayerLevelRewardUpdateNotify_CmdId)
	*p = x
	return p
}

func (x PlayerLevelRewardUpdateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerLevelRewardUpdateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerLevelRewardUpdateNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerLevelRewardUpdateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerLevelRewardUpdateNotify_proto_enumTypes[0]
}

func (x PlayerLevelRewardUpdateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerLevelRewardUpdateNotify_CmdId.Descriptor instead.
func (PlayerLevelRewardUpdateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerLevelRewardUpdateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerLevelRewardUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelList []uint32 `protobuf:"varint,1,rep,packed,name=level_list,json=levelList,proto3" json:"level_list,omitempty"`
}

func (x *PlayerLevelRewardUpdateNotify) Reset() {
	*x = PlayerLevelRewardUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLevelRewardUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLevelRewardUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLevelRewardUpdateNotify) ProtoMessage() {}

func (x *PlayerLevelRewardUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLevelRewardUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLevelRewardUpdateNotify.ProtoReflect.Descriptor instead.
func (*PlayerLevelRewardUpdateNotify) Descriptor() ([]byte, []int) {
	return file_PlayerLevelRewardUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLevelRewardUpdateNotify) GetLevelList() []uint32 {
	if x != nil {
		return x.LevelList
	}
	return nil
}

var File_PlayerLevelRewardUpdateNotify_proto protoreflect.FileDescriptor

var file_PlayerLevelRewardUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a,
	0x1d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x98, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerLevelRewardUpdateNotify_proto_rawDescOnce sync.Once
	file_PlayerLevelRewardUpdateNotify_proto_rawDescData = file_PlayerLevelRewardUpdateNotify_proto_rawDesc
)

func file_PlayerLevelRewardUpdateNotify_proto_rawDescGZIP() []byte {
	file_PlayerLevelRewardUpdateNotify_proto_rawDescOnce.Do(func() {
		file_PlayerLevelRewardUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLevelRewardUpdateNotify_proto_rawDescData)
	})
	return file_PlayerLevelRewardUpdateNotify_proto_rawDescData
}

var file_PlayerLevelRewardUpdateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerLevelRewardUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLevelRewardUpdateNotify_proto_goTypes = []interface{}{
	(PlayerLevelRewardUpdateNotify_CmdId)(0), // 0: proto.PlayerLevelRewardUpdateNotify.CmdId
	(*PlayerLevelRewardUpdateNotify)(nil),    // 1: proto.PlayerLevelRewardUpdateNotify
}
var file_PlayerLevelRewardUpdateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerLevelRewardUpdateNotify_proto_init() }
func file_PlayerLevelRewardUpdateNotify_proto_init() {
	if File_PlayerLevelRewardUpdateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerLevelRewardUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLevelRewardUpdateNotify); i {
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
			RawDescriptor: file_PlayerLevelRewardUpdateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLevelRewardUpdateNotify_proto_goTypes,
		DependencyIndexes: file_PlayerLevelRewardUpdateNotify_proto_depIdxs,
		EnumInfos:         file_PlayerLevelRewardUpdateNotify_proto_enumTypes,
		MessageInfos:      file_PlayerLevelRewardUpdateNotify_proto_msgTypes,
	}.Build()
	File_PlayerLevelRewardUpdateNotify_proto = out.File
	file_PlayerLevelRewardUpdateNotify_proto_rawDesc = nil
	file_PlayerLevelRewardUpdateNotify_proto_goTypes = nil
	file_PlayerLevelRewardUpdateNotify_proto_depIdxs = nil
}
