// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: WeaponAwakenReq.proto

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

type WeaponAwakenReq_CmdId int32

const (
	WeaponAwakenReq_ENET_CHANNEL_ID  WeaponAwakenReq_CmdId = 0
	WeaponAwakenReq_NONE             WeaponAwakenReq_CmdId = 0
	WeaponAwakenReq_ENET_IS_RELIABLE WeaponAwakenReq_CmdId = 1
	WeaponAwakenReq_IS_ALLOW_CLIENT  WeaponAwakenReq_CmdId = 1
	WeaponAwakenReq_CMD_ID           WeaponAwakenReq_CmdId = 656
)

// Enum value maps for WeaponAwakenReq_CmdId.
var (
	WeaponAwakenReq_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		656: "CMD_ID",
	}
	WeaponAwakenReq_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           656,
	}
)

func (x WeaponAwakenReq_CmdId) Enum() *WeaponAwakenReq_CmdId {
	p := new(WeaponAwakenReq_CmdId)
	*p = x
	return p
}

func (x WeaponAwakenReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeaponAwakenReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WeaponAwakenReq_proto_enumTypes[0].Descriptor()
}

func (WeaponAwakenReq_CmdId) Type() protoreflect.EnumType {
	return &file_WeaponAwakenReq_proto_enumTypes[0]
}

func (x WeaponAwakenReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeaponAwakenReq_CmdId.Descriptor instead.
func (WeaponAwakenReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WeaponAwakenReq_proto_rawDescGZIP(), []int{0, 0}
}

type WeaponAwakenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetWeaponGuid uint64            `protobuf:"varint,1,opt,name=target_weapon_guid,json=targetWeaponGuid,proto3" json:"target_weapon_guid,omitempty"`
	ItemGuid         uint64            `protobuf:"varint,2,opt,name=item_guid,json=itemGuid,proto3" json:"item_guid,omitempty"`
	AffixLevelMap    map[uint32]uint32 `protobuf:"bytes,3,rep,name=affix_level_map,json=affixLevelMap,proto3" json:"affix_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *WeaponAwakenReq) Reset() {
	*x = WeaponAwakenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WeaponAwakenReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeaponAwakenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeaponAwakenReq) ProtoMessage() {}

func (x *WeaponAwakenReq) ProtoReflect() protoreflect.Message {
	mi := &file_WeaponAwakenReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeaponAwakenReq.ProtoReflect.Descriptor instead.
func (*WeaponAwakenReq) Descriptor() ([]byte, []int) {
	return file_WeaponAwakenReq_proto_rawDescGZIP(), []int{0}
}

func (x *WeaponAwakenReq) GetTargetWeaponGuid() uint64 {
	if x != nil {
		return x.TargetWeaponGuid
	}
	return 0
}

func (x *WeaponAwakenReq) GetItemGuid() uint64 {
	if x != nil {
		return x.ItemGuid
	}
	return 0
}

func (x *WeaponAwakenReq) GetAffixLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.AffixLevelMap
	}
	return nil
}

var File_WeaponAwakenReq_proto protoreflect.FileDescriptor

var file_WeaponAwakenReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x41, 0x77, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x02, 0x0a, 0x0f, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x41, 0x77, 0x61, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x12, 0x51, 0x0a,
	0x0f, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x41, 0x77, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x41,
	0x66, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0d, 0x61, 0x66, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x1a, 0x40, 0x0a, 0x12, 0x41, 0x66, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x90, 0x05, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WeaponAwakenReq_proto_rawDescOnce sync.Once
	file_WeaponAwakenReq_proto_rawDescData = file_WeaponAwakenReq_proto_rawDesc
)

func file_WeaponAwakenReq_proto_rawDescGZIP() []byte {
	file_WeaponAwakenReq_proto_rawDescOnce.Do(func() {
		file_WeaponAwakenReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_WeaponAwakenReq_proto_rawDescData)
	})
	return file_WeaponAwakenReq_proto_rawDescData
}

var file_WeaponAwakenReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WeaponAwakenReq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_WeaponAwakenReq_proto_goTypes = []interface{}{
	(WeaponAwakenReq_CmdId)(0), // 0: proto.WeaponAwakenReq.CmdId
	(*WeaponAwakenReq)(nil),    // 1: proto.WeaponAwakenReq
	nil,                        // 2: proto.WeaponAwakenReq.AffixLevelMapEntry
}
var file_WeaponAwakenReq_proto_depIdxs = []int32{
	2, // 0: proto.WeaponAwakenReq.affix_level_map:type_name -> proto.WeaponAwakenReq.AffixLevelMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WeaponAwakenReq_proto_init() }
func file_WeaponAwakenReq_proto_init() {
	if File_WeaponAwakenReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WeaponAwakenReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeaponAwakenReq); i {
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
			RawDescriptor: file_WeaponAwakenReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WeaponAwakenReq_proto_goTypes,
		DependencyIndexes: file_WeaponAwakenReq_proto_depIdxs,
		EnumInfos:         file_WeaponAwakenReq_proto_enumTypes,
		MessageInfos:      file_WeaponAwakenReq_proto_msgTypes,
	}.Build()
	File_WeaponAwakenReq_proto = out.File
	file_WeaponAwakenReq_proto_rawDesc = nil
	file_WeaponAwakenReq_proto_goTypes = nil
	file_WeaponAwakenReq_proto_depIdxs = nil
}
