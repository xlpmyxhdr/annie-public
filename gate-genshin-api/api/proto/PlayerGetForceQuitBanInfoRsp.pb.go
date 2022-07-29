// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PlayerGetForceQuitBanInfoRsp.proto

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

type PlayerGetForceQuitBanInfoRsp_CmdId int32

const (
	PlayerGetForceQuitBanInfoRsp_NONE             PlayerGetForceQuitBanInfoRsp_CmdId = 0
	PlayerGetForceQuitBanInfoRsp_ENET_CHANNEL_ID  PlayerGetForceQuitBanInfoRsp_CmdId = 0
	PlayerGetForceQuitBanInfoRsp_ENET_IS_RELIABLE PlayerGetForceQuitBanInfoRsp_CmdId = 1
	PlayerGetForceQuitBanInfoRsp_CMD_ID           PlayerGetForceQuitBanInfoRsp_CmdId = 4188
)

// Enum value maps for PlayerGetForceQuitBanInfoRsp_CmdId.
var (
	PlayerGetForceQuitBanInfoRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		4188: "CMD_ID",
	}
	PlayerGetForceQuitBanInfoRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           4188,
	}
)

func (x PlayerGetForceQuitBanInfoRsp_CmdId) Enum() *PlayerGetForceQuitBanInfoRsp_CmdId {
	p := new(PlayerGetForceQuitBanInfoRsp_CmdId)
	*p = x
	return p
}

func (x PlayerGetForceQuitBanInfoRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerGetForceQuitBanInfoRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerGetForceQuitBanInfoRsp_proto_enumTypes[0].Descriptor()
}

func (PlayerGetForceQuitBanInfoRsp_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerGetForceQuitBanInfoRsp_proto_enumTypes[0]
}

func (x PlayerGetForceQuitBanInfoRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerGetForceQuitBanInfoRsp_CmdId.Descriptor instead.
func (PlayerGetForceQuitBanInfoRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerGetForceQuitBanInfoRsp_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerGetForceQuitBanInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MatchId    uint32 `protobuf:"varint,2,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	ExpireTime uint32 `protobuf:"varint,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *PlayerGetForceQuitBanInfoRsp) Reset() {
	*x = PlayerGetForceQuitBanInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerGetForceQuitBanInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerGetForceQuitBanInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerGetForceQuitBanInfoRsp) ProtoMessage() {}

func (x *PlayerGetForceQuitBanInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerGetForceQuitBanInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerGetForceQuitBanInfoRsp.ProtoReflect.Descriptor instead.
func (*PlayerGetForceQuitBanInfoRsp) Descriptor() ([]byte, []int) {
	return file_PlayerGetForceQuitBanInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerGetForceQuitBanInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlayerGetForceQuitBanInfoRsp) GetMatchId() uint32 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *PlayerGetForceQuitBanInfoRsp) GetExpireTime() uint32 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

var File_PlayerGetForceQuitBanInfoRsp_proto protoreflect.FileDescriptor

var file_PlayerGetForceQuitBanInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x51, 0x75, 0x69, 0x74, 0x42, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x1c,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x51, 0x75,
	0x69, 0x74, 0x42, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xdc, 0x20, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerGetForceQuitBanInfoRsp_proto_rawDescOnce sync.Once
	file_PlayerGetForceQuitBanInfoRsp_proto_rawDescData = file_PlayerGetForceQuitBanInfoRsp_proto_rawDesc
)

func file_PlayerGetForceQuitBanInfoRsp_proto_rawDescGZIP() []byte {
	file_PlayerGetForceQuitBanInfoRsp_proto_rawDescOnce.Do(func() {
		file_PlayerGetForceQuitBanInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerGetForceQuitBanInfoRsp_proto_rawDescData)
	})
	return file_PlayerGetForceQuitBanInfoRsp_proto_rawDescData
}

var file_PlayerGetForceQuitBanInfoRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerGetForceQuitBanInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerGetForceQuitBanInfoRsp_proto_goTypes = []interface{}{
	(PlayerGetForceQuitBanInfoRsp_CmdId)(0), // 0: proto.PlayerGetForceQuitBanInfoRsp.CmdId
	(*PlayerGetForceQuitBanInfoRsp)(nil),    // 1: proto.PlayerGetForceQuitBanInfoRsp
}
var file_PlayerGetForceQuitBanInfoRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerGetForceQuitBanInfoRsp_proto_init() }
func file_PlayerGetForceQuitBanInfoRsp_proto_init() {
	if File_PlayerGetForceQuitBanInfoRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerGetForceQuitBanInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerGetForceQuitBanInfoRsp); i {
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
			RawDescriptor: file_PlayerGetForceQuitBanInfoRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerGetForceQuitBanInfoRsp_proto_goTypes,
		DependencyIndexes: file_PlayerGetForceQuitBanInfoRsp_proto_depIdxs,
		EnumInfos:         file_PlayerGetForceQuitBanInfoRsp_proto_enumTypes,
		MessageInfos:      file_PlayerGetForceQuitBanInfoRsp_proto_msgTypes,
	}.Build()
	File_PlayerGetForceQuitBanInfoRsp_proto = out.File
	file_PlayerGetForceQuitBanInfoRsp_proto_rawDesc = nil
	file_PlayerGetForceQuitBanInfoRsp_proto_goTypes = nil
	file_PlayerGetForceQuitBanInfoRsp_proto_depIdxs = nil
}
