// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: redpanda/api/dataplane/v1alpha1/error.proto

package dataplanev1alpha1

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

type Reason int32

const (
	Reason_REASON_UNSPECIFIED              Reason = 0
	Reason_REASON_FEATURE_NOT_CONFIGURED   Reason = 1
	Reason_REASON_CONSOLE_ERROR            Reason = 2
	Reason_REASON_REDPANDA_ADMIN_API_ERROR Reason = 3
	Reason_REASON_KAFKA_API_ERROR          Reason = 4
	Reason_REASON_KAFKA_CONNECT_API_ERROR  Reason = 5
)

// Enum value maps for Reason.
var (
	Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "REASON_FEATURE_NOT_CONFIGURED",
		2: "REASON_CONSOLE_ERROR",
		3: "REASON_REDPANDA_ADMIN_API_ERROR",
		4: "REASON_KAFKA_API_ERROR",
		5: "REASON_KAFKA_CONNECT_API_ERROR",
	}
	Reason_value = map[string]int32{
		"REASON_UNSPECIFIED":              0,
		"REASON_FEATURE_NOT_CONFIGURED":   1,
		"REASON_CONSOLE_ERROR":            2,
		"REASON_REDPANDA_ADMIN_API_ERROR": 3,
		"REASON_KAFKA_API_ERROR":          4,
		"REASON_KAFKA_CONNECT_API_ERROR":  5,
	}
)

func (x Reason) Enum() *Reason {
	p := new(Reason)
	*p = x
	return p
}

func (x Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_dataplane_v1alpha1_error_proto_enumTypes[0].Descriptor()
}

func (Reason) Type() protoreflect.EnumType {
	return &file_redpanda_api_dataplane_v1alpha1_error_proto_enumTypes[0]
}

func (x Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reason.Descriptor instead.
func (Reason) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescGZIP(), []int{0}
}

var File_redpanda_api_dataplane_v1alpha1_error_proto protoreflect.FileDescriptor

var file_redpanda_api_dataplane_v1alpha1_error_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x72,
	0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2a, 0xc2,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43,
	0x4f, 0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x23,
	0x0a, 0x1f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x44, 0x50, 0x41, 0x4e, 0x44,
	0x41, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4b, 0x41,
	0x46, 0x4b, 0x41, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12,
	0x22, 0x0a, 0x1e, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x05, 0x42, 0xc5, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x52, 0x41, 0x44, 0xaa, 0x02, 0x1f, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e,
	0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2b, 0x52, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescOnce sync.Once
	file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescData = file_redpanda_api_dataplane_v1alpha1_error_proto_rawDesc
)

func file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescGZIP() []byte {
	file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescOnce.Do(func() {
		file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescData)
	})
	return file_redpanda_api_dataplane_v1alpha1_error_proto_rawDescData
}

var file_redpanda_api_dataplane_v1alpha1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_redpanda_api_dataplane_v1alpha1_error_proto_goTypes = []interface{}{
	(Reason)(0), // 0: redpanda.api.dataplane.v1alpha1.Reason
}
var file_redpanda_api_dataplane_v1alpha1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_redpanda_api_dataplane_v1alpha1_error_proto_init() }
func file_redpanda_api_dataplane_v1alpha1_error_proto_init() {
	if File_redpanda_api_dataplane_v1alpha1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_dataplane_v1alpha1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_api_dataplane_v1alpha1_error_proto_goTypes,
		DependencyIndexes: file_redpanda_api_dataplane_v1alpha1_error_proto_depIdxs,
		EnumInfos:         file_redpanda_api_dataplane_v1alpha1_error_proto_enumTypes,
	}.Build()
	File_redpanda_api_dataplane_v1alpha1_error_proto = out.File
	file_redpanda_api_dataplane_v1alpha1_error_proto_rawDesc = nil
	file_redpanda_api_dataplane_v1alpha1_error_proto_goTypes = nil
	file_redpanda_api_dataplane_v1alpha1_error_proto_depIdxs = nil
}
