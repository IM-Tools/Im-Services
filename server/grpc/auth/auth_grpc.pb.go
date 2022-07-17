// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcAuth

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

// ImAuthClient is the client API for ImAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImAuthClient interface {
	CheckAuth(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error)
}

type imAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewImAuthClient(cc grpc.ClientConnInterface) ImAuthClient {
	return &imAuthClient{cc}
}

func (c *imAuthClient) CheckAuth(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error) {
	out := new(CheckAuthResponse)
	err := c.cc.Invoke(ctx, "/ImAuth/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImAuthServer is the server API for ImAuth service.
// All implementations must embed UnimplementedImAuthServer
// for forward compatibility
type ImAuthServer interface {
	CheckAuth(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error)
	mustEmbedUnimplementedImAuthServer()
}

// UnimplementedImAuthServer must be embedded to have forward compatible implementations.
type UnimplementedImAuthServer struct {
}

func (UnimplementedImAuthServer) CheckAuth(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedImAuthServer) mustEmbedUnimplementedImAuthServer() {}

// UnsafeImAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImAuthServer will
// result in compilation errors.
type UnsafeImAuthServer interface {
	mustEmbedUnimplementedImAuthServer()
}

func RegisterImAuthServer(s grpc.ServiceRegistrar, srv ImAuthServer) {
	s.RegisterService(&ImAuth_ServiceDesc, srv)
}

func _ImAuth_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImAuthServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImAuth/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImAuthServer).CheckAuth(ctx, req.(*CheckAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImAuth_ServiceDesc is the grpc.ServiceDesc for ImAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImAuth",
	HandlerType: (*ImAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _ImAuth_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/protos/auth.proto",
}
