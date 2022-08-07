// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: MusicBeatmapList.proto

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

type MusicBeatmapList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeatmapNoteList []*MusicBeatmapNote `protobuf:"bytes,1,rep,name=beatmap_note_list,json=beatmapNoteList,proto3" json:"beatmap_note_list,omitempty"`
}

func (x *MusicBeatmapList) Reset() {
	*x = MusicBeatmapList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MusicBeatmapList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MusicBeatmapList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MusicBeatmapList) ProtoMessage() {}

func (x *MusicBeatmapList) ProtoReflect() protoreflect.Message {
	mi := &file_MusicBeatmapList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MusicBeatmapList.ProtoReflect.Descriptor instead.
func (*MusicBeatmapList) Descriptor() ([]byte, []int) {
	return file_MusicBeatmapList_proto_rawDescGZIP(), []int{0}
}

func (x *MusicBeatmapList) GetBeatmapNoteList() []*MusicBeatmapNote {
	if x != nil {
		return x.BeatmapNoteList
	}
	return nil
}

var File_MusicBeatmapList_proto protoreflect.FileDescriptor

var file_MusicBeatmapList_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4e, 0x6f, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x10, 0x4d, 0x75, 0x73, 0x69, 0x63,
	0x42, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x11, 0x62,
	0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x42, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x0f, 0x62, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MusicBeatmapList_proto_rawDescOnce sync.Once
	file_MusicBeatmapList_proto_rawDescData = file_MusicBeatmapList_proto_rawDesc
)

func file_MusicBeatmapList_proto_rawDescGZIP() []byte {
	file_MusicBeatmapList_proto_rawDescOnce.Do(func() {
		file_MusicBeatmapList_proto_rawDescData = protoimpl.X.CompressGZIP(file_MusicBeatmapList_proto_rawDescData)
	})
	return file_MusicBeatmapList_proto_rawDescData
}

var file_MusicBeatmapList_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MusicBeatmapList_proto_goTypes = []interface{}{
	(*MusicBeatmapList)(nil), // 0: proto.MusicBeatmapList
	(*MusicBeatmapNote)(nil), // 1: proto.MusicBeatmapNote
}
var file_MusicBeatmapList_proto_depIdxs = []int32{
	1, // 0: proto.MusicBeatmapList.beatmap_note_list:type_name -> proto.MusicBeatmapNote
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MusicBeatmapList_proto_init() }
func file_MusicBeatmapList_proto_init() {
	if File_MusicBeatmapList_proto != nil {
		return
	}
	file_MusicBeatmapNote_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MusicBeatmapList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MusicBeatmapList); i {
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
			RawDescriptor: file_MusicBeatmapList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MusicBeatmapList_proto_goTypes,
		DependencyIndexes: file_MusicBeatmapList_proto_depIdxs,
		MessageInfos:      file_MusicBeatmapList_proto_msgTypes,
	}.Build()
	File_MusicBeatmapList_proto = out.File
	file_MusicBeatmapList_proto_rawDesc = nil
	file_MusicBeatmapList_proto_goTypes = nil
	file_MusicBeatmapList_proto_depIdxs = nil
}
