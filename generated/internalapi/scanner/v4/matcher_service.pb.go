// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: internalapi/scanner/v4/matcher_service.proto

package v4

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetVulnerabilitiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HashId        string                 `protobuf:"bytes,1,opt,name=hash_id,json=hashId,proto3" json:"hash_id,omitempty"`
	Contents      *Contents              `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVulnerabilitiesRequest) Reset() {
	*x = GetVulnerabilitiesRequest{}
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVulnerabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVulnerabilitiesRequest) ProtoMessage() {}

func (x *GetVulnerabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVulnerabilitiesRequest.ProtoReflect.Descriptor instead.
func (*GetVulnerabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetVulnerabilitiesRequest) GetHashId() string {
	if x != nil {
		return x.HashId
	}
	return ""
}

func (x *GetVulnerabilitiesRequest) GetContents() *Contents {
	if x != nil {
		return x.Contents
	}
	return nil
}

type GetSBOMRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id is a unique identifier (ie: sha256 digest) that represents the contents.
	// For SPDX 2.3 this will be the document name.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the user requested name for the contents (ie: for images this may be the images full name with tag).
	// For SPDX 2.3 this will be included to the document comment.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// uri is a unique absolute Uniform Resource Identifier (URI) for this document (if applicable).
	// For SPDX 2.3 this will be the document namespace.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// contents is the packages, versions, etc. that will be converted to an SBOM.
	Contents      *Contents `protobuf:"bytes,4,opt,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSBOMRequest) Reset() {
	*x = GetSBOMRequest{}
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSBOMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSBOMRequest) ProtoMessage() {}

func (x *GetSBOMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSBOMRequest.ProtoReflect.Descriptor instead.
func (*GetSBOMRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSBOMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSBOMRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSBOMRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *GetSBOMRequest) GetContents() *Contents {
	if x != nil {
		return x.Contents
	}
	return nil
}

type GetSBOMResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sbom          []byte                 `protobuf:"bytes,1,opt,name=sbom,proto3" json:"sbom,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSBOMResponse) Reset() {
	*x = GetSBOMResponse{}
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSBOMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSBOMResponse) ProtoMessage() {}

func (x *GetSBOMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSBOMResponse.ProtoReflect.Descriptor instead.
func (*GetSBOMResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetSBOMResponse) GetSbom() []byte {
	if x != nil {
		return x.Sbom
	}
	return nil
}

type Metadata struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	LastVulnerabilityUpdate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=LastVulnerabilityUpdate,proto3" json:"LastVulnerabilityUpdate,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP(), []int{3}
}

func (x *Metadata) GetLastVulnerabilityUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastVulnerabilityUpdate
	}
	return nil
}

var File_internalapi_scanner_v4_matcher_service_proto protoreflect.FileDescriptor

var file_internalapi_scanner_v4_matcher_service_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x78, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x42, 0x4f, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x42, 0x4f, 0x4d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x62, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x62, 0x6f, 0x6d, 0x22, 0x60, 0x0a, 0x08, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x17, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x17, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x32, 0xe8, 0x01, 0x0a, 0x07,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x25, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x34, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x42, 0x4f, 0x4d, 0x12, 0x1a, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x42,
	0x4f, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x42, 0x4f, 0x4d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f,
	0x76, 0x34, 0x3b, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internalapi_scanner_v4_matcher_service_proto_rawDescOnce sync.Once
	file_internalapi_scanner_v4_matcher_service_proto_rawDescData []byte
)

func file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP() []byte {
	file_internalapi_scanner_v4_matcher_service_proto_rawDescOnce.Do(func() {
		file_internalapi_scanner_v4_matcher_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internalapi_scanner_v4_matcher_service_proto_rawDesc), len(file_internalapi_scanner_v4_matcher_service_proto_rawDesc)))
	})
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescData
}

var file_internalapi_scanner_v4_matcher_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_scanner_v4_matcher_service_proto_goTypes = []any{
	(*GetVulnerabilitiesRequest)(nil), // 0: scanner.v4.GetVulnerabilitiesRequest
	(*GetSBOMRequest)(nil),            // 1: scanner.v4.GetSBOMRequest
	(*GetSBOMResponse)(nil),           // 2: scanner.v4.GetSBOMResponse
	(*Metadata)(nil),                  // 3: scanner.v4.Metadata
	(*Contents)(nil),                  // 4: scanner.v4.Contents
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
	(*VulnerabilityReport)(nil),       // 7: scanner.v4.VulnerabilityReport
}
var file_internalapi_scanner_v4_matcher_service_proto_depIdxs = []int32{
	4, // 0: scanner.v4.GetVulnerabilitiesRequest.contents:type_name -> scanner.v4.Contents
	4, // 1: scanner.v4.GetSBOMRequest.contents:type_name -> scanner.v4.Contents
	5, // 2: scanner.v4.Metadata.LastVulnerabilityUpdate:type_name -> google.protobuf.Timestamp
	0, // 3: scanner.v4.Matcher.GetVulnerabilities:input_type -> scanner.v4.GetVulnerabilitiesRequest
	6, // 4: scanner.v4.Matcher.GetMetadata:input_type -> google.protobuf.Empty
	1, // 5: scanner.v4.Matcher.GetSBOM:input_type -> scanner.v4.GetSBOMRequest
	7, // 6: scanner.v4.Matcher.GetVulnerabilities:output_type -> scanner.v4.VulnerabilityReport
	3, // 7: scanner.v4.Matcher.GetMetadata:output_type -> scanner.v4.Metadata
	2, // 8: scanner.v4.Matcher.GetSBOM:output_type -> scanner.v4.GetSBOMResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internalapi_scanner_v4_matcher_service_proto_init() }
func file_internalapi_scanner_v4_matcher_service_proto_init() {
	if File_internalapi_scanner_v4_matcher_service_proto != nil {
		return
	}
	file_internalapi_scanner_v4_common_proto_init()
	file_internalapi_scanner_v4_vulnerability_report_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internalapi_scanner_v4_matcher_service_proto_rawDesc), len(file_internalapi_scanner_v4_matcher_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_scanner_v4_matcher_service_proto_goTypes,
		DependencyIndexes: file_internalapi_scanner_v4_matcher_service_proto_depIdxs,
		MessageInfos:      file_internalapi_scanner_v4_matcher_service_proto_msgTypes,
	}.Build()
	File_internalapi_scanner_v4_matcher_service_proto = out.File
	file_internalapi_scanner_v4_matcher_service_proto_goTypes = nil
	file_internalapi_scanner_v4_matcher_service_proto_depIdxs = nil
}
