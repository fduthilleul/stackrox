// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: storage/signature_integration.proto

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

type SignatureIntegration struct {
	state              protoimpl.MessageState           `protogen:"open.v1"`
	Id                 string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`     // @gotags: sql:"pk"
	Name               string                           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique"` // @gotags: sql:"unique"
	Cosign             *CosignPublicKeyVerification     `protobuf:"bytes,3,opt,name=cosign,proto3" json:"cosign,omitempty"`
	CosignCertificates []*CosignCertificateVerification `protobuf:"bytes,4,rep,name=cosign_certificates,json=cosignCertificates,proto3" json:"cosign_certificates,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SignatureIntegration) Reset() {
	*x = SignatureIntegration{}
	mi := &file_storage_signature_integration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignatureIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureIntegration) ProtoMessage() {}

func (x *SignatureIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_signature_integration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureIntegration.ProtoReflect.Descriptor instead.
func (*SignatureIntegration) Descriptor() ([]byte, []int) {
	return file_storage_signature_integration_proto_rawDescGZIP(), []int{0}
}

func (x *SignatureIntegration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SignatureIntegration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignatureIntegration) GetCosign() *CosignPublicKeyVerification {
	if x != nil {
		return x.Cosign
	}
	return nil
}

func (x *SignatureIntegration) GetCosignCertificates() []*CosignCertificateVerification {
	if x != nil {
		return x.CosignCertificates
	}
	return nil
}

type CosignPublicKeyVerification struct {
	state         protoimpl.MessageState                   `protogen:"open.v1"`
	PublicKeys    []*CosignPublicKeyVerification_PublicKey `protobuf:"bytes,3,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CosignPublicKeyVerification) Reset() {
	*x = CosignPublicKeyVerification{}
	mi := &file_storage_signature_integration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CosignPublicKeyVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosignPublicKeyVerification) ProtoMessage() {}

func (x *CosignPublicKeyVerification) ProtoReflect() protoreflect.Message {
	mi := &file_storage_signature_integration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosignPublicKeyVerification.ProtoReflect.Descriptor instead.
func (*CosignPublicKeyVerification) Descriptor() ([]byte, []int) {
	return file_storage_signature_integration_proto_rawDescGZIP(), []int{1}
}

func (x *CosignPublicKeyVerification) GetPublicKeys() []*CosignPublicKeyVerification_PublicKey {
	if x != nil {
		return x.PublicKeys
	}
	return nil
}

// Holds all verification data for verifying certificates attached to cosign signatures.
// If only the certificate is given, the Fulcio trusted root chain will be assumed and verified against.
// If only the chain is given, this will be used over the Fulcio trusted root chain for verification.
// If no certificate or chain is given, the Fulcio trusted root chain will be assumed and verified against.
type CosignCertificateVerification struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PEM encoded certificate to use for verification.
	CertificatePemEnc string `protobuf:"bytes,1,opt,name=certificate_pem_enc,json=certificatePemEnc,proto3" json:"certificate_pem_enc,omitempty"`
	// PEM encoded certificate chain to use for verification.
	CertificateChainPemEnc string `protobuf:"bytes,2,opt,name=certificate_chain_pem_enc,json=certificateChainPemEnc,proto3" json:"certificate_chain_pem_enc,omitempty"`
	// Certificate OIDC issuer to verify against.
	// This supports regular expressions following the RE2 syntax: https://github.com/google/re2/wiki/Syntax.
	// In case the certificate does not specify an OIDC issuer, you may use '.*' as the OIDC issuer. However,
	// it is recommended to use Fulcio compatible certificates according to the specification:
	// https://github.com/sigstore/fulcio/blob/main/docs/certificate-specification.md.
	CertificateOidcIssuer string `protobuf:"bytes,3,opt,name=certificate_oidc_issuer,json=certificateOidcIssuer,proto3" json:"certificate_oidc_issuer,omitempty"`
	// Certificate identity to verify against.
	// This supports regular expressions following the RE2 syntax: https://github.com/google/re2/wiki/Syntax.
	// In case the certificate does not specify an identity, you may use '.*' as the identity. However, it is
	// recommended to use Fulcio compatible certificates according to the specification:
	// https://github.com/sigstore/fulcio/blob/main/docs/certificate-specification.md.
	CertificateIdentity string `protobuf:"bytes,4,opt,name=certificate_identity,json=certificateIdentity,proto3" json:"certificate_identity,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *CosignCertificateVerification) Reset() {
	*x = CosignCertificateVerification{}
	mi := &file_storage_signature_integration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CosignCertificateVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosignCertificateVerification) ProtoMessage() {}

func (x *CosignCertificateVerification) ProtoReflect() protoreflect.Message {
	mi := &file_storage_signature_integration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosignCertificateVerification.ProtoReflect.Descriptor instead.
func (*CosignCertificateVerification) Descriptor() ([]byte, []int) {
	return file_storage_signature_integration_proto_rawDescGZIP(), []int{2}
}

func (x *CosignCertificateVerification) GetCertificatePemEnc() string {
	if x != nil {
		return x.CertificatePemEnc
	}
	return ""
}

func (x *CosignCertificateVerification) GetCertificateChainPemEnc() string {
	if x != nil {
		return x.CertificateChainPemEnc
	}
	return ""
}

func (x *CosignCertificateVerification) GetCertificateOidcIssuer() string {
	if x != nil {
		return x.CertificateOidcIssuer
	}
	return ""
}

func (x *CosignCertificateVerification) GetCertificateIdentity() string {
	if x != nil {
		return x.CertificateIdentity
	}
	return ""
}

type CosignPublicKeyVerification_PublicKey struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKeyPemEnc string                 `protobuf:"bytes,2,opt,name=public_key_pem_enc,json=publicKeyPemEnc,proto3" json:"public_key_pem_enc,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CosignPublicKeyVerification_PublicKey) Reset() {
	*x = CosignPublicKeyVerification_PublicKey{}
	mi := &file_storage_signature_integration_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CosignPublicKeyVerification_PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosignPublicKeyVerification_PublicKey) ProtoMessage() {}

func (x *CosignPublicKeyVerification_PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_storage_signature_integration_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosignPublicKeyVerification_PublicKey.ProtoReflect.Descriptor instead.
func (*CosignPublicKeyVerification_PublicKey) Descriptor() ([]byte, []int) {
	return file_storage_signature_integration_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CosignPublicKeyVerification_PublicKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CosignPublicKeyVerification_PublicKey) GetPublicKeyPemEnc() string {
	if x != nil {
		return x.PublicKeyPemEnc
	}
	return ""
}

var File_storage_signature_integration_proto protoreflect.FileDescriptor

var file_storage_signature_integration_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xd1,
	0x01, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x63,
	0x6f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x63, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x57, 0x0a, 0x13, 0x63, 0x6f, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12,
	0x63, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x73, 0x1a, 0x4c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d, 0x5f, 0x65, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x6d, 0x45, 0x6e,
	0x63, 0x22, 0xf5, 0x01, 0x0a, 0x1d, 0x43, 0x6f, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x65, 0x6d, 0x5f, 0x65, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6d,
	0x45, 0x6e, 0x63, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x65, 0x6d, 0x5f, 0x65, 0x6e, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x65, 0x6d, 0x45, 0x6e, 0x63, 0x12, 0x36,
	0x0a, 0x17, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x69,
	0x64, 0x63, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4f, 0x69, 0x64, 0x63,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_storage_signature_integration_proto_rawDescOnce sync.Once
	file_storage_signature_integration_proto_rawDescData []byte
)

func file_storage_signature_integration_proto_rawDescGZIP() []byte {
	file_storage_signature_integration_proto_rawDescOnce.Do(func() {
		file_storage_signature_integration_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_signature_integration_proto_rawDesc), len(file_storage_signature_integration_proto_rawDesc)))
	})
	return file_storage_signature_integration_proto_rawDescData
}

var file_storage_signature_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_signature_integration_proto_goTypes = []any{
	(*SignatureIntegration)(nil),                  // 0: storage.SignatureIntegration
	(*CosignPublicKeyVerification)(nil),           // 1: storage.CosignPublicKeyVerification
	(*CosignCertificateVerification)(nil),         // 2: storage.CosignCertificateVerification
	(*CosignPublicKeyVerification_PublicKey)(nil), // 3: storage.CosignPublicKeyVerification.PublicKey
}
var file_storage_signature_integration_proto_depIdxs = []int32{
	1, // 0: storage.SignatureIntegration.cosign:type_name -> storage.CosignPublicKeyVerification
	2, // 1: storage.SignatureIntegration.cosign_certificates:type_name -> storage.CosignCertificateVerification
	3, // 2: storage.CosignPublicKeyVerification.public_keys:type_name -> storage.CosignPublicKeyVerification.PublicKey
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_signature_integration_proto_init() }
func file_storage_signature_integration_proto_init() {
	if File_storage_signature_integration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_signature_integration_proto_rawDesc), len(file_storage_signature_integration_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_signature_integration_proto_goTypes,
		DependencyIndexes: file_storage_signature_integration_proto_depIdxs,
		MessageInfos:      file_storage_signature_integration_proto_msgTypes,
	}.Build()
	File_storage_signature_integration_proto = out.File
	file_storage_signature_integration_proto_goTypes = nil
	file_storage_signature_integration_proto_depIdxs = nil
}
