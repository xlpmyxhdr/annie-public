// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: SetPlayerPropRsp.proto

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

type SetPlayerPropRsp_CmdId int32

const (
	SetPlayerPropRsp_NONE             SetPlayerPropRsp_CmdId = 0
	SetPlayerPropRsp_ENET_CHANNEL_ID  SetPlayerPropRsp_CmdId = 0
	SetPlayerPropRsp_ENET_IS_RELIABLE SetPlayerPropRsp_CmdId = 1
	SetPlayerPropRsp_CMD_ID           SetPlayerPropRsp_CmdId = 197
)

// Enum value maps for SetPlayerPropRsp_CmdId.
var (
	SetPlayerPropRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		197: "CMD_ID",
	}
	SetPlayerPropRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           197,
	}
)

func (x SetPlayerPropRsp_CmdId) Enum() *SetPlayerPropRsp_CmdId {
	p := new(SetPlayerPropRsp_CmdId)
	*p = x
	return p
}

func (x SetPlayerPropRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetPlayerPropRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SetPlayerPropRsp_proto_enumTypes[0].Descriptor()
}

func (SetPlayerPropRsp_CmdId) Type() protoreflect.EnumType {
	return &file_SetPlayerPropRsp_proto_enumTypes[0]
}

func (x SetPlayerPropRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetPlayerPropRsp_CmdId.Descriptor instead.
func (SetPlayerPropRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SetPlayerPropRsp_proto_rawDescGZIP(), []int{0, 0}
}

type SetPlayerPropRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetPlayerPropRsp) Reset() {
	*x = SetPlayerPropRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetPlayerPropRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerPropRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerPropRsp) ProtoMessage() {}

func (x *SetPlayerPropRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetPlayerPropRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerPropRsp.ProtoReflect.Descriptor instead.
func (*SetPlayerPropRsp) Descriptor() ([]byte, []int) {
	return file_SetPlayerPropRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetPlayerPropRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetPlayerPropRsp_proto protoreflect.FileDescriptor

var file_SetPlayerPropRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7b, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xc5, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetPlayerPropRsp_proto_rawDescOnce sync.Once
	file_SetPlayerPropRsp_proto_rawDescData = file_SetPlayerPropRsp_proto_rawDesc
)

func file_SetPlayerPropRsp_proto_rawDescGZIP() []byte {
	file_SetPlayerPropRsp_proto_rawDescOnce.Do(func() {
		file_SetPlayerPropRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetPlayerPropRsp_proto_rawDescData)
	})
	return file_SetPlayerPropRsp_proto_rawDescData
}

var file_SetPlayerPropRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SetPlayerPropRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetPlayerPropRsp_proto_goTypes = []interface{}{
	(SetPlayerPropRsp_CmdId)(0), // 0: proto.SetPlayerPropRsp.CmdId
	(*SetPlayerPropRsp)(nil),    // 1: proto.SetPlayerPropRsp
}
var file_SetPlayerPropRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetPlayerPropRsp_proto_init() }
func file_SetPlayerPropRsp_proto_init() {
	if File_SetPlayerPropRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetPlayerPropRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerPropRsp); i {
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
			RawDescriptor: file_SetPlayerPropRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetPlayerPropRsp_proto_goTypes,
		DependencyIndexes: file_SetPlayerPropRsp_proto_depIdxs,
		EnumInfos:         file_SetPlayerPropRsp_proto_enumTypes,
		MessageInfos:      file_SetPlayerPropRsp_proto_msgTypes,
	}.Build()
	File_SetPlayerPropRsp_proto = out.File
	file_SetPlayerPropRsp_proto_rawDesc = nil
	file_SetPlayerPropRsp_proto_goTypes = nil
	file_SetPlayerPropRsp_proto_depIdxs = nil
}
