// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/product.proto

package genproto

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

type ProductCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float32  `protobuf:"fixed32,3,opt,name=Price,proto3" json:"Price,omitempty"`
	PictureUrls []string `protobuf:"bytes,4,rep,name=PictureUrls,proto3" json:"PictureUrls,omitempty"`
}

func (x *ProductCreate) Reset() {
	*x = ProductCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreate) ProtoMessage() {}

func (x *ProductCreate) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreate.ProtoReflect.Descriptor instead.
func (*ProductCreate) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductCreate) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductCreate) GetPictureUrls() []string {
	if x != nil {
		return x.PictureUrls
	}
	return nil
}

type ProductUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ProductUpt) Reset() {
	*x = ProductUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpt) ProtoMessage() {}

func (x *ProductUpt) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpt.ProtoReflect.Descriptor instead.
func (*ProductUpt) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductUpt) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductUpt) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductUpt) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Body *ProductUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *ProductUpdate) Reset() {
	*x = ProductUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdate) ProtoMessage() {}

func (x *ProductUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdate.ProtoReflect.Descriptor instead.
func (*ProductUpdate) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductUpdate) GetBody() *ProductUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type ProductPicture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId  string `protobuf:"bytes,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	PictureUrl string `protobuf:"bytes,2,opt,name=PictureUrl,proto3" json:"PictureUrl,omitempty"`
}

func (x *ProductPicture) Reset() {
	*x = ProductPicture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPicture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPicture) ProtoMessage() {}

func (x *ProductPicture) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPicture.ProtoReflect.Descriptor instead.
func (*ProductPicture) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductPicture) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductPicture) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

type ProductGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float32  `protobuf:"fixed32,4,opt,name=Price,proto3" json:"Price,omitempty"`
	PictureUrls []string `protobuf:"bytes,5,rep,name=PictureUrls,proto3" json:"PictureUrls,omitempty"`
	CreatedAt   string   `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *ProductGet) Reset() {
	*x = ProductGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductGet) ProtoMessage() {}

func (x *ProductGet) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductGet.ProtoReflect.Descriptor instead.
func (*ProductGet) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductGet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductGet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductGet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductGet) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductGet) GetPictureUrls() []string {
	if x != nil {
		return x.PictureUrls
	}
	return nil
}

func (x *ProductGet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ProductFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string      `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	PriceFrom  float32     `protobuf:"fixed32,2,opt,name=PriceFrom,proto3" json:"PriceFrom,omitempty"`
	PriceTo    float32     `protobuf:"fixed32,3,opt,name=PriceTo,proto3" json:"PriceTo,omitempty"`
	Pagination *Pagination `protobuf:"bytes,4,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *ProductFilter) Reset() {
	*x = ProductFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductFilter) ProtoMessage() {}

func (x *ProductFilter) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductFilter.ProtoReflect.Descriptor instead.
func (*ProductFilter) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{5}
}

func (x *ProductFilter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductFilter) GetPriceFrom() float32 {
	if x != nil {
		return x.PriceFrom
	}
	return 0
}

func (x *ProductFilter) GetPriceTo() float32 {
	if x != nil {
		return x.PriceTo
	}
	return 0
}

func (x *ProductFilter) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products   []*ProductGet `protobuf:"bytes,1,rep,name=Products,proto3" json:"Products,omitempty"`
	TotalCount int64         `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	Pagination *Pagination   `protobuf:"bytes,3,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductList) GetProducts() []*ProductGet {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductList) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ProductList) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_muallimah_submodule_protos_product_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_product_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x1a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x5a, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x47, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x74, 0x52, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x4e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72,
	0x6c, 0x22, 0xaa, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91,
	0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x12, 0x32,
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x80, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x34, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74, 0x12, 0x32,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_product_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_product_proto_rawDescData = file_muallimah_submodule_protos_product_proto_rawDesc
)

func file_muallimah_submodule_protos_product_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_product_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_product_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_product_proto_rawDescData
}

var file_muallimah_submodule_protos_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_muallimah_submodule_protos_product_proto_goTypes = []any{
	(*ProductCreate)(nil),  // 0: protos.ProductCreate
	(*ProductUpt)(nil),     // 1: protos.ProductUpt
	(*ProductUpdate)(nil),  // 2: protos.ProductUpdate
	(*ProductPicture)(nil), // 3: protos.ProductPicture
	(*ProductGet)(nil),     // 4: protos.ProductGet
	(*ProductFilter)(nil),  // 5: protos.ProductFilter
	(*ProductList)(nil),    // 6: protos.ProductList
	(*Pagination)(nil),     // 7: protos.Pagination
	(*ById)(nil),           // 8: protos.ById
	(*Void)(nil),           // 9: protos.Void
}
var file_muallimah_submodule_protos_product_proto_depIdxs = []int32{
	1,  // 0: protos.ProductUpdate.Body:type_name -> protos.ProductUpt
	7,  // 1: protos.ProductFilter.Pagination:type_name -> protos.Pagination
	4,  // 2: protos.ProductList.Products:type_name -> protos.ProductGet
	7,  // 3: protos.ProductList.Pagination:type_name -> protos.Pagination
	0,  // 4: protos.ProductService.CreateProduct:input_type -> protos.ProductCreate
	2,  // 5: protos.ProductService.UpdateProduct:input_type -> protos.ProductUpdate
	8,  // 6: protos.ProductService.DeleteProduct:input_type -> protos.ById
	5,  // 7: protos.ProductService.ListProducts:input_type -> protos.ProductFilter
	8,  // 8: protos.ProductService.GetProduct:input_type -> protos.ById
	3,  // 9: protos.ProductService.AddPicture:input_type -> protos.ProductPicture
	3,  // 10: protos.ProductService.DeletePicture:input_type -> protos.ProductPicture
	9,  // 11: protos.ProductService.CreateProduct:output_type -> protos.Void
	9,  // 12: protos.ProductService.UpdateProduct:output_type -> protos.Void
	9,  // 13: protos.ProductService.DeleteProduct:output_type -> protos.Void
	6,  // 14: protos.ProductService.ListProducts:output_type -> protos.ProductList
	4,  // 15: protos.ProductService.GetProduct:output_type -> protos.ProductGet
	9,  // 16: protos.ProductService.AddPicture:output_type -> protos.Void
	9,  // 17: protos.ProductService.DeletePicture:output_type -> protos.Void
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_product_proto_init() }
func file_muallimah_submodule_protos_product_proto_init() {
	if File_muallimah_submodule_protos_product_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_product_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProductCreate); i {
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
		file_muallimah_submodule_protos_product_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProductUpt); i {
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
		file_muallimah_submodule_protos_product_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ProductUpdate); i {
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
		file_muallimah_submodule_protos_product_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ProductPicture); i {
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
		file_muallimah_submodule_protos_product_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ProductGet); i {
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
		file_muallimah_submodule_protos_product_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ProductFilter); i {
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
		file_muallimah_submodule_protos_product_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ProductList); i {
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
			RawDescriptor: file_muallimah_submodule_protos_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_product_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_product_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_product_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_product_proto = out.File
	file_muallimah_submodule_protos_product_proto_rawDesc = nil
	file_muallimah_submodule_protos_product_proto_goTypes = nil
	file_muallimah_submodule_protos_product_proto_depIdxs = nil
}
