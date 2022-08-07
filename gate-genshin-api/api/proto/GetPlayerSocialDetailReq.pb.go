// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: GetPlayerSocialDetailReq.proto

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

type GetPlayerSocialDetailReq_CmdId int32

const (
	GetPlayerSocialDetailReq_NONE             GetPlayerSocialDetailReq_CmdId = 0
	GetPlayerSocialDetailReq_ENET_CHANNEL_ID  GetPlayerSocialDetailReq_CmdId = 0
	GetPlayerSocialDetailReq_ENET_IS_RELIABLE GetPlayerSocialDetailReq_CmdId = 1
	GetPlayerSocialDetailReq_IS_ALLOW_CLIENT  GetPlayerSocialDetailReq_CmdId = 1
	GetPlayerSocialDetailReq_CMD_ID           GetPlayerSocialDetailReq_CmdId = 4049
)

// Enum value maps for GetPlayerSocialDetailReq_CmdId.
var (
	GetPlayerSocialDetailReq_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		4049: "CMD_ID",
	}
	GetPlayerSocialDetailReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           4049,
	}
)

func (x GetPlayerSocialDetailReq_CmdId) Enum() *GetPlayerSocialDetailReq_CmdId {
	p := new(GetPlayerSocialDetailReq_CmdId)
	*p = x
	return p
}

func (x GetPlayerSocialDetailReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPlayerSocialDetailReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetPlayerSocialDetailReq_proto_enumTypes[0].Descriptor()
}

func (GetPlayerSocialDetailReq_CmdId) Type() protoreflect.EnumType {
	return &file_GetPlayerSocialDetailReq_proto_enumTypes[0]
}

func (x GetPlayerSocialDetailReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPlayerSocialDetailReq_CmdId.Descriptor instead.
func (GetPlayerSocialDetailReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetPlayerSocialDetailReq_proto_rawDescGZIP(), []int{0, 0}
}

type GetPlayerSocialDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,14,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetPlayerSocialDetailReq) Reset() {
	*x = GetPlayerSocialDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPlayerSocialDetailReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerSocialDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerSocialDetailReq) ProtoMessage() {}

func (x *GetPlayerSocialDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetPlayerSocialDetailReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerSocialDetailReq.ProtoReflect.Descriptor instead.
func (*GetPlayerSocialDetailReq) Descriptor() ([]byte, []int) {
	return file_GetPlayerSocialDetailReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerSocialDetailReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_GetPlayerSocialDetailReq_proto protoreflect.FileDescriptor

var file_GetPlayerSocialDetailReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xd1, 0x1f, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPlayerSocialDetailReq_proto_rawDescOnce sync.Once
	file_GetPlayerSocialDetailReq_proto_rawDescData = file_GetPlayerSocialDetailReq_proto_rawDesc
)

func file_GetPlayerSocialDetailReq_proto_rawDescGZIP() []byte {
	file_GetPlayerSocialDetailReq_proto_rawDescOnce.Do(func() {
		file_GetPlayerSocialDetailReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPlayerSocialDetailReq_proto_rawDescData)
	})
	return file_GetPlayerSocialDetailReq_proto_rawDescData
}

var file_GetPlayerSocialDetailReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetPlayerSocialDetailReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPlayerSocialDetailReq_proto_goTypes = []interface{}{
	(GetPlayerSocialDetailReq_CmdId)(0), // 0: proto.GetPlayerSocialDetailReq.CmdId
	(*GetPlayerSocialDetailReq)(nil),    // 1: proto.GetPlayerSocialDetailReq
}
var file_GetPlayerSocialDetailReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetPlayerSocialDetailReq_proto_init() }
func file_GetPlayerSocialDetailReq_proto_init() {
	if File_GetPlayerSocialDetailReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetPlayerSocialDetailReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerSocialDetailReq); i {
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
			RawDescriptor: file_GetPlayerSocialDetailReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPlayerSocialDetailReq_proto_goTypes,
		DependencyIndexes: file_GetPlayerSocialDetailReq_proto_depIdxs,
		EnumInfos:         file_GetPlayerSocialDetailReq_proto_enumTypes,
		MessageInfos:      file_GetPlayerSocialDetailReq_proto_msgTypes,
	}.Build()
	File_GetPlayerSocialDetailReq_proto = out.File
	file_GetPlayerSocialDetailReq_proto_rawDesc = nil
	file_GetPlayerSocialDetailReq_proto_goTypes = nil
	file_GetPlayerSocialDetailReq_proto_depIdxs = nil
}
