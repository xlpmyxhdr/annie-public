// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: AvatarLifeStateChangeNotify.proto

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

// CmdId: 1235
// EnetChannelId: 0
// EnetIsReliable: true
type AvatarLifeStateChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid      uint64        `protobuf:"varint,8,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	LifeState       uint32        `protobuf:"varint,15,opt,name=life_state,json=lifeState,proto3" json:"life_state,omitempty"`
	SourceEntityId  uint32        `protobuf:"varint,1,opt,name=source_entity_id,json=sourceEntityId,proto3" json:"source_entity_id,omitempty"`
	AttackTag       string        `protobuf:"bytes,10,opt,name=attack_tag,json=attackTag,proto3" json:"attack_tag,omitempty"`
	DieType         PlayerDieType `protobuf:"varint,11,opt,name=die_type,json=dieType,proto3,enum=proto.PlayerDieType" json:"die_type,omitempty"`
	MoveReliableSeq uint32        `protobuf:"varint,7,opt,name=move_reliable_seq,json=moveReliableSeq,proto3" json:"move_reliable_seq,omitempty"`
	ServerBuffList  []*ServerBuff `protobuf:"bytes,6,rep,name=server_buff_list,json=serverBuffList,proto3" json:"server_buff_list,omitempty"`
}

func (x *AvatarLifeStateChangeNotify) Reset() {
	*x = AvatarLifeStateChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarLifeStateChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarLifeStateChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarLifeStateChangeNotify) ProtoMessage() {}

func (x *AvatarLifeStateChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarLifeStateChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarLifeStateChangeNotify.ProtoReflect.Descriptor instead.
func (*AvatarLifeStateChangeNotify) Descriptor() ([]byte, []int) {
	return file_AvatarLifeStateChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarLifeStateChangeNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarLifeStateChangeNotify) GetLifeState() uint32 {
	if x != nil {
		return x.LifeState
	}
	return 0
}

func (x *AvatarLifeStateChangeNotify) GetSourceEntityId() uint32 {
	if x != nil {
		return x.SourceEntityId
	}
	return 0
}

func (x *AvatarLifeStateChangeNotify) GetAttackTag() string {
	if x != nil {
		return x.AttackTag
	}
	return ""
}

func (x *AvatarLifeStateChangeNotify) GetDieType() PlayerDieType {
	if x != nil {
		return x.DieType
	}
	return PlayerDieType_PLAYER_DIE_TYPE_NONE
}

func (x *AvatarLifeStateChangeNotify) GetMoveReliableSeq() uint32 {
	if x != nil {
		return x.MoveReliableSeq
	}
	return 0
}

func (x *AvatarLifeStateChangeNotify) GetServerBuffList() []*ServerBuff {
	if x != nil {
		return x.ServerBuffList
	}
	return nil
}

var File_AvatarLifeStateChangeNotify_proto protoreflect.FileDescriptor

var file_AvatarLifeStateChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x66, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x67, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x69,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x64, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x6c, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x71, 0x12, 0x3b, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x42, 0x75, 0x66, 0x66, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarLifeStateChangeNotify_proto_rawDescOnce sync.Once
	file_AvatarLifeStateChangeNotify_proto_rawDescData = file_AvatarLifeStateChangeNotify_proto_rawDesc
)

func file_AvatarLifeStateChangeNotify_proto_rawDescGZIP() []byte {
	file_AvatarLifeStateChangeNotify_proto_rawDescOnce.Do(func() {
		file_AvatarLifeStateChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarLifeStateChangeNotify_proto_rawDescData)
	})
	return file_AvatarLifeStateChangeNotify_proto_rawDescData
}

var file_AvatarLifeStateChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarLifeStateChangeNotify_proto_goTypes = []interface{}{
	(*AvatarLifeStateChangeNotify)(nil), // 0: proto.AvatarLifeStateChangeNotify
	(PlayerDieType)(0),                  // 1: proto.PlayerDieType
	(*ServerBuff)(nil),                  // 2: proto.ServerBuff
}
var file_AvatarLifeStateChangeNotify_proto_depIdxs = []int32{
	1, // 0: proto.AvatarLifeStateChangeNotify.die_type:type_name -> proto.PlayerDieType
	2, // 1: proto.AvatarLifeStateChangeNotify.server_buff_list:type_name -> proto.ServerBuff
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarLifeStateChangeNotify_proto_init() }
func file_AvatarLifeStateChangeNotify_proto_init() {
	if File_AvatarLifeStateChangeNotify_proto != nil {
		return
	}
	file_PlayerDieType_proto_init()
	file_ServerBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarLifeStateChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarLifeStateChangeNotify); i {
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
			RawDescriptor: file_AvatarLifeStateChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarLifeStateChangeNotify_proto_goTypes,
		DependencyIndexes: file_AvatarLifeStateChangeNotify_proto_depIdxs,
		MessageInfos:      file_AvatarLifeStateChangeNotify_proto_msgTypes,
	}.Build()
	File_AvatarLifeStateChangeNotify_proto = out.File
	file_AvatarLifeStateChangeNotify_proto_rawDesc = nil
	file_AvatarLifeStateChangeNotify_proto_goTypes = nil
	file_AvatarLifeStateChangeNotify_proto_depIdxs = nil
}
