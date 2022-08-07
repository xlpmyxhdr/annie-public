// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: WidgetSlotChangeNotify.proto

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

// CmdId: 4274
// EnetChannelId: 0
// EnetIsReliable: true
type WidgetSlotChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   WidgetSlotOp    `protobuf:"varint,6,opt,name=op,proto3,enum=proto.WidgetSlotOp" json:"op,omitempty"`
	Slot *WidgetSlotData `protobuf:"bytes,8,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *WidgetSlotChangeNotify) Reset() {
	*x = WidgetSlotChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetSlotChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetSlotChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetSlotChangeNotify) ProtoMessage() {}

func (x *WidgetSlotChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetSlotChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetSlotChangeNotify.ProtoReflect.Descriptor instead.
func (*WidgetSlotChangeNotify) Descriptor() ([]byte, []int) {
	return file_WidgetSlotChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetSlotChangeNotify) GetOp() WidgetSlotOp {
	if x != nil {
		return x.Op
	}
	return WidgetSlotOp_WIDGET_SLOT_OP_ATTACH
}

func (x *WidgetSlotChangeNotify) GetSlot() *WidgetSlotData {
	if x != nil {
		return x.Slot
	}
	return nil
}

var File_WidgetSlotChangeNotify_proto protoreflect.FileDescriptor

var file_WidgetSlotChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x4f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x16, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x29,
	0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetSlotChangeNotify_proto_rawDescOnce sync.Once
	file_WidgetSlotChangeNotify_proto_rawDescData = file_WidgetSlotChangeNotify_proto_rawDesc
)

func file_WidgetSlotChangeNotify_proto_rawDescGZIP() []byte {
	file_WidgetSlotChangeNotify_proto_rawDescOnce.Do(func() {
		file_WidgetSlotChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetSlotChangeNotify_proto_rawDescData)
	})
	return file_WidgetSlotChangeNotify_proto_rawDescData
}

var file_WidgetSlotChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetSlotChangeNotify_proto_goTypes = []interface{}{
	(*WidgetSlotChangeNotify)(nil), // 0: proto.WidgetSlotChangeNotify
	(WidgetSlotOp)(0),              // 1: proto.WidgetSlotOp
	(*WidgetSlotData)(nil),         // 2: proto.WidgetSlotData
}
var file_WidgetSlotChangeNotify_proto_depIdxs = []int32{
	1, // 0: proto.WidgetSlotChangeNotify.op:type_name -> proto.WidgetSlotOp
	2, // 1: proto.WidgetSlotChangeNotify.slot:type_name -> proto.WidgetSlotData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WidgetSlotChangeNotify_proto_init() }
func file_WidgetSlotChangeNotify_proto_init() {
	if File_WidgetSlotChangeNotify_proto != nil {
		return
	}
	file_WidgetSlotData_proto_init()
	file_WidgetSlotOp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WidgetSlotChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetSlotChangeNotify); i {
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
			RawDescriptor: file_WidgetSlotChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetSlotChangeNotify_proto_goTypes,
		DependencyIndexes: file_WidgetSlotChangeNotify_proto_depIdxs,
		MessageInfos:      file_WidgetSlotChangeNotify_proto_msgTypes,
	}.Build()
	File_WidgetSlotChangeNotify_proto = out.File
	file_WidgetSlotChangeNotify_proto_rawDesc = nil
	file_WidgetSlotChangeNotify_proto_goTypes = nil
	file_WidgetSlotChangeNotify_proto_depIdxs = nil
}
