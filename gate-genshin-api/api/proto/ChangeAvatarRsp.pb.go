// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChangeAvatarRsp.proto

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

type ChangeAvatarRsp_CmdId int32

const (
	ChangeAvatarRsp_NONE             ChangeAvatarRsp_CmdId = 0
	ChangeAvatarRsp_ENET_CHANNEL_ID  ChangeAvatarRsp_CmdId = 0
	ChangeAvatarRsp_ENET_IS_RELIABLE ChangeAvatarRsp_CmdId = 1
	ChangeAvatarRsp_CMD_ID           ChangeAvatarRsp_CmdId = 1634
)

// Enum value maps for ChangeAvatarRsp_CmdId.
var (
	ChangeAvatarRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1634: "CMD_ID",
	}
	ChangeAvatarRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1634,
	}
)

func (x ChangeAvatarRsp_CmdId) Enum() *ChangeAvatarRsp_CmdId {
	p := new(ChangeAvatarRsp_CmdId)
	*p = x
	return p
}

func (x ChangeAvatarRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeAvatarRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ChangeAvatarRsp_proto_enumTypes[0].Descriptor()
}

func (ChangeAvatarRsp_CmdId) Type() protoreflect.EnumType {
	return &file_ChangeAvatarRsp_proto_enumTypes[0]
}

func (x ChangeAvatarRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeAvatarRsp_CmdId.Descriptor instead.
func (ChangeAvatarRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ChangeAvatarRsp_proto_rawDescGZIP(), []int{0, 0}
}

type ChangeAvatarRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CurGuid uint64 `protobuf:"varint,2,opt,name=cur_guid,json=curGuid,proto3" json:"cur_guid,omitempty"`
	SkillId uint32 `protobuf:"varint,3,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *ChangeAvatarRsp) Reset() {
	*x = ChangeAvatarRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChangeAvatarRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAvatarRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAvatarRsp) ProtoMessage() {}

func (x *ChangeAvatarRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChangeAvatarRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAvatarRsp.ProtoReflect.Descriptor instead.
func (*ChangeAvatarRsp) Descriptor() ([]byte, []int) {
	return file_ChangeAvatarRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeAvatarRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChangeAvatarRsp) GetCurGuid() uint64 {
	if x != nil {
		return x.CurGuid
	}
	return 0
}

func (x *ChangeAvatarRsp) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

var File_ChangeAvatarRsp_proto protoreflect.FileDescriptor

var file_ChangeAvatarRsp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x63, 0x75, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xe2, 0x0c, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChangeAvatarRsp_proto_rawDescOnce sync.Once
	file_ChangeAvatarRsp_proto_rawDescData = file_ChangeAvatarRsp_proto_rawDesc
)

func file_ChangeAvatarRsp_proto_rawDescGZIP() []byte {
	file_ChangeAvatarRsp_proto_rawDescOnce.Do(func() {
		file_ChangeAvatarRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeAvatarRsp_proto_rawDescData)
	})
	return file_ChangeAvatarRsp_proto_rawDescData
}

var file_ChangeAvatarRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChangeAvatarRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChangeAvatarRsp_proto_goTypes = []interface{}{
	(ChangeAvatarRsp_CmdId)(0), // 0: proto.ChangeAvatarRsp.CmdId
	(*ChangeAvatarRsp)(nil),    // 1: proto.ChangeAvatarRsp
}
var file_ChangeAvatarRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChangeAvatarRsp_proto_init() }
func file_ChangeAvatarRsp_proto_init() {
	if File_ChangeAvatarRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChangeAvatarRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAvatarRsp); i {
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
			RawDescriptor: file_ChangeAvatarRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeAvatarRsp_proto_goTypes,
		DependencyIndexes: file_ChangeAvatarRsp_proto_depIdxs,
		EnumInfos:         file_ChangeAvatarRsp_proto_enumTypes,
		MessageInfos:      file_ChangeAvatarRsp_proto_msgTypes,
	}.Build()
	File_ChangeAvatarRsp_proto = out.File
	file_ChangeAvatarRsp_proto_rawDesc = nil
	file_ChangeAvatarRsp_proto_goTypes = nil
	file_ChangeAvatarRsp_proto_depIdxs = nil
}
