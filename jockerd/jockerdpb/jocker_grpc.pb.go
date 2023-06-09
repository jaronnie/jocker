// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.16.0
// source: jocker.proto

package jockerdpb

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
	Jockerd_ServerVersion_FullMethodName = "/jockerdpb.Jockerd/ServerVersion"
	Jockerd_Container_FullMethodName     = "/jockerdpb.Jockerd/container"
)

// JockerdClient is the client API for Jockerd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JockerdClient interface {
	ServerVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerVersionResponse, error)
	Container(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Container, error)
}

type jockerdClient struct {
	cc grpc.ClientConnInterface
}

func NewJockerdClient(cc grpc.ClientConnInterface) JockerdClient {
	return &jockerdClient{cc}
}

func (c *jockerdClient) ServerVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerVersionResponse, error) {
	out := new(ServerVersionResponse)
	err := c.cc.Invoke(ctx, Jockerd_ServerVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jockerdClient) Container(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Container, error) {
	out := new(Container)
	err := c.cc.Invoke(ctx, Jockerd_Container_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JockerdServer is the server API for Jockerd service.
// All implementations must embed UnimplementedJockerdServer
// for forward compatibility
type JockerdServer interface {
	ServerVersion(context.Context, *Empty) (*ServerVersionResponse, error)
	Container(context.Context, *Empty) (*Container, error)
	mustEmbedUnimplementedJockerdServer()
}

// UnimplementedJockerdServer must be embedded to have forward compatible implementations.
type UnimplementedJockerdServer struct {
}

func (UnimplementedJockerdServer) ServerVersion(context.Context, *Empty) (*ServerVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerVersion not implemented")
}
func (UnimplementedJockerdServer) Container(context.Context, *Empty) (*Container, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Container not implemented")
}
func (UnimplementedJockerdServer) mustEmbedUnimplementedJockerdServer() {}

// UnsafeJockerdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JockerdServer will
// result in compilation errors.
type UnsafeJockerdServer interface {
	mustEmbedUnimplementedJockerdServer()
}

func RegisterJockerdServer(s grpc.ServiceRegistrar, srv JockerdServer) {
	s.RegisterService(&Jockerd_ServiceDesc, srv)
}

func _Jockerd_ServerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JockerdServer).ServerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jockerd_ServerVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JockerdServer).ServerVersion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jockerd_Container_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JockerdServer).Container(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jockerd_Container_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JockerdServer).Container(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Jockerd_ServiceDesc is the grpc.ServiceDesc for Jockerd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jockerd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jockerdpb.Jockerd",
	HandlerType: (*JockerdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerVersion",
			Handler:    _Jockerd_ServerVersion_Handler,
		},
		{
			MethodName: "container",
			Handler:    _Jockerd_Container_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jocker.proto",
}
