// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: v1/session.proto

package session

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
	Session_CreateSession_FullMethodName  = "/api.session.v1.Session/CreateSession"
	Session_DeleteSession_FullMethodName  = "/api.session.v1.Session/DeleteSession"
	Session_GetSession_FullMethodName     = "/api.session.v1.Session/GetSession"
	Session_ListSession_FullMethodName    = "/api.session.v1.Session/ListSession"
	Session_CreateRelation_FullMethodName = "/api.session.v1.Session/CreateRelation"
	Session_GetRelation_FullMethodName    = "/api.session.v1.Session/GetRelation"
)

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionReply, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionReply, error)
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionReply, error)
	ListSession(ctx context.Context, in *ListSessionRequest, opts ...grpc.CallOption) (*ListSessionReply, error)
	CreateRelation(ctx context.Context, in *CreateRelationReq, opts ...grpc.CallOption) (*CreateRelationReply, error)
	GetRelation(ctx context.Context, in *GetRelationReq, opts ...grpc.CallOption) (*GetRelationReply, error)
}

type sessionClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionClient(cc grpc.ClientConnInterface) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionReply, error) {
	out := new(CreateSessionReply)
	err := c.cc.Invoke(ctx, Session_CreateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionReply, error) {
	out := new(DeleteSessionReply)
	err := c.cc.Invoke(ctx, Session_DeleteSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionReply, error) {
	out := new(GetSessionReply)
	err := c.cc.Invoke(ctx, Session_GetSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) ListSession(ctx context.Context, in *ListSessionRequest, opts ...grpc.CallOption) (*ListSessionReply, error) {
	out := new(ListSessionReply)
	err := c.cc.Invoke(ctx, Session_ListSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) CreateRelation(ctx context.Context, in *CreateRelationReq, opts ...grpc.CallOption) (*CreateRelationReply, error) {
	out := new(CreateRelationReply)
	err := c.cc.Invoke(ctx, Session_CreateRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) GetRelation(ctx context.Context, in *GetRelationReq, opts ...grpc.CallOption) (*GetRelationReply, error) {
	out := new(GetRelationReply)
	err := c.cc.Invoke(ctx, Session_GetRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
// All implementations must embed UnimplementedSessionServer
// for forward compatibility
type SessionServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionReply, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionReply, error)
	GetSession(context.Context, *GetSessionRequest) (*GetSessionReply, error)
	ListSession(context.Context, *ListSessionRequest) (*ListSessionReply, error)
	CreateRelation(context.Context, *CreateRelationReq) (*CreateRelationReply, error)
	GetRelation(context.Context, *GetRelationReq) (*GetRelationReply, error)
	mustEmbedUnimplementedSessionServer()
}

// UnimplementedSessionServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (UnimplementedSessionServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedSessionServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedSessionServer) GetSession(context.Context, *GetSessionRequest) (*GetSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedSessionServer) ListSession(context.Context, *ListSessionRequest) (*ListSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSession not implemented")
}
func (UnimplementedSessionServer) CreateRelation(context.Context, *CreateRelationReq) (*CreateRelationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelation not implemented")
}
func (UnimplementedSessionServer) GetRelation(context.Context, *GetRelationReq) (*GetRelationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelation not implemented")
}
func (UnimplementedSessionServer) mustEmbedUnimplementedSessionServer() {}

// UnsafeSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServer will
// result in compilation errors.
type UnsafeSessionServer interface {
	mustEmbedUnimplementedSessionServer()
}

func RegisterSessionServer(s grpc.ServiceRegistrar, srv SessionServer) {
	s.RegisterService(&Session_ServiceDesc, srv)
}

func _Session_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_DeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_GetSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_ListSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).ListSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_ListSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).ListSession(ctx, req.(*ListSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_CreateRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).CreateRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_CreateRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).CreateRelation(ctx, req.(*CreateRelationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_GetRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).GetRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_GetRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).GetRelation(ctx, req.(*GetRelationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Session_ServiceDesc is the grpc.ServiceDesc for Session service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Session_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.session.v1.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _Session_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _Session_DeleteSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _Session_GetSession_Handler,
		},
		{
			MethodName: "ListSession",
			Handler:    _Session_ListSession_Handler,
		},
		{
			MethodName: "CreateRelation",
			Handler:    _Session_CreateRelation_Handler,
		},
		{
			MethodName: "GetRelation",
			Handler:    _Session_GetRelation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/session.proto",
}
