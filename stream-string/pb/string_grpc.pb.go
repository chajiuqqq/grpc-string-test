// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: string.proto

package pb

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

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StringServiceClient interface {
	Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
	Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
	ConcatServerStream(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (StringService_ConcatServerStreamClient, error)
	ConcatClientStream(ctx context.Context, opts ...grpc.CallOption) (StringService_ConcatClientStreamClient, error)
	ConcatDoubleStream(ctx context.Context, opts ...grpc.CallOption) (StringService_ConcatDoubleStreamClient, error)
}

type stringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStringServiceClient(cc grpc.ClientConnInterface) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) ConcatServerStream(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (StringService_ConcatServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StringService_ServiceDesc.Streams[0], "/pb.StringService/ConcatServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stringServiceConcatServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StringService_ConcatServerStreamClient interface {
	Recv() (*StringResponse, error)
	grpc.ClientStream
}

type stringServiceConcatServerStreamClient struct {
	grpc.ClientStream
}

func (x *stringServiceConcatServerStreamClient) Recv() (*StringResponse, error) {
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stringServiceClient) ConcatClientStream(ctx context.Context, opts ...grpc.CallOption) (StringService_ConcatClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StringService_ServiceDesc.Streams[1], "/pb.StringService/ConcatClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stringServiceConcatClientStreamClient{stream}
	return x, nil
}

type StringService_ConcatClientStreamClient interface {
	Send(*StringRequest) error
	CloseAndRecv() (*StringResponse, error)
	grpc.ClientStream
}

type stringServiceConcatClientStreamClient struct {
	grpc.ClientStream
}

func (x *stringServiceConcatClientStreamClient) Send(m *StringRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stringServiceConcatClientStreamClient) CloseAndRecv() (*StringResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stringServiceClient) ConcatDoubleStream(ctx context.Context, opts ...grpc.CallOption) (StringService_ConcatDoubleStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StringService_ServiceDesc.Streams[2], "/pb.StringService/ConcatDoubleStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stringServiceConcatDoubleStreamClient{stream}
	return x, nil
}

type StringService_ConcatDoubleStreamClient interface {
	Send(*StringRequest) error
	Recv() (*StringResponse, error)
	grpc.ClientStream
}

type stringServiceConcatDoubleStreamClient struct {
	grpc.ClientStream
}

func (x *stringServiceConcatDoubleStreamClient) Send(m *StringRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stringServiceConcatDoubleStreamClient) Recv() (*StringResponse, error) {
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StringServiceServer is the server API for StringService service.
// All implementations must embed UnimplementedStringServiceServer
// for forward compatibility
type StringServiceServer interface {
	Concat(context.Context, *StringRequest) (*StringResponse, error)
	Diff(context.Context, *StringRequest) (*StringResponse, error)
	ConcatServerStream(*StringRequest, StringService_ConcatServerStreamServer) error
	ConcatClientStream(StringService_ConcatClientStreamServer) error
	ConcatDoubleStream(StringService_ConcatDoubleStreamServer) error
	mustEmbedUnimplementedStringServiceServer()
}

// UnimplementedStringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (UnimplementedStringServiceServer) Concat(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}
func (UnimplementedStringServiceServer) Diff(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedStringServiceServer) ConcatServerStream(*StringRequest, StringService_ConcatServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConcatServerStream not implemented")
}
func (UnimplementedStringServiceServer) ConcatClientStream(StringService_ConcatClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConcatClientStream not implemented")
}
func (UnimplementedStringServiceServer) ConcatDoubleStream(StringService_ConcatDoubleStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConcatDoubleStream not implemented")
}
func (UnimplementedStringServiceServer) mustEmbedUnimplementedStringServiceServer() {}

// UnsafeStringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StringServiceServer will
// result in compilation errors.
type UnsafeStringServiceServer interface {
	mustEmbedUnimplementedStringServiceServer()
}

func RegisterStringServiceServer(s grpc.ServiceRegistrar, srv StringServiceServer) {
	s.RegisterService(&StringService_ServiceDesc, srv)
}

func _StringService_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Concat(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Diff(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_ConcatServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StringServiceServer).ConcatServerStream(m, &stringServiceConcatServerStreamServer{stream})
}

type StringService_ConcatServerStreamServer interface {
	Send(*StringResponse) error
	grpc.ServerStream
}

type stringServiceConcatServerStreamServer struct {
	grpc.ServerStream
}

func (x *stringServiceConcatServerStreamServer) Send(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StringService_ConcatClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StringServiceServer).ConcatClientStream(&stringServiceConcatClientStreamServer{stream})
}

type StringService_ConcatClientStreamServer interface {
	SendAndClose(*StringResponse) error
	Recv() (*StringRequest, error)
	grpc.ServerStream
}

type stringServiceConcatClientStreamServer struct {
	grpc.ServerStream
}

func (x *stringServiceConcatClientStreamServer) SendAndClose(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stringServiceConcatClientStreamServer) Recv() (*StringRequest, error) {
	m := new(StringRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StringService_ConcatDoubleStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StringServiceServer).ConcatDoubleStream(&stringServiceConcatDoubleStreamServer{stream})
}

type StringService_ConcatDoubleStreamServer interface {
	Send(*StringResponse) error
	Recv() (*StringRequest, error)
	grpc.ServerStream
}

type stringServiceConcatDoubleStreamServer struct {
	grpc.ServerStream
}

func (x *stringServiceConcatDoubleStreamServer) Send(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stringServiceConcatDoubleStreamServer) Recv() (*StringRequest, error) {
	m := new(StringRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StringService_ServiceDesc is the grpc.ServiceDesc for StringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Concat",
			Handler:    _StringService_Concat_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _StringService_Diff_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConcatServerStream",
			Handler:       _StringService_ConcatServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConcatClientStream",
			Handler:       _StringService_ConcatClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ConcatDoubleStream",
			Handler:       _StringService_ConcatDoubleStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "string.proto",
}
