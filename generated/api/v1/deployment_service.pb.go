// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.3
// source: api/v1/deployment_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
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

type DeploymentLabelsResponse struct {
	state         protoimpl.MessageState                           `protogen:"open.v1"`
	Labels        map[string]*DeploymentLabelsResponse_LabelValues `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values        []string                                         `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentLabelsResponse) Reset() {
	*x = DeploymentLabelsResponse{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentLabelsResponse) ProtoMessage() {}

func (x *DeploymentLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentLabelsResponse.ProtoReflect.Descriptor instead.
func (*DeploymentLabelsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentLabelsResponse) GetLabels() map[string]*DeploymentLabelsResponse_LabelValues {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DeploymentLabelsResponse) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type ListDeploymentsResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Deployments   []*storage.ListDeployment `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDeploymentsResponse) Reset() {
	*x = ListDeploymentsResponse{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeploymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsResponse) ProtoMessage() {}

func (x *ListDeploymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsResponse.ProtoReflect.Descriptor instead.
func (*ListDeploymentsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListDeploymentsResponse) GetDeployments() []*storage.ListDeployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type CountDeploymentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountDeploymentsResponse) Reset() {
	*x = CountDeploymentsResponse{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountDeploymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountDeploymentsResponse) ProtoMessage() {}

func (x *CountDeploymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountDeploymentsResponse.ProtoReflect.Descriptor instead.
func (*CountDeploymentsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{2}
}

func (x *CountDeploymentsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListDeploymentsWithProcessInfoResponse struct {
	state         protoimpl.MessageState                                              `protogen:"open.v1"`
	Deployments   []*ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDeploymentsWithProcessInfoResponse) Reset() {
	*x = ListDeploymentsWithProcessInfoResponse{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeploymentsWithProcessInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsWithProcessInfoResponse) ProtoMessage() {}

func (x *ListDeploymentsWithProcessInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsWithProcessInfoResponse.ProtoReflect.Descriptor instead.
func (*ListDeploymentsWithProcessInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListDeploymentsWithProcessInfoResponse) GetDeployments() []*ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type GetDeploymentWithRiskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Deployment    *storage.Deployment    `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Risk          *storage.Risk          `protobuf:"bytes,2,opt,name=risk,proto3" json:"risk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDeploymentWithRiskResponse) Reset() {
	*x = GetDeploymentWithRiskResponse{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDeploymentWithRiskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentWithRiskResponse) ProtoMessage() {}

func (x *GetDeploymentWithRiskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentWithRiskResponse.ProtoReflect.Descriptor instead.
func (*GetDeploymentWithRiskResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeploymentWithRiskResponse) GetDeployment() *storage.Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *GetDeploymentWithRiskResponse) GetRisk() *storage.Risk {
	if x != nil {
		return x.Risk
	}
	return nil
}

type ExportDeploymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timeout       int32                  `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Query         string                 `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportDeploymentRequest) Reset() {
	*x = ExportDeploymentRequest{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDeploymentRequest) ProtoMessage() {}

func (x *ExportDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDeploymentRequest.ProtoReflect.Descriptor instead.
func (*ExportDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{5}
}

func (x *ExportDeploymentRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *ExportDeploymentRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type ExportDeploymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Deployment    *storage.Deployment    `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportDeploymentResponse) Reset() {
	*x = ExportDeploymentResponse{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDeploymentResponse) ProtoMessage() {}

func (x *ExportDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDeploymentResponse.ProtoReflect.Descriptor instead.
func (*ExportDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{6}
}

func (x *ExportDeploymentResponse) GetDeployment() *storage.Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

type DeploymentLabelsResponse_LabelValues struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []string               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentLabelsResponse_LabelValues) Reset() {
	*x = DeploymentLabelsResponse_LabelValues{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentLabelsResponse_LabelValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentLabelsResponse_LabelValues) ProtoMessage() {}

func (x *DeploymentLabelsResponse_LabelValues) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentLabelsResponse_LabelValues.ProtoReflect.Descriptor instead.
func (*DeploymentLabelsResponse_LabelValues) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DeploymentLabelsResponse_LabelValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo struct {
	state            protoimpl.MessageState                    `protogen:"open.v1"`
	Deployment       *storage.ListDeployment                   `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	BaselineStatuses []*storage.ContainerNameAndBaselineStatus `protobuf:"bytes,3,rep,name=baseline_statuses,json=baselineStatuses,proto3" json:"baseline_statuses,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) Reset() {
	*x = ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo{}
	mi := &file_api_v1_deployment_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) ProtoMessage() {}

func (x *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_deployment_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo.ProtoReflect.Descriptor instead.
func (*ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) Descriptor() ([]byte, []int) {
	return file_api_v1_deployment_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) GetDeployment() *storage.ListDeployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo) GetBaselineStatuses() []*storage.ContainerNameAndBaselineStatus {
	if x != nil {
		return x.BaselineStatuses
	}
	return nil
}

var File_api_v1_deployment_service_proto protoreflect.FileDescriptor

var file_api_v1_deployment_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x69, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x18, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x25,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x63, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x54, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x30, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xc3, 0x02, 0x0a, 0x26, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x44, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xb0, 0x01, 0x0a, 0x19, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x11,
	0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x41,
	0x6e, 0x64, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x77, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x69, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x04, 0x72, 0x69, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x52, 0x04, 0x72, 0x69, 0x73,
	0x6b, 0x22, 0x49, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x18,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xe8, 0x05,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x69,
	0x73, 0x6b, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x69, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x77, 0x69, 0x74, 0x68, 0x72, 0x69, 0x73, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a,
	0x10, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1b, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x2a, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x77, 0x69, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x5d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x70, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x30, 0x01, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x58, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_deployment_service_proto_rawDescOnce sync.Once
	file_api_v1_deployment_service_proto_rawDescData = file_api_v1_deployment_service_proto_rawDesc
)

func file_api_v1_deployment_service_proto_rawDescGZIP() []byte {
	file_api_v1_deployment_service_proto_rawDescOnce.Do(func() {
		file_api_v1_deployment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_deployment_service_proto_rawDescData)
	})
	return file_api_v1_deployment_service_proto_rawDescData
}

var file_api_v1_deployment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1_deployment_service_proto_goTypes = []any{
	(*DeploymentLabelsResponse)(nil),               // 0: v1.DeploymentLabelsResponse
	(*ListDeploymentsResponse)(nil),                // 1: v1.ListDeploymentsResponse
	(*CountDeploymentsResponse)(nil),               // 2: v1.CountDeploymentsResponse
	(*ListDeploymentsWithProcessInfoResponse)(nil), // 3: v1.ListDeploymentsWithProcessInfoResponse
	(*GetDeploymentWithRiskResponse)(nil),          // 4: v1.GetDeploymentWithRiskResponse
	(*ExportDeploymentRequest)(nil),                // 5: v1.ExportDeploymentRequest
	(*ExportDeploymentResponse)(nil),               // 6: v1.ExportDeploymentResponse
	(*DeploymentLabelsResponse_LabelValues)(nil),   // 7: v1.DeploymentLabelsResponse.LabelValues
	nil, // 8: v1.DeploymentLabelsResponse.LabelsEntry
	(*ListDeploymentsWithProcessInfoResponse_DeploymentWithProcessInfo)(nil), // 9: v1.ListDeploymentsWithProcessInfoResponse.DeploymentWithProcessInfo
	(*storage.ListDeployment)(nil),                                           // 10: storage.ListDeployment
	(*storage.Deployment)(nil),                                               // 11: storage.Deployment
	(*storage.Risk)(nil),                                                     // 12: storage.Risk
	(*storage.ContainerNameAndBaselineStatus)(nil),                           // 13: storage.ContainerNameAndBaselineStatus
	(*ResourceByID)(nil),                                                     // 14: v1.ResourceByID
	(*RawQuery)(nil),                                                         // 15: v1.RawQuery
	(*Empty)(nil),                                                            // 16: v1.Empty
}
var file_api_v1_deployment_service_proto_depIdxs = []int32{
	8,  // 0: v1.DeploymentLabelsResponse.labels:type_name -> v1.DeploymentLabelsResponse.LabelsEntry
	10, // 1: v1.ListDeploymentsResponse.deployments:type_name -> storage.ListDeployment
	9,  // 2: v1.ListDeploymentsWithProcessInfoResponse.deployments:type_name -> v1.ListDeploymentsWithProcessInfoResponse.DeploymentWithProcessInfo
	11, // 3: v1.GetDeploymentWithRiskResponse.deployment:type_name -> storage.Deployment
	12, // 4: v1.GetDeploymentWithRiskResponse.risk:type_name -> storage.Risk
	11, // 5: v1.ExportDeploymentResponse.deployment:type_name -> storage.Deployment
	7,  // 6: v1.DeploymentLabelsResponse.LabelsEntry.value:type_name -> v1.DeploymentLabelsResponse.LabelValues
	10, // 7: v1.ListDeploymentsWithProcessInfoResponse.DeploymentWithProcessInfo.deployment:type_name -> storage.ListDeployment
	13, // 8: v1.ListDeploymentsWithProcessInfoResponse.DeploymentWithProcessInfo.baseline_statuses:type_name -> storage.ContainerNameAndBaselineStatus
	14, // 9: v1.DeploymentService.GetDeployment:input_type -> v1.ResourceByID
	14, // 10: v1.DeploymentService.GetDeploymentWithRisk:input_type -> v1.ResourceByID
	15, // 11: v1.DeploymentService.CountDeployments:input_type -> v1.RawQuery
	15, // 12: v1.DeploymentService.ListDeployments:input_type -> v1.RawQuery
	15, // 13: v1.DeploymentService.ListDeploymentsWithProcessInfo:input_type -> v1.RawQuery
	16, // 14: v1.DeploymentService.GetLabels:input_type -> v1.Empty
	5,  // 15: v1.DeploymentService.ExportDeployments:input_type -> v1.ExportDeploymentRequest
	11, // 16: v1.DeploymentService.GetDeployment:output_type -> storage.Deployment
	4,  // 17: v1.DeploymentService.GetDeploymentWithRisk:output_type -> v1.GetDeploymentWithRiskResponse
	2,  // 18: v1.DeploymentService.CountDeployments:output_type -> v1.CountDeploymentsResponse
	1,  // 19: v1.DeploymentService.ListDeployments:output_type -> v1.ListDeploymentsResponse
	3,  // 20: v1.DeploymentService.ListDeploymentsWithProcessInfo:output_type -> v1.ListDeploymentsWithProcessInfoResponse
	0,  // 21: v1.DeploymentService.GetLabels:output_type -> v1.DeploymentLabelsResponse
	6,  // 22: v1.DeploymentService.ExportDeployments:output_type -> v1.ExportDeploymentResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1_deployment_service_proto_init() }
func file_api_v1_deployment_service_proto_init() {
	if File_api_v1_deployment_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_deployment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_deployment_service_proto_goTypes,
		DependencyIndexes: file_api_v1_deployment_service_proto_depIdxs,
		MessageInfos:      file_api_v1_deployment_service_proto_msgTypes,
	}.Build()
	File_api_v1_deployment_service_proto = out.File
	file_api_v1_deployment_service_proto_rawDesc = nil
	file_api_v1_deployment_service_proto_goTypes = nil
	file_api_v1_deployment_service_proto_depIdxs = nil
}
