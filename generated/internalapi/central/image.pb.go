// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: internalapi/central/image.proto

package central

import (
	storage "github.com/stackrox/rox/generated/storage"
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

// ScanImage is sent to sensor to request a local scan of an image.
type ScanImage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// request id is used to map scan results to a waiting goroutine.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// full image name ie: <registry>/something/nginx:1.2.3.
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	// force will cause central and sensor caches to be ignored.
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	// namespace is used by sensor to pull additional secrets for registry authentication.
	Namespace     string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScanImage) Reset() {
	*x = ScanImage{}
	mi := &file_internalapi_central_image_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScanImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanImage) ProtoMessage() {}

func (x *ScanImage) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_image_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanImage.ProtoReflect.Descriptor instead.
func (*ScanImage) Descriptor() ([]byte, []int) {
	return file_internalapi_central_image_proto_rawDescGZIP(), []int{0}
}

func (x *ScanImage) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ScanImage) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *ScanImage) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *ScanImage) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// ImageIntegrations contains a list of integrations sensor should upsert and/or
// delete from its internal store.
type ImageIntegrations struct {
	state                 protoimpl.MessageState      `protogen:"open.v1"`
	UpdatedIntegrations   []*storage.ImageIntegration `protobuf:"bytes,1,rep,name=updated_integrations,json=updatedIntegrations,proto3" json:"updated_integrations,omitempty"`
	DeletedIntegrationIds []string                    `protobuf:"bytes,2,rep,name=deleted_integration_ids,json=deletedIntegrationIds,proto3" json:"deleted_integration_ids,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ImageIntegrations) Reset() {
	*x = ImageIntegrations{}
	mi := &file_internalapi_central_image_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageIntegrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageIntegrations) ProtoMessage() {}

func (x *ImageIntegrations) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_image_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageIntegrations.ProtoReflect.Descriptor instead.
func (*ImageIntegrations) Descriptor() ([]byte, []int) {
	return file_internalapi_central_image_proto_rawDescGZIP(), []int{1}
}

func (x *ImageIntegrations) GetUpdatedIntegrations() []*storage.ImageIntegration {
	if x != nil {
		return x.UpdatedIntegrations
	}
	return nil
}

func (x *ImageIntegrations) GetDeletedIntegrationIds() []string {
	if x != nil {
		return x.DeletedIntegrationIds
	}
	return nil
}

var File_internalapi_central_image_proto protoreflect.FileDescriptor

var file_internalapi_central_image_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x1a, 0x1f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x09, 0x53,
	0x63, 0x61, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4c, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36,
	0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x3b,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internalapi_central_image_proto_rawDescOnce sync.Once
	file_internalapi_central_image_proto_rawDescData []byte
)

func file_internalapi_central_image_proto_rawDescGZIP() []byte {
	file_internalapi_central_image_proto_rawDescOnce.Do(func() {
		file_internalapi_central_image_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internalapi_central_image_proto_rawDesc), len(file_internalapi_central_image_proto_rawDesc)))
	})
	return file_internalapi_central_image_proto_rawDescData
}

var file_internalapi_central_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internalapi_central_image_proto_goTypes = []any{
	(*ScanImage)(nil),                // 0: central.ScanImage
	(*ImageIntegrations)(nil),        // 1: central.ImageIntegrations
	(*storage.ImageIntegration)(nil), // 2: storage.ImageIntegration
}
var file_internalapi_central_image_proto_depIdxs = []int32{
	2, // 0: central.ImageIntegrations.updated_integrations:type_name -> storage.ImageIntegration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_central_image_proto_init() }
func file_internalapi_central_image_proto_init() {
	if File_internalapi_central_image_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internalapi_central_image_proto_rawDesc), len(file_internalapi_central_image_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_image_proto_goTypes,
		DependencyIndexes: file_internalapi_central_image_proto_depIdxs,
		MessageInfos:      file_internalapi_central_image_proto_msgTypes,
	}.Build()
	File_internalapi_central_image_proto = out.File
	file_internalapi_central_image_proto_goTypes = nil
	file_internalapi_central_image_proto_depIdxs = nil
}
