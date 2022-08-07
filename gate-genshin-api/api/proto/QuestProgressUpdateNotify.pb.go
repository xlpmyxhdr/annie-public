// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: QuestProgressUpdateNotify.proto

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

type QuestProgressUpdateNotify_CmdId int32

const (
	QuestProgressUpdateNotify_NONE             QuestProgressUpdateNotify_CmdId = 0
	QuestProgressUpdateNotify_ENET_CHANNEL_ID  QuestProgressUpdateNotify_CmdId = 0
	QuestProgressUpdateNotify_ENET_IS_RELIABLE QuestProgressUpdateNotify_CmdId = 1
	QuestProgressUpdateNotify_CMD_ID           QuestProgressUpdateNotify_CmdId = 445
)

// Enum value maps for QuestProgressUpdateNotify_CmdId.
var (
	QuestProgressUpdateNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		445: "CMD_ID",
	}
	QuestProgressUpdateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           445,
	}
)

func (x QuestProgressUpdateNotify_CmdId) Enum() *QuestProgressUpdateNotify_CmdId {
	p := new(QuestProgressUpdateNotify_CmdId)
	*p = x
	return p
}

func (x QuestProgressUpdateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestProgressUpdateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_QuestProgressUpdateNotify_proto_enumTypes[0].Descriptor()
}

func (QuestProgressUpdateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_QuestProgressUpdateNotify_proto_enumTypes[0]
}

func (x QuestProgressUpdateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestProgressUpdateNotify_CmdId.Descriptor instead.
func (QuestProgressUpdateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_QuestProgressUpdateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type QuestProgressUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestId            uint32   `protobuf:"varint,14,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	FinishProgressList []uint32 `protobuf:"varint,7,rep,packed,name=finish_progress_list,json=finishProgressList,proto3" json:"finish_progress_list,omitempty"`
	FailProgressList   []uint32 `protobuf:"varint,12,rep,packed,name=fail_progress_list,json=failProgressList,proto3" json:"fail_progress_list,omitempty"`
}

func (x *QuestProgressUpdateNotify) Reset() {
	*x = QuestProgressUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestProgressUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestProgressUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestProgressUpdateNotify) ProtoMessage() {}

func (x *QuestProgressUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_QuestProgressUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestProgressUpdateNotify.ProtoReflect.Descriptor instead.
func (*QuestProgressUpdateNotify) Descriptor() ([]byte, []int) {
	return file_QuestProgressUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *QuestProgressUpdateNotify) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *QuestProgressUpdateNotify) GetFinishProgressList() []uint32 {
	if x != nil {
		return x.FinishProgressList
	}
	return nil
}

func (x *QuestProgressUpdateNotify) GetFailProgressList() []uint32 {
	if x != nil {
		return x.FailProgressList
	}
	return nil
}

var File_QuestProgressUpdateNotify_proto protoreflect.FileDescriptor

var file_QuestProgressUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x19, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x12, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x10, 0x66, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xbd, 0x03, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QuestProgressUpdateNotify_proto_rawDescOnce sync.Once
	file_QuestProgressUpdateNotify_proto_rawDescData = file_QuestProgressUpdateNotify_proto_rawDesc
)

func file_QuestProgressUpdateNotify_proto_rawDescGZIP() []byte {
	file_QuestProgressUpdateNotify_proto_rawDescOnce.Do(func() {
		file_QuestProgressUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestProgressUpdateNotify_proto_rawDescData)
	})
	return file_QuestProgressUpdateNotify_proto_rawDescData
}

var file_QuestProgressUpdateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_QuestProgressUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestProgressUpdateNotify_proto_goTypes = []interface{}{
	(QuestProgressUpdateNotify_CmdId)(0), // 0: proto.QuestProgressUpdateNotify.CmdId
	(*QuestProgressUpdateNotify)(nil),    // 1: proto.QuestProgressUpdateNotify
}
var file_QuestProgressUpdateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QuestProgressUpdateNotify_proto_init() }
func file_QuestProgressUpdateNotify_proto_init() {
	if File_QuestProgressUpdateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QuestProgressUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestProgressUpdateNotify); i {
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
			RawDescriptor: file_QuestProgressUpdateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestProgressUpdateNotify_proto_goTypes,
		DependencyIndexes: file_QuestProgressUpdateNotify_proto_depIdxs,
		EnumInfos:         file_QuestProgressUpdateNotify_proto_enumTypes,
		MessageInfos:      file_QuestProgressUpdateNotify_proto_msgTypes,
	}.Build()
	File_QuestProgressUpdateNotify_proto = out.File
	file_QuestProgressUpdateNotify_proto_rawDesc = nil
	file_QuestProgressUpdateNotify_proto_goTypes = nil
	file_QuestProgressUpdateNotify_proto_depIdxs = nil
}
