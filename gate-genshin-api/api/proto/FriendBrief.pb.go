// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: FriendBrief.proto

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

type FriendBrief struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                   uint32                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname              string                  `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Level                 uint32                  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	AvatarId              uint32                  `protobuf:"varint,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	WorldLevel            uint32                  `protobuf:"varint,5,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	Signature             string                  `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	OnlineState           FriendOnlineState       `protobuf:"varint,7,opt,name=online_state,json=onlineState,proto3,enum=proto.FriendOnlineState" json:"online_state,omitempty"`
	Param                 uint32                  `protobuf:"varint,8,opt,name=param,proto3" json:"param,omitempty"`
	IsMpModeAvailable     bool                    `protobuf:"varint,10,opt,name=is_mp_mode_available,json=isMpModeAvailable,proto3" json:"is_mp_mode_available,omitempty"`
	OnlineId              string                  `protobuf:"bytes,11,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	LastActiveTime        uint32                  `protobuf:"varint,12,opt,name=last_active_time,json=lastActiveTime,proto3" json:"last_active_time,omitempty"`
	NameCardId            uint32                  `protobuf:"varint,13,opt,name=name_card_id,json=nameCardId,proto3" json:"name_card_id,omitempty"`
	MpPlayerNum           uint32                  `protobuf:"varint,14,opt,name=mp_player_num,json=mpPlayerNum,proto3" json:"mp_player_num,omitempty"`
	IsChatNoDisturb       bool                    `protobuf:"varint,15,opt,name=is_chat_no_disturb,json=isChatNoDisturb,proto3" json:"is_chat_no_disturb,omitempty"`
	ChatSequence          uint32                  `protobuf:"varint,16,opt,name=chat_sequence,json=chatSequence,proto3" json:"chat_sequence,omitempty"`
	RemarkName            string                  `protobuf:"bytes,17,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	ShowAvatarInfoList    []*SocialShowAvatarInfo `protobuf:"bytes,22,rep,name=show_avatar_info_list,json=showAvatarInfoList,proto3" json:"show_avatar_info_list,omitempty"`
	FriendEnterHomeOption FriendEnterHomeOption   `protobuf:"varint,23,opt,name=friend_enter_home_option,json=friendEnterHomeOption,proto3,enum=proto.FriendEnterHomeOption" json:"friend_enter_home_option,omitempty"`
	ProfilePicture        *ProfilePicture         `protobuf:"bytes,24,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	IsGameSource          bool                    `protobuf:"varint,25,opt,name=is_game_source,json=isGameSource,proto3" json:"is_game_source,omitempty"`
	IsPsnSource           bool                    `protobuf:"varint,26,opt,name=is_psn_source,json=isPsnSource,proto3" json:"is_psn_source,omitempty"`
	PlatformType          PlatformType            `protobuf:"varint,27,opt,name=platform_type,json=platformType,proto3,enum=proto.PlatformType" json:"platform_type,omitempty"`
}

func (x *FriendBrief) Reset() {
	*x = FriendBrief{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FriendBrief_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendBrief) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendBrief) ProtoMessage() {}

func (x *FriendBrief) ProtoReflect() protoreflect.Message {
	mi := &file_FriendBrief_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendBrief.ProtoReflect.Descriptor instead.
func (*FriendBrief) Descriptor() ([]byte, []int) {
	return file_FriendBrief_proto_rawDescGZIP(), []int{0}
}

func (x *FriendBrief) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FriendBrief) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *FriendBrief) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *FriendBrief) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *FriendBrief) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *FriendBrief) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *FriendBrief) GetOnlineState() FriendOnlineState {
	if x != nil {
		return x.OnlineState
	}
	return FriendOnlineState_FREIEND_DISCONNECT
}

func (x *FriendBrief) GetParam() uint32 {
	if x != nil {
		return x.Param
	}
	return 0
}

func (x *FriendBrief) GetIsMpModeAvailable() bool {
	if x != nil {
		return x.IsMpModeAvailable
	}
	return false
}

func (x *FriendBrief) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *FriendBrief) GetLastActiveTime() uint32 {
	if x != nil {
		return x.LastActiveTime
	}
	return 0
}

func (x *FriendBrief) GetNameCardId() uint32 {
	if x != nil {
		return x.NameCardId
	}
	return 0
}

func (x *FriendBrief) GetMpPlayerNum() uint32 {
	if x != nil {
		return x.MpPlayerNum
	}
	return 0
}

func (x *FriendBrief) GetIsChatNoDisturb() bool {
	if x != nil {
		return x.IsChatNoDisturb
	}
	return false
}

func (x *FriendBrief) GetChatSequence() uint32 {
	if x != nil {
		return x.ChatSequence
	}
	return 0
}

func (x *FriendBrief) GetRemarkName() string {
	if x != nil {
		return x.RemarkName
	}
	return ""
}

func (x *FriendBrief) GetShowAvatarInfoList() []*SocialShowAvatarInfo {
	if x != nil {
		return x.ShowAvatarInfoList
	}
	return nil
}

func (x *FriendBrief) GetFriendEnterHomeOption() FriendEnterHomeOption {
	if x != nil {
		return x.FriendEnterHomeOption
	}
	return FriendEnterHomeOption_NEED_CONFIRM
}

func (x *FriendBrief) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *FriendBrief) GetIsGameSource() bool {
	if x != nil {
		return x.IsGameSource
	}
	return false
}

func (x *FriendBrief) GetIsPsnSource() bool {
	if x != nil {
		return x.IsPsnSource
	}
	return false
}

func (x *FriendBrief) GetPlatformType() PlatformType {
	if x != nil {
		return x.PlatformType
	}
	return PlatformType_EDITOR
}

var File_FriendBrief_proto protoreflect.FileDescriptor

var file_FriendBrief_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x07, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f,
	0x6d, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x4d, 0x70, 0x4d, 0x6f, 0x64,
	0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x70, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6e, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x44, 0x69, 0x73,
	0x74, 0x75, 0x72, 0x62, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x61,
	0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x73, 0x68,
	0x6f, 0x77, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x18, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3e, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x70, 0x73,
	0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x50, 0x73, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FriendBrief_proto_rawDescOnce sync.Once
	file_FriendBrief_proto_rawDescData = file_FriendBrief_proto_rawDesc
)

func file_FriendBrief_proto_rawDescGZIP() []byte {
	file_FriendBrief_proto_rawDescOnce.Do(func() {
		file_FriendBrief_proto_rawDescData = protoimpl.X.CompressGZIP(file_FriendBrief_proto_rawDescData)
	})
	return file_FriendBrief_proto_rawDescData
}

var file_FriendBrief_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FriendBrief_proto_goTypes = []interface{}{
	(*FriendBrief)(nil),          // 0: proto.FriendBrief
	(FriendOnlineState)(0),       // 1: proto.FriendOnlineState
	(*SocialShowAvatarInfo)(nil), // 2: proto.SocialShowAvatarInfo
	(FriendEnterHomeOption)(0),   // 3: proto.FriendEnterHomeOption
	(*ProfilePicture)(nil),       // 4: proto.ProfilePicture
	(PlatformType)(0),            // 5: proto.PlatformType
}
var file_FriendBrief_proto_depIdxs = []int32{
	1, // 0: proto.FriendBrief.online_state:type_name -> proto.FriendOnlineState
	2, // 1: proto.FriendBrief.show_avatar_info_list:type_name -> proto.SocialShowAvatarInfo
	3, // 2: proto.FriendBrief.friend_enter_home_option:type_name -> proto.FriendEnterHomeOption
	4, // 3: proto.FriendBrief.profile_picture:type_name -> proto.ProfilePicture
	5, // 4: proto.FriendBrief.platform_type:type_name -> proto.PlatformType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_FriendBrief_proto_init() }
func file_FriendBrief_proto_init() {
	if File_FriendBrief_proto != nil {
		return
	}
	file_FriendOnlineState_proto_init()
	file_FriendEnterHomeOption_proto_init()
	file_ProfilePicture_proto_init()
	file_PlatformType_proto_init()
	file_SocialShowAvatarInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FriendBrief_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendBrief); i {
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
			RawDescriptor: file_FriendBrief_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FriendBrief_proto_goTypes,
		DependencyIndexes: file_FriendBrief_proto_depIdxs,
		MessageInfos:      file_FriendBrief_proto_msgTypes,
	}.Build()
	File_FriendBrief_proto = out.File
	file_FriendBrief_proto_rawDesc = nil
	file_FriendBrief_proto_goTypes = nil
	file_FriendBrief_proto_depIdxs = nil
}
