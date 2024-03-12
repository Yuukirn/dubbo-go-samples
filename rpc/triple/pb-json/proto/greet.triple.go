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
// Source: greet.proto
package greet

import (
	"context"
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
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "greet.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServiceGreetProcedure is the fully-qualified name of the GreetService's Greet RPC.
	GreetServiceGreetProcedure = "/greet.GreetService/Greet"
)

var (
	_ GreetService = (*GreetServiceImpl)(nil)
)

// GreetService is a client for the greet.GreetService service.
type GreetService interface {
	Greet(ctx context.Context, req *GreetRequest, opts ...client.CallOption) (*GreetResponse, error)
}

// NewGreetService constructs a client for the greet.GreetService service.
func NewGreetService(cli *client.Client, opts ...client.ReferenceOption) (GreetService, error) {
	conn, err := cli.DialWithInfo("greet.GreetService", &GreetService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &GreetServiceImpl{
		conn: conn,
	}, nil
}

func SetConsumerService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &GreetService_ClientInfo)
}

// GreetServiceImpl implements GreetService.
type GreetServiceImpl struct {
	conn *client.Connection
}

func (c *GreetServiceImpl) Greet(ctx context.Context, req *GreetRequest, opts ...client.CallOption) (*GreetResponse, error) {
	resp := new(GreetResponse)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "Greet", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

var GreetService_ClientInfo = client.ClientInfo{
	InterfaceName: "greet.GreetService",
	MethodNames:   []string{"Greet"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*GreetServiceImpl)
		dubboCli.conn = conn
	},
}

// GreetServiceHandler is an implementation of the greet.GreetService service.
type GreetServiceHandler interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterGreetServiceHandler(srv *server.Server, hdlr GreetServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &GreetService_ServiceInfo, opts...)
}

func SetProviderService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &GreetService_ServiceInfo)
}

var GreetService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "greet.GreetService",
	ServiceType:   (*GreetServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "Greet",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(GreetRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*GreetRequest)
				res, err := handler.(GreetServiceHandler).Greet(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
