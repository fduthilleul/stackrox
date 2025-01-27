// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: api/v1/node_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ListNodesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClusterId     string                 `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNodesRequest) Reset() {
	*x = ListNodesRequest{}
	mi := &file_api_v1_node_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodesRequest) ProtoMessage() {}

func (x *ListNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_node_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodesRequest.ProtoReflect.Descriptor instead.
func (*ListNodesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_node_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListNodesRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type ListNodesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nodes         []*storage.Node        `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNodesResponse) Reset() {
	*x = ListNodesResponse{}
	mi := &file_api_v1_node_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodesResponse) ProtoMessage() {}

func (x *ListNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_node_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodesResponse.ProtoReflect.Descriptor instead.
func (*ListNodesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_node_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListNodesResponse) GetNodes() []*storage.Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClusterId     string                 `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	NodeId        string                 `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	mi := &file_api_v1_node_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_node_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_node_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodeRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetNodeRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type ExportNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timeout       int32                  `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Query         string                 `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportNodeRequest) Reset() {
	*x = ExportNodeRequest{}
	mi := &file_api_v1_node_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportNodeRequest) ProtoMessage() {}

func (x *ExportNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_node_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportNodeRequest.ProtoReflect.Descriptor instead.
func (*ExportNodeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_node_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExportNodeRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *ExportNodeRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type ExportNodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Node          *storage.Node          `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportNodeResponse) Reset() {
	*x = ExportNodeResponse{}
	mi := &file_api_v1_node_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportNodeResponse) ProtoMessage() {}

func (x *ExportNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_node_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportNodeResponse.ProtoReflect.Descriptor instead.
func (*ExportNodeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_node_service_proto_rawDescGZIP(), []int{4}
}

func (x *ExportNodeResponse) GetNode() *storage.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_api_v1_node_service_proto protoreflect.FileDescriptor

var file_api_v1_node_service_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x31, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x48,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x37, 0x0a,
	0x12, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x32, 0x99, 0x02, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x58, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x30, 0x01, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b,
	0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1_node_service_proto_rawDescOnce sync.Once
	file_api_v1_node_service_proto_rawDescData []byte
)

func file_api_v1_node_service_proto_rawDescGZIP() []byte {
	file_api_v1_node_service_proto_rawDescOnce.Do(func() {
		file_api_v1_node_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_node_service_proto_rawDesc), len(file_api_v1_node_service_proto_rawDesc)))
	})
	return file_api_v1_node_service_proto_rawDescData
}

var file_api_v1_node_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_node_service_proto_goTypes = []any{
	(*ListNodesRequest)(nil),   // 0: v1.ListNodesRequest
	(*ListNodesResponse)(nil),  // 1: v1.ListNodesResponse
	(*GetNodeRequest)(nil),     // 2: v1.GetNodeRequest
	(*ExportNodeRequest)(nil),  // 3: v1.ExportNodeRequest
	(*ExportNodeResponse)(nil), // 4: v1.ExportNodeResponse
	(*storage.Node)(nil),       // 5: storage.Node
}
var file_api_v1_node_service_proto_depIdxs = []int32{
	5, // 0: v1.ListNodesResponse.nodes:type_name -> storage.Node
	5, // 1: v1.ExportNodeResponse.node:type_name -> storage.Node
	0, // 2: v1.NodeService.ListNodes:input_type -> v1.ListNodesRequest
	2, // 3: v1.NodeService.GetNode:input_type -> v1.GetNodeRequest
	3, // 4: v1.NodeService.ExportNodes:input_type -> v1.ExportNodeRequest
	1, // 5: v1.NodeService.ListNodes:output_type -> v1.ListNodesResponse
	5, // 6: v1.NodeService.GetNode:output_type -> storage.Node
	4, // 7: v1.NodeService.ExportNodes:output_type -> v1.ExportNodeResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_node_service_proto_init() }
func file_api_v1_node_service_proto_init() {
	if File_api_v1_node_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_node_service_proto_rawDesc), len(file_api_v1_node_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_node_service_proto_goTypes,
		DependencyIndexes: file_api_v1_node_service_proto_depIdxs,
		MessageInfos:      file_api_v1_node_service_proto_msgTypes,
	}.Build()
	File_api_v1_node_service_proto = out.File
	file_api_v1_node_service_proto_goTypes = nil
	file_api_v1_node_service_proto_depIdxs = nil
}
