// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: randomgenerator.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RandomTextGeneratorService_GenerateTextEN_FullMethodName  = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateTextEN"
	RandomTextGeneratorService_GenerateTextTR_FullMethodName  = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateTextTR"
	RandomTextGeneratorService_GenerateTextsEN_FullMethodName = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateTextsEN"
	RandomTextGeneratorService_GenerateTextsTR_FullMethodName = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateTextsTR"
	RandomTextGeneratorService_GetTextBound_FullMethodName    = "/org.csystem.generator.random.RandomTextGeneratorService/GetTextBound"
	RandomTextGeneratorService_GenerateInt32_FullMethodName   = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateInt32"
	RandomTextGeneratorService_GenerateDouble_FullMethodName  = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateDouble"
	RandomTextGeneratorService_GenerateInt64_FullMethodName   = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateInt64"
	RandomTextGeneratorService_GenerateInt32S_FullMethodName  = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateInt32s"
	RandomTextGeneratorService_GenerateDoubles_FullMethodName = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateDoubles"
	RandomTextGeneratorService_GenerateInt64S_FullMethodName  = "/org.csystem.generator.random.RandomTextGeneratorService/GenerateInt64s"
)

// RandomTextGeneratorServiceClient is the client API for RandomTextGeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomTextGeneratorServiceClient interface {
	GenerateTextEN(ctx context.Context, in *TextGenerateInfo, opts ...grpc.CallOption) (*TextInfo, error)
	GenerateTextTR(ctx context.Context, in *TextGenerateInfo, opts ...grpc.CallOption) (*TextInfo, error)
	GenerateTextsEN(ctx context.Context, in *TextsGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TextInfo], error)
	GenerateTextsTR(ctx context.Context, in *TextsGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TextInfo], error)
	GetTextBound(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*TextBound, error)
	GenerateInt32(ctx context.Context, in *Int32GenerateInfo, opts ...grpc.CallOption) (*Int32Result, error)
	GenerateDouble(ctx context.Context, in *DoubleGenerateInfo, opts ...grpc.CallOption) (*DoubleResult, error)
	GenerateInt64(ctx context.Context, in *Int64GenerateInfo, opts ...grpc.CallOption) (*Int64Result, error)
	GenerateInt32S(ctx context.Context, in *Int32SGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Int32Result], error)
	GenerateDoubles(ctx context.Context, in *DoublesGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DoubleResult], error)
	GenerateInt64S(ctx context.Context, in *Int64SGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Int64Result], error)
}

type randomTextGeneratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomTextGeneratorServiceClient(cc grpc.ClientConnInterface) RandomTextGeneratorServiceClient {
	return &randomTextGeneratorServiceClient{cc}
}

func (c *randomTextGeneratorServiceClient) GenerateTextEN(ctx context.Context, in *TextGenerateInfo, opts ...grpc.CallOption) (*TextInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextInfo)
	err := c.cc.Invoke(ctx, RandomTextGeneratorService_GenerateTextEN_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomTextGeneratorServiceClient) GenerateTextTR(ctx context.Context, in *TextGenerateInfo, opts ...grpc.CallOption) (*TextInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextInfo)
	err := c.cc.Invoke(ctx, RandomTextGeneratorService_GenerateTextTR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomTextGeneratorServiceClient) GenerateTextsEN(ctx context.Context, in *TextsGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TextInfo], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RandomTextGeneratorService_ServiceDesc.Streams[0], RandomTextGeneratorService_GenerateTextsEN_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TextsGenerateInfo, TextInfo]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateTextsENClient = grpc.ServerStreamingClient[TextInfo]

func (c *randomTextGeneratorServiceClient) GenerateTextsTR(ctx context.Context, in *TextsGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TextInfo], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RandomTextGeneratorService_ServiceDesc.Streams[1], RandomTextGeneratorService_GenerateTextsTR_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TextsGenerateInfo, TextInfo]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateTextsTRClient = grpc.ServerStreamingClient[TextInfo]

func (c *randomTextGeneratorServiceClient) GetTextBound(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*TextBound, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextBound)
	err := c.cc.Invoke(ctx, RandomTextGeneratorService_GetTextBound_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomTextGeneratorServiceClient) GenerateInt32(ctx context.Context, in *Int32GenerateInfo, opts ...grpc.CallOption) (*Int32Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Result)
	err := c.cc.Invoke(ctx, RandomTextGeneratorService_GenerateInt32_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomTextGeneratorServiceClient) GenerateDouble(ctx context.Context, in *DoubleGenerateInfo, opts ...grpc.CallOption) (*DoubleResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoubleResult)
	err := c.cc.Invoke(ctx, RandomTextGeneratorService_GenerateDouble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomTextGeneratorServiceClient) GenerateInt64(ctx context.Context, in *Int64GenerateInfo, opts ...grpc.CallOption) (*Int64Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int64Result)
	err := c.cc.Invoke(ctx, RandomTextGeneratorService_GenerateInt64_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomTextGeneratorServiceClient) GenerateInt32S(ctx context.Context, in *Int32SGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Int32Result], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RandomTextGeneratorService_ServiceDesc.Streams[2], RandomTextGeneratorService_GenerateInt32S_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Int32SGenerateInfo, Int32Result]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateInt32SClient = grpc.ServerStreamingClient[Int32Result]

func (c *randomTextGeneratorServiceClient) GenerateDoubles(ctx context.Context, in *DoublesGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DoubleResult], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RandomTextGeneratorService_ServiceDesc.Streams[3], RandomTextGeneratorService_GenerateDoubles_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DoublesGenerateInfo, DoubleResult]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateDoublesClient = grpc.ServerStreamingClient[DoubleResult]

func (c *randomTextGeneratorServiceClient) GenerateInt64S(ctx context.Context, in *Int64SGenerateInfo, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Int64Result], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RandomTextGeneratorService_ServiceDesc.Streams[4], RandomTextGeneratorService_GenerateInt64S_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Int64SGenerateInfo, Int64Result]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateInt64SClient = grpc.ServerStreamingClient[Int64Result]

// RandomTextGeneratorServiceServer is the server API for RandomTextGeneratorService service.
// All implementations must embed UnimplementedRandomTextGeneratorServiceServer
// for forward compatibility.
type RandomTextGeneratorServiceServer interface {
	GenerateTextEN(context.Context, *TextGenerateInfo) (*TextInfo, error)
	GenerateTextTR(context.Context, *TextGenerateInfo) (*TextInfo, error)
	GenerateTextsEN(*TextsGenerateInfo, grpc.ServerStreamingServer[TextInfo]) error
	GenerateTextsTR(*TextsGenerateInfo, grpc.ServerStreamingServer[TextInfo]) error
	GetTextBound(context.Context, *NoParam) (*TextBound, error)
	GenerateInt32(context.Context, *Int32GenerateInfo) (*Int32Result, error)
	GenerateDouble(context.Context, *DoubleGenerateInfo) (*DoubleResult, error)
	GenerateInt64(context.Context, *Int64GenerateInfo) (*Int64Result, error)
	GenerateInt32S(*Int32SGenerateInfo, grpc.ServerStreamingServer[Int32Result]) error
	GenerateDoubles(*DoublesGenerateInfo, grpc.ServerStreamingServer[DoubleResult]) error
	GenerateInt64S(*Int64SGenerateInfo, grpc.ServerStreamingServer[Int64Result]) error
	mustEmbedUnimplementedRandomTextGeneratorServiceServer()
}

// UnimplementedRandomTextGeneratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRandomTextGeneratorServiceServer struct{}

func (UnimplementedRandomTextGeneratorServiceServer) GenerateTextEN(context.Context, *TextGenerateInfo) (*TextInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTextEN not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateTextTR(context.Context, *TextGenerateInfo) (*TextInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTextTR not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateTextsEN(*TextsGenerateInfo, grpc.ServerStreamingServer[TextInfo]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateTextsEN not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateTextsTR(*TextsGenerateInfo, grpc.ServerStreamingServer[TextInfo]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateTextsTR not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GetTextBound(context.Context, *NoParam) (*TextBound, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTextBound not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateInt32(context.Context, *Int32GenerateInfo) (*Int32Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInt32 not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateDouble(context.Context, *DoubleGenerateInfo) (*DoubleResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDouble not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateInt64(context.Context, *Int64GenerateInfo) (*Int64Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInt64 not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateInt32S(*Int32SGenerateInfo, grpc.ServerStreamingServer[Int32Result]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateInt32S not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateDoubles(*DoublesGenerateInfo, grpc.ServerStreamingServer[DoubleResult]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateDoubles not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) GenerateInt64S(*Int64SGenerateInfo, grpc.ServerStreamingServer[Int64Result]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateInt64S not implemented")
}
func (UnimplementedRandomTextGeneratorServiceServer) mustEmbedUnimplementedRandomTextGeneratorServiceServer() {
}
func (UnimplementedRandomTextGeneratorServiceServer) testEmbeddedByValue() {}

// UnsafeRandomTextGeneratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomTextGeneratorServiceServer will
// result in compilation errors.
type UnsafeRandomTextGeneratorServiceServer interface {
	mustEmbedUnimplementedRandomTextGeneratorServiceServer()
}

func RegisterRandomTextGeneratorServiceServer(s grpc.ServiceRegistrar, srv RandomTextGeneratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedRandomTextGeneratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RandomTextGeneratorService_ServiceDesc, srv)
}

func _RandomTextGeneratorService_GenerateTextEN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextGenerateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomTextGeneratorServiceServer).GenerateTextEN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomTextGeneratorService_GenerateTextEN_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomTextGeneratorServiceServer).GenerateTextEN(ctx, req.(*TextGenerateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomTextGeneratorService_GenerateTextTR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextGenerateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomTextGeneratorServiceServer).GenerateTextTR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomTextGeneratorService_GenerateTextTR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomTextGeneratorServiceServer).GenerateTextTR(ctx, req.(*TextGenerateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomTextGeneratorService_GenerateTextsEN_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TextsGenerateInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomTextGeneratorServiceServer).GenerateTextsEN(m, &grpc.GenericServerStream[TextsGenerateInfo, TextInfo]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateTextsENServer = grpc.ServerStreamingServer[TextInfo]

func _RandomTextGeneratorService_GenerateTextsTR_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TextsGenerateInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomTextGeneratorServiceServer).GenerateTextsTR(m, &grpc.GenericServerStream[TextsGenerateInfo, TextInfo]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateTextsTRServer = grpc.ServerStreamingServer[TextInfo]

func _RandomTextGeneratorService_GetTextBound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomTextGeneratorServiceServer).GetTextBound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomTextGeneratorService_GetTextBound_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomTextGeneratorServiceServer).GetTextBound(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomTextGeneratorService_GenerateInt32_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32GenerateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomTextGeneratorServiceServer).GenerateInt32(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomTextGeneratorService_GenerateInt32_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomTextGeneratorServiceServer).GenerateInt32(ctx, req.(*Int32GenerateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomTextGeneratorService_GenerateDouble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoubleGenerateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomTextGeneratorServiceServer).GenerateDouble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomTextGeneratorService_GenerateDouble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomTextGeneratorServiceServer).GenerateDouble(ctx, req.(*DoubleGenerateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomTextGeneratorService_GenerateInt64_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64GenerateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomTextGeneratorServiceServer).GenerateInt64(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomTextGeneratorService_GenerateInt64_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomTextGeneratorServiceServer).GenerateInt64(ctx, req.(*Int64GenerateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomTextGeneratorService_GenerateInt32S_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Int32SGenerateInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomTextGeneratorServiceServer).GenerateInt32S(m, &grpc.GenericServerStream[Int32SGenerateInfo, Int32Result]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateInt32SServer = grpc.ServerStreamingServer[Int32Result]

func _RandomTextGeneratorService_GenerateDoubles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DoublesGenerateInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomTextGeneratorServiceServer).GenerateDoubles(m, &grpc.GenericServerStream[DoublesGenerateInfo, DoubleResult]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateDoublesServer = grpc.ServerStreamingServer[DoubleResult]

func _RandomTextGeneratorService_GenerateInt64S_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Int64SGenerateInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomTextGeneratorServiceServer).GenerateInt64S(m, &grpc.GenericServerStream[Int64SGenerateInfo, Int64Result]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomTextGeneratorService_GenerateInt64SServer = grpc.ServerStreamingServer[Int64Result]

// RandomTextGeneratorService_ServiceDesc is the grpc.ServiceDesc for RandomTextGeneratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandomTextGeneratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.csystem.generator.random.RandomTextGeneratorService",
	HandlerType: (*RandomTextGeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTextEN",
			Handler:    _RandomTextGeneratorService_GenerateTextEN_Handler,
		},
		{
			MethodName: "GenerateTextTR",
			Handler:    _RandomTextGeneratorService_GenerateTextTR_Handler,
		},
		{
			MethodName: "GetTextBound",
			Handler:    _RandomTextGeneratorService_GetTextBound_Handler,
		},
		{
			MethodName: "GenerateInt32",
			Handler:    _RandomTextGeneratorService_GenerateInt32_Handler,
		},
		{
			MethodName: "GenerateDouble",
			Handler:    _RandomTextGeneratorService_GenerateDouble_Handler,
		},
		{
			MethodName: "GenerateInt64",
			Handler:    _RandomTextGeneratorService_GenerateInt64_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateTextsEN",
			Handler:       _RandomTextGeneratorService_GenerateTextsEN_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GenerateTextsTR",
			Handler:       _RandomTextGeneratorService_GenerateTextsTR_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GenerateInt32s",
			Handler:       _RandomTextGeneratorService_GenerateInt32S_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GenerateDoubles",
			Handler:       _RandomTextGeneratorService_GenerateDoubles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GenerateInt64s",
			Handler:       _RandomTextGeneratorService_GenerateInt64S_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "randomgenerator.proto",
}
