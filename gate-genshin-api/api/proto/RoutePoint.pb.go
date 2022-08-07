// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: RoutePoint.proto

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

type RoutePoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position    *Vector `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	ArriveRange float32 `protobuf:"fixed32,2,opt,name=arrive_range,json=arriveRange,proto3" json:"arrive_range,omitempty"`
	// Types that are assignable to MoveParams:
	//	*RoutePoint_Velocity
	//	*RoutePoint_Time
	MoveParams isRoutePoint_MoveParams `protobuf_oneof:"move_params"`
	// Types that are assignable to RotateParams:
	//	*RoutePoint_Rotation
	//	*RoutePoint_RotationSpeed
	//	*RoutePoint_AxisSpeed
	RotateParams isRoutePoint_RotateParams `protobuf_oneof:"rotate_params"`
}

func (x *RoutePoint) Reset() {
	*x = RoutePoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RoutePoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutePoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutePoint) ProtoMessage() {}

func (x *RoutePoint) ProtoReflect() protoreflect.Message {
	mi := &file_RoutePoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutePoint.ProtoReflect.Descriptor instead.
func (*RoutePoint) Descriptor() ([]byte, []int) {
	return file_RoutePoint_proto_rawDescGZIP(), []int{0}
}

func (x *RoutePoint) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *RoutePoint) GetArriveRange() float32 {
	if x != nil {
		return x.ArriveRange
	}
	return 0
}

func (m *RoutePoint) GetMoveParams() isRoutePoint_MoveParams {
	if m != nil {
		return m.MoveParams
	}
	return nil
}

func (x *RoutePoint) GetVelocity() float32 {
	if x, ok := x.GetMoveParams().(*RoutePoint_Velocity); ok {
		return x.Velocity
	}
	return 0
}

func (x *RoutePoint) GetTime() float32 {
	if x, ok := x.GetMoveParams().(*RoutePoint_Time); ok {
		return x.Time
	}
	return 0
}

func (m *RoutePoint) GetRotateParams() isRoutePoint_RotateParams {
	if m != nil {
		return m.RotateParams
	}
	return nil
}

func (x *RoutePoint) GetRotation() *Vector {
	if x, ok := x.GetRotateParams().(*RoutePoint_Rotation); ok {
		return x.Rotation
	}
	return nil
}

func (x *RoutePoint) GetRotationSpeed() *MathQuaternion {
	if x, ok := x.GetRotateParams().(*RoutePoint_RotationSpeed); ok {
		return x.RotationSpeed
	}
	return nil
}

func (x *RoutePoint) GetAxisSpeed() *MathQuaternion {
	if x, ok := x.GetRotateParams().(*RoutePoint_AxisSpeed); ok {
		return x.AxisSpeed
	}
	return nil
}

type isRoutePoint_MoveParams interface {
	isRoutePoint_MoveParams()
}

type RoutePoint_Velocity struct {
	Velocity float32 `protobuf:"fixed32,11,opt,name=velocity,proto3,oneof"`
}

type RoutePoint_Time struct {
	Time float32 `protobuf:"fixed32,12,opt,name=time,proto3,oneof"`
}

func (*RoutePoint_Velocity) isRoutePoint_MoveParams() {}

func (*RoutePoint_Time) isRoutePoint_MoveParams() {}

type isRoutePoint_RotateParams interface {
	isRoutePoint_RotateParams()
}

type RoutePoint_Rotation struct {
	Rotation *Vector `protobuf:"bytes,21,opt,name=rotation,proto3,oneof"`
}

type RoutePoint_RotationSpeed struct {
	RotationSpeed *MathQuaternion `protobuf:"bytes,22,opt,name=rotation_speed,json=rotationSpeed,proto3,oneof"`
}

type RoutePoint_AxisSpeed struct {
	AxisSpeed *MathQuaternion `protobuf:"bytes,23,opt,name=axis_speed,json=axisSpeed,proto3,oneof"`
}

func (*RoutePoint_Rotation) isRoutePoint_RotateParams() {}

func (*RoutePoint_RotationSpeed) isRoutePoint_RotateParams() {}

func (*RoutePoint_AxisSpeed) isRoutePoint_RotateParams() {}

var File_RoutePoint_proto protoreflect.FileDescriptor

var file_RoutePoint_proto_rawDesc = []byte{
	0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61,
	0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02,
	0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x61,
	0x72, 0x72, 0x69, 0x76, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x76, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x08,
	0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48,
	0x01, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0e, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x68,
	0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0d, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x0a, 0x61,
	0x78, 0x69, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61, 0x74,
	0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x09, 0x61, 0x78, 0x69, 0x73, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RoutePoint_proto_rawDescOnce sync.Once
	file_RoutePoint_proto_rawDescData = file_RoutePoint_proto_rawDesc
)

func file_RoutePoint_proto_rawDescGZIP() []byte {
	file_RoutePoint_proto_rawDescOnce.Do(func() {
		file_RoutePoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_RoutePoint_proto_rawDescData)
	})
	return file_RoutePoint_proto_rawDescData
}

var file_RoutePoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RoutePoint_proto_goTypes = []interface{}{
	(*RoutePoint)(nil),     // 0: proto.RoutePoint
	(*Vector)(nil),         // 1: proto.Vector
	(*MathQuaternion)(nil), // 2: proto.MathQuaternion
}
var file_RoutePoint_proto_depIdxs = []int32{
	1, // 0: proto.RoutePoint.position:type_name -> proto.Vector
	1, // 1: proto.RoutePoint.rotation:type_name -> proto.Vector
	2, // 2: proto.RoutePoint.rotation_speed:type_name -> proto.MathQuaternion
	2, // 3: proto.RoutePoint.axis_speed:type_name -> proto.MathQuaternion
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RoutePoint_proto_init() }
func file_RoutePoint_proto_init() {
	if File_RoutePoint_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_MathQuaternion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RoutePoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutePoint); i {
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
	file_RoutePoint_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RoutePoint_Velocity)(nil),
		(*RoutePoint_Time)(nil),
		(*RoutePoint_Rotation)(nil),
		(*RoutePoint_RotationSpeed)(nil),
		(*RoutePoint_AxisSpeed)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RoutePoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RoutePoint_proto_goTypes,
		DependencyIndexes: file_RoutePoint_proto_depIdxs,
		MessageInfos:      file_RoutePoint_proto_msgTypes,
	}.Build()
	File_RoutePoint_proto = out.File
	file_RoutePoint_proto_rawDesc = nil
	file_RoutePoint_proto_goTypes = nil
	file_RoutePoint_proto_depIdxs = nil
}
