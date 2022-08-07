// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: GetAllMailRsp.proto

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

type GetAllMailRsp_CmdId int32

const (
	GetAllMailRsp_NONE             GetAllMailRsp_CmdId = 0
	GetAllMailRsp_ENET_CHANNEL_ID  GetAllMailRsp_CmdId = 0
	GetAllMailRsp_ENET_IS_RELIABLE GetAllMailRsp_CmdId = 1
	GetAllMailRsp_CMD_ID           GetAllMailRsp_CmdId = 1427
)

// Enum value maps for GetAllMailRsp_CmdId.
var (
	GetAllMailRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:    "ENET_IS_RELIABLE",
		1427: "CMD_ID",
	}
	GetAllMailRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           1427,
	}
)

func (x GetAllMailRsp_CmdId) Enum() *GetAllMailRsp_CmdId {
	p := new(GetAllMailRsp_CmdId)
	*p = x
	return p
}

func (x GetAllMailRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetAllMailRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetAllMailRsp_proto_enumTypes[0].Descriptor()
}

func (GetAllMailRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetAllMailRsp_proto_enumTypes[0]
}

func (x GetAllMailRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetAllMailRsp_CmdId.Descriptor instead.
func (GetAllMailRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetAllMailRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetAllMailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     int32       `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MailList    []*MailData `protobuf:"bytes,9,rep,name=mail_list,json=mailList,proto3" json:"mail_list,omitempty"`
	IsTruncated bool        `protobuf:"varint,5,opt,name=is_truncated,json=isTruncated,proto3" json:"is_truncated,omitempty"`
	ANKKGPJCINB bool        `protobuf:"varint,7,opt,name=ANKKGPJCINB,proto3" json:"ANKKGPJCINB,omitempty"`
}

func (x *GetAllMailRsp) Reset() {
	*x = GetAllMailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAllMailRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailRsp) ProtoMessage() {}

func (x *GetAllMailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAllMailRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailRsp.ProtoReflect.Descriptor instead.
func (*GetAllMailRsp) Descriptor() ([]byte, []int) {
	return file_GetAllMailRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllMailRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAllMailRsp) GetMailList() []*MailData {
	if x != nil {
		return x.MailList
	}
	return nil
}

func (x *GetAllMailRsp) GetIsTruncated() bool {
	if x != nil {
		return x.IsTruncated
	}
	return false
}

func (x *GetAllMailRsp) GetANKKGPJCINB() bool {
	if x != nil {
		return x.ANKKGPJCINB
	}
	return false
}

var File_GetAllMailRsp_proto protoreflect.FileDescriptor

var file_GetAllMailRsp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x4d, 0x61,
	0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x61,
	0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x75,
	0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73,
	0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4e, 0x4b,
	0x4b, 0x47, 0x50, 0x4a, 0x43, 0x49, 0x4e, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x41, 0x4e, 0x4b, 0x4b, 0x47, 0x50, 0x4a, 0x43, 0x49, 0x4e, 0x42, 0x22, 0x4d, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0x93, 0x0b, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAllMailRsp_proto_rawDescOnce sync.Once
	file_GetAllMailRsp_proto_rawDescData = file_GetAllMailRsp_proto_rawDesc
)

func file_GetAllMailRsp_proto_rawDescGZIP() []byte {
	file_GetAllMailRsp_proto_rawDescOnce.Do(func() {
		file_GetAllMailRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAllMailRsp_proto_rawDescData)
	})
	return file_GetAllMailRsp_proto_rawDescData
}

var file_GetAllMailRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetAllMailRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAllMailRsp_proto_goTypes = []interface{}{
	(GetAllMailRsp_CmdId)(0), // 0: proto.GetAllMailRsp.CmdId
	(*GetAllMailRsp)(nil),    // 1: proto.GetAllMailRsp
	(*MailData)(nil),         // 2: proto.MailData
}
var file_GetAllMailRsp_proto_depIdxs = []int32{
	2, // 0: proto.GetAllMailRsp.mail_list:type_name -> proto.MailData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetAllMailRsp_proto_init() }
func file_GetAllMailRsp_proto_init() {
	if File_GetAllMailRsp_proto != nil {
		return
	}
	file_MailData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetAllMailRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMailRsp); i {
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
			RawDescriptor: file_GetAllMailRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAllMailRsp_proto_goTypes,
		DependencyIndexes: file_GetAllMailRsp_proto_depIdxs,
		EnumInfos:         file_GetAllMailRsp_proto_enumTypes,
		MessageInfos:      file_GetAllMailRsp_proto_msgTypes,
	}.Build()
	File_GetAllMailRsp_proto = out.File
	file_GetAllMailRsp_proto_rawDesc = nil
	file_GetAllMailRsp_proto_goTypes = nil
	file_GetAllMailRsp_proto_depIdxs = nil
}
