// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: fiamma/bitvmstaker/tx.proto

package bitvmstaker

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
	Msg_UpdateParams_FullMethodName           = "/fiamma.bitvmstaker.Msg/UpdateParams"
	Msg_CreateStaker_FullMethodName           = "/fiamma.bitvmstaker.Msg/CreateStaker"
	Msg_RemoveStaker_FullMethodName           = "/fiamma.bitvmstaker.Msg/RemoveStaker"
	Msg_UpdateCommitteeAddress_FullMethodName = "/fiamma.bitvmstaker.Msg/UpdateCommitteeAddress"
	Msg_RegisterVK_FullMethodName             = "/fiamma.bitvmstaker.Msg/RegisterVK"
	Msg_RemoveVK_FullMethodName               = "/fiamma.bitvmstaker.Msg/RemoveVK"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateStaker(ctx context.Context, in *MsgCreateStaker, opts ...grpc.CallOption) (*MsgCreateStakerResponse, error)
	RemoveStaker(ctx context.Context, in *MsgRemoveStaker, opts ...grpc.CallOption) (*MsgRemoveStakerResponse, error)
	UpdateCommitteeAddress(ctx context.Context, in *MsgUpdateCommitteeAddress, opts ...grpc.CallOption) (*MsgUpdateCommitteeAddressResponse, error)
	RegisterVK(ctx context.Context, in *MsgRegisterVK, opts ...grpc.CallOption) (*MsgRegisterVKResponse, error)
	RemoveVK(ctx context.Context, in *MsgRemoveVK, opts ...grpc.CallOption) (*MsgRemoveVKResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateStaker(ctx context.Context, in *MsgCreateStaker, opts ...grpc.CallOption) (*MsgCreateStakerResponse, error) {
	out := new(MsgCreateStakerResponse)
	err := c.cc.Invoke(ctx, Msg_CreateStaker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveStaker(ctx context.Context, in *MsgRemoveStaker, opts ...grpc.CallOption) (*MsgRemoveStakerResponse, error) {
	out := new(MsgRemoveStakerResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveStaker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateCommitteeAddress(ctx context.Context, in *MsgUpdateCommitteeAddress, opts ...grpc.CallOption) (*MsgUpdateCommitteeAddressResponse, error) {
	out := new(MsgUpdateCommitteeAddressResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateCommitteeAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterVK(ctx context.Context, in *MsgRegisterVK, opts ...grpc.CallOption) (*MsgRegisterVKResponse, error) {
	out := new(MsgRegisterVKResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterVK_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveVK(ctx context.Context, in *MsgRemoveVK, opts ...grpc.CallOption) (*MsgRemoveVKResponse, error) {
	out := new(MsgRemoveVKResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveVK_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateStaker(context.Context, *MsgCreateStaker) (*MsgCreateStakerResponse, error)
	RemoveStaker(context.Context, *MsgRemoveStaker) (*MsgRemoveStakerResponse, error)
	UpdateCommitteeAddress(context.Context, *MsgUpdateCommitteeAddress) (*MsgUpdateCommitteeAddressResponse, error)
	RegisterVK(context.Context, *MsgRegisterVK) (*MsgRegisterVKResponse, error)
	RemoveVK(context.Context, *MsgRemoveVK) (*MsgRemoveVKResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateStaker(context.Context, *MsgCreateStaker) (*MsgCreateStakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaker not implemented")
}
func (UnimplementedMsgServer) RemoveStaker(context.Context, *MsgRemoveStaker) (*MsgRemoveStakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStaker not implemented")
}
func (UnimplementedMsgServer) UpdateCommitteeAddress(context.Context, *MsgUpdateCommitteeAddress) (*MsgUpdateCommitteeAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommitteeAddress not implemented")
}
func (UnimplementedMsgServer) RegisterVK(context.Context, *MsgRegisterVK) (*MsgRegisterVKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterVK not implemented")
}
func (UnimplementedMsgServer) RemoveVK(context.Context, *MsgRemoveVK) (*MsgRemoveVKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVK not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateStaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateStaker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStaker(ctx, req.(*MsgCreateStaker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveStaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveStaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveStaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveStaker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveStaker(ctx, req.(*MsgRemoveStaker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateCommitteeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateCommitteeAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateCommitteeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateCommitteeAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateCommitteeAddress(ctx, req.(*MsgUpdateCommitteeAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterVK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterVK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterVK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterVK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterVK(ctx, req.(*MsgRegisterVK))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveVK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveVK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveVK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveVK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveVK(ctx, req.(*MsgRemoveVK))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fiamma.bitvmstaker.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateStaker",
			Handler:    _Msg_CreateStaker_Handler,
		},
		{
			MethodName: "RemoveStaker",
			Handler:    _Msg_RemoveStaker_Handler,
		},
		{
			MethodName: "UpdateCommitteeAddress",
			Handler:    _Msg_UpdateCommitteeAddress_Handler,
		},
		{
			MethodName: "RegisterVK",
			Handler:    _Msg_RegisterVK_Handler,
		},
		{
			MethodName: "RemoveVK",
			Handler:    _Msg_RemoveVK_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fiamma/bitvmstaker/tx.proto",
}
