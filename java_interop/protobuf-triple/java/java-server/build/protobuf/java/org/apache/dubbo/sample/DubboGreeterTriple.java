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

    package org.apache.dubbo.sample;

import org.apache.dubbo.common.stream.StreamObserver;
import org.apache.dubbo.common.URL;
import org.apache.dubbo.rpc.Invoker;
import org.apache.dubbo.rpc.PathResolver;
import org.apache.dubbo.rpc.RpcException;
import org.apache.dubbo.rpc.ServerService;
import org.apache.dubbo.rpc.TriRpcStatus;
import org.apache.dubbo.rpc.model.MethodDescriptor;
import org.apache.dubbo.rpc.model.ServiceDescriptor;
import org.apache.dubbo.rpc.model.StubMethodDescriptor;
import org.apache.dubbo.rpc.model.StubServiceDescriptor;
import org.apache.dubbo.rpc.stub.BiStreamMethodHandler;
import org.apache.dubbo.rpc.stub.ServerStreamMethodHandler;
import org.apache.dubbo.rpc.stub.StubInvocationUtil;
import org.apache.dubbo.rpc.stub.StubInvoker;
import org.apache.dubbo.rpc.stub.StubMethodHandler;
import org.apache.dubbo.rpc.stub.StubSuppliers;
import org.apache.dubbo.rpc.stub.UnaryStubMethodHandler;

import com.google.protobuf.Message;

import java.util.HashMap;
import java.util.Map;
import java.util.function.BiConsumer;
import java.util.concurrent.CompletableFuture;

public final class DubboGreeterTriple {

    public static final String SERVICE_NAME = Greeter.SERVICE_NAME;

    private static final StubServiceDescriptor serviceDescriptor = new StubServiceDescriptor(SERVICE_NAME,Greeter.class);

    static {
        org.apache.dubbo.rpc.protocol.tri.service.SchemaDescriptorRegistry.addSchemaDescriptor(SERVICE_NAME,HelloWorldProto.getDescriptor());
        StubSuppliers.addSupplier(SERVICE_NAME, DubboGreeterTriple::newStub);
        StubSuppliers.addSupplier(Greeter.JAVA_SERVICE_NAME,  DubboGreeterTriple::newStub);
        StubSuppliers.addDescriptor(SERVICE_NAME, serviceDescriptor);
        StubSuppliers.addDescriptor(Greeter.JAVA_SERVICE_NAME, serviceDescriptor);
    }

    @SuppressWarnings("all")
    public static Greeter newStub(Invoker<?> invoker) {
        return new GreeterStub((Invoker<Greeter>)invoker);
    }

    /**
         * <pre>
         *  Sends a greeting
         * </pre>
         */
    private static final StubMethodDescriptor sayHelloMethod = new StubMethodDescriptor("SayHello",
    org.apache.dubbo.sample.HelloRequest.class, org.apache.dubbo.sample.HelloReply.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), org.apache.dubbo.sample.HelloRequest::parseFrom,
    org.apache.dubbo.sample.HelloReply::parseFrom);

    private static final StubMethodDescriptor sayHelloAsyncMethod = new StubMethodDescriptor("SayHello",
    org.apache.dubbo.sample.HelloRequest.class, java.util.concurrent.CompletableFuture.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), org.apache.dubbo.sample.HelloRequest::parseFrom,
    org.apache.dubbo.sample.HelloReply::parseFrom);

    private static final StubMethodDescriptor sayHelloProxyAsyncMethod = new StubMethodDescriptor("SayHelloAsync",
    org.apache.dubbo.sample.HelloRequest.class, org.apache.dubbo.sample.HelloReply.class, serviceDescriptor, MethodDescriptor.RpcType.UNARY,
    obj -> ((Message) obj).toByteArray(), obj -> ((Message) obj).toByteArray(), org.apache.dubbo.sample.HelloRequest::parseFrom,
    org.apache.dubbo.sample.HelloReply::parseFrom);





    public static class GreeterStub implements Greeter{
        private final Invoker<Greeter> invoker;

        public GreeterStub(Invoker<Greeter> invoker) {
            this.invoker = invoker;
        }

            /**
         * <pre>
         *  Sends a greeting
         * </pre>
         */
        @Override
        public org.apache.dubbo.sample.HelloReply sayHello(org.apache.dubbo.sample.HelloRequest request){
            return StubInvocationUtil.unaryCall(invoker, sayHelloMethod, request);
        }

        public CompletableFuture<org.apache.dubbo.sample.HelloReply> sayHelloAsync(org.apache.dubbo.sample.HelloRequest request){
            return StubInvocationUtil.unaryCall(invoker, sayHelloAsyncMethod, request);
        }

            /**
         * <pre>
         *  Sends a greeting
         * </pre>
         */
        @Override
        public void sayHello(org.apache.dubbo.sample.HelloRequest request, StreamObserver<org.apache.dubbo.sample.HelloReply> responseObserver){
            StubInvocationUtil.unaryCall(invoker, sayHelloMethod , request, responseObserver);
        }



    }

    public static abstract class GreeterImplBase implements Greeter, ServerService<Greeter> {

        private <T, R> BiConsumer<T, StreamObserver<R>> syncToAsync(java.util.function.Function<T, R> syncFun) {
            return new BiConsumer<T, StreamObserver<R>>() {
                @Override
                public void accept(T t, StreamObserver<R> observer) {
                    try {
                        R ret = syncFun.apply(t);
                        observer.onNext(ret);
                        observer.onCompleted();
                    } catch (Throwable e) {
                        observer.onError(e);
                    }
                }
            };
        }

        @Override
        public final Invoker<Greeter> getInvoker(URL url) {
            PathResolver pathResolver = url.getOrDefaultFrameworkModel()
            .getExtensionLoader(PathResolver.class)
            .getDefaultExtension();
            Map<String,StubMethodHandler<?, ?>> handlers = new HashMap<>();

            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/SayHello" );
            pathResolver.addNativeStub( "/" + SERVICE_NAME + "/SayHelloAsync" );

            BiConsumer<org.apache.dubbo.sample.HelloRequest, StreamObserver<org.apache.dubbo.sample.HelloReply>> sayHelloFunc = this::sayHello;
            handlers.put(sayHelloMethod.getMethodName(), new UnaryStubMethodHandler<>(sayHelloFunc));
            BiConsumer<org.apache.dubbo.sample.HelloRequest, StreamObserver<org.apache.dubbo.sample.HelloReply>> sayHelloAsyncFunc = syncToAsync(this::sayHello);
            handlers.put(sayHelloProxyAsyncMethod.getMethodName(), new UnaryStubMethodHandler<>(sayHelloAsyncFunc));




            return new StubInvoker<>(this, url, Greeter.class, handlers);
        }


        @Override
        public org.apache.dubbo.sample.HelloReply sayHello(org.apache.dubbo.sample.HelloRequest request){
            throw unimplementedMethodException(sayHelloMethod);
        }





        @Override
        public final ServiceDescriptor getServiceDescriptor() {
            return serviceDescriptor;
        }
        private RpcException unimplementedMethodException(StubMethodDescriptor methodDescriptor) {
            return TriRpcStatus.UNIMPLEMENTED.withDescription(String.format("Method %s is unimplemented",
                "/" + serviceDescriptor.getInterfaceName() + "/" + methodDescriptor.getMethodName())).asException();
        }
    }

}
