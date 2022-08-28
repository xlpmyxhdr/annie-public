// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AbilityAppliedModifier.proto

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

type AbilityAppliedModifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifierLocalId           int32                    `protobuf:"varint,1,opt,name=modifier_local_id,json=modifierLocalId,proto3" json:"modifier_local_id,omitempty"`
	ParentAbilityEntityId     uint32                   `protobuf:"varint,2,opt,name=parent_ability_entity_id,json=parentAbilityEntityId,proto3" json:"parent_ability_entity_id,omitempty"`
	ParentAbilityName         *AbilityString           `protobuf:"bytes,3,opt,name=parent_ability_name,json=parentAbilityName,proto3" json:"parent_ability_name,omitempty"`
	ParentAbilityOverride     *AbilityString           `protobuf:"bytes,4,opt,name=parent_ability_override,json=parentAbilityOverride,proto3" json:"parent_ability_override,omitempty"`
	InstancedAbilityId        uint32                   `protobuf:"varint,5,opt,name=instanced_ability_id,json=instancedAbilityId,proto3" json:"instanced_ability_id,omitempty"`
	InstancedModifierId       uint32                   `protobuf:"varint,6,opt,name=instanced_modifier_id,json=instancedModifierId,proto3" json:"instanced_modifier_id,omitempty"`
	ExistDuration             float32                  `protobuf:"fixed32,7,opt,name=exist_duration,json=existDuration,proto3" json:"exist_duration,omitempty"`
	AttachedInstancedModifier *AbilityAttachedModifier `protobuf:"bytes,8,opt,name=attached_instanced_modifier,json=attachedInstancedModifier,proto3" json:"attached_instanced_modifier,omitempty"`
	ApplyEntityId             uint32                   `protobuf:"varint,9,opt,name=apply_entity_id,json=applyEntityId,proto3" json:"apply_entity_id,omitempty"`
	IsAttachedParentAbility   bool                     `protobuf:"varint,10,opt,name=is_attached_parent_ability,json=isAttachedParentAbility,proto3" json:"is_attached_parent_ability,omitempty"`
	ModifierDurability        *ModifierDurability      `protobuf:"bytes,11,opt,name=modifier_durability,json=modifierDurability,proto3" json:"modifier_durability,omitempty"`
	SbuffUid                  uint32                   `protobuf:"varint,12,opt,name=sbuff_uid,json=sbuffUid,proto3" json:"sbuff_uid,omitempty"`
	IsServerbuffModifier      bool                     `protobuf:"varint,13,opt,name=is_serverbuff_modifier,json=isServerbuffModifier,proto3" json:"is_serverbuff_modifier,omitempty"`
}

func (x *AbilityAppliedModifier) Reset() {
	*x = AbilityAppliedModifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityAppliedModifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityAppliedModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityAppliedModifier) ProtoMessage() {}

func (x *AbilityAppliedModifier) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityAppliedModifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityAppliedModifier.ProtoReflect.Descriptor instead.
func (*AbilityAppliedModifier) Descriptor() ([]byte, []int) {
	return file_AbilityAppliedModifier_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityAppliedModifier) GetModifierLocalId() int32 {
	if x != nil {
		return x.ModifierLocalId
	}
	return 0
}

func (x *AbilityAppliedModifier) GetParentAbilityEntityId() uint32 {
	if x != nil {
		return x.ParentAbilityEntityId
	}
	return 0
}

func (x *AbilityAppliedModifier) GetParentAbilityName() *AbilityString {
	if x != nil {
		return x.ParentAbilityName
	}
	return nil
}

func (x *AbilityAppliedModifier) GetParentAbilityOverride() *AbilityString {
	if x != nil {
		return x.ParentAbilityOverride
	}
	return nil
}

func (x *AbilityAppliedModifier) GetInstancedAbilityId() uint32 {
	if x != nil {
		return x.InstancedAbilityId
	}
	return 0
}

func (x *AbilityAppliedModifier) GetInstancedModifierId() uint32 {
	if x != nil {
		return x.InstancedModifierId
	}
	return 0
}

func (x *AbilityAppliedModifier) GetExistDuration() float32 {
	if x != nil {
		return x.ExistDuration
	}
	return 0
}

func (x *AbilityAppliedModifier) GetAttachedInstancedModifier() *AbilityAttachedModifier {
	if x != nil {
		return x.AttachedInstancedModifier
	}
	return nil
}

func (x *AbilityAppliedModifier) GetApplyEntityId() uint32 {
	if x != nil {
		return x.ApplyEntityId
	}
	return 0
}

func (x *AbilityAppliedModifier) GetIsAttachedParentAbility() bool {
	if x != nil {
		return x.IsAttachedParentAbility
	}
	return false
}

func (x *AbilityAppliedModifier) GetModifierDurability() *ModifierDurability {
	if x != nil {
		return x.ModifierDurability
	}
	return nil
}

func (x *AbilityAppliedModifier) GetSbuffUid() uint32 {
	if x != nil {
		return x.SbuffUid
	}
	return 0
}

func (x *AbilityAppliedModifier) GetIsServerbuffModifier() bool {
	if x != nil {
		return x.IsServerbuffModifier
	}
	return false
}

var File_AbilityAppliedModifier_proto protoreflect.FileDescriptor

var file_AbilityAppliedModifier_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x05, 0x0a,
	0x16, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x13,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x17,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x15, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x65, 0x78, 0x69, 0x73, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x58, 0x0a, 0x1b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x19, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69, 0x73, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x44, 0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x62, 0x75, 0x66, 0x66,
	0x55, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x62, 0x75, 0x66, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x75, 0x66,
	0x66, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityAppliedModifier_proto_rawDescOnce sync.Once
	file_AbilityAppliedModifier_proto_rawDescData = file_AbilityAppliedModifier_proto_rawDesc
)

func file_AbilityAppliedModifier_proto_rawDescGZIP() []byte {
	file_AbilityAppliedModifier_proto_rawDescOnce.Do(func() {
		file_AbilityAppliedModifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityAppliedModifier_proto_rawDescData)
	})
	return file_AbilityAppliedModifier_proto_rawDescData
}

var file_AbilityAppliedModifier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityAppliedModifier_proto_goTypes = []interface{}{
	(*AbilityAppliedModifier)(nil),  // 0: AbilityAppliedModifier
	(*AbilityString)(nil),           // 1: AbilityString
	(*AbilityAttachedModifier)(nil), // 2: AbilityAttachedModifier
	(*ModifierDurability)(nil),      // 3: ModifierDurability
}
var file_AbilityAppliedModifier_proto_depIdxs = []int32{
	1, // 0: AbilityAppliedModifier.parent_ability_name:type_name -> AbilityString
	1, // 1: AbilityAppliedModifier.parent_ability_override:type_name -> AbilityString
	2, // 2: AbilityAppliedModifier.attached_instanced_modifier:type_name -> AbilityAttachedModifier
	3, // 3: AbilityAppliedModifier.modifier_durability:type_name -> ModifierDurability
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AbilityAppliedModifier_proto_init() }
func file_AbilityAppliedModifier_proto_init() {
	if File_AbilityAppliedModifier_proto != nil {
		return
	}
	file_AbilityAttachedModifier_proto_init()
	file_AbilityString_proto_init()
	file_ModifierDurability_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityAppliedModifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityAppliedModifier); i {
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
			RawDescriptor: file_AbilityAppliedModifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityAppliedModifier_proto_goTypes,
		DependencyIndexes: file_AbilityAppliedModifier_proto_depIdxs,
		MessageInfos:      file_AbilityAppliedModifier_proto_msgTypes,
	}.Build()
	File_AbilityAppliedModifier_proto = out.File
	file_AbilityAppliedModifier_proto_rawDesc = nil
	file_AbilityAppliedModifier_proto_goTypes = nil
	file_AbilityAppliedModifier_proto_depIdxs = nil
}
