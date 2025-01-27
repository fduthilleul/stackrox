// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: internalapi/central/sensor_upgrade_ctrl_iservice.proto

package central

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type UpgradeCheckInFromUpgraderRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ClusterId         string                 `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	UpgradeProcessId  string                 `protobuf:"bytes,1,opt,name=upgrade_process_id,json=upgradeProcessId,proto3" json:"upgrade_process_id,omitempty"`
	CurrentWorkflow   string                 `protobuf:"bytes,2,opt,name=current_workflow,json=currentWorkflow,proto3" json:"current_workflow,omitempty"`
	LastExecutedStage string                 `protobuf:"bytes,3,opt,name=last_executed_stage,json=lastExecutedStage,proto3" json:"last_executed_stage,omitempty"`
	// The error from the last executed stage, if any.
	// If this is empty, that implies that the last stage
	// was succesfully executed.
	LastExecutedStageError string `protobuf:"bytes,4,opt,name=last_executed_stage_error,json=lastExecutedStageError,proto3" json:"last_executed_stage_error,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *UpgradeCheckInFromUpgraderRequest) Reset() {
	*x = UpgradeCheckInFromUpgraderRequest{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInFromUpgraderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInFromUpgraderRequest) ProtoMessage() {}

func (x *UpgradeCheckInFromUpgraderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInFromUpgraderRequest.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInFromUpgraderRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{0}
}

func (x *UpgradeCheckInFromUpgraderRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpgradeCheckInFromUpgraderRequest) GetUpgradeProcessId() string {
	if x != nil {
		return x.UpgradeProcessId
	}
	return ""
}

func (x *UpgradeCheckInFromUpgraderRequest) GetCurrentWorkflow() string {
	if x != nil {
		return x.CurrentWorkflow
	}
	return ""
}

func (x *UpgradeCheckInFromUpgraderRequest) GetLastExecutedStage() string {
	if x != nil {
		return x.LastExecutedStage
	}
	return ""
}

func (x *UpgradeCheckInFromUpgraderRequest) GetLastExecutedStageError() string {
	if x != nil {
		return x.LastExecutedStageError
	}
	return ""
}

type UpgradeCheckInFromUpgraderResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	WorkflowToExecute string                 `protobuf:"bytes,1,opt,name=workflow_to_execute,json=workflowToExecute,proto3" json:"workflow_to_execute,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UpgradeCheckInFromUpgraderResponse) Reset() {
	*x = UpgradeCheckInFromUpgraderResponse{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInFromUpgraderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInFromUpgraderResponse) ProtoMessage() {}

func (x *UpgradeCheckInFromUpgraderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInFromUpgraderResponse.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInFromUpgraderResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{1}
}

func (x *UpgradeCheckInFromUpgraderResponse) GetWorkflowToExecute() string {
	if x != nil {
		return x.WorkflowToExecute
	}
	return ""
}

type UpgradeCheckInFromSensorRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ClusterId        string                 `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	UpgradeProcessId string                 `protobuf:"bytes,1,opt,name=upgrade_process_id,json=upgradeProcessId,proto3" json:"upgrade_process_id,omitempty"`
	// Types that are valid to be assigned to State:
	//
	//	*UpgradeCheckInFromSensorRequest_LaunchError
	//	*UpgradeCheckInFromSensorRequest_PodStates
	//	*UpgradeCheckInFromSensorRequest_DeploymentGone
	State         isUpgradeCheckInFromSensorRequest_State `protobuf_oneof:"state"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeCheckInFromSensorRequest) Reset() {
	*x = UpgradeCheckInFromSensorRequest{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInFromSensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInFromSensorRequest) ProtoMessage() {}

func (x *UpgradeCheckInFromSensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInFromSensorRequest.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInFromSensorRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{2}
}

func (x *UpgradeCheckInFromSensorRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpgradeCheckInFromSensorRequest) GetUpgradeProcessId() string {
	if x != nil {
		return x.UpgradeProcessId
	}
	return ""
}

func (x *UpgradeCheckInFromSensorRequest) GetState() isUpgradeCheckInFromSensorRequest_State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *UpgradeCheckInFromSensorRequest) GetLaunchError() string {
	if x != nil {
		if x, ok := x.State.(*UpgradeCheckInFromSensorRequest_LaunchError); ok {
			return x.LaunchError
		}
	}
	return ""
}

func (x *UpgradeCheckInFromSensorRequest) GetPodStates() *UpgradeCheckInFromSensorRequest_UpgraderPodStates {
	if x != nil {
		if x, ok := x.State.(*UpgradeCheckInFromSensorRequest_PodStates); ok {
			return x.PodStates
		}
	}
	return nil
}

func (x *UpgradeCheckInFromSensorRequest) GetDeploymentGone() bool {
	if x != nil {
		if x, ok := x.State.(*UpgradeCheckInFromSensorRequest_DeploymentGone); ok {
			return x.DeploymentGone
		}
	}
	return false
}

type isUpgradeCheckInFromSensorRequest_State interface {
	isUpgradeCheckInFromSensorRequest_State()
}

type UpgradeCheckInFromSensorRequest_LaunchError struct {
	LaunchError string `protobuf:"bytes,2,opt,name=launch_error,json=launchError,proto3,oneof"`
}

type UpgradeCheckInFromSensorRequest_PodStates struct {
	PodStates *UpgradeCheckInFromSensorRequest_UpgraderPodStates `protobuf:"bytes,3,opt,name=pod_states,json=podStates,proto3,oneof"`
}

type UpgradeCheckInFromSensorRequest_DeploymentGone struct {
	DeploymentGone bool `protobuf:"varint,4,opt,name=deployment_gone,json=deploymentGone,proto3,oneof"`
}

func (*UpgradeCheckInFromSensorRequest_LaunchError) isUpgradeCheckInFromSensorRequest_State() {}

func (*UpgradeCheckInFromSensorRequest_PodStates) isUpgradeCheckInFromSensorRequest_State() {}

func (*UpgradeCheckInFromSensorRequest_DeploymentGone) isUpgradeCheckInFromSensorRequest_State() {}

// UpgradeCheckInResponseDetails contains proto messages that are added to details
// when returning errors from these endpoints.
type UpgradeCheckInResponseDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeCheckInResponseDetails) Reset() {
	*x = UpgradeCheckInResponseDetails{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInResponseDetails) ProtoMessage() {}

func (x *UpgradeCheckInResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInResponseDetails.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInResponseDetails) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{3}
}

type UpgradeCheckInFromSensorRequest_PodErrorCondition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ImageRelated  bool                   `protobuf:"varint,2,opt,name=image_related,json=imageRelated,proto3" json:"image_related,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeCheckInFromSensorRequest_PodErrorCondition) Reset() {
	*x = UpgradeCheckInFromSensorRequest_PodErrorCondition{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInFromSensorRequest_PodErrorCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInFromSensorRequest_PodErrorCondition) ProtoMessage() {}

func (x *UpgradeCheckInFromSensorRequest_PodErrorCondition) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInFromSensorRequest_PodErrorCondition.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInFromSensorRequest_PodErrorCondition) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{2, 0}
}

func (x *UpgradeCheckInFromSensorRequest_PodErrorCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpgradeCheckInFromSensorRequest_PodErrorCondition) GetImageRelated() bool {
	if x != nil {
		return x.ImageRelated
	}
	return false
}

type UpgradeCheckInFromSensorRequest_UpgraderPodState struct {
	state         protoimpl.MessageState                             `protogen:"open.v1"`
	PodName       string                                             `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Started       bool                                               `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	Error         *UpgradeCheckInFromSensorRequest_PodErrorCondition `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodState) Reset() {
	*x = UpgradeCheckInFromSensorRequest_UpgraderPodState{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInFromSensorRequest_UpgraderPodState) ProtoMessage() {}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodState) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInFromSensorRequest_UpgraderPodState.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInFromSensorRequest_UpgraderPodState) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{2, 1}
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodState) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodState) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodState) GetError() *UpgradeCheckInFromSensorRequest_PodErrorCondition {
	if x != nil {
		return x.Error
	}
	return nil
}

type UpgradeCheckInFromSensorRequest_UpgraderPodStates struct {
	state         protoimpl.MessageState                              `protogen:"open.v1"`
	States        []*UpgradeCheckInFromSensorRequest_UpgraderPodState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodStates) Reset() {
	*x = UpgradeCheckInFromSensorRequest_UpgraderPodStates{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodStates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInFromSensorRequest_UpgraderPodStates) ProtoMessage() {}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodStates) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInFromSensorRequest_UpgraderPodStates.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInFromSensorRequest_UpgraderPodStates) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{2, 2}
}

func (x *UpgradeCheckInFromSensorRequest_UpgraderPodStates) GetStates() []*UpgradeCheckInFromSensorRequest_UpgraderPodState {
	if x != nil {
		return x.States
	}
	return nil
}

type UpgradeCheckInResponseDetails_NoUpgradeInProgress struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeCheckInResponseDetails_NoUpgradeInProgress) Reset() {
	*x = UpgradeCheckInResponseDetails_NoUpgradeInProgress{}
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeCheckInResponseDetails_NoUpgradeInProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeCheckInResponseDetails_NoUpgradeInProgress) ProtoMessage() {}

func (x *UpgradeCheckInResponseDetails_NoUpgradeInProgress) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeCheckInResponseDetails_NoUpgradeInProgress.ProtoReflect.Descriptor instead.
func (*UpgradeCheckInResponseDetails_NoUpgradeInProgress) Descriptor() ([]byte, []int) {
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP(), []int{3, 0}
}

var File_internalapi_central_sensor_upgrade_ctrl_iservice_proto protoreflect.FileDescriptor

var file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDesc = string([]byte{
	0x0a, 0x36, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x75, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x74, 0x72, 0x6c, 0x5f, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86,
	0x02, 0x0a, 0x21, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x2e, 0x0a, 0x13,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x19,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x16, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x22, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x22, 0xfc, 0x04,
	0x0a, 0x1f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x48, 0x00, 0x52, 0x09, 0x70, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x29, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x6f, 0x6e, 0x65, 0x1a, 0x52, 0x0a, 0x11, 0x50,
	0x6f, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x1a,
	0x99, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x6f, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x66, 0x0a, 0x11, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x51, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x1d,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x15, 0x0a,
	0x13, 0x4e, 0x6f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x32, 0xf2, 0x01, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x1a, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x2a, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x18, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x28, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x3b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescOnce sync.Once
	file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescData []byte
)

func file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDesc), len(file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDesc)))
	})
	return file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDescData
}

var file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_goTypes = []any{
	(*UpgradeCheckInFromUpgraderRequest)(nil),                 // 0: central.UpgradeCheckInFromUpgraderRequest
	(*UpgradeCheckInFromUpgraderResponse)(nil),                // 1: central.UpgradeCheckInFromUpgraderResponse
	(*UpgradeCheckInFromSensorRequest)(nil),                   // 2: central.UpgradeCheckInFromSensorRequest
	(*UpgradeCheckInResponseDetails)(nil),                     // 3: central.UpgradeCheckInResponseDetails
	(*UpgradeCheckInFromSensorRequest_PodErrorCondition)(nil), // 4: central.UpgradeCheckInFromSensorRequest.PodErrorCondition
	(*UpgradeCheckInFromSensorRequest_UpgraderPodState)(nil),  // 5: central.UpgradeCheckInFromSensorRequest.UpgraderPodState
	(*UpgradeCheckInFromSensorRequest_UpgraderPodStates)(nil), // 6: central.UpgradeCheckInFromSensorRequest.UpgraderPodStates
	(*UpgradeCheckInResponseDetails_NoUpgradeInProgress)(nil), // 7: central.UpgradeCheckInResponseDetails.NoUpgradeInProgress
	(*emptypb.Empty)(nil),                                     // 8: google.protobuf.Empty
}
var file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_depIdxs = []int32{
	6, // 0: central.UpgradeCheckInFromSensorRequest.pod_states:type_name -> central.UpgradeCheckInFromSensorRequest.UpgraderPodStates
	4, // 1: central.UpgradeCheckInFromSensorRequest.UpgraderPodState.error:type_name -> central.UpgradeCheckInFromSensorRequest.PodErrorCondition
	5, // 2: central.UpgradeCheckInFromSensorRequest.UpgraderPodStates.states:type_name -> central.UpgradeCheckInFromSensorRequest.UpgraderPodState
	0, // 3: central.SensorUpgradeControlService.UpgradeCheckInFromUpgrader:input_type -> central.UpgradeCheckInFromUpgraderRequest
	2, // 4: central.SensorUpgradeControlService.UpgradeCheckInFromSensor:input_type -> central.UpgradeCheckInFromSensorRequest
	1, // 5: central.SensorUpgradeControlService.UpgradeCheckInFromUpgrader:output_type -> central.UpgradeCheckInFromUpgraderResponse
	8, // 6: central.SensorUpgradeControlService.UpgradeCheckInFromSensor:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_init() }
func file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_init() {
	if File_internalapi_central_sensor_upgrade_ctrl_iservice_proto != nil {
		return
	}
	file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes[2].OneofWrappers = []any{
		(*UpgradeCheckInFromSensorRequest_LaunchError)(nil),
		(*UpgradeCheckInFromSensorRequest_PodStates)(nil),
		(*UpgradeCheckInFromSensorRequest_DeploymentGone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDesc), len(file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_central_sensor_upgrade_ctrl_iservice_proto = out.File
	file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_goTypes = nil
	file_internalapi_central_sensor_upgrade_ctrl_iservice_proto_depIdxs = nil
}
