// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ShowEquip.proto

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

type ShowEquip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Detail:
	//	*ShowEquip_Reliquary
	//	*ShowEquip_Weapon
	Detail isShowEquip_Detail `protobuf_oneof:"detail"`
	ItemId uint32             `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *ShowEquip) Reset() {
	*x = ShowEquip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShowEquip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowEquip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowEquip) ProtoMessage() {}

func (x *ShowEquip) ProtoReflect() protoreflect.Message {
	mi := &file_ShowEquip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowEquip.ProtoReflect.Descriptor instead.
func (*ShowEquip) Descriptor() ([]byte, []int) {
	return file_ShowEquip_proto_rawDescGZIP(), []int{0}
}

func (m *ShowEquip) GetDetail() isShowEquip_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *ShowEquip) GetReliquary() *Reliquary {
	if x, ok := x.GetDetail().(*ShowEquip_Reliquary); ok {
		return x.Reliquary
	}
	return nil
}

func (x *ShowEquip) GetWeapon() *Weapon {
	if x, ok := x.GetDetail().(*ShowEquip_Weapon); ok {
		return x.Weapon
	}
	return nil
}

func (x *ShowEquip) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type isShowEquip_Detail interface {
	isShowEquip_Detail()
}

type ShowEquip_Reliquary struct {
	Reliquary *Reliquary `protobuf:"bytes,2,opt,name=reliquary,proto3,oneof"`
}

type ShowEquip_Weapon struct {
	Weapon *Weapon `protobuf:"bytes,3,opt,name=weapon,proto3,oneof"`
}

func (*ShowEquip_Reliquary) isShowEquip_Detail() {}

func (*ShowEquip_Weapon) isShowEquip_Detail() {}

var File_ShowEquip_proto protoreflect.FileDescriptor

var file_ShowEquip_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x71, 0x75, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75,
	0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x57, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x77,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65,
	0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShowEquip_proto_rawDescOnce sync.Once
	file_ShowEquip_proto_rawDescData = file_ShowEquip_proto_rawDesc
)

func file_ShowEquip_proto_rawDescGZIP() []byte {
	file_ShowEquip_proto_rawDescOnce.Do(func() {
		file_ShowEquip_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShowEquip_proto_rawDescData)
	})
	return file_ShowEquip_proto_rawDescData
}

var file_ShowEquip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ShowEquip_proto_goTypes = []interface{}{
	(*ShowEquip)(nil), // 0: proto.ShowEquip
	(*Reliquary)(nil), // 1: proto.Reliquary
	(*Weapon)(nil),    // 2: proto.Weapon
}
var file_ShowEquip_proto_depIdxs = []int32{
	1, // 0: proto.ShowEquip.reliquary:type_name -> proto.Reliquary
	2, // 1: proto.ShowEquip.weapon:type_name -> proto.Weapon
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ShowEquip_proto_init() }
func file_ShowEquip_proto_init() {
	if File_ShowEquip_proto != nil {
		return
	}
	file_Reliquary_proto_init()
	file_Weapon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ShowEquip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowEquip); i {
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
	file_ShowEquip_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ShowEquip_Reliquary)(nil),
		(*ShowEquip_Weapon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ShowEquip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShowEquip_proto_goTypes,
		DependencyIndexes: file_ShowEquip_proto_depIdxs,
		MessageInfos:      file_ShowEquip_proto_msgTypes,
	}.Build()
	File_ShowEquip_proto = out.File
	file_ShowEquip_proto_rawDesc = nil
	file_ShowEquip_proto_goTypes = nil
	file_ShowEquip_proto_depIdxs = nil
}
