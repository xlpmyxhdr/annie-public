// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChallengeDataNotify.proto

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

type ChallengeDataNotify_CmdId int32

const (
	ChallengeDataNotify_NONE             ChallengeDataNotify_CmdId = 0
	ChallengeDataNotify_ENET_CHANNEL_ID  ChallengeDataNotify_CmdId = 0
	ChallengeDataNotify_ENET_IS_RELIABLE ChallengeDataNotify_CmdId = 1
	ChallengeDataNotify_CMD_ID           ChallengeDataNotify_CmdId = 963
)

// Enum value maps for ChallengeDataNotify_CmdId.
var (
	ChallengeDataNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		963: "CMD_ID",
	}
	ChallengeDataNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           963,
	}
)

func (x ChallengeDataNotify_CmdId) Enum() *ChallengeDataNotify_CmdId {
	p := new(ChallengeDataNotify_CmdId)
	*p = x
	return p
}

func (x ChallengeDataNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChallengeDataNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ChallengeDataNotify_proto_enumTypes[0].Descriptor()
}

func (ChallengeDataNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ChallengeDataNotify_proto_enumTypes[0]
}

func (x ChallengeDataNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChallengeDataNotify_CmdId.Descriptor instead.
func (ChallengeDataNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ChallengeDataNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ChallengeDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeIndex uint32 `protobuf:"varint,8,opt,name=challenge_index,json=challengeIndex,proto3" json:"challenge_index,omitempty"`
	ParamIndex     uint32 `protobuf:"varint,9,opt,name=param_index,json=paramIndex,proto3" json:"param_index,omitempty"`
	Value          uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ChallengeDataNotify) Reset() {
	*x = ChallengeDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeDataNotify) ProtoMessage() {}

func (x *ChallengeDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeDataNotify.ProtoReflect.Descriptor instead.
func (*ChallengeDataNotify) Descriptor() ([]byte, []int) {
	return file_ChallengeDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeDataNotify) GetChallengeIndex() uint32 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

func (x *ChallengeDataNotify) GetParamIndex() uint32 {
	if x != nil {
		return x.ParamIndex
	}
	return 0
}

func (x *ChallengeDataNotify) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_ChallengeDataNotify_proto protoreflect.FileDescriptor

var file_ChallengeDataNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d,
	0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xc3, 0x07, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeDataNotify_proto_rawDescOnce sync.Once
	file_ChallengeDataNotify_proto_rawDescData = file_ChallengeDataNotify_proto_rawDesc
)

func file_ChallengeDataNotify_proto_rawDescGZIP() []byte {
	file_ChallengeDataNotify_proto_rawDescOnce.Do(func() {
		file_ChallengeDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeDataNotify_proto_rawDescData)
	})
	return file_ChallengeDataNotify_proto_rawDescData
}

var file_ChallengeDataNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChallengeDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeDataNotify_proto_goTypes = []interface{}{
	(ChallengeDataNotify_CmdId)(0), // 0: proto.ChallengeDataNotify.CmdId
	(*ChallengeDataNotify)(nil),    // 1: proto.ChallengeDataNotify
}
var file_ChallengeDataNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChallengeDataNotify_proto_init() }
func file_ChallengeDataNotify_proto_init() {
	if File_ChallengeDataNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChallengeDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeDataNotify); i {
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
			RawDescriptor: file_ChallengeDataNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeDataNotify_proto_goTypes,
		DependencyIndexes: file_ChallengeDataNotify_proto_depIdxs,
		EnumInfos:         file_ChallengeDataNotify_proto_enumTypes,
		MessageInfos:      file_ChallengeDataNotify_proto_msgTypes,
	}.Build()
	File_ChallengeDataNotify_proto = out.File
	file_ChallengeDataNotify_proto_rawDesc = nil
	file_ChallengeDataNotify_proto_goTypes = nil
	file_ChallengeDataNotify_proto_depIdxs = nil
}
