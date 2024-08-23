// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: muallimah-submodule/protos/course.proto

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

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string  `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	CategoryId  string  `protobuf:"bytes,6,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CreatedAt   string  `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Course) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Course) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Course) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Course) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CourseCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string  `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	CategoryId  string  `protobuf:"bytes,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CourseCreateReq) Reset() {
	*x = CourseCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseCreateReq) ProtoMessage() {}

func (x *CourseCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseCreateReq.ProtoReflect.Descriptor instead.
func (*CourseCreateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_course_proto_rawDescGZIP(), []int{1}
}

func (x *CourseCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CourseCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CourseCreateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CourseCreateReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CourseCreateReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type CourseListsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string      `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Filter     *Pagination `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *CourseListsReq) Reset() {
	*x = CourseListsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseListsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseListsReq) ProtoMessage() {}

func (x *CourseListsReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseListsReq.ProtoReflect.Descriptor instead.
func (*CourseListsReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_course_proto_rawDescGZIP(), []int{2}
}

func (x *CourseListsReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CourseListsReq) GetFilter() *Pagination {
	if x != nil {
		return x.Filter
	}
	return nil
}

type CourseListsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses    []*Course   `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	TotalCount int32       `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CourseListsRes) Reset() {
	*x = CourseListsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseListsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseListsRes) ProtoMessage() {}

func (x *CourseListsRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseListsRes.ProtoReflect.Descriptor instead.
func (*CourseListsRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_course_proto_rawDescGZIP(), []int{3}
}

func (x *CourseListsRes) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

func (x *CourseListsRes) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *CourseListsRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CourseUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string  `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *CourseUpt) Reset() {
	*x = CourseUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseUpt) ProtoMessage() {}

func (x *CourseUpt) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseUpt.ProtoReflect.Descriptor instead.
func (*CourseUpt) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_course_proto_rawDescGZIP(), []int{4}
}

func (x *CourseUpt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CourseUpt) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CourseUpt) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CourseUpt) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type CourseUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Body *CourseUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *CourseUpdateReq) Reset() {
	*x = CourseUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_course_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseUpdateReq) ProtoMessage() {}

func (x *CourseUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_course_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseUpdateReq.ProtoReflect.Descriptor instead.
func (*CourseUpdateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_course_proto_rawDescGZIP(), []int{5}
}

func (x *CourseUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CourseUpdateReq) GetBody() *CourseUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_muallimah_submodule_protos_course_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_course_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x1a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9b, 0x01,
	0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x0e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x74, 0x0a, 0x09,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x55, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0x48, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x55, 0x70, 0x74, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x32, 0x93, 0x02, 0x0a,
	0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_course_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_course_proto_rawDescData = file_muallimah_submodule_protos_course_proto_rawDesc
)

func file_muallimah_submodule_protos_course_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_course_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_course_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_course_proto_rawDescData
}

var file_muallimah_submodule_protos_course_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_muallimah_submodule_protos_course_proto_goTypes = []any{
	(*Course)(nil),          // 0: protos.Course
	(*CourseCreateReq)(nil), // 1: protos.CourseCreateReq
	(*CourseListsReq)(nil),  // 2: protos.CourseListsReq
	(*CourseListsRes)(nil),  // 3: protos.CourseListsRes
	(*CourseUpt)(nil),       // 4: protos.CourseUpt
	(*CourseUpdateReq)(nil), // 5: protos.CourseUpdateReq
	(*Pagination)(nil),      // 6: protos.Pagination
	(*ById)(nil),            // 7: protos.ById
	(*Void)(nil),            // 8: protos.Void
}
var file_muallimah_submodule_protos_course_proto_depIdxs = []int32{
	6, // 0: protos.CourseListsReq.filter:type_name -> protos.Pagination
	0, // 1: protos.CourseListsRes.courses:type_name -> protos.Course
	6, // 2: protos.CourseListsRes.pagination:type_name -> protos.Pagination
	4, // 3: protos.CourseUpdateReq.Body:type_name -> protos.CourseUpt
	1, // 4: protos.CourseService.CreateCourse:input_type -> protos.CourseCreateReq
	7, // 5: protos.CourseService.GetCourse:input_type -> protos.ById
	5, // 6: protos.CourseService.UpdateCourse:input_type -> protos.CourseUpdateReq
	7, // 7: protos.CourseService.DeleteCourse:input_type -> protos.ById
	2, // 8: protos.CourseService.ListCourses:input_type -> protos.CourseListsReq
	8, // 9: protos.CourseService.CreateCourse:output_type -> protos.Void
	0, // 10: protos.CourseService.GetCourse:output_type -> protos.Course
	8, // 11: protos.CourseService.UpdateCourse:output_type -> protos.Void
	8, // 12: protos.CourseService.DeleteCourse:output_type -> protos.Void
	3, // 13: protos.CourseService.ListCourses:output_type -> protos.CourseListsRes
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_course_proto_init() }
func file_muallimah_submodule_protos_course_proto_init() {
	if File_muallimah_submodule_protos_course_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_course_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Course); i {
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
		file_muallimah_submodule_protos_course_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CourseCreateReq); i {
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
		file_muallimah_submodule_protos_course_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CourseListsReq); i {
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
		file_muallimah_submodule_protos_course_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CourseListsRes); i {
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
		file_muallimah_submodule_protos_course_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CourseUpt); i {
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
		file_muallimah_submodule_protos_course_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CourseUpdateReq); i {
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
			RawDescriptor: file_muallimah_submodule_protos_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_course_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_course_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_course_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_course_proto = out.File
	file_muallimah_submodule_protos_course_proto_rawDesc = nil
	file_muallimah_submodule_protos_course_proto_goTypes = nil
	file_muallimah_submodule_protos_course_proto_depIdxs = nil
}
