// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: FireWorkData.proto

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

type FireWorkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FireWorkInstance []*FireWorkInstance `protobuf:"bytes,2,rep,name=fireWorkInstance,proto3" json:"fireWorkInstance,omitempty"`
}

func (x *FireWorkData) Reset() {
	*x = FireWorkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FireWorkData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FireWorkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FireWorkData) ProtoMessage() {}

func (x *FireWorkData) ProtoReflect() protoreflect.Message {
	mi := &file_FireWorkData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FireWorkData.ProtoReflect.Descriptor instead.
func (*FireWorkData) Descriptor() ([]byte, []int) {
	return file_FireWorkData_proto_rawDescGZIP(), []int{0}
}

func (x *FireWorkData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FireWorkData) GetFireWorkInstance() []*FireWorkInstance {
	if x != nil {
		return x.FireWorkInstance
	}
	return nil
}

var File_FireWorkData_proto protoreflect.FileDescriptor

var file_FireWorkData_proto_rawDesc = []byte{
	0x0a, 0x12, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x46, 0x69, 0x72,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0c, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x66, 0x69, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FireWorkData_proto_rawDescOnce sync.Once
	file_FireWorkData_proto_rawDescData = file_FireWorkData_proto_rawDesc
)

func file_FireWorkData_proto_rawDescGZIP() []byte {
	file_FireWorkData_proto_rawDescOnce.Do(func() {
		file_FireWorkData_proto_rawDescData = protoimpl.X.CompressGZIP(file_FireWorkData_proto_rawDescData)
	})
	return file_FireWorkData_proto_rawDescData
}

var file_FireWorkData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FireWorkData_proto_goTypes = []interface{}{
	(*FireWorkData)(nil),     // 0: proto.FireWorkData
	(*FireWorkInstance)(nil), // 1: proto.FireWorkInstance
}
var file_FireWorkData_proto_depIdxs = []int32{
	1, // 0: proto.FireWorkData.fireWorkInstance:type_name -> proto.FireWorkInstance
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FireWorkData_proto_init() }
func file_FireWorkData_proto_init() {
	if File_FireWorkData_proto != nil {
		return
	}
	file_FireWorkInstance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FireWorkData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FireWorkData); i {
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
			RawDescriptor: file_FireWorkData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FireWorkData_proto_goTypes,
		DependencyIndexes: file_FireWorkData_proto_depIdxs,
		MessageInfos:      file_FireWorkData_proto_msgTypes,
	}.Build()
	File_FireWorkData_proto = out.File
	file_FireWorkData_proto_rawDesc = nil
	file_FireWorkData_proto_goTypes = nil
	file_FireWorkData_proto_depIdxs = nil
}
