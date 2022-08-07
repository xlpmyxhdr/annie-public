// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: DoGachaReq.proto

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

type DoGachaReq_CmdId int32

const (
	DoGachaReq_NONE             DoGachaReq_CmdId = 0
	DoGachaReq_ENET_CHANNEL_ID  DoGachaReq_CmdId = 0
	DoGachaReq_ENET_IS_RELIABLE DoGachaReq_CmdId = 1
	DoGachaReq_IS_ALLOW_CLIENT  DoGachaReq_CmdId = 1
	DoGachaReq_CMD_ID           DoGachaReq_CmdId = 1541
)

// Enum value maps for DoGachaReq_CmdId.
var (
	DoGachaReq_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		1541: "CMD_ID",
	}
	DoGachaReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           1541,
	}
)

func (x DoGachaReq_CmdId) Enum() *DoGachaReq_CmdId {
	p := new(DoGachaReq_CmdId)
	*p = x
	return p
}

func (x DoGachaReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DoGachaReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DoGachaReq_proto_enumTypes[0].Descriptor()
}

func (DoGachaReq_CmdId) Type() protoreflect.EnumType {
	return &file_DoGachaReq_proto_enumTypes[0]
}

func (x DoGachaReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DoGachaReq_CmdId.Descriptor instead.
func (DoGachaReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DoGachaReq_proto_rawDescGZIP(), []int{0, 0}
}

type DoGachaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaType       uint32 `protobuf:"varint,10,opt,name=gacha_type,json=gachaType,proto3" json:"gacha_type,omitempty"`
	GachaTimes      uint32 `protobuf:"varint,15,opt,name=gacha_times,json=gachaTimes,proto3" json:"gacha_times,omitempty"`
	GachaRandom     uint32 `protobuf:"varint,5,opt,name=gacha_random,json=gachaRandom,proto3" json:"gacha_random,omitempty"`
	GachaScheduleId uint32 `protobuf:"varint,6,opt,name=gacha_schedule_id,json=gachaScheduleId,proto3" json:"gacha_schedule_id,omitempty"`
	GachaTag        string `protobuf:"bytes,4,opt,name=gacha_tag,json=gachaTag,proto3" json:"gacha_tag,omitempty"`
}

func (x *DoGachaReq) Reset() {
	*x = DoGachaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DoGachaReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoGachaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoGachaReq) ProtoMessage() {}

func (x *DoGachaReq) ProtoReflect() protoreflect.Message {
	mi := &file_DoGachaReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoGachaReq.ProtoReflect.Descriptor instead.
func (*DoGachaReq) Descriptor() ([]byte, []int) {
	return file_DoGachaReq_proto_rawDescGZIP(), []int{0}
}

func (x *DoGachaReq) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *DoGachaReq) GetGachaTimes() uint32 {
	if x != nil {
		return x.GachaTimes
	}
	return 0
}

func (x *DoGachaReq) GetGachaRandom() uint32 {
	if x != nil {
		return x.GachaRandom
	}
	return 0
}

func (x *DoGachaReq) GetGachaScheduleId() uint32 {
	if x != nil {
		return x.GachaScheduleId
	}
	return 0
}

func (x *DoGachaReq) GetGachaTag() string {
	if x != nil {
		return x.GachaTag
	}
	return ""
}

var File_DoGachaReq_proto protoreflect.FileDescriptor

var file_DoGachaReq_proto_rawDesc = []byte{
	0x0a, 0x10, 0x44, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x0a, 0x44, 0x6f,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61,
	0x63, 0x68, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x61,
	0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x54, 0x61, 0x67, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49,
	0x44, 0x10, 0x85, 0x0c, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DoGachaReq_proto_rawDescOnce sync.Once
	file_DoGachaReq_proto_rawDescData = file_DoGachaReq_proto_rawDesc
)

func file_DoGachaReq_proto_rawDescGZIP() []byte {
	file_DoGachaReq_proto_rawDescOnce.Do(func() {
		file_DoGachaReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_DoGachaReq_proto_rawDescData)
	})
	return file_DoGachaReq_proto_rawDescData
}

var file_DoGachaReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DoGachaReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DoGachaReq_proto_goTypes = []interface{}{
	(DoGachaReq_CmdId)(0), // 0: proto.DoGachaReq.CmdId
	(*DoGachaReq)(nil),    // 1: proto.DoGachaReq
}
var file_DoGachaReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DoGachaReq_proto_init() }
func file_DoGachaReq_proto_init() {
	if File_DoGachaReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DoGachaReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoGachaReq); i {
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
			RawDescriptor: file_DoGachaReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DoGachaReq_proto_goTypes,
		DependencyIndexes: file_DoGachaReq_proto_depIdxs,
		EnumInfos:         file_DoGachaReq_proto_enumTypes,
		MessageInfos:      file_DoGachaReq_proto_msgTypes,
	}.Build()
	File_DoGachaReq_proto = out.File
	file_DoGachaReq_proto_rawDesc = nil
	file_DoGachaReq_proto_goTypes = nil
	file_DoGachaReq_proto_depIdxs = nil
}
