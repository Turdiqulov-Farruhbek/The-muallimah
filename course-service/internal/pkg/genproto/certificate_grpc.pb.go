// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: muallimah-submodule/protos/certificate.proto

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
	CertificateService_CreateCertificate_FullMethodName = "/protos.CertificateService/CreateCertificate"
	CertificateService_GetCertificate_FullMethodName    = "/protos.CertificateService/GetCertificate"
	CertificateService_UpdateCertificate_FullMethodName = "/protos.CertificateService/UpdateCertificate"
	CertificateService_DeleteCertificate_FullMethodName = "/protos.CertificateService/DeleteCertificate"
	CertificateService_ListCertificates_FullMethodName  = "/protos.CertificateService/ListCertificates"
)

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateServiceClient interface {
	CreateCertificate(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetCertificate(ctx context.Context, in *ById, opts ...grpc.CallOption) (*CertificateGet, error)
	UpdateCertificate(ctx context.Context, in *CertificateUpdate, opts ...grpc.CallOption) (*Void, error)
	DeleteCertificate(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	ListCertificates(ctx context.Context, in *CertificateFilter, opts ...grpc.CallOption) (*CertificateList, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) CreateCertificate(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, CertificateService_CreateCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) GetCertificate(ctx context.Context, in *ById, opts ...grpc.CallOption) (*CertificateGet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateGet)
	err := c.cc.Invoke(ctx, CertificateService_GetCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) UpdateCertificate(ctx context.Context, in *CertificateUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, CertificateService_UpdateCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) DeleteCertificate(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, CertificateService_DeleteCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) ListCertificates(ctx context.Context, in *CertificateFilter, opts ...grpc.CallOption) (*CertificateList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateList)
	err := c.cc.Invoke(ctx, CertificateService_ListCertificates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations must embed UnimplementedCertificateServiceServer
// for forward compatibility
type CertificateServiceServer interface {
	CreateCertificate(context.Context, *ById) (*Void, error)
	GetCertificate(context.Context, *ById) (*CertificateGet, error)
	UpdateCertificate(context.Context, *CertificateUpdate) (*Void, error)
	DeleteCertificate(context.Context, *ById) (*Void, error)
	ListCertificates(context.Context, *CertificateFilter) (*CertificateList, error)
	mustEmbedUnimplementedCertificateServiceServer()
}

// UnimplementedCertificateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (UnimplementedCertificateServiceServer) CreateCertificate(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) GetCertificate(context.Context, *ById) (*CertificateGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) UpdateCertificate(context.Context, *CertificateUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) DeleteCertificate(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) ListCertificates(context.Context, *CertificateFilter) (*CertificateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCertificates not implemented")
}
func (UnimplementedCertificateServiceServer) mustEmbedUnimplementedCertificateServiceServer() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_CreateCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_GetCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_GetCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetCertificate(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_UpdateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).UpdateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_UpdateCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).UpdateCertificate(ctx, req.(*CertificateUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_DeleteCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).DeleteCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_DeleteCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).DeleteCertificate(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_ListCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ListCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_ListCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ListCertificates(ctx, req.(*CertificateFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificateService_CreateCertificate_Handler,
		},
		{
			MethodName: "GetCertificate",
			Handler:    _CertificateService_GetCertificate_Handler,
		},
		{
			MethodName: "UpdateCertificate",
			Handler:    _CertificateService_UpdateCertificate_Handler,
		},
		{
			MethodName: "DeleteCertificate",
			Handler:    _CertificateService_DeleteCertificate_Handler,
		},
		{
			MethodName: "ListCertificates",
			Handler:    _CertificateService_ListCertificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "muallimah-submodule/protos/certificate.proto",
}
