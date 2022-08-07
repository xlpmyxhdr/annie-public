// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: HomeUpdateArrangementInfoReq.proto

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

// CmdId: 4472
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type HomeUpdateArrangementInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneArrangementInfo *HomeSceneArrangementInfo `protobuf:"bytes,12,opt,name=scene_arrangement_info,json=sceneArrangementInfo,proto3" json:"scene_arrangement_info,omitempty"`
}

func (x *HomeUpdateArrangementInfoReq) Reset() {
	*x = HomeUpdateArrangementInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeUpdateArrangementInfoReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeUpdateArrangementInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeUpdateArrangementInfoReq) ProtoMessage() {}

func (x *HomeUpdateArrangementInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeUpdateArrangementInfoReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeUpdateArrangementInfoReq.ProtoReflect.Descriptor instead.
func (*HomeUpdateArrangementInfoReq) Descriptor() ([]byte, []int) {
	return file_HomeUpdateArrangementInfoReq_proto_rawDescGZIP(), []int{0}
}

func (x *HomeUpdateArrangementInfoReq) GetSceneArrangementInfo() *HomeSceneArrangementInfo {
	if x != nil {
		return x.SceneArrangementInfo
	}
	return nil
}

var File_HomeUpdateArrangementInfoReq_proto protoreflect.FileDescriptor

var file_HomeUpdateArrangementInfoReq_proto_rawDesc = []byte{
	0x0a, 0x22, 0x48, 0x6f, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x48, 0x6f, 0x6d,
	0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x1c, 0x48,
	0x6f, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x55, 0x0a, 0x16, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeUpdateArrangementInfoReq_proto_rawDescOnce sync.Once
	file_HomeUpdateArrangementInfoReq_proto_rawDescData = file_HomeUpdateArrangementInfoReq_proto_rawDesc
)

func file_HomeUpdateArrangementInfoReq_proto_rawDescGZIP() []byte {
	file_HomeUpdateArrangementInfoReq_proto_rawDescOnce.Do(func() {
		file_HomeUpdateArrangementInfoReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeUpdateArrangementInfoReq_proto_rawDescData)
	})
	return file_HomeUpdateArrangementInfoReq_proto_rawDescData
}

var file_HomeUpdateArrangementInfoReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeUpdateArrangementInfoReq_proto_goTypes = []interface{}{
	(*HomeUpdateArrangementInfoReq)(nil), // 0: proto.HomeUpdateArrangementInfoReq
	(*HomeSceneArrangementInfo)(nil),     // 1: proto.HomeSceneArrangementInfo
}
var file_HomeUpdateArrangementInfoReq_proto_depIdxs = []int32{
	1, // 0: proto.HomeUpdateArrangementInfoReq.scene_arrangement_info:type_name -> proto.HomeSceneArrangementInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeUpdateArrangementInfoReq_proto_init() }
func file_HomeUpdateArrangementInfoReq_proto_init() {
	if File_HomeUpdateArrangementInfoReq_proto != nil {
		return
	}
	file_HomeSceneArrangementInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeUpdateArrangementInfoReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeUpdateArrangementInfoReq); i {
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
			RawDescriptor: file_HomeUpdateArrangementInfoReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeUpdateArrangementInfoReq_proto_goTypes,
		DependencyIndexes: file_HomeUpdateArrangementInfoReq_proto_depIdxs,
		MessageInfos:      file_HomeUpdateArrangementInfoReq_proto_msgTypes,
	}.Build()
	File_HomeUpdateArrangementInfoReq_proto = out.File
	file_HomeUpdateArrangementInfoReq_proto_rawDesc = nil
	file_HomeUpdateArrangementInfoReq_proto_goTypes = nil
	file_HomeUpdateArrangementInfoReq_proto_depIdxs = nil
}
