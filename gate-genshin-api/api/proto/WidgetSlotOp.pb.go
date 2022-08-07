// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: WidgetSlotOp.proto

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

type WidgetSlotOp int32

const (
	WidgetSlotOp_WIDGET_SLOT_OP_ATTACH WidgetSlotOp = 0
	WidgetSlotOp_WIDGET_SLOT_OP_DETACH WidgetSlotOp = 1
)

// Enum value maps for WidgetSlotOp.
var (
	WidgetSlotOp_name = map[int32]string{
		0: "WIDGET_SLOT_OP_ATTACH",
		1: "WIDGET_SLOT_OP_DETACH",
	}
	WidgetSlotOp_value = map[string]int32{
		"WIDGET_SLOT_OP_ATTACH": 0,
		"WIDGET_SLOT_OP_DETACH": 1,
	}
)

func (x WidgetSlotOp) Enum() *WidgetSlotOp {
	p := new(WidgetSlotOp)
	*p = x
	return p
}

func (x WidgetSlotOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WidgetSlotOp) Descriptor() protoreflect.EnumDescriptor {
	return file_WidgetSlotOp_proto_enumTypes[0].Descriptor()
}

func (WidgetSlotOp) Type() protoreflect.EnumType {
	return &file_WidgetSlotOp_proto_enumTypes[0]
}

func (x WidgetSlotOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WidgetSlotOp.Descriptor instead.
func (WidgetSlotOp) EnumDescriptor() ([]byte, []int) {
	return file_WidgetSlotOp_proto_rawDescGZIP(), []int{0}
}

var File_WidgetSlotOp_proto protoreflect.FileDescriptor

var file_WidgetSlotOp_proto_rawDesc = []byte{
	0x0a, 0x12, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x4f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x44, 0x0a, 0x0c, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x4f, 0x70, 0x12, 0x19, 0x0a, 0x15, 0x57,
	0x49, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x4c, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x5f, 0x41, 0x54,
	0x54, 0x41, 0x43, 0x48, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x49, 0x44, 0x47, 0x45, 0x54,
	0x5f, 0x53, 0x4c, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x43, 0x48, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetSlotOp_proto_rawDescOnce sync.Once
	file_WidgetSlotOp_proto_rawDescData = file_WidgetSlotOp_proto_rawDesc
)

func file_WidgetSlotOp_proto_rawDescGZIP() []byte {
	file_WidgetSlotOp_proto_rawDescOnce.Do(func() {
		file_WidgetSlotOp_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetSlotOp_proto_rawDescData)
	})
	return file_WidgetSlotOp_proto_rawDescData
}

var file_WidgetSlotOp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WidgetSlotOp_proto_goTypes = []interface{}{
	(WidgetSlotOp)(0), // 0: proto.WidgetSlotOp
}
var file_WidgetSlotOp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WidgetSlotOp_proto_init() }
func file_WidgetSlotOp_proto_init() {
	if File_WidgetSlotOp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_WidgetSlotOp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetSlotOp_proto_goTypes,
		DependencyIndexes: file_WidgetSlotOp_proto_depIdxs,
		EnumInfos:         file_WidgetSlotOp_proto_enumTypes,
	}.Build()
	File_WidgetSlotOp_proto = out.File
	file_WidgetSlotOp_proto_rawDesc = nil
	file_WidgetSlotOp_proto_goTypes = nil
	file_WidgetSlotOp_proto_depIdxs = nil
}
