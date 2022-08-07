// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarWearFlycloakReq.proto

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

type AvatarWearFlycloakReq_CmdId int32

const (
	AvatarWearFlycloakReq_NONE             AvatarWearFlycloakReq_CmdId = 0
	AvatarWearFlycloakReq_ENET_CHANNEL_ID  AvatarWearFlycloakReq_CmdId = 0
	AvatarWearFlycloakReq_ENET_IS_RELIABLE AvatarWearFlycloakReq_CmdId = 1
	AvatarWearFlycloakReq_IS_ALLOW_CLIENT  AvatarWearFlycloakReq_CmdId = 1
	AvatarWearFlycloakReq_CMD_ID           AvatarWearFlycloakReq_CmdId = 1728
)

// Enum value maps for AvatarWearFlycloakReq_CmdId.
var (
	AvatarWearFlycloakReq_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		1728: "CMD_ID",
	}
	AvatarWearFlycloakReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           1728,
	}
)

func (x AvatarWearFlycloakReq_CmdId) Enum() *AvatarWearFlycloakReq_CmdId {
	p := new(AvatarWearFlycloakReq_CmdId)
	*p = x
	return p
}

func (x AvatarWearFlycloakReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarWearFlycloakReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarWearFlycloakReq_proto_enumTypes[0].Descriptor()
}

func (AvatarWearFlycloakReq_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarWearFlycloakReq_proto_enumTypes[0]
}

func (x AvatarWearFlycloakReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarWearFlycloakReq_CmdId.Descriptor instead.
func (AvatarWearFlycloakReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarWearFlycloakReq_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarWearFlycloakReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid uint64 `protobuf:"varint,11,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	FlycloakId uint32 `protobuf:"varint,9,opt,name=flycloak_id,json=flycloakId,proto3" json:"flycloak_id,omitempty"`
}

func (x *AvatarWearFlycloakReq) Reset() {
	*x = AvatarWearFlycloakReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarWearFlycloakReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarWearFlycloakReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarWearFlycloakReq) ProtoMessage() {}

func (x *AvatarWearFlycloakReq) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarWearFlycloakReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarWearFlycloakReq.ProtoReflect.Descriptor instead.
func (*AvatarWearFlycloakReq) Descriptor() ([]byte, []int) {
	return file_AvatarWearFlycloakReq_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarWearFlycloakReq) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarWearFlycloakReq) GetFlycloakId() uint32 {
	if x != nil {
		return x.FlycloakId
	}
	return 0
}

var File_AvatarWearFlycloakReq_proto protoreflect.FileDescriptor

var file_AvatarWearFlycloakReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x57, 0x65, 0x61, 0x72, 0x46, 0x6c, 0x79, 0x63,
	0x6c, 0x6f, 0x61, 0x6b, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x57,
	0x65, 0x61, 0x72, 0x46, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x49, 0x64,
	0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xc0, 0x0d,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarWearFlycloakReq_proto_rawDescOnce sync.Once
	file_AvatarWearFlycloakReq_proto_rawDescData = file_AvatarWearFlycloakReq_proto_rawDesc
)

func file_AvatarWearFlycloakReq_proto_rawDescGZIP() []byte {
	file_AvatarWearFlycloakReq_proto_rawDescOnce.Do(func() {
		file_AvatarWearFlycloakReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarWearFlycloakReq_proto_rawDescData)
	})
	return file_AvatarWearFlycloakReq_proto_rawDescData
}

var file_AvatarWearFlycloakReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarWearFlycloakReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarWearFlycloakReq_proto_goTypes = []interface{}{
	(AvatarWearFlycloakReq_CmdId)(0), // 0: proto.AvatarWearFlycloakReq.CmdId
	(*AvatarWearFlycloakReq)(nil),    // 1: proto.AvatarWearFlycloakReq
}
var file_AvatarWearFlycloakReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarWearFlycloakReq_proto_init() }
func file_AvatarWearFlycloakReq_proto_init() {
	if File_AvatarWearFlycloakReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarWearFlycloakReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarWearFlycloakReq); i {
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
			RawDescriptor: file_AvatarWearFlycloakReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarWearFlycloakReq_proto_goTypes,
		DependencyIndexes: file_AvatarWearFlycloakReq_proto_depIdxs,
		EnumInfos:         file_AvatarWearFlycloakReq_proto_enumTypes,
		MessageInfos:      file_AvatarWearFlycloakReq_proto_msgTypes,
	}.Build()
	File_AvatarWearFlycloakReq_proto = out.File
	file_AvatarWearFlycloakReq_proto_rawDesc = nil
	file_AvatarWearFlycloakReq_proto_goTypes = nil
	file_AvatarWearFlycloakReq_proto_depIdxs = nil
}
