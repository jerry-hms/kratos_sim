// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: v1/ws.proto

package v1

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

const (
	Ws_Send_FullMethodName          = "/api.ws.service.v1.Ws/Send"
	Ws_Bind_FullMethodName          = "/api.ws.service.v1.Ws/Bind"
	Ws_CreateMessage_FullMethodName = "/api.ws.service.v1.Ws/CreateMessage"
)

// WsClient is the client API for Ws service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WsClient interface {
	Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendReply, error)
	Bind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*BindReply, error)
	CreateMessage(ctx context.Context, in *CreateMessageReq, opts ...grpc.CallOption) (*CreateMessageReply, error)
}

type wsClient struct {
	cc grpc.ClientConnInterface
}

func NewWsClient(cc grpc.ClientConnInterface) WsClient {
	return &wsClient{cc}
}

func (c *wsClient) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := c.cc.Invoke(ctx, Ws_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wsClient) Bind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*BindReply, error) {
	out := new(BindReply)
	err := c.cc.Invoke(ctx, Ws_Bind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wsClient) CreateMessage(ctx context.Context, in *CreateMessageReq, opts ...grpc.CallOption) (*CreateMessageReply, error) {
	out := new(CreateMessageReply)
	err := c.cc.Invoke(ctx, Ws_CreateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WsServer is the server API for Ws service.
// All implementations must embed UnimplementedWsServer
// for forward compatibility
type WsServer interface {
	Send(context.Context, *SendReq) (*SendReply, error)
	Bind(context.Context, *BindReq) (*BindReply, error)
	CreateMessage(context.Context, *CreateMessageReq) (*CreateMessageReply, error)
	mustEmbedUnimplementedWsServer()
}

// UnimplementedWsServer must be embedded to have forward compatible implementations.
type UnimplementedWsServer struct {
}

func (UnimplementedWsServer) Send(context.Context, *SendReq) (*SendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedWsServer) Bind(context.Context, *BindReq) (*BindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bind not implemented")
}
func (UnimplementedWsServer) CreateMessage(context.Context, *CreateMessageReq) (*CreateMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedWsServer) mustEmbedUnimplementedWsServer() {}

// UnsafeWsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WsServer will
// result in compilation errors.
type UnsafeWsServer interface {
	mustEmbedUnimplementedWsServer()
}

func RegisterWsServer(s grpc.ServiceRegistrar, srv WsServer) {
	s.RegisterService(&Ws_ServiceDesc, srv)
}

func _Ws_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ws_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsServer).Send(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ws_Bind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsServer).Bind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ws_Bind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsServer).Bind(ctx, req.(*BindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ws_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ws_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsServer).CreateMessage(ctx, req.(*CreateMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ws_ServiceDesc is the grpc.ServiceDesc for Ws service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ws_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ws.service.v1.Ws",
	HandlerType: (*WsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Ws_Send_Handler,
		},
		{
			MethodName: "Bind",
			Handler:    _Ws_Bind_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _Ws_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/ws.proto",
}
