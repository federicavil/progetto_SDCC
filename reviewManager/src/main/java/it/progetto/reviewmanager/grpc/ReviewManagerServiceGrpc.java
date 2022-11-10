package it.progetto.reviewmanager.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.50.2)",
    comments = "Source: src/main/resources/reviewManager.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class ReviewManagerServiceGrpc {

  private ReviewManagerServiceGrpc() {}

  public static final String SERVICE_NAME = "it.progetto.progetto_sdcc.proto.ReviewManagerService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<it.progetto.reviewmanager.grpc.AddReviewRequest,
      it.progetto.reviewmanager.grpc.AddReviewResponse> getAddReviewMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "addReview",
      requestType = it.progetto.reviewmanager.grpc.AddReviewRequest.class,
      responseType = it.progetto.reviewmanager.grpc.AddReviewResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.reviewmanager.grpc.AddReviewRequest,
      it.progetto.reviewmanager.grpc.AddReviewResponse> getAddReviewMethod() {
    io.grpc.MethodDescriptor<it.progetto.reviewmanager.grpc.AddReviewRequest, it.progetto.reviewmanager.grpc.AddReviewResponse> getAddReviewMethod;
    if ((getAddReviewMethod = ReviewManagerServiceGrpc.getAddReviewMethod) == null) {
      synchronized (ReviewManagerServiceGrpc.class) {
        if ((getAddReviewMethod = ReviewManagerServiceGrpc.getAddReviewMethod) == null) {
          ReviewManagerServiceGrpc.getAddReviewMethod = getAddReviewMethod =
              io.grpc.MethodDescriptor.<it.progetto.reviewmanager.grpc.AddReviewRequest, it.progetto.reviewmanager.grpc.AddReviewResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "addReview"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.reviewmanager.grpc.AddReviewRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.reviewmanager.grpc.AddReviewResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ReviewManagerServiceMethodDescriptorSupplier("addReview"))
              .build();
        }
      }
    }
    return getAddReviewMethod;
  }

  private static volatile io.grpc.MethodDescriptor<it.progetto.reviewmanager.grpc.GetReviewsRequest,
      it.progetto.reviewmanager.grpc.GetReviewsResponse> getGetReviewsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "getReviews",
      requestType = it.progetto.reviewmanager.grpc.GetReviewsRequest.class,
      responseType = it.progetto.reviewmanager.grpc.GetReviewsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<it.progetto.reviewmanager.grpc.GetReviewsRequest,
      it.progetto.reviewmanager.grpc.GetReviewsResponse> getGetReviewsMethod() {
    io.grpc.MethodDescriptor<it.progetto.reviewmanager.grpc.GetReviewsRequest, it.progetto.reviewmanager.grpc.GetReviewsResponse> getGetReviewsMethod;
    if ((getGetReviewsMethod = ReviewManagerServiceGrpc.getGetReviewsMethod) == null) {
      synchronized (ReviewManagerServiceGrpc.class) {
        if ((getGetReviewsMethod = ReviewManagerServiceGrpc.getGetReviewsMethod) == null) {
          ReviewManagerServiceGrpc.getGetReviewsMethod = getGetReviewsMethod =
              io.grpc.MethodDescriptor.<it.progetto.reviewmanager.grpc.GetReviewsRequest, it.progetto.reviewmanager.grpc.GetReviewsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "getReviews"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.reviewmanager.grpc.GetReviewsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  it.progetto.reviewmanager.grpc.GetReviewsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ReviewManagerServiceMethodDescriptorSupplier("getReviews"))
              .build();
        }
      }
    }
    return getGetReviewsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ReviewManagerServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ReviewManagerServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ReviewManagerServiceStub>() {
        @java.lang.Override
        public ReviewManagerServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ReviewManagerServiceStub(channel, callOptions);
        }
      };
    return ReviewManagerServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ReviewManagerServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ReviewManagerServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ReviewManagerServiceBlockingStub>() {
        @java.lang.Override
        public ReviewManagerServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ReviewManagerServiceBlockingStub(channel, callOptions);
        }
      };
    return ReviewManagerServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ReviewManagerServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ReviewManagerServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ReviewManagerServiceFutureStub>() {
        @java.lang.Override
        public ReviewManagerServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ReviewManagerServiceFutureStub(channel, callOptions);
        }
      };
    return ReviewManagerServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class ReviewManagerServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void addReview(it.progetto.reviewmanager.grpc.AddReviewRequest request,
        io.grpc.stub.StreamObserver<it.progetto.reviewmanager.grpc.AddReviewResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAddReviewMethod(), responseObserver);
    }

    /**
     */
    public void getReviews(it.progetto.reviewmanager.grpc.GetReviewsRequest request,
        io.grpc.stub.StreamObserver<it.progetto.reviewmanager.grpc.GetReviewsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetReviewsMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getAddReviewMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.reviewmanager.grpc.AddReviewRequest,
                it.progetto.reviewmanager.grpc.AddReviewResponse>(
                  this, METHODID_ADD_REVIEW)))
          .addMethod(
            getGetReviewsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                it.progetto.reviewmanager.grpc.GetReviewsRequest,
                it.progetto.reviewmanager.grpc.GetReviewsResponse>(
                  this, METHODID_GET_REVIEWS)))
          .build();
    }
  }

  /**
   */
  public static final class ReviewManagerServiceStub extends io.grpc.stub.AbstractAsyncStub<ReviewManagerServiceStub> {
    private ReviewManagerServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReviewManagerServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ReviewManagerServiceStub(channel, callOptions);
    }

    /**
     */
    public void addReview(it.progetto.reviewmanager.grpc.AddReviewRequest request,
        io.grpc.stub.StreamObserver<it.progetto.reviewmanager.grpc.AddReviewResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAddReviewMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getReviews(it.progetto.reviewmanager.grpc.GetReviewsRequest request,
        io.grpc.stub.StreamObserver<it.progetto.reviewmanager.grpc.GetReviewsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetReviewsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ReviewManagerServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<ReviewManagerServiceBlockingStub> {
    private ReviewManagerServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReviewManagerServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ReviewManagerServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public it.progetto.reviewmanager.grpc.AddReviewResponse addReview(it.progetto.reviewmanager.grpc.AddReviewRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAddReviewMethod(), getCallOptions(), request);
    }

    /**
     */
    public it.progetto.reviewmanager.grpc.GetReviewsResponse getReviews(it.progetto.reviewmanager.grpc.GetReviewsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetReviewsMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ReviewManagerServiceFutureStub extends io.grpc.stub.AbstractFutureStub<ReviewManagerServiceFutureStub> {
    private ReviewManagerServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReviewManagerServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ReviewManagerServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.reviewmanager.grpc.AddReviewResponse> addReview(
        it.progetto.reviewmanager.grpc.AddReviewRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAddReviewMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<it.progetto.reviewmanager.grpc.GetReviewsResponse> getReviews(
        it.progetto.reviewmanager.grpc.GetReviewsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetReviewsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_ADD_REVIEW = 0;
  private static final int METHODID_GET_REVIEWS = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ReviewManagerServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ReviewManagerServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_ADD_REVIEW:
          serviceImpl.addReview((it.progetto.reviewmanager.grpc.AddReviewRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.reviewmanager.grpc.AddReviewResponse>) responseObserver);
          break;
        case METHODID_GET_REVIEWS:
          serviceImpl.getReviews((it.progetto.reviewmanager.grpc.GetReviewsRequest) request,
              (io.grpc.stub.StreamObserver<it.progetto.reviewmanager.grpc.GetReviewsResponse>) responseObserver);
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

  private static abstract class ReviewManagerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ReviewManagerServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return it.progetto.reviewmanager.grpc.ReviewManager.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ReviewManagerService");
    }
  }

  private static final class ReviewManagerServiceFileDescriptorSupplier
      extends ReviewManagerServiceBaseDescriptorSupplier {
    ReviewManagerServiceFileDescriptorSupplier() {}
  }

  private static final class ReviewManagerServiceMethodDescriptorSupplier
      extends ReviewManagerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ReviewManagerServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (ReviewManagerServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ReviewManagerServiceFileDescriptorSupplier())
              .addMethod(getAddReviewMethod())
              .addMethod(getGetReviewsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
