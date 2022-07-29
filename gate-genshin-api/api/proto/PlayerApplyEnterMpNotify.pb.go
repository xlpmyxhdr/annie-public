// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PlayerApplyEnterMpNotify.proto

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

type PlayerApplyEnterMpNotify_CmdId int32

const (
	PlayerApplyEnterMpNotify_NONE             PlayerApplyEnterMpNotify_CmdId = 0
	PlayerApplyEnterMpNotify_ENET_CHANNEL_ID  PlayerApplyEnterMpNotify_CmdId = 0
	PlayerApplyEnterMpNotify_ENET_IS_RELIABLE PlayerApplyEnterMpNotify_CmdId = 1
	PlayerApplyEnterMpNotify_CMD_ID           PlayerApplyEnterMpNotify_CmdId = 1803
)

// Enum value maps for PlayerApplyEnterMpNotify_CmdId.
var (
	PlayerApplyEnterMpNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1803: "CMD_ID",
	}
	PlayerApplyEnterMpNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1803,
	}
)

func (x PlayerApplyEnterMpNotify_CmdId) Enum() *PlayerApplyEnterMpNotify_CmdId {
	p := new(PlayerApplyEnterMpNotify_CmdId)
	*p = x
	return p
}

func (x PlayerApplyEnterMpNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerApplyEnterMpNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerApplyEnterMpNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerApplyEnterMpNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerApplyEnterMpNotify_proto_enumTypes[0]
}

func (x PlayerApplyEnterMpNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerApplyEnterMpNotify_CmdId.Descriptor instead.
func (PlayerApplyEnterMpNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerApplyEnterMpNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerApplyEnterMpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcPlayerInfo  *OnlinePlayerInfo `protobuf:"bytes,1,opt,name=src_player_info,json=srcPlayerInfo,proto3" json:"src_player_info,omitempty"`
	SrcAppId       uint32            `protobuf:"varint,2,opt,name=src_app_id,json=srcAppId,proto3" json:"src_app_id,omitempty"`
	SrcThreadIndex uint32            `protobuf:"varint,3,opt,name=src_thread_index,json=srcThreadIndex,proto3" json:"src_thread_index,omitempty"`
}

func (x *PlayerApplyEnterMpNotify) Reset() {
	*x = PlayerApplyEnterMpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerApplyEnterMpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerApplyEnterMpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerApplyEnterMpNotify) ProtoMessage() {}

func (x *PlayerApplyEnterMpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerApplyEnterMpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerApplyEnterMpNotify.ProtoReflect.Descriptor instead.
func (*PlayerApplyEnterMpNotify) Descriptor() ([]byte, []int) {
	return file_PlayerApplyEnterMpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerApplyEnterMpNotify) GetSrcPlayerInfo() *OnlinePlayerInfo {
	if x != nil {
		return x.SrcPlayerInfo
	}
	return nil
}

func (x *PlayerApplyEnterMpNotify) GetSrcAppId() uint32 {
	if x != nil {
		return x.SrcAppId
	}
	return 0
}

func (x *PlayerApplyEnterMpNotify) GetSrcThreadIndex() uint32 {
	if x != nil {
		return x.SrcThreadIndex
	}
	return 0
}

var File_PlayerApplyEnterMpNotify_proto protoreflect.FileDescriptor

var file_PlayerApplyEnterMpNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf2, 0x01, 0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3f, 0x0a, 0x0f,
	0x73, 0x72, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x73, 0x72, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x0a, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x72, 0x63, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x72, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x72, 0x63, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x8b, 0x0e,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerApplyEnterMpNotify_proto_rawDescOnce sync.Once
	file_PlayerApplyEnterMpNotify_proto_rawDescData = file_PlayerApplyEnterMpNotify_proto_rawDesc
)

func file_PlayerApplyEnterMpNotify_proto_rawDescGZIP() []byte {
	file_PlayerApplyEnterMpNotify_proto_rawDescOnce.Do(func() {
		file_PlayerApplyEnterMpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerApplyEnterMpNotify_proto_rawDescData)
	})
	return file_PlayerApplyEnterMpNotify_proto_rawDescData
}

var file_PlayerApplyEnterMpNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerApplyEnterMpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerApplyEnterMpNotify_proto_goTypes = []interface{}{
	(PlayerApplyEnterMpNotify_CmdId)(0), // 0: proto.PlayerApplyEnterMpNotify.CmdId
	(*PlayerApplyEnterMpNotify)(nil),    // 1: proto.PlayerApplyEnterMpNotify
	(*OnlinePlayerInfo)(nil),            // 2: proto.OnlinePlayerInfo
}
var file_PlayerApplyEnterMpNotify_proto_depIdxs = []int32{
	2, // 0: proto.PlayerApplyEnterMpNotify.src_player_info:type_name -> proto.OnlinePlayerInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerApplyEnterMpNotify_proto_init() }
func file_PlayerApplyEnterMpNotify_proto_init() {
	if File_PlayerApplyEnterMpNotify_proto != nil {
		return
	}
	file_OnlinePlayerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerApplyEnterMpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerApplyEnterMpNotify); i {
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
			RawDescriptor: file_PlayerApplyEnterMpNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerApplyEnterMpNotify_proto_goTypes,
		DependencyIndexes: file_PlayerApplyEnterMpNotify_proto_depIdxs,
		EnumInfos:         file_PlayerApplyEnterMpNotify_proto_enumTypes,
		MessageInfos:      file_PlayerApplyEnterMpNotify_proto_msgTypes,
	}.Build()
	File_PlayerApplyEnterMpNotify_proto = out.File
	file_PlayerApplyEnterMpNotify_proto_rawDesc = nil
	file_PlayerApplyEnterMpNotify_proto_goTypes = nil
	file_PlayerApplyEnterMpNotify_proto_depIdxs = nil
}
