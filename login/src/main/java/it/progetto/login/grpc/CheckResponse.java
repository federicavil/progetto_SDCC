// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: src/main/resources/loginService.proto

package it.progetto.login.grpc;

/**
 * Protobuf type {@code it.progetto.progetto_sdcc.proto.CheckResponse}
 */
public final class CheckResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:it.progetto.progetto_sdcc.proto.CheckResponse)
    CheckResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CheckResponse.newBuilder() to construct.
  private CheckResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CheckResponse() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new CheckResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return it.progetto.login.grpc.LoginServiceOuterClass.internal_static_it_progetto_progetto_sdcc_proto_CheckResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return it.progetto.login.grpc.LoginServiceOuterClass.internal_static_it_progetto_progetto_sdcc_proto_CheckResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            it.progetto.login.grpc.CheckResponse.class, it.progetto.login.grpc.CheckResponse.Builder.class);
  }

  public static final int ISLOGGED_FIELD_NUMBER = 1;
  private boolean isLogged_;
  /**
   * <code>bool isLogged = 1;</code>
   * @return The isLogged.
   */
  @java.lang.Override
  public boolean getIsLogged() {
    return isLogged_;
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
    if (isLogged_ != false) {
      output.writeBool(1, isLogged_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (isLogged_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(1, isLogged_);
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
    if (!(obj instanceof it.progetto.login.grpc.CheckResponse)) {
      return super.equals(obj);
    }
    it.progetto.login.grpc.CheckResponse other = (it.progetto.login.grpc.CheckResponse) obj;

    if (getIsLogged()
        != other.getIsLogged()) return false;
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
    hash = (37 * hash) + ISLOGGED_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getIsLogged());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static it.progetto.login.grpc.CheckResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static it.progetto.login.grpc.CheckResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static it.progetto.login.grpc.CheckResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static it.progetto.login.grpc.CheckResponse parseFrom(
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
  public static Builder newBuilder(it.progetto.login.grpc.CheckResponse prototype) {
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
   * Protobuf type {@code it.progetto.progetto_sdcc.proto.CheckResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:it.progetto.progetto_sdcc.proto.CheckResponse)
      it.progetto.login.grpc.CheckResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return it.progetto.login.grpc.LoginServiceOuterClass.internal_static_it_progetto_progetto_sdcc_proto_CheckResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return it.progetto.login.grpc.LoginServiceOuterClass.internal_static_it_progetto_progetto_sdcc_proto_CheckResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              it.progetto.login.grpc.CheckResponse.class, it.progetto.login.grpc.CheckResponse.Builder.class);
    }

    // Construct using it.progetto.login.grpc.CheckResponse.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      isLogged_ = false;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return it.progetto.login.grpc.LoginServiceOuterClass.internal_static_it_progetto_progetto_sdcc_proto_CheckResponse_descriptor;
    }

    @java.lang.Override
    public it.progetto.login.grpc.CheckResponse getDefaultInstanceForType() {
      return it.progetto.login.grpc.CheckResponse.getDefaultInstance();
    }

    @java.lang.Override
    public it.progetto.login.grpc.CheckResponse build() {
      it.progetto.login.grpc.CheckResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public it.progetto.login.grpc.CheckResponse buildPartial() {
      it.progetto.login.grpc.CheckResponse result = new it.progetto.login.grpc.CheckResponse(this);
      result.isLogged_ = isLogged_;
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
      if (other instanceof it.progetto.login.grpc.CheckResponse) {
        return mergeFrom((it.progetto.login.grpc.CheckResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(it.progetto.login.grpc.CheckResponse other) {
      if (other == it.progetto.login.grpc.CheckResponse.getDefaultInstance()) return this;
      if (other.getIsLogged() != false) {
        setIsLogged(other.getIsLogged());
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
            case 8: {
              isLogged_ = input.readBool();

              break;
            } // case 8
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

    private boolean isLogged_ ;
    /**
     * <code>bool isLogged = 1;</code>
     * @return The isLogged.
     */
    @java.lang.Override
    public boolean getIsLogged() {
      return isLogged_;
    }
    /**
     * <code>bool isLogged = 1;</code>
     * @param value The isLogged to set.
     * @return This builder for chaining.
     */
    public Builder setIsLogged(boolean value) {
      
      isLogged_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bool isLogged = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearIsLogged() {
      
      isLogged_ = false;
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


    // @@protoc_insertion_point(builder_scope:it.progetto.progetto_sdcc.proto.CheckResponse)
  }

  // @@protoc_insertion_point(class_scope:it.progetto.progetto_sdcc.proto.CheckResponse)
  private static final it.progetto.login.grpc.CheckResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new it.progetto.login.grpc.CheckResponse();
  }

  public static it.progetto.login.grpc.CheckResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CheckResponse>
      PARSER = new com.google.protobuf.AbstractParser<CheckResponse>() {
    @java.lang.Override
    public CheckResponse parsePartialFrom(
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

  public static com.google.protobuf.Parser<CheckResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CheckResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public it.progetto.login.grpc.CheckResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

