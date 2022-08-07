// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: SyncTeamEntityNotify.proto

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

type SyncTeamEntityNotify_CmdId int32

const (
	SyncTeamEntityNotify_NONE             SyncTeamEntityNotify_CmdId = 0
	SyncTeamEntityNotify_ENET_CHANNEL_ID  SyncTeamEntityNotify_CmdId = 0
	SyncTeamEntityNotify_ENET_IS_RELIABLE SyncTeamEntityNotify_CmdId = 1
	SyncTeamEntityNotify_CMD_ID           SyncTeamEntityNotify_CmdId = 334
)

// Enum value maps for SyncTeamEntityNotify_CmdId.
var (
	SyncTeamEntityNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		334: "CMD_ID",
	}
	SyncTeamEntityNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           334,
	}
)

func (x SyncTeamEntityNotify_CmdId) Enum() *SyncTeamEntityNotify_CmdId {
	p := new(SyncTeamEntityNotify_CmdId)
	*p = x
	return p
}

func (x SyncTeamEntityNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncTeamEntityNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SyncTeamEntityNotify_proto_enumTypes[0].Descriptor()
}

func (SyncTeamEntityNotify_CmdId) Type() protoreflect.EnumType {
	return &file_SyncTeamEntityNotify_proto_enumTypes[0]
}

func (x SyncTeamEntityNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SyncTeamEntityNotify_CmdId.Descriptor instead.
func (SyncTeamEntityNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SyncTeamEntityNotify_proto_rawDescGZIP(), []int{0, 0}
}

type SyncTeamEntityNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId            uint32            `protobuf:"varint,13,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	TeamEntityInfoList []*TeamEntityInfo `protobuf:"bytes,2,rep,name=team_entity_info_list,json=teamEntityInfoList,proto3" json:"team_entity_info_list,omitempty"`
}

func (x *SyncTeamEntityNotify) Reset() {
	*x = SyncTeamEntityNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncTeamEntityNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTeamEntityNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTeamEntityNotify) ProtoMessage() {}

func (x *SyncTeamEntityNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncTeamEntityNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTeamEntityNotify.ProtoReflect.Descriptor instead.
func (*SyncTeamEntityNotify) Descriptor() ([]byte, []int) {
	return file_SyncTeamEntityNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncTeamEntityNotify) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *SyncTeamEntityNotify) GetTeamEntityInfoList() []*TeamEntityInfo {
	if x != nil {
		return x.TeamEntityInfoList
	}
	return nil
}

var File_SyncTeamEntityNotify_proto protoreflect.FileDescriptor

var file_SyncTeamEntityNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x53, 0x79,
	0x6e, 0x63, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x48, 0x0a,
	0x15, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x12, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xce, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncTeamEntityNotify_proto_rawDescOnce sync.Once
	file_SyncTeamEntityNotify_proto_rawDescData = file_SyncTeamEntityNotify_proto_rawDesc
)

func file_SyncTeamEntityNotify_proto_rawDescGZIP() []byte {
	file_SyncTeamEntityNotify_proto_rawDescOnce.Do(func() {
		file_SyncTeamEntityNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncTeamEntityNotify_proto_rawDescData)
	})
	return file_SyncTeamEntityNotify_proto_rawDescData
}

var file_SyncTeamEntityNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SyncTeamEntityNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncTeamEntityNotify_proto_goTypes = []interface{}{
	(SyncTeamEntityNotify_CmdId)(0), // 0: proto.SyncTeamEntityNotify.CmdId
	(*SyncTeamEntityNotify)(nil),    // 1: proto.SyncTeamEntityNotify
	(*TeamEntityInfo)(nil),          // 2: proto.TeamEntityInfo
}
var file_SyncTeamEntityNotify_proto_depIdxs = []int32{
	2, // 0: proto.SyncTeamEntityNotify.team_entity_info_list:type_name -> proto.TeamEntityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncTeamEntityNotify_proto_init() }
func file_SyncTeamEntityNotify_proto_init() {
	if File_SyncTeamEntityNotify_proto != nil {
		return
	}
	file_TeamEntityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncTeamEntityNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTeamEntityNotify); i {
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
			RawDescriptor: file_SyncTeamEntityNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncTeamEntityNotify_proto_goTypes,
		DependencyIndexes: file_SyncTeamEntityNotify_proto_depIdxs,
		EnumInfos:         file_SyncTeamEntityNotify_proto_enumTypes,
		MessageInfos:      file_SyncTeamEntityNotify_proto_msgTypes,
	}.Build()
	File_SyncTeamEntityNotify_proto = out.File
	file_SyncTeamEntityNotify_proto_rawDesc = nil
	file_SyncTeamEntityNotify_proto_goTypes = nil
	file_SyncTeamEntityNotify_proto_depIdxs = nil
}
