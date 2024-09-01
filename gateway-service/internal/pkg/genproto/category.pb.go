// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/category.proto

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

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_category_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CategoryCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CategoryCreateReq) Reset() {
	*x = CategoryCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryCreateReq) ProtoMessage() {}

func (x *CategoryCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryCreateReq.ProtoReflect.Descriptor instead.
func (*CategoryCreateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_category_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CategoryUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CategoryUpt) Reset() {
	*x = CategoryUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpt) ProtoMessage() {}

func (x *CategoryUpt) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpt.ProtoReflect.Descriptor instead.
func (*CategoryUpt) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_category_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryUpt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CategoryUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body *CategoryUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *CategoryUpdateReq) Reset() {
	*x = CategoryUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpdateReq) ProtoMessage() {}

func (x *CategoryUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpdateReq.ProtoReflect.Descriptor instead.
func (*CategoryUpdateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_category_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryUpdateReq) GetBody() *CategoryUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type CategoryListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	TotalCount int32       `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CategoryListRes) Reset() {
	*x = CategoryListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListRes) ProtoMessage() {}

func (x *CategoryListRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListRes.ProtoReflect.Descriptor instead.
func (*CategoryListRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_category_proto_rawDescGZIP(), []int{4}
}

func (x *CategoryListRes) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *CategoryListRes) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *CategoryListRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_muallimah_submodule_protos_category_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_category_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x1a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73,
	0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x55, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x11, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x74,
	0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x98, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xa3, 0x02, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x12, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x39, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_category_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_category_proto_rawDescData = file_muallimah_submodule_protos_category_proto_rawDesc
)

func file_muallimah_submodule_protos_category_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_category_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_category_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_category_proto_rawDescData
}

var file_muallimah_submodule_protos_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_muallimah_submodule_protos_category_proto_goTypes = []any{
	(*Category)(nil),          // 0: protos.Category
	(*CategoryCreateReq)(nil), // 1: protos.CategoryCreateReq
	(*CategoryUpt)(nil),       // 2: protos.CategoryUpt
	(*CategoryUpdateReq)(nil), // 3: protos.CategoryUpdateReq
	(*CategoryListRes)(nil),   // 4: protos.CategoryListRes
	(*Pagination)(nil),        // 5: protos.Pagination
	(*ById)(nil),              // 6: protos.ById
	(*Void)(nil),              // 7: protos.Void
}
var file_muallimah_submodule_protos_category_proto_depIdxs = []int32{
	2, // 0: protos.CategoryUpdateReq.Body:type_name -> protos.CategoryUpt
	0, // 1: protos.CategoryListRes.categories:type_name -> protos.Category
	5, // 2: protos.CategoryListRes.pagination:type_name -> protos.Pagination
	1, // 3: protos.CategoryService.CreateCategory:input_type -> protos.CategoryCreateReq
	6, // 4: protos.CategoryService.GetCategory:input_type -> protos.ById
	3, // 5: protos.CategoryService.UpdateCategory:input_type -> protos.CategoryUpdateReq
	6, // 6: protos.CategoryService.DeleteCategory:input_type -> protos.ById
	5, // 7: protos.CategoryService.ListCategories:input_type -> protos.Pagination
	7, // 8: protos.CategoryService.CreateCategory:output_type -> protos.Void
	0, // 9: protos.CategoryService.GetCategory:output_type -> protos.Category
	7, // 10: protos.CategoryService.UpdateCategory:output_type -> protos.Void
	7, // 11: protos.CategoryService.DeleteCategory:output_type -> protos.Void
	4, // 12: protos.CategoryService.ListCategories:output_type -> protos.CategoryListRes
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_category_proto_init() }
func file_muallimah_submodule_protos_category_proto_init() {
	if File_muallimah_submodule_protos_category_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_category_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Category); i {
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
		file_muallimah_submodule_protos_category_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryCreateReq); i {
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
		file_muallimah_submodule_protos_category_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryUpt); i {
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
		file_muallimah_submodule_protos_category_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryUpdateReq); i {
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
		file_muallimah_submodule_protos_category_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryListRes); i {
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
			RawDescriptor: file_muallimah_submodule_protos_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_category_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_category_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_category_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_category_proto = out.File
	file_muallimah_submodule_protos_category_proto_rawDesc = nil
	file_muallimah_submodule_protos_category_proto_goTypes = nil
	file_muallimah_submodule_protos_category_proto_depIdxs = nil
}