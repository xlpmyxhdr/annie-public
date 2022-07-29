// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: CalcWeaponUpgradeReturnItemsReq.proto

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

type CalcWeaponUpgradeReturnItemsReq_CmdId int32

const (
	CalcWeaponUpgradeReturnItemsReq_ENET_CHANNEL_ID  CalcWeaponUpgradeReturnItemsReq_CmdId = 0
	CalcWeaponUpgradeReturnItemsReq_NONE             CalcWeaponUpgradeReturnItemsReq_CmdId = 0
	CalcWeaponUpgradeReturnItemsReq_ENET_IS_RELIABLE CalcWeaponUpgradeReturnItemsReq_CmdId = 1
	CalcWeaponUpgradeReturnItemsReq_IS_ALLOW_CLIENT  CalcWeaponUpgradeReturnItemsReq_CmdId = 1
	CalcWeaponUpgradeReturnItemsReq_CMD_ID           CalcWeaponUpgradeReturnItemsReq_CmdId = 669
)

// Enum value maps for CalcWeaponUpgradeReturnItemsReq_CmdId.
var (
	CalcWeaponUpgradeReturnItemsReq_CmdId_name = map[int32]string{
		0: "ENET_CHANNEL_ID",
		// Duplicate value: 0: "NONE",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
		669: "CMD_ID",
	}
	CalcWeaponUpgradeReturnItemsReq_CmdId_value = map[string]int32{
		"ENET_CHANNEL_ID":  0,
		"NONE":             0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
		"CMD_ID":           669,
	}
)

func (x CalcWeaponUpgradeReturnItemsReq_CmdId) Enum() *CalcWeaponUpgradeReturnItemsReq_CmdId {
	p := new(CalcWeaponUpgradeReturnItemsReq_CmdId)
	*p = x
	return p
}

func (x CalcWeaponUpgradeReturnItemsReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CalcWeaponUpgradeReturnItemsReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_CalcWeaponUpgradeReturnItemsReq_proto_enumTypes[0].Descriptor()
}

func (CalcWeaponUpgradeReturnItemsReq_CmdId) Type() protoreflect.EnumType {
	return &file_CalcWeaponUpgradeReturnItemsReq_proto_enumTypes[0]
}

func (x CalcWeaponUpgradeReturnItemsReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CalcWeaponUpgradeReturnItemsReq_CmdId.Descriptor instead.
func (CalcWeaponUpgradeReturnItemsReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescGZIP(), []int{0, 0}
}

type CalcWeaponUpgradeReturnItemsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetWeaponGuid   uint64       `protobuf:"varint,1,opt,name=target_weapon_guid,json=targetWeaponGuid,proto3" json:"target_weapon_guid,omitempty"`
	FoodWeaponGuidList []uint64     `protobuf:"varint,2,rep,packed,name=food_weapon_guid_list,json=foodWeaponGuidList,proto3" json:"food_weapon_guid_list,omitempty"`
	ItemParamList      []*ItemParam `protobuf:"bytes,3,rep,name=item_param_list,json=itemParamList,proto3" json:"item_param_list,omitempty"`
}

func (x *CalcWeaponUpgradeReturnItemsReq) Reset() {
	*x = CalcWeaponUpgradeReturnItemsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CalcWeaponUpgradeReturnItemsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcWeaponUpgradeReturnItemsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcWeaponUpgradeReturnItemsReq) ProtoMessage() {}

func (x *CalcWeaponUpgradeReturnItemsReq) ProtoReflect() protoreflect.Message {
	mi := &file_CalcWeaponUpgradeReturnItemsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcWeaponUpgradeReturnItemsReq.ProtoReflect.Descriptor instead.
func (*CalcWeaponUpgradeReturnItemsReq) Descriptor() ([]byte, []int) {
	return file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescGZIP(), []int{0}
}

func (x *CalcWeaponUpgradeReturnItemsReq) GetTargetWeaponGuid() uint64 {
	if x != nil {
		return x.TargetWeaponGuid
	}
	return 0
}

func (x *CalcWeaponUpgradeReturnItemsReq) GetFoodWeaponGuidList() []uint64 {
	if x != nil {
		return x.FoodWeaponGuidList
	}
	return nil
}

func (x *CalcWeaponUpgradeReturnItemsReq) GetItemParamList() []*ItemParam {
	if x != nil {
		return x.ItemParamList
	}
	return nil
}

var File_CalcWeaponUpgradeReturnItemsReq_proto protoreflect.FileDescriptor

var file_CalcWeaponUpgradeReturnItemsReq_proto_rawDesc = []byte{
	0x0a, 0x25, 0x43, 0x61, 0x6c, 0x63, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x02, 0x0a, 0x1f, 0x43, 0x61, 0x6c, 0x63, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x65,
	0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69,
	0x64, 0x12, 0x31, 0x0a, 0x15, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x12, 0x66, 0x6f, 0x6f, 0x64, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52,
	0x0d, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9d, 0x05, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescOnce sync.Once
	file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescData = file_CalcWeaponUpgradeReturnItemsReq_proto_rawDesc
)

func file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescGZIP() []byte {
	file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescOnce.Do(func() {
		file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescData)
	})
	return file_CalcWeaponUpgradeReturnItemsReq_proto_rawDescData
}

var file_CalcWeaponUpgradeReturnItemsReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CalcWeaponUpgradeReturnItemsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CalcWeaponUpgradeReturnItemsReq_proto_goTypes = []interface{}{
	(CalcWeaponUpgradeReturnItemsReq_CmdId)(0), // 0: proto.CalcWeaponUpgradeReturnItemsReq.CmdId
	(*CalcWeaponUpgradeReturnItemsReq)(nil),    // 1: proto.CalcWeaponUpgradeReturnItemsReq
	(*ItemParam)(nil),                          // 2: proto.ItemParam
}
var file_CalcWeaponUpgradeReturnItemsReq_proto_depIdxs = []int32{
	2, // 0: proto.CalcWeaponUpgradeReturnItemsReq.item_param_list:type_name -> proto.ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CalcWeaponUpgradeReturnItemsReq_proto_init() }
func file_CalcWeaponUpgradeReturnItemsReq_proto_init() {
	if File_CalcWeaponUpgradeReturnItemsReq_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CalcWeaponUpgradeReturnItemsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcWeaponUpgradeReturnItemsReq); i {
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
			RawDescriptor: file_CalcWeaponUpgradeReturnItemsReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CalcWeaponUpgradeReturnItemsReq_proto_goTypes,
		DependencyIndexes: file_CalcWeaponUpgradeReturnItemsReq_proto_depIdxs,
		EnumInfos:         file_CalcWeaponUpgradeReturnItemsReq_proto_enumTypes,
		MessageInfos:      file_CalcWeaponUpgradeReturnItemsReq_proto_msgTypes,
	}.Build()
	File_CalcWeaponUpgradeReturnItemsReq_proto = out.File
	file_CalcWeaponUpgradeReturnItemsReq_proto_rawDesc = nil
	file_CalcWeaponUpgradeReturnItemsReq_proto_goTypes = nil
	file_CalcWeaponUpgradeReturnItemsReq_proto_depIdxs = nil
}
