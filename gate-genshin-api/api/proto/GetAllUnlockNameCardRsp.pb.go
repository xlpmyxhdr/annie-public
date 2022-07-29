// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: GetAllUnlockNameCardRsp.proto

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

type GetAllUnlockNameCardRsp_CmdId int32

const (
	GetAllUnlockNameCardRsp_ENET_CHANNEL_ID  GetAllUnlockNameCardRsp_CmdId = 0
	GetAllUnlockNameCardRsp_NONE             GetAllUnlockNameCardRsp_CmdId = 0
	GetAllUnlockNameCardRsp_ENET_IS_RELIABLE GetAllUnlockNameCardRsp_CmdId = 1
	GetAllUnlockNameCardRsp_IS_ALLOW_CLIENT  GetAllUnlockNameCardRsp_CmdId = 1
	GetAllUnlockNameCardRsp_CMD_ID           GetAllUnlockNameCardRsp_CmdId = 4076
)

// Enum value maps for GetAllUnlockNameCardRsp_CmdId.
var (
	GetAllUnlockNameCardRsp_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		4076: "CMD_ID",
	}
	GetAllUnlockNameCardRsp_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           4076,
	}
)

func (x GetAllUnlockNameCardRsp_CmdId) Enum() *GetAllUnlockNameCardRsp_CmdId {
	p := new(GetAllUnlockNameCardRsp_CmdId)
	*p = x
	return p
}

func (x GetAllUnlockNameCardRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetAllUnlockNameCardRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetAllUnlockNameCardRsp_proto_enumTypes[0].Descriptor()
}

func (GetAllUnlockNameCardRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetAllUnlockNameCardRsp_proto_enumTypes[0]
}

func (x GetAllUnlockNameCardRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetAllUnlockNameCardRsp_CmdId.Descriptor instead.
func (GetAllUnlockNameCardRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetAllUnlockNameCardRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetAllUnlockNameCardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode      int32    `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	NameCardList []uint32 `protobuf:"varint,2,rep,packed,name=name_card_list,json=nameCardList,proto3" json:"name_card_list,omitempty"`
}

func (x *GetAllUnlockNameCardRsp) Reset() {
	*x = GetAllUnlockNameCardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAllUnlockNameCardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUnlockNameCardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUnlockNameCardRsp) ProtoMessage() {}

func (x *GetAllUnlockNameCardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAllUnlockNameCardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUnlockNameCardRsp.ProtoReflect.Descriptor instead.
func (*GetAllUnlockNameCardRsp) Descriptor() ([]byte, []int) {
	return file_GetAllUnlockNameCardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllUnlockNameCardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAllUnlockNameCardRsp) GetNameCardList() []uint32 {
	if x != nil {
		return x.NameCardList
	}
	return nil
}

var File_GetAllUnlockNameCardRsp_proto protoreflect.FileDescriptor

var file_GetAllUnlockNameCardRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xec, 0x1f, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAllUnlockNameCardRsp_proto_rawDescOnce sync.Once
	file_GetAllUnlockNameCardRsp_proto_rawDescData = file_GetAllUnlockNameCardRsp_proto_rawDesc
)

func file_GetAllUnlockNameCardRsp_proto_rawDescGZIP() []byte {
	file_GetAllUnlockNameCardRsp_proto_rawDescOnce.Do(func() {
		file_GetAllUnlockNameCardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAllUnlockNameCardRsp_proto_rawDescData)
	})
	return file_GetAllUnlockNameCardRsp_proto_rawDescData
}

var file_GetAllUnlockNameCardRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetAllUnlockNameCardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAllUnlockNameCardRsp_proto_goTypes = []interface{}{
	(GetAllUnlockNameCardRsp_CmdId)(0), // 0: proto.GetAllUnlockNameCardRsp.CmdId
	(*GetAllUnlockNameCardRsp)(nil),    // 1: proto.GetAllUnlockNameCardRsp
}
var file_GetAllUnlockNameCardRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetAllUnlockNameCardRsp_proto_init() }
func file_GetAllUnlockNameCardRsp_proto_init() {
	if File_GetAllUnlockNameCardRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetAllUnlockNameCardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUnlockNameCardRsp); i {
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
			RawDescriptor: file_GetAllUnlockNameCardRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAllUnlockNameCardRsp_proto_goTypes,
		DependencyIndexes: file_GetAllUnlockNameCardRsp_proto_depIdxs,
		EnumInfos:         file_GetAllUnlockNameCardRsp_proto_enumTypes,
		MessageInfos:      file_GetAllUnlockNameCardRsp_proto_msgTypes,
	}.Build()
	File_GetAllUnlockNameCardRsp_proto = out.File
	file_GetAllUnlockNameCardRsp_proto_rawDesc = nil
	file_GetAllUnlockNameCardRsp_proto_goTypes = nil
	file_GetAllUnlockNameCardRsp_proto_depIdxs = nil
}
