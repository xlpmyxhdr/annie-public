// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: StoreItemDelNotify.proto

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

type StoreItemDelNotify_CmdId int32

const (
	StoreItemDelNotify_NONE             StoreItemDelNotify_CmdId = 0
	StoreItemDelNotify_ENET_CHANNEL_ID  StoreItemDelNotify_CmdId = 0
	StoreItemDelNotify_ENET_IS_RELIABLE StoreItemDelNotify_CmdId = 1
	StoreItemDelNotify_CMD_ID           StoreItemDelNotify_CmdId = 684
)

// Enum value maps for StoreItemDelNotify_CmdId.
var (
	StoreItemDelNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		684: "CMD_ID",
	}
	StoreItemDelNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           684,
	}
)

func (x StoreItemDelNotify_CmdId) Enum() *StoreItemDelNotify_CmdId {
	p := new(StoreItemDelNotify_CmdId)
	*p = x
	return p
}

func (x StoreItemDelNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoreItemDelNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_StoreItemDelNotify_proto_enumTypes[0].Descriptor()
}

func (StoreItemDelNotify_CmdId) Type() protoreflect.EnumType {
	return &file_StoreItemDelNotify_proto_enumTypes[0]
}

func (x StoreItemDelNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoreItemDelNotify_CmdId.Descriptor instead.
func (StoreItemDelNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_StoreItemDelNotify_proto_rawDescGZIP(), []int{0, 0}
}

type StoreItemDelNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreType StoreType `protobuf:"varint,1,opt,name=store_type,json=storeType,proto3,enum=proto.StoreType" json:"store_type,omitempty"`
	GuidList  []uint64  `protobuf:"varint,2,rep,packed,name=guid_list,json=guidList,proto3" json:"guid_list,omitempty"`
}

func (x *StoreItemDelNotify) Reset() {
	*x = StoreItemDelNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StoreItemDelNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreItemDelNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreItemDelNotify) ProtoMessage() {}

func (x *StoreItemDelNotify) ProtoReflect() protoreflect.Message {
	mi := &file_StoreItemDelNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreItemDelNotify.ProtoReflect.Descriptor instead.
func (*StoreItemDelNotify) Descriptor() ([]byte, []int) {
	return file_StoreItemDelNotify_proto_rawDescGZIP(), []int{0}
}

func (x *StoreItemDelNotify) GetStoreType() StoreType {
	if x != nil {
		return x.StoreType
	}
	return StoreType_STORE_NONE
}

func (x *StoreItemDelNotify) GetGuidList() []uint64 {
	if x != nil {
		return x.GuidList
	}
	return nil
}

var File_StoreItemDelNotify_proto protoreflect.FileDescriptor

var file_StoreItemDelNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2f, 0x0a, 0x0a, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x75,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x67,
	0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xac, 0x05, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StoreItemDelNotify_proto_rawDescOnce sync.Once
	file_StoreItemDelNotify_proto_rawDescData = file_StoreItemDelNotify_proto_rawDesc
)

func file_StoreItemDelNotify_proto_rawDescGZIP() []byte {
	file_StoreItemDelNotify_proto_rawDescOnce.Do(func() {
		file_StoreItemDelNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_StoreItemDelNotify_proto_rawDescData)
	})
	return file_StoreItemDelNotify_proto_rawDescData
}

var file_StoreItemDelNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_StoreItemDelNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StoreItemDelNotify_proto_goTypes = []interface{}{
	(StoreItemDelNotify_CmdId)(0), // 0: proto.StoreItemDelNotify.CmdId
	(*StoreItemDelNotify)(nil),    // 1: proto.StoreItemDelNotify
	(StoreType)(0),                // 2: proto.StoreType
}
var file_StoreItemDelNotify_proto_depIdxs = []int32{
	2, // 0: proto.StoreItemDelNotify.store_type:type_name -> proto.StoreType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_StoreItemDelNotify_proto_init() }
func file_StoreItemDelNotify_proto_init() {
	if File_StoreItemDelNotify_proto != nil {
		return
	}
	file_StoreType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StoreItemDelNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreItemDelNotify); i {
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
			RawDescriptor: file_StoreItemDelNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StoreItemDelNotify_proto_goTypes,
		DependencyIndexes: file_StoreItemDelNotify_proto_depIdxs,
		EnumInfos:         file_StoreItemDelNotify_proto_enumTypes,
		MessageInfos:      file_StoreItemDelNotify_proto_msgTypes,
	}.Build()
	File_StoreItemDelNotify_proto = out.File
	file_StoreItemDelNotify_proto_rawDesc = nil
	file_StoreItemDelNotify_proto_goTypes = nil
	file_StoreItemDelNotify_proto_depIdxs = nil
}
