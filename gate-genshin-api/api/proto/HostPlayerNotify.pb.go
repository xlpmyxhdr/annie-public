// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: HostPlayerNotify.proto

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

type HostPlayerNotify_CmdId int32

const (
	HostPlayerNotify_NONE             HostPlayerNotify_CmdId = 0
	HostPlayerNotify_ENET_CHANNEL_ID  HostPlayerNotify_CmdId = 0
	HostPlayerNotify_ENET_IS_RELIABLE HostPlayerNotify_CmdId = 1
	HostPlayerNotify_CMD_ID           HostPlayerNotify_CmdId = 341
)

// Enum value maps for HostPlayerNotify_CmdId.
var (
	HostPlayerNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		341: "CMD_ID",
	}
	HostPlayerNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           341,
	}
)

func (x HostPlayerNotify_CmdId) Enum() *HostPlayerNotify_CmdId {
	p := new(HostPlayerNotify_CmdId)
	*p = x
	return p
}

func (x HostPlayerNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostPlayerNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_HostPlayerNotify_proto_enumTypes[0].Descriptor()
}

func (HostPlayerNotify_CmdId) Type() protoreflect.EnumType {
	return &file_HostPlayerNotify_proto_enumTypes[0]
}

func (x HostPlayerNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostPlayerNotify_CmdId.Descriptor instead.
func (HostPlayerNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_HostPlayerNotify_proto_rawDescGZIP(), []int{0, 0}
}

type HostPlayerNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostUid    uint32 `protobuf:"varint,10,opt,name=host_uid,json=hostUid,proto3" json:"host_uid,omitempty"`
	HostPeerId uint32 `protobuf:"varint,7,opt,name=host_peer_id,json=hostPeerId,proto3" json:"host_peer_id,omitempty"`
}

func (x *HostPlayerNotify) Reset() {
	*x = HostPlayerNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HostPlayerNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostPlayerNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostPlayerNotify) ProtoMessage() {}

func (x *HostPlayerNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HostPlayerNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostPlayerNotify.ProtoReflect.Descriptor instead.
func (*HostPlayerNotify) Descriptor() ([]byte, []int) {
	return file_HostPlayerNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HostPlayerNotify) GetHostUid() uint32 {
	if x != nil {
		return x.HostUid
	}
	return 0
}

func (x *HostPlayerNotify) GetHostPeerId() uint32 {
	if x != nil {
		return x.HostPeerId
	}
	return 0
}

var File_HostPlayerNotify_proto protoreflect.FileDescriptor

var file_HostPlayerNotify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x10, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd5, 0x02, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HostPlayerNotify_proto_rawDescOnce sync.Once
	file_HostPlayerNotify_proto_rawDescData = file_HostPlayerNotify_proto_rawDesc
)

func file_HostPlayerNotify_proto_rawDescGZIP() []byte {
	file_HostPlayerNotify_proto_rawDescOnce.Do(func() {
		file_HostPlayerNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HostPlayerNotify_proto_rawDescData)
	})
	return file_HostPlayerNotify_proto_rawDescData
}

var file_HostPlayerNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HostPlayerNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HostPlayerNotify_proto_goTypes = []interface{}{
	(HostPlayerNotify_CmdId)(0), // 0: proto.HostPlayerNotify.CmdId
	(*HostPlayerNotify)(nil),    // 1: proto.HostPlayerNotify
}
var file_HostPlayerNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HostPlayerNotify_proto_init() }
func file_HostPlayerNotify_proto_init() {
	if File_HostPlayerNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HostPlayerNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostPlayerNotify); i {
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
			RawDescriptor: file_HostPlayerNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HostPlayerNotify_proto_goTypes,
		DependencyIndexes: file_HostPlayerNotify_proto_depIdxs,
		EnumInfos:         file_HostPlayerNotify_proto_enumTypes,
		MessageInfos:      file_HostPlayerNotify_proto_msgTypes,
	}.Build()
	File_HostPlayerNotify_proto = out.File
	file_HostPlayerNotify_proto_rawDesc = nil
	file_HostPlayerNotify_proto_goTypes = nil
	file_HostPlayerNotify_proto_depIdxs = nil
}
