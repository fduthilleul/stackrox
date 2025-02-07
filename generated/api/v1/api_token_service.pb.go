// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: api/v1/api_token_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GenerateTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Marked as deprecated in api/v1/api_token_service.proto.
	Role          string                 `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Roles         []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1/api_token_service.proto.
func (x *GenerateTokenRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GenerateTokenRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GenerateTokenRequest) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Metadata      *storage.TokenMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GenerateTokenResponse) GetMetadata() *storage.TokenMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type GetAPITokensRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RevokedOneof:
	//
	//	*GetAPITokensRequest_Revoked
	RevokedOneof  isGetAPITokensRequest_RevokedOneof `protobuf_oneof:"revoked_oneof"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAPITokensRequest) Reset() {
	*x = GetAPITokensRequest{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAPITokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPITokensRequest) ProtoMessage() {}

func (x *GetAPITokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPITokensRequest.ProtoReflect.Descriptor instead.
func (*GetAPITokensRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAPITokensRequest) GetRevokedOneof() isGetAPITokensRequest_RevokedOneof {
	if x != nil {
		return x.RevokedOneof
	}
	return nil
}

func (x *GetAPITokensRequest) GetRevoked() bool {
	if x != nil {
		if x, ok := x.RevokedOneof.(*GetAPITokensRequest_Revoked); ok {
			return x.Revoked
		}
	}
	return false
}

type isGetAPITokensRequest_RevokedOneof interface {
	isGetAPITokensRequest_RevokedOneof()
}

type GetAPITokensRequest_Revoked struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked,proto3,oneof"`
}

func (*GetAPITokensRequest_Revoked) isGetAPITokensRequest_RevokedOneof() {}

type GetAPITokensResponse struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Tokens        []*storage.TokenMetadata `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAPITokensResponse) Reset() {
	*x = GetAPITokensResponse{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAPITokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPITokensResponse) ProtoMessage() {}

func (x *GetAPITokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPITokensResponse.ProtoReflect.Descriptor instead.
func (*GetAPITokensResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAPITokensResponse) GetTokens() []*storage.TokenMetadata {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type ListAllowedTokenRolesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleNames     []string               `protobuf:"bytes,1,rep,name=roleNames,proto3" json:"roleNames,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllowedTokenRolesResponse) Reset() {
	*x = ListAllowedTokenRolesResponse{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllowedTokenRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllowedTokenRolesResponse) ProtoMessage() {}

func (x *ListAllowedTokenRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllowedTokenRolesResponse.ProtoReflect.Descriptor instead.
func (*ListAllowedTokenRolesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllowedTokenRolesResponse) GetRoleNames() []string {
	if x != nil {
		return x.RoleNames
	}
	return nil
}

var File_api_v1_api_token_service_proto protoreflect.FileDescriptor

var file_api_v1_api_token_service_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x15,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x3d, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0xed, 0x03, 0x0a, 0x0f, 0x41,
	0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a,
	0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x58, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x67, 0x0a,
	0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x09,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1_api_token_service_proto_rawDescOnce sync.Once
	file_api_v1_api_token_service_proto_rawDescData []byte
)

func file_api_v1_api_token_service_proto_rawDescGZIP() []byte {
	file_api_v1_api_token_service_proto_rawDescOnce.Do(func() {
		file_api_v1_api_token_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_api_token_service_proto_rawDesc), len(file_api_v1_api_token_service_proto_rawDesc)))
	})
	return file_api_v1_api_token_service_proto_rawDescData
}

var file_api_v1_api_token_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_api_token_service_proto_goTypes = []any{
	(*GenerateTokenRequest)(nil),          // 0: v1.GenerateTokenRequest
	(*GenerateTokenResponse)(nil),         // 1: v1.GenerateTokenResponse
	(*GetAPITokensRequest)(nil),           // 2: v1.GetAPITokensRequest
	(*GetAPITokensResponse)(nil),          // 3: v1.GetAPITokensResponse
	(*ListAllowedTokenRolesResponse)(nil), // 4: v1.ListAllowedTokenRolesResponse
	(*timestamppb.Timestamp)(nil),         // 5: google.protobuf.Timestamp
	(*storage.TokenMetadata)(nil),         // 6: storage.TokenMetadata
	(*ResourceByID)(nil),                  // 7: v1.ResourceByID
	(*Empty)(nil),                         // 8: v1.Empty
}
var file_api_v1_api_token_service_proto_depIdxs = []int32{
	5, // 0: v1.GenerateTokenRequest.expiration:type_name -> google.protobuf.Timestamp
	6, // 1: v1.GenerateTokenResponse.metadata:type_name -> storage.TokenMetadata
	6, // 2: v1.GetAPITokensResponse.tokens:type_name -> storage.TokenMetadata
	7, // 3: v1.APITokenService.GetAPIToken:input_type -> v1.ResourceByID
	2, // 4: v1.APITokenService.GetAPITokens:input_type -> v1.GetAPITokensRequest
	0, // 5: v1.APITokenService.GenerateToken:input_type -> v1.GenerateTokenRequest
	7, // 6: v1.APITokenService.RevokeToken:input_type -> v1.ResourceByID
	8, // 7: v1.APITokenService.ListAllowedTokenRoles:input_type -> v1.Empty
	6, // 8: v1.APITokenService.GetAPIToken:output_type -> storage.TokenMetadata
	3, // 9: v1.APITokenService.GetAPITokens:output_type -> v1.GetAPITokensResponse
	1, // 10: v1.APITokenService.GenerateToken:output_type -> v1.GenerateTokenResponse
	8, // 11: v1.APITokenService.RevokeToken:output_type -> v1.Empty
	4, // 12: v1.APITokenService.ListAllowedTokenRoles:output_type -> v1.ListAllowedTokenRolesResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_api_token_service_proto_init() }
func file_api_v1_api_token_service_proto_init() {
	if File_api_v1_api_token_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_api_token_service_proto_msgTypes[2].OneofWrappers = []any{
		(*GetAPITokensRequest_Revoked)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_api_token_service_proto_rawDesc), len(file_api_v1_api_token_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_api_token_service_proto_goTypes,
		DependencyIndexes: file_api_v1_api_token_service_proto_depIdxs,
		MessageInfos:      file_api_v1_api_token_service_proto_msgTypes,
	}.Build()
	File_api_v1_api_token_service_proto = out.File
	file_api_v1_api_token_service_proto_goTypes = nil
	file_api_v1_api_token_service_proto_depIdxs = nil
}
