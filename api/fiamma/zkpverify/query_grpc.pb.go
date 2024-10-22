// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: fiamma/zkpverify/query.proto

package zkpverify

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
	Query_Params_FullMethodName                   = "/fiamma.zkpverify.Query/Params"
	Query_PendingProof_FullMethodName             = "/fiamma.zkpverify.Query/PendingProof"
	Query_ProofData_FullMethodName                = "/fiamma.zkpverify.Query/ProofData"
	Query_VerifyResult_FullMethodName             = "/fiamma.zkpverify.Query/VerifyResult"
	Query_VerifyResultsByNamespace_FullMethodName = "/fiamma.zkpverify.Query/VerifyResultsByNamespace"
	Query_PendingProofByNamespace_FullMethodName  = "/fiamma.zkpverify.Query/PendingProofByNamespace"
	Query_DASubmitter_FullMethodName              = "/fiamma.zkpverify.Query/DASubmitter"
	Query_DASubmissionQueue_FullMethodName        = "/fiamma.zkpverify.Query/DASubmissionQueue"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of PendingProof items.
	PendingProof(ctx context.Context, in *QueryPendingProofRequest, opts ...grpc.CallOption) (*QueryPendingProofResponse, error)
	// Queries proof data by proof id
	ProofData(ctx context.Context, in *QueryProofDataRequest, opts ...grpc.CallOption) (*QueryProofDataResponse, error)
	VerifyResult(ctx context.Context, in *QueryVerifyResultRequest, opts ...grpc.CallOption) (*QueryVerifyResultResponse, error)
	// Queries a list of VerifyResult items by namespace.
	VerifyResultsByNamespace(ctx context.Context, in *QueryVerifyResultsByNamespaceRequest, opts ...grpc.CallOption) (*QueryVerifyResultsByNamespaceResponse, error)
	// Queries a list of PendingProofByNamespace items.
	PendingProofByNamespace(ctx context.Context, in *QueryPendingProofByNamespaceRequest, opts ...grpc.CallOption) (*QueryPendingProofByNamespaceResponse, error)
	DASubmitter(ctx context.Context, in *QueryDASubmitterRequest, opts ...grpc.CallOption) (*QueryDASubmitterResponse, error)
	DASubmissionQueue(ctx context.Context, in *QueryDASubmissionQueueRequest, opts ...grpc.CallOption) (*QueryDASubmissionQueueResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PendingProof(ctx context.Context, in *QueryPendingProofRequest, opts ...grpc.CallOption) (*QueryPendingProofResponse, error) {
	out := new(QueryPendingProofResponse)
	err := c.cc.Invoke(ctx, Query_PendingProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProofData(ctx context.Context, in *QueryProofDataRequest, opts ...grpc.CallOption) (*QueryProofDataResponse, error) {
	out := new(QueryProofDataResponse)
	err := c.cc.Invoke(ctx, Query_ProofData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VerifyResult(ctx context.Context, in *QueryVerifyResultRequest, opts ...grpc.CallOption) (*QueryVerifyResultResponse, error) {
	out := new(QueryVerifyResultResponse)
	err := c.cc.Invoke(ctx, Query_VerifyResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VerifyResultsByNamespace(ctx context.Context, in *QueryVerifyResultsByNamespaceRequest, opts ...grpc.CallOption) (*QueryVerifyResultsByNamespaceResponse, error) {
	out := new(QueryVerifyResultsByNamespaceResponse)
	err := c.cc.Invoke(ctx, Query_VerifyResultsByNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PendingProofByNamespace(ctx context.Context, in *QueryPendingProofByNamespaceRequest, opts ...grpc.CallOption) (*QueryPendingProofByNamespaceResponse, error) {
	out := new(QueryPendingProofByNamespaceResponse)
	err := c.cc.Invoke(ctx, Query_PendingProofByNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DASubmitter(ctx context.Context, in *QueryDASubmitterRequest, opts ...grpc.CallOption) (*QueryDASubmitterResponse, error) {
	out := new(QueryDASubmitterResponse)
	err := c.cc.Invoke(ctx, Query_DASubmitter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DASubmissionQueue(ctx context.Context, in *QueryDASubmissionQueueRequest, opts ...grpc.CallOption) (*QueryDASubmissionQueueResponse, error) {
	out := new(QueryDASubmissionQueueResponse)
	err := c.cc.Invoke(ctx, Query_DASubmissionQueue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of PendingProof items.
	PendingProof(context.Context, *QueryPendingProofRequest) (*QueryPendingProofResponse, error)
	// Queries proof data by proof id
	ProofData(context.Context, *QueryProofDataRequest) (*QueryProofDataResponse, error)
	VerifyResult(context.Context, *QueryVerifyResultRequest) (*QueryVerifyResultResponse, error)
	// Queries a list of VerifyResult items by namespace.
	VerifyResultsByNamespace(context.Context, *QueryVerifyResultsByNamespaceRequest) (*QueryVerifyResultsByNamespaceResponse, error)
	// Queries a list of PendingProofByNamespace items.
	PendingProofByNamespace(context.Context, *QueryPendingProofByNamespaceRequest) (*QueryPendingProofByNamespaceResponse, error)
	DASubmitter(context.Context, *QueryDASubmitterRequest) (*QueryDASubmitterResponse, error)
	DASubmissionQueue(context.Context, *QueryDASubmissionQueueRequest) (*QueryDASubmissionQueueResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) PendingProof(context.Context, *QueryPendingProofRequest) (*QueryPendingProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingProof not implemented")
}
func (UnimplementedQueryServer) ProofData(context.Context, *QueryProofDataRequest) (*QueryProofDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProofData not implemented")
}
func (UnimplementedQueryServer) VerifyResult(context.Context, *QueryVerifyResultRequest) (*QueryVerifyResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyResult not implemented")
}
func (UnimplementedQueryServer) VerifyResultsByNamespace(context.Context, *QueryVerifyResultsByNamespaceRequest) (*QueryVerifyResultsByNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyResultsByNamespace not implemented")
}
func (UnimplementedQueryServer) PendingProofByNamespace(context.Context, *QueryPendingProofByNamespaceRequest) (*QueryPendingProofByNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingProofByNamespace not implemented")
}
func (UnimplementedQueryServer) DASubmitter(context.Context, *QueryDASubmitterRequest) (*QueryDASubmitterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DASubmitter not implemented")
}
func (UnimplementedQueryServer) DASubmissionQueue(context.Context, *QueryDASubmissionQueueRequest) (*QueryDASubmissionQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DASubmissionQueue not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PendingProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPendingProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PendingProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PendingProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PendingProof(ctx, req.(*QueryPendingProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProofData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProofDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProofData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ProofData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProofData(ctx, req.(*QueryProofDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VerifyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVerifyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VerifyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VerifyResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VerifyResult(ctx, req.(*QueryVerifyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VerifyResultsByNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVerifyResultsByNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VerifyResultsByNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VerifyResultsByNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VerifyResultsByNamespace(ctx, req.(*QueryVerifyResultsByNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PendingProofByNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPendingProofByNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PendingProofByNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PendingProofByNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PendingProofByNamespace(ctx, req.(*QueryPendingProofByNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DASubmitter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDASubmitterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DASubmitter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DASubmitter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DASubmitter(ctx, req.(*QueryDASubmitterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DASubmissionQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDASubmissionQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DASubmissionQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DASubmissionQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DASubmissionQueue(ctx, req.(*QueryDASubmissionQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fiamma.zkpverify.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "PendingProof",
			Handler:    _Query_PendingProof_Handler,
		},
		{
			MethodName: "ProofData",
			Handler:    _Query_ProofData_Handler,
		},
		{
			MethodName: "VerifyResult",
			Handler:    _Query_VerifyResult_Handler,
		},
		{
			MethodName: "VerifyResultsByNamespace",
			Handler:    _Query_VerifyResultsByNamespace_Handler,
		},
		{
			MethodName: "PendingProofByNamespace",
			Handler:    _Query_PendingProofByNamespace_Handler,
		},
		{
			MethodName: "DASubmitter",
			Handler:    _Query_DASubmitter_Handler,
		},
		{
			MethodName: "DASubmissionQueue",
			Handler:    _Query_DASubmissionQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fiamma/zkpverify/query.proto",
}
