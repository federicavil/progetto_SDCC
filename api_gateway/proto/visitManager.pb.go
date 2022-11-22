// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: visitManager.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Difficulty int32

const (
	Difficulty_T   Difficulty = 0
	Difficulty_E   Difficulty = 1
	Difficulty_EE  Difficulty = 2
	Difficulty_EEA Difficulty = 3
)

// Enum value maps for Difficulty.
var (
	Difficulty_name = map[int32]string{
		0: "T",
		1: "E",
		2: "EE",
		3: "EEA",
	}
	Difficulty_value = map[string]int32{
		"T":   0,
		"E":   1,
		"EE":  2,
		"EEA": 3,
	}
)

func (x Difficulty) Enum() *Difficulty {
	p := new(Difficulty)
	*p = x
	return p
}

func (x Difficulty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Difficulty) Descriptor() protoreflect.EnumDescriptor {
	return file_visitManager_proto_enumTypes[0].Descriptor()
}

func (Difficulty) Type() protoreflect.EnumType {
	return &file_visitManager_proto_enumTypes[0]
}

func (x Difficulty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Difficulty) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Difficulty(num)
	return nil
}

// Deprecated: Use Difficulty.Descriptor instead.
func (Difficulty) EnumDescriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{0}
}

type Invite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID_Visit *string `protobuf:"bytes,1,req,name=ID_Visit,json=IDVisit" json:"ID_Visit,omitempty"`
	Username *string `protobuf:"bytes,2,req,name=Username" json:"Username,omitempty"`
}

func (x *Invite) Reset() {
	*x = Invite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invite) ProtoMessage() {}

func (x *Invite) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invite.ProtoReflect.Descriptor instead.
func (*Invite) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{0}
}

func (x *Invite) GetID_Visit() string {
	if x != nil && x.ID_Visit != nil {
		return *x.ID_Visit
	}
	return ""
}

func (x *Invite) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

type InviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID_Visit *string `protobuf:"bytes,1,req,name=ID_Visit,json=IDVisit" json:"ID_Visit,omitempty"`
	Username *string `protobuf:"bytes,2,req,name=Username" json:"Username,omitempty"`
	Response *string `protobuf:"bytes,3,req,name=Response" json:"Response,omitempty"`
}

func (x *InviteResponse) Reset() {
	*x = InviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteResponse) ProtoMessage() {}

func (x *InviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteResponse.ProtoReflect.Descriptor instead.
func (*InviteResponse) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{1}
}

func (x *InviteResponse) GetID_Visit() string {
	if x != nil && x.ID_Visit != nil {
		return *x.ID_Visit
	}
	return ""
}

func (x *InviteResponse) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *InviteResponse) GetResponse() string {
	if x != nil && x.Response != nil {
		return *x.Response
	}
	return ""
}

type Visits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visit []*Visit `protobuf:"bytes,1,rep,name=Visit" json:"Visit,omitempty"`
}

func (x *Visits) Reset() {
	*x = Visits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Visits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Visits) ProtoMessage() {}

func (x *Visits) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Visits.ProtoReflect.Descriptor instead.
func (*Visits) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{2}
}

func (x *Visits) GetVisit() []*Visit {
	if x != nil {
		return x.Visit
	}
	return nil
}

type ID_Visit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *string `protobuf:"bytes,1,req,name=Value" json:"Value,omitempty"`
}

func (x *ID_Visit) Reset() {
	*x = ID_Visit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID_Visit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID_Visit) ProtoMessage() {}

func (x *ID_Visit) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID_Visit.ProtoReflect.Descriptor instead.
func (*ID_Visit) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{3}
}

func (x *ID_Visit) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type Visit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID_Visit     *uint32  `protobuf:"varint,1,req,name=ID_Visit,json=IDVisit" json:"ID_Visit,omitempty"`
	ID_Path      *string  `protobuf:"bytes,2,req,name=ID_Path,json=IDPath" json:"ID_Path,omitempty"`
	Creator      *string  `protobuf:"bytes,3,req,name=Creator" json:"Creator,omitempty"`
	Timestamp    *string  `protobuf:"bytes,4,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Participants []string `protobuf:"bytes,5,rep,name=Participants" json:"Participants,omitempty"`
}

func (x *Visit) Reset() {
	*x = Visit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Visit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Visit) ProtoMessage() {}

func (x *Visit) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Visit.ProtoReflect.Descriptor instead.
func (*Visit) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{4}
}

func (x *Visit) GetID_Visit() uint32 {
	if x != nil && x.ID_Visit != nil {
		return *x.ID_Visit
	}
	return 0
}

func (x *Visit) GetID_Path() string {
	if x != nil && x.ID_Path != nil {
		return *x.ID_Path
	}
	return ""
}

func (x *Visit) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *Visit) GetTimestamp() string {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return ""
}

func (x *Visit) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

type InputVisit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  *string `protobuf:"bytes,1,req,name=Username" json:"Username,omitempty"`
	Pathname  *string `protobuf:"bytes,2,req,name=Pathname" json:"Pathname,omitempty"`
	Timestamp *string `protobuf:"bytes,3,req,name=Timestamp" json:"Timestamp,omitempty"`
}

func (x *InputVisit) Reset() {
	*x = InputVisit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputVisit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputVisit) ProtoMessage() {}

func (x *InputVisit) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputVisit.ProtoReflect.Descriptor instead.
func (*InputVisit) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{5}
}

func (x *InputVisit) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *InputVisit) GetPathname() string {
	if x != nil && x.Pathname != nil {
		return *x.Pathname
	}
	return ""
}

func (x *InputVisit) GetTimestamp() string {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return ""
}

type Return struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32 `protobuf:"varint,1,req,name=Ret" json:"Ret,omitempty"`
}

func (x *Return) Reset() {
	*x = Return{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Return) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Return) ProtoMessage() {}

func (x *Return) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Return.ProtoReflect.Descriptor instead.
func (*Return) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{6}
}

func (x *Return) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID *string `protobuf:"bytes,1,req,name=ID" json:"ID,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_visitManager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_visitManager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_visitManager_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

var File_visitManager_proto protoreflect.FileDescriptor

var file_visitManager_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x69, 0x73, 0x69, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x49, 0x44, 0x5f, 0x56, 0x69, 0x73, 0x69, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x07, 0x49, 0x44, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x49, 0x44, 0x5f, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x49, 0x44, 0x56, 0x69, 0x73,
	0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x06, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x05, 0x56, 0x69, 0x73, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x05, 0x56, 0x69, 0x73,
	0x69, 0x74, 0x22, 0x20, 0x0a, 0x08, 0x49, 0x44, 0x5f, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x49, 0x44, 0x5f, 0x56, 0x69, 0x73, 0x69, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x07, 0x49, 0x44, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x49, 0x44, 0x5f,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x49, 0x44, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x62,
	0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x74, 0x68,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x1a, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x52, 0x65, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x52, 0x65, 0x74, 0x22, 0x16,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x2a, 0x2b, 0x0a, 0x0a, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x12, 0x05, 0x0a, 0x01, 0x54, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x45,
	0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x45, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x45,
	0x41, 0x10, 0x03, 0x32, 0xd8, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x56, 0x69, 0x73,
	0x69, 0x74, 0x12, 0x0b, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74, 0x1a,
	0x07, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x56, 0x69, 0x73, 0x69, 0x74, 0x73, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x07, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x73, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x09, 0x2e, 0x49,
	0x44, 0x5f, 0x56, 0x69, 0x73, 0x69, 0x74, 0x1a, 0x06, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x22,
	0x00, 0x12, 0x27, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x07, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x1a,
	0x07, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x14, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x52, 0x65, 0x66, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f,
}

var (
	file_visitManager_proto_rawDescOnce sync.Once
	file_visitManager_proto_rawDescData = file_visitManager_proto_rawDesc
)

func file_visitManager_proto_rawDescGZIP() []byte {
	file_visitManager_proto_rawDescOnce.Do(func() {
		file_visitManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_visitManager_proto_rawDescData)
	})
	return file_visitManager_proto_rawDescData
}

var file_visitManager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_visitManager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_visitManager_proto_goTypes = []interface{}{
	(Difficulty)(0),        // 0: Difficulty
	(*Invite)(nil),         // 1: Invite
	(*InviteResponse)(nil), // 2: InviteResponse
	(*Visits)(nil),         // 3: Visits
	(*ID_Visit)(nil),       // 4: ID_Visit
	(*Visit)(nil),          // 5: Visit
	(*InputVisit)(nil),     // 6: InputVisit
	(*Return)(nil),         // 7: Return
	(*User)(nil),           // 8: User
}
var file_visitManager_proto_depIdxs = []int32{
	5, // 0: Visits.Visit:type_name -> Visit
	6, // 1: ManageVisit.AddNewVisit:input_type -> InputVisit
	8, // 2: ManageVisit.GetAllVisits:input_type -> User
	4, // 3: ManageVisit.GetVisitByID:input_type -> ID_Visit
	1, // 4: ManageVisit.InviteUserToVisit:input_type -> Invite
	2, // 5: ManageVisit.AcceptOrRefuseInvite:input_type -> InviteResponse
	7, // 6: ManageVisit.AddNewVisit:output_type -> Return
	3, // 7: ManageVisit.GetAllVisits:output_type -> Visits
	5, // 8: ManageVisit.GetVisitByID:output_type -> Visit
	7, // 9: ManageVisit.InviteUserToVisit:output_type -> Return
	7, // 10: ManageVisit.AcceptOrRefuseInvite:output_type -> Return
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_visitManager_proto_init() }
func file_visitManager_proto_init() {
	if File_visitManager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_visitManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Visits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID_Visit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Visit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputVisit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Return); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_visitManager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_visitManager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_visitManager_proto_goTypes,
		DependencyIndexes: file_visitManager_proto_depIdxs,
		EnumInfos:         file_visitManager_proto_enumTypes,
		MessageInfos:      file_visitManager_proto_msgTypes,
	}.Build()
	File_visitManager_proto = out.File
	file_visitManager_proto_rawDesc = nil
	file_visitManager_proto_goTypes = nil
	file_visitManager_proto_depIdxs = nil
}
