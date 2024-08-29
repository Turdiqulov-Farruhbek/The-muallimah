// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/post.proto

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

type PostCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content     string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	PictureUrls []string `protobuf:"bytes,3,rep,name=PictureUrls,proto3" json:"PictureUrls,omitempty"`
}

func (x *PostCreate) Reset() {
	*x = PostCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCreate) ProtoMessage() {}

func (x *PostCreate) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCreate.ProtoReflect.Descriptor instead.
func (*PostCreate) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostCreate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostCreate) GetPictureUrls() []string {
	if x != nil {
		return x.PictureUrls
	}
	return nil
}

type PostGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	PictureUrls []string `protobuf:"bytes,4,rep,name=PictureUrls,proto3" json:"PictureUrls,omitempty"`
	CreatedAt   string   `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *PostGet) Reset() {
	*x = PostGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGet) ProtoMessage() {}

func (x *PostGet) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGet.ProtoReflect.Descriptor instead.
func (*PostGet) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_post_proto_rawDescGZIP(), []int{1}
}

func (x *PostGet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostGet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostGet) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostGet) GetPictureUrls() []string {
	if x != nil {
		return x.PictureUrls
	}
	return nil
}

func (x *PostGet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PostUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *PostUpt) Reset() {
	*x = PostUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUpt) ProtoMessage() {}

func (x *PostUpt) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUpt.ProtoReflect.Descriptor instead.
func (*PostUpt) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_post_proto_rawDescGZIP(), []int{2}
}

func (x *PostUpt) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostUpt) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PostUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body *PostUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *PostUpdate) Reset() {
	*x = PostUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUpdate) ProtoMessage() {}

func (x *PostUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUpdate.ProtoReflect.Descriptor instead.
func (*PostUpdate) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_post_proto_rawDescGZIP(), []int{3}
}

func (x *PostUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostUpdate) GetBody() *PostUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type PostList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts      []*PostGet  `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts,omitempty"`
	TotalCount int32       `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	Pagination *Pagination `protobuf:"bytes,3,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *PostList) Reset() {
	*x = PostList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostList) ProtoMessage() {}

func (x *PostList) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostList.ProtoReflect.Descriptor instead.
func (*PostList) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_post_proto_rawDescGZIP(), []int{4}
}

func (x *PostList) GetPosts() []*PostGet {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *PostList) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *PostList) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type PostPicture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId     string `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	PictureUrl string `protobuf:"bytes,2,opt,name=PictureUrl,proto3" json:"PictureUrl,omitempty"`
}

func (x *PostPicture) Reset() {
	*x = PostPicture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_muallimah_submodule_protos_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPicture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPicture) ProtoMessage() {}

func (x *PostPicture) ProtoReflect() protoreflect.Message {
	mi := &file_muallimah_submodule_protos_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPicture.ProtoReflect.Descriptor instead.
func (*PostPicture) Descriptor() ([]byte, []int) {
	return file_muallimah_submodule_protos_post_proto_rawDescGZIP(), []int{5}
}

func (x *PostPicture) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostPicture) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

var File_muallimah_submodule_protos_post_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_post_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a,
	0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x73,
	0x74, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55,
	0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x70, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x41, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x70, 0x74, 0x52, 0x04, 0x42, 0x6f,
	0x64, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x0b, 0x50, 0x6f,
	0x73, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72,
	0x6c, 0x32, 0xe2, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x36,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_muallimah_submodule_protos_post_proto_rawDescOnce sync.Once
	file_muallimah_submodule_protos_post_proto_rawDescData = file_muallimah_submodule_protos_post_proto_rawDesc
)

func file_muallimah_submodule_protos_post_proto_rawDescGZIP() []byte {
	file_muallimah_submodule_protos_post_proto_rawDescOnce.Do(func() {
		file_muallimah_submodule_protos_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_muallimah_submodule_protos_post_proto_rawDescData)
	})
	return file_muallimah_submodule_protos_post_proto_rawDescData
}

var file_muallimah_submodule_protos_post_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_muallimah_submodule_protos_post_proto_goTypes = []interface{}{
	(*PostCreate)(nil),  // 0: protos.PostCreate
	(*PostGet)(nil),     // 1: protos.PostGet
	(*PostUpt)(nil),     // 2: protos.PostUpt
	(*PostUpdate)(nil),  // 3: protos.PostUpdate
	(*PostList)(nil),    // 4: protos.PostList
	(*PostPicture)(nil), // 5: protos.PostPicture
	(*Pagination)(nil),  // 6: protos.Pagination
	(*ById)(nil),        // 7: protos.ById
	(*Void)(nil),        // 8: protos.Void
}
var file_muallimah_submodule_protos_post_proto_depIdxs = []int32{
	2,  // 0: protos.PostUpdate.Body:type_name -> protos.PostUpt
	1,  // 1: protos.PostList.Posts:type_name -> protos.PostGet
	6,  // 2: protos.PostList.Pagination:type_name -> protos.Pagination
	0,  // 3: protos.PostService.CreatePost:input_type -> protos.PostCreate
	7,  // 4: protos.PostService.GetPost:input_type -> protos.ById
	3,  // 5: protos.PostService.UpdatePost:input_type -> protos.PostUpdate
	7,  // 6: protos.PostService.DeletePost:input_type -> protos.ById
	6,  // 7: protos.PostService.GetPosts:input_type -> protos.Pagination
	5,  // 8: protos.PostService.AddPostPicture:input_type -> protos.PostPicture
	5,  // 9: protos.PostService.DeletePostPicture:input_type -> protos.PostPicture
	8,  // 10: protos.PostService.CreatePost:output_type -> protos.Void
	1,  // 11: protos.PostService.GetPost:output_type -> protos.PostGet
	8,  // 12: protos.PostService.UpdatePost:output_type -> protos.Void
	8,  // 13: protos.PostService.DeletePost:output_type -> protos.Void
	4,  // 14: protos.PostService.GetPosts:output_type -> protos.PostList
	8,  // 15: protos.PostService.AddPostPicture:output_type -> protos.Void
	8,  // 16: protos.PostService.DeletePostPicture:output_type -> protos.Void
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_post_proto_init() }
func file_muallimah_submodule_protos_post_proto_init() {
	if File_muallimah_submodule_protos_post_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_muallimah_submodule_protos_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCreate); i {
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
		file_muallimah_submodule_protos_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostGet); i {
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
		file_muallimah_submodule_protos_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostUpt); i {
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
		file_muallimah_submodule_protos_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostUpdate); i {
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
		file_muallimah_submodule_protos_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostList); i {
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
		file_muallimah_submodule_protos_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPicture); i {
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
			RawDescriptor: file_muallimah_submodule_protos_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_post_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_post_proto_depIdxs,
		MessageInfos:      file_muallimah_submodule_protos_post_proto_msgTypes,
	}.Build()
	File_muallimah_submodule_protos_post_proto = out.File
	file_muallimah_submodule_protos_post_proto_rawDesc = nil
	file_muallimah_submodule_protos_post_proto_goTypes = nil
	file_muallimah_submodule_protos_post_proto_depIdxs = nil
}