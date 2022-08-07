// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: WorldPlayerReviveReq.proto

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

type WorldPlayerReviveReq_CmdId int32

const (
	WorldPlayerReviveReq_NONE             WorldPlayerReviveReq_CmdId = 0
	WorldPlayerReviveReq_ENET_CHANNEL_ID  WorldPlayerReviveReq_CmdId = 0
	WorldPlayerReviveReq_ENET_IS_RELIABLE WorldPlayerReviveReq_CmdId = 1
	WorldPlayerReviveReq_IS_ALLOW_CLIENT  WorldPlayerReviveReq_CmdId = 1
	WorldPlayerReviveReq_CMD_ID           WorldPlayerReviveReq_CmdId = 247
)

// Enum value maps for WorldPlayerReviveReq_CmdId.
var (
	WorldPlayerReviveReq_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		247: "CMD_ID",
	}
	WorldPlayerReviveReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           247,
	}
)

func (x WorldPlayerReviveReq_CmdId) Enum() *WorldPlayerReviveReq_CmdId {
	p := new(WorldPlayerReviveReq_CmdId)
	*p = x
	return p
}

func (x WorldPlayerReviveReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorldPlayerReviveReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WorldPlayerReviveReq_proto_enumTypes[0].Descriptor()
}

func (WorldPlayerReviveReq_CmdId) Type() protoreflect.EnumType {
	return &file_WorldPlayerReviveReq_proto_enumTypes[0]
}

func (x WorldPlayerReviveReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorldPlayerReviveReq_CmdId.Descriptor instead.
func (WorldPlayerReviveReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WorldPlayerReviveReq_proto_rawDescGZIP(), []int{0, 0}
}

type WorldPlayerReviveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorldPlayerReviveReq) Reset() {
	*x = WorldPlayerReviveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldPlayerReviveReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldPlayerReviveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldPlayerReviveReq) ProtoMessage() {}

func (x *WorldPlayerReviveReq) ProtoReflect() protoreflect.Message {
	mi := &file_WorldPlayerReviveReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldPlayerReviveReq.ProtoReflect.Descriptor instead.
func (*WorldPlayerReviveReq) Descriptor() ([]byte, []int) {
	return file_WorldPlayerReviveReq_proto_rawDescGZIP(), []int{0}
}

var File_WorldPlayerReviveReq_proto protoreflect.FileDescriptor

var file_WorldPlayerReviveReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x76,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x22, 0x62, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf7, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_WorldPlayerReviveReq_proto_rawDescOnce sync.Once
	file_WorldPlayerReviveReq_proto_rawDescData = file_WorldPlayerReviveReq_proto_rawDesc
)

func file_WorldPlayerReviveReq_proto_rawDescGZIP() []byte {
	file_WorldPlayerReviveReq_proto_rawDescOnce.Do(func() {
		file_WorldPlayerReviveReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldPlayerReviveReq_proto_rawDescData)
	})
	return file_WorldPlayerReviveReq_proto_rawDescData
}

var file_WorldPlayerReviveReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WorldPlayerReviveReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorldPlayerReviveReq_proto_goTypes = []interface{}{
	(WorldPlayerReviveReq_CmdId)(0), // 0: proto.WorldPlayerReviveReq.CmdId
	(*WorldPlayerReviveReq)(nil),    // 1: proto.WorldPlayerReviveReq
}
var file_WorldPlayerReviveReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WorldPlayerReviveReq_proto_init() }
func file_WorldPlayerReviveReq_proto_init() {
	if File_WorldPlayerReviveReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WorldPlayerReviveReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldPlayerReviveReq); i {
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
			RawDescriptor: file_WorldPlayerReviveReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldPlayerReviveReq_proto_goTypes,
		DependencyIndexes: file_WorldPlayerReviveReq_proto_depIdxs,
		EnumInfos:         file_WorldPlayerReviveReq_proto_enumTypes,
		MessageInfos:      file_WorldPlayerReviveReq_proto_msgTypes,
	}.Build()
	File_WorldPlayerReviveReq_proto = out.File
	file_WorldPlayerReviveReq_proto_rawDesc = nil
	file_WorldPlayerReviveReq_proto_goTypes = nil
	file_WorldPlayerReviveReq_proto_depIdxs = nil
}
