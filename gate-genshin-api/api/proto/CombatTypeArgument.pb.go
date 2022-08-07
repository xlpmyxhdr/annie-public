// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: CombatTypeArgument.proto

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

type CombatTypeArgument int32

const (
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_NONE                       CombatTypeArgument = 0
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_EVT_BEING_HIT              CombatTypeArgument = 1
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_ANIMATOR_STATE_CHANGED     CombatTypeArgument = 2
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_FACE_TO_DIR                CombatTypeArgument = 3
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_SET_ATTACK_TARGET          CombatTypeArgument = 4
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_RUSH_MOVE                  CombatTypeArgument = 5
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_ANIMATOR_PARAMETER_CHANGED CombatTypeArgument = 6
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_ENTITY_MOVE                CombatTypeArgument = 7
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_SYNC_ENTITY_POSITION       CombatTypeArgument = 8
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_STEER_MOTION_INFO          CombatTypeArgument = 9
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_FORCE_SET_POS_INFO         CombatTypeArgument = 10
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_COMPENSATE_POS_DIFF        CombatTypeArgument = 11
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_MONSTER_DO_BLINK           CombatTypeArgument = 12
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_FIXED_RUSH_MOVE            CombatTypeArgument = 13
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_SYNC_TRANSFORM             CombatTypeArgument = 14
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_LIGHT_CORE_MOVE            CombatTypeArgument = 15
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_KPDNFKCMKPG                CombatTypeArgument = 16
	CombatTypeArgument_COMBAT_TYPE_ARGUMENT_KPLOMOIALGF                CombatTypeArgument = 17
)

// Enum value maps for CombatTypeArgument.
var (
	CombatTypeArgument_name = map[int32]string{
		0:  "COMBAT_TYPE_ARGUMENT_NONE",
		1:  "COMBAT_TYPE_ARGUMENT_EVT_BEING_HIT",
		2:  "COMBAT_TYPE_ARGUMENT_ANIMATOR_STATE_CHANGED",
		3:  "COMBAT_TYPE_ARGUMENT_FACE_TO_DIR",
		4:  "COMBAT_TYPE_ARGUMENT_SET_ATTACK_TARGET",
		5:  "COMBAT_TYPE_ARGUMENT_RUSH_MOVE",
		6:  "COMBAT_TYPE_ARGUMENT_ANIMATOR_PARAMETER_CHANGED",
		7:  "COMBAT_TYPE_ARGUMENT_ENTITY_MOVE",
		8:  "COMBAT_TYPE_ARGUMENT_SYNC_ENTITY_POSITION",
		9:  "COMBAT_TYPE_ARGUMENT_STEER_MOTION_INFO",
		10: "COMBAT_TYPE_ARGUMENT_FORCE_SET_POS_INFO",
		11: "COMBAT_TYPE_ARGUMENT_COMPENSATE_POS_DIFF",
		12: "COMBAT_TYPE_ARGUMENT_MONSTER_DO_BLINK",
		13: "COMBAT_TYPE_ARGUMENT_FIXED_RUSH_MOVE",
		14: "COMBAT_TYPE_ARGUMENT_SYNC_TRANSFORM",
		15: "COMBAT_TYPE_ARGUMENT_LIGHT_CORE_MOVE",
		16: "COMBAT_TYPE_ARGUMENT_KPDNFKCMKPG",
		17: "COMBAT_TYPE_ARGUMENT_KPLOMOIALGF",
	}
	CombatTypeArgument_value = map[string]int32{
		"COMBAT_TYPE_ARGUMENT_NONE":                       0,
		"COMBAT_TYPE_ARGUMENT_EVT_BEING_HIT":              1,
		"COMBAT_TYPE_ARGUMENT_ANIMATOR_STATE_CHANGED":     2,
		"COMBAT_TYPE_ARGUMENT_FACE_TO_DIR":                3,
		"COMBAT_TYPE_ARGUMENT_SET_ATTACK_TARGET":          4,
		"COMBAT_TYPE_ARGUMENT_RUSH_MOVE":                  5,
		"COMBAT_TYPE_ARGUMENT_ANIMATOR_PARAMETER_CHANGED": 6,
		"COMBAT_TYPE_ARGUMENT_ENTITY_MOVE":                7,
		"COMBAT_TYPE_ARGUMENT_SYNC_ENTITY_POSITION":       8,
		"COMBAT_TYPE_ARGUMENT_STEER_MOTION_INFO":          9,
		"COMBAT_TYPE_ARGUMENT_FORCE_SET_POS_INFO":         10,
		"COMBAT_TYPE_ARGUMENT_COMPENSATE_POS_DIFF":        11,
		"COMBAT_TYPE_ARGUMENT_MONSTER_DO_BLINK":           12,
		"COMBAT_TYPE_ARGUMENT_FIXED_RUSH_MOVE":            13,
		"COMBAT_TYPE_ARGUMENT_SYNC_TRANSFORM":             14,
		"COMBAT_TYPE_ARGUMENT_LIGHT_CORE_MOVE":            15,
		"COMBAT_TYPE_ARGUMENT_KPDNFKCMKPG":                16,
		"COMBAT_TYPE_ARGUMENT_KPLOMOIALGF":                17,
	}
)

func (x CombatTypeArgument) Enum() *CombatTypeArgument {
	p := new(CombatTypeArgument)
	*p = x
	return p
}

func (x CombatTypeArgument) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CombatTypeArgument) Descriptor() protoreflect.EnumDescriptor {
	return file_CombatTypeArgument_proto_enumTypes[0].Descriptor()
}

func (CombatTypeArgument) Type() protoreflect.EnumType {
	return &file_CombatTypeArgument_proto_enumTypes[0]
}

func (x CombatTypeArgument) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CombatTypeArgument.Descriptor instead.
func (CombatTypeArgument) EnumDescriptor() ([]byte, []int) {
	return file_CombatTypeArgument_proto_rawDescGZIP(), []int{0}
}

var File_CombatTypeArgument_proto protoreflect.FileDescriptor

var file_CombatTypeArgument_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x41, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x87, 0x06, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x42,
	0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x4d, 0x42, 0x41,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x45, 0x56, 0x54, 0x5f, 0x42, 0x45, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x01, 0x12,
	0x2f, 0x0a, 0x2b, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4e, 0x49, 0x4d, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x4f,
	0x5f, 0x44, 0x49, 0x52, 0x10, 0x03, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x45, 0x54, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54,
	0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x55, 0x53, 0x48, 0x5f,
	0x4d, 0x4f, 0x56, 0x45, 0x10, 0x05, 0x12, 0x33, 0x0a, 0x2f, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x4e, 0x49, 0x4d, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45,
	0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x06, 0x12, 0x24, 0x0a, 0x20, 0x43,
	0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10,
	0x07, 0x12, 0x2d, 0x0a, 0x29, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08,
	0x12, 0x2a, 0x0a, 0x26, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x45, 0x45, 0x52, 0x5f, 0x4d,
	0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x09, 0x12, 0x2b, 0x0a, 0x27,
	0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x50,
	0x4f, 0x53, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0a, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x4f, 0x4d,
	0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x45, 0x4e, 0x53, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4f, 0x53,
	0x5f, 0x44, 0x49, 0x46, 0x46, 0x10, 0x0b, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x4f, 0x4d, 0x42, 0x41,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x5f, 0x42, 0x4c, 0x49, 0x4e, 0x4b,
	0x10, 0x0c, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x58, 0x45, 0x44,
	0x5f, 0x52, 0x55, 0x53, 0x48, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x0d, 0x12, 0x27, 0x0a, 0x23,
	0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46,
	0x4f, 0x52, 0x4d, 0x10, 0x0e, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x0f, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x50, 0x44, 0x4e, 0x46, 0x4b, 0x43, 0x4d,
	0x4b, 0x50, 0x47, 0x10, 0x10, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4d, 0x42, 0x41, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x50,
	0x4c, 0x4f, 0x4d, 0x4f, 0x49, 0x41, 0x4c, 0x47, 0x46, 0x10, 0x11, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CombatTypeArgument_proto_rawDescOnce sync.Once
	file_CombatTypeArgument_proto_rawDescData = file_CombatTypeArgument_proto_rawDesc
)

func file_CombatTypeArgument_proto_rawDescGZIP() []byte {
	file_CombatTypeArgument_proto_rawDescOnce.Do(func() {
		file_CombatTypeArgument_proto_rawDescData = protoimpl.X.CompressGZIP(file_CombatTypeArgument_proto_rawDescData)
	})
	return file_CombatTypeArgument_proto_rawDescData
}

var file_CombatTypeArgument_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CombatTypeArgument_proto_goTypes = []interface{}{
	(CombatTypeArgument)(0), // 0: proto.CombatTypeArgument
}
var file_CombatTypeArgument_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CombatTypeArgument_proto_init() }
func file_CombatTypeArgument_proto_init() {
	if File_CombatTypeArgument_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CombatTypeArgument_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CombatTypeArgument_proto_goTypes,
		DependencyIndexes: file_CombatTypeArgument_proto_depIdxs,
		EnumInfos:         file_CombatTypeArgument_proto_enumTypes,
	}.Build()
	File_CombatTypeArgument_proto = out.File
	file_CombatTypeArgument_proto_rawDesc = nil
	file_CombatTypeArgument_proto_goTypes = nil
	file_CombatTypeArgument_proto_depIdxs = nil
}
