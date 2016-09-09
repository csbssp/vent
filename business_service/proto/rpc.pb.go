// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Register service

type RegisterClient interface {
	Register(ctx context.Context, in *RegisterC2S, opts ...grpc.CallOption) (*RegisterS2C, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *RegisterC2S, opts ...grpc.CallOption) (*RegisterS2C, error) {
	out := new(RegisterS2C)
	err := grpc.Invoke(ctx, "/proto.Register/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Register service

type RegisterServer interface {
	Register(context.Context, *RegisterC2S) (*RegisterS2C, error)
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Register/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*RegisterC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for Login service

type LoginClient interface {
	Login(ctx context.Context, in *LoginC2S, opts ...grpc.CallOption) (*LoginS2C, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginC2S, opts ...grpc.CallOption) (*LoginS2C, error) {
	out := new(LoginS2C)
	err := grpc.Invoke(ctx, "/proto.Login/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginServer interface {
	Login(context.Context, *LoginC2S) (*LoginS2C, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for UserInfoManager service

type UserInfoManagerClient interface {
	UserInfoModify(ctx context.Context, in *UserInfoModifyC2S, opts ...grpc.CallOption) (*UserInfoModifyS2C, error)
	UserInfoGet(ctx context.Context, in *UserInfoGetC2S, opts ...grpc.CallOption) (*UserInfoGetS2C, error)
}

type userInfoManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoManagerClient(cc *grpc.ClientConn) UserInfoManagerClient {
	return &userInfoManagerClient{cc}
}

func (c *userInfoManagerClient) UserInfoModify(ctx context.Context, in *UserInfoModifyC2S, opts ...grpc.CallOption) (*UserInfoModifyS2C, error) {
	out := new(UserInfoModifyS2C)
	err := grpc.Invoke(ctx, "/proto.UserInfoManager/UserInfoModify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoManagerClient) UserInfoGet(ctx context.Context, in *UserInfoGetC2S, opts ...grpc.CallOption) (*UserInfoGetS2C, error) {
	out := new(UserInfoGetS2C)
	err := grpc.Invoke(ctx, "/proto.UserInfoManager/UserInfoGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserInfoManager service

type UserInfoManagerServer interface {
	UserInfoModify(context.Context, *UserInfoModifyC2S) (*UserInfoModifyS2C, error)
	UserInfoGet(context.Context, *UserInfoGetC2S) (*UserInfoGetS2C, error)
}

func RegisterUserInfoManagerServer(s *grpc.Server, srv UserInfoManagerServer) {
	s.RegisterService(&_UserInfoManager_serviceDesc, srv)
}

func _UserInfoManager_UserInfoModify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoModifyC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoManagerServer).UserInfoModify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserInfoManager/UserInfoModify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoManagerServer).UserInfoModify(ctx, req.(*UserInfoModifyC2S))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoManager_UserInfoGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoGetC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoManagerServer).UserInfoGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserInfoManager/UserInfoGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoManagerServer).UserInfoGet(ctx, req.(*UserInfoGetC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserInfoManager",
	HandlerType: (*UserInfoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfoModify",
			Handler:    _UserInfoManager_UserInfoModify_Handler,
		},
		{
			MethodName: "UserInfoGet",
			Handler:    _UserInfoManager_UserInfoGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto1.RegisterFile("rpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x3c, 0xc9, 0x39, 0x99, 0xa9, 0x79,
	0x25, 0x10, 0x41, 0x23, 0x07, 0x2e, 0x8e, 0xa0, 0xd4, 0xf4, 0xcc, 0xe2, 0x92, 0xd4, 0x22, 0x21,
	0x13, 0x24, 0xb6, 0x10, 0x44, 0x5e, 0x0f, 0x26, 0xe0, 0x6c, 0x54, 0x2c, 0x85, 0x2e, 0x16, 0x6c,
	0x94, 0xac, 0xc4, 0x60, 0x64, 0xc2, 0xc5, 0xea, 0x93, 0x9f, 0x9e, 0x99, 0x27, 0xa4, 0x0d, 0x63,
	0xf0, 0x43, 0xd5, 0x81, 0x79, 0x20, 0x8d, 0x28, 0x02, 0x10, 0x5d, 0x33, 0x18, 0xb9, 0xf8, 0x43,
	0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x7d, 0x13, 0xf3, 0x12, 0xd3, 0x81, 0x76, 0xba, 0x71,
	0xf1, 0xc1, 0x85, 0xf2, 0x53, 0x32, 0xd3, 0x2a, 0x85, 0x24, 0xa0, 0x1a, 0x51, 0x85, 0x9d, 0x8d,
	0x82, 0xa5, 0xb0, 0xcb, 0x04, 0x1b, 0x39, 0x2b, 0x31, 0x08, 0xd9, 0x72, 0x71, 0xc3, 0x84, 0xdd,
	0x53, 0x4b, 0x84, 0x44, 0xd1, 0x94, 0x02, 0xc5, 0x40, 0x26, 0x60, 0x11, 0x06, 0x6b, 0x4f, 0x62,
	0x03, 0x8b, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x1e, 0x8e, 0x4c, 0x3b, 0x01, 0x00,
	0x00,
}
