// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: SetPlayerSignatureRsp.proto

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

type SetPlayerSignatureRsp_CmdId int32

const (
	SetPlayerSignatureRsp_NONE             SetPlayerSignatureRsp_CmdId = 0
	SetPlayerSignatureRsp_ENET_CHANNEL_ID  SetPlayerSignatureRsp_CmdId = 0
	SetPlayerSignatureRsp_ENET_IS_RELIABLE SetPlayerSignatureRsp_CmdId = 1
	SetPlayerSignatureRsp_CMD_ID           SetPlayerSignatureRsp_CmdId = 4019
)

// Enum value maps for SetPlayerSignatureRsp_CmdId.
var (
	SetPlayerSignatureRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		4019: "CMD_ID",
	}
	SetPlayerSignatureRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           4019,
	}
)

func (x SetPlayerSignatureRsp_CmdId) Enum() *SetPlayerSignatureRsp_CmdId {
	p := new(SetPlayerSignatureRsp_CmdId)
	*p = x
	return p
}

func (x SetPlayerSignatureRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetPlayerSignatureRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SetPlayerSignatureRsp_proto_enumTypes[0].Descriptor()
}

func (SetPlayerSignatureRsp_CmdId) Type() protoreflect.EnumType {
	return &file_SetPlayerSignatureRsp_proto_enumTypes[0]
}

func (x SetPlayerSignatureRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetPlayerSignatureRsp_CmdId.Descriptor instead.
func (SetPlayerSignatureRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SetPlayerSignatureRsp_proto_rawDescGZIP(), []int{0, 0}
}

type SetPlayerSignatureRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SetPlayerSignatureRsp) Reset() {
	*x = SetPlayerSignatureRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetPlayerSignatureRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerSignatureRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerSignatureRsp) ProtoMessage() {}

func (x *SetPlayerSignatureRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetPlayerSignatureRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerSignatureRsp.ProtoReflect.Descriptor instead.
func (*SetPlayerSignatureRsp) Descriptor() ([]byte, []int) {
	return file_SetPlayerSignatureRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetPlayerSignatureRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SetPlayerSignatureRsp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_SetPlayerSignatureRsp_proto protoreflect.FileDescriptor

var file_SetPlayerSignatureRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb3,
	0x1f, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetPlayerSignatureRsp_proto_rawDescOnce sync.Once
	file_SetPlayerSignatureRsp_proto_rawDescData = file_SetPlayerSignatureRsp_proto_rawDesc
)

func file_SetPlayerSignatureRsp_proto_rawDescGZIP() []byte {
	file_SetPlayerSignatureRsp_proto_rawDescOnce.Do(func() {
		file_SetPlayerSignatureRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetPlayerSignatureRsp_proto_rawDescData)
	})
	return file_SetPlayerSignatureRsp_proto_rawDescData
}

var file_SetPlayerSignatureRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SetPlayerSignatureRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetPlayerSignatureRsp_proto_goTypes = []interface{}{
	(SetPlayerSignatureRsp_CmdId)(0), // 0: proto.SetPlayerSignatureRsp.CmdId
	(*SetPlayerSignatureRsp)(nil),    // 1: proto.SetPlayerSignatureRsp
}
var file_SetPlayerSignatureRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetPlayerSignatureRsp_proto_init() }
func file_SetPlayerSignatureRsp_proto_init() {
	if File_SetPlayerSignatureRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetPlayerSignatureRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerSignatureRsp); i {
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
			RawDescriptor: file_SetPlayerSignatureRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetPlayerSignatureRsp_proto_goTypes,
		DependencyIndexes: file_SetPlayerSignatureRsp_proto_depIdxs,
		EnumInfos:         file_SetPlayerSignatureRsp_proto_enumTypes,
		MessageInfos:      file_SetPlayerSignatureRsp_proto_msgTypes,
	}.Build()
	File_SetPlayerSignatureRsp_proto = out.File
	file_SetPlayerSignatureRsp_proto_rawDesc = nil
	file_SetPlayerSignatureRsp_proto_goTypes = nil
	file_SetPlayerSignatureRsp_proto_depIdxs = nil
}
