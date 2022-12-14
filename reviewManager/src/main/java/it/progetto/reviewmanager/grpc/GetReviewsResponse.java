// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: src/main/resources/reviewManager.proto

package it.progetto.reviewmanager.grpc;

/**
 * Protobuf type {@code it.progetto.progetto_sdcc.proto.GetReviewsResponse}
 */
public final class GetReviewsResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:it.progetto.progetto_sdcc.proto.GetReviewsResponse)
    GetReviewsResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetReviewsResponse.newBuilder() to construct.
  private GetReviewsResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetReviewsResponse() {
    reviewsList_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new GetReviewsResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return it.progetto.reviewmanager.grpc.ReviewManager.internal_static_it_progetto_progetto_sdcc_proto_GetReviewsResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return it.progetto.reviewmanager.grpc.ReviewManager.internal_static_it_progetto_progetto_sdcc_proto_GetReviewsResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            it.progetto.reviewmanager.grpc.GetReviewsResponse.class, it.progetto.reviewmanager.grpc.GetReviewsResponse.Builder.class);
  }

  public static final int REVIEWSLIST_FIELD_NUMBER = 1;
  private volatile java.lang.Object reviewsList_;
  /**
   * <code>string reviewsList = 1;</code>
   * @return The reviewsList.
   */
  @java.lang.Override
  public java.lang.String getReviewsList() {
    java.lang.Object ref = reviewsList_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      reviewsList_ = s;
      return s;
    }
  }
  /**
   * <code>string reviewsList = 1;</code>
   * @return The bytes for reviewsList.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getReviewsListBytes() {
    java.lang.Object ref = reviewsList_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      reviewsList_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(reviewsList_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, reviewsList_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(reviewsList_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, reviewsList_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof it.progetto.reviewmanager.grpc.GetReviewsResponse)) {
      return super.equals(obj);
    }
    it.progetto.reviewmanager.grpc.GetReviewsResponse other = (it.progetto.reviewmanager.grpc.GetReviewsResponse) obj;

    if (!getReviewsList()
        .equals(other.getReviewsList())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + REVIEWSLIST_FIELD_NUMBER;
    hash = (53 * hash) + getReviewsList().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static it.progetto.reviewmanager.grpc.GetReviewsResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(it.progetto.reviewmanager.grpc.GetReviewsResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code it.progetto.progetto_sdcc.proto.GetReviewsResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:it.progetto.progetto_sdcc.proto.GetReviewsResponse)
      it.progetto.reviewmanager.grpc.GetReviewsResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return it.progetto.reviewmanager.grpc.ReviewManager.internal_static_it_progetto_progetto_sdcc_proto_GetReviewsResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return it.progetto.reviewmanager.grpc.ReviewManager.internal_static_it_progetto_progetto_sdcc_proto_GetReviewsResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              it.progetto.reviewmanager.grpc.GetReviewsResponse.class, it.progetto.reviewmanager.grpc.GetReviewsResponse.Builder.class);
    }

    // Construct using it.progetto.reviewmanager.grpc.GetReviewsResponse.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      reviewsList_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return it.progetto.reviewmanager.grpc.ReviewManager.internal_static_it_progetto_progetto_sdcc_proto_GetReviewsResponse_descriptor;
    }

    @java.lang.Override
    public it.progetto.reviewmanager.grpc.GetReviewsResponse getDefaultInstanceForType() {
      return it.progetto.reviewmanager.grpc.GetReviewsResponse.getDefaultInstance();
    }

    @java.lang.Override
    public it.progetto.reviewmanager.grpc.GetReviewsResponse build() {
      it.progetto.reviewmanager.grpc.GetReviewsResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public it.progetto.reviewmanager.grpc.GetReviewsResponse buildPartial() {
      it.progetto.reviewmanager.grpc.GetReviewsResponse result = new it.progetto.reviewmanager.grpc.GetReviewsResponse(this);
      result.reviewsList_ = reviewsList_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof it.progetto.reviewmanager.grpc.GetReviewsResponse) {
        return mergeFrom((it.progetto.reviewmanager.grpc.GetReviewsResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(it.progetto.reviewmanager.grpc.GetReviewsResponse other) {
      if (other == it.progetto.reviewmanager.grpc.GetReviewsResponse.getDefaultInstance()) return this;
      if (!other.getReviewsList().isEmpty()) {
        reviewsList_ = other.reviewsList_;
        onChanged();
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              reviewsList_ = input.readStringRequireUtf8();

              break;
            } // case 10
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private java.lang.Object reviewsList_ = "";
    /**
     * <code>string reviewsList = 1;</code>
     * @return The reviewsList.
     */
    public java.lang.String getReviewsList() {
      java.lang.Object ref = reviewsList_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        reviewsList_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string reviewsList = 1;</code>
     * @return The bytes for reviewsList.
     */
    public com.google.protobuf.ByteString
        getReviewsListBytes() {
      java.lang.Object ref = reviewsList_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        reviewsList_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string reviewsList = 1;</code>
     * @param value The reviewsList to set.
     * @return This builder for chaining.
     */
    public Builder setReviewsList(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      reviewsList_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string reviewsList = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearReviewsList() {
      
      reviewsList_ = getDefaultInstance().getReviewsList();
      onChanged();
      return this;
    }
    /**
     * <code>string reviewsList = 1;</code>
     * @param value The bytes for reviewsList to set.
     * @return This builder for chaining.
     */
    public Builder setReviewsListBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      reviewsList_ = value;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:it.progetto.progetto_sdcc.proto.GetReviewsResponse)
  }

  // @@protoc_insertion_point(class_scope:it.progetto.progetto_sdcc.proto.GetReviewsResponse)
  private static final it.progetto.reviewmanager.grpc.GetReviewsResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new it.progetto.reviewmanager.grpc.GetReviewsResponse();
  }

  public static it.progetto.reviewmanager.grpc.GetReviewsResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetReviewsResponse>
      PARSER = new com.google.protobuf.AbstractParser<GetReviewsResponse>() {
    @java.lang.Override
    public GetReviewsResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<GetReviewsResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetReviewsResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public it.progetto.reviewmanager.grpc.GetReviewsResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

