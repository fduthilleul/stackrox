// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.3
// source: api/v1/delegated_registry_config_service.proto

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

type DelegatedRegistryConfig_EnabledFor int32

const (
	// Scan all images via central services except for images from the OCP integrated registry
	DelegatedRegistryConfig_NONE DelegatedRegistryConfig_EnabledFor = 0
	// Scan all images via the secured clusters
	DelegatedRegistryConfig_ALL DelegatedRegistryConfig_EnabledFor = 1
	// Scan images that match `registries` or are from the OCP integrated registry via the secured clusters
	// otherwise scan via central services
	DelegatedRegistryConfig_SPECIFIC DelegatedRegistryConfig_EnabledFor = 2
)

// Enum value maps for DelegatedRegistryConfig_EnabledFor.
var (
	DelegatedRegistryConfig_EnabledFor_name = map[int32]string{
		0: "NONE",
		1: "ALL",
		2: "SPECIFIC",
	}
	DelegatedRegistryConfig_EnabledFor_value = map[string]int32{
		"NONE":     0,
		"ALL":      1,
		"SPECIFIC": 2,
	}
)

func (x DelegatedRegistryConfig_EnabledFor) Enum() *DelegatedRegistryConfig_EnabledFor {
	p := new(DelegatedRegistryConfig_EnabledFor)
	*p = x
	return p
}

func (x DelegatedRegistryConfig_EnabledFor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DelegatedRegistryConfig_EnabledFor) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_delegated_registry_config_service_proto_enumTypes[0].Descriptor()
}

func (DelegatedRegistryConfig_EnabledFor) Type() protoreflect.EnumType {
	return &file_api_v1_delegated_registry_config_service_proto_enumTypes[0]
}

func (x DelegatedRegistryConfig_EnabledFor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DelegatedRegistryConfig_EnabledFor.Descriptor instead.
func (DelegatedRegistryConfig_EnabledFor) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_delegated_registry_config_service_proto_rawDescGZIP(), []int{2, 0}
}

type DelegatedRegistryClustersResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Clusters      []*DelegatedRegistryCluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelegatedRegistryClustersResponse) Reset() {
	*x = DelegatedRegistryClustersResponse{}
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelegatedRegistryClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatedRegistryClustersResponse) ProtoMessage() {}

func (x *DelegatedRegistryClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegatedRegistryClustersResponse.ProtoReflect.Descriptor instead.
func (*DelegatedRegistryClustersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_delegated_registry_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *DelegatedRegistryClustersResponse) GetClusters() []*DelegatedRegistryCluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DelegatedRegistryCluster struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsValid       bool                   `protobuf:"varint,3,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelegatedRegistryCluster) Reset() {
	*x = DelegatedRegistryCluster{}
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelegatedRegistryCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatedRegistryCluster) ProtoMessage() {}

func (x *DelegatedRegistryCluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegatedRegistryCluster.ProtoReflect.Descriptor instead.
func (*DelegatedRegistryCluster) Descriptor() ([]byte, []int) {
	return file_api_v1_delegated_registry_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *DelegatedRegistryCluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DelegatedRegistryCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DelegatedRegistryCluster) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

// DelegatedRegistryConfig determines if and where scan requests are delegated to, such as kept in
// central services or sent to particular secured clusters.
type DelegatedRegistryConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Determines if delegation is enabled for no registries, all registries, or specific registries
	EnabledFor DelegatedRegistryConfig_EnabledFor `protobuf:"varint,1,opt,name=enabled_for,json=enabledFor,proto3,enum=v1.DelegatedRegistryConfig_EnabledFor" json:"enabled_for,omitempty"`
	// The default cluster to delegate ad-hoc requests to if no match found in registries, not used
	// if `enabled for` is NONE
	DefaultClusterId string `protobuf:"bytes,2,opt,name=default_cluster_id,json=defaultClusterId,proto3" json:"default_cluster_id,omitempty"`
	// If `enabled for` is NONE registries has no effect.
	//
	// If `ALL` registries directs ad-hoc requests to the specified secured clusters if the path matches.
	//
	// If `SPECIFIC` registries directs ad-hoc requests to the specified secured clusters just like with `ALL`,
	// but in addition images that match the specified paths will be scanned locally by the secured clusters
	// (images from the OCP integrated registry are always scanned locally). Images that do not match a path
	// will be scanned via central services
	Registries    []*DelegatedRegistryConfig_DelegatedRegistry `protobuf:"bytes,3,rep,name=registries,proto3" json:"registries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelegatedRegistryConfig) Reset() {
	*x = DelegatedRegistryConfig{}
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelegatedRegistryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatedRegistryConfig) ProtoMessage() {}

func (x *DelegatedRegistryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegatedRegistryConfig.ProtoReflect.Descriptor instead.
func (*DelegatedRegistryConfig) Descriptor() ([]byte, []int) {
	return file_api_v1_delegated_registry_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *DelegatedRegistryConfig) GetEnabledFor() DelegatedRegistryConfig_EnabledFor {
	if x != nil {
		return x.EnabledFor
	}
	return DelegatedRegistryConfig_NONE
}

func (x *DelegatedRegistryConfig) GetDefaultClusterId() string {
	if x != nil {
		return x.DefaultClusterId
	}
	return ""
}

func (x *DelegatedRegistryConfig) GetRegistries() []*DelegatedRegistryConfig_DelegatedRegistry {
	if x != nil {
		return x.Registries
	}
	return nil
}

type DelegatedRegistryConfig_DelegatedRegistry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Registry + optional path, ie: quay.example.com/prod
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ID of the cluster to delegate ad-hoc requests to
	ClusterId     string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelegatedRegistryConfig_DelegatedRegistry) Reset() {
	*x = DelegatedRegistryConfig_DelegatedRegistry{}
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelegatedRegistryConfig_DelegatedRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatedRegistryConfig_DelegatedRegistry) ProtoMessage() {}

func (x *DelegatedRegistryConfig_DelegatedRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delegated_registry_config_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegatedRegistryConfig_DelegatedRegistry.ProtoReflect.Descriptor instead.
func (*DelegatedRegistryConfig_DelegatedRegistry) Descriptor() ([]byte, []int) {
	return file_api_v1_delegated_registry_config_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DelegatedRegistryConfig_DelegatedRegistry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DelegatedRegistryConfig_DelegatedRegistry) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

var File_api_v1_delegated_registry_config_service_proto protoreflect.FileDescriptor

var file_api_v1_delegated_registry_config_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x59, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x22, 0xd6, 0x02, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x0b,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x1a, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0a, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x10, 0x02, 0x32, 0xdb, 0x02, 0x0a, 0x1e, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12,
	0x24, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x70, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x1a, 0x1b, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_delegated_registry_config_service_proto_rawDescOnce sync.Once
	file_api_v1_delegated_registry_config_service_proto_rawDescData = file_api_v1_delegated_registry_config_service_proto_rawDesc
)

func file_api_v1_delegated_registry_config_service_proto_rawDescGZIP() []byte {
	file_api_v1_delegated_registry_config_service_proto_rawDescOnce.Do(func() {
		file_api_v1_delegated_registry_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_delegated_registry_config_service_proto_rawDescData)
	})
	return file_api_v1_delegated_registry_config_service_proto_rawDescData
}

var file_api_v1_delegated_registry_config_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_delegated_registry_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_delegated_registry_config_service_proto_goTypes = []any{
	(DelegatedRegistryConfig_EnabledFor)(0),           // 0: v1.DelegatedRegistryConfig.EnabledFor
	(*DelegatedRegistryClustersResponse)(nil),         // 1: v1.DelegatedRegistryClustersResponse
	(*DelegatedRegistryCluster)(nil),                  // 2: v1.DelegatedRegistryCluster
	(*DelegatedRegistryConfig)(nil),                   // 3: v1.DelegatedRegistryConfig
	(*DelegatedRegistryConfig_DelegatedRegistry)(nil), // 4: v1.DelegatedRegistryConfig.DelegatedRegistry
	(*Empty)(nil), // 5: v1.Empty
}
var file_api_v1_delegated_registry_config_service_proto_depIdxs = []int32{
	2, // 0: v1.DelegatedRegistryClustersResponse.clusters:type_name -> v1.DelegatedRegistryCluster
	0, // 1: v1.DelegatedRegistryConfig.enabled_for:type_name -> v1.DelegatedRegistryConfig.EnabledFor
	4, // 2: v1.DelegatedRegistryConfig.registries:type_name -> v1.DelegatedRegistryConfig.DelegatedRegistry
	5, // 3: v1.DelegatedRegistryConfigService.GetConfig:input_type -> v1.Empty
	5, // 4: v1.DelegatedRegistryConfigService.GetClusters:input_type -> v1.Empty
	3, // 5: v1.DelegatedRegistryConfigService.UpdateConfig:input_type -> v1.DelegatedRegistryConfig
	3, // 6: v1.DelegatedRegistryConfigService.GetConfig:output_type -> v1.DelegatedRegistryConfig
	1, // 7: v1.DelegatedRegistryConfigService.GetClusters:output_type -> v1.DelegatedRegistryClustersResponse
	3, // 8: v1.DelegatedRegistryConfigService.UpdateConfig:output_type -> v1.DelegatedRegistryConfig
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_delegated_registry_config_service_proto_init() }
func file_api_v1_delegated_registry_config_service_proto_init() {
	if File_api_v1_delegated_registry_config_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_delegated_registry_config_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_delegated_registry_config_service_proto_goTypes,
		DependencyIndexes: file_api_v1_delegated_registry_config_service_proto_depIdxs,
		EnumInfos:         file_api_v1_delegated_registry_config_service_proto_enumTypes,
		MessageInfos:      file_api_v1_delegated_registry_config_service_proto_msgTypes,
	}.Build()
	File_api_v1_delegated_registry_config_service_proto = out.File
	file_api_v1_delegated_registry_config_service_proto_rawDesc = nil
	file_api_v1_delegated_registry_config_service_proto_goTypes = nil
	file_api_v1_delegated_registry_config_service_proto_depIdxs = nil
}
