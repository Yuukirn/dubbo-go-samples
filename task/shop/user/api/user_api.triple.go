/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: user_api.proto
package api

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// This is a compile-time assertion to ensure that this generated file and the Triple package
// are compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of Triple newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of Triple or updating the Triple
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "org.apache.dubbogo.samples.shop.user.api.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceRegisterProcedure is the fully-qualified name of the UserService's Register RPC.
	UserServiceRegisterProcedure = "/org.apache.dubbogo.samples.shop.user.api.UserService/Register"
	// UserServiceLoginProcedure is the fully-qualified name of the UserService's Login RPC.
	UserServiceLoginProcedure = "/org.apache.dubbogo.samples.shop.user.api.UserService/Login"
	// UserServiceTimeoutLoginProcedure is the fully-qualified name of the UserService's TimeoutLogin RPC.
	UserServiceTimeoutLoginProcedure = "/org.apache.dubbogo.samples.shop.user.api.UserService/TimeoutLogin"
	// UserServiceGetInfoProcedure is the fully-qualified name of the UserService's GetInfo RPC.
	UserServiceGetInfoProcedure = "/org.apache.dubbogo.samples.shop.user.api.UserService/GetInfo"
)

var (
	_ UserService = (*UserServiceImpl)(nil)
)

// UserService is a client for the org.apache.dubbogo.samples.shop.user.api.UserService service.
type UserService interface {
	Register(ctx context.Context, req *User, opts ...client.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, req *LoginReq, opts ...client.CallOption) (*User, error)
	TimeoutLogin(ctx context.Context, req *LoginReq, opts ...client.CallOption) (*User, error)
	GetInfo(ctx context.Context, req *GetInfoReq, opts ...client.CallOption) (*User, error)
}

// NewUserService constructs a client for the api.UserService service.
func NewUserService(cli *client.Client, opts ...client.ReferenceOption) (UserService, error) {
	conn, err := cli.DialWithInfo("org.apache.dubbogo.samples.shop.user.api.UserService", &UserService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &UserServiceImpl{
		conn: conn,
	}, nil
}

func SetConsumerService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &UserService_ClientInfo)
}

// UserServiceImpl implements UserService.
type UserServiceImpl struct {
	conn *client.Connection
}

func (c *UserServiceImpl) Register(ctx context.Context, req *User, opts ...client.CallOption) (*RegisterResp, error) {
	resp := new(RegisterResp)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "Register", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *UserServiceImpl) Login(ctx context.Context, req *LoginReq, opts ...client.CallOption) (*User, error) {
	resp := new(User)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "Login", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *UserServiceImpl) TimeoutLogin(ctx context.Context, req *LoginReq, opts ...client.CallOption) (*User, error) {
	resp := new(User)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "TimeoutLogin", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *UserServiceImpl) GetInfo(ctx context.Context, req *GetInfoReq, opts ...client.CallOption) (*User, error) {
	resp := new(User)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "GetInfo", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

var UserService_ClientInfo = client.ClientInfo{
	InterfaceName: "org.apache.dubbogo.samples.shop.user.api.UserService",
	MethodNames:   []string{"Register", "Login", "TimeoutLogin", "GetInfo"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*UserServiceImpl)
		dubboCli.conn = conn
	},
}

// UserServiceHandler is an implementation of the org.apache.dubbogo.samples.shop.user.api.UserService service.
type UserServiceHandler interface {
	Register(context.Context, *User) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*User, error)
	TimeoutLogin(context.Context, *LoginReq) (*User, error)
	GetInfo(context.Context, *GetInfoReq) (*User, error)
}

func RegisterUserServiceHandler(srv *server.Server, hdlr UserServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &UserService_ServiceInfo, opts...)
}

func SetProviderService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &UserService_ServiceInfo)
}

var UserService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "org.apache.dubbogo.samples.shop.user.api.UserService",
	ServiceType:   (*UserServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "Register",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(User)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*User)
				res, err := handler.(UserServiceHandler).Register(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "Login",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(LoginReq)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*LoginReq)
				res, err := handler.(UserServiceHandler).Login(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "TimeoutLogin",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(LoginReq)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*LoginReq)
				res, err := handler.(UserServiceHandler).TimeoutLogin(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "GetInfo",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(GetInfoReq)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*GetInfoReq)
				res, err := handler.(UserServiceHandler).GetInfo(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}