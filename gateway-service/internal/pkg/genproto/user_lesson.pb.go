// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/user_lesson.proto

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

type UserLesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LessonId   string      `protobuf:"bytes,2,opt,name=LessonId,proto3" json:"LessonId,omitempty"`
	UserCourse *UserCourse `protobuf:"bytes,3,opt,name=UserCourse,proto3" json:"UserCourse,omitempty"`
	Status     string      `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"` // Enum type: progress_status_type
}

func (x *UserLesson) Reset() {
	*x = UserLesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLesson) ProtoMessage() {}

func (x *UserLesson) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLesson.ProtoReflect.Descriptor instead.
func (*UserLesson) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP(), []int{0}
}

func (x *UserLesson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserLesson) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *UserLesson) GetUserCourse() *UserCourse {
	if x != nil {
		return x.UserCourse
	}
	return nil
}

func (x *UserLesson) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UserLessonCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId     string `protobuf:"bytes,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	UserCourseId string `protobuf:"bytes,2,opt,name=UserCourseId,proto3" json:"UserCourseId,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserLessonCreateReq) Reset() {
	*x = UserLessonCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLessonCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLessonCreateReq) ProtoMessage() {}

func (x *UserLessonCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLessonCreateReq.ProtoReflect.Descriptor instead.
func (*UserLessonCreateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP(), []int{1}
}

func (x *UserLessonCreateReq) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *UserLessonCreateReq) GetUserCourseId() string {
	if x != nil {
		return x.UserCourseId
	}
	return ""
}

func (x *UserLessonCreateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UserLessonUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body *UserUptbody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UserLessonUpdateReq) Reset() {
	*x = UserLessonUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLessonUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLessonUpdateReq) ProtoMessage() {}

func (x *UserLessonUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLessonUpdateReq.ProtoReflect.Descriptor instead.
func (*UserLessonUpdateReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP(), []int{2}
}

func (x *UserLessonUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserLessonUpdateReq) GetBody() *UserUptbody {
	if x != nil {
		return x.Body
	}
	return nil
}

type UserUptbody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserUptbody) Reset() {
	*x = UserUptbody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUptbody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUptbody) ProtoMessage() {}

func (x *UserUptbody) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUptbody.ProtoReflect.Descriptor instead.
func (*UserUptbody) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP(), []int{3}
}

func (x *UserUptbody) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UserLessonListsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCourseId string      `protobuf:"bytes,1,opt,name=UserCourseId,proto3" json:"UserCourseId,omitempty"`
	Filter       *Pagination `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *UserLessonListsReq) Reset() {
	*x = UserLessonListsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLessonListsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLessonListsReq) ProtoMessage() {}

func (x *UserLessonListsReq) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLessonListsReq.ProtoReflect.Descriptor instead.
func (*UserLessonListsReq) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP(), []int{4}
}

func (x *UserLessonListsReq) GetUserCourseId() string {
	if x != nil {
		return x.UserCourseId
	}
	return ""
}

func (x *UserLessonListsReq) GetFilter() *Pagination {
	if x != nil {
		return x.Filter
	}
	return nil
}

type UserLessonListsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLessons []*UserLesson `protobuf:"bytes,1,rep,name=user_lessons,json=userLessons,proto3" json:"user_lessons,omitempty"`
	Pagination  *Pagination   `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	TotalCount  int32         `protobuf:"varint,3,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
}

func (x *UserLessonListsRes) Reset() {
	*x = UserLessonListsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLessonListsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLessonListsRes) ProtoMessage() {}

func (x *UserLessonListsRes) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_user_lesson_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLessonListsRes.ProtoReflect.Descriptor instead.
func (*UserLessonListsRes) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP(), []int{5}
}

func (x *UserLessonListsRes) GetUserLessons() []*UserLesson {
	if x != nil {
		return x.UserLessons
	}
	return nil
}

func (x *UserLessonListsRes) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *UserLessonListsRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_muallimah_submodule_protos_user_lesson_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_user_lesson_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61,
	0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84,
	0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6e, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x25, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x74,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x64, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x9f, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbf, 0x02, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_user_lesson_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_user_lesson_proto_rawDescData = file_muallimah_submodule_protos_user_lesson_proto_rawDesc
)

func file_muallimah_submodule_protos_user_lesson_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_user_lesson_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_user_lesson_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_user_lesson_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_user_lesson_proto_rawDescData
}

var file_muallimah_submodule_protos_user_lesson_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_muallimah_submodule_protos_user_lesson_proto_goTypes = []any{
	(*UserLesson)(nil),          // 0: protos.UserLesson
	(*UserLessonCreateReq)(nil), // 1: protos.UserLessonCreateReq
	(*UserLessonUpdateReq)(nil), // 2: protos.UserLessonUpdateReq
	(*UserUptbody)(nil),         // 3: protos.UserUptbody
	(*UserLessonListsReq)(nil),  // 4: protos.UserLessonListsReq
	(*UserLessonListsRes)(nil),  // 5: protos.UserLessonListsRes
	(*UserCourse)(nil),          // 6: protos.UserCourse
	(*Pagination)(nil),          // 7: protos.Pagination
	(*ById)(nil),                // 8: protos.ById
	(*Void)(nil),                // 9: protos.Void
}
var file_muallimah_submodule_protos_user_lesson_proto_depIdxs = []int32{
	6,  // 0: protos.UserLesson.UserCourse:type_name -> protos.UserCourse
	3,  // 1: protos.UserLessonUpdateReq.body:type_name -> protos.UserUptbody
	7,  // 2: protos.UserLessonListsReq.filter:type_name -> protos.Pagination
	0,  // 3: protos.UserLessonListsRes.user_lessons:type_name -> protos.UserLesson
	7,  // 4: protos.UserLessonListsRes.pagination:type_name -> protos.Pagination
	1,  // 5: protos.UserLessonService.CreateUserLesson:input_type -> protos.UserLessonCreateReq
	8,  // 6: protos.UserLessonService.GetUserLesson:input_type -> protos.ById
	2,  // 7: protos.UserLessonService.UpdateUserLesson:input_type -> protos.UserLessonUpdateReq
	8,  // 8: protos.UserLessonService.DeleteUserLesson:input_type -> protos.ById
	4,  // 9: protos.UserLessonService.ListUserLessons:input_type -> protos.UserLessonListsReq
	9,  // 10: protos.UserLessonService.CreateUserLesson:output_type -> protos.Void
	0,  // 11: protos.UserLessonService.GetUserLesson:output_type -> protos.UserLesson
	9,  // 12: protos.UserLessonService.UpdateUserLesson:output_type -> protos.Void
	9,  // 13: protos.UserLessonService.DeleteUserLesson:output_type -> protos.Void
	5,  // 14: protos.UserLessonService.ListUserLessons:output_type -> protos.UserLessonListsRes
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_user_lesson_proto_init() }
func file_muallimah_submodule_protos_user_lesson_proto_init() {
	if File_muallimah_submodule_protos_user_lesson_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	file_muallimah_submodule_protos_users_course_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_user_lesson_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserLesson); i {
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
		file_muallimah_submodule_protos_user_lesson_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserLessonCreateReq); i {
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
		file_muallimah_submodule_protos_user_lesson_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserLessonUpdateReq); i {
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
		file_muallimah_submodule_protos_user_lesson_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserUptbody); i {
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
		file_muallimah_submodule_protos_user_lesson_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserLessonListsReq); i {
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
		file_muallimah_submodule_protos_user_lesson_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UserLessonListsRes); i {
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
			RawDescriptor: file_muallimah_submodule_protos_user_lesson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_user_lesson_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_user_lesson_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_user_lesson_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_user_lesson_proto = out.File
	file_muallimah_submodule_protos_user_lesson_proto_rawDesc = nil
	file_muallimah_submodule_protos_user_lesson_proto_goTypes = nil
	file_muallimah_submodule_protos_user_lesson_proto_depIdxs = nil
}
