// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: SceneEntityDrownReq.proto

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

type SceneEntityDrownReq_CmdId int32

const (
	SceneEntityDrownReq_NONE             SceneEntityDrownReq_CmdId = 0
	SceneEntityDrownReq_ENET_CHANNEL_ID  SceneEntityDrownReq_CmdId = 0
	SceneEntityDrownReq_ENET_IS_RELIABLE SceneEntityDrownReq_CmdId = 1
	SceneEntityDrownReq_IS_ALLOW_CLIENT  SceneEntityDrownReq_CmdId = 1
	SceneEntityDrownReq_CMD_ID           SceneEntityDrownReq_CmdId = 212
)

// Enum value maps for SceneEntityDrownReq_CmdId.
var (
	SceneEntityDrownReq_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		212: "CMD_ID",
	}
	SceneEntityDrownReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           212,
	}
)

func (x SceneEntityDrownReq_CmdId) Enum() *SceneEntityDrownReq_CmdId {
	p := new(SceneEntityDrownReq_CmdId)
	*p = x
	return p
}

func (x SceneEntityDrownReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneEntityDrownReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneEntityDrownReq_proto_enumTypes[0].Descriptor()
}

func (SceneEntityDrownReq_CmdId) Type() protoreflect.EnumType {
	return &file_SceneEntityDrownReq_proto_enumTypes[0]
}

func (x SceneEntityDrownReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneEntityDrownReq_CmdId.Descriptor instead.
func (SceneEntityDrownReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneEntityDrownReq_proto_rawDescGZIP(), []int{0, 0}
}

type SceneEntityDrownReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId uint32 `protobuf:"varint,14,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *SceneEntityDrownReq) Reset() {
	*x = SceneEntityDrownReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityDrownReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityDrownReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityDrownReq) ProtoMessage() {}

func (x *SceneEntityDrownReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityDrownReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityDrownReq.ProtoReflect.Descriptor instead.
func (*SceneEntityDrownReq) Descriptor() ([]byte, []int) {
	return file_SceneEntityDrownReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityDrownReq) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_SceneEntityDrownReq_proto protoreflect.FileDescriptor

var file_SceneEntityDrownReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x72, 0x6f,
	0x77, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x13, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x44, 0x72, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0xd4, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityDrownReq_proto_rawDescOnce sync.Once
	file_SceneEntityDrownReq_proto_rawDescData = file_SceneEntityDrownReq_proto_rawDesc
)

func file_SceneEntityDrownReq_proto_rawDescGZIP() []byte {
	file_SceneEntityDrownReq_proto_rawDescOnce.Do(func() {
		file_SceneEntityDrownReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityDrownReq_proto_rawDescData)
	})
	return file_SceneEntityDrownReq_proto_rawDescData
}

var file_SceneEntityDrownReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneEntityDrownReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityDrownReq_proto_goTypes = []interface{}{
	(SceneEntityDrownReq_CmdId)(0), // 0: proto.SceneEntityDrownReq.CmdId
	(*SceneEntityDrownReq)(nil),    // 1: proto.SceneEntityDrownReq
}
var file_SceneEntityDrownReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneEntityDrownReq_proto_init() }
func file_SceneEntityDrownReq_proto_init() {
	if File_SceneEntityDrownReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityDrownReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityDrownReq); i {
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
			RawDescriptor: file_SceneEntityDrownReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityDrownReq_proto_goTypes,
		DependencyIndexes: file_SceneEntityDrownReq_proto_depIdxs,
		EnumInfos:         file_SceneEntityDrownReq_proto_enumTypes,
		MessageInfos:      file_SceneEntityDrownReq_proto_msgTypes,
	}.Build()
	File_SceneEntityDrownReq_proto = out.File
	file_SceneEntityDrownReq_proto_rawDesc = nil
	file_SceneEntityDrownReq_proto_goTypes = nil
	file_SceneEntityDrownReq_proto_depIdxs = nil
}
