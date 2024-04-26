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

package main

import (
	"context"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"github.com/dubbogo/gost/log/logger"
)

import (
	greet_gen "github.com/apache/dubbo-go-samples/non_idl_with_java/proto"
)

func main() {
	ctx := context.Background()

	ins, err := dubbo.NewInstance(
		dubbo.WithName("dubbo_rpc_hessian2_client"),
		dubbo.WithRegistry(
			registry.WithZookeeper(),
			registry.WithAddress("127.0.0.1:2181"),
		),
	)
	if err != nil {
		panic(err)
	}

	cli, err := ins.NewClient(
		client.WithClientURL("127.0.0.1:20000"),
		client.WithClientProtocolDubbo(),
		client.WithClientSerialization(constant.Hessian2Serialization),
	)
	if err != nil {
		panic(err)
	}

	svc, err := greet_gen.NewGreetService(cli)
	if err != nil {
		panic(err)
	}

	resp, err := svc.Greet(ctx, &greet_gen.GreetRequest{
		Name: "dubbo-go",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof("Greet response: %s", resp.Greeting)

	resp2, err := svc.Greet2(ctx, &greet_gen.Greet2Request{Name: "dubbo-go"})
	if err != nil {
		panic(err)
	}
	logger.Infof("Greet response: %s", resp2.Greeting)

	resp3, err := svc.Greet3(ctx, "dubbo-go", 20)
	if err != nil {
		panic(err)
	}
	logger.Infof("Greet3 response: %s", resp3)
}
