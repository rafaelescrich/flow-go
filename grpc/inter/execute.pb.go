// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inter/execute.proto

package bamboo_proto

import (
	fmt "fmt"
	shared "github.com/dapperlabs/bamboo-node/grpc/shared"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("inter/execute.proto", fileDescriptor_0d73587ec8f364c7) }

var fileDescriptor_0d73587ec8f364c7 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8d, 0xbd, 0x4a, 0xc5, 0x40,
	0x10, 0x85, 0x0b, 0xc5, 0x62, 0x4d, 0x35, 0x17, 0x8b, 0x9b, 0x42, 0xd1, 0x4a, 0x11, 0x12, 0xd0,
	0xda, 0x42, 0x41, 0xb4, 0x8a, 0x12, 0xf0, 0x01, 0xf2, 0x33, 0x6e, 0x56, 0x4d, 0x26, 0xee, 0x4c,
	0x40, 0x9f, 0xd1, 0x97, 0x12, 0x93, 0x8d, 0x64, 0x21, 0xb9, 0xa4, 0x9c, 0x39, 0xdf, 0xf9, 0x8e,
	0xda, 0x98, 0x46, 0xd0, 0xc6, 0xf8, 0x85, 0x45, 0x27, 0x18, 0xb5, 0x96, 0x84, 0x20, 0xc8, 0xb3,
	0x3a, 0x27, 0x1a, 0xae, 0xf0, 0x88, 0xab, 0xcc, 0x62, 0x19, 0xd7, 0xc8, 0x9c, 0x69, 0xe4, 0xe1,
	0x7d, 0xf5, 0xb3, 0xa7, 0x0e, 0xef, 0x87, 0x5a, 0x42, 0x25, 0xc2, 0x8d, 0xda, 0x7f, 0x36, 0x8d,
	0x86, 0x6d, 0x34, 0x6d, 0x47, 0x7f, 0xbf, 0x14, 0x3f, 0x3b, 0x64, 0x09, 0xc3, 0xb9, 0x88, 0x5b,
	0x6a, 0x18, 0xe1, 0x45, 0x05, 0xce, 0x76, 0xf7, 0x41, 0xc5, 0x3b, 0x9c, 0xfa, 0xec, 0x34, 0x1b,
	0x75, 0x67, 0xbb, 0x10, 0xa7, 0x7d, 0x53, 0x9b, 0x84, 0xc4, 0xbc, 0x7e, 0xf7, 0x6f, 0x87, 0x94,
	0x70, 0xee, 0x57, 0x67, 0x90, 0x71, 0xe4, 0x62, 0x05, 0xe9, 0xb6, 0x9e, 0x54, 0xf0, 0x80, 0x92,
	0xa2, 0x36, 0x2c, 0x68, 0x19, 0x8e, 0xfd, 0xea, 0x7f, 0x30, 0xaa, 0x4f, 0x16, 0x73, 0x27, 0x34,
	0x6a, 0x3b, 0x15, 0xde, 0x4a, 0xbf, 0xfb, 0x88, 0x46, 0x57, 0x02, 0x97, 0x0b, 0x6d, 0x8f, 0x5a,
	0x3b, 0x95, 0x1f, 0xf4, 0xc1, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x3b, 0x76, 0x35,
	0x10, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecuteNodeClient is the client API for ExecuteNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecuteNodeClient interface {
	Ping(ctx context.Context, in *shared.PingRequest, opts ...grpc.CallOption) (*shared.PingResponse, error)
	ExecuteBlock(ctx context.Context, in *shared.ExecuteBlockRequest, opts ...grpc.CallOption) (*shared.ExecuteBlockResponse, error)
	NotifyBlockExecuted(ctx context.Context, in *shared.NotifyBlockExecutedRequest, opts ...grpc.CallOption) (*shared.NotifyBlockExecutedResponse, error)
	// Providing register values and metadata.
	GetRegisters(ctx context.Context, in *shared.RegistersRequest, opts ...grpc.CallOption) (*shared.RegistersResponse, error)
	GetRegistersAtBlockHeight(ctx context.Context, in *shared.RegistersAtBlockHeightRequest, opts ...grpc.CallOption) (*shared.RegistersResponse, error)
}

type executeNodeClient struct {
	cc *grpc.ClientConn
}

func NewExecuteNodeClient(cc *grpc.ClientConn) ExecuteNodeClient {
	return &executeNodeClient{cc}
}

func (c *executeNodeClient) Ping(ctx context.Context, in *shared.PingRequest, opts ...grpc.CallOption) (*shared.PingResponse, error) {
	out := new(shared.PingResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.ExecuteNode/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeNodeClient) ExecuteBlock(ctx context.Context, in *shared.ExecuteBlockRequest, opts ...grpc.CallOption) (*shared.ExecuteBlockResponse, error) {
	out := new(shared.ExecuteBlockResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.ExecuteNode/ExecuteBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeNodeClient) NotifyBlockExecuted(ctx context.Context, in *shared.NotifyBlockExecutedRequest, opts ...grpc.CallOption) (*shared.NotifyBlockExecutedResponse, error) {
	out := new(shared.NotifyBlockExecutedResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.ExecuteNode/NotifyBlockExecuted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeNodeClient) GetRegisters(ctx context.Context, in *shared.RegistersRequest, opts ...grpc.CallOption) (*shared.RegistersResponse, error) {
	out := new(shared.RegistersResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.ExecuteNode/GetRegisters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeNodeClient) GetRegistersAtBlockHeight(ctx context.Context, in *shared.RegistersAtBlockHeightRequest, opts ...grpc.CallOption) (*shared.RegistersResponse, error) {
	out := new(shared.RegistersResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.ExecuteNode/GetRegistersAtBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuteNodeServer is the server API for ExecuteNode service.
type ExecuteNodeServer interface {
	Ping(context.Context, *shared.PingRequest) (*shared.PingResponse, error)
	ExecuteBlock(context.Context, *shared.ExecuteBlockRequest) (*shared.ExecuteBlockResponse, error)
	NotifyBlockExecuted(context.Context, *shared.NotifyBlockExecutedRequest) (*shared.NotifyBlockExecutedResponse, error)
	// Providing register values and metadata.
	GetRegisters(context.Context, *shared.RegistersRequest) (*shared.RegistersResponse, error)
	GetRegistersAtBlockHeight(context.Context, *shared.RegistersAtBlockHeightRequest) (*shared.RegistersResponse, error)
}

// UnimplementedExecuteNodeServer can be embedded to have forward compatible implementations.
type UnimplementedExecuteNodeServer struct {
}

func (*UnimplementedExecuteNodeServer) Ping(ctx context.Context, req *shared.PingRequest) (*shared.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedExecuteNodeServer) ExecuteBlock(ctx context.Context, req *shared.ExecuteBlockRequest) (*shared.ExecuteBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteBlock not implemented")
}
func (*UnimplementedExecuteNodeServer) NotifyBlockExecuted(ctx context.Context, req *shared.NotifyBlockExecutedRequest) (*shared.NotifyBlockExecutedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyBlockExecuted not implemented")
}
func (*UnimplementedExecuteNodeServer) GetRegisters(ctx context.Context, req *shared.RegistersRequest) (*shared.RegistersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisters not implemented")
}
func (*UnimplementedExecuteNodeServer) GetRegistersAtBlockHeight(ctx context.Context, req *shared.RegistersAtBlockHeightRequest) (*shared.RegistersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistersAtBlockHeight not implemented")
}

func RegisterExecuteNodeServer(s *grpc.Server, srv ExecuteNodeServer) {
	s.RegisterService(&_ExecuteNode_serviceDesc, srv)
}

func _ExecuteNode_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteNodeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.ExecuteNode/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteNodeServer).Ping(ctx, req.(*shared.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteNode_ExecuteBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.ExecuteBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteNodeServer).ExecuteBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.ExecuteNode/ExecuteBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteNodeServer).ExecuteBlock(ctx, req.(*shared.ExecuteBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteNode_NotifyBlockExecuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.NotifyBlockExecutedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteNodeServer).NotifyBlockExecuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.ExecuteNode/NotifyBlockExecuted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteNodeServer).NotifyBlockExecuted(ctx, req.(*shared.NotifyBlockExecutedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteNode_GetRegisters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.RegistersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteNodeServer).GetRegisters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.ExecuteNode/GetRegisters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteNodeServer).GetRegisters(ctx, req.(*shared.RegistersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteNode_GetRegistersAtBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.RegistersAtBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteNodeServer).GetRegistersAtBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.ExecuteNode/GetRegistersAtBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteNodeServer).GetRegistersAtBlockHeight(ctx, req.(*shared.RegistersAtBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecuteNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bamboo.proto.ExecuteNode",
	HandlerType: (*ExecuteNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ExecuteNode_Ping_Handler,
		},
		{
			MethodName: "ExecuteBlock",
			Handler:    _ExecuteNode_ExecuteBlock_Handler,
		},
		{
			MethodName: "NotifyBlockExecuted",
			Handler:    _ExecuteNode_NotifyBlockExecuted_Handler,
		},
		{
			MethodName: "GetRegisters",
			Handler:    _ExecuteNode_GetRegisters_Handler,
		},
		{
			MethodName: "GetRegistersAtBlockHeight",
			Handler:    _ExecuteNode_GetRegistersAtBlockHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inter/execute.proto",
}
