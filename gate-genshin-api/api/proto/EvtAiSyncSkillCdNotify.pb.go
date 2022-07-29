// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: EvtAiSyncSkillCdNotify.proto

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

type EvtAiSyncSkillCdNotify_CmdId int32

const (
	EvtAiSyncSkillCdNotify_ENET_CHANNEL_ID  EvtAiSyncSkillCdNotify_CmdId = 0
	EvtAiSyncSkillCdNotify_NONE             EvtAiSyncSkillCdNotify_CmdId = 0
	EvtAiSyncSkillCdNotify_ENET_IS_RELIABLE EvtAiSyncSkillCdNotify_CmdId = 1
	EvtAiSyncSkillCdNotify_IS_ALLOW_CLIENT  EvtAiSyncSkillCdNotify_CmdId = 1
	EvtAiSyncSkillCdNotify_CMD_ID           EvtAiSyncSkillCdNotify_CmdId = 399
)

// Enum value maps for EvtAiSyncSkillCdNotify_CmdId.
var (
	EvtAiSyncSkillCdNotify_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		399: "CMD_ID",
	}
	EvtAiSyncSkillCdNotify_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           399,
	}
)

func (x EvtAiSyncSkillCdNotify_CmdId) Enum() *EvtAiSyncSkillCdNotify_CmdId {
	p := new(EvtAiSyncSkillCdNotify_CmdId)
	*p = x
	return p
}

func (x EvtAiSyncSkillCdNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvtAiSyncSkillCdNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EvtAiSyncSkillCdNotify_proto_enumTypes[0].Descriptor()
}

func (EvtAiSyncSkillCdNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EvtAiSyncSkillCdNotify_proto_enumTypes[0]
}

func (x EvtAiSyncSkillCdNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvtAiSyncSkillCdNotify_CmdId.Descriptor instead.
func (EvtAiSyncSkillCdNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EvtAiSyncSkillCdNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EvtAiSyncSkillCdNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AiCdMap map[uint32]*AiSkillCdInfo `protobuf:"bytes,1,rep,name=ai_cd_map,json=aiCdMap,proto3" json:"ai_cd_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EvtAiSyncSkillCdNotify) Reset() {
	*x = EvtAiSyncSkillCdNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtAiSyncSkillCdNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtAiSyncSkillCdNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtAiSyncSkillCdNotify) ProtoMessage() {}

func (x *EvtAiSyncSkillCdNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtAiSyncSkillCdNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtAiSyncSkillCdNotify.ProtoReflect.Descriptor instead.
func (*EvtAiSyncSkillCdNotify) Descriptor() ([]byte, []int) {
	return file_EvtAiSyncSkillCdNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtAiSyncSkillCdNotify) GetAiCdMap() map[uint32]*AiSkillCdInfo {
	if x != nil {
		return x.AiCdMap
	}
	return nil
}

var File_EvtAiSyncSkillCdNotify_proto protoreflect.FileDescriptor

var file_EvtAiSyncSkillCdNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x45, 0x76, 0x74, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x43, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x41, 0x69, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x16, 0x45,
	0x76, 0x74, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x46, 0x0a, 0x09, 0x61, 0x69, 0x5f, 0x63, 0x64, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x76, 0x74, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x41, 0x69, 0x43, 0x64, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x69, 0x43, 0x64, 0x4d, 0x61, 0x70, 0x1a, 0x50, 0x0a,
	0x0c, 0x41, 0x69, 0x43, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x69, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x8f, 0x03, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtAiSyncSkillCdNotify_proto_rawDescOnce sync.Once
	file_EvtAiSyncSkillCdNotify_proto_rawDescData = file_EvtAiSyncSkillCdNotify_proto_rawDesc
)

func file_EvtAiSyncSkillCdNotify_proto_rawDescGZIP() []byte {
	file_EvtAiSyncSkillCdNotify_proto_rawDescOnce.Do(func() {
		file_EvtAiSyncSkillCdNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtAiSyncSkillCdNotify_proto_rawDescData)
	})
	return file_EvtAiSyncSkillCdNotify_proto_rawDescData
}

var file_EvtAiSyncSkillCdNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EvtAiSyncSkillCdNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_EvtAiSyncSkillCdNotify_proto_goTypes = []interface{}{
	(EvtAiSyncSkillCdNotify_CmdId)(0), // 0: proto.EvtAiSyncSkillCdNotify.CmdId
	(*EvtAiSyncSkillCdNotify)(nil),    // 1: proto.EvtAiSyncSkillCdNotify
	nil,                               // 2: proto.EvtAiSyncSkillCdNotify.AiCdMapEntry
	(*AiSkillCdInfo)(nil),             // 3: proto.AiSkillCdInfo
}
var file_EvtAiSyncSkillCdNotify_proto_depIdxs = []int32{
	2, // 0: proto.EvtAiSyncSkillCdNotify.ai_cd_map:type_name -> proto.EvtAiSyncSkillCdNotify.AiCdMapEntry
	3, // 1: proto.EvtAiSyncSkillCdNotify.AiCdMapEntry.value:type_name -> proto.AiSkillCdInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtAiSyncSkillCdNotify_proto_init() }
func file_EvtAiSyncSkillCdNotify_proto_init() {
	if File_EvtAiSyncSkillCdNotify_proto != nil {
		return
	}
	file_AiSkillCdInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtAiSyncSkillCdNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtAiSyncSkillCdNotify); i {
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
			RawDescriptor: file_EvtAiSyncSkillCdNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtAiSyncSkillCdNotify_proto_goTypes,
		DependencyIndexes: file_EvtAiSyncSkillCdNotify_proto_depIdxs,
		EnumInfos:         file_EvtAiSyncSkillCdNotify_proto_enumTypes,
		MessageInfos:      file_EvtAiSyncSkillCdNotify_proto_msgTypes,
	}.Build()
	File_EvtAiSyncSkillCdNotify_proto = out.File
	file_EvtAiSyncSkillCdNotify_proto_rawDesc = nil
	file_EvtAiSyncSkillCdNotify_proto_goTypes = nil
	file_EvtAiSyncSkillCdNotify_proto_depIdxs = nil
}
