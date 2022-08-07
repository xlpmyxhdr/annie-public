// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: ChangeHpReason.proto

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

type ChangeHpReason int32

const (
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_NONE                      ChangeHpReason = 0
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_AVATAR                ChangeHpReason = 1
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_MONSTER               ChangeHpReason = 2
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_GEAR                  ChangeHpReason = 3
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_ENVIR                 ChangeHpReason = 4
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_FALL                  ChangeHpReason = 5
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_DRAWN                 ChangeHpReason = 6
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_ABYSS                 ChangeHpReason = 7
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_ABILITY               ChangeHpReason = 8
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_SUMMON                ChangeHpReason = 9
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_SCRIPT                ChangeHpReason = 10
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_GM                    ChangeHpReason = 11
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_KILL_SELF             ChangeHpReason = 12
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_CLIMATE_COLD          ChangeHpReason = 13
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_STORM_LIGHTNING       ChangeHpReason = 14
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_KILL_SERVER_GADGET    ChangeHpReason = 15
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_REPLACE               ChangeHpReason = 16
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_SUB_PLAYER_LEAVE          ChangeHpReason = 17
	ChangeHpReason_CHANGE_HP_REASON_CIKCDBOJGDK                         ChangeHpReason = 18
	ChangeHpReason_CHANGE_HP_REASON_HEKLBLFBJJK                         ChangeHpReason = 19
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_BY_LUA                    ChangeHpReason = 51
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_ABILITY               ChangeHpReason = 101
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_ITEM                  ChangeHpReason = 102
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_REVIVE                ChangeHpReason = 103
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_UPGRADE               ChangeHpReason = 104
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_STATUE                ChangeHpReason = 105
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_BACKGROUND            ChangeHpReason = 106
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_GM                    ChangeHpReason = 107
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_TRIAL_AVATAR_ACTIVITY ChangeHpReason = 108
	ChangeHpReason_CHANGE_HP_REASON_CHANGE_HP_ADD_ROUGUELIKE_SPRING     ChangeHpReason = 109
)

// Enum value maps for ChangeHpReason.
var (
	ChangeHpReason_name = map[int32]string{
		0:   "CHANGE_HP_REASON_CHANGE_HP_NONE",
		1:   "CHANGE_HP_REASON_CHANGE_HP_SUB_AVATAR",
		2:   "CHANGE_HP_REASON_CHANGE_HP_SUB_MONSTER",
		3:   "CHANGE_HP_REASON_CHANGE_HP_SUB_GEAR",
		4:   "CHANGE_HP_REASON_CHANGE_HP_SUB_ENVIR",
		5:   "CHANGE_HP_REASON_CHANGE_HP_SUB_FALL",
		6:   "CHANGE_HP_REASON_CHANGE_HP_SUB_DRAWN",
		7:   "CHANGE_HP_REASON_CHANGE_HP_SUB_ABYSS",
		8:   "CHANGE_HP_REASON_CHANGE_HP_SUB_ABILITY",
		9:   "CHANGE_HP_REASON_CHANGE_HP_SUB_SUMMON",
		10:  "CHANGE_HP_REASON_CHANGE_HP_SUB_SCRIPT",
		11:  "CHANGE_HP_REASON_CHANGE_HP_SUB_GM",
		12:  "CHANGE_HP_REASON_CHANGE_HP_SUB_KILL_SELF",
		13:  "CHANGE_HP_REASON_CHANGE_HP_SUB_CLIMATE_COLD",
		14:  "CHANGE_HP_REASON_CHANGE_HP_SUB_STORM_LIGHTNING",
		15:  "CHANGE_HP_REASON_CHANGE_HP_SUB_KILL_SERVER_GADGET",
		16:  "CHANGE_HP_REASON_CHANGE_HP_SUB_REPLACE",
		17:  "CHANGE_HP_REASON_CHANGE_HP_SUB_PLAYER_LEAVE",
		18:  "CHANGE_HP_REASON_CIKCDBOJGDK",
		19:  "CHANGE_HP_REASON_HEKLBLFBJJK",
		51:  "CHANGE_HP_REASON_CHANGE_HP_BY_LUA",
		101: "CHANGE_HP_REASON_CHANGE_HP_ADD_ABILITY",
		102: "CHANGE_HP_REASON_CHANGE_HP_ADD_ITEM",
		103: "CHANGE_HP_REASON_CHANGE_HP_ADD_REVIVE",
		104: "CHANGE_HP_REASON_CHANGE_HP_ADD_UPGRADE",
		105: "CHANGE_HP_REASON_CHANGE_HP_ADD_STATUE",
		106: "CHANGE_HP_REASON_CHANGE_HP_ADD_BACKGROUND",
		107: "CHANGE_HP_REASON_CHANGE_HP_ADD_GM",
		108: "CHANGE_HP_REASON_CHANGE_HP_ADD_TRIAL_AVATAR_ACTIVITY",
		109: "CHANGE_HP_REASON_CHANGE_HP_ADD_ROUGUELIKE_SPRING",
	}
	ChangeHpReason_value = map[string]int32{
		"CHANGE_HP_REASON_CHANGE_HP_NONE":                      0,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_AVATAR":                1,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_MONSTER":               2,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_GEAR":                  3,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_ENVIR":                 4,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_FALL":                  5,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_DRAWN":                 6,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_ABYSS":                 7,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_ABILITY":               8,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_SUMMON":                9,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_SCRIPT":                10,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_GM":                    11,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_KILL_SELF":             12,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_CLIMATE_COLD":          13,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_STORM_LIGHTNING":       14,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_KILL_SERVER_GADGET":    15,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_REPLACE":               16,
		"CHANGE_HP_REASON_CHANGE_HP_SUB_PLAYER_LEAVE":          17,
		"CHANGE_HP_REASON_CIKCDBOJGDK":                         18,
		"CHANGE_HP_REASON_HEKLBLFBJJK":                         19,
		"CHANGE_HP_REASON_CHANGE_HP_BY_LUA":                    51,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_ABILITY":               101,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_ITEM":                  102,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_REVIVE":                103,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_UPGRADE":               104,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_STATUE":                105,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_BACKGROUND":            106,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_GM":                    107,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_TRIAL_AVATAR_ACTIVITY": 108,
		"CHANGE_HP_REASON_CHANGE_HP_ADD_ROUGUELIKE_SPRING":     109,
	}
)

func (x ChangeHpReason) Enum() *ChangeHpReason {
	p := new(ChangeHpReason)
	*p = x
	return p
}

func (x ChangeHpReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeHpReason) Descriptor() protoreflect.EnumDescriptor {
	return file_ChangeHpReason_proto_enumTypes[0].Descriptor()
}

func (ChangeHpReason) Type() protoreflect.EnumType {
	return &file_ChangeHpReason_proto_enumTypes[0]
}

func (x ChangeHpReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeHpReason.Descriptor instead.
func (ChangeHpReason) EnumDescriptor() ([]byte, []int) {
	return file_ChangeHpReason_proto_rawDescGZIP(), []int{0}
}

var File_ChangeHpReason_proto protoreflect.FileDescriptor

var file_ChangeHpReason_proto_rawDesc = []byte{
	0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb4, 0x0a,
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x1f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x10, 0x01,
	0x12, 0x2a, 0x0a, 0x26, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53,
	0x55, 0x42, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x47,
	0x45, 0x41, 0x52, 0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x10, 0x04, 0x12,
	0x27, 0x0a, 0x23, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55,
	0x42, 0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x10, 0x05, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x44, 0x52, 0x41, 0x57, 0x4e,
	0x10, 0x06, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50,
	0x5f, 0x53, 0x55, 0x42, 0x5f, 0x41, 0x42, 0x59, 0x53, 0x53, 0x10, 0x07, 0x12, 0x2a, 0x0a, 0x26,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x08, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x4f,
	0x4e, 0x10, 0x09, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48,
	0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x0a, 0x12, 0x25,
	0x0a, 0x21, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42,
	0x5f, 0x47, 0x4d, 0x10, 0x0b, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x45, 0x4c,
	0x46, 0x10, 0x0c, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48,
	0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x43, 0x4c, 0x49, 0x4d, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f,
	0x4c, 0x44, 0x10, 0x0d, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48,
	0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x4d, 0x5f, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0e, 0x12, 0x35, 0x0a, 0x31, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x10, 0x0f, 0x12,
	0x2a, 0x0a, 0x26, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55,
	0x42, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x10, 0x12, 0x2f, 0x0a, 0x2b, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x50, 0x4c,
	0x41, 0x59, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x11, 0x12, 0x20, 0x0a, 0x1c,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x43, 0x49, 0x4b, 0x43, 0x44, 0x42, 0x4f, 0x4a, 0x47, 0x44, 0x4b, 0x10, 0x12, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x48, 0x45, 0x4b, 0x4c, 0x42, 0x4c, 0x46, 0x42, 0x4a, 0x4a, 0x4b, 0x10, 0x13,
	0x12, 0x25, 0x0a, 0x21, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x42,
	0x59, 0x5f, 0x4c, 0x55, 0x41, 0x10, 0x33, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x10, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48,
	0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x66, 0x12, 0x29, 0x0a, 0x25,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x56, 0x45, 0x10, 0x67, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44,
	0x45, 0x10, 0x68, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48,
	0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x45, 0x10, 0x69, 0x12, 0x2d,
	0x0a, 0x29, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x41, 0x44, 0x44,
	0x5f, 0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x6a, 0x12, 0x25, 0x0a,
	0x21, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x47, 0x4d, 0x10, 0x6b, 0x12, 0x38, 0x0a, 0x34, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48,
	0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x48, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x56, 0x41,
	0x54, 0x41, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x6c, 0x12, 0x34,
	0x0a, 0x30, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x48, 0x50, 0x5f, 0x41, 0x44, 0x44,
	0x5f, 0x52, 0x4f, 0x55, 0x47, 0x55, 0x45, 0x4c, 0x49, 0x4b, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x6d, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChangeHpReason_proto_rawDescOnce sync.Once
	file_ChangeHpReason_proto_rawDescData = file_ChangeHpReason_proto_rawDesc
)

func file_ChangeHpReason_proto_rawDescGZIP() []byte {
	file_ChangeHpReason_proto_rawDescOnce.Do(func() {
		file_ChangeHpReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeHpReason_proto_rawDescData)
	})
	return file_ChangeHpReason_proto_rawDescData
}

var file_ChangeHpReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChangeHpReason_proto_goTypes = []interface{}{
	(ChangeHpReason)(0), // 0: proto.ChangeHpReason
}
var file_ChangeHpReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChangeHpReason_proto_init() }
func file_ChangeHpReason_proto_init() {
	if File_ChangeHpReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChangeHpReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeHpReason_proto_goTypes,
		DependencyIndexes: file_ChangeHpReason_proto_depIdxs,
		EnumInfos:         file_ChangeHpReason_proto_enumTypes,
	}.Build()
	File_ChangeHpReason_proto = out.File
	file_ChangeHpReason_proto_rawDesc = nil
	file_ChangeHpReason_proto_goTypes = nil
	file_ChangeHpReason_proto_depIdxs = nil
}
