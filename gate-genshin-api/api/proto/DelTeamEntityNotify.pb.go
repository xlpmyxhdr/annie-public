// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: DelTeamEntityNotify.proto

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

type DelTeamEntityNotify_CmdId int32

const (
	DelTeamEntityNotify_NONE             DelTeamEntityNotify_CmdId = 0
	DelTeamEntityNotify_ENET_CHANNEL_ID  DelTeamEntityNotify_CmdId = 0
	DelTeamEntityNotify_ENET_IS_RELIABLE DelTeamEntityNotify_CmdId = 1
	DelTeamEntityNotify_CMD_ID           DelTeamEntityNotify_CmdId = 302
)

// Enum value maps for DelTeamEntityNotify_CmdId.
var (
	DelTeamEntityNotify_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		302: "CMD_ID",
	}
	DelTeamEntityNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           302,
	}
)

func (x DelTeamEntityNotify_CmdId) Enum() *DelTeamEntityNotify_CmdId {
	p := new(DelTeamEntityNotify_CmdId)
	*p = x
	return p
}

func (x DelTeamEntityNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DelTeamEntityNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DelTeamEntityNotify_proto_enumTypes[0].Descriptor()
}

func (DelTeamEntityNotify_CmdId) Type() protoreflect.EnumType {
	return &file_DelTeamEntityNotify_proto_enumTypes[0]
}

func (x DelTeamEntityNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DelTeamEntityNotify_CmdId.Descriptor instead.
func (DelTeamEntityNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DelTeamEntityNotify_proto_rawDescGZIP(), []int{0, 0}
}

type DelTeamEntityNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId         uint32   `protobuf:"varint,8,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	DelEntityIdList []uint32 `protobuf:"varint,15,rep,packed,name=del_entity_id_list,json=delEntityIdList,proto3" json:"del_entity_id_list,omitempty"`
}

func (x *DelTeamEntityNotify) Reset() {
	*x = DelTeamEntityNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DelTeamEntityNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelTeamEntityNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelTeamEntityNotify) ProtoMessage() {}

func (x *DelTeamEntityNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DelTeamEntityNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelTeamEntityNotify.ProtoReflect.Descriptor instead.
func (*DelTeamEntityNotify) Descriptor() ([]byte, []int) {
	return file_DelTeamEntityNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DelTeamEntityNotify) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *DelTeamEntityNotify) GetDelEntityIdList() []uint32 {
	if x != nil {
		return x.DelEntityIdList
	}
	return nil
}

var File_DelTeamEntityNotify_proto protoreflect.FileDescriptor

var file_DelTeamEntityNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xae, 0x02, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DelTeamEntityNotify_proto_rawDescOnce sync.Once
	file_DelTeamEntityNotify_proto_rawDescData = file_DelTeamEntityNotify_proto_rawDesc
)

func file_DelTeamEntityNotify_proto_rawDescGZIP() []byte {
	file_DelTeamEntityNotify_proto_rawDescOnce.Do(func() {
		file_DelTeamEntityNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DelTeamEntityNotify_proto_rawDescData)
	})
	return file_DelTeamEntityNotify_proto_rawDescData
}

var file_DelTeamEntityNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DelTeamEntityNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DelTeamEntityNotify_proto_goTypes = []interface{}{
	(DelTeamEntityNotify_CmdId)(0), // 0: proto.DelTeamEntityNotify.CmdId
	(*DelTeamEntityNotify)(nil),    // 1: proto.DelTeamEntityNotify
}
var file_DelTeamEntityNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DelTeamEntityNotify_proto_init() }
func file_DelTeamEntityNotify_proto_init() {
	if File_DelTeamEntityNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DelTeamEntityNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelTeamEntityNotify); i {
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
			RawDescriptor: file_DelTeamEntityNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DelTeamEntityNotify_proto_goTypes,
		DependencyIndexes: file_DelTeamEntityNotify_proto_depIdxs,
		EnumInfos:         file_DelTeamEntityNotify_proto_enumTypes,
		MessageInfos:      file_DelTeamEntityNotify_proto_msgTypes,
	}.Build()
	File_DelTeamEntityNotify_proto = out.File
	file_DelTeamEntityNotify_proto_rawDesc = nil
	file_DelTeamEntityNotify_proto_goTypes = nil
	file_DelTeamEntityNotify_proto_depIdxs = nil
}
