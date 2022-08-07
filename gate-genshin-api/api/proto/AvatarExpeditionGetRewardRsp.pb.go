// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarExpeditionGetRewardRsp.proto

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

// CmdId: 1646
// EnetChannelId: 0
// EnetIsReliable: true
type AvatarExpeditionGetRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode           int32                            `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ExpeditionInfoMap map[uint64]*AvatarExpeditionInfo `protobuf:"bytes,5,rep,name=expedition_info_map,json=expeditionInfoMap,proto3" json:"expedition_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ItemList          []*ItemParam                     `protobuf:"bytes,6,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	BNFDDKNNJJH       []*ItemParam                     `protobuf:"bytes,2,rep,name=BNFDDKNNJJH,proto3" json:"BNFDDKNNJJH,omitempty"`
}

func (x *AvatarExpeditionGetRewardRsp) Reset() {
	*x = AvatarExpeditionGetRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarExpeditionGetRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExpeditionGetRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExpeditionGetRewardRsp) ProtoMessage() {}

func (x *AvatarExpeditionGetRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarExpeditionGetRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExpeditionGetRewardRsp.ProtoReflect.Descriptor instead.
func (*AvatarExpeditionGetRewardRsp) Descriptor() ([]byte, []int) {
	return file_AvatarExpeditionGetRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExpeditionGetRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AvatarExpeditionGetRewardRsp) GetExpeditionInfoMap() map[uint64]*AvatarExpeditionInfo {
	if x != nil {
		return x.ExpeditionInfoMap
	}
	return nil
}

func (x *AvatarExpeditionGetRewardRsp) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *AvatarExpeditionGetRewardRsp) GetBNFDDKNNJJH() []*ItemParam {
	if x != nil {
		return x.BNFDDKNNJJH
	}
	return nil
}

var File_AvatarExpeditionGetRewardRsp_proto protoreflect.FileDescriptor

var file_AvatarExpeditionGetRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x1c, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x6a, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45,
	0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x65, 0x78,
	0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12,
	0x2d, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x0b, 0x42, 0x4e, 0x46, 0x44, 0x44, 0x4b, 0x4e, 0x4e, 0x4a, 0x4a, 0x48, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0b, 0x42, 0x4e, 0x46, 0x44, 0x44, 0x4b, 0x4e, 0x4e, 0x4a,
	0x4a, 0x48, 0x1a, 0x61, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExpeditionGetRewardRsp_proto_rawDescOnce sync.Once
	file_AvatarExpeditionGetRewardRsp_proto_rawDescData = file_AvatarExpeditionGetRewardRsp_proto_rawDesc
)

func file_AvatarExpeditionGetRewardRsp_proto_rawDescGZIP() []byte {
	file_AvatarExpeditionGetRewardRsp_proto_rawDescOnce.Do(func() {
		file_AvatarExpeditionGetRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExpeditionGetRewardRsp_proto_rawDescData)
	})
	return file_AvatarExpeditionGetRewardRsp_proto_rawDescData
}

var file_AvatarExpeditionGetRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarExpeditionGetRewardRsp_proto_goTypes = []interface{}{
	(*AvatarExpeditionGetRewardRsp)(nil), // 0: proto.AvatarExpeditionGetRewardRsp
	nil,                                  // 1: proto.AvatarExpeditionGetRewardRsp.ExpeditionInfoMapEntry
	(*ItemParam)(nil),                    // 2: proto.ItemParam
	(*AvatarExpeditionInfo)(nil),         // 3: proto.AvatarExpeditionInfo
}
var file_AvatarExpeditionGetRewardRsp_proto_depIdxs = []int32{
	1, // 0: proto.AvatarExpeditionGetRewardRsp.expedition_info_map:type_name -> proto.AvatarExpeditionGetRewardRsp.ExpeditionInfoMapEntry
	2, // 1: proto.AvatarExpeditionGetRewardRsp.item_list:type_name -> proto.ItemParam
	2, // 2: proto.AvatarExpeditionGetRewardRsp.BNFDDKNNJJH:type_name -> proto.ItemParam
	3, // 3: proto.AvatarExpeditionGetRewardRsp.ExpeditionInfoMapEntry.value:type_name -> proto.AvatarExpeditionInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AvatarExpeditionGetRewardRsp_proto_init() }
func file_AvatarExpeditionGetRewardRsp_proto_init() {
	if File_AvatarExpeditionGetRewardRsp_proto != nil {
		return
	}
	file_AvatarExpeditionInfo_proto_init()
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarExpeditionGetRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExpeditionGetRewardRsp); i {
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
			RawDescriptor: file_AvatarExpeditionGetRewardRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExpeditionGetRewardRsp_proto_goTypes,
		DependencyIndexes: file_AvatarExpeditionGetRewardRsp_proto_depIdxs,
		MessageInfos:      file_AvatarExpeditionGetRewardRsp_proto_msgTypes,
	}.Build()
	File_AvatarExpeditionGetRewardRsp_proto = out.File
	file_AvatarExpeditionGetRewardRsp_proto_rawDesc = nil
	file_AvatarExpeditionGetRewardRsp_proto_goTypes = nil
	file_AvatarExpeditionGetRewardRsp_proto_depIdxs = nil
}
