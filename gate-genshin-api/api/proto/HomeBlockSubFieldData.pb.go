// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: HomeBlockSubFieldData.proto

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

type HomeBlockSubFieldData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos *Vector `protobuf:"bytes,12,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot *Vector `protobuf:"bytes,14,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *HomeBlockSubFieldData) Reset() {
	*x = HomeBlockSubFieldData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeBlockSubFieldData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeBlockSubFieldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeBlockSubFieldData) ProtoMessage() {}

func (x *HomeBlockSubFieldData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeBlockSubFieldData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeBlockSubFieldData.ProtoReflect.Descriptor instead.
func (*HomeBlockSubFieldData) Descriptor() ([]byte, []int) {
	return file_HomeBlockSubFieldData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeBlockSubFieldData) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *HomeBlockSubFieldData) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_HomeBlockSubFieldData_proto protoreflect.FileDescriptor

var file_HomeBlockSubFieldData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x75, 0x62, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x59, 0x0a, 0x15, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x75, 0x62, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x03, 0x70,
	0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x03,
	0x72, 0x6f, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_HomeBlockSubFieldData_proto_rawDescOnce sync.Once
	file_HomeBlockSubFieldData_proto_rawDescData = file_HomeBlockSubFieldData_proto_rawDesc
)

func file_HomeBlockSubFieldData_proto_rawDescGZIP() []byte {
	file_HomeBlockSubFieldData_proto_rawDescOnce.Do(func() {
		file_HomeBlockSubFieldData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeBlockSubFieldData_proto_rawDescData)
	})
	return file_HomeBlockSubFieldData_proto_rawDescData
}

var file_HomeBlockSubFieldData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeBlockSubFieldData_proto_goTypes = []interface{}{
	(*HomeBlockSubFieldData)(nil), // 0: proto.HomeBlockSubFieldData
	(*Vector)(nil),                // 1: proto.Vector
}
var file_HomeBlockSubFieldData_proto_depIdxs = []int32{
	1, // 0: proto.HomeBlockSubFieldData.pos:type_name -> proto.Vector
	1, // 1: proto.HomeBlockSubFieldData.rot:type_name -> proto.Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HomeBlockSubFieldData_proto_init() }
func file_HomeBlockSubFieldData_proto_init() {
	if File_HomeBlockSubFieldData_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeBlockSubFieldData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeBlockSubFieldData); i {
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
			RawDescriptor: file_HomeBlockSubFieldData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeBlockSubFieldData_proto_goTypes,
		DependencyIndexes: file_HomeBlockSubFieldData_proto_depIdxs,
		MessageInfos:      file_HomeBlockSubFieldData_proto_msgTypes,
	}.Build()
	File_HomeBlockSubFieldData_proto = out.File
	file_HomeBlockSubFieldData_proto_rawDesc = nil
	file_HomeBlockSubFieldData_proto_goTypes = nil
	file_HomeBlockSubFieldData_proto_depIdxs = nil
}
