// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AbilitySyncStateInfo.proto

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

type AbilitySyncStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsInited           bool                       `protobuf:"varint,1,opt,name=is_inited,json=isInited,proto3" json:"is_inited,omitempty"`
	DynamicValueMap    []*AbilityScalarValueEntry `protobuf:"bytes,2,rep,name=dynamic_value_map,json=dynamicValueMap,proto3" json:"dynamic_value_map,omitempty"`
	AppliedAbilities   []*AbilityAppliedAbility   `protobuf:"bytes,3,rep,name=applied_abilities,json=appliedAbilities,proto3" json:"applied_abilities,omitempty"`
	AppliedModifiers   []*AbilityAppliedModifier  `protobuf:"bytes,4,rep,name=applied_modifiers,json=appliedModifiers,proto3" json:"applied_modifiers,omitempty"`
	MixinRecoverInfos  []*AbilityMixinRecoverInfo `protobuf:"bytes,5,rep,name=mixin_recover_infos,json=mixinRecoverInfos,proto3" json:"mixin_recover_infos,omitempty"`
	SgvDynamicValueMap []*AbilityScalarValueEntry `protobuf:"bytes,6,rep,name=sgv_dynamic_value_map,json=sgvDynamicValueMap,proto3" json:"sgv_dynamic_value_map,omitempty"`
}

func (x *AbilitySyncStateInfo) Reset() {
	*x = AbilitySyncStateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilitySyncStateInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilitySyncStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilitySyncStateInfo) ProtoMessage() {}

func (x *AbilitySyncStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AbilitySyncStateInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilitySyncStateInfo.ProtoReflect.Descriptor instead.
func (*AbilitySyncStateInfo) Descriptor() ([]byte, []int) {
	return file_AbilitySyncStateInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AbilitySyncStateInfo) GetIsInited() bool {
	if x != nil {
		return x.IsInited
	}
	return false
}

func (x *AbilitySyncStateInfo) GetDynamicValueMap() []*AbilityScalarValueEntry {
	if x != nil {
		return x.DynamicValueMap
	}
	return nil
}

func (x *AbilitySyncStateInfo) GetAppliedAbilities() []*AbilityAppliedAbility {
	if x != nil {
		return x.AppliedAbilities
	}
	return nil
}

func (x *AbilitySyncStateInfo) GetAppliedModifiers() []*AbilityAppliedModifier {
	if x != nil {
		return x.AppliedModifiers
	}
	return nil
}

func (x *AbilitySyncStateInfo) GetMixinRecoverInfos() []*AbilityMixinRecoverInfo {
	if x != nil {
		return x.MixinRecoverInfos
	}
	return nil
}

func (x *AbilitySyncStateInfo) GetSgvDynamicValueMap() []*AbilityScalarValueEntry {
	if x != nil {
		return x.SgvDynamicValueMap
	}
	return nil
}

var File_AbilitySyncStateInfo_proto protoreflect.FileDescriptor

var file_AbilitySyncStateInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x64, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x03, 0x0a,
	0x14, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x49, 0x6e, 0x69, 0x74,
	0x65, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x49,
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x11, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x13, 0x6d, 0x69, 0x78, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x11, 0x6d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x51, 0x0a, 0x15, 0x73, 0x67, 0x76, 0x5f, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x73, 0x67, 0x76, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilitySyncStateInfo_proto_rawDescOnce sync.Once
	file_AbilitySyncStateInfo_proto_rawDescData = file_AbilitySyncStateInfo_proto_rawDesc
)

func file_AbilitySyncStateInfo_proto_rawDescGZIP() []byte {
	file_AbilitySyncStateInfo_proto_rawDescOnce.Do(func() {
		file_AbilitySyncStateInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilitySyncStateInfo_proto_rawDescData)
	})
	return file_AbilitySyncStateInfo_proto_rawDescData
}

var file_AbilitySyncStateInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilitySyncStateInfo_proto_goTypes = []interface{}{
	(*AbilitySyncStateInfo)(nil),    // 0: proto.AbilitySyncStateInfo
	(*AbilityScalarValueEntry)(nil), // 1: proto.AbilityScalarValueEntry
	(*AbilityAppliedAbility)(nil),   // 2: proto.AbilityAppliedAbility
	(*AbilityAppliedModifier)(nil),  // 3: proto.AbilityAppliedModifier
	(*AbilityMixinRecoverInfo)(nil), // 4: proto.AbilityMixinRecoverInfo
}
var file_AbilitySyncStateInfo_proto_depIdxs = []int32{
	1, // 0: proto.AbilitySyncStateInfo.dynamic_value_map:type_name -> proto.AbilityScalarValueEntry
	2, // 1: proto.AbilitySyncStateInfo.applied_abilities:type_name -> proto.AbilityAppliedAbility
	3, // 2: proto.AbilitySyncStateInfo.applied_modifiers:type_name -> proto.AbilityAppliedModifier
	4, // 3: proto.AbilitySyncStateInfo.mixin_recover_infos:type_name -> proto.AbilityMixinRecoverInfo
	1, // 4: proto.AbilitySyncStateInfo.sgv_dynamic_value_map:type_name -> proto.AbilityScalarValueEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_AbilitySyncStateInfo_proto_init() }
func file_AbilitySyncStateInfo_proto_init() {
	if File_AbilitySyncStateInfo_proto != nil {
		return
	}
	file_AbilityAppliedAbility_proto_init()
	file_AbilityAppliedModifier_proto_init()
	file_AbilityMixinRecoverInfo_proto_init()
	file_AbilityScalarValueEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilitySyncStateInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilitySyncStateInfo); i {
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
			RawDescriptor: file_AbilitySyncStateInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilitySyncStateInfo_proto_goTypes,
		DependencyIndexes: file_AbilitySyncStateInfo_proto_depIdxs,
		MessageInfos:      file_AbilitySyncStateInfo_proto_msgTypes,
	}.Build()
	File_AbilitySyncStateInfo_proto = out.File
	file_AbilitySyncStateInfo_proto_rawDesc = nil
	file_AbilitySyncStateInfo_proto_goTypes = nil
	file_AbilitySyncStateInfo_proto_depIdxs = nil
}
