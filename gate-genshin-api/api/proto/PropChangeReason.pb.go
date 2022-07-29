// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PropChangeReason.proto

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

type PropChangeReason int32

const (
	PropChangeReason_PROP_CHANGE_NONE                      PropChangeReason = 0
	PropChangeReason_PROP_CHANGE_STATUE_RECOVER            PropChangeReason = 1
	PropChangeReason_PROP_CHANGE_ENERGY_BALL               PropChangeReason = 2
	PropChangeReason_PROP_CHANGE_ABILITY                   PropChangeReason = 3
	PropChangeReason_PROP_CHANGE_LEVELUP                   PropChangeReason = 4
	PropChangeReason_PROP_CHANGE_ITEM                      PropChangeReason = 5
	PropChangeReason_PROP_CHANGE_AVATAR_CARD               PropChangeReason = 6
	PropChangeReason_PROP_CHANGE_CITY_LEVELUP              PropChangeReason = 7
	PropChangeReason_PROP_CHANGE_AVATAR_UPGRADE            PropChangeReason = 8
	PropChangeReason_PROP_CHANGE_AVATAR_PROMOTE            PropChangeReason = 9
	PropChangeReason_PROP_CHANGE_PLAYER_ADD_EXP            PropChangeReason = 10
	PropChangeReason_PROP_CHANGE_FINISH_QUEST              PropChangeReason = 11
	PropChangeReason_PROP_CHANGE_GM                        PropChangeReason = 12
	PropChangeReason_PROP_CHANGE_MANUAL_ADJUST_WORLD_LEVEL PropChangeReason = 13
)

// Enum value maps for PropChangeReason.
var (
	PropChangeReason_name = map[int32]string{
		0:  "PROP_CHANGE_NONE",
		1:  "PROP_CHANGE_STATUE_RECOVER",
		2:  "PROP_CHANGE_ENERGY_BALL",
		3:  "PROP_CHANGE_ABILITY",
		4:  "PROP_CHANGE_LEVELUP",
		5:  "PROP_CHANGE_ITEM",
		6:  "PROP_CHANGE_AVATAR_CARD",
		7:  "PROP_CHANGE_CITY_LEVELUP",
		8:  "PROP_CHANGE_AVATAR_UPGRADE",
		9:  "PROP_CHANGE_AVATAR_PROMOTE",
		10: "PROP_CHANGE_PLAYER_ADD_EXP",
		11: "PROP_CHANGE_FINISH_QUEST",
		12: "PROP_CHANGE_GM",
		13: "PROP_CHANGE_MANUAL_ADJUST_WORLD_LEVEL",
	}
	PropChangeReason_value = map[string]int32{
		"PROP_CHANGE_NONE":                      0,
		"PROP_CHANGE_STATUE_RECOVER":            1,
		"PROP_CHANGE_ENERGY_BALL":               2,
		"PROP_CHANGE_ABILITY":                   3,
		"PROP_CHANGE_LEVELUP":                   4,
		"PROP_CHANGE_ITEM":                      5,
		"PROP_CHANGE_AVATAR_CARD":               6,
		"PROP_CHANGE_CITY_LEVELUP":              7,
		"PROP_CHANGE_AVATAR_UPGRADE":            8,
		"PROP_CHANGE_AVATAR_PROMOTE":            9,
		"PROP_CHANGE_PLAYER_ADD_EXP":            10,
		"PROP_CHANGE_FINISH_QUEST":              11,
		"PROP_CHANGE_GM":                        12,
		"PROP_CHANGE_MANUAL_ADJUST_WORLD_LEVEL": 13,
	}
)

func (x PropChangeReason) Enum() *PropChangeReason {
	p := new(PropChangeReason)
	*p = x
	return p
}

func (x PropChangeReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PropChangeReason) Descriptor() protoreflect.EnumDescriptor {
	return file_PropChangeReason_proto_enumTypes[0].Descriptor()
}

func (PropChangeReason) Type() protoreflect.EnumType {
	return &file_PropChangeReason_proto_enumTypes[0]
}

func (x PropChangeReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PropChangeReason.Descriptor instead.
func (PropChangeReason) EnumDescriptor() ([]byte, []int) {
	return file_PropChangeReason_proto_rawDescGZIP(), []int{0}
}

var File_PropChangeReason_proto protoreflect.FileDescriptor

var file_PropChangeReason_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xa5, 0x03, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52,
	0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x45,
	0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52,
	0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59,
	0x5f, 0x42, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x50, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x03,
	0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x55, 0x50, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f,
	0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x05, 0x12,
	0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41,
	0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18,
	0x50, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x43, 0x49, 0x54, 0x59,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x55, 0x50, 0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52,
	0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52,
	0x5f, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52,
	0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52,
	0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52,
	0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x58, 0x50, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52,
	0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x50,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x47, 0x4d, 0x10, 0x0c, 0x12, 0x29, 0x0a, 0x25,
	0x50, 0x52, 0x4f, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x55,
	0x41, 0x4c, 0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x0d, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PropChangeReason_proto_rawDescOnce sync.Once
	file_PropChangeReason_proto_rawDescData = file_PropChangeReason_proto_rawDesc
)

func file_PropChangeReason_proto_rawDescGZIP() []byte {
	file_PropChangeReason_proto_rawDescOnce.Do(func() {
		file_PropChangeReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_PropChangeReason_proto_rawDescData)
	})
	return file_PropChangeReason_proto_rawDescData
}

var file_PropChangeReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PropChangeReason_proto_goTypes = []interface{}{
	(PropChangeReason)(0), // 0: proto.PropChangeReason
}
var file_PropChangeReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PropChangeReason_proto_init() }
func file_PropChangeReason_proto_init() {
	if File_PropChangeReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PropChangeReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PropChangeReason_proto_goTypes,
		DependencyIndexes: file_PropChangeReason_proto_depIdxs,
		EnumInfos:         file_PropChangeReason_proto_enumTypes,
	}.Build()
	File_PropChangeReason_proto = out.File
	file_PropChangeReason_proto_rawDesc = nil
	file_PropChangeReason_proto_goTypes = nil
	file_PropChangeReason_proto_depIdxs = nil
}
