// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: internalapi/central/development_service.proto

// This is an internal service which contains tools intended to be used only for testing.
// It will NOT be available in Central in production builds.

package central

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type URLHasValidCertResponse_URLResult int32

const (
	URLHasValidCertResponse_UNSET                                        URLHasValidCertResponse_URLResult = 0
	URLHasValidCertResponse_CERT_SIGNED_BY_UNKNOWN_AUTHORITY             URLHasValidCertResponse_URLResult = 1
	URLHasValidCertResponse_CERT_SIGNING_AUTHORITY_VALID_BUT_OTHER_ERROR URLHasValidCertResponse_URLResult = 2
	URLHasValidCertResponse_OTHER_GET_ERROR                              URLHasValidCertResponse_URLResult = 3
	URLHasValidCertResponse_REQUEST_SUCCEEDED                            URLHasValidCertResponse_URLResult = 4
)

// Enum value maps for URLHasValidCertResponse_URLResult.
var (
	URLHasValidCertResponse_URLResult_name = map[int32]string{
		0: "UNSET",
		1: "CERT_SIGNED_BY_UNKNOWN_AUTHORITY",
		2: "CERT_SIGNING_AUTHORITY_VALID_BUT_OTHER_ERROR",
		3: "OTHER_GET_ERROR",
		4: "REQUEST_SUCCEEDED",
	}
	URLHasValidCertResponse_URLResult_value = map[string]int32{
		"UNSET":                            0,
		"CERT_SIGNED_BY_UNKNOWN_AUTHORITY": 1,
		"CERT_SIGNING_AUTHORITY_VALID_BUT_OTHER_ERROR": 2,
		"OTHER_GET_ERROR":   3,
		"REQUEST_SUCCEEDED": 4,
	}
)

func (x URLHasValidCertResponse_URLResult) Enum() *URLHasValidCertResponse_URLResult {
	p := new(URLHasValidCertResponse_URLResult)
	*p = x
	return p
}

func (x URLHasValidCertResponse_URLResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (URLHasValidCertResponse_URLResult) Descriptor() protoreflect.EnumDescriptor {
	return file_internalapi_central_development_service_proto_enumTypes[0].Descriptor()
}

func (URLHasValidCertResponse_URLResult) Type() protoreflect.EnumType {
	return &file_internalapi_central_development_service_proto_enumTypes[0]
}

func (x URLHasValidCertResponse_URLResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use URLHasValidCertResponse_URLResult.Descriptor instead.
func (URLHasValidCertResponse_URLResult) EnumDescriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{1, 0}
}

type URLHasValidCertRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CertPEM       string                 `protobuf:"bytes,2,opt,name=certPEM,proto3" json:"certPEM,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *URLHasValidCertRequest) Reset() {
	*x = URLHasValidCertRequest{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *URLHasValidCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLHasValidCertRequest) ProtoMessage() {}

func (x *URLHasValidCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLHasValidCertRequest.ProtoReflect.Descriptor instead.
func (*URLHasValidCertRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{0}
}

func (x *URLHasValidCertRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *URLHasValidCertRequest) GetCertPEM() string {
	if x != nil {
		return x.CertPEM
	}
	return ""
}

type URLHasValidCertResponse struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Result        URLHasValidCertResponse_URLResult `protobuf:"varint,1,opt,name=result,proto3,enum=central.URLHasValidCertResponse_URLResult" json:"result,omitempty"`
	Details       string                            `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *URLHasValidCertResponse) Reset() {
	*x = URLHasValidCertResponse{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *URLHasValidCertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLHasValidCertResponse) ProtoMessage() {}

func (x *URLHasValidCertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLHasValidCertResponse.ProtoReflect.Descriptor instead.
func (*URLHasValidCertResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{1}
}

func (x *URLHasValidCertResponse) GetResult() URLHasValidCertResponse_URLResult {
	if x != nil {
		return x.Result
	}
	return URLHasValidCertResponse_UNSET
}

func (x *URLHasValidCertResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type RandomDataRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Size          int32                  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RandomDataRequest) Reset() {
	*x = RandomDataRequest{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RandomDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomDataRequest) ProtoMessage() {}

func (x *RandomDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomDataRequest.ProtoReflect.Descriptor instead.
func (*RandomDataRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{2}
}

func (x *RandomDataRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type RandomDataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RandomDataResponse) Reset() {
	*x = RandomDataResponse{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RandomDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomDataResponse) ProtoMessage() {}

func (x *RandomDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomDataResponse.ProtoReflect.Descriptor instead.
func (*RandomDataResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{3}
}

func (x *RandomDataResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EnvVarsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnvVars       []string               `protobuf:"bytes,1,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnvVarsResponse) Reset() {
	*x = EnvVarsResponse{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnvVarsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVarsResponse) ProtoMessage() {}

func (x *EnvVarsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvVarsResponse.ProtoReflect.Descriptor instead.
func (*EnvVarsResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{4}
}

func (x *EnvVarsResponse) GetEnvVars() []string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type ReconciliationStatsByClusterResponse struct {
	state         protoimpl.MessageState                                                `protogen:"open.v1"`
	Stats         []*ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReconciliationStatsByClusterResponse) Reset() {
	*x = ReconciliationStatsByClusterResponse{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReconciliationStatsByClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconciliationStatsByClusterResponse) ProtoMessage() {}

func (x *ReconciliationStatsByClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconciliationStatsByClusterResponse.ProtoReflect.Descriptor instead.
func (*ReconciliationStatsByClusterResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{5}
}

func (x *ReconciliationStatsByClusterResponse) GetStats() []*ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster {
	if x != nil {
		return x.Stats
	}
	return nil
}

type ReplicateImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Times         int32                  `protobuf:"varint,2,opt,name=times,proto3" json:"times,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplicateImageRequest) Reset() {
	*x = ReplicateImageRequest{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplicateImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicateImageRequest) ProtoMessage() {}

func (x *ReplicateImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicateImageRequest.ProtoReflect.Descriptor instead.
func (*ReplicateImageRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReplicateImageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReplicateImageRequest) GetTimes() int32 {
	if x != nil {
		return x.Times
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{7}
}

type ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	ClusterId            string                 `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ReconciliationDone   bool                   `protobuf:"varint,2,opt,name=reconciliation_done,json=reconciliationDone,proto3" json:"reconciliation_done,omitempty"`
	DeletedObjectsByType map[string]int32       `protobuf:"bytes,3,rep,name=deleted_objects_by_type,json=deletedObjectsByType,proto3" json:"deleted_objects_by_type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) Reset() {
	*x = ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster{}
	mi := &file_internalapi_central_development_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) ProtoMessage() {}

func (x *ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_development_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster.ProtoReflect.Descriptor instead.
func (*ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) Descriptor() ([]byte, []int) {
	return file_internalapi_central_development_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) GetReconciliationDone() bool {
	if x != nil {
		return x.ReconciliationDone
	}
	return false
}

func (x *ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster) GetDeletedObjectsByType() map[string]int32 {
	if x != nil {
		return x.DeletedObjectsByType
	}
	return nil
}

var File_internalapi_central_development_service_proto protoreflect.FileDescriptor

var file_internalapi_central_development_service_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x16, 0x55, 0x52, 0x4c, 0x48, 0x61, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x65, 0x72, 0x74, 0x50, 0x45, 0x4d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x50, 0x45, 0x4d, 0x22, 0x94, 0x02, 0x0a,
	0x17, 0x55, 0x52, 0x4c, 0x48, 0x61, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x48, 0x61, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x09, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x5f, 0x42,
	0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52,
	0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x53, 0x49,
	0x47, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x42, 0x55, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x5f, 0x47, 0x45, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45,
	0x44, 0x10, 0x04, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x28, 0x0a, 0x12,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x76,
	0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x76,
	0x56, 0x61, 0x72, 0x73, 0x22, 0xe3, 0x03, 0x0a, 0x24, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x46,
	0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x1a, 0xd7, 0x02, 0x0a, 0x1d, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f,
	0x6e, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x65, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x42, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x47, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0xca, 0x04, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x0f, 0x55, 0x52, 0x4c, 0x48,
	0x61, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x48, 0x61, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x48, 0x61, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x66, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x64, 0x65, 0x76, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x55, 0x0a, 0x07, 0x45, 0x6e,
	0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x0e, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e,
	0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x65, 0x6e, 0x76, 0x76, 0x61, 0x72,
	0x73, 0x12, 0x8b, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x2d, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x42, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42,
	0x46, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x3b,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x58, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_internalapi_central_development_service_proto_rawDescOnce sync.Once
	file_internalapi_central_development_service_proto_rawDescData []byte
)

func file_internalapi_central_development_service_proto_rawDescGZIP() []byte {
	file_internalapi_central_development_service_proto_rawDescOnce.Do(func() {
		file_internalapi_central_development_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internalapi_central_development_service_proto_rawDesc), len(file_internalapi_central_development_service_proto_rawDesc)))
	})
	return file_internalapi_central_development_service_proto_rawDescData
}

var file_internalapi_central_development_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internalapi_central_development_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internalapi_central_development_service_proto_goTypes = []any{
	(URLHasValidCertResponse_URLResult)(0),       // 0: central.URLHasValidCertResponse.URLResult
	(*URLHasValidCertRequest)(nil),               // 1: central.URLHasValidCertRequest
	(*URLHasValidCertResponse)(nil),              // 2: central.URLHasValidCertResponse
	(*RandomDataRequest)(nil),                    // 3: central.RandomDataRequest
	(*RandomDataResponse)(nil),                   // 4: central.RandomDataResponse
	(*EnvVarsResponse)(nil),                      // 5: central.EnvVarsResponse
	(*ReconciliationStatsByClusterResponse)(nil), // 6: central.ReconciliationStatsByClusterResponse
	(*ReplicateImageRequest)(nil),                // 7: central.ReplicateImageRequest
	(*Empty)(nil),                                // 8: central.Empty
	(*ReconciliationStatsByClusterResponse_ReconciliationStatsForCluster)(nil), // 9: central.ReconciliationStatsByClusterResponse.ReconciliationStatsForCluster
	nil, // 10: central.ReconciliationStatsByClusterResponse.ReconciliationStatsForCluster.DeletedObjectsByTypeEntry
}
var file_internalapi_central_development_service_proto_depIdxs = []int32{
	0,  // 0: central.URLHasValidCertResponse.result:type_name -> central.URLHasValidCertResponse.URLResult
	9,  // 1: central.ReconciliationStatsByClusterResponse.stats:type_name -> central.ReconciliationStatsByClusterResponse.ReconciliationStatsForCluster
	10, // 2: central.ReconciliationStatsByClusterResponse.ReconciliationStatsForCluster.deleted_objects_by_type:type_name -> central.ReconciliationStatsByClusterResponse.ReconciliationStatsForCluster.DeletedObjectsByTypeEntry
	7,  // 3: central.DevelopmentService.ReplicateImage:input_type -> central.ReplicateImageRequest
	1,  // 4: central.DevelopmentService.URLHasValidCert:input_type -> central.URLHasValidCertRequest
	3,  // 5: central.DevelopmentService.RandomData:input_type -> central.RandomDataRequest
	8,  // 6: central.DevelopmentService.EnvVars:input_type -> central.Empty
	8,  // 7: central.DevelopmentService.ReconciliationStatsByCluster:input_type -> central.Empty
	8,  // 8: central.DevelopmentService.ReplicateImage:output_type -> central.Empty
	2,  // 9: central.DevelopmentService.URLHasValidCert:output_type -> central.URLHasValidCertResponse
	4,  // 10: central.DevelopmentService.RandomData:output_type -> central.RandomDataResponse
	5,  // 11: central.DevelopmentService.EnvVars:output_type -> central.EnvVarsResponse
	6,  // 12: central.DevelopmentService.ReconciliationStatsByCluster:output_type -> central.ReconciliationStatsByClusterResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_internalapi_central_development_service_proto_init() }
func file_internalapi_central_development_service_proto_init() {
	if File_internalapi_central_development_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internalapi_central_development_service_proto_rawDesc), len(file_internalapi_central_development_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_central_development_service_proto_goTypes,
		DependencyIndexes: file_internalapi_central_development_service_proto_depIdxs,
		EnumInfos:         file_internalapi_central_development_service_proto_enumTypes,
		MessageInfos:      file_internalapi_central_development_service_proto_msgTypes,
	}.Build()
	File_internalapi_central_development_service_proto = out.File
	file_internalapi_central_development_service_proto_goTypes = nil
	file_internalapi_central_development_service_proto_depIdxs = nil
}
