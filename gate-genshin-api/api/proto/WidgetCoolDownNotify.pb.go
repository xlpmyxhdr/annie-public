// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: WidgetCoolDownNotify.proto

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

// CmdId: 4263
// EnetChannelId: 0
// EnetIsReliable: true
type WidgetCoolDownNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCoolDownDataList  []*WidgetCoolDownData `protobuf:"bytes,4,rep,name=group_cool_down_data_list,json=groupCoolDownDataList,proto3" json:"group_cool_down_data_list,omitempty"`
	NormalCoolDownDataList []*WidgetCoolDownData `protobuf:"bytes,8,rep,name=normal_cool_down_data_list,json=normalCoolDownDataList,proto3" json:"normal_cool_down_data_list,omitempty"`
}

func (x *WidgetCoolDownNotify) Reset() {
	*x = WidgetCoolDownNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetCoolDownNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetCoolDownNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetCoolDownNotify) ProtoMessage() {}

func (x *WidgetCoolDownNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetCoolDownNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetCoolDownNotify.ProtoReflect.Descriptor instead.
func (*WidgetCoolDownNotify) Descriptor() ([]byte, []int) {
	return file_WidgetCoolDownNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetCoolDownNotify) GetGroupCoolDownDataList() []*WidgetCoolDownData {
	if x != nil {
		return x.GroupCoolDownDataList
	}
	return nil
}

func (x *WidgetCoolDownNotify) GetNormalCoolDownDataList() []*WidgetCoolDownData {
	if x != nil {
		return x.NormalCoolDownDataList
	}
	return nil
}

var File_WidgetCoolDownNotify_proto protoreflect.FileDescriptor

var file_WidgetCoolDownNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44,
	0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01,
	0x0a, 0x14, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x53, 0x0a, 0x19, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x63, 0x6f, 0x6f, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x15, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6f, 0x6c, 0x44,
	0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x1a, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6f, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x16, 0x6e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetCoolDownNotify_proto_rawDescOnce sync.Once
	file_WidgetCoolDownNotify_proto_rawDescData = file_WidgetCoolDownNotify_proto_rawDesc
)

func file_WidgetCoolDownNotify_proto_rawDescGZIP() []byte {
	file_WidgetCoolDownNotify_proto_rawDescOnce.Do(func() {
		file_WidgetCoolDownNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetCoolDownNotify_proto_rawDescData)
	})
	return file_WidgetCoolDownNotify_proto_rawDescData
}

var file_WidgetCoolDownNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetCoolDownNotify_proto_goTypes = []interface{}{
	(*WidgetCoolDownNotify)(nil), // 0: proto.WidgetCoolDownNotify
	(*WidgetCoolDownData)(nil),   // 1: proto.WidgetCoolDownData
}
var file_WidgetCoolDownNotify_proto_depIdxs = []int32{
	1, // 0: proto.WidgetCoolDownNotify.group_cool_down_data_list:type_name -> proto.WidgetCoolDownData
	1, // 1: proto.WidgetCoolDownNotify.normal_cool_down_data_list:type_name -> proto.WidgetCoolDownData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WidgetCoolDownNotify_proto_init() }
func file_WidgetCoolDownNotify_proto_init() {
	if File_WidgetCoolDownNotify_proto != nil {
		return
	}
	file_WidgetCoolDownData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WidgetCoolDownNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetCoolDownNotify); i {
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
			RawDescriptor: file_WidgetCoolDownNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetCoolDownNotify_proto_goTypes,
		DependencyIndexes: file_WidgetCoolDownNotify_proto_depIdxs,
		MessageInfos:      file_WidgetCoolDownNotify_proto_msgTypes,
	}.Build()
	File_WidgetCoolDownNotify_proto = out.File
	file_WidgetCoolDownNotify_proto_rawDesc = nil
	file_WidgetCoolDownNotify_proto_goTypes = nil
	file_WidgetCoolDownNotify_proto_depIdxs = nil
}
