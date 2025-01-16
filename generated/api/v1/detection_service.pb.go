// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.3
// source: api/v1/detection_service.proto

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

type BuildDetectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Resource:
	//
	//	*BuildDetectionRequest_Image
	//	*BuildDetectionRequest_ImageName
	Resource           isBuildDetectionRequest_Resource `protobuf_oneof:"Resource"`
	NoExternalMetadata bool                             `protobuf:"varint,2,opt,name=no_external_metadata,json=noExternalMetadata,proto3" json:"no_external_metadata,omitempty"`
	SendNotifications  bool                             `protobuf:"varint,4,opt,name=send_notifications,json=sendNotifications,proto3" json:"send_notifications,omitempty"`
	Force              bool                             `protobuf:"varint,6,opt,name=force,proto3" json:"force,omitempty"`
	PolicyCategories   []string                         `protobuf:"bytes,5,rep,name=policy_categories,json=policyCategories,proto3" json:"policy_categories,omitempty"`
	// Cluster to delegate scan to, may be the cluster's name or ID.
	Cluster       string `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuildDetectionRequest) Reset() {
	*x = BuildDetectionRequest{}
	mi := &file_api_v1_detection_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildDetectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildDetectionRequest) ProtoMessage() {}

func (x *BuildDetectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildDetectionRequest.ProtoReflect.Descriptor instead.
func (*BuildDetectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{0}
}

func (x *BuildDetectionRequest) GetResource() isBuildDetectionRequest_Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *BuildDetectionRequest) GetImage() *storage.ContainerImage {
	if x != nil {
		if x, ok := x.Resource.(*BuildDetectionRequest_Image); ok {
			return x.Image
		}
	}
	return nil
}

func (x *BuildDetectionRequest) GetImageName() string {
	if x != nil {
		if x, ok := x.Resource.(*BuildDetectionRequest_ImageName); ok {
			return x.ImageName
		}
	}
	return ""
}

func (x *BuildDetectionRequest) GetNoExternalMetadata() bool {
	if x != nil {
		return x.NoExternalMetadata
	}
	return false
}

func (x *BuildDetectionRequest) GetSendNotifications() bool {
	if x != nil {
		return x.SendNotifications
	}
	return false
}

func (x *BuildDetectionRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *BuildDetectionRequest) GetPolicyCategories() []string {
	if x != nil {
		return x.PolicyCategories
	}
	return nil
}

func (x *BuildDetectionRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type isBuildDetectionRequest_Resource interface {
	isBuildDetectionRequest_Resource()
}

type BuildDetectionRequest_Image struct {
	Image *storage.ContainerImage `protobuf:"bytes,1,opt,name=image,proto3,oneof"`
}

type BuildDetectionRequest_ImageName struct {
	ImageName string `protobuf:"bytes,3,opt,name=image_name,json=imageName,proto3,oneof"`
}

func (*BuildDetectionRequest_Image) isBuildDetectionRequest_Resource() {}

func (*BuildDetectionRequest_ImageName) isBuildDetectionRequest_Resource() {}

type BuildDetectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alerts        []*storage.Alert       `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuildDetectionResponse) Reset() {
	*x = BuildDetectionResponse{}
	mi := &file_api_v1_detection_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildDetectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildDetectionResponse) ProtoMessage() {}

func (x *BuildDetectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildDetectionResponse.ProtoReflect.Descriptor instead.
func (*BuildDetectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{1}
}

func (x *BuildDetectionResponse) GetAlerts() []*storage.Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type DeployDetectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Resource:
	//
	//	*DeployDetectionRequest_Deployment
	Resource           isDeployDetectionRequest_Resource `protobuf_oneof:"Resource"`
	NoExternalMetadata bool                              `protobuf:"varint,2,opt,name=no_external_metadata,json=noExternalMetadata,proto3" json:"no_external_metadata,omitempty"`
	EnforcementOnly    bool                              `protobuf:"varint,3,opt,name=enforcement_only,json=enforcementOnly,proto3" json:"enforcement_only,omitempty"`
	ClusterId          string                            `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DeployDetectionRequest) Reset() {
	*x = DeployDetectionRequest{}
	mi := &file_api_v1_detection_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployDetectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployDetectionRequest) ProtoMessage() {}

func (x *DeployDetectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployDetectionRequest.ProtoReflect.Descriptor instead.
func (*DeployDetectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeployDetectionRequest) GetResource() isDeployDetectionRequest_Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *DeployDetectionRequest) GetDeployment() *storage.Deployment {
	if x != nil {
		if x, ok := x.Resource.(*DeployDetectionRequest_Deployment); ok {
			return x.Deployment
		}
	}
	return nil
}

func (x *DeployDetectionRequest) GetNoExternalMetadata() bool {
	if x != nil {
		return x.NoExternalMetadata
	}
	return false
}

func (x *DeployDetectionRequest) GetEnforcementOnly() bool {
	if x != nil {
		return x.EnforcementOnly
	}
	return false
}

func (x *DeployDetectionRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type isDeployDetectionRequest_Resource interface {
	isDeployDetectionRequest_Resource()
}

type DeployDetectionRequest_Deployment struct {
	Deployment *storage.Deployment `protobuf:"bytes,1,opt,name=deployment,proto3,oneof"`
}

func (*DeployDetectionRequest_Deployment) isDeployDetectionRequest_Resource() {}

type DeployYAMLDetectionRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Yaml               string                 `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
	NoExternalMetadata bool                   `protobuf:"varint,2,opt,name=no_external_metadata,json=noExternalMetadata,proto3" json:"no_external_metadata,omitempty"`
	EnforcementOnly    bool                   `protobuf:"varint,3,opt,name=enforcement_only,json=enforcementOnly,proto3" json:"enforcement_only,omitempty"`
	Force              bool                   `protobuf:"varint,5,opt,name=force,proto3" json:"force,omitempty"`
	PolicyCategories   []string               `protobuf:"bytes,4,rep,name=policy_categories,json=policyCategories,proto3" json:"policy_categories,omitempty"`
	// Cluster to delegate scan to, may be the cluster's name or ID.
	Cluster       string `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace     string `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeployYAMLDetectionRequest) Reset() {
	*x = DeployYAMLDetectionRequest{}
	mi := &file_api_v1_detection_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployYAMLDetectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployYAMLDetectionRequest) ProtoMessage() {}

func (x *DeployYAMLDetectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployYAMLDetectionRequest.ProtoReflect.Descriptor instead.
func (*DeployYAMLDetectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeployYAMLDetectionRequest) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

func (x *DeployYAMLDetectionRequest) GetNoExternalMetadata() bool {
	if x != nil {
		return x.NoExternalMetadata
	}
	return false
}

func (x *DeployYAMLDetectionRequest) GetEnforcementOnly() bool {
	if x != nil {
		return x.EnforcementOnly
	}
	return false
}

func (x *DeployYAMLDetectionRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *DeployYAMLDetectionRequest) GetPolicyCategories() []string {
	if x != nil {
		return x.PolicyCategories
	}
	return nil
}

func (x *DeployYAMLDetectionRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *DeployYAMLDetectionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type DeployDetectionResponse struct {
	state protoimpl.MessageState         `protogen:"open.v1"`
	Runs  []*DeployDetectionResponse_Run `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
	// The reference will be in the format: namespace/name[<group>/<version>, Kind=<kind>].
	IgnoredObjectRefs []string                 `protobuf:"bytes,2,rep,name=ignored_object_refs,json=ignoredObjectRefs,proto3" json:"ignored_object_refs,omitempty"`
	Remarks           []*DeployDetectionRemark `protobuf:"bytes,3,rep,name=remarks,proto3" json:"remarks,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DeployDetectionResponse) Reset() {
	*x = DeployDetectionResponse{}
	mi := &file_api_v1_detection_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployDetectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployDetectionResponse) ProtoMessage() {}

func (x *DeployDetectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployDetectionResponse.ProtoReflect.Descriptor instead.
func (*DeployDetectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeployDetectionResponse) GetRuns() []*DeployDetectionResponse_Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

func (x *DeployDetectionResponse) GetIgnoredObjectRefs() []string {
	if x != nil {
		return x.IgnoredObjectRefs
	}
	return nil
}

func (x *DeployDetectionResponse) GetRemarks() []*DeployDetectionRemark {
	if x != nil {
		return x.Remarks
	}
	return nil
}

type DeployDetectionRemark struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	Name                   string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PermissionLevel        string                 `protobuf:"bytes,2,opt,name=permission_level,json=permissionLevel,proto3" json:"permission_level,omitempty"`
	AppliedNetworkPolicies []string               `protobuf:"bytes,3,rep,name=applied_network_policies,json=appliedNetworkPolicies,proto3" json:"applied_network_policies,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *DeployDetectionRemark) Reset() {
	*x = DeployDetectionRemark{}
	mi := &file_api_v1_detection_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployDetectionRemark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployDetectionRemark) ProtoMessage() {}

func (x *DeployDetectionRemark) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployDetectionRemark.ProtoReflect.Descriptor instead.
func (*DeployDetectionRemark) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeployDetectionRemark) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeployDetectionRemark) GetPermissionLevel() string {
	if x != nil {
		return x.PermissionLevel
	}
	return ""
}

func (x *DeployDetectionRemark) GetAppliedNetworkPolicies() []string {
	if x != nil {
		return x.AppliedNetworkPolicies
	}
	return nil
}

// This is a helper message for the roxctl JSON report, as jsonpb can only serialize protobuf messages
type ResultAggregation struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Alerts        []*storage.Alert         `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	Remarks       []*DeployDetectionRemark `protobuf:"bytes,2,rep,name=remarks,proto3" json:"remarks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResultAggregation) Reset() {
	*x = ResultAggregation{}
	mi := &file_api_v1_detection_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultAggregation) ProtoMessage() {}

func (x *ResultAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultAggregation.ProtoReflect.Descriptor instead.
func (*ResultAggregation) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{6}
}

func (x *ResultAggregation) GetAlerts() []*storage.Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *ResultAggregation) GetRemarks() []*DeployDetectionRemark {
	if x != nil {
		return x.Remarks
	}
	return nil
}

type DeployDetectionResponse_Run struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Alerts        []*storage.Alert       `protobuf:"bytes,3,rep,name=alerts,proto3" json:"alerts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeployDetectionResponse_Run) Reset() {
	*x = DeployDetectionResponse_Run{}
	mi := &file_api_v1_detection_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployDetectionResponse_Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployDetectionResponse_Run) ProtoMessage() {}

func (x *DeployDetectionResponse_Run) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_detection_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployDetectionResponse_Run.ProtoReflect.Descriptor instead.
func (*DeployDetectionResponse_Run) Descriptor() ([]byte, []int) {
	return file_api_v1_detection_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *DeployDetectionResponse_Run) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeployDetectionResponse_Run) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeployDetectionResponse_Run) GetAlerts() []*storage.Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_api_v1_detection_service_proto protoreflect.FileDescriptor

var file_api_v1_detection_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x15, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x6e, 0x6f, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6e, 0x6f, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x2d, 0x0a, 0x12, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x65, 0x6e,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x16, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0xd7, 0x01, 0x0a, 0x16, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6e,
	0x6f, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6e, 0x6f, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a,
	0x10, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x1a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x59, 0x41,
	0x4d, 0x4c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x6e, 0x6f, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6e, 0x6f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x6e, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x8a,
	0x02, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x75,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x12,
	0x2e, 0x0a, 0x13, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x73, 0x12,
	0x33, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x07, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x73, 0x1a, 0x55, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x15,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x70,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x32, 0xe0, 0x02, 0x0a, 0x10, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x69, 0x0a, 0x10,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x7a, 0x0a, 0x18, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x59,
	0x41, 0x4d, 0x4c, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x59,
	0x41, 0x4d, 0x4c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x79,
	0x61, 0x6d, 0x6c, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a,
	0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_detection_service_proto_rawDescOnce sync.Once
	file_api_v1_detection_service_proto_rawDescData = file_api_v1_detection_service_proto_rawDesc
)

func file_api_v1_detection_service_proto_rawDescGZIP() []byte {
	file_api_v1_detection_service_proto_rawDescOnce.Do(func() {
		file_api_v1_detection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_detection_service_proto_rawDescData)
	})
	return file_api_v1_detection_service_proto_rawDescData
}

var file_api_v1_detection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_detection_service_proto_goTypes = []any{
	(*BuildDetectionRequest)(nil),       // 0: v1.BuildDetectionRequest
	(*BuildDetectionResponse)(nil),      // 1: v1.BuildDetectionResponse
	(*DeployDetectionRequest)(nil),      // 2: v1.DeployDetectionRequest
	(*DeployYAMLDetectionRequest)(nil),  // 3: v1.DeployYAMLDetectionRequest
	(*DeployDetectionResponse)(nil),     // 4: v1.DeployDetectionResponse
	(*DeployDetectionRemark)(nil),       // 5: v1.DeployDetectionRemark
	(*ResultAggregation)(nil),           // 6: v1.ResultAggregation
	(*DeployDetectionResponse_Run)(nil), // 7: v1.DeployDetectionResponse.Run
	(*storage.ContainerImage)(nil),      // 8: storage.ContainerImage
	(*storage.Alert)(nil),               // 9: storage.Alert
	(*storage.Deployment)(nil),          // 10: storage.Deployment
}
var file_api_v1_detection_service_proto_depIdxs = []int32{
	8,  // 0: v1.BuildDetectionRequest.image:type_name -> storage.ContainerImage
	9,  // 1: v1.BuildDetectionResponse.alerts:type_name -> storage.Alert
	10, // 2: v1.DeployDetectionRequest.deployment:type_name -> storage.Deployment
	7,  // 3: v1.DeployDetectionResponse.runs:type_name -> v1.DeployDetectionResponse.Run
	5,  // 4: v1.DeployDetectionResponse.remarks:type_name -> v1.DeployDetectionRemark
	9,  // 5: v1.ResultAggregation.alerts:type_name -> storage.Alert
	5,  // 6: v1.ResultAggregation.remarks:type_name -> v1.DeployDetectionRemark
	9,  // 7: v1.DeployDetectionResponse.Run.alerts:type_name -> storage.Alert
	0,  // 8: v1.DetectionService.DetectBuildTime:input_type -> v1.BuildDetectionRequest
	2,  // 9: v1.DetectionService.DetectDeployTime:input_type -> v1.DeployDetectionRequest
	3,  // 10: v1.DetectionService.DetectDeployTimeFromYAML:input_type -> v1.DeployYAMLDetectionRequest
	1,  // 11: v1.DetectionService.DetectBuildTime:output_type -> v1.BuildDetectionResponse
	4,  // 12: v1.DetectionService.DetectDeployTime:output_type -> v1.DeployDetectionResponse
	4,  // 13: v1.DetectionService.DetectDeployTimeFromYAML:output_type -> v1.DeployDetectionResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1_detection_service_proto_init() }
func file_api_v1_detection_service_proto_init() {
	if File_api_v1_detection_service_proto != nil {
		return
	}
	file_api_v1_detection_service_proto_msgTypes[0].OneofWrappers = []any{
		(*BuildDetectionRequest_Image)(nil),
		(*BuildDetectionRequest_ImageName)(nil),
	}
	file_api_v1_detection_service_proto_msgTypes[2].OneofWrappers = []any{
		(*DeployDetectionRequest_Deployment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_detection_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_detection_service_proto_goTypes,
		DependencyIndexes: file_api_v1_detection_service_proto_depIdxs,
		MessageInfos:      file_api_v1_detection_service_proto_msgTypes,
	}.Build()
	File_api_v1_detection_service_proto = out.File
	file_api_v1_detection_service_proto_rawDesc = nil
	file_api_v1_detection_service_proto_goTypes = nil
	file_api_v1_detection_service_proto_depIdxs = nil
}
