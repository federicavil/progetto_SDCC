package it.progetto.login.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.50.2)",
    comments = "Source: src/main/resources/loginService.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class LoginServiceGrpc {

  private LoginServiceGrpc() {}

  public static final String SERVICE_NAME = "it.progetto.progetto_sdcc.proto.LoginService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<it.progetto.login.grpc.LoginRequest,
      it.progetto.login.grpc.LoginResponse> getLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "login",
      requestType = it.progetto.login.grpc.LoginRequest.class,
      responseType = it.progetto.login.grpc.LoginResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.login.grpc.LoginRequest,
      it.progetto.login.grpc.LoginResponse> getLoginMethod() {
    io.grpc.MethodDescriptor<it.progetto.login.grpc.LoginRequest, it.progetto.login.grpc.LoginResponse> getLoginMethod;
    if ((getLoginMethod = LoginServiceGrpc.getLoginMethod) == null) {
      synchronized (LoginServiceGrpc.class) {
        if ((getLoginMethod = LoginServiceGrpc.getLoginMethod) == null) {
          LoginServiceGrpc.getLoginMethod = getLoginMethod =
              io.grpc.MethodDescriptor.<it.progetto.login.grpc.LoginRequest, it.progetto.login.grpc.LoginResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "login"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.LoginRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.LoginResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LoginServiceMethodDescriptorSupplier("login"))
              .build();
        }
      }
    }
    return getLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<it.progetto.login.grpc.LoginRequest,
      it.progetto.login.grpc.LoginResponse> getSigninMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "signin",
      requestType = it.progetto.login.grpc.LoginRequest.class,
      responseType = it.progetto.login.grpc.LoginResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.login.grpc.LoginRequest,
      it.progetto.login.grpc.LoginResponse> getSigninMethod() {
    io.grpc.MethodDescriptor<it.progetto.login.grpc.LoginRequest, it.progetto.login.grpc.LoginResponse> getSigninMethod;
    if ((getSigninMethod = LoginServiceGrpc.getSigninMethod) == null) {
      synchronized (LoginServiceGrpc.class) {
        if ((getSigninMethod = LoginServiceGrpc.getSigninMethod) == null) {
          LoginServiceGrpc.getSigninMethod = getSigninMethod =
              io.grpc.MethodDescriptor.<it.progetto.login.grpc.LoginRequest, it.progetto.login.grpc.LoginResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "signin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.LoginRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.LoginResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LoginServiceMethodDescriptorSupplier("signin"))
              .build();
        }
      }
    }
    return getSigninMethod;
  }

  private static volatile io.grpc.MethodDescriptor<it.progetto.login.grpc.CheckRequest,
      it.progetto.login.grpc.CheckResponse> getCheckLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "checkLogin",
      requestType = it.progetto.login.grpc.CheckRequest.class,
      responseType = it.progetto.login.grpc.CheckResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.login.grpc.CheckRequest,
      it.progetto.login.grpc.CheckResponse> getCheckLoginMethod() {
    io.grpc.MethodDescriptor<it.progetto.login.grpc.CheckRequest, it.progetto.login.grpc.CheckResponse> getCheckLoginMethod;
    if ((getCheckLoginMethod = LoginServiceGrpc.getCheckLoginMethod) == null) {
      synchronized (LoginServiceGrpc.class) {
        if ((getCheckLoginMethod = LoginServiceGrpc.getCheckLoginMethod) == null) {
          LoginServiceGrpc.getCheckLoginMethod = getCheckLoginMethod =
              io.grpc.MethodDescriptor.<it.progetto.login.grpc.CheckRequest, it.progetto.login.grpc.CheckResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "checkLogin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.CheckRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.CheckResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LoginServiceMethodDescriptorSupplier("checkLogin"))
              .build();
        }
      }
    }
    return getCheckLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<it.progetto.login.grpc.CheckRequest,
      it.progetto.login.grpc.CheckResponse> getLogOutMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "logOut",
      requestType = it.progetto.login.grpc.CheckRequest.class,
      responseType = it.progetto.login.grpc.CheckResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.login.grpc.CheckRequest,
      it.progetto.login.grpc.CheckResponse> getLogOutMethod() {
    io.grpc.MethodDescriptor<it.progetto.login.grpc.CheckRequest, it.progetto.login.grpc.CheckResponse> getLogOutMethod;
    if ((getLogOutMethod = LoginServiceGrpc.getLogOutMethod) == null) {
      synchronized (LoginServiceGrpc.class) {
        if ((getLogOutMethod = LoginServiceGrpc.getLogOutMethod) == null) {
          LoginServiceGrpc.getLogOutMethod = getLogOutMethod =
              io.grpc.MethodDescriptor.<it.progetto.login.grpc.CheckRequest, it.progetto.login.grpc.CheckResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "logOut"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.CheckRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.CheckResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LoginServiceMethodDescriptorSupplier("logOut"))
              .build();
        }
      }
    }
    return getLogOutMethod;
  }

  private static volatile io.grpc.MethodDescriptor<it.progetto.login.grpc.ProfileRequest,
      it.progetto.login.grpc.ProfileResponse> getGetProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "getProfile",
      requestType = it.progetto.login.grpc.ProfileRequest.class,
      responseType = it.progetto.login.grpc.ProfileResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.login.grpc.ProfileRequest,
      it.progetto.login.grpc.ProfileResponse> getGetProfileMethod() {
    io.grpc.MethodDescriptor<it.progetto.login.grpc.ProfileRequest, it.progetto.login.grpc.ProfileResponse> getGetProfileMethod;
    if ((getGetProfileMethod = LoginServiceGrpc.getGetProfileMethod) == null) {
      synchronized (LoginServiceGrpc.class) {
        if ((getGetProfileMethod = LoginServiceGrpc.getGetProfileMethod) == null) {
          LoginServiceGrpc.getGetProfileMethod = getGetProfileMethod =
              io.grpc.MethodDescriptor.<it.progetto.login.grpc.ProfileRequest, it.progetto.login.grpc.ProfileResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "getProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.ProfileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.ProfileResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LoginServiceMethodDescriptorSupplier("getProfile"))
              .build();
        }
      }
    }
    return getGetProfileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<it.progetto.login.grpc.ProfileRequest,
      it.progetto.login.grpc.ProfileResponse> getUpdateProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "updateProfile",
      requestType = it.progetto.login.grpc.ProfileRequest.class,
      responseType = it.progetto.login.grpc.ProfileResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.login.grpc.ProfileRequest,
      it.progetto.login.grpc.ProfileResponse> getUpdateProfileMethod() {
    io.grpc.MethodDescriptor<it.progetto.login.grpc.ProfileRequest, it.progetto.login.grpc.ProfileResponse> getUpdateProfileMethod;
    if ((getUpdateProfileMethod = LoginServiceGrpc.getUpdateProfileMethod) == null) {
      synchronized (LoginServiceGrpc.class) {
        if ((getUpdateProfileMethod = LoginServiceGrpc.getUpdateProfileMethod) == null) {
          LoginServiceGrpc.getUpdateProfileMethod = getUpdateProfileMethod =
              io.grpc.MethodDescriptor.<it.progetto.login.grpc.ProfileRequest, it.progetto.login.grpc.ProfileResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "updateProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.ProfileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.login.grpc.ProfileResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LoginServiceMethodDescriptorSupplier("updateProfile"))
              .build();
        }
      }
    }
    return getUpdateProfileMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static LoginServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LoginServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LoginServiceStub>() {
        @java.lang.Override
        public LoginServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LoginServiceStub(channel, callOptions);
        }
      };
    return LoginServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static LoginServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LoginServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LoginServiceBlockingStub>() {
        @java.lang.Override
        public LoginServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LoginServiceBlockingStub(channel, callOptions);
        }
      };
    return LoginServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static LoginServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LoginServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LoginServiceFutureStub>() {
        @java.lang.Override
        public LoginServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LoginServiceFutureStub(channel, callOptions);
        }
      };
    return LoginServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class LoginServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void login(it.progetto.login.grpc.LoginRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.LoginResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLoginMethod(), responseObserver);
    }

    /**
     */
    public void signin(it.progetto.login.grpc.LoginRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.LoginResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSigninMethod(), responseObserver);
    }

    /**
     */
    public void checkLogin(it.progetto.login.grpc.CheckRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.CheckResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCheckLoginMethod(), responseObserver);
    }

    /**
     */
    public void logOut(it.progetto.login.grpc.CheckRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.CheckResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLogOutMethod(), responseObserver);
    }

    /**
     */
    public void getProfile(it.progetto.login.grpc.ProfileRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.ProfileResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetProfileMethod(), responseObserver);
    }

    /**
     */
    public void updateProfile(it.progetto.login.grpc.ProfileRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.ProfileResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateProfileMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getLoginMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.login.grpc.LoginRequest,
                it.progetto.login.grpc.LoginResponse>(
                  this, METHODID_LOGIN)))
          .addMethod(
            getSigninMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.login.grpc.LoginRequest,
                it.progetto.login.grpc.LoginResponse>(
                  this, METHODID_SIGNIN)))
          .addMethod(
            getCheckLoginMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.login.grpc.CheckRequest,
                it.progetto.login.grpc.CheckResponse>(
                  this, METHODID_CHECK_LOGIN)))
          .addMethod(
            getLogOutMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.login.grpc.CheckRequest,
                it.progetto.login.grpc.CheckResponse>(
                  this, METHODID_LOG_OUT)))
          .addMethod(
            getGetProfileMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.login.grpc.ProfileRequest,
                it.progetto.login.grpc.ProfileResponse>(
                  this, METHODID_GET_PROFILE)))
          .addMethod(
            getUpdateProfileMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.login.grpc.ProfileRequest,
                it.progetto.login.grpc.ProfileResponse>(
                  this, METHODID_UPDATE_PROFILE)))
          .build();
    }
  }

  /**
   */
  public static final class LoginServiceStub extends io.grpc.stub.AbstractAsyncStub<LoginServiceStub> {
    private LoginServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LoginServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LoginServiceStub(channel, callOptions);
    }

    /**
     */
    public void login(it.progetto.login.grpc.LoginRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.LoginResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void signin(it.progetto.login.grpc.LoginRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.LoginResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSigninMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void checkLogin(it.progetto.login.grpc.CheckRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.CheckResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCheckLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void logOut(it.progetto.login.grpc.CheckRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.CheckResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLogOutMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getProfile(it.progetto.login.grpc.ProfileRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.ProfileResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetProfileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateProfile(it.progetto.login.grpc.ProfileRequest request,
        io.grpc.stub.StreamObserver<it.progetto.login.grpc.ProfileResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateProfileMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class LoginServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<LoginServiceBlockingStub> {
    private LoginServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LoginServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LoginServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public it.progetto.login.grpc.LoginResponse login(it.progetto.login.grpc.LoginRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLoginMethod(), getCallOptions(), request);
    }

    /**
     */
    public it.progetto.login.grpc.LoginResponse signin(it.progetto.login.grpc.LoginRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSigninMethod(), getCallOptions(), request);
    }

    /**
     */
    public it.progetto.login.grpc.CheckResponse checkLogin(it.progetto.login.grpc.CheckRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCheckLoginMethod(), getCallOptions(), request);
    }

    /**
     */
    public it.progetto.login.grpc.CheckResponse logOut(it.progetto.login.grpc.CheckRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLogOutMethod(), getCallOptions(), request);
    }

    /**
     */
    public it.progetto.login.grpc.ProfileResponse getProfile(it.progetto.login.grpc.ProfileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetProfileMethod(), getCallOptions(), request);
    }

    /**
     */
    public it.progetto.login.grpc.ProfileResponse updateProfile(it.progetto.login.grpc.ProfileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateProfileMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class LoginServiceFutureStub extends io.grpc.stub.AbstractFutureStub<LoginServiceFutureStub> {
    private LoginServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LoginServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LoginServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.login.grpc.LoginResponse> login(
        it.progetto.login.grpc.LoginRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLoginMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.login.grpc.LoginResponse> signin(
        it.progetto.login.grpc.LoginRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSigninMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.login.grpc.CheckResponse> checkLogin(
        it.progetto.login.grpc.CheckRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCheckLoginMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.login.grpc.CheckResponse> logOut(
        it.progetto.login.grpc.CheckRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLogOutMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.login.grpc.ProfileResponse> getProfile(
        it.progetto.login.grpc.ProfileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetProfileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.login.grpc.ProfileResponse> updateProfile(
        it.progetto.login.grpc.ProfileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateProfileMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LOGIN = 0;
  private static final int METHODID_SIGNIN = 1;
  private static final int METHODID_CHECK_LOGIN = 2;
  private static final int METHODID_LOG_OUT = 3;
  private static final int METHODID_GET_PROFILE = 4;
  private static final int METHODID_UPDATE_PROFILE = 5;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final LoginServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(LoginServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LOGIN:
          serviceImpl.login((it.progetto.login.grpc.LoginRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.login.grpc.LoginResponse>) responseObserver);
          break;
        case METHODID_SIGNIN:
          serviceImpl.signin((it.progetto.login.grpc.LoginRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.login.grpc.LoginResponse>) responseObserver);
          break;
        case METHODID_CHECK_LOGIN:
          serviceImpl.checkLogin((it.progetto.login.grpc.CheckRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.login.grpc.CheckResponse>) responseObserver);
          break;
        case METHODID_LOG_OUT:
          serviceImpl.logOut((it.progetto.login.grpc.CheckRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.login.grpc.CheckResponse>) responseObserver);
          break;
        case METHODID_GET_PROFILE:
          serviceImpl.getProfile((it.progetto.login.grpc.ProfileRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.login.grpc.ProfileResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PROFILE:
          serviceImpl.updateProfile((it.progetto.login.grpc.ProfileRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.login.grpc.ProfileResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class LoginServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    LoginServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return it.progetto.login.grpc.LoginServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("LoginService");
    }
  }

  private static final class LoginServiceFileDescriptorSupplier
      extends LoginServiceBaseDescriptorSupplier {
    LoginServiceFileDescriptorSupplier() {}
  }

  private static final class LoginServiceMethodDescriptorSupplier
      extends LoginServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    LoginServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (LoginServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new LoginServiceFileDescriptorSupplier())
              .addMethod(getLoginMethod())
              .addMethod(getSigninMethod())
              .addMethod(getCheckLoginMethod())
              .addMethod(getLogOutMethod())
              .addMethod(getGetProfileMethod())
              .addMethod(getUpdateProfileMethod())
              .build();
        }
      }
    }
    return result;
  }
}
