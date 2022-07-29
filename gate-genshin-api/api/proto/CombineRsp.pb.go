// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: CombineRsp.proto

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

type CombineRsp_CmdId int32

const (
	CombineRsp_NONE             CombineRsp_CmdId = 0
	CombineRsp_ENET_CHANNEL_ID  CombineRsp_CmdId = 0
	CombineRsp_ENET_IS_RELIABLE CombineRsp_CmdId = 1
	CombineRsp_CMD_ID           CombineRsp_CmdId = 695
)

// Enum value maps for CombineRsp_CmdId.
var (
	CombineRsp_CmdId_name = map[int32]string{
		0: "NONE",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1:   "ENET_IS_RELIABLE",
		695: "CMD_ID",
	}
	CombineRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"CMD_ID":           695,
	}
)

func (x CombineRsp_CmdId) Enum() *CombineRsp_CmdId {
	p := new(CombineRsp_CmdId)
	*p = x
	return p
}

func (x CombineRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CombineRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_CombineRsp_proto_enumTypes[0].Descriptor()
}

func (CombineRsp_CmdId) Type() protoreflect.EnumType {
	return &file_CombineRsp_proto_enumTypes[0]
}

func (x CombineRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CombineRsp_CmdId.Descriptor instead.
func (CombineRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_CombineRsp_proto_rawDescGZIP(), []int{0, 0}
}

type CombineRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             int32        `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CombineId           uint32       `protobuf:"varint,2,opt,name=combine_id,json=combineId,proto3" json:"combine_id,omitempty"`
	CombineCount        uint32       `protobuf:"varint,3,opt,name=combine_count,json=combineCount,proto3" json:"combine_count,omitempty"`
	AvatarGuid          uint64       `protobuf:"varint,4,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	CostItemList        []*ItemParam `protobuf:"bytes,5,rep,name=cost_item_list,json=costItemList,proto3" json:"cost_item_list,omitempty"`
	ResultItemList      []*ItemParam `protobuf:"bytes,6,rep,name=result_item_list,json=resultItemList,proto3" json:"result_item_list,omitempty"`
	TotalRandomItemList []*ItemParam `protobuf:"bytes,7,rep,name=total_random_item_list,json=totalRandomItemList,proto3" json:"total_random_item_list,omitempty"`
	TotalReturnItemList []*ItemParam `protobuf:"bytes,8,rep,name=total_return_item_list,json=totalReturnItemList,proto3" json:"total_return_item_list,omitempty"`
	TotalExtraItemList  []*ItemParam `protobuf:"bytes,9,rep,name=total_extra_item_list,json=totalExtraItemList,proto3" json:"total_extra_item_list,omitempty"`
}

func (x *CombineRsp) Reset() {
	*x = CombineRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CombineRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombineRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombineRsp) ProtoMessage() {}

func (x *CombineRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CombineRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombineRsp.ProtoReflect.Descriptor instead.
func (*CombineRsp) Descriptor() ([]byte, []int) {
	return file_CombineRsp_proto_rawDescGZIP(), []int{0}
}

func (x *CombineRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *CombineRsp) GetCombineId() uint32 {
	if x != nil {
		return x.CombineId
	}
	return 0
}

func (x *CombineRsp) GetCombineCount() uint32 {
	if x != nil {
		return x.CombineCount
	}
	return 0
}

func (x *CombineRsp) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *CombineRsp) GetCostItemList() []*ItemParam {
	if x != nil {
		return x.CostItemList
	}
	return nil
}

func (x *CombineRsp) GetResultItemList() []*ItemParam {
	if x != nil {
		return x.ResultItemList
	}
	return nil
}

func (x *CombineRsp) GetTotalRandomItemList() []*ItemParam {
	if x != nil {
		return x.TotalRandomItemList
	}
	return nil
}

func (x *CombineRsp) GetTotalReturnItemList() []*ItemParam {
	if x != nil {
		return x.TotalReturnItemList
	}
	return nil
}

func (x *CombineRsp) GetTotalExtraItemList() []*ItemParam {
	if x != nil {
		return x.TotalExtraItemList
	}
	return nil
}

var File_CombineRsp_proto protoreflect.FileDescriptor

var file_CombineRsp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x0a, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x16,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x13,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x15, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x12, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb7, 0x05, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CombineRsp_proto_rawDescOnce sync.Once
	file_CombineRsp_proto_rawDescData = file_CombineRsp_proto_rawDesc
)

func file_CombineRsp_proto_rawDescGZIP() []byte {
	file_CombineRsp_proto_rawDescOnce.Do(func() {
		file_CombineRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_CombineRsp_proto_rawDescData)
	})
	return file_CombineRsp_proto_rawDescData
}

var file_CombineRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CombineRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CombineRsp_proto_goTypes = []interface{}{
	(CombineRsp_CmdId)(0), // 0: proto.CombineRsp.CmdId
	(*CombineRsp)(nil),    // 1: proto.CombineRsp
	(*ItemParam)(nil),     // 2: proto.ItemParam
}
var file_CombineRsp_proto_depIdxs = []int32{
	2, // 0: proto.CombineRsp.cost_item_list:type_name -> proto.ItemParam
	2, // 1: proto.CombineRsp.result_item_list:type_name -> proto.ItemParam
	2, // 2: proto.CombineRsp.total_random_item_list:type_name -> proto.ItemParam
	2, // 3: proto.CombineRsp.total_return_item_list:type_name -> proto.ItemParam
	2, // 4: proto.CombineRsp.total_extra_item_list:type_name -> proto.ItemParam
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_CombineRsp_proto_init() }
func file_CombineRsp_proto_init() {
	if File_CombineRsp_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CombineRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombineRsp); i {
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
			RawDescriptor: file_CombineRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CombineRsp_proto_goTypes,
		DependencyIndexes: file_CombineRsp_proto_depIdxs,
		EnumInfos:         file_CombineRsp_proto_enumTypes,
		MessageInfos:      file_CombineRsp_proto_msgTypes,
	}.Build()
	File_CombineRsp_proto = out.File
	file_CombineRsp_proto_rawDesc = nil
	file_CombineRsp_proto_goTypes = nil
	file_CombineRsp_proto_depIdxs = nil
}
