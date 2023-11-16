// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: property/v1/user.proto

package v1

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[0]
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
	return file_property_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mobile string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      *User  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code      int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_property_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserReply) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetUserReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUserReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetUserReply) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListUserRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	CurrentPage int64 `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	PerPage     int64 `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	HasNext     int64 `protobuf:"varint,4,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_property_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *Meta) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Meta) GetCurrentPage() int64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Meta) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *Meta) GetHasNext() int64 {
	if x != nil {
		return x.HasNext
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*User `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Meta *Meta   `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_property_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *Data) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Data) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      *Data  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_property_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *ListUserReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListUserReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListUserReply) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *ListUserReply) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_property_v1_user_proto protoreflect.FileDescriptor

var file_property_v1_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x53, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0x7d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x42, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x75, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x22, 0x5c, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x7e, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x86, 0x02, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x76, 0x63, 0x12, 0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x2d, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x68, 0x6b, 0x35, 0x39, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_property_v1_user_proto_rawDescOnce sync.Once
	file_property_v1_user_proto_rawDescData = file_property_v1_user_proto_rawDesc
)

func file_property_v1_user_proto_rawDescGZIP() []byte {
	file_property_v1_user_proto_rawDescOnce.Do(func() {
		file_property_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_property_v1_user_proto_rawDescData)
	})
	return file_property_v1_user_proto_rawDescData
}

var file_property_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_property_v1_user_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: api.property.v1.User
	(*GetUserRequest)(nil),  // 1: api.property.v1.GetUserRequest
	(*GetUserReply)(nil),    // 2: api.property.v1.GetUserReply
	(*ListUserRequest)(nil), // 3: api.property.v1.ListUserRequest
	(*Meta)(nil),            // 4: api.property.v1.Meta
	(*Data)(nil),            // 5: api.property.v1.Data
	(*ListUserReply)(nil),   // 6: api.property.v1.ListUserReply
}
var file_property_v1_user_proto_depIdxs = []int32{
	0, // 0: api.property.v1.GetUserReply.data:type_name -> api.property.v1.User
	0, // 1: api.property.v1.Data.data:type_name -> api.property.v1.User
	4, // 2: api.property.v1.Data.meta:type_name -> api.property.v1.Meta
	5, // 3: api.property.v1.ListUserReply.data:type_name -> api.property.v1.Data
	1, // 4: api.property.v1.UserSvc.GetUserByUserId:input_type -> api.property.v1.GetUserRequest
	1, // 5: api.property.v1.UserSvc.GetUserByMobile:input_type -> api.property.v1.GetUserRequest
	3, // 6: api.property.v1.UserSvc.GetPaginatedUsers:input_type -> api.property.v1.ListUserRequest
	2, // 7: api.property.v1.UserSvc.GetUserByUserId:output_type -> api.property.v1.GetUserReply
	2, // 8: api.property.v1.UserSvc.GetUserByMobile:output_type -> api.property.v1.GetUserReply
	6, // 9: api.property.v1.UserSvc.GetPaginatedUsers:output_type -> api.property.v1.ListUserReply
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_property_v1_user_proto_init() }
func file_property_v1_user_proto_init() {
	if File_property_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_property_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_property_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_property_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_property_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_property_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_property_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_property_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply); i {
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
			RawDescriptor: file_property_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_property_v1_user_proto_goTypes,
		DependencyIndexes: file_property_v1_user_proto_depIdxs,
		MessageInfos:      file_property_v1_user_proto_msgTypes,
	}.Build()
	File_property_v1_user_proto = out.File
	file_property_v1_user_proto_rawDesc = nil
	file_property_v1_user_proto_goTypes = nil
	file_property_v1_user_proto_depIdxs = nil
}
