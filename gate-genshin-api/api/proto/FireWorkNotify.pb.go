// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: FireWorkNotify.proto

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

// CmdId: 6079
// EnetChannelId: 0
// EnetIsReliable: true
type FireWorkNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FireWorkData []*FireWorkData `protobuf:"bytes,1,rep,name=fireWorkData,proto3" json:"fireWorkData,omitempty"`
}

func (x *FireWorkNotify) Reset() {
	*x = FireWorkNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FireWorkNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FireWorkNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FireWorkNotify) ProtoMessage() {}

func (x *FireWorkNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FireWorkNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FireWorkNotify.ProtoReflect.Descriptor instead.
func (*FireWorkNotify) Descriptor() ([]byte, []int) {
	return file_FireWorkNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FireWorkNotify) GetFireWorkData() []*FireWorkData {
	if x != nil {
		return x.FireWorkData
	}
	return nil
}

var File_FireWorkNotify_proto protoreflect.FileDescriptor

var file_FireWorkNotify_proto_rawDesc = []byte{
	0x0a, 0x14, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x46,
	0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x0e, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x37, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c,
	0x66, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FireWorkNotify_proto_rawDescOnce sync.Once
	file_FireWorkNotify_proto_rawDescData = file_FireWorkNotify_proto_rawDesc
)

func file_FireWorkNotify_proto_rawDescGZIP() []byte {
	file_FireWorkNotify_proto_rawDescOnce.Do(func() {
		file_FireWorkNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FireWorkNotify_proto_rawDescData)
	})
	return file_FireWorkNotify_proto_rawDescData
}

var file_FireWorkNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FireWorkNotify_proto_goTypes = []interface{}{
	(*FireWorkNotify)(nil), // 0: proto.FireWorkNotify
	(*FireWorkData)(nil),   // 1: proto.FireWorkData
}
var file_FireWorkNotify_proto_depIdxs = []int32{
	1, // 0: proto.FireWorkNotify.fireWorkData:type_name -> proto.FireWorkData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FireWorkNotify_proto_init() }
func file_FireWorkNotify_proto_init() {
	if File_FireWorkNotify_proto != nil {
		return
	}
	file_FireWorkData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FireWorkNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FireWorkNotify); i {
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
			RawDescriptor: file_FireWorkNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FireWorkNotify_proto_goTypes,
		DependencyIndexes: file_FireWorkNotify_proto_depIdxs,
		MessageInfos:      file_FireWorkNotify_proto_msgTypes,
	}.Build()
	File_FireWorkNotify_proto = out.File
	file_FireWorkNotify_proto_rawDesc = nil
	file_FireWorkNotify_proto_goTypes = nil
	file_FireWorkNotify_proto_depIdxs = nil
}
