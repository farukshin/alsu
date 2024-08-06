// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alsu.proto

package alsu

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("alsu.proto", fileDescriptor_95c8adaf113370ad)
}

var fileDescriptor_95c8adaf113370ad = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcc, 0x29, 0x2e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0xe4, 0xd2, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0xc1, 0x62, 0x49, 0xa5, 0x69, 0xfa, 0xe5, 0x45, 0x89, 0x05, 0x05, 0xa9, 0x45,
	0xc5, 0x10, 0x55, 0x46, 0xe1, 0x5c, 0xdc, 0x8e, 0x39, 0xc5, 0xa5, 0xc1, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x42, 0x1e, 0x5c, 0x5c, 0x05, 0x99, 0x79, 0xe9, 0x20, 0x6e, 0x6a, 0x91, 0x90, 0x8c,
	0x1e, 0x44, 0xb7, 0x1e, 0x4c, 0xb7, 0x5e, 0x70, 0x49, 0x51, 0x66, 0x5e, 0x7a, 0x58, 0x62, 0x4e,
	0x69, 0xaa, 0x14, 0x5e, 0xd9, 0x24, 0x36, 0xb0, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdb,
	0x6a, 0x5a, 0xce, 0x93, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlsuServiceClient is the client API for AlsuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlsuServiceClient interface {
	PingServer(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
}

type alsuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlsuServiceClient(cc grpc.ClientConnInterface) AlsuServiceClient {
	return &alsuServiceClient{cc}
}

func (c *alsuServiceClient) PingServer(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/alsu.AlsuService/pingServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlsuServiceServer is the server API for AlsuService service.
type AlsuServiceServer interface {
	PingServer(context.Context, *wrapperspb.StringValue) (*wrapperspb.StringValue, error)
}

// UnimplementedAlsuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlsuServiceServer struct {
}

func (*UnimplementedAlsuServiceServer) PingServer(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingServer not implemented")
}

func RegisterAlsuServiceServer(s *grpc.Server, srv AlsuServiceServer) {
	s.RegisterService(&_AlsuService_serviceDesc, srv)
}

func _AlsuService_PingServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlsuServiceServer).PingServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alsu.AlsuService/PingServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlsuServiceServer).PingServer(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlsuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alsu.AlsuService",
	HandlerType: (*AlsuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pingServer",
			Handler:    _AlsuService_PingServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alsu.proto",
}