// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: notificationManager.proto

package proto

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

// NotificationManagerClient is the client API for NotificationManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationManagerClient interface {
	GetInvites(ctx context.Context, in *InviteInput, opts ...grpc.CallOption) (*InviteOutput, error)
}

type notificationManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationManagerClient(cc grpc.ClientConnInterface) NotificationManagerClient {
	return &notificationManagerClient{cc}
}

func (c *notificationManagerClient) GetInvites(ctx context.Context, in *InviteInput, opts ...grpc.CallOption) (*InviteOutput, error) {
	out := new(InviteOutput)
	err := c.cc.Invoke(ctx, "/NotificationManager/GetInvites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationManagerServer is the server API for NotificationManager service.
// All implementations must embed UnimplementedNotificationManagerServer
// for forward compatibility
type NotificationManagerServer interface {
	GetInvites(context.Context, *InviteInput) (*InviteOutput, error)
	mustEmbedUnimplementedNotificationManagerServer()
}

// UnimplementedNotificationManagerServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationManagerServer struct {
}

func (UnimplementedNotificationManagerServer) GetInvites(context.Context, *InviteInput) (*InviteOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvites not implemented")
}
func (UnimplementedNotificationManagerServer) mustEmbedUnimplementedNotificationManagerServer() {}

// UnsafeNotificationManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationManagerServer will
// result in compilation errors.
type UnsafeNotificationManagerServer interface {
	mustEmbedUnimplementedNotificationManagerServer()
}

func RegisterNotificationManagerServer(s grpc.ServiceRegistrar, srv NotificationManagerServer) {
	s.RegisterService(&NotificationManager_ServiceDesc, srv)
}

func _NotificationManager_GetInvites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).GetInvites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationManager/GetInvites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).GetInvites(ctx, req.(*InviteInput))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationManager_ServiceDesc is the grpc.ServiceDesc for NotificationManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotificationManager",
	HandlerType: (*NotificationManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInvites",
			Handler:    _NotificationManager_GetInvites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notificationManager.proto",
}
