// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ActivityNullDetailInfo.proto

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

type ActivityNullDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActivityNullDetailInfo) Reset() {
	*x = ActivityNullDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityNullDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityNullDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityNullDetailInfo) ProtoMessage() {}

func (x *ActivityNullDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityNullDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityNullDetailInfo.ProtoReflect.Descriptor instead.
func (*ActivityNullDetailInfo) Descriptor() ([]byte, []int) {
	return file_ActivityNullDetailInfo_proto_rawDescGZIP(), []int{0}
}

var File_ActivityNullDetailInfo_proto protoreflect.FileDescriptor

var file_ActivityNullDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6c, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4e, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ActivityNullDetailInfo_proto_rawDescOnce sync.Once
	file_ActivityNullDetailInfo_proto_rawDescData = file_ActivityNullDetailInfo_proto_rawDesc
)

func file_ActivityNullDetailInfo_proto_rawDescGZIP() []byte {
	file_ActivityNullDetailInfo_proto_rawDescOnce.Do(func() {
		file_ActivityNullDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityNullDetailInfo_proto_rawDescData)
	})
	return file_ActivityNullDetailInfo_proto_rawDescData
}

var file_ActivityNullDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActivityNullDetailInfo_proto_goTypes = []interface{}{
	(*ActivityNullDetailInfo)(nil), // 0: proto.ActivityNullDetailInfo
}
var file_ActivityNullDetailInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ActivityNullDetailInfo_proto_init() }
func file_ActivityNullDetailInfo_proto_init() {
	if File_ActivityNullDetailInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ActivityNullDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityNullDetailInfo); i {
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
			RawDescriptor: file_ActivityNullDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityNullDetailInfo_proto_goTypes,
		DependencyIndexes: file_ActivityNullDetailInfo_proto_depIdxs,
		MessageInfos:      file_ActivityNullDetailInfo_proto_msgTypes,
	}.Build()
	File_ActivityNullDetailInfo_proto = out.File
	file_ActivityNullDetailInfo_proto_rawDesc = nil
	file_ActivityNullDetailInfo_proto_goTypes = nil
	file_ActivityNullDetailInfo_proto_depIdxs = nil
}
