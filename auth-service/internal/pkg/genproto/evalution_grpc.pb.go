// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: muallimah-submodule/protos/evalution.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EvaluationService_CreateEvaluation_FullMethodName = "/protos.EvaluationService/CreateEvaluation"
	EvaluationService_GetEvaluation_FullMethodName    = "/protos.EvaluationService/GetEvaluation"
	EvaluationService_DeleteEvaluation_FullMethodName = "/protos.EvaluationService/DeleteEvaluation"
	EvaluationService_ListEvaluations_FullMethodName  = "/protos.EvaluationService/ListEvaluations"
)

// EvaluationServiceClient is the client API for EvaluationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvaluationServiceClient interface {
	CreateEvaluation(ctx context.Context, in *EvaluationCreate, opts ...grpc.CallOption) (*Void, error)
	GetEvaluation(ctx context.Context, in *ById, opts ...grpc.CallOption) (*EvaluationGet, error)
	DeleteEvaluation(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	ListEvaluations(ctx context.Context, in *EvaluationFilter, opts ...grpc.CallOption) (*EvaluationList, error)
}

type evaluationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEvaluationServiceClient(cc grpc.ClientConnInterface) EvaluationServiceClient {
	return &evaluationServiceClient{cc}
}

func (c *evaluationServiceClient) CreateEvaluation(ctx context.Context, in *EvaluationCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, EvaluationService_CreateEvaluation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluationServiceClient) GetEvaluation(ctx context.Context, in *ById, opts ...grpc.CallOption) (*EvaluationGet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EvaluationGet)
	err := c.cc.Invoke(ctx, EvaluationService_GetEvaluation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluationServiceClient) DeleteEvaluation(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, EvaluationService_DeleteEvaluation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaluationServiceClient) ListEvaluations(ctx context.Context, in *EvaluationFilter, opts ...grpc.CallOption) (*EvaluationList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EvaluationList)
	err := c.cc.Invoke(ctx, EvaluationService_ListEvaluations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluationServiceServer is the server API for EvaluationService service.
// All implementations must embed UnimplementedEvaluationServiceServer
// for forward compatibility
type EvaluationServiceServer interface {
	CreateEvaluation(context.Context, *EvaluationCreate) (*Void, error)
	GetEvaluation(context.Context, *ById) (*EvaluationGet, error)
	DeleteEvaluation(context.Context, *ById) (*Void, error)
	ListEvaluations(context.Context, *EvaluationFilter) (*EvaluationList, error)
	mustEmbedUnimplementedEvaluationServiceServer()
}

// UnimplementedEvaluationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEvaluationServiceServer struct {
}

func (UnimplementedEvaluationServiceServer) CreateEvaluation(context.Context, *EvaluationCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvaluation not implemented")
}
func (UnimplementedEvaluationServiceServer) GetEvaluation(context.Context, *ById) (*EvaluationGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvaluation not implemented")
}
func (UnimplementedEvaluationServiceServer) DeleteEvaluation(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvaluation not implemented")
}
func (UnimplementedEvaluationServiceServer) ListEvaluations(context.Context, *EvaluationFilter) (*EvaluationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvaluations not implemented")
}
func (UnimplementedEvaluationServiceServer) mustEmbedUnimplementedEvaluationServiceServer() {}

// UnsafeEvaluationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvaluationServiceServer will
// result in compilation errors.
type UnsafeEvaluationServiceServer interface {
	mustEmbedUnimplementedEvaluationServiceServer()
}

func RegisterEvaluationServiceServer(s grpc.ServiceRegistrar, srv EvaluationServiceServer) {
	s.RegisterService(&EvaluationService_ServiceDesc, srv)
}

func _EvaluationService_CreateEvaluation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluationServiceServer).CreateEvaluation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaluationService_CreateEvaluation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluationServiceServer).CreateEvaluation(ctx, req.(*EvaluationCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluationService_GetEvaluation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluationServiceServer).GetEvaluation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaluationService_GetEvaluation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluationServiceServer).GetEvaluation(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluationService_DeleteEvaluation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluationServiceServer).DeleteEvaluation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaluationService_DeleteEvaluation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluationServiceServer).DeleteEvaluation(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaluationService_ListEvaluations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluationServiceServer).ListEvaluations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaluationService_ListEvaluations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluationServiceServer).ListEvaluations(ctx, req.(*EvaluationFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// EvaluationService_ServiceDesc is the grpc.ServiceDesc for EvaluationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EvaluationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.EvaluationService",
	HandlerType: (*EvaluationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvaluation",
			Handler:    _EvaluationService_CreateEvaluation_Handler,
		},
		{
			MethodName: "GetEvaluation",
			Handler:    _EvaluationService_GetEvaluation_Handler,
		},
		{
			MethodName: "DeleteEvaluation",
			Handler:    _EvaluationService_DeleteEvaluation_Handler,
		},
		{
			MethodName: "ListEvaluations",
			Handler:    _EvaluationService_ListEvaluations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "muallimah-submodule/protos/evalution.proto",
}
