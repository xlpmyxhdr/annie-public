// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: CustomGadgetTreeInfo.proto

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

type CustomGadgetTreeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeList []*CustomCommonNodeInfo `protobuf:"bytes,1,rep,name=node_list,json=nodeList,proto3" json:"node_list,omitempty"`
}

func (x *CustomGadgetTreeInfo) Reset() {
	*x = CustomGadgetTreeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CustomGadgetTreeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomGadgetTreeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomGadgetTreeInfo) ProtoMessage() {}

func (x *CustomGadgetTreeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CustomGadgetTreeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomGadgetTreeInfo.ProtoReflect.Descriptor instead.
func (*CustomGadgetTreeInfo) Descriptor() ([]byte, []int) {
	return file_CustomGadgetTreeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CustomGadgetTreeInfo) GetNodeList() []*CustomCommonNodeInfo {
	if x != nil {
		return x.NodeList
	}
	return nil
}

var File_CustomGadgetTreeInfo_proto protoreflect.FileDescriptor

var file_CustomGadgetTreeInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x54, 0x72,
	0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x50, 0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x54,
	0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CustomGadgetTreeInfo_proto_rawDescOnce sync.Once
	file_CustomGadgetTreeInfo_proto_rawDescData = file_CustomGadgetTreeInfo_proto_rawDesc
)

func file_CustomGadgetTreeInfo_proto_rawDescGZIP() []byte {
	file_CustomGadgetTreeInfo_proto_rawDescOnce.Do(func() {
		file_CustomGadgetTreeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CustomGadgetTreeInfo_proto_rawDescData)
	})
	return file_CustomGadgetTreeInfo_proto_rawDescData
}

var file_CustomGadgetTreeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CustomGadgetTreeInfo_proto_goTypes = []interface{}{
	(*CustomGadgetTreeInfo)(nil), // 0: proto.CustomGadgetTreeInfo
	(*CustomCommonNodeInfo)(nil), // 1: proto.CustomCommonNodeInfo
}
var file_CustomGadgetTreeInfo_proto_depIdxs = []int32{
	1, // 0: proto.CustomGadgetTreeInfo.node_list:type_name -> proto.CustomCommonNodeInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CustomGadgetTreeInfo_proto_init() }
func file_CustomGadgetTreeInfo_proto_init() {
	if File_CustomGadgetTreeInfo_proto != nil {
		return
	}
	file_CustomCommonNodeInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CustomGadgetTreeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomGadgetTreeInfo); i {
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
			RawDescriptor: file_CustomGadgetTreeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CustomGadgetTreeInfo_proto_goTypes,
		DependencyIndexes: file_CustomGadgetTreeInfo_proto_depIdxs,
		MessageInfos:      file_CustomGadgetTreeInfo_proto_msgTypes,
	}.Build()
	File_CustomGadgetTreeInfo_proto = out.File
	file_CustomGadgetTreeInfo_proto_rawDesc = nil
	file_CustomGadgetTreeInfo_proto_goTypes = nil
	file_CustomGadgetTreeInfo_proto_depIdxs = nil
}
