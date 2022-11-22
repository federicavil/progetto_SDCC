// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: visitManager.proto

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

// ManageVisitClient is the client API for ManageVisit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageVisitClient interface {
	AddNewVisit(ctx context.Context, in *InputVisit, opts ...grpc.CallOption) (*Return, error)
	GetAllVisits(ctx context.Context, in *User, opts ...grpc.CallOption) (*Visits, error)
	GetVisitByID(ctx context.Context, in *ID_Visit, opts ...grpc.CallOption) (*Visit, error)
	InviteUserToVisit(ctx context.Context, in *Invite, opts ...grpc.CallOption) (*Return, error)
	AcceptOrRefuseInvite(ctx context.Context, in *InviteResponse, opts ...grpc.CallOption) (*Return, error)
}

type manageVisitClient struct {
	cc grpc.ClientConnInterface
}

func NewManageVisitClient(cc grpc.ClientConnInterface) ManageVisitClient {
	return &manageVisitClient{cc}
}

func (c *manageVisitClient) AddNewVisit(ctx context.Context, in *InputVisit, opts ...grpc.CallOption) (*Return, error) {
	out := new(Return)
	err := c.cc.Invoke(ctx, "/ManageVisit/AddNewVisit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVisitClient) GetAllVisits(ctx context.Context, in *User, opts ...grpc.CallOption) (*Visits, error) {
	out := new(Visits)
	err := c.cc.Invoke(ctx, "/ManageVisit/GetAllVisits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVisitClient) GetVisitByID(ctx context.Context, in *ID_Visit, opts ...grpc.CallOption) (*Visit, error) {
	out := new(Visit)
	err := c.cc.Invoke(ctx, "/ManageVisit/GetVisitByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVisitClient) InviteUserToVisit(ctx context.Context, in *Invite, opts ...grpc.CallOption) (*Return, error) {
	out := new(Return)
	err := c.cc.Invoke(ctx, "/ManageVisit/InviteUserToVisit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageVisitClient) AcceptOrRefuseInvite(ctx context.Context, in *InviteResponse, opts ...grpc.CallOption) (*Return, error) {
	out := new(Return)
	err := c.cc.Invoke(ctx, "/ManageVisit/AcceptOrRefuseInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageVisitServer is the server API for ManageVisit service.
// All implementations must embed UnimplementedManageVisitServer
// for forward compatibility
type ManageVisitServer interface {
	AddNewVisit(context.Context, *InputVisit) (*Return, error)
	GetAllVisits(context.Context, *User) (*Visits, error)
	GetVisitByID(context.Context, *ID_Visit) (*Visit, error)
	InviteUserToVisit(context.Context, *Invite) (*Return, error)
	AcceptOrRefuseInvite(context.Context, *InviteResponse) (*Return, error)
	mustEmbedUnimplementedManageVisitServer()
}

// UnimplementedManageVisitServer must be embedded to have forward compatible implementations.
type UnimplementedManageVisitServer struct {
}

func (UnimplementedManageVisitServer) AddNewVisit(context.Context, *InputVisit) (*Return, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewVisit not implemented")
}
func (UnimplementedManageVisitServer) GetAllVisits(context.Context, *User) (*Visits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVisits not implemented")
}
func (UnimplementedManageVisitServer) GetVisitByID(context.Context, *ID_Visit) (*Visit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVisitByID not implemented")
}
func (UnimplementedManageVisitServer) InviteUserToVisit(context.Context, *Invite) (*Return, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUserToVisit not implemented")
}
func (UnimplementedManageVisitServer) AcceptOrRefuseInvite(context.Context, *InviteResponse) (*Return, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOrRefuseInvite not implemented")
}
func (UnimplementedManageVisitServer) mustEmbedUnimplementedManageVisitServer() {}

// UnsafeManageVisitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageVisitServer will
// result in compilation errors.
type UnsafeManageVisitServer interface {
	mustEmbedUnimplementedManageVisitServer()
}

func RegisterManageVisitServer(s grpc.ServiceRegistrar, srv ManageVisitServer) {
	s.RegisterService(&ManageVisit_ServiceDesc, srv)
}

func _ManageVisit_AddNewVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputVisit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVisitServer).AddNewVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManageVisit/AddNewVisit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVisitServer).AddNewVisit(ctx, req.(*InputVisit))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVisit_GetAllVisits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVisitServer).GetAllVisits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManageVisit/GetAllVisits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVisitServer).GetAllVisits(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVisit_GetVisitByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID_Visit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVisitServer).GetVisitByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManageVisit/GetVisitByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVisitServer).GetVisitByID(ctx, req.(*ID_Visit))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVisit_InviteUserToVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Invite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVisitServer).InviteUserToVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManageVisit/InviteUserToVisit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVisitServer).InviteUserToVisit(ctx, req.(*Invite))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageVisit_AcceptOrRefuseInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageVisitServer).AcceptOrRefuseInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManageVisit/AcceptOrRefuseInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageVisitServer).AcceptOrRefuseInvite(ctx, req.(*InviteResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// ManageVisit_ServiceDesc is the grpc.ServiceDesc for ManageVisit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManageVisit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ManageVisit",
	HandlerType: (*ManageVisitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewVisit",
			Handler:    _ManageVisit_AddNewVisit_Handler,
		},
		{
			MethodName: "GetAllVisits",
			Handler:    _ManageVisit_GetAllVisits_Handler,
		},
		{
			MethodName: "GetVisitByID",
			Handler:    _ManageVisit_GetVisitByID_Handler,
		},
		{
			MethodName: "InviteUserToVisit",
			Handler:    _ManageVisit_InviteUserToVisit_Handler,
		},
		{
			MethodName: "AcceptOrRefuseInvite",
			Handler:    _ManageVisit_AcceptOrRefuseInvite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "visitManager.proto",
}
