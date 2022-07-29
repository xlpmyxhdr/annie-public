// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: SceneTeamUpdateNotify.proto

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

type SceneTeamUpdateNotify_CmdId int32

const (
	SceneTeamUpdateNotify_NONE             SceneTeamUpdateNotify_CmdId = 0
	SceneTeamUpdateNotify_ENET_CHANNEL_ID  SceneTeamUpdateNotify_CmdId = 0
	SceneTeamUpdateNotify_ENET_IS_RELIABLE SceneTeamUpdateNotify_CmdId = 1
	SceneTeamUpdateNotify_CMD_ID           SceneTeamUpdateNotify_CmdId = 1793
)

// Enum value maps for SceneTeamUpdateNotify_CmdId.
var (
	SceneTeamUpdateNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1793: "CMD_ID",
	}
	SceneTeamUpdateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1793,
	}
)

func (x SceneTeamUpdateNotify_CmdId) Enum() *SceneTeamUpdateNotify_CmdId {
	p := new(SceneTeamUpdateNotify_CmdId)
	*p = x
	return p
}

func (x SceneTeamUpdateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneTeamUpdateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneTeamUpdateNotify_proto_enumTypes[0].Descriptor()
}

func (SceneTeamUpdateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_SceneTeamUpdateNotify_proto_enumTypes[0]
}

func (x SceneTeamUpdateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneTeamUpdateNotify_CmdId.Descriptor instead.
func (SceneTeamUpdateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneTeamUpdateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type SceneTeamUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneTeamAvatarList []*SceneTeamAvatar `protobuf:"bytes,1,rep,name=scene_team_avatar_list,json=sceneTeamAvatarList,proto3" json:"scene_team_avatar_list,omitempty"`
	IsInMp              bool               `protobuf:"varint,3,opt,name=is_in_mp,json=isInMp,proto3" json:"is_in_mp,omitempty"`
}

func (x *SceneTeamUpdateNotify) Reset() {
	*x = SceneTeamUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneTeamUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneTeamUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneTeamUpdateNotify) ProtoMessage() {}

func (x *SceneTeamUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneTeamUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneTeamUpdateNotify.ProtoReflect.Descriptor instead.
func (*SceneTeamUpdateNotify) Descriptor() ([]byte, []int) {
	return file_SceneTeamUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneTeamUpdateNotify) GetSceneTeamAvatarList() []*SceneTeamAvatar {
	if x != nil {
		return x.SceneTeamAvatarList
	}
	return nil
}

func (x *SceneTeamUpdateNotify) GetIsInMp() bool {
	if x != nil {
		return x.IsInMp
	}
	return false
}

var File_SceneTeamUpdateNotify_proto protoreflect.FileDescriptor

var file_SceneTeamUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x15,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x4b, 0x0a, 0x16, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x13, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x4d, 0x70, 0x22, 0x4d, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0x81, 0x0e, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneTeamUpdateNotify_proto_rawDescOnce sync.Once
	file_SceneTeamUpdateNotify_proto_rawDescData = file_SceneTeamUpdateNotify_proto_rawDesc
)

func file_SceneTeamUpdateNotify_proto_rawDescGZIP() []byte {
	file_SceneTeamUpdateNotify_proto_rawDescOnce.Do(func() {
		file_SceneTeamUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneTeamUpdateNotify_proto_rawDescData)
	})
	return file_SceneTeamUpdateNotify_proto_rawDescData
}

var file_SceneTeamUpdateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneTeamUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneTeamUpdateNotify_proto_goTypes = []interface{}{
	(SceneTeamUpdateNotify_CmdId)(0), // 0: proto.SceneTeamUpdateNotify.CmdId
	(*SceneTeamUpdateNotify)(nil),    // 1: proto.SceneTeamUpdateNotify
	(*SceneTeamAvatar)(nil),          // 2: proto.SceneTeamAvatar
}
var file_SceneTeamUpdateNotify_proto_depIdxs = []int32{
	2, // 0: proto.SceneTeamUpdateNotify.scene_team_avatar_list:type_name -> proto.SceneTeamAvatar
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneTeamUpdateNotify_proto_init() }
func file_SceneTeamUpdateNotify_proto_init() {
	if File_SceneTeamUpdateNotify_proto != nil {
		return
	}
	file_SceneTeamAvatar_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneTeamUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneTeamUpdateNotify); i {
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
			RawDescriptor: file_SceneTeamUpdateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneTeamUpdateNotify_proto_goTypes,
		DependencyIndexes: file_SceneTeamUpdateNotify_proto_depIdxs,
		EnumInfos:         file_SceneTeamUpdateNotify_proto_enumTypes,
		MessageInfos:      file_SceneTeamUpdateNotify_proto_msgTypes,
	}.Build()
	File_SceneTeamUpdateNotify_proto = out.File
	file_SceneTeamUpdateNotify_proto_rawDesc = nil
	file_SceneTeamUpdateNotify_proto_goTypes = nil
	file_SceneTeamUpdateNotify_proto_depIdxs = nil
}
