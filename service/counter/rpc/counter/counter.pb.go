// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: counter.proto

package counter

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReqUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId    int64  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`         //事件拥有者id
	EventType  string `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`    //事件类型
	EventField string `protobuf:"bytes,3,opt,name=event_field,json=eventField,proto3" json:"event_field,omitempty"` //事件子域
	Value      int64  `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`                            //事件值
}

func (x *ReqUpdate) Reset() {
	*x = ReqUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdate) ProtoMessage() {}

func (x *ReqUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdate.ProtoReflect.Descriptor instead.
func (*ReqUpdate) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{0}
}

func (x *ReqUpdate) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *ReqUpdate) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *ReqUpdate) GetEventField() string {
	if x != nil {
		return x.EventField
	}
	return ""
}

func (x *ReqUpdate) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ResUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResUpdate) Reset() {
	*x = ResUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResUpdate) ProtoMessage() {}

func (x *ResUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResUpdate.ProtoReflect.Descriptor instead.
func (*ResUpdate) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{1}
}

type ReqQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId       int64  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`                   //事件拥有者id
	EventType     string `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`              //事件类型
	EventField    string `protobuf:"bytes,3,opt,name=event_field,json=eventField,proto3" json:"event_field,omitempty"`           //事件子域
	TimeDimension int32  `protobuf:"varint,4,opt,name=time_dimension,json=timeDimension,proto3" json:"time_dimension,omitempty"` //时间维度
	TimeFlag      int64  `protobuf:"varint,5,opt,name=time_flag,json=timeFlag,proto3" json:"time_flag,omitempty"`                //时间标记
}

func (x *ReqQuery) Reset() {
	*x = ReqQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQuery) ProtoMessage() {}

func (x *ReqQuery) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQuery.ProtoReflect.Descriptor instead.
func (*ReqQuery) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{2}
}

func (x *ReqQuery) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *ReqQuery) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *ReqQuery) GetEventField() string {
	if x != nil {
		return x.EventField
	}
	return ""
}

func (x *ReqQuery) GetTimeDimension() int32 {
	if x != nil {
		return x.TimeDimension
	}
	return 0
}

func (x *ReqQuery) GetTimeFlag() int64 {
	if x != nil {
		return x.TimeFlag
	}
	return 0
}

type ResQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"` //计数
	Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"` //记量
	Max   int64 `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`     //最大值
	Min   int64 `protobuf:"varint,4,opt,name=min,proto3" json:"min,omitempty"`     //最小值
}

func (x *ResQuery) Reset() {
	*x = ResQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResQuery) ProtoMessage() {}

func (x *ResQuery) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResQuery.ProtoReflect.Descriptor instead.
func (*ResQuery) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{3}
}

func (x *ResQuery) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ResQuery) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ResQuery) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *ResQuery) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

var File_counter_proto protoreflect.FileDescriptor

var file_counter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x7c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x22,
	0x5a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x32, 0x6a, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x11, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counter_proto_rawDescOnce sync.Once
	file_counter_proto_rawDescData = file_counter_proto_rawDesc
)

func file_counter_proto_rawDescGZIP() []byte {
	file_counter_proto_rawDescOnce.Do(func() {
		file_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_counter_proto_rawDescData)
	})
	return file_counter_proto_rawDescData
}

var file_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_counter_proto_goTypes = []interface{}{
	(*ReqUpdate)(nil), // 0: counter.ReqUpdate
	(*ResUpdate)(nil), // 1: counter.ResUpdate
	(*ReqQuery)(nil),  // 2: counter.ReqQuery
	(*ResQuery)(nil),  // 3: counter.ResQuery
}
var file_counter_proto_depIdxs = []int32{
	2, // 0: counter.Counter.Query:input_type -> counter.ReqQuery
	0, // 1: counter.Counter.Update:input_type -> counter.ReqUpdate
	3, // 2: counter.Counter.Query:output_type -> counter.ResQuery
	1, // 3: counter.Counter.Update:output_type -> counter.ResUpdate
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_counter_proto_init() }
func file_counter_proto_init() {
	if File_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdate); i {
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
		file_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResUpdate); i {
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
		file_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQuery); i {
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
		file_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResQuery); i {
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
			RawDescriptor: file_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counter_proto_goTypes,
		DependencyIndexes: file_counter_proto_depIdxs,
		MessageInfos:      file_counter_proto_msgTypes,
	}.Build()
	File_counter_proto = out.File
	file_counter_proto_rawDesc = nil
	file_counter_proto_goTypes = nil
	file_counter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterClient interface {
	Query(ctx context.Context, in *ReqQuery, opts ...grpc.CallOption) (*ResQuery, error)
	Update(ctx context.Context, in *ReqUpdate, opts ...grpc.CallOption) (*ResUpdate, error)
}

type counterClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterClient(cc grpc.ClientConnInterface) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) Query(ctx context.Context, in *ReqQuery, opts ...grpc.CallOption) (*ResQuery, error) {
	out := new(ResQuery)
	err := c.cc.Invoke(ctx, "/counter.Counter/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) Update(ctx context.Context, in *ReqUpdate, opts ...grpc.CallOption) (*ResUpdate, error) {
	out := new(ResUpdate)
	err := c.cc.Invoke(ctx, "/counter.Counter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServer is the server API for Counter service.
type CounterServer interface {
	Query(context.Context, *ReqQuery) (*ResQuery, error)
	Update(context.Context, *ReqUpdate) (*ResUpdate, error)
}

// UnimplementedCounterServer can be embedded to have forward compatible implementations.
type UnimplementedCounterServer struct {
}

func (*UnimplementedCounterServer) Query(context.Context, *ReqQuery) (*ResQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedCounterServer) Update(context.Context, *ReqUpdate) (*ResUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterCounterServer(s *grpc.Server, srv CounterServer) {
	s.RegisterService(&_Counter_serviceDesc, srv)
}

func _Counter_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Counter/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).Query(ctx, req.(*ReqQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Counter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).Update(ctx, req.(*ReqUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Counter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "counter.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Counter_Query_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Counter_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
