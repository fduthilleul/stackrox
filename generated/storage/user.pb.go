// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: storage/user.proto

package storage

import (
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

type SlimUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"User ID"`     // @gotags: search:"User ID"
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"User Name"` // @gotags: search:"User Name"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlimUser) Reset() {
	*x = SlimUser{}
	mi := &file_storage_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlimUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlimUser) ProtoMessage() {}

func (x *SlimUser) ProtoReflect() protoreflect.Message {
	mi := &file_storage_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlimUser.ProtoReflect.Descriptor instead.
func (*SlimUser) Descriptor() ([]byte, []int) {
	return file_storage_user_proto_rawDescGZIP(), []int{0}
}

func (x *SlimUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SlimUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// User is an object that allows us to track the roles a user is tied to, and how they logged in.
type User struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthProviderId string                 `protobuf:"bytes,2,opt,name=auth_provider_id,json=authProviderId,proto3" json:"auth_provider_id,omitempty"`
	Attributes     []*UserAttribute       `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	IdpToken       string                 `protobuf:"bytes,4,opt,name=idp_token,json=idpToken,proto3" json:"idp_token,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_storage_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_storage_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_storage_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetAuthProviderId() string {
	if x != nil {
		return x.AuthProviderId
	}
	return ""
}

func (x *User) GetAttributes() []*UserAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *User) GetIdpToken() string {
	if x != nil {
		return x.IdpToken
	}
	return ""
}

type UserAttribute struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAttribute) Reset() {
	*x = UserAttribute{}
	mi := &file_storage_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttribute) ProtoMessage() {}

func (x *UserAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_storage_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttribute.ProtoReflect.Descriptor instead.
func (*UserAttribute) Descriptor() ([]byte, []int) {
	return file_storage_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserAttribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UserAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Username      string                     `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FriendlyName  string                     `protobuf:"bytes,2,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Permissions   *UserInfo_ResourceToAccess `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Roles         []*UserInfo_Role           `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_storage_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_storage_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_storage_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *UserInfo) GetPermissions() *UserInfo_ResourceToAccess {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *UserInfo) GetRoles() []*UserInfo_Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

// Role is wire compatible with the old format of storage.Role and
// hence only includes role name and associated permissions.
type UserInfo_Role struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ResourceToAccess map[string]Access      `protobuf:"bytes,3,rep,name=resource_to_access,json=resourceToAccess,proto3" json:"resource_to_access,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=storage.Access"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UserInfo_Role) Reset() {
	*x = UserInfo_Role{}
	mi := &file_storage_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo_Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo_Role) ProtoMessage() {}

func (x *UserInfo_Role) ProtoReflect() protoreflect.Message {
	mi := &file_storage_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo_Role.ProtoReflect.Descriptor instead.
func (*UserInfo_Role) Descriptor() ([]byte, []int) {
	return file_storage_user_proto_rawDescGZIP(), []int{3, 0}
}

func (x *UserInfo_Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo_Role) GetResourceToAccess() map[string]Access {
	if x != nil {
		return x.ResourceToAccess
	}
	return nil
}

// ResourceToAccess represents a collection of permissions. It is wire
// compatible with the old format of storage.Role and replaces it in
// places where only aggregated permissions are required.
type UserInfo_ResourceToAccess struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ResourceToAccess map[string]Access      `protobuf:"bytes,3,rep,name=resource_to_access,json=resourceToAccess,proto3" json:"resource_to_access,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=storage.Access"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UserInfo_ResourceToAccess) Reset() {
	*x = UserInfo_ResourceToAccess{}
	mi := &file_storage_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo_ResourceToAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo_ResourceToAccess) ProtoMessage() {}

func (x *UserInfo_ResourceToAccess) ProtoReflect() protoreflect.Message {
	mi := &file_storage_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo_ResourceToAccess.ProtoReflect.Descriptor instead.
func (*UserInfo_ResourceToAccess) Descriptor() ([]byte, []int) {
	return file_storage_user_proto_rawDescGZIP(), []int{3, 1}
}

func (x *UserInfo_ResourceToAccess) GetResourceToAccess() map[string]Access {
	if x != nil {
		return x.ResourceToAccess
	}
	return nil
}

var File_storage_user_proto protoreflect.FileDescriptor

var file_storage_user_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x53, 0x6c, 0x69, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x95, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x64, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x64, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x37, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xf9, 0x04, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x44, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x1a, 0xd2, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x5a, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x6f,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x54, 0x0a,
	0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x1a, 0xdc, 0x01, 0x0a, 0x10, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x66,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x54, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x42, 0x2e,
	0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_user_proto_rawDescOnce sync.Once
	file_storage_user_proto_rawDescData []byte
)

func file_storage_user_proto_rawDescGZIP() []byte {
	file_storage_user_proto_rawDescOnce.Do(func() {
		file_storage_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_user_proto_rawDesc), len(file_storage_user_proto_rawDesc)))
	})
	return file_storage_user_proto_rawDescData
}

var file_storage_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_storage_user_proto_goTypes = []any{
	(*SlimUser)(nil),                  // 0: storage.SlimUser
	(*User)(nil),                      // 1: storage.User
	(*UserAttribute)(nil),             // 2: storage.UserAttribute
	(*UserInfo)(nil),                  // 3: storage.UserInfo
	(*UserInfo_Role)(nil),             // 4: storage.UserInfo.Role
	(*UserInfo_ResourceToAccess)(nil), // 5: storage.UserInfo.ResourceToAccess
	nil,                               // 6: storage.UserInfo.Role.ResourceToAccessEntry
	nil,                               // 7: storage.UserInfo.ResourceToAccess.ResourceToAccessEntry
	(Access)(0),                       // 8: storage.Access
}
var file_storage_user_proto_depIdxs = []int32{
	2, // 0: storage.User.attributes:type_name -> storage.UserAttribute
	5, // 1: storage.UserInfo.permissions:type_name -> storage.UserInfo.ResourceToAccess
	4, // 2: storage.UserInfo.roles:type_name -> storage.UserInfo.Role
	6, // 3: storage.UserInfo.Role.resource_to_access:type_name -> storage.UserInfo.Role.ResourceToAccessEntry
	7, // 4: storage.UserInfo.ResourceToAccess.resource_to_access:type_name -> storage.UserInfo.ResourceToAccess.ResourceToAccessEntry
	8, // 5: storage.UserInfo.Role.ResourceToAccessEntry.value:type_name -> storage.Access
	8, // 6: storage.UserInfo.ResourceToAccess.ResourceToAccessEntry.value:type_name -> storage.Access
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_storage_user_proto_init() }
func file_storage_user_proto_init() {
	if File_storage_user_proto != nil {
		return
	}
	file_storage_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_user_proto_rawDesc), len(file_storage_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_user_proto_goTypes,
		DependencyIndexes: file_storage_user_proto_depIdxs,
		MessageInfos:      file_storage_user_proto_msgTypes,
	}.Build()
	File_storage_user_proto = out.File
	file_storage_user_proto_goTypes = nil
	file_storage_user_proto_depIdxs = nil
}
