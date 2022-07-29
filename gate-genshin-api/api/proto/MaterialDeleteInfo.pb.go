// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: MaterialDeleteInfo.proto

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

type MaterialDeleteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to DeleteInfo:
	//	*MaterialDeleteInfo_CountDownDelete_
	//	*MaterialDeleteInfo_DateDelete
	//	*MaterialDeleteInfo_DelayWeekCountDownDelete_
	DeleteInfo      isMaterialDeleteInfo_DeleteInfo `protobuf_oneof:"DeleteInfo"`
	HasDeleteConfig bool                            `protobuf:"varint,1,opt,name=has_delete_config,json=hasDeleteConfig,proto3" json:"has_delete_config,omitempty"`
}

func (x *MaterialDeleteInfo) Reset() {
	*x = MaterialDeleteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaterialDeleteInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialDeleteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialDeleteInfo) ProtoMessage() {}

func (x *MaterialDeleteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MaterialDeleteInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialDeleteInfo.ProtoReflect.Descriptor instead.
func (*MaterialDeleteInfo) Descriptor() ([]byte, []int) {
	return file_MaterialDeleteInfo_proto_rawDescGZIP(), []int{0}
}

func (m *MaterialDeleteInfo) GetDeleteInfo() isMaterialDeleteInfo_DeleteInfo {
	if m != nil {
		return m.DeleteInfo
	}
	return nil
}

func (x *MaterialDeleteInfo) GetCountDownDelete() *MaterialDeleteInfo_CountDownDelete {
	if x, ok := x.GetDeleteInfo().(*MaterialDeleteInfo_CountDownDelete_); ok {
		return x.CountDownDelete
	}
	return nil
}

func (x *MaterialDeleteInfo) GetDateDelete() *MaterialDeleteInfo_DateTimeDelete {
	if x, ok := x.GetDeleteInfo().(*MaterialDeleteInfo_DateDelete); ok {
		return x.DateDelete
	}
	return nil
}

func (x *MaterialDeleteInfo) GetDelayWeekCountDownDelete() *MaterialDeleteInfo_DelayWeekCountDownDelete {
	if x, ok := x.GetDeleteInfo().(*MaterialDeleteInfo_DelayWeekCountDownDelete_); ok {
		return x.DelayWeekCountDownDelete
	}
	return nil
}

func (x *MaterialDeleteInfo) GetHasDeleteConfig() bool {
	if x != nil {
		return x.HasDeleteConfig
	}
	return false
}

type isMaterialDeleteInfo_DeleteInfo interface {
	isMaterialDeleteInfo_DeleteInfo()
}

type MaterialDeleteInfo_CountDownDelete_ struct {
	CountDownDelete *MaterialDeleteInfo_CountDownDelete `protobuf:"bytes,2,opt,name=count_down_delete,json=countDownDelete,proto3,oneof"`
}

type MaterialDeleteInfo_DateDelete struct {
	DateDelete *MaterialDeleteInfo_DateTimeDelete `protobuf:"bytes,3,opt,name=date_delete,json=dateDelete,proto3,oneof"`
}

type MaterialDeleteInfo_DelayWeekCountDownDelete_ struct {
	DelayWeekCountDownDelete *MaterialDeleteInfo_DelayWeekCountDownDelete `protobuf:"bytes,4,opt,name=delay_week_count_down_delete,json=delayWeekCountDownDelete,proto3,oneof"`
}

func (*MaterialDeleteInfo_CountDownDelete_) isMaterialDeleteInfo_DeleteInfo() {}

func (*MaterialDeleteInfo_DateDelete) isMaterialDeleteInfo_DeleteInfo() {}

func (*MaterialDeleteInfo_DelayWeekCountDownDelete_) isMaterialDeleteInfo_DeleteInfo() {}

type MaterialDeleteInfo_CountDownDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteTimeNumMap    map[uint32]uint32 `protobuf:"bytes,1,rep,name=delete_time_num_map,json=deleteTimeNumMap,proto3" json:"delete_time_num_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ConfigCountDownTime uint32            `protobuf:"varint,2,opt,name=config_count_down_time,json=configCountDownTime,proto3" json:"config_count_down_time,omitempty"`
}

func (x *MaterialDeleteInfo_CountDownDelete) Reset() {
	*x = MaterialDeleteInfo_CountDownDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaterialDeleteInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialDeleteInfo_CountDownDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialDeleteInfo_CountDownDelete) ProtoMessage() {}

func (x *MaterialDeleteInfo_CountDownDelete) ProtoReflect() protoreflect.Message {
	mi := &file_MaterialDeleteInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialDeleteInfo_CountDownDelete.ProtoReflect.Descriptor instead.
func (*MaterialDeleteInfo_CountDownDelete) Descriptor() ([]byte, []int) {
	return file_MaterialDeleteInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MaterialDeleteInfo_CountDownDelete) GetDeleteTimeNumMap() map[uint32]uint32 {
	if x != nil {
		return x.DeleteTimeNumMap
	}
	return nil
}

func (x *MaterialDeleteInfo_CountDownDelete) GetConfigCountDownTime() uint32 {
	if x != nil {
		return x.ConfigCountDownTime
	}
	return 0
}

type MaterialDeleteInfo_DateTimeDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteTime uint32 `protobuf:"varint,1,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
}

func (x *MaterialDeleteInfo_DateTimeDelete) Reset() {
	*x = MaterialDeleteInfo_DateTimeDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaterialDeleteInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialDeleteInfo_DateTimeDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialDeleteInfo_DateTimeDelete) ProtoMessage() {}

func (x *MaterialDeleteInfo_DateTimeDelete) ProtoReflect() protoreflect.Message {
	mi := &file_MaterialDeleteInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialDeleteInfo_DateTimeDelete.ProtoReflect.Descriptor instead.
func (*MaterialDeleteInfo_DateTimeDelete) Descriptor() ([]byte, []int) {
	return file_MaterialDeleteInfo_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MaterialDeleteInfo_DateTimeDelete) GetDeleteTime() uint32 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

type MaterialDeleteInfo_DelayWeekCountDownDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteTimeNumMap    map[uint32]uint32 `protobuf:"bytes,1,rep,name=delete_time_num_map,json=deleteTimeNumMap,proto3" json:"delete_time_num_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ConfigDelayWeek     uint32            `protobuf:"varint,2,opt,name=config_delay_week,json=configDelayWeek,proto3" json:"config_delay_week,omitempty"`
	ConfigCountDownTime uint32            `protobuf:"varint,3,opt,name=config_count_down_time,json=configCountDownTime,proto3" json:"config_count_down_time,omitempty"`
}

func (x *MaterialDeleteInfo_DelayWeekCountDownDelete) Reset() {
	*x = MaterialDeleteInfo_DelayWeekCountDownDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaterialDeleteInfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialDeleteInfo_DelayWeekCountDownDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialDeleteInfo_DelayWeekCountDownDelete) ProtoMessage() {}

func (x *MaterialDeleteInfo_DelayWeekCountDownDelete) ProtoReflect() protoreflect.Message {
	mi := &file_MaterialDeleteInfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialDeleteInfo_DelayWeekCountDownDelete.ProtoReflect.Descriptor instead.
func (*MaterialDeleteInfo_DelayWeekCountDownDelete) Descriptor() ([]byte, []int) {
	return file_MaterialDeleteInfo_proto_rawDescGZIP(), []int{0, 2}
}

func (x *MaterialDeleteInfo_DelayWeekCountDownDelete) GetDeleteTimeNumMap() map[uint32]uint32 {
	if x != nil {
		return x.DeleteTimeNumMap
	}
	return nil
}

func (x *MaterialDeleteInfo_DelayWeekCountDownDelete) GetConfigDelayWeek() uint32 {
	if x != nil {
		return x.ConfigDelayWeek
	}
	return 0
}

func (x *MaterialDeleteInfo_DelayWeekCountDownDelete) GetConfigCountDownTime() uint32 {
	if x != nil {
		return x.ConfigCountDownTime
	}
	return 0
}

var File_MaterialDeleteInfo_proto protoreflect.FileDescriptor

var file_MaterialDeleteInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd7, 0x07, 0x0a, 0x12, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x57, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x74,
	0x0a, 0x1c, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x57, 0x65, 0x65, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f,
	0x77, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x18, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x57, 0x65, 0x65, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x68, 0x61, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0xfb, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x6e, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x4d, 0x61, 0x70, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x43, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x31,
	0x0a, 0x0e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x1a, 0xb9, 0x02, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x57, 0x65, 0x65, 0x6b, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x77,
	0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x57, 0x65, 0x65, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x43, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MaterialDeleteInfo_proto_rawDescOnce sync.Once
	file_MaterialDeleteInfo_proto_rawDescData = file_MaterialDeleteInfo_proto_rawDesc
)

func file_MaterialDeleteInfo_proto_rawDescGZIP() []byte {
	file_MaterialDeleteInfo_proto_rawDescOnce.Do(func() {
		file_MaterialDeleteInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MaterialDeleteInfo_proto_rawDescData)
	})
	return file_MaterialDeleteInfo_proto_rawDescData
}

var file_MaterialDeleteInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_MaterialDeleteInfo_proto_goTypes = []interface{}{
	(*MaterialDeleteInfo)(nil),                          // 0: proto.MaterialDeleteInfo
	(*MaterialDeleteInfo_CountDownDelete)(nil),          // 1: proto.MaterialDeleteInfo.CountDownDelete
	(*MaterialDeleteInfo_DateTimeDelete)(nil),           // 2: proto.MaterialDeleteInfo.DateTimeDelete
	(*MaterialDeleteInfo_DelayWeekCountDownDelete)(nil), // 3: proto.MaterialDeleteInfo.DelayWeekCountDownDelete
	nil, // 4: proto.MaterialDeleteInfo.CountDownDelete.DeleteTimeNumMapEntry
	nil, // 5: proto.MaterialDeleteInfo.DelayWeekCountDownDelete.DeleteTimeNumMapEntry
}
var file_MaterialDeleteInfo_proto_depIdxs = []int32{
	1, // 0: proto.MaterialDeleteInfo.count_down_delete:type_name -> proto.MaterialDeleteInfo.CountDownDelete
	2, // 1: proto.MaterialDeleteInfo.date_delete:type_name -> proto.MaterialDeleteInfo.DateTimeDelete
	3, // 2: proto.MaterialDeleteInfo.delay_week_count_down_delete:type_name -> proto.MaterialDeleteInfo.DelayWeekCountDownDelete
	4, // 3: proto.MaterialDeleteInfo.CountDownDelete.delete_time_num_map:type_name -> proto.MaterialDeleteInfo.CountDownDelete.DeleteTimeNumMapEntry
	5, // 4: proto.MaterialDeleteInfo.DelayWeekCountDownDelete.delete_time_num_map:type_name -> proto.MaterialDeleteInfo.DelayWeekCountDownDelete.DeleteTimeNumMapEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_MaterialDeleteInfo_proto_init() }
func file_MaterialDeleteInfo_proto_init() {
	if File_MaterialDeleteInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MaterialDeleteInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialDeleteInfo); i {
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
		file_MaterialDeleteInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialDeleteInfo_CountDownDelete); i {
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
		file_MaterialDeleteInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialDeleteInfo_DateTimeDelete); i {
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
		file_MaterialDeleteInfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialDeleteInfo_DelayWeekCountDownDelete); i {
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
	file_MaterialDeleteInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MaterialDeleteInfo_CountDownDelete_)(nil),
		(*MaterialDeleteInfo_DateDelete)(nil),
		(*MaterialDeleteInfo_DelayWeekCountDownDelete_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MaterialDeleteInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MaterialDeleteInfo_proto_goTypes,
		DependencyIndexes: file_MaterialDeleteInfo_proto_depIdxs,
		MessageInfos:      file_MaterialDeleteInfo_proto_msgTypes,
	}.Build()
	File_MaterialDeleteInfo_proto = out.File
	file_MaterialDeleteInfo_proto_rawDesc = nil
	file_MaterialDeleteInfo_proto_goTypes = nil
	file_MaterialDeleteInfo_proto_depIdxs = nil
}
