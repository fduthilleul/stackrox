// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: api/v1/cve_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SuppressCVERequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// These are (NVD) vulnerability identifiers, `cve` field of `storage.CVE`, and *not* the `id` field.
	// For example, CVE-2021-44832.
	Cves []string `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by the number of seconds,
	// with nanoseconds expressed as fractional seconds.
	// For example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as "3.000000001s",
	// and 3 seconds and 1 microsecond should be expressed in JSON format as "3.000001s".
	Duration      *durationpb.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuppressCVERequest) Reset() {
	*x = SuppressCVERequest{}
	mi := &file_api_v1_cve_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuppressCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuppressCVERequest) ProtoMessage() {}

func (x *SuppressCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cve_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuppressCVERequest.ProtoReflect.Descriptor instead.
func (*SuppressCVERequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cve_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuppressCVERequest) GetCves() []string {
	if x != nil {
		return x.Cves
	}
	return nil
}

func (x *SuppressCVERequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type UnsuppressCVERequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// These are (NVD) vulnerability identifiers, `cve` field of `storage.CVE`, and *not* the `id` field.
	// For example, CVE-2021-44832.
	Cves          []string `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnsuppressCVERequest) Reset() {
	*x = UnsuppressCVERequest{}
	mi := &file_api_v1_cve_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsuppressCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsuppressCVERequest) ProtoMessage() {}

func (x *UnsuppressCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cve_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsuppressCVERequest.ProtoReflect.Descriptor instead.
func (*UnsuppressCVERequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cve_service_proto_rawDescGZIP(), []int{1}
}

func (x *UnsuppressCVERequest) GetCves() []string {
	if x != nil {
		return x.Cves
	}
	return nil
}

var File_api_v1_cve_service_proto protoreflect.FileDescriptor

var file_api_v1_cve_service_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x12,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x76, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x2a, 0x0a, 0x14, 0x55, 0x6e, 0x73, 0x75, 0x70,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x76, 0x65, 0x73, 0x32, 0xc3, 0x01, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x56, 0x45,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0c, 0x53, 0x75, 0x70, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x73, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x32, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x63, 0x76, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5a, 0x0a,
	0x0e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x73, 0x12,
	0x18, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x32,
	0x18, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x63, 0x76, 0x65, 0x73, 0x2f, 0x75,
	0x6e, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x32, 0xc0, 0x01, 0x0a, 0x0e, 0x4e, 0x6f,
	0x64, 0x65, 0x43, 0x56, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c,
	0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x73, 0x12, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x32, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x63, 0x76, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x59, 0x0a, 0x0e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x56, 0x45, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x3a, 0x01, 0x2a, 0x32, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x63, 0x76, 0x65,
	0x73, 0x2f, 0x75, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x32, 0xc9, 0x01, 0x0a,
	0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x56, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x56, 0x0a, 0x0c, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56,
	0x45, 0x73, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a,
	0x32, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x76, 0x65,
	0x73, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5c, 0x0a, 0x0e, 0x55, 0x6e,
	0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x73, 0x12, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x56, 0x45, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x32, 0x1a, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x76, 0x65, 0x73, 0x2f, 0x75, 0x6e,
	0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1_cve_service_proto_rawDescOnce sync.Once
	file_api_v1_cve_service_proto_rawDescData []byte
)

func file_api_v1_cve_service_proto_rawDescGZIP() []byte {
	file_api_v1_cve_service_proto_rawDescOnce.Do(func() {
		file_api_v1_cve_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_cve_service_proto_rawDesc), len(file_api_v1_cve_service_proto_rawDesc)))
	})
	return file_api_v1_cve_service_proto_rawDescData
}

var file_api_v1_cve_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_cve_service_proto_goTypes = []any{
	(*SuppressCVERequest)(nil),   // 0: v1.SuppressCVERequest
	(*UnsuppressCVERequest)(nil), // 1: v1.UnsuppressCVERequest
	(*durationpb.Duration)(nil),  // 2: google.protobuf.Duration
	(*Empty)(nil),                // 3: v1.Empty
}
var file_api_v1_cve_service_proto_depIdxs = []int32{
	2, // 0: v1.SuppressCVERequest.duration:type_name -> google.protobuf.Duration
	0, // 1: v1.ImageCVEService.SuppressCVEs:input_type -> v1.SuppressCVERequest
	1, // 2: v1.ImageCVEService.UnsuppressCVEs:input_type -> v1.UnsuppressCVERequest
	0, // 3: v1.NodeCVEService.SuppressCVEs:input_type -> v1.SuppressCVERequest
	1, // 4: v1.NodeCVEService.UnsuppressCVEs:input_type -> v1.UnsuppressCVERequest
	0, // 5: v1.ClusterCVEService.SuppressCVEs:input_type -> v1.SuppressCVERequest
	1, // 6: v1.ClusterCVEService.UnsuppressCVEs:input_type -> v1.UnsuppressCVERequest
	3, // 7: v1.ImageCVEService.SuppressCVEs:output_type -> v1.Empty
	3, // 8: v1.ImageCVEService.UnsuppressCVEs:output_type -> v1.Empty
	3, // 9: v1.NodeCVEService.SuppressCVEs:output_type -> v1.Empty
	3, // 10: v1.NodeCVEService.UnsuppressCVEs:output_type -> v1.Empty
	3, // 11: v1.ClusterCVEService.SuppressCVEs:output_type -> v1.Empty
	3, // 12: v1.ClusterCVEService.UnsuppressCVEs:output_type -> v1.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_cve_service_proto_init() }
func file_api_v1_cve_service_proto_init() {
	if File_api_v1_cve_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_cve_service_proto_rawDesc), len(file_api_v1_cve_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_api_v1_cve_service_proto_goTypes,
		DependencyIndexes: file_api_v1_cve_service_proto_depIdxs,
		MessageInfos:      file_api_v1_cve_service_proto_msgTypes,
	}.Build()
	File_api_v1_cve_service_proto = out.File
	file_api_v1_cve_service_proto_goTypes = nil
	file_api_v1_cve_service_proto_depIdxs = nil
}
