// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/api.proto

package api_pb

import (
	rss_feed "/blogs/rss_feed"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_proto_api_proto protoreflect.FileDescriptor

var file_proto_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x72, 0x73, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x79, 0x0a, 0x13, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x12, 0x62, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x53, 0x53, 0x46, 0x65, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x53, 0x53, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x22, 0x0a, 0x2f, 0x72, 0x73, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x42, 0x15,
	0x5a, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0xea, 0x02, 0x05,
	0x41, 0x70, 0x69, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_api_proto_goTypes = []interface{}{
	(*rss_feed.CreateRSSFeedRequest)(nil), // 0: portfolio.blogs.CreateRSSFeedRequest
	(*empty.Empty)(nil),                   // 1: google.protobuf.Empty
}
var file_proto_api_proto_depIdxs = []int32{
	0, // 0: portfolio.PortfolioBackendApi.CreateRSSFeed:input_type -> portfolio.blogs.CreateRSSFeedRequest
	1, // 1: portfolio.PortfolioBackendApi.CreateRSSFeed:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_api_proto_init() }
func file_proto_api_proto_init() {
	if File_proto_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_proto_goTypes,
		DependencyIndexes: file_proto_api_proto_depIdxs,
	}.Build()
	File_proto_api_proto = out.File
	file_proto_api_proto_rawDesc = nil
	file_proto_api_proto_goTypes = nil
	file_proto_api_proto_depIdxs = nil
}
