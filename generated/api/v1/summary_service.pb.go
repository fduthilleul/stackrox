// TODO(ROX-24528): This API is deprecated, remove it in 4.7

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.3
// source: api/v1/summary_service.proto

package v1

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

type SummaryCountsResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	NumAlerts      int64                  `protobuf:"varint,1,opt,name=num_alerts,json=numAlerts,proto3" json:"num_alerts,omitempty"`
	NumClusters    int64                  `protobuf:"varint,2,opt,name=num_clusters,json=numClusters,proto3" json:"num_clusters,omitempty"`
	NumDeployments int64                  `protobuf:"varint,3,opt,name=num_deployments,json=numDeployments,proto3" json:"num_deployments,omitempty"`
	NumImages      int64                  `protobuf:"varint,4,opt,name=num_images,json=numImages,proto3" json:"num_images,omitempty"`
	NumSecrets     int64                  `protobuf:"varint,5,opt,name=num_secrets,json=numSecrets,proto3" json:"num_secrets,omitempty"`
	NumNodes       int64                  `protobuf:"varint,6,opt,name=num_nodes,json=numNodes,proto3" json:"num_nodes,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SummaryCountsResponse) Reset() {
	*x = SummaryCountsResponse{}
	mi := &file_api_v1_summary_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummaryCountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryCountsResponse) ProtoMessage() {}

func (x *SummaryCountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_summary_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryCountsResponse.ProtoReflect.Descriptor instead.
func (*SummaryCountsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_summary_service_proto_rawDescGZIP(), []int{0}
}

func (x *SummaryCountsResponse) GetNumAlerts() int64 {
	if x != nil {
		return x.NumAlerts
	}
	return 0
}

func (x *SummaryCountsResponse) GetNumClusters() int64 {
	if x != nil {
		return x.NumClusters
	}
	return 0
}

func (x *SummaryCountsResponse) GetNumDeployments() int64 {
	if x != nil {
		return x.NumDeployments
	}
	return 0
}

func (x *SummaryCountsResponse) GetNumImages() int64 {
	if x != nil {
		return x.NumImages
	}
	return 0
}

func (x *SummaryCountsResponse) GetNumSecrets() int64 {
	if x != nil {
		return x.NumSecrets
	}
	return 0
}

func (x *SummaryCountsResponse) GetNumNodes() int64 {
	if x != nil {
		return x.NumNodes
	}
	return 0
}

var File_api_v1_summary_service_proto protoreflect.FileDescriptor

var file_api_v1_summary_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x15, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e,
	0x75, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e,
	0x75, 0x6d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75,
	0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x66, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x09, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x27,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_summary_service_proto_rawDescOnce sync.Once
	file_api_v1_summary_service_proto_rawDescData = file_api_v1_summary_service_proto_rawDesc
)

func file_api_v1_summary_service_proto_rawDescGZIP() []byte {
	file_api_v1_summary_service_proto_rawDescOnce.Do(func() {
		file_api_v1_summary_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_summary_service_proto_rawDescData)
	})
	return file_api_v1_summary_service_proto_rawDescData
}

var file_api_v1_summary_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_summary_service_proto_goTypes = []any{
	(*SummaryCountsResponse)(nil), // 0: v1.SummaryCountsResponse
	(*Empty)(nil),                 // 1: v1.Empty
}
var file_api_v1_summary_service_proto_depIdxs = []int32{
	1, // 0: v1.SummaryService.GetSummaryCounts:input_type -> v1.Empty
	0, // 1: v1.SummaryService.GetSummaryCounts:output_type -> v1.SummaryCountsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_summary_service_proto_init() }
func file_api_v1_summary_service_proto_init() {
	if File_api_v1_summary_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_summary_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_summary_service_proto_goTypes,
		DependencyIndexes: file_api_v1_summary_service_proto_depIdxs,
		MessageInfos:      file_api_v1_summary_service_proto_msgTypes,
	}.Build()
	File_api_v1_summary_service_proto = out.File
	file_api_v1_summary_service_proto_rawDesc = nil
	file_api_v1_summary_service_proto_goTypes = nil
	file_api_v1_summary_service_proto_depIdxs = nil
}
