// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.3
// source: api/v1/network_baseline_service.proto

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

// Status of this peer connection. As of now we only have two statuses:
//   - BASELINE: the connection is in the current deployment baseline
//   - ANOMALOUS: the connection is not recognized by the current deployment baseline
type NetworkBaselinePeerStatus_Status int32

const (
	NetworkBaselinePeerStatus_BASELINE  NetworkBaselinePeerStatus_Status = 0
	NetworkBaselinePeerStatus_ANOMALOUS NetworkBaselinePeerStatus_Status = 1
)

// Enum value maps for NetworkBaselinePeerStatus_Status.
var (
	NetworkBaselinePeerStatus_Status_name = map[int32]string{
		0: "BASELINE",
		1: "ANOMALOUS",
	}
	NetworkBaselinePeerStatus_Status_value = map[string]int32{
		"BASELINE":  0,
		"ANOMALOUS": 1,
	}
)

func (x NetworkBaselinePeerStatus_Status) Enum() *NetworkBaselinePeerStatus_Status {
	p := new(NetworkBaselinePeerStatus_Status)
	*p = x
	return p
}

func (x NetworkBaselinePeerStatus_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkBaselinePeerStatus_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_network_baseline_service_proto_enumTypes[0].Descriptor()
}

func (NetworkBaselinePeerStatus_Status) Type() protoreflect.EnumType {
	return &file_api_v1_network_baseline_service_proto_enumTypes[0]
}

func (x NetworkBaselinePeerStatus_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkBaselinePeerStatus_Status.Descriptor instead.
func (NetworkBaselinePeerStatus_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{2, 0}
}

type NetworkBaselinePeerEntity struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Id            string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          storage.NetworkEntityInfo_Type `protobuf:"varint,2,opt,name=type,proto3,enum=storage.NetworkEntityInfo_Type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkBaselinePeerEntity) Reset() {
	*x = NetworkBaselinePeerEntity{}
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkBaselinePeerEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBaselinePeerEntity) ProtoMessage() {}

func (x *NetworkBaselinePeerEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBaselinePeerEntity.ProtoReflect.Descriptor instead.
func (*NetworkBaselinePeerEntity) Descriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkBaselinePeerEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkBaselinePeerEntity) GetType() storage.NetworkEntityInfo_Type {
	if x != nil {
		return x.Type
	}
	return storage.NetworkEntityInfo_Type(0)
}

type NetworkBaselineStatusPeer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The peer entity of the connection. This can be constructed from the
	// entity object of the networkgraph API. Only the ID and type are required.
	Entity *NetworkBaselinePeerEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// The port and protocol of the destination of the given connection.
	Port     uint32             `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Protocol storage.L4Protocol `protobuf:"varint,3,opt,name=protocol,proto3,enum=storage.L4Protocol" json:"protocol,omitempty"`
	// A boolean representing whether the query is for an ingress or egress
	// connection. This is defined with respect to the current deployment.
	// Thus:
	//   - If the connection in question is in the outEdges of the current deployment,
	//     this should be false.
	//   - If it is in the outEdges of the peer deployment, this
	//     should be true.
	Ingress       bool `protobuf:"varint,4,opt,name=ingress,proto3" json:"ingress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkBaselineStatusPeer) Reset() {
	*x = NetworkBaselineStatusPeer{}
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkBaselineStatusPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBaselineStatusPeer) ProtoMessage() {}

func (x *NetworkBaselineStatusPeer) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBaselineStatusPeer.ProtoReflect.Descriptor instead.
func (*NetworkBaselineStatusPeer) Descriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkBaselineStatusPeer) GetEntity() *NetworkBaselinePeerEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *NetworkBaselineStatusPeer) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *NetworkBaselineStatusPeer) GetProtocol() storage.L4Protocol {
	if x != nil {
		return x.Protocol
	}
	return storage.L4Protocol(0)
}

func (x *NetworkBaselineStatusPeer) GetIngress() bool {
	if x != nil {
		return x.Ingress
	}
	return false
}

type NetworkBaselinePeerStatus struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Peer          *NetworkBaselineStatusPeer       `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Status        NetworkBaselinePeerStatus_Status `protobuf:"varint,2,opt,name=status,proto3,enum=v1.NetworkBaselinePeerStatus_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkBaselinePeerStatus) Reset() {
	*x = NetworkBaselinePeerStatus{}
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkBaselinePeerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBaselinePeerStatus) ProtoMessage() {}

func (x *NetworkBaselinePeerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBaselinePeerStatus.ProtoReflect.Descriptor instead.
func (*NetworkBaselinePeerStatus) Descriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkBaselinePeerStatus) GetPeer() *NetworkBaselineStatusPeer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *NetworkBaselinePeerStatus) GetStatus() NetworkBaselinePeerStatus_Status {
	if x != nil {
		return x.Status
	}
	return NetworkBaselinePeerStatus_BASELINE
}

type NetworkBaselineStatusRequest struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	DeploymentId  string                       `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Peers         []*NetworkBaselineStatusPeer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkBaselineStatusRequest) Reset() {
	*x = NetworkBaselineStatusRequest{}
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkBaselineStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBaselineStatusRequest) ProtoMessage() {}

func (x *NetworkBaselineStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBaselineStatusRequest.ProtoReflect.Descriptor instead.
func (*NetworkBaselineStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkBaselineStatusRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *NetworkBaselineStatusRequest) GetPeers() []*NetworkBaselineStatusPeer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type NetworkBaselineStatusResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Statuses      []*NetworkBaselinePeerStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkBaselineStatusResponse) Reset() {
	*x = NetworkBaselineStatusResponse{}
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkBaselineStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBaselineStatusResponse) ProtoMessage() {}

func (x *NetworkBaselineStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBaselineStatusResponse.ProtoReflect.Descriptor instead.
func (*NetworkBaselineStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkBaselineStatusResponse) GetStatuses() []*NetworkBaselinePeerStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type ModifyBaselineStatusForPeersRequest struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	DeploymentId  string                       `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Peers         []*NetworkBaselinePeerStatus `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifyBaselineStatusForPeersRequest) Reset() {
	*x = ModifyBaselineStatusForPeersRequest{}
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyBaselineStatusForPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyBaselineStatusForPeersRequest) ProtoMessage() {}

func (x *ModifyBaselineStatusForPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_network_baseline_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyBaselineStatusForPeersRequest.ProtoReflect.Descriptor instead.
func (*ModifyBaselineStatusForPeersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_network_baseline_service_proto_rawDescGZIP(), []int{5}
}

func (x *ModifyBaselineStatusForPeersRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ModifyBaselineStatusForPeersRequest) GetPeers() []*NetworkBaselinePeerStatus {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_api_v1_network_baseline_service_proto protoreflect.FileDescriptor

var file_api_v1_network_baseline_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60,
	0x0a, 0x19, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xb1, 0x01, 0x0a, 0x19, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x65, 0x72, 0x12, 0x35,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x34, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x19, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x25, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a,
	0x08, 0x42, 0x41, 0x53, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x4e, 0x4f, 0x4d, 0x41, 0x4c, 0x4f, 0x55, 0x53, 0x10, 0x01, 0x22, 0x78, 0x0a, 0x1c, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x1d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42,
	0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x22, 0x7f, 0x0a, 0x23, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x32, 0xe8, 0x04, 0x0a, 0x16, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x6f, 0x72, 0x46, 0x6c, 0x6f, 0x77,
	0x73, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61,
	0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01,
	0x2a, 0x22, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61,
	0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x62, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x88, 0x01, 0x0a, 0x1c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x42, 0x61, 0x73, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x65, 0x65,
	0x72, 0x73, 0x12, 0x27, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x42, 0x61,
	0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x6f, 0x72, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01,
	0x2a, 0x32, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61,
	0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x13,
	0x4c, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x32, 0x1d, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x60, 0x0a, 0x15, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x32, 0x1f, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x27, 0x0a, 0x18,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_network_baseline_service_proto_rawDescOnce sync.Once
	file_api_v1_network_baseline_service_proto_rawDescData = file_api_v1_network_baseline_service_proto_rawDesc
)

func file_api_v1_network_baseline_service_proto_rawDescGZIP() []byte {
	file_api_v1_network_baseline_service_proto_rawDescOnce.Do(func() {
		file_api_v1_network_baseline_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_network_baseline_service_proto_rawDescData)
	})
	return file_api_v1_network_baseline_service_proto_rawDescData
}

var file_api_v1_network_baseline_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_network_baseline_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_network_baseline_service_proto_goTypes = []any{
	(NetworkBaselinePeerStatus_Status)(0),       // 0: v1.NetworkBaselinePeerStatus.Status
	(*NetworkBaselinePeerEntity)(nil),           // 1: v1.NetworkBaselinePeerEntity
	(*NetworkBaselineStatusPeer)(nil),           // 2: v1.NetworkBaselineStatusPeer
	(*NetworkBaselinePeerStatus)(nil),           // 3: v1.NetworkBaselinePeerStatus
	(*NetworkBaselineStatusRequest)(nil),        // 4: v1.NetworkBaselineStatusRequest
	(*NetworkBaselineStatusResponse)(nil),       // 5: v1.NetworkBaselineStatusResponse
	(*ModifyBaselineStatusForPeersRequest)(nil), // 6: v1.ModifyBaselineStatusForPeersRequest
	(storage.NetworkEntityInfo_Type)(0),         // 7: storage.NetworkEntityInfo.Type
	(storage.L4Protocol)(0),                     // 8: storage.L4Protocol
	(*ResourceByID)(nil),                        // 9: v1.ResourceByID
	(*storage.NetworkBaseline)(nil),             // 10: storage.NetworkBaseline
	(*Empty)(nil),                               // 11: v1.Empty
}
var file_api_v1_network_baseline_service_proto_depIdxs = []int32{
	7,  // 0: v1.NetworkBaselinePeerEntity.type:type_name -> storage.NetworkEntityInfo.Type
	1,  // 1: v1.NetworkBaselineStatusPeer.entity:type_name -> v1.NetworkBaselinePeerEntity
	8,  // 2: v1.NetworkBaselineStatusPeer.protocol:type_name -> storage.L4Protocol
	2,  // 3: v1.NetworkBaselinePeerStatus.peer:type_name -> v1.NetworkBaselineStatusPeer
	0,  // 4: v1.NetworkBaselinePeerStatus.status:type_name -> v1.NetworkBaselinePeerStatus.Status
	2,  // 5: v1.NetworkBaselineStatusRequest.peers:type_name -> v1.NetworkBaselineStatusPeer
	3,  // 6: v1.NetworkBaselineStatusResponse.statuses:type_name -> v1.NetworkBaselinePeerStatus
	3,  // 7: v1.ModifyBaselineStatusForPeersRequest.peers:type_name -> v1.NetworkBaselinePeerStatus
	4,  // 8: v1.NetworkBaselineService.GetNetworkBaselineStatusForFlows:input_type -> v1.NetworkBaselineStatusRequest
	9,  // 9: v1.NetworkBaselineService.GetNetworkBaseline:input_type -> v1.ResourceByID
	6,  // 10: v1.NetworkBaselineService.ModifyBaselineStatusForPeers:input_type -> v1.ModifyBaselineStatusForPeersRequest
	9,  // 11: v1.NetworkBaselineService.LockNetworkBaseline:input_type -> v1.ResourceByID
	9,  // 12: v1.NetworkBaselineService.UnlockNetworkBaseline:input_type -> v1.ResourceByID
	5,  // 13: v1.NetworkBaselineService.GetNetworkBaselineStatusForFlows:output_type -> v1.NetworkBaselineStatusResponse
	10, // 14: v1.NetworkBaselineService.GetNetworkBaseline:output_type -> storage.NetworkBaseline
	11, // 15: v1.NetworkBaselineService.ModifyBaselineStatusForPeers:output_type -> v1.Empty
	11, // 16: v1.NetworkBaselineService.LockNetworkBaseline:output_type -> v1.Empty
	11, // 17: v1.NetworkBaselineService.UnlockNetworkBaseline:output_type -> v1.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1_network_baseline_service_proto_init() }
func file_api_v1_network_baseline_service_proto_init() {
	if File_api_v1_network_baseline_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_network_baseline_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_network_baseline_service_proto_goTypes,
		DependencyIndexes: file_api_v1_network_baseline_service_proto_depIdxs,
		EnumInfos:         file_api_v1_network_baseline_service_proto_enumTypes,
		MessageInfos:      file_api_v1_network_baseline_service_proto_msgTypes,
	}.Build()
	File_api_v1_network_baseline_service_proto = out.File
	file_api_v1_network_baseline_service_proto_rawDesc = nil
	file_api_v1_network_baseline_service_proto_goTypes = nil
	file_api_v1_network_baseline_service_proto_depIdxs = nil
}
