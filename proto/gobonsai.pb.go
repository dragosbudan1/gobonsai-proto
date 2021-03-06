// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: gobonsai.proto

package gobonsai

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SetPumpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PumpId int32 `protobuf:"varint,1,opt,name=pumpId,proto3" json:"pumpId,omitempty"`
	Active bool  `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *SetPumpRequest) Reset() {
	*x = SetPumpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobonsai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPumpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpRequest) ProtoMessage() {}

func (x *SetPumpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gobonsai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpRequest.ProtoReflect.Descriptor instead.
func (*SetPumpRequest) Descriptor() ([]byte, []int) {
	return file_gobonsai_proto_rawDescGZIP(), []int{0}
}

func (x *SetPumpRequest) GetPumpId() int32 {
	if x != nil {
		return x.PumpId
	}
	return 0
}

func (x *SetPumpRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type SetPumpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PumpId  int32 `protobuf:"varint,1,opt,name=pumpId,proto3" json:"pumpId,omitempty"`
	Active  bool  `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	Success bool  `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SetPumpReply) Reset() {
	*x = SetPumpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobonsai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPumpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpReply) ProtoMessage() {}

func (x *SetPumpReply) ProtoReflect() protoreflect.Message {
	mi := &file_gobonsai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpReply.ProtoReflect.Descriptor instead.
func (*SetPumpReply) Descriptor() ([]byte, []int) {
	return file_gobonsai_proto_rawDescGZIP(), []int{1}
}

func (x *SetPumpReply) GetPumpId() int32 {
	if x != nil {
		return x.PumpId
	}
	return 0
}

func (x *SetPumpReply) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *SetPumpReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PumpData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PumpId int32 `protobuf:"varint,1,opt,name=pumpId,proto3" json:"pumpId,omitempty"`
	Active bool  `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *PumpData) Reset() {
	*x = PumpData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobonsai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PumpData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PumpData) ProtoMessage() {}

func (x *PumpData) ProtoReflect() protoreflect.Message {
	mi := &file_gobonsai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PumpData.ProtoReflect.Descriptor instead.
func (*PumpData) Descriptor() ([]byte, []int) {
	return file_gobonsai_proto_rawDescGZIP(), []int{2}
}

func (x *PumpData) GetPumpId() int32 {
	if x != nil {
		return x.PumpId
	}
	return 0
}

func (x *PumpData) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type GetPumpsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pumps []*PumpData `protobuf:"bytes,1,rep,name=Pumps,json=pumps,proto3" json:"Pumps,omitempty"`
}

func (x *GetPumpsResponse) Reset() {
	*x = GetPumpsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobonsai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPumpsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPumpsResponse) ProtoMessage() {}

func (x *GetPumpsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gobonsai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPumpsResponse.ProtoReflect.Descriptor instead.
func (*GetPumpsResponse) Descriptor() ([]byte, []int) {
	return file_gobonsai_proto_rawDescGZIP(), []int{3}
}

func (x *GetPumpsResponse) GetPumps() []*PumpData {
	if x != nil {
		return x.Pumps
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gobonsai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gobonsai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gobonsai_proto_rawDescGZIP(), []int{4}
}

var File_gobonsai_proto protoreflect.FileDescriptor

var file_gobonsai_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x6f, 0x62, 0x6f, 0x6e, 0x73, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x6f, 0x62, 0x6f, 0x6e, 0x73, 0x61, 0x69, 0x22, 0x40, 0x0a, 0x0e, 0x53, 0x65,
	0x74, 0x50, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x75, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x75,
	0x6d, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x58, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x75, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x75,
	0x6d, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3a, 0x0a, 0x08, 0x50, 0x75, 0x6d, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x75, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x62, 0x6f, 0x6e, 0x73, 0x61, 0x69,
	0x2e, 0x50, 0x75, 0x6d, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x70, 0x75, 0x6d, 0x70, 0x73,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x8b, 0x01, 0x0a, 0x0f, 0x47, 0x6f,
	0x42, 0x6f, 0x6e, 0x73, 0x61, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x12, 0x18, 0x2e, 0x67, 0x6f, 0x62, 0x6f, 0x6e,
	0x73, 0x61, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x62, 0x6f, 0x6e, 0x73, 0x61, 0x69, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x12, 0x0f, 0x2e, 0x67, 0x6f, 0x62, 0x6f, 0x6e,
	0x73, 0x61, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x62, 0x6f,
	0x6e, 0x73, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gobonsai_proto_rawDescOnce sync.Once
	file_gobonsai_proto_rawDescData = file_gobonsai_proto_rawDesc
)

func file_gobonsai_proto_rawDescGZIP() []byte {
	file_gobonsai_proto_rawDescOnce.Do(func() {
		file_gobonsai_proto_rawDescData = protoimpl.X.CompressGZIP(file_gobonsai_proto_rawDescData)
	})
	return file_gobonsai_proto_rawDescData
}

var file_gobonsai_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gobonsai_proto_goTypes = []interface{}{
	(*SetPumpRequest)(nil),   // 0: gobonsai.SetPumpRequest
	(*SetPumpReply)(nil),     // 1: gobonsai.SetPumpReply
	(*PumpData)(nil),         // 2: gobonsai.PumpData
	(*GetPumpsResponse)(nil), // 3: gobonsai.GetPumpsResponse
	(*Empty)(nil),            // 4: gobonsai.Empty
}
var file_gobonsai_proto_depIdxs = []int32{
	2, // 0: gobonsai.GetPumpsResponse.Pumps:type_name -> gobonsai.PumpData
	0, // 1: gobonsai.GoBonsaiService.SetPump:input_type -> gobonsai.SetPumpRequest
	4, // 2: gobonsai.GoBonsaiService.GetPumps:input_type -> gobonsai.Empty
	1, // 3: gobonsai.GoBonsaiService.SetPump:output_type -> gobonsai.SetPumpReply
	3, // 4: gobonsai.GoBonsaiService.GetPumps:output_type -> gobonsai.GetPumpsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gobonsai_proto_init() }
func file_gobonsai_proto_init() {
	if File_gobonsai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gobonsai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPumpRequest); i {
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
		file_gobonsai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPumpReply); i {
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
		file_gobonsai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PumpData); i {
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
		file_gobonsai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPumpsResponse); i {
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
		file_gobonsai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_gobonsai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gobonsai_proto_goTypes,
		DependencyIndexes: file_gobonsai_proto_depIdxs,
		MessageInfos:      file_gobonsai_proto_msgTypes,
	}.Build()
	File_gobonsai_proto = out.File
	file_gobonsai_proto_rawDesc = nil
	file_gobonsai_proto_goTypes = nil
	file_gobonsai_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GoBonsaiServiceClient is the client API for GoBonsaiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoBonsaiServiceClient interface {
	SetPump(ctx context.Context, in *SetPumpRequest, opts ...grpc.CallOption) (*SetPumpReply, error)
	GetPumps(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPumpsResponse, error)
}

type goBonsaiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoBonsaiServiceClient(cc grpc.ClientConnInterface) GoBonsaiServiceClient {
	return &goBonsaiServiceClient{cc}
}

func (c *goBonsaiServiceClient) SetPump(ctx context.Context, in *SetPumpRequest, opts ...grpc.CallOption) (*SetPumpReply, error) {
	out := new(SetPumpReply)
	err := c.cc.Invoke(ctx, "/gobonsai.GoBonsaiService/SetPump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goBonsaiServiceClient) GetPumps(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPumpsResponse, error) {
	out := new(GetPumpsResponse)
	err := c.cc.Invoke(ctx, "/gobonsai.GoBonsaiService/GetPumps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoBonsaiServiceServer is the server API for GoBonsaiService service.
type GoBonsaiServiceServer interface {
	SetPump(context.Context, *SetPumpRequest) (*SetPumpReply, error)
	GetPumps(context.Context, *Empty) (*GetPumpsResponse, error)
}

// UnimplementedGoBonsaiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGoBonsaiServiceServer struct {
}

func (*UnimplementedGoBonsaiServiceServer) SetPump(context.Context, *SetPumpRequest) (*SetPumpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPump not implemented")
}
func (*UnimplementedGoBonsaiServiceServer) GetPumps(context.Context, *Empty) (*GetPumpsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPumps not implemented")
}

func RegisterGoBonsaiServiceServer(s *grpc.Server, srv GoBonsaiServiceServer) {
	s.RegisterService(&_GoBonsaiService_serviceDesc, srv)
}

func _GoBonsaiService_SetPump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoBonsaiServiceServer).SetPump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gobonsai.GoBonsaiService/SetPump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoBonsaiServiceServer).SetPump(ctx, req.(*SetPumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoBonsaiService_GetPumps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoBonsaiServiceServer).GetPumps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gobonsai.GoBonsaiService/GetPumps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoBonsaiServiceServer).GetPumps(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoBonsaiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gobonsai.GoBonsaiService",
	HandlerType: (*GoBonsaiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPump",
			Handler:    _GoBonsaiService_SetPump_Handler,
		},
		{
			MethodName: "GetPumps",
			Handler:    _GoBonsaiService_GetPumps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gobonsai.proto",
}
