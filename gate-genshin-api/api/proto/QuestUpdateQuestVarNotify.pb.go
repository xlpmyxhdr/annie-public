// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: QuestUpdateQuestVarNotify.proto

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

type QuestUpdateQuestVarNotify_CmdId int32

const (
	QuestUpdateQuestVarNotify_NONE             QuestUpdateQuestVarNotify_CmdId = 0
	QuestUpdateQuestVarNotify_ENET_CHANNEL_ID  QuestUpdateQuestVarNotify_CmdId = 0
	QuestUpdateQuestVarNotify_ENET_IS_RELIABLE QuestUpdateQuestVarNotify_CmdId = 1
	QuestUpdateQuestVarNotify_CMD_ID           QuestUpdateQuestVarNotify_CmdId = 463
)

// Enum value maps for QuestUpdateQuestVarNotify_CmdId.
var (
	QuestUpdateQuestVarNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		463: "CMD_ID",
	}
	QuestUpdateQuestVarNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           463,
	}
)

func (x QuestUpdateQuestVarNotify_CmdId) Enum() *QuestUpdateQuestVarNotify_CmdId {
	p := new(QuestUpdateQuestVarNotify_CmdId)
	*p = x
	return p
}

func (x QuestUpdateQuestVarNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestUpdateQuestVarNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_QuestUpdateQuestVarNotify_proto_enumTypes[0].Descriptor()
}

func (QuestUpdateQuestVarNotify_CmdId) Type() protoreflect.EnumType {
	return &file_QuestUpdateQuestVarNotify_proto_enumTypes[0]
}

func (x QuestUpdateQuestVarNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestUpdateQuestVarNotify_CmdId.Descriptor instead.
func (QuestUpdateQuestVarNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_QuestUpdateQuestVarNotify_proto_rawDescGZIP(), []int{0, 0}
}

type QuestUpdateQuestVarNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentQuestId     uint32  `protobuf:"varint,3,opt,name=parent_quest_id,json=parentQuestId,proto3" json:"parent_quest_id,omitempty"`
	QuestVar          []int32 `protobuf:"varint,6,rep,packed,name=quest_var,json=questVar,proto3" json:"quest_var,omitempty"`
	ParentQuestVarSeq uint32  `protobuf:"varint,15,opt,name=parent_quest_var_seq,json=parentQuestVarSeq,proto3" json:"parent_quest_var_seq,omitempty"`
}

func (x *QuestUpdateQuestVarNotify) Reset() {
	*x = QuestUpdateQuestVarNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestUpdateQuestVarNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestUpdateQuestVarNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestUpdateQuestVarNotify) ProtoMessage() {}

func (x *QuestUpdateQuestVarNotify) ProtoReflect() protoreflect.Message {
	mi := &file_QuestUpdateQuestVarNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestUpdateQuestVarNotify.ProtoReflect.Descriptor instead.
func (*QuestUpdateQuestVarNotify) Descriptor() ([]byte, []int) {
	return file_QuestUpdateQuestVarNotify_proto_rawDescGZIP(), []int{0}
}

func (x *QuestUpdateQuestVarNotify) GetParentQuestId() uint32 {
	if x != nil {
		return x.ParentQuestId
	}
	return 0
}

func (x *QuestUpdateQuestVarNotify) GetQuestVar() []int32 {
	if x != nil {
		return x.QuestVar
	}
	return nil
}

func (x *QuestUpdateQuestVarNotify) GetParentQuestVarSeq() uint32 {
	if x != nil {
		return x.ParentQuestVarSeq
	}
	return 0
}

var File_QuestUpdateQuestVarNotify_proto protoreflect.FileDescriptor

var file_QuestUpdateQuestVarNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x19, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72, 0x12, 0x2f, 0x0a, 0x14, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x5f,
	0x73, 0x65, 0x71, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72, 0x53, 0x65, 0x71, 0x22, 0x4d, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0xcf, 0x03, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QuestUpdateQuestVarNotify_proto_rawDescOnce sync.Once
	file_QuestUpdateQuestVarNotify_proto_rawDescData = file_QuestUpdateQuestVarNotify_proto_rawDesc
)

func file_QuestUpdateQuestVarNotify_proto_rawDescGZIP() []byte {
	file_QuestUpdateQuestVarNotify_proto_rawDescOnce.Do(func() {
		file_QuestUpdateQuestVarNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestUpdateQuestVarNotify_proto_rawDescData)
	})
	return file_QuestUpdateQuestVarNotify_proto_rawDescData
}

var file_QuestUpdateQuestVarNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_QuestUpdateQuestVarNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestUpdateQuestVarNotify_proto_goTypes = []interface{}{
	(QuestUpdateQuestVarNotify_CmdId)(0), // 0: proto.QuestUpdateQuestVarNotify.CmdId
	(*QuestUpdateQuestVarNotify)(nil),    // 1: proto.QuestUpdateQuestVarNotify
}
var file_QuestUpdateQuestVarNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QuestUpdateQuestVarNotify_proto_init() }
func file_QuestUpdateQuestVarNotify_proto_init() {
	if File_QuestUpdateQuestVarNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QuestUpdateQuestVarNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestUpdateQuestVarNotify); i {
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
			RawDescriptor: file_QuestUpdateQuestVarNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestUpdateQuestVarNotify_proto_goTypes,
		DependencyIndexes: file_QuestUpdateQuestVarNotify_proto_depIdxs,
		EnumInfos:         file_QuestUpdateQuestVarNotify_proto_enumTypes,
		MessageInfos:      file_QuestUpdateQuestVarNotify_proto_msgTypes,
	}.Build()
	File_QuestUpdateQuestVarNotify_proto = out.File
	file_QuestUpdateQuestVarNotify_proto_rawDesc = nil
	file_QuestUpdateQuestVarNotify_proto_goTypes = nil
	file_QuestUpdateQuestVarNotify_proto_depIdxs = nil
}
