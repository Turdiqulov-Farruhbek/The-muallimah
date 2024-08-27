// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/transaction.proto

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

type TransactionCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCourseId string  `protobuf:"bytes,1,opt,name=user_course_id,json=userCourseId,proto3" json:"user_course_id,omitempty"`
	Amount       float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Type         string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TransactionCreateReq) Reset() {
	*x = TransactionCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionCreateReq) ProtoMessage() {}

func (x *TransactionCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionCreateReq.ProtoReflect.Descriptor instead.
func (*TransactionCreateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionCreateReq) GetUserCourseId() string {
	if x != nil {
		return x.UserCourseId
	}
	return ""
}

func (x *TransactionCreateReq) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionCreateReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type TransactionGetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserCourseId string  `protobuf:"bytes,2,opt,name=user_course_id,json=userCourseId,proto3" json:"user_course_id,omitempty"`
	Amount       float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Type         string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TransactionGetRes) Reset() {
	*x = TransactionGetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetRes) ProtoMessage() {}

func (x *TransactionGetRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetRes.ProtoReflect.Descriptor instead.
func (*TransactionGetRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionGetRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionGetRes) GetUserCourseId() string {
	if x != nil {
		return x.UserCourseId
	}
	return ""
}

func (x *TransactionGetRes) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionGetRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type TransactionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCourseId string  `protobuf:"bytes,1,opt,name=user_course_id,json=userCourseId,proto3" json:"user_course_id,omitempty"`
	AmountFrom   float32 `protobuf:"fixed32,2,opt,name=amount_from,json=amountFrom,proto3" json:"amount_from,omitempty"`
	AmountTo     float32 `protobuf:"fixed32,3,opt,name=amount_to,json=amountTo,proto3" json:"amount_to,omitempty"`
	Type         string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TransactionListReq) Reset() {
	*x = TransactionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionListReq) ProtoMessage() {}

func (x *TransactionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionListReq.ProtoReflect.Descriptor instead.
func (*TransactionListReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionListReq) GetUserCourseId() string {
	if x != nil {
		return x.UserCourseId
	}
	return ""
}

func (x *TransactionListReq) GetAmountFrom() float32 {
	if x != nil {
		return x.AmountFrom
	}
	return 0
}

func (x *TransactionListReq) GetAmountTo() float32 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *TransactionListReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type TransactionListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*TransactionGetRes `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionListRes) Reset() {
	*x = TransactionListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionListRes) ProtoMessage() {}

func (x *TransactionListRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionListRes.ProtoReflect.Descriptor instead.
func (*TransactionListRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionListRes) GetTransactions() []*TransactionGetRes {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_muallimah_submodule_protos_transaction_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_transaction_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61,
	0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x75, 0x0a, 0x11, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x8c, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x53, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0xdc, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_transaction_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_transaction_proto_rawDescData = file_muallimah_submodule_protos_transaction_proto_rawDesc
)

func file_muallimah_submodule_protos_transaction_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_transaction_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_transaction_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_transaction_proto_rawDescData
}

var file_muallimah_submodule_protos_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_muallimah_submodule_protos_transaction_proto_goTypes = []any{
	(*TransactionCreateReq)(nil), // 0: protos.TransactionCreateReq
	(*TransactionGetRes)(nil),    // 1: protos.TransactionGetRes
	(*TransactionListReq)(nil),   // 2: protos.TransactionListReq
	(*TransactionListRes)(nil),   // 3: protos.TransactionListRes
	(*ById)(nil),                 // 4: protos.ById
	(*Void)(nil),                 // 5: protos.Void
}
var file_muallimah_submodule_protos_transaction_proto_depIdxs = []int32{
	1, // 0: protos.TransactionListRes.transactions:type_name -> protos.TransactionGetRes
	0, // 1: protos.TransactionService.CreateTransaction:input_type -> protos.TransactionCreateReq
	4, // 2: protos.TransactionService.GetTransaction:input_type -> protos.ById
	2, // 3: protos.TransactionService.ListTransactions:input_type -> protos.TransactionListReq
	5, // 4: protos.TransactionService.CreateTransaction:output_type -> protos.Void
	1, // 5: protos.TransactionService.GetTransaction:output_type -> protos.TransactionGetRes
	3, // 6: protos.TransactionService.ListTransactions:output_type -> protos.TransactionListRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_transaction_proto_init() }
func file_muallimah_submodule_protos_transaction_proto_init() {
	if File_muallimah_submodule_protos_transaction_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_transaction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionCreateReq); i {
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
		file_muallimah_submodule_protos_transaction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionGetRes); i {
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
		file_muallimah_submodule_protos_transaction_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionListReq); i {
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
		file_muallimah_submodule_protos_transaction_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionListRes); i {
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
			RawDescriptor: file_muallimah_submodule_protos_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_transaction_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_transaction_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_transaction_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_transaction_proto = out.File
	file_muallimah_submodule_protos_transaction_proto_rawDesc = nil
	file_muallimah_submodule_protos_transaction_proto_goTypes = nil
	file_muallimah_submodule_protos_transaction_proto_depIdxs = nil
}
