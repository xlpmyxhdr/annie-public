// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: BossChestInfo.proto

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

type BossChestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonsterConfigId uint32                                  `protobuf:"varint,1,opt,name=monster_config_id,json=monsterConfigId,proto3" json:"monster_config_id,omitempty"`
	Resin           uint32                                  `protobuf:"varint,2,opt,name=resin,proto3" json:"resin,omitempty"`
	RemainUidList   []uint32                                `protobuf:"varint,3,rep,packed,name=remain_uid_list,json=remainUidList,proto3" json:"remain_uid_list,omitempty"`
	QualifyUidList  []uint32                                `protobuf:"varint,4,rep,packed,name=qualify_uid_list,json=qualifyUidList,proto3" json:"qualify_uid_list,omitempty"`
	UidDiscountMap  map[uint32]*WeeklyBossResinDiscountInfo `protobuf:"bytes,5,rep,name=uid_discount_map,json=uidDiscountMap,proto3" json:"uid_discount_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BossChestInfo) Reset() {
	*x = BossChestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BossChestInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BossChestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BossChestInfo) ProtoMessage() {}

func (x *BossChestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BossChestInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BossChestInfo.ProtoReflect.Descriptor instead.
func (*BossChestInfo) Descriptor() ([]byte, []int) {
	return file_BossChestInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BossChestInfo) GetMonsterConfigId() uint32 {
	if x != nil {
		return x.MonsterConfigId
	}
	return 0
}

func (x *BossChestInfo) GetResin() uint32 {
	if x != nil {
		return x.Resin
	}
	return 0
}

func (x *BossChestInfo) GetRemainUidList() []uint32 {
	if x != nil {
		return x.RemainUidList
	}
	return nil
}

func (x *BossChestInfo) GetQualifyUidList() []uint32 {
	if x != nil {
		return x.QualifyUidList
	}
	return nil
}

func (x *BossChestInfo) GetUidDiscountMap() map[uint32]*WeeklyBossResinDiscountInfo {
	if x != nil {
		return x.UidDiscountMap
	}
	return nil
}

var File_BossChestInfo_proto protoreflect.FileDescriptor

var file_BossChestInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x57, 0x65,
	0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xde, 0x02, 0x0a, 0x0d, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x73, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65,
	0x73, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x75, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x55, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x66, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x79, 0x55, 0x69,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x10, 0x75, 0x69, 0x64, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x69, 0x64, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x75, 0x69, 0x64, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x65, 0x0a, 0x13, 0x55, 0x69, 0x64,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x42, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BossChestInfo_proto_rawDescOnce sync.Once
	file_BossChestInfo_proto_rawDescData = file_BossChestInfo_proto_rawDesc
)

func file_BossChestInfo_proto_rawDescGZIP() []byte {
	file_BossChestInfo_proto_rawDescOnce.Do(func() {
		file_BossChestInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BossChestInfo_proto_rawDescData)
	})
	return file_BossChestInfo_proto_rawDescData
}

var file_BossChestInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_BossChestInfo_proto_goTypes = []interface{}{
	(*BossChestInfo)(nil),               // 0: proto.BossChestInfo
	nil,                                 // 1: proto.BossChestInfo.UidDiscountMapEntry
	(*WeeklyBossResinDiscountInfo)(nil), // 2: proto.WeeklyBossResinDiscountInfo
}
var file_BossChestInfo_proto_depIdxs = []int32{
	1, // 0: proto.BossChestInfo.uid_discount_map:type_name -> proto.BossChestInfo.UidDiscountMapEntry
	2, // 1: proto.BossChestInfo.UidDiscountMapEntry.value:type_name -> proto.WeeklyBossResinDiscountInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BossChestInfo_proto_init() }
func file_BossChestInfo_proto_init() {
	if File_BossChestInfo_proto != nil {
		return
	}
	file_WeeklyBossResinDiscountInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BossChestInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BossChestInfo); i {
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
			RawDescriptor: file_BossChestInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BossChestInfo_proto_goTypes,
		DependencyIndexes: file_BossChestInfo_proto_depIdxs,
		MessageInfos:      file_BossChestInfo_proto_msgTypes,
	}.Build()
	File_BossChestInfo_proto = out.File
	file_BossChestInfo_proto_rawDesc = nil
	file_BossChestInfo_proto_goTypes = nil
	file_BossChestInfo_proto_depIdxs = nil
}
