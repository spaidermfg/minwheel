// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: hello.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	HelloProtobuf(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// stream启用流特性
	Channel(ctx context.Context, opts ...grpc.CallOption) (HelloService_ChannelClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) HelloProtobuf(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/main.HelloService/HelloProtobuf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Channel(ctx context.Context, opts ...grpc.CallOption) (HelloService_ChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/main.HelloService/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceChannelClient{stream}
	return x, nil
}

type HelloService_ChannelClient interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ClientStream
}

type helloServiceChannelClient struct {
	grpc.ClientStream
}

func (x *helloServiceChannelClient) Send(m *String) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceChannelClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	HelloProtobuf(context.Context, *String) (*String, error)
	// stream启用流特性
	Channel(HelloService_ChannelServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) HelloProtobuf(context.Context, *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloProtobuf not implemented")
}
func (UnimplementedHelloServiceServer) Channel(HelloService_ChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method Channel not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_HelloProtobuf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).HelloProtobuf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.HelloService/HelloProtobuf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).HelloProtobuf(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).Channel(&helloServiceChannelServer{stream})
}

type HelloService_ChannelServer interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ServerStream
}

type helloServiceChannelServer struct {
	grpc.ServerStream
}

func (x *helloServiceChannelServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceChannelServer) Recv() (*String, error) {
	m := new(String)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloProtobuf",
			Handler:    _HelloService_HelloProtobuf_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _HelloService_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}

// PubsubServiceClient is the client API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
}

type pubsubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubsubServiceClient(cc grpc.ClientConnInterface) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/main.PubsubService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubsubService_ServiceDesc.Streams[0], "/main.PubsubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubscribeClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubsubServiceServer is the server API for PubsubService service.
// All implementations must embed UnimplementedPubsubServiceServer
// for forward compatibility
type PubsubServiceServer interface {
	Publish(context.Context, *String) (*String, error)
	Subscribe(*String, PubsubService_SubscribeServer) error
	mustEmbedUnimplementedPubsubServiceServer()
}

// UnimplementedPubsubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPubsubServiceServer struct {
}

func (UnimplementedPubsubServiceServer) Publish(context.Context, *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPubsubServiceServer) Subscribe(*String, PubsubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubsubServiceServer) mustEmbedUnimplementedPubsubServiceServer() {}

// UnsafePubsubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubsubServiceServer will
// result in compilation errors.
type UnsafePubsubServiceServer interface {
	mustEmbedUnimplementedPubsubServiceServer()
}

func RegisterPubsubServiceServer(s grpc.ServiceRegistrar, srv PubsubServiceServer) {
	s.RegisterService(&PubsubService_ServiceDesc, srv)
}

func _PubsubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.PubsubService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Publish(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubsubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(String)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subscribe(m, &pubsubServiceSubscribeServer{stream})
}

type PubsubService_SubscribeServer interface {
	Send(*String) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

// PubsubService_ServiceDesc is the grpc.ServiceDesc for PubsubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubsubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubsubService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello.proto",
}
