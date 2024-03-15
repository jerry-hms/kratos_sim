// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/sim.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar    string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Mobile    string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[0]
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
	return file_v1_sim_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Mobile   string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RegisterReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *User `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReply) GetUserInfo() *User {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *User  `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReply) GetUserInfo() *User {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecvId  int64  `protobuf:"varint,1,opt,name=recv_id,json=recvId,proto3" json:"recv_id,omitempty"`
	Scene   string `protobuf:"bytes,2,opt,name=scene,proto3" json:"scene,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Url     string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	MsgType string `protobuf:"bytes,5,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
}

func (x *ChatReq) Reset() {
	*x = ChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReq) ProtoMessage() {}

func (x *ChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReq.ProtoReflect.Descriptor instead.
func (*ChatReq) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{5}
}

func (x *ChatReq) GetRecvId() int64 {
	if x != nil {
		return x.RecvId
	}
	return 0
}

func (x *ChatReq) GetScene() string {
	if x != nil {
		return x.Scene
	}
	return ""
}

func (x *ChatReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ChatReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ChatReq) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

type ChatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatReply) Reset() {
	*x = ChatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReply) ProtoMessage() {}

func (x *ChatReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReply.ProtoReflect.Descriptor instead.
func (*ChatReply) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{6}
}

type BindWsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *BindWsReq) Reset() {
	*x = BindWsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindWsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindWsReq) ProtoMessage() {}

func (x *BindWsReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindWsReq.ProtoReflect.Descriptor instead.
func (*BindWsReq) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{7}
}

func (x *BindWsReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BindWsReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type BindWsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindWsReply) Reset() {
	*x = BindWsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sim_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindWsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindWsReply) ProtoMessage() {}

func (x *BindWsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sim_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindWsReply.ProtoReflect.Descriptor instead.
func (*BindWsReply) Descriptor() ([]byte, []int) {
	return file_v1_sim_proto_rawDescGZIP(), []int{8}
}

var File_v1_sim_proto protoreflect.FileDescriptor

var file_v1_sim_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x22, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x51, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x41, 0x0a, 0x09, 0x42, 0x69, 0x6e, 0x64, 0x57, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x57, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0xd1, 0x02, 0x0a, 0x03, 0x53, 0x69, 0x6d, 0x12, 0x56, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x4a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x3a, 0x01, 0x2a, 0x22, 0x08, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x46, 0x0a,
	0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x12, 0x5e, 0x0a, 0x06, 0x42, 0x69, 0x6e, 0x64, 0x57, 0x73, 0x12,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e,
	0x64, 0x57, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x57, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69,
	0x6e, 0x64, 0x2d, 0x77, 0x73, 0x42, 0x3b, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x27, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x5f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_sim_proto_rawDescOnce sync.Once
	file_v1_sim_proto_rawDescData = file_v1_sim_proto_rawDesc
)

func file_v1_sim_proto_rawDescGZIP() []byte {
	file_v1_sim_proto_rawDescOnce.Do(func() {
		file_v1_sim_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_sim_proto_rawDescData)
	})
	return file_v1_sim_proto_rawDescData
}

var file_v1_sim_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_sim_proto_goTypes = []interface{}{
	(*User)(nil),          // 0: api.sim.v1.User
	(*RegisterReq)(nil),   // 1: api.sim.v1.RegisterReq
	(*RegisterReply)(nil), // 2: api.sim.v1.RegisterReply
	(*LoginReq)(nil),      // 3: api.sim.v1.LoginReq
	(*LoginReply)(nil),    // 4: api.sim.v1.LoginReply
	(*ChatReq)(nil),       // 5: api.sim.v1.ChatReq
	(*ChatReply)(nil),     // 6: api.sim.v1.ChatReply
	(*BindWsReq)(nil),     // 7: api.sim.v1.BindWsReq
	(*BindWsReply)(nil),   // 8: api.sim.v1.BindWsReply
}
var file_v1_sim_proto_depIdxs = []int32{
	0, // 0: api.sim.v1.RegisterReply.user_info:type_name -> api.sim.v1.User
	0, // 1: api.sim.v1.LoginReply.user_info:type_name -> api.sim.v1.User
	1, // 2: api.sim.v1.Sim.Register:input_type -> api.sim.v1.RegisterReq
	3, // 3: api.sim.v1.Sim.Login:input_type -> api.sim.v1.LoginReq
	5, // 4: api.sim.v1.Sim.Chat:input_type -> api.sim.v1.ChatReq
	7, // 5: api.sim.v1.Sim.BindWs:input_type -> api.sim.v1.BindWsReq
	2, // 6: api.sim.v1.Sim.Register:output_type -> api.sim.v1.RegisterReply
	4, // 7: api.sim.v1.Sim.Login:output_type -> api.sim.v1.LoginReply
	6, // 8: api.sim.v1.Sim.Chat:output_type -> api.sim.v1.ChatReply
	8, // 9: api.sim.v1.Sim.BindWs:output_type -> api.sim.v1.BindWsReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_sim_proto_init() }
func file_v1_sim_proto_init() {
	if File_v1_sim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_sim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_sim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_v1_sim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_v1_sim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_v1_sim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_v1_sim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatReq); i {
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
		file_v1_sim_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatReply); i {
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
		file_v1_sim_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindWsReq); i {
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
		file_v1_sim_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindWsReply); i {
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
			RawDescriptor: file_v1_sim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_sim_proto_goTypes,
		DependencyIndexes: file_v1_sim_proto_depIdxs,
		MessageInfos:      file_v1_sim_proto_msgTypes,
	}.Build()
	File_v1_sim_proto = out.File
	file_v1_sim_proto_rawDesc = nil
	file_v1_sim_proto_goTypes = nil
	file_v1_sim_proto_depIdxs = nil
}
