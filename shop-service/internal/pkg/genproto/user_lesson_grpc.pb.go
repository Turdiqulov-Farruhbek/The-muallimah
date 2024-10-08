// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: muallimah-submodule/protos/user_lesson.proto

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
	UserLessonService_CreateUserLesson_FullMethodName = "/protos.UserLessonService/CreateUserLesson"
	UserLessonService_GetUserLesson_FullMethodName    = "/protos.UserLessonService/GetUserLesson"
	UserLessonService_UpdateUserLesson_FullMethodName = "/protos.UserLessonService/UpdateUserLesson"
	UserLessonService_DeleteUserLesson_FullMethodName = "/protos.UserLessonService/DeleteUserLesson"
	UserLessonService_ListUserLessons_FullMethodName  = "/protos.UserLessonService/ListUserLessons"
)

// UserLessonServiceClient is the client API for UserLessonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLessonServiceClient interface {
	CreateUserLesson(ctx context.Context, in *UserLessonCreateReq, opts ...grpc.CallOption) (*Void, error)
	GetUserLesson(ctx context.Context, in *UserLessonGetReq, opts ...grpc.CallOption) (*UserLessonGetRes, error)
	UpdateUserLesson(ctx context.Context, in *UserLessonUpdateReq, opts ...grpc.CallOption) (*Void, error)
	DeleteUserLesson(ctx context.Context, in *UserLessonDeleteReq, opts ...grpc.CallOption) (*Void, error)
	ListUserLessons(ctx context.Context, in *UserLessonListsReq, opts ...grpc.CallOption) (*UserLessonListsRes, error)
}

type userLessonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLessonServiceClient(cc grpc.ClientConnInterface) UserLessonServiceClient {
	return &userLessonServiceClient{cc}
}

func (c *userLessonServiceClient) CreateUserLesson(ctx context.Context, in *UserLessonCreateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_CreateUserLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) GetUserLesson(ctx context.Context, in *UserLessonGetReq, opts ...grpc.CallOption) (*UserLessonGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLessonGetRes)
	err := c.cc.Invoke(ctx, UserLessonService_GetUserLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) UpdateUserLesson(ctx context.Context, in *UserLessonUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_UpdateUserLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) DeleteUserLesson(ctx context.Context, in *UserLessonDeleteReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserLessonService_DeleteUserLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLessonServiceClient) ListUserLessons(ctx context.Context, in *UserLessonListsReq, opts ...grpc.CallOption) (*UserLessonListsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLessonListsRes)
	err := c.cc.Invoke(ctx, UserLessonService_ListUserLessons_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLessonServiceServer is the server API for UserLessonService service.
// All implementations must embed UnimplementedUserLessonServiceServer
// for forward compatibility
type UserLessonServiceServer interface {
	CreateUserLesson(context.Context, *UserLessonCreateReq) (*Void, error)
	GetUserLesson(context.Context, *UserLessonGetReq) (*UserLessonGetRes, error)
	UpdateUserLesson(context.Context, *UserLessonUpdateReq) (*Void, error)
	DeleteUserLesson(context.Context, *UserLessonDeleteReq) (*Void, error)
	ListUserLessons(context.Context, *UserLessonListsReq) (*UserLessonListsRes, error)
	mustEmbedUnimplementedUserLessonServiceServer()
}

// UnimplementedUserLessonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserLessonServiceServer struct {
}

func (UnimplementedUserLessonServiceServer) CreateUserLesson(context.Context, *UserLessonCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserLesson not implemented")
}
func (UnimplementedUserLessonServiceServer) GetUserLesson(context.Context, *UserLessonGetReq) (*UserLessonGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLesson not implemented")
}
func (UnimplementedUserLessonServiceServer) UpdateUserLesson(context.Context, *UserLessonUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserLesson not implemented")
}
func (UnimplementedUserLessonServiceServer) DeleteUserLesson(context.Context, *UserLessonDeleteReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserLesson not implemented")
}
func (UnimplementedUserLessonServiceServer) ListUserLessons(context.Context, *UserLessonListsReq) (*UserLessonListsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserLessons not implemented")
}
func (UnimplementedUserLessonServiceServer) mustEmbedUnimplementedUserLessonServiceServer() {}

// UnsafeUserLessonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLessonServiceServer will
// result in compilation errors.
type UnsafeUserLessonServiceServer interface {
	mustEmbedUnimplementedUserLessonServiceServer()
}

func RegisterUserLessonServiceServer(s grpc.ServiceRegistrar, srv UserLessonServiceServer) {
	s.RegisterService(&UserLessonService_ServiceDesc, srv)
}

func _UserLessonService_CreateUserLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).CreateUserLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_CreateUserLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).CreateUserLesson(ctx, req.(*UserLessonCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_GetUserLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).GetUserLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_GetUserLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).GetUserLesson(ctx, req.(*UserLessonGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_UpdateUserLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).UpdateUserLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_UpdateUserLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).UpdateUserLesson(ctx, req.(*UserLessonUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_DeleteUserLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).DeleteUserLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_DeleteUserLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).DeleteUserLesson(ctx, req.(*UserLessonDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLessonService_ListUserLessons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLessonListsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLessonServiceServer).ListUserLessons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLessonService_ListUserLessons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLessonServiceServer).ListUserLessons(ctx, req.(*UserLessonListsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLessonService_ServiceDesc is the grpc.ServiceDesc for UserLessonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLessonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserLessonService",
	HandlerType: (*UserLessonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserLesson",
			Handler:    _UserLessonService_CreateUserLesson_Handler,
		},
		{
			MethodName: "GetUserLesson",
			Handler:    _UserLessonService_GetUserLesson_Handler,
		},
		{
			MethodName: "UpdateUserLesson",
			Handler:    _UserLessonService_UpdateUserLesson_Handler,
		},
		{
			MethodName: "DeleteUserLesson",
			Handler:    _UserLessonService_DeleteUserLesson_Handler,
		},
		{
			MethodName: "ListUserLessons",
			Handler:    _UserLessonService_ListUserLessons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "muallimah-submodule/protos/user_lesson.proto",
}
