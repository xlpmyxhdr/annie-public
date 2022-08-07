// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: PrivateChatReq.proto

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

type PrivateChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid uint32 `protobuf:"varint,2,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	// Types that are assignable to Content:
	//	*PrivateChatReq_Text
	//	*PrivateChatReq_Icon
	Content isPrivateChatReq_Content `protobuf_oneof:"content"`
}

func (x *PrivateChatReq) Reset() {
	*x = PrivateChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PrivateChatReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateChatReq) ProtoMessage() {}

func (x *PrivateChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_PrivateChatReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateChatReq.ProtoReflect.Descriptor instead.
func (*PrivateChatReq) Descriptor() ([]byte, []int) {
	return file_PrivateChatReq_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateChatReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (m *PrivateChatReq) GetContent() isPrivateChatReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *PrivateChatReq) GetText() string {
	if x, ok := x.GetContent().(*PrivateChatReq_Text); ok {
		return x.Text
	}
	return ""
}

func (x *PrivateChatReq) GetIcon() uint32 {
	if x, ok := x.GetContent().(*PrivateChatReq_Icon); ok {
		return x.Icon
	}
	return 0
}

type isPrivateChatReq_Content interface {
	isPrivateChatReq_Content()
}

type PrivateChatReq_Text struct {
	Text string `protobuf:"bytes,8,opt,name=text,proto3,oneof"`
}

type PrivateChatReq_Icon struct {
	Icon uint32 `protobuf:"varint,6,opt,name=icon,proto3,oneof"`
}

func (*PrivateChatReq_Text) isPrivateChatReq_Content() {}

func (*PrivateChatReq_Icon) isPrivateChatReq_Content() {}

var File_PrivateChatReq_proto protoreflect.FileDescriptor

var file_PrivateChatReq_proto_rawDesc = []byte{
	0x0a, 0x14, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a,
	0x0e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PrivateChatReq_proto_rawDescOnce sync.Once
	file_PrivateChatReq_proto_rawDescData = file_PrivateChatReq_proto_rawDesc
)

func file_PrivateChatReq_proto_rawDescGZIP() []byte {
	file_PrivateChatReq_proto_rawDescOnce.Do(func() {
		file_PrivateChatReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PrivateChatReq_proto_rawDescData)
	})
	return file_PrivateChatReq_proto_rawDescData
}

var file_PrivateChatReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PrivateChatReq_proto_goTypes = []interface{}{
	(*PrivateChatReq)(nil), // 0: proto.PrivateChatReq
}
var file_PrivateChatReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PrivateChatReq_proto_init() }
func file_PrivateChatReq_proto_init() {
	if File_PrivateChatReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PrivateChatReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateChatReq); i {
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
	file_PrivateChatReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PrivateChatReq_Text)(nil),
		(*PrivateChatReq_Icon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PrivateChatReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PrivateChatReq_proto_goTypes,
		DependencyIndexes: file_PrivateChatReq_proto_depIdxs,
		MessageInfos:      file_PrivateChatReq_proto_msgTypes,
	}.Build()
	File_PrivateChatReq_proto = out.File
	file_PrivateChatReq_proto_rawDesc = nil
	file_PrivateChatReq_proto_goTypes = nil
	file_PrivateChatReq_proto_depIdxs = nil
}
