// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: WidgetActiveChangeNotify.proto

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

// CmdId: 4295
// EnetChannelId: 0
// EnetIsReliable: true
type WidgetActiveChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WidgetDataList []*WidgetSlotData `protobuf:"bytes,4,rep,name=widget_data_list,json=widgetDataList,proto3" json:"widget_data_list,omitempty"`
}

func (x *WidgetActiveChangeNotify) Reset() {
	*x = WidgetActiveChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetActiveChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetActiveChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetActiveChangeNotify) ProtoMessage() {}

func (x *WidgetActiveChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetActiveChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetActiveChangeNotify.ProtoReflect.Descriptor instead.
func (*WidgetActiveChangeNotify) Descriptor() ([]byte, []int) {
	return file_WidgetActiveChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetActiveChangeNotify) GetWidgetDataList() []*WidgetSlotData {
	if x != nil {
		return x.WidgetDataList
	}
	return nil
}

var File_WidgetActiveChangeNotify_proto protoreflect.FileDescriptor

var file_WidgetActiveChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53,
	0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a,
	0x18, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3f, 0x0a, 0x10, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetActiveChangeNotify_proto_rawDescOnce sync.Once
	file_WidgetActiveChangeNotify_proto_rawDescData = file_WidgetActiveChangeNotify_proto_rawDesc
)

func file_WidgetActiveChangeNotify_proto_rawDescGZIP() []byte {
	file_WidgetActiveChangeNotify_proto_rawDescOnce.Do(func() {
		file_WidgetActiveChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetActiveChangeNotify_proto_rawDescData)
	})
	return file_WidgetActiveChangeNotify_proto_rawDescData
}

var file_WidgetActiveChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetActiveChangeNotify_proto_goTypes = []interface{}{
	(*WidgetActiveChangeNotify)(nil), // 0: proto.WidgetActiveChangeNotify
	(*WidgetSlotData)(nil),           // 1: proto.WidgetSlotData
}
var file_WidgetActiveChangeNotify_proto_depIdxs = []int32{
	1, // 0: proto.WidgetActiveChangeNotify.widget_data_list:type_name -> proto.WidgetSlotData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WidgetActiveChangeNotify_proto_init() }
func file_WidgetActiveChangeNotify_proto_init() {
	if File_WidgetActiveChangeNotify_proto != nil {
		return
	}
	file_WidgetSlotData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WidgetActiveChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetActiveChangeNotify); i {
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
			RawDescriptor: file_WidgetActiveChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetActiveChangeNotify_proto_goTypes,
		DependencyIndexes: file_WidgetActiveChangeNotify_proto_depIdxs,
		MessageInfos:      file_WidgetActiveChangeNotify_proto_msgTypes,
	}.Build()
	File_WidgetActiveChangeNotify_proto = out.File
	file_WidgetActiveChangeNotify_proto_rawDesc = nil
	file_WidgetActiveChangeNotify_proto_goTypes = nil
	file_WidgetActiveChangeNotify_proto_depIdxs = nil
}
