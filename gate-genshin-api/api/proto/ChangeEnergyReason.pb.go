// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChangeEnergyReason.proto

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

type ChangeEnergyReason int32

const (
	ChangeEnergyReason_CHANGE_ENERGY_REASON_NONE        ChangeEnergyReason = 0
	ChangeEnergyReason_CHANGE_ENERGY_REASON_SKILL_START ChangeEnergyReason = 1
)

// Enum value maps for ChangeEnergyReason.
var (
	ChangeEnergyReason_name = map[int32]string{
		0: "CHANGE_ENERGY_REASON_NONE",
		1: "CHANGE_ENERGY_REASON_SKILL_START",
	}
	ChangeEnergyReason_value = map[string]int32{
		"CHANGE_ENERGY_REASON_NONE":        0,
		"CHANGE_ENERGY_REASON_SKILL_START": 1,
	}
)

func (x ChangeEnergyReason) Enum() *ChangeEnergyReason {
	p := new(ChangeEnergyReason)
	*p = x
	return p
}

func (x ChangeEnergyReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeEnergyReason) Descriptor() protoreflect.EnumDescriptor {
	return file_ChangeEnergyReason_proto_enumTypes[0].Descriptor()
}

func (ChangeEnergyReason) Type() protoreflect.EnumType {
	return &file_ChangeEnergyReason_proto_enumTypes[0]
}

func (x ChangeEnergyReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeEnergyReason.Descriptor instead.
func (ChangeEnergyReason) EnumDescriptor() ([]byte, []int) {
	return file_ChangeEnergyReason_proto_rawDescGZIP(), []int{0}
}

var File_ChangeEnergyReason_proto protoreflect.FileDescriptor

var file_ChangeEnergyReason_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x59, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x65, 0x72, 0x67,
	0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53,
	0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChangeEnergyReason_proto_rawDescOnce sync.Once
	file_ChangeEnergyReason_proto_rawDescData = file_ChangeEnergyReason_proto_rawDesc
)

func file_ChangeEnergyReason_proto_rawDescGZIP() []byte {
	file_ChangeEnergyReason_proto_rawDescOnce.Do(func() {
		file_ChangeEnergyReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeEnergyReason_proto_rawDescData)
	})
	return file_ChangeEnergyReason_proto_rawDescData
}

var file_ChangeEnergyReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChangeEnergyReason_proto_goTypes = []interface{}{
	(ChangeEnergyReason)(0), // 0: proto.ChangeEnergyReason
}
var file_ChangeEnergyReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChangeEnergyReason_proto_init() }
func file_ChangeEnergyReason_proto_init() {
	if File_ChangeEnergyReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChangeEnergyReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeEnergyReason_proto_goTypes,
		DependencyIndexes: file_ChangeEnergyReason_proto_depIdxs,
		EnumInfos:         file_ChangeEnergyReason_proto_enumTypes,
	}.Build()
	File_ChangeEnergyReason_proto = out.File
	file_ChangeEnergyReason_proto_rawDesc = nil
	file_ChangeEnergyReason_proto_goTypes = nil
	file_ChangeEnergyReason_proto_depIdxs = nil
}
