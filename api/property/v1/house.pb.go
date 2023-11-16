// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: property/v1/house.proto

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

type House struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price    string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	CoverUrl string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
}

func (x *House) Reset() {
	*x = House{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *House) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*House) ProtoMessage() {}

func (x *House) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use House.ProtoReflect.Descriptor instead.
func (*House) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{0}
}

func (x *House) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *House) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *House) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *House) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

type CreateHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHouseRequest) Reset() {
	*x = CreateHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHouseRequest) ProtoMessage() {}

func (x *CreateHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHouseRequest.ProtoReflect.Descriptor instead.
func (*CreateHouseRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{1}
}

type CreateHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHouseReply) Reset() {
	*x = CreateHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHouseReply) ProtoMessage() {}

func (x *CreateHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHouseReply.ProtoReflect.Descriptor instead.
func (*CreateHouseReply) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{2}
}

type UpdateHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateHouseRequest) Reset() {
	*x = UpdateHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHouseRequest) ProtoMessage() {}

func (x *UpdateHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHouseRequest.ProtoReflect.Descriptor instead.
func (*UpdateHouseRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{3}
}

type UpdateHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateHouseReply) Reset() {
	*x = UpdateHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHouseReply) ProtoMessage() {}

func (x *UpdateHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHouseReply.ProtoReflect.Descriptor instead.
func (*UpdateHouseReply) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{4}
}

type DeleteHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHouseRequest) Reset() {
	*x = DeleteHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHouseRequest) ProtoMessage() {}

func (x *DeleteHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHouseRequest.ProtoReflect.Descriptor instead.
func (*DeleteHouseRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{5}
}

type DeleteHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHouseReply) Reset() {
	*x = DeleteHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHouseReply) ProtoMessage() {}

func (x *DeleteHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHouseReply.ProtoReflect.Descriptor instead.
func (*DeleteHouseReply) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{6}
}

type GetHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHouseRequest) Reset() {
	*x = GetHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHouseRequest) ProtoMessage() {}

func (x *GetHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHouseRequest.ProtoReflect.Descriptor instead.
func (*GetHouseRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{7}
}

type GetHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHouseReply) Reset() {
	*x = GetHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHouseReply) ProtoMessage() {}

func (x *GetHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHouseReply.ProtoReflect.Descriptor instead.
func (*GetHouseReply) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{8}
}

type PagintatedListHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PagintatedListHouseRequest) Reset() {
	*x = PagintatedListHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagintatedListHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagintatedListHouseRequest) ProtoMessage() {}

func (x *PagintatedListHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagintatedListHouseRequest.ProtoReflect.Descriptor instead.
func (*PagintatedListHouseRequest) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{9}
}

func (x *PagintatedListHouseRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PagintatedListHouseRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PagintatedListHouseRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PaginatedListListHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp string   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      []*House `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PaginatedListListHouseReply) Reset() {
	*x = PaginatedListListHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_v1_house_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginatedListListHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginatedListListHouseReply) ProtoMessage() {}

func (x *PaginatedListListHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_property_v1_house_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginatedListListHouseReply.ProtoReflect.Descriptor instead.
func (*PaginatedListListHouseReply) Descriptor() ([]byte, []int) {
	return file_property_v1_house_proto_rawDescGZIP(), []int{10}
}

func (x *PaginatedListListHouseReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PaginatedListListHouseReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PaginatedListListHouseReply) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *PaginatedListListHouseReply) GetData() []*House {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_property_v1_house_proto protoreflect.FileDescriptor

var file_property_v1_house_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x60, 0x0a, 0x05, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x14, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x61,
	0x0a, 0x1a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x74, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x8d, 0x01, 0x0a, 0x1b, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xce, 0x03, 0x0a, 0x08, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x76, 0x63, 0x12, 0x55,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x6f, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x2d, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x68, 0x6b, 0x35, 0x39, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_property_v1_house_proto_rawDescOnce sync.Once
	file_property_v1_house_proto_rawDescData = file_property_v1_house_proto_rawDesc
)

func file_property_v1_house_proto_rawDescGZIP() []byte {
	file_property_v1_house_proto_rawDescOnce.Do(func() {
		file_property_v1_house_proto_rawDescData = protoimpl.X.CompressGZIP(file_property_v1_house_proto_rawDescData)
	})
	return file_property_v1_house_proto_rawDescData
}

var file_property_v1_house_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_property_v1_house_proto_goTypes = []interface{}{
	(*House)(nil),                       // 0: api.property.v1.House
	(*CreateHouseRequest)(nil),          // 1: api.property.v1.CreateHouseRequest
	(*CreateHouseReply)(nil),            // 2: api.property.v1.CreateHouseReply
	(*UpdateHouseRequest)(nil),          // 3: api.property.v1.UpdateHouseRequest
	(*UpdateHouseReply)(nil),            // 4: api.property.v1.UpdateHouseReply
	(*DeleteHouseRequest)(nil),          // 5: api.property.v1.DeleteHouseRequest
	(*DeleteHouseReply)(nil),            // 6: api.property.v1.DeleteHouseReply
	(*GetHouseRequest)(nil),             // 7: api.property.v1.GetHouseRequest
	(*GetHouseReply)(nil),               // 8: api.property.v1.GetHouseReply
	(*PagintatedListHouseRequest)(nil),  // 9: api.property.v1.PagintatedListHouseRequest
	(*PaginatedListListHouseReply)(nil), // 10: api.property.v1.PaginatedListListHouseReply
}
var file_property_v1_house_proto_depIdxs = []int32{
	0,  // 0: api.property.v1.PaginatedListListHouseReply.data:type_name -> api.property.v1.House
	1,  // 1: api.property.v1.HouseSvc.CreateHouse:input_type -> api.property.v1.CreateHouseRequest
	3,  // 2: api.property.v1.HouseSvc.UpdateHouse:input_type -> api.property.v1.UpdateHouseRequest
	5,  // 3: api.property.v1.HouseSvc.DeleteHouse:input_type -> api.property.v1.DeleteHouseRequest
	7,  // 4: api.property.v1.HouseSvc.GetHouse:input_type -> api.property.v1.GetHouseRequest
	9,  // 5: api.property.v1.HouseSvc.PaginatedListHouse:input_type -> api.property.v1.PagintatedListHouseRequest
	2,  // 6: api.property.v1.HouseSvc.CreateHouse:output_type -> api.property.v1.CreateHouseReply
	4,  // 7: api.property.v1.HouseSvc.UpdateHouse:output_type -> api.property.v1.UpdateHouseReply
	6,  // 8: api.property.v1.HouseSvc.DeleteHouse:output_type -> api.property.v1.DeleteHouseReply
	8,  // 9: api.property.v1.HouseSvc.GetHouse:output_type -> api.property.v1.GetHouseReply
	10, // 10: api.property.v1.HouseSvc.PaginatedListHouse:output_type -> api.property.v1.PaginatedListListHouseReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_property_v1_house_proto_init() }
func file_property_v1_house_proto_init() {
	if File_property_v1_house_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_property_v1_house_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*House); i {
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
		file_property_v1_house_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHouseRequest); i {
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
		file_property_v1_house_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHouseReply); i {
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
		file_property_v1_house_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHouseRequest); i {
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
		file_property_v1_house_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHouseReply); i {
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
		file_property_v1_house_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHouseRequest); i {
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
		file_property_v1_house_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHouseReply); i {
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
		file_property_v1_house_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHouseRequest); i {
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
		file_property_v1_house_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHouseReply); i {
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
		file_property_v1_house_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagintatedListHouseRequest); i {
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
		file_property_v1_house_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginatedListListHouseReply); i {
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
			RawDescriptor: file_property_v1_house_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_property_v1_house_proto_goTypes,
		DependencyIndexes: file_property_v1_house_proto_depIdxs,
		MessageInfos:      file_property_v1_house_proto_msgTypes,
	}.Build()
	File_property_v1_house_proto = out.File
	file_property_v1_house_proto_rawDesc = nil
	file_property_v1_house_proto_goTypes = nil
	file_property_v1_house_proto_depIdxs = nil
}
