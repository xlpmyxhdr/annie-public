// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PrivateChatNotify.proto

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

type PrivateChatNotify_CmdId int32

const (
	PrivateChatNotify_NONE             PrivateChatNotify_CmdId = 0
	PrivateChatNotify_ENET_CHANNEL_ID  PrivateChatNotify_CmdId = 0
	PrivateChatNotify_ENET_IS_RELIABLE PrivateChatNotify_CmdId = 1
	PrivateChatNotify_CMD_ID           PrivateChatNotify_CmdId = 4991
)

// Enum value maps for PrivateChatNotify_CmdId.
var (
	PrivateChatNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		4991: "CMD_ID",
	}
	PrivateChatNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           4991,
	}
)

func (x PrivateChatNotify_CmdId) Enum() *PrivateChatNotify_CmdId {
	p := new(PrivateChatNotify_CmdId)
	*p = x
	return p
}

func (x PrivateChatNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivateChatNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PrivateChatNotify_proto_enumTypes[0].Descriptor()
}

func (PrivateChatNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PrivateChatNotify_proto_enumTypes[0]
}

func (x PrivateChatNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivateChatNotify_CmdId.Descriptor instead.
func (PrivateChatNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PrivateChatNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PrivateChatNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo *ChatInfo `protobuf:"bytes,1,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
}

func (x *PrivateChatNotify) Reset() {
	*x = PrivateChatNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PrivateChatNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateChatNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateChatNotify) ProtoMessage() {}

func (x *PrivateChatNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PrivateChatNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateChatNotify.ProtoReflect.Descriptor instead.
func (*PrivateChatNotify) Descriptor() ([]byte, []int) {
	return file_PrivateChatNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateChatNotify) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

var File_PrivateChatNotify_proto protoreflect.FileDescriptor

var file_PrivateChatNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x90, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xff, 0x26, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PrivateChatNotify_proto_rawDescOnce sync.Once
	file_PrivateChatNotify_proto_rawDescData = file_PrivateChatNotify_proto_rawDesc
)

func file_PrivateChatNotify_proto_rawDescGZIP() []byte {
	file_PrivateChatNotify_proto_rawDescOnce.Do(func() {
		file_PrivateChatNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PrivateChatNotify_proto_rawDescData)
	})
	return file_PrivateChatNotify_proto_rawDescData
}

var file_PrivateChatNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PrivateChatNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PrivateChatNotify_proto_goTypes = []interface{}{
	(PrivateChatNotify_CmdId)(0), // 0: proto.PrivateChatNotify.CmdId
	(*PrivateChatNotify)(nil),    // 1: proto.PrivateChatNotify
	(*ChatInfo)(nil),             // 2: proto.ChatInfo
}
var file_PrivateChatNotify_proto_depIdxs = []int32{
	2, // 0: proto.PrivateChatNotify.chat_info:type_name -> proto.ChatInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PrivateChatNotify_proto_init() }
func file_PrivateChatNotify_proto_init() {
	if File_PrivateChatNotify_proto != nil {
		return
	}
	file_ChatInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PrivateChatNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateChatNotify); i {
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
			RawDescriptor: file_PrivateChatNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PrivateChatNotify_proto_goTypes,
		DependencyIndexes: file_PrivateChatNotify_proto_depIdxs,
		EnumInfos:         file_PrivateChatNotify_proto_enumTypes,
		MessageInfos:      file_PrivateChatNotify_proto_msgTypes,
	}.Build()
	File_PrivateChatNotify_proto = out.File
	file_PrivateChatNotify_proto_rawDesc = nil
	file_PrivateChatNotify_proto_goTypes = nil
	file_PrivateChatNotify_proto_depIdxs = nil
}
