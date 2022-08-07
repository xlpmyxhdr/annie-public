// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ClientCollectorData.proto

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

type ClientCollectorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaterialId uint32 `protobuf:"varint,6,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
	MaxPoints  uint32 `protobuf:"varint,2,opt,name=max_points,json=maxPoints,proto3" json:"max_points,omitempty"`
	CurrPoints uint32 `protobuf:"varint,8,opt,name=curr_points,json=currPoints,proto3" json:"curr_points,omitempty"`
}

func (x *ClientCollectorData) Reset() {
	*x = ClientCollectorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientCollectorData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCollectorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCollectorData) ProtoMessage() {}

func (x *ClientCollectorData) ProtoReflect() protoreflect.Message {
	mi := &file_ClientCollectorData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCollectorData.ProtoReflect.Descriptor instead.
func (*ClientCollectorData) Descriptor() ([]byte, []int) {
	return file_ClientCollectorData_proto_rawDescGZIP(), []int{0}
}

func (x *ClientCollectorData) GetMaterialId() uint32 {
	if x != nil {
		return x.MaterialId
	}
	return 0
}

func (x *ClientCollectorData) GetMaxPoints() uint32 {
	if x != nil {
		return x.MaxPoints
	}
	return 0
}

func (x *ClientCollectorData) GetCurrPoints() uint32 {
	if x != nil {
		return x.CurrPoints
	}
	return 0
}

var File_ClientCollectorData_proto protoreflect.FileDescriptor

var file_ClientCollectorData_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x76, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72,
	0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x63, 0x75, 0x72, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientCollectorData_proto_rawDescOnce sync.Once
	file_ClientCollectorData_proto_rawDescData = file_ClientCollectorData_proto_rawDesc
)

func file_ClientCollectorData_proto_rawDescGZIP() []byte {
	file_ClientCollectorData_proto_rawDescOnce.Do(func() {
		file_ClientCollectorData_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientCollectorData_proto_rawDescData)
	})
	return file_ClientCollectorData_proto_rawDescData
}

var file_ClientCollectorData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientCollectorData_proto_goTypes = []interface{}{
	(*ClientCollectorData)(nil), // 0: proto.ClientCollectorData
}
var file_ClientCollectorData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClientCollectorData_proto_init() }
func file_ClientCollectorData_proto_init() {
	if File_ClientCollectorData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientCollectorData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCollectorData); i {
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
			RawDescriptor: file_ClientCollectorData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientCollectorData_proto_goTypes,
		DependencyIndexes: file_ClientCollectorData_proto_depIdxs,
		MessageInfos:      file_ClientCollectorData_proto_msgTypes,
	}.Build()
	File_ClientCollectorData_proto = out.File
	file_ClientCollectorData_proto_rawDesc = nil
	file_ClientCollectorData_proto_goTypes = nil
	file_ClientCollectorData_proto_depIdxs = nil
}
