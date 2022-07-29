// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChangeAvatarReq.proto

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

type ChangeAvatarReq_CmdId int32

const (
	ChangeAvatarReq_ENET_CHANNEL_ID  ChangeAvatarReq_CmdId = 0
	ChangeAvatarReq_NONE             ChangeAvatarReq_CmdId = 0
	ChangeAvatarReq_ENET_IS_RELIABLE ChangeAvatarReq_CmdId = 1
	ChangeAvatarReq_IS_ALLOW_CLIENT  ChangeAvatarReq_CmdId = 1
	ChangeAvatarReq_CMD_ID           ChangeAvatarReq_CmdId = 1669
)

// Enum value maps for ChangeAvatarReq_CmdId.
var (
	ChangeAvatarReq_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		1669: "CMD_ID",
	}
	ChangeAvatarReq_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           1669,
	}
)

func (x ChangeAvatarReq_CmdId) Enum() *ChangeAvatarReq_CmdId {
	p := new(ChangeAvatarReq_CmdId)
	*p = x
	return p
}

func (x ChangeAvatarReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeAvatarReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ChangeAvatarReq_proto_enumTypes[0].Descriptor()
}

func (ChangeAvatarReq_CmdId) Type() protoreflect.EnumType {
	return &file_ChangeAvatarReq_proto_enumTypes[0]
}

func (x ChangeAvatarReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeAvatarReq_CmdId.Descriptor instead.
func (ChangeAvatarReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ChangeAvatarReq_proto_rawDescGZIP(), []int{0, 0}
}

type ChangeAvatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid    uint64  `protobuf:"varint,1,opt,name=guid,proto3" json:"guid,omitempty"`
	SkillId uint32  `protobuf:"varint,2,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	IsMove  bool    `protobuf:"varint,3,opt,name=is_move,json=isMove,proto3" json:"is_move,omitempty"`
	MovePos *Vector `protobuf:"bytes,4,opt,name=move_pos,json=movePos,proto3" json:"move_pos,omitempty"`
}

func (x *ChangeAvatarReq) Reset() {
	*x = ChangeAvatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChangeAvatarReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAvatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAvatarReq) ProtoMessage() {}

func (x *ChangeAvatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_ChangeAvatarReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAvatarReq.ProtoReflect.Descriptor instead.
func (*ChangeAvatarReq) Descriptor() ([]byte, []int) {
	return file_ChangeAvatarReq_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeAvatarReq) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *ChangeAvatarReq) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *ChangeAvatarReq) GetIsMove() bool {
	if x != nil {
		return x.IsMove
	}
	return false
}

func (x *ChangeAvatarReq) GetMovePos() *Vector {
	if x != nil {
		return x.MovePos
	}
	return nil
}

var File_ChangeAvatarReq_proto protoreflect.FileDescriptor

var file_ChangeAvatarReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a,
	0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x67, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x65,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x6f, 0x73, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x85, 0x0d, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChangeAvatarReq_proto_rawDescOnce sync.Once
	file_ChangeAvatarReq_proto_rawDescData = file_ChangeAvatarReq_proto_rawDesc
)

func file_ChangeAvatarReq_proto_rawDescGZIP() []byte {
	file_ChangeAvatarReq_proto_rawDescOnce.Do(func() {
		file_ChangeAvatarReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeAvatarReq_proto_rawDescData)
	})
	return file_ChangeAvatarReq_proto_rawDescData
}

var file_ChangeAvatarReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChangeAvatarReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChangeAvatarReq_proto_goTypes = []interface{}{
	(ChangeAvatarReq_CmdId)(0), // 0: proto.ChangeAvatarReq.CmdId
	(*ChangeAvatarReq)(nil),    // 1: proto.ChangeAvatarReq
	(*Vector)(nil),             // 2: proto.Vector
}
var file_ChangeAvatarReq_proto_depIdxs = []int32{
	2, // 0: proto.ChangeAvatarReq.move_pos:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChangeAvatarReq_proto_init() }
func file_ChangeAvatarReq_proto_init() {
	if File_ChangeAvatarReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChangeAvatarReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAvatarReq); i {
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
			RawDescriptor: file_ChangeAvatarReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeAvatarReq_proto_goTypes,
		DependencyIndexes: file_ChangeAvatarReq_proto_depIdxs,
		EnumInfos:         file_ChangeAvatarReq_proto_enumTypes,
		MessageInfos:      file_ChangeAvatarReq_proto_msgTypes,
	}.Build()
	File_ChangeAvatarReq_proto = out.File
	file_ChangeAvatarReq_proto_rawDesc = nil
	file_ChangeAvatarReq_proto_goTypes = nil
	file_ChangeAvatarReq_proto_depIdxs = nil
}
