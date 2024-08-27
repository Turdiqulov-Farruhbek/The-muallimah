// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/order.proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId     string `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Quantity   int32  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalPrice string `protobuf:"bytes,6,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status     string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Order) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Order) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order) GetTotalPrice() string {
	if x != nil {
		return x.TotalPrice
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId     string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Quantity   int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalPrice string `protobuf:"bytes,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status     string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderCreateReq) Reset() {
	*x = OrderCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateReq) ProtoMessage() {}

func (x *OrderCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateReq.ProtoReflect.Descriptor instead.
func (*OrderCreateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCreateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderCreateReq) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *OrderCreateReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderCreateReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderCreateReq) GetTotalPrice() string {
	if x != nil {
		return x.TotalPrice
	}
	return ""
}

func (x *OrderCreateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderGetReq) Reset() {
	*x = OrderGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetReq) ProtoMessage() {}

func (x *OrderGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetReq.ProtoReflect.Descriptor instead.
func (*OrderGetReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderGetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrderGetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderGetRes) Reset() {
	*x = OrderGetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetRes) ProtoMessage() {}

func (x *OrderGetRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetRes.ProtoReflect.Descriptor instead.
func (*OrderGetRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderGetRes) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body *OrderUpt `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *OrderUpdateReq) Reset() {
	*x = OrderUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateReq) ProtoMessage() {}

func (x *OrderUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateReq.ProtoReflect.Descriptor instead.
func (*OrderUpdateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderUpdateReq) GetBody() *OrderUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type OrderUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId     string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Quantity   int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalPrice string `protobuf:"bytes,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status     string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderUpt) Reset() {
	*x = OrderUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpt) ProtoMessage() {}

func (x *OrderUpt) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpt.ProtoReflect.Descriptor instead.
func (*OrderUpt) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderUpt) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *OrderUpt) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderUpt) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderUpt) GetTotalPrice() string {
	if x != nil {
		return x.TotalPrice
	}
	return ""
}

func (x *OrderUpt) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderDeleteReq) Reset() {
	*x = OrderDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDeleteReq) ProtoMessage() {}

func (x *OrderDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDeleteReq.ProtoReflect.Descriptor instead.
func (*OrderDeleteReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderDeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrderListsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Filter *Pagination `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *OrderListsReq) Reset() {
	*x = OrderListsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListsReq) ProtoMessage() {}

func (x *OrderListsReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListsReq.ProtoReflect.Descriptor instead.
func (*OrderListsReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderListsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderListsReq) GetFilter() *Pagination {
	if x != nil {
		return x.Filter
	}
	return nil
}

type OrderListsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders     []*Order    `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	TotalCount int32       `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *OrderListsRes) Reset() {
	*x = OrderListsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListsRes) ProtoMessage() {}

func (x *OrderListsRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListsRes.ProtoReflect.Descriptor instead.
func (*OrderListsRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_order_proto_rawDescGZIP(), []int{8}
}

func (x *OrderListsRes) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *OrderListsRes) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *OrderListsRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_muallimah_submodule_protos_order_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_order_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x1a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xab,
	0x01, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x46, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70,
	0x74, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x70, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8b,
	0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9f, 0x02, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x33, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x42, 0x18,
	0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_order_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_order_proto_rawDescData = file_muallimah_submodule_protos_order_proto_rawDesc
)

func file_muallimah_submodule_protos_order_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_order_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_order_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_order_proto_rawDescData
}

var file_muallimah_submodule_protos_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_muallimah_submodule_protos_order_proto_goTypes = []any{
	(*Order)(nil),          // 0: protos.Order
	(*OrderCreateReq)(nil), // 1: protos.OrderCreateReq
	(*OrderGetReq)(nil),    // 2: protos.OrderGetReq
	(*OrderGetRes)(nil),    // 3: protos.OrderGetRes
	(*OrderUpdateReq)(nil), // 4: protos.OrderUpdateReq
	(*OrderUpt)(nil),       // 5: protos.OrderUpt
	(*OrderDeleteReq)(nil), // 6: protos.OrderDeleteReq
	(*OrderListsReq)(nil),  // 7: protos.OrderListsReq
	(*OrderListsRes)(nil),  // 8: protos.OrderListsRes
	(*Pagination)(nil),     // 9: protos.Pagination
	(*Void)(nil),           // 10: protos.Void
}
var file_muallimah_submodule_protos_order_proto_depIdxs = []int32{
	0,  // 0: protos.OrderGetRes.order:type_name -> protos.Order
	5,  // 1: protos.OrderUpdateReq.body:type_name -> protos.OrderUpt
	9,  // 2: protos.OrderListsReq.filter:type_name -> protos.Pagination
	0,  // 3: protos.OrderListsRes.orders:type_name -> protos.Order
	9,  // 4: protos.OrderListsRes.pagination:type_name -> protos.Pagination
	1,  // 5: protos.OrderService.CreateOrder:input_type -> protos.OrderCreateReq
	2,  // 6: protos.OrderService.GetOrder:input_type -> protos.OrderGetReq
	4,  // 7: protos.OrderService.UpdateOrder:input_type -> protos.OrderUpdateReq
	6,  // 8: protos.OrderService.DeleteOrder:input_type -> protos.OrderDeleteReq
	7,  // 9: protos.OrderService.ListOrders:input_type -> protos.OrderListsReq
	10, // 10: protos.OrderService.CreateOrder:output_type -> protos.Void
	3,  // 11: protos.OrderService.GetOrder:output_type -> protos.OrderGetRes
	10, // 12: protos.OrderService.UpdateOrder:output_type -> protos.Void
	10, // 13: protos.OrderService.DeleteOrder:output_type -> protos.Void
	8,  // 14: protos.OrderService.ListOrders:output_type -> protos.OrderListsRes
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_order_proto_init() }
func file_muallimah_submodule_protos_order_proto_init() {
	if File_muallimah_submodule_protos_order_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_order_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Order); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OrderCreateReq); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OrderGetReq); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*OrderGetRes); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OrderUpdateReq); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OrderUpt); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*OrderDeleteReq); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*OrderListsReq); i {
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
		file_muallimah_submodule_protos_order_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*OrderListsRes); i {
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
			RawDescriptor: file_muallimah_submodule_protos_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_order_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_order_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_order_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_order_proto = out.File
	file_muallimah_submodule_protos_order_proto_rawDesc = nil
	file_muallimah_submodule_protos_order_proto_goTypes = nil
	file_muallimah_submodule_protos_order_proto_depIdxs = nil
}
