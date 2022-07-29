// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: DeleteFriendReq.proto

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

type DeleteFriendReq_CmdId int32

const (
	DeleteFriendReq_ENET_CHANNEL_ID  DeleteFriendReq_CmdId = 0
	DeleteFriendReq_NONE             DeleteFriendReq_CmdId = 0
	DeleteFriendReq_ENET_IS_RELIABLE DeleteFriendReq_CmdId = 1
	DeleteFriendReq_IS_ALLOW_CLIENT  DeleteFriendReq_CmdId = 1
	DeleteFriendReq_CMD_ID           DeleteFriendReq_CmdId = 4012
)

// Enum value maps for DeleteFriendReq_CmdId.
var (
	DeleteFriendReq_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		4012: "CMD_ID",
	}
	DeleteFriendReq_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           4012,
	}
)

func (x DeleteFriendReq_CmdId) Enum() *DeleteFriendReq_CmdId {
	p := new(DeleteFriendReq_CmdId)
	*p = x
	return p
}

func (x DeleteFriendReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteFriendReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DeleteFriendReq_proto_enumTypes[0].Descriptor()
}

func (DeleteFriendReq_CmdId) Type() protoreflect.EnumType {
	return &file_DeleteFriendReq_proto_enumTypes[0]
}

func (x DeleteFriendReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteFriendReq_CmdId.Descriptor instead.
func (DeleteFriendReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DeleteFriendReq_proto_rawDescGZIP(), []int{0, 0}
}

type DeleteFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid uint32 `protobuf:"varint,1,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *DeleteFriendReq) Reset() {
	*x = DeleteFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DeleteFriendReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendReq) ProtoMessage() {}

func (x *DeleteFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_DeleteFriendReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendReq.ProtoReflect.Descriptor instead.
func (*DeleteFriendReq) Descriptor() ([]byte, []int) {
	return file_DeleteFriendReq_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteFriendReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

var File_DeleteFriendReq_proto protoreflect.FileDescriptor

var file_DeleteFriendReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69,
	0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xac,
	0x1f, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DeleteFriendReq_proto_rawDescOnce sync.Once
	file_DeleteFriendReq_proto_rawDescData = file_DeleteFriendReq_proto_rawDesc
)

func file_DeleteFriendReq_proto_rawDescGZIP() []byte {
	file_DeleteFriendReq_proto_rawDescOnce.Do(func() {
		file_DeleteFriendReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_DeleteFriendReq_proto_rawDescData)
	})
	return file_DeleteFriendReq_proto_rawDescData
}

var file_DeleteFriendReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DeleteFriendReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DeleteFriendReq_proto_goTypes = []interface{}{
	(DeleteFriendReq_CmdId)(0), // 0: proto.DeleteFriendReq.CmdId
	(*DeleteFriendReq)(nil),    // 1: proto.DeleteFriendReq
}
var file_DeleteFriendReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DeleteFriendReq_proto_init() }
func file_DeleteFriendReq_proto_init() {
	if File_DeleteFriendReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DeleteFriendReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendReq); i {
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
			RawDescriptor: file_DeleteFriendReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DeleteFriendReq_proto_goTypes,
		DependencyIndexes: file_DeleteFriendReq_proto_depIdxs,
		EnumInfos:         file_DeleteFriendReq_proto_enumTypes,
		MessageInfos:      file_DeleteFriendReq_proto_msgTypes,
	}.Build()
	File_DeleteFriendReq_proto = out.File
	file_DeleteFriendReq_proto_rawDesc = nil
	file_DeleteFriendReq_proto_goTypes = nil
	file_DeleteFriendReq_proto_depIdxs = nil
}
