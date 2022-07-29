// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: EvtAiSyncCombatThreatInfoNotify.proto

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

type EvtAiSyncCombatThreatInfoNotify_CmdId int32

const (
	EvtAiSyncCombatThreatInfoNotify_ENET_CHANNEL_ID  EvtAiSyncCombatThreatInfoNotify_CmdId = 0
	EvtAiSyncCombatThreatInfoNotify_NONE             EvtAiSyncCombatThreatInfoNotify_CmdId = 0
	EvtAiSyncCombatThreatInfoNotify_ENET_IS_RELIABLE EvtAiSyncCombatThreatInfoNotify_CmdId = 1
	EvtAiSyncCombatThreatInfoNotify_IS_ALLOW_CLIENT  EvtAiSyncCombatThreatInfoNotify_CmdId = 1
	EvtAiSyncCombatThreatInfoNotify_CMD_ID           EvtAiSyncCombatThreatInfoNotify_CmdId = 373
)

// Enum value maps for EvtAiSyncCombatThreatInfoNotify_CmdId.
var (
	EvtAiSyncCombatThreatInfoNotify_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		373: "CMD_ID",
	}
	EvtAiSyncCombatThreatInfoNotify_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           373,
	}
)

func (x EvtAiSyncCombatThreatInfoNotify_CmdId) Enum() *EvtAiSyncCombatThreatInfoNotify_CmdId {
	p := new(EvtAiSyncCombatThreatInfoNotify_CmdId)
	*p = x
	return p
}

func (x EvtAiSyncCombatThreatInfoNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvtAiSyncCombatThreatInfoNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EvtAiSyncCombatThreatInfoNotify_proto_enumTypes[0].Descriptor()
}

func (EvtAiSyncCombatThreatInfoNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EvtAiSyncCombatThreatInfoNotify_proto_enumTypes[0]
}

func (x EvtAiSyncCombatThreatInfoNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvtAiSyncCombatThreatInfoNotify_CmdId.Descriptor instead.
func (EvtAiSyncCombatThreatInfoNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EvtAiSyncCombatThreatInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CombatThreatInfoMap map[uint32]*AiThreatInfo `protobuf:"bytes,1,rep,name=combat_threat_info_map,json=combatThreatInfoMap,proto3" json:"combat_threat_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EvtAiSyncCombatThreatInfoNotify) Reset() {
	*x = EvtAiSyncCombatThreatInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtAiSyncCombatThreatInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtAiSyncCombatThreatInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtAiSyncCombatThreatInfoNotify) ProtoMessage() {}

func (x *EvtAiSyncCombatThreatInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtAiSyncCombatThreatInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtAiSyncCombatThreatInfoNotify.ProtoReflect.Descriptor instead.
func (*EvtAiSyncCombatThreatInfoNotify) Descriptor() ([]byte, []int) {
	return file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtAiSyncCombatThreatInfoNotify) GetCombatThreatInfoMap() map[uint32]*AiThreatInfo {
	if x != nil {
		return x.CombatThreatInfoMap
	}
	return nil
}

var File_EvtAiSyncCombatThreatInfoNotify_proto protoreflect.FileDescriptor

var file_EvtAiSyncCombatThreatInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x45, 0x76, 0x74, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x41, 0x69, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x1f, 0x45, 0x76, 0x74, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63,
	0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x74, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x76, 0x74, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x43,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x1a, 0x5b, 0x0a, 0x18,
	0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x69, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64,
	0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06,
	0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf5, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescOnce sync.Once
	file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescData = file_EvtAiSyncCombatThreatInfoNotify_proto_rawDesc
)

func file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescGZIP() []byte {
	file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescOnce.Do(func() {
		file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescData)
	})
	return file_EvtAiSyncCombatThreatInfoNotify_proto_rawDescData
}

var file_EvtAiSyncCombatThreatInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EvtAiSyncCombatThreatInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_EvtAiSyncCombatThreatInfoNotify_proto_goTypes = []interface{}{
	(EvtAiSyncCombatThreatInfoNotify_CmdId)(0), // 0: proto.EvtAiSyncCombatThreatInfoNotify.CmdId
	(*EvtAiSyncCombatThreatInfoNotify)(nil),    // 1: proto.EvtAiSyncCombatThreatInfoNotify
	nil,                                        // 2: proto.EvtAiSyncCombatThreatInfoNotify.CombatThreatInfoMapEntry
	(*AiThreatInfo)(nil),                       // 3: proto.AiThreatInfo
}
var file_EvtAiSyncCombatThreatInfoNotify_proto_depIdxs = []int32{
	2, // 0: proto.EvtAiSyncCombatThreatInfoNotify.combat_threat_info_map:type_name -> proto.EvtAiSyncCombatThreatInfoNotify.CombatThreatInfoMapEntry
	3, // 1: proto.EvtAiSyncCombatThreatInfoNotify.CombatThreatInfoMapEntry.value:type_name -> proto.AiThreatInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtAiSyncCombatThreatInfoNotify_proto_init() }
func file_EvtAiSyncCombatThreatInfoNotify_proto_init() {
	if File_EvtAiSyncCombatThreatInfoNotify_proto != nil {
		return
	}
	file_AiThreatInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtAiSyncCombatThreatInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtAiSyncCombatThreatInfoNotify); i {
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
			RawDescriptor: file_EvtAiSyncCombatThreatInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtAiSyncCombatThreatInfoNotify_proto_goTypes,
		DependencyIndexes: file_EvtAiSyncCombatThreatInfoNotify_proto_depIdxs,
		EnumInfos:         file_EvtAiSyncCombatThreatInfoNotify_proto_enumTypes,
		MessageInfos:      file_EvtAiSyncCombatThreatInfoNotify_proto_msgTypes,
	}.Build()
	File_EvtAiSyncCombatThreatInfoNotify_proto = out.File
	file_EvtAiSyncCombatThreatInfoNotify_proto_rawDesc = nil
	file_EvtAiSyncCombatThreatInfoNotify_proto_goTypes = nil
	file_EvtAiSyncCombatThreatInfoNotify_proto_depIdxs = nil
}
