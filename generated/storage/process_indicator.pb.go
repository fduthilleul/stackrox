// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: storage/process_indicator.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Next available tag: 13
type ProcessIndicator struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique UUID for the Indicator message
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Process ID,store,hidden" sql:"pk,type(uuid)"`                                            // @gotags: search:"Process ID,store,hidden"  sql:"pk,type(uuid)"
	DeploymentId  string `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty" search:"Deployment ID,store,hidden" policy:",prefer-parent" sql:"index=hash,fk(Deployment:id),no-fk-constraint,type(uuid)"`    // @gotags: search:"Deployment ID,store,hidden" policy:",prefer-parent" sql:"index=hash,fk(Deployment:id),no-fk-constraint,type(uuid)"
	ContainerName string `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty" search:"Container Name,hidden" policy:",prefer-parent"` // @gotags: search:"Container Name,hidden" policy:",prefer-parent"
	// Pod name
	PodId  string `protobuf:"bytes,4,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty" search:"Pod ID,hidden"`     // @gotags: search:"Pod ID,hidden"
	PodUid string `protobuf:"bytes,11,opt,name=pod_uid,json=podUid,proto3" json:"pod_uid,omitempty" search:"Pod UID,hidden" sql:"index=hash,type(uuid)"` // @gotags: search:"Pod UID,hidden" sql:"index=hash,type(uuid)"
	// A process signal message passed from Collector to Sensor
	Signal             *ProcessSignal         `protobuf:"bytes,6,opt,name=signal,proto3" json:"signal,omitempty"`
	ClusterId          string                 `protobuf:"bytes,7,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden" policy:",prefer-parent" sql:"type(uuid)"` // @gotags: search:"Cluster ID,store,hidden" policy:",prefer-parent" sql:"type(uuid)"
	Namespace          string                 `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store,hidden" policy:",prefer-parent"`                  // @gotags: search:"Namespace,store,hidden" policy:",prefer-parent"
	ContainerStartTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=container_start_time,json=containerStartTime,proto3" json:"container_start_time,omitempty"`
	ImageId            string                 `protobuf:"bytes,12,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ProcessIndicator) Reset() {
	*x = ProcessIndicator{}
	mi := &file_storage_process_indicator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessIndicator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessIndicator) ProtoMessage() {}

func (x *ProcessIndicator) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_indicator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessIndicator.ProtoReflect.Descriptor instead.
func (*ProcessIndicator) Descriptor() ([]byte, []int) {
	return file_storage_process_indicator_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessIndicator) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcessIndicator) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ProcessIndicator) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ProcessIndicator) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *ProcessIndicator) GetPodUid() string {
	if x != nil {
		return x.PodUid
	}
	return ""
}

func (x *ProcessIndicator) GetSignal() *ProcessSignal {
	if x != nil {
		return x.Signal
	}
	return nil
}

func (x *ProcessIndicator) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ProcessIndicator) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ProcessIndicator) GetContainerStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ContainerStartTime
	}
	return nil
}

func (x *ProcessIndicator) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

// This is the unique key we store process indicators under in Bolt.
// This is stored in the DB, so please follow proto compatibility rules for it,
// OR discard existing values and repopulate it on startup.
type ProcessIndicatorUniqueKey struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	PodId               string                 `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	ContainerName       string                 `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	ProcessName         string                 `protobuf:"bytes,3,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	ProcessExecFilePath string                 `protobuf:"bytes,4,opt,name=process_exec_file_path,json=processExecFilePath,proto3" json:"process_exec_file_path,omitempty"`
	ProcessArgs         string                 `protobuf:"bytes,5,opt,name=process_args,json=processArgs,proto3" json:"process_args,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ProcessIndicatorUniqueKey) Reset() {
	*x = ProcessIndicatorUniqueKey{}
	mi := &file_storage_process_indicator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessIndicatorUniqueKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessIndicatorUniqueKey) ProtoMessage() {}

func (x *ProcessIndicatorUniqueKey) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_indicator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessIndicatorUniqueKey.ProtoReflect.Descriptor instead.
func (*ProcessIndicatorUniqueKey) Descriptor() ([]byte, []int) {
	return file_storage_process_indicator_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessIndicatorUniqueKey) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *ProcessIndicatorUniqueKey) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ProcessIndicatorUniqueKey) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *ProcessIndicatorUniqueKey) GetProcessExecFilePath() string {
	if x != nil {
		return x.ProcessExecFilePath
	}
	return ""
}

func (x *ProcessIndicatorUniqueKey) GetProcessArgs() string {
	if x != nil {
		return x.ProcessArgs
	}
	return ""
}

// This is the processes information which is added to endpoint data
type NetworkProcessUniqueKey struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	ProcessName         string                 `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	ProcessExecFilePath string                 `protobuf:"bytes,2,opt,name=process_exec_file_path,json=processExecFilePath,proto3" json:"process_exec_file_path,omitempty"`
	ProcessArgs         string                 `protobuf:"bytes,3,opt,name=process_args,json=processArgs,proto3" json:"process_args,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *NetworkProcessUniqueKey) Reset() {
	*x = NetworkProcessUniqueKey{}
	mi := &file_storage_process_indicator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkProcessUniqueKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkProcessUniqueKey) ProtoMessage() {}

func (x *NetworkProcessUniqueKey) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_indicator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkProcessUniqueKey.ProtoReflect.Descriptor instead.
func (*NetworkProcessUniqueKey) Descriptor() ([]byte, []int) {
	return file_storage_process_indicator_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkProcessUniqueKey) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *NetworkProcessUniqueKey) GetProcessExecFilePath() string {
	if x != nil {
		return x.ProcessExecFilePath
	}
	return ""
}

func (x *NetworkProcessUniqueKey) GetProcessArgs() string {
	if x != nil {
		return x.ProcessArgs
	}
	return ""
}

type ProcessSignal struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique UUID for identifying the message
	// We have this here instead of at the top level
	// because we want to have each message to be
	// self contained.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of container associated with this process
	ContainerId string `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty" search:"Container ID,hidden"` // @gotags: search:"Container ID,hidden"
	// Process creation time
	Time *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty" search:"Process Creation Time,hidden" sql:"index=btree"` // @gotags: search:"Process Creation Time,hidden" sql:"index=btree"
	// Process name
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" search:"Process Name"` // @gotags: search:"Process Name"
	// Process arguments
	Args string `protobuf:"bytes,5,opt,name=args,proto3" json:"args,omitempty" search:"Process Arguments"` // @gotags: search:"Process Arguments"
	// Process executable file path
	ExecFilePath string `protobuf:"bytes,6,opt,name=exec_file_path,json=execFilePath,proto3" json:"exec_file_path,omitempty" search:"Process Path"` // @gotags: search:"Process Path"
	// Host process ID
	Pid uint32 `protobuf:"varint,7,opt,name=pid,proto3" json:"pid,omitempty"`
	// Real user ID
	Uid uint32 `protobuf:"varint,8,opt,name=uid,proto3" json:"uid,omitempty" search:"Process UID"` // @gotags: search:"Process UID"
	// Real group ID
	Gid uint32 `protobuf:"varint,9,opt,name=gid,proto3" json:"gid,omitempty"`
	// Process Lineage
	//
	// Deprecated: Marked as deprecated in storage/process_indicator.proto.
	Lineage []string `protobuf:"bytes,10,rep,name=lineage,proto3" json:"lineage,omitempty"`
	// Signal origin
	Scraped bool `protobuf:"varint,11,opt,name=scraped,proto3" json:"scraped,omitempty"`
	// Process LineageInfo
	LineageInfo   []*ProcessSignal_LineageInfo `protobuf:"bytes,12,rep,name=lineage_info,json=lineageInfo,proto3" json:"lineage_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessSignal) Reset() {
	*x = ProcessSignal{}
	mi := &file_storage_process_indicator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSignal) ProtoMessage() {}

func (x *ProcessSignal) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_indicator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSignal.ProtoReflect.Descriptor instead.
func (*ProcessSignal) Descriptor() ([]byte, []int) {
	return file_storage_process_indicator_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessSignal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcessSignal) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ProcessSignal) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ProcessSignal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessSignal) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *ProcessSignal) GetExecFilePath() string {
	if x != nil {
		return x.ExecFilePath
	}
	return ""
}

func (x *ProcessSignal) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProcessSignal) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ProcessSignal) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

// Deprecated: Marked as deprecated in storage/process_indicator.proto.
func (x *ProcessSignal) GetLineage() []string {
	if x != nil {
		return x.Lineage
	}
	return nil
}

func (x *ProcessSignal) GetScraped() bool {
	if x != nil {
		return x.Scraped
	}
	return false
}

func (x *ProcessSignal) GetLineageInfo() []*ProcessSignal_LineageInfo {
	if x != nil {
		return x.LineageInfo
	}
	return nil
}

type ProcessSignal_LineageInfo struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	ParentUid          uint32                 `protobuf:"varint,1,opt,name=parent_uid,json=parentUid,proto3" json:"parent_uid,omitempty"`
	ParentExecFilePath string                 `protobuf:"bytes,2,opt,name=parent_exec_file_path,json=parentExecFilePath,proto3" json:"parent_exec_file_path,omitempty" policy:"Process Ancestor"` // @gotags: policy:"Process Ancestor"
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ProcessSignal_LineageInfo) Reset() {
	*x = ProcessSignal_LineageInfo{}
	mi := &file_storage_process_indicator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessSignal_LineageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSignal_LineageInfo) ProtoMessage() {}

func (x *ProcessSignal_LineageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_storage_process_indicator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSignal_LineageInfo.ProtoReflect.Descriptor instead.
func (*ProcessSignal_LineageInfo) Descriptor() ([]byte, []int) {
	return file_storage_process_indicator_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ProcessSignal_LineageInfo) GetParentUid() uint32 {
	if x != nil {
		return x.ParentUid
	}
	return 0
}

func (x *ProcessSignal_LineageInfo) GetParentExecFilePath() string {
	if x != nil {
		return x.ParentExecFilePath
	}
	return ""
}

var File_storage_process_indicator_proto protoreflect.FileDescriptor

var file_storage_process_indicator_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x64, 0x55, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0xd4,
	0x01, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06,
	0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x16, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x65, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x72,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x41, 0x72, 0x67, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x17, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x65, 0x78, 0x65, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x65,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x72, 0x67, 0x73, 0x22, 0xd6, 0x03, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x78, 0x65,
	0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x64, 0x12, 0x45, 0x0a,
	0x0c, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x6e, 0x65,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x5f, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x65, 0x63, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x42, 0x31, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_process_indicator_proto_rawDescOnce sync.Once
	file_storage_process_indicator_proto_rawDescData []byte
)

func file_storage_process_indicator_proto_rawDescGZIP() []byte {
	file_storage_process_indicator_proto_rawDescOnce.Do(func() {
		file_storage_process_indicator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_process_indicator_proto_rawDesc), len(file_storage_process_indicator_proto_rawDesc)))
	})
	return file_storage_process_indicator_proto_rawDescData
}

var file_storage_process_indicator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storage_process_indicator_proto_goTypes = []any{
	(*ProcessIndicator)(nil),          // 0: storage.ProcessIndicator
	(*ProcessIndicatorUniqueKey)(nil), // 1: storage.ProcessIndicatorUniqueKey
	(*NetworkProcessUniqueKey)(nil),   // 2: storage.NetworkProcessUniqueKey
	(*ProcessSignal)(nil),             // 3: storage.ProcessSignal
	(*ProcessSignal_LineageInfo)(nil), // 4: storage.ProcessSignal.LineageInfo
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
}
var file_storage_process_indicator_proto_depIdxs = []int32{
	3, // 0: storage.ProcessIndicator.signal:type_name -> storage.ProcessSignal
	5, // 1: storage.ProcessIndicator.container_start_time:type_name -> google.protobuf.Timestamp
	5, // 2: storage.ProcessSignal.time:type_name -> google.protobuf.Timestamp
	4, // 3: storage.ProcessSignal.lineage_info:type_name -> storage.ProcessSignal.LineageInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_process_indicator_proto_init() }
func file_storage_process_indicator_proto_init() {
	if File_storage_process_indicator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_process_indicator_proto_rawDesc), len(file_storage_process_indicator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_process_indicator_proto_goTypes,
		DependencyIndexes: file_storage_process_indicator_proto_depIdxs,
		MessageInfos:      file_storage_process_indicator_proto_msgTypes,
	}.Build()
	File_storage_process_indicator_proto = out.File
	file_storage_process_indicator_proto_goTypes = nil
	file_storage_process_indicator_proto_depIdxs = nil
}
