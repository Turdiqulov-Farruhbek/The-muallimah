// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: muallimah-submodule/protos/book.proto

package genproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_muallimah_submodule_protos_book_proto protoreflect.FileDescriptor

var file_muallimah_submodule_protos_book_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a,
	0x27, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x75, 0x61, 0x6c, 0x6c, 0x69,
	0x6d, 0x61, 0x68, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd9, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x31, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x32, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_muallimah_submodule_protos_book_proto_goTypes = []any{
	(*BookCreate)(nil),  // 0: protos.BookCreate
	(*BookUpdate)(nil),  // 1: protos.BookUpdate
	(*ById)(nil),        // 2: protos.ById
	(*BookFilter)(nil),  // 3: protos.BookFilter
	(*BookPicture)(nil), // 4: protos.BookPicture
	(*Void)(nil),        // 5: protos.Void
	(*BookList)(nil),    // 6: protos.BookList
	(*BookGet)(nil),     // 7: protos.BookGet
}
var file_muallimah_submodule_protos_book_proto_depIdxs = []int32{
	0, // 0: protos.BookService.CreateBook:input_type -> protos.BookCreate
	1, // 1: protos.BookService.UpdateBook:input_type -> protos.BookUpdate
	2, // 2: protos.BookService.DeleteBook:input_type -> protos.ById
	3, // 3: protos.BookService.ListBooks:input_type -> protos.BookFilter
	2, // 4: protos.BookService.GetBook:input_type -> protos.ById
	4, // 5: protos.BookService.AddPicture:input_type -> protos.BookPicture
	4, // 6: protos.BookService.DeletePicture:input_type -> protos.BookPicture
	5, // 7: protos.BookService.CreateBook:output_type -> protos.Void
	5, // 8: protos.BookService.UpdateBook:output_type -> protos.Void
	5, // 9: protos.BookService.DeleteBook:output_type -> protos.Void
	6, // 10: protos.BookService.ListBooks:output_type -> protos.BookList
	7, // 11: protos.BookService.GetBook:output_type -> protos.BookGet
	5, // 12: protos.BookService.AddPicture:output_type -> protos.Void
	5, // 13: protos.BookService.DeletePicture:output_type -> protos.Void
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_muallimah_submodule_protos_book_proto_init() }
func file_muallimah_submodule_protos_book_proto_init() {
	if File_muallimah_submodule_protos_book_proto != nil {
		return
	}
	file_muallimah_submodule_protos_common_proto_init()
	file_muallimah_submodule_protos_book_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_muallimah_submodule_protos_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_muallimah_submodule_protos_book_proto_goTypes,
		DependencyIndexes: file_muallimah_submodule_protos_book_proto_depIdxs,
	}.Build()
	File_muallimah_submodule_protos_book_proto = out.File
	file_muallimah_submodule_protos_book_proto_rawDesc = nil
	file_muallimah_submodule_protos_book_proto_goTypes = nil
	file_muallimah_submodule_protos_book_proto_depIdxs = nil
}