// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: EvtBeingHitsCombineNotify.proto

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

type EvtBeingHitsCombineNotify_CmdId int32

const (
	EvtBeingHitsCombineNotify_NONE             EvtBeingHitsCombineNotify_CmdId = 0
	EvtBeingHitsCombineNotify_ENET_CHANNEL_ID  EvtBeingHitsCombineNotify_CmdId = 0
	EvtBeingHitsCombineNotify_ENET_IS_RELIABLE EvtBeingHitsCombineNotify_CmdId = 1
	EvtBeingHitsCombineNotify_IS_ALLOW_CLIENT  EvtBeingHitsCombineNotify_CmdId = 1
	EvtBeingHitsCombineNotify_CMD_ID           EvtBeingHitsCombineNotify_CmdId = 336
)

// Enum value maps for EvtBeingHitsCombineNotify_CmdId.
var (
	EvtBeingHitsCombineNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		336: "CMD_ID",
	}
	EvtBeingHitsCombineNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           336,
	}
)

func (x EvtBeingHitsCombineNotify_CmdId) Enum() *EvtBeingHitsCombineNotify_CmdId {
	p := new(EvtBeingHitsCombineNotify_CmdId)
	*p = x
	return p
}

func (x EvtBeingHitsCombineNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvtBeingHitsCombineNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EvtBeingHitsCombineNotify_proto_enumTypes[0].Descriptor()
}

func (EvtBeingHitsCombineNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EvtBeingHitsCombineNotify_proto_enumTypes[0]
}

func (x EvtBeingHitsCombineNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvtBeingHitsCombineNotify_CmdId.Descriptor instead.
func (EvtBeingHitsCombineNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EvtBeingHitsCombineNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EvtBeingHitsCombineNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType         ForwardType        `protobuf:"varint,1,opt,name=forward_type,json=forwardType,proto3,enum=proto.ForwardType" json:"forward_type,omitempty"`
	EvtBeingHitInfoList []*EvtBeingHitInfo `protobuf:"bytes,14,rep,name=evt_being_hit_info_list,json=evtBeingHitInfoList,proto3" json:"evt_being_hit_info_list,omitempty"`
}

func (x *EvtBeingHitsCombineNotify) Reset() {
	*x = EvtBeingHitsCombineNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtBeingHitsCombineNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtBeingHitsCombineNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtBeingHitsCombineNotify) ProtoMessage() {}

func (x *EvtBeingHitsCombineNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtBeingHitsCombineNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtBeingHitsCombineNotify.ProtoReflect.Descriptor instead.
func (*EvtBeingHitsCombineNotify) Descriptor() ([]byte, []int) {
	return file_EvtBeingHitsCombineNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtBeingHitsCombineNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

func (x *EvtBeingHitsCombineNotify) GetEvtBeingHitInfoList() []*EvtBeingHitInfo {
	if x != nil {
		return x.EvtBeingHitInfoList
	}
	return nil
}

var File_EvtBeingHitsCombineNotify_proto protoreflect.FileDescriptor

var file_EvtBeingHitsCombineNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x74, 0x73, 0x43, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x45, 0x76, 0x74,
	0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x19, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x35, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x17, 0x65, 0x76, 0x74, 0x5f, 0x62,
	0x65, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x13, 0x65, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xd0, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtBeingHitsCombineNotify_proto_rawDescOnce sync.Once
	file_EvtBeingHitsCombineNotify_proto_rawDescData = file_EvtBeingHitsCombineNotify_proto_rawDesc
)

func file_EvtBeingHitsCombineNotify_proto_rawDescGZIP() []byte {
	file_EvtBeingHitsCombineNotify_proto_rawDescOnce.Do(func() {
		file_EvtBeingHitsCombineNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtBeingHitsCombineNotify_proto_rawDescData)
	})
	return file_EvtBeingHitsCombineNotify_proto_rawDescData
}

var file_EvtBeingHitsCombineNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EvtBeingHitsCombineNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtBeingHitsCombineNotify_proto_goTypes = []interface{}{
	(EvtBeingHitsCombineNotify_CmdId)(0), // 0: proto.EvtBeingHitsCombineNotify.CmdId
	(*EvtBeingHitsCombineNotify)(nil),    // 1: proto.EvtBeingHitsCombineNotify
	(ForwardType)(0),                     // 2: proto.ForwardType
	(*EvtBeingHitInfo)(nil),              // 3: proto.EvtBeingHitInfo
}
var file_EvtBeingHitsCombineNotify_proto_depIdxs = []int32{
	2, // 0: proto.EvtBeingHitsCombineNotify.forward_type:type_name -> proto.ForwardType
	3, // 1: proto.EvtBeingHitsCombineNotify.evt_being_hit_info_list:type_name -> proto.EvtBeingHitInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtBeingHitsCombineNotify_proto_init() }
func file_EvtBeingHitsCombineNotify_proto_init() {
	if File_EvtBeingHitsCombineNotify_proto != nil {
		return
	}
	file_ForwardType_proto_init()
	file_EvtBeingHitInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtBeingHitsCombineNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtBeingHitsCombineNotify); i {
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
			RawDescriptor: file_EvtBeingHitsCombineNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtBeingHitsCombineNotify_proto_goTypes,
		DependencyIndexes: file_EvtBeingHitsCombineNotify_proto_depIdxs,
		EnumInfos:         file_EvtBeingHitsCombineNotify_proto_enumTypes,
		MessageInfos:      file_EvtBeingHitsCombineNotify_proto_msgTypes,
	}.Build()
	File_EvtBeingHitsCombineNotify_proto = out.File
	file_EvtBeingHitsCombineNotify_proto_rawDesc = nil
	file_EvtBeingHitsCombineNotify_proto_goTypes = nil
	file_EvtBeingHitsCombineNotify_proto_depIdxs = nil
}
