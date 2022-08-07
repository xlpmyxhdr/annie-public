// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarGainCostumeNotify.proto

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

// CmdId: 1670
// EnetChannelId: 0
// EnetIsReliable: true
type AvatarGainCostumeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostumeId uint32 `protobuf:"varint,15,opt,name=costume_id,json=costumeId,proto3" json:"costume_id,omitempty"`
}

func (x *AvatarGainCostumeNotify) Reset() {
	*x = AvatarGainCostumeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarGainCostumeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarGainCostumeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarGainCostumeNotify) ProtoMessage() {}

func (x *AvatarGainCostumeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarGainCostumeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarGainCostumeNotify.ProtoReflect.Descriptor instead.
func (*AvatarGainCostumeNotify) Descriptor() ([]byte, []int) {
	return file_AvatarGainCostumeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarGainCostumeNotify) GetCostumeId() uint32 {
	if x != nil {
		return x.CostumeId
	}
	return 0
}

var File_AvatarGainCostumeNotify_proto protoreflect.FileDescriptor

var file_AvatarGainCostumeNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x73, 0x74,
	0x75, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x17, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x47, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarGainCostumeNotify_proto_rawDescOnce sync.Once
	file_AvatarGainCostumeNotify_proto_rawDescData = file_AvatarGainCostumeNotify_proto_rawDesc
)

func file_AvatarGainCostumeNotify_proto_rawDescGZIP() []byte {
	file_AvatarGainCostumeNotify_proto_rawDescOnce.Do(func() {
		file_AvatarGainCostumeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarGainCostumeNotify_proto_rawDescData)
	})
	return file_AvatarGainCostumeNotify_proto_rawDescData
}

var file_AvatarGainCostumeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarGainCostumeNotify_proto_goTypes = []interface{}{
	(*AvatarGainCostumeNotify)(nil), // 0: proto.AvatarGainCostumeNotify
}
var file_AvatarGainCostumeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarGainCostumeNotify_proto_init() }
func file_AvatarGainCostumeNotify_proto_init() {
	if File_AvatarGainCostumeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarGainCostumeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarGainCostumeNotify); i {
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
			RawDescriptor: file_AvatarGainCostumeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarGainCostumeNotify_proto_goTypes,
		DependencyIndexes: file_AvatarGainCostumeNotify_proto_depIdxs,
		MessageInfos:      file_AvatarGainCostumeNotify_proto_msgTypes,
	}.Build()
	File_AvatarGainCostumeNotify_proto = out.File
	file_AvatarGainCostumeNotify_proto_rawDesc = nil
	file_AvatarGainCostumeNotify_proto_goTypes = nil
	file_AvatarGainCostumeNotify_proto_depIdxs = nil
}
