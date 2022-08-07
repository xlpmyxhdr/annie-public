// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarExpeditionState.proto

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

type AvatarExpeditionState int32

const (
	AvatarExpeditionState_AVATAR_EXPEDITION_STATE_NONE                 AvatarExpeditionState = 0
	AvatarExpeditionState_AVATAR_EXPEDITION_STATE_DOING                AvatarExpeditionState = 1
	AvatarExpeditionState_AVATAR_EXPEDITION_STATE_FINISH_WAIT_REWARD   AvatarExpeditionState = 2
	AvatarExpeditionState_AVATAR_EXPEDITION_STATE_CALLBACK_WAIT_REWARD AvatarExpeditionState = 3
	AvatarExpeditionState_AVATAR_EXPEDITION_STATE_LOCKED               AvatarExpeditionState = 4
)

// Enum value maps for AvatarExpeditionState.
var (
	AvatarExpeditionState_name = map[int32]string{
		0: "AVATAR_EXPEDITION_STATE_NONE",
		1: "AVATAR_EXPEDITION_STATE_DOING",
		2: "AVATAR_EXPEDITION_STATE_FINISH_WAIT_REWARD",
		3: "AVATAR_EXPEDITION_STATE_CALLBACK_WAIT_REWARD",
		4: "AVATAR_EXPEDITION_STATE_LOCKED",
	}
	AvatarExpeditionState_value = map[string]int32{
		"AVATAR_EXPEDITION_STATE_NONE":                 0,
		"AVATAR_EXPEDITION_STATE_DOING":                1,
		"AVATAR_EXPEDITION_STATE_FINISH_WAIT_REWARD":   2,
		"AVATAR_EXPEDITION_STATE_CALLBACK_WAIT_REWARD": 3,
		"AVATAR_EXPEDITION_STATE_LOCKED":               4,
	}
)

func (x AvatarExpeditionState) Enum() *AvatarExpeditionState {
	p := new(AvatarExpeditionState)
	*p = x
	return p
}

func (x AvatarExpeditionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarExpeditionState) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarExpeditionState_proto_enumTypes[0].Descriptor()
}

func (AvatarExpeditionState) Type() protoreflect.EnumType {
	return &file_AvatarExpeditionState_proto_enumTypes[0]
}

func (x AvatarExpeditionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarExpeditionState.Descriptor instead.
func (AvatarExpeditionState) EnumDescriptor() ([]byte, []int) {
	return file_AvatarExpeditionState_proto_rawDescGZIP(), []int{0}
}

var File_AvatarExpeditionState_proto protoreflect.FileDescriptor

var file_AvatarExpeditionState_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe2, 0x01, 0x0a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45,
	0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x1c, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x21, 0x0a, 0x1d, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x44,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x2e, 0x0a, 0x2a, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x45, 0x58,
	0x50, 0x45, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46,
	0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52,
	0x44, 0x10, 0x02, 0x12, 0x30, 0x0a, 0x2c, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x45, 0x58,
	0x50, 0x45, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x57,
	0x41, 0x52, 0x44, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f,
	0x45, 0x58, 0x50, 0x45, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x04, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExpeditionState_proto_rawDescOnce sync.Once
	file_AvatarExpeditionState_proto_rawDescData = file_AvatarExpeditionState_proto_rawDesc
)

func file_AvatarExpeditionState_proto_rawDescGZIP() []byte {
	file_AvatarExpeditionState_proto_rawDescOnce.Do(func() {
		file_AvatarExpeditionState_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExpeditionState_proto_rawDescData)
	})
	return file_AvatarExpeditionState_proto_rawDescData
}

var file_AvatarExpeditionState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarExpeditionState_proto_goTypes = []interface{}{
	(AvatarExpeditionState)(0), // 0: proto.AvatarExpeditionState
}
var file_AvatarExpeditionState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarExpeditionState_proto_init() }
func file_AvatarExpeditionState_proto_init() {
	if File_AvatarExpeditionState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AvatarExpeditionState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExpeditionState_proto_goTypes,
		DependencyIndexes: file_AvatarExpeditionState_proto_depIdxs,
		EnumInfos:         file_AvatarExpeditionState_proto_enumTypes,
	}.Build()
	File_AvatarExpeditionState_proto = out.File
	file_AvatarExpeditionState_proto_rawDesc = nil
	file_AvatarExpeditionState_proto_goTypes = nil
	file_AvatarExpeditionState_proto_depIdxs = nil
}
