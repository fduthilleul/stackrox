// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: internalapi/sensor/image_iservice.proto

package sensor

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

type GetImageRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Image         *storage.ContainerImage `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Namespace     string                  `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ScanInline    bool                    `protobuf:"varint,2,opt,name=scan_inline,json=scanInline,proto3" json:"scan_inline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	mi := &file_internalapi_sensor_image_iservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_image_iservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_image_iservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetImageRequest) GetImage() *storage.ContainerImage {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GetImageRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetImageRequest) GetScanInline() bool {
	if x != nil {
		return x.ScanInline
	}
	return false
}

type GetImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Image         *storage.Image         `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	mi := &file_internalapi_sensor_image_iservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_image_iservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_image_iservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetImageResponse) GetImage() *storage.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_internalapi_sensor_image_iservice_proto protoreflect.FileDescriptor

var file_internalapi_sensor_image_iservice_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x1a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x49, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0x4d, 0x0a, 0x0c, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x3b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_internalapi_sensor_image_iservice_proto_rawDescOnce sync.Once
	file_internalapi_sensor_image_iservice_proto_rawDescData []byte
)

func file_internalapi_sensor_image_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_image_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_image_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internalapi_sensor_image_iservice_proto_rawDesc), len(file_internalapi_sensor_image_iservice_proto_rawDesc)))
	})
	return file_internalapi_sensor_image_iservice_proto_rawDescData
}

var file_internalapi_sensor_image_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internalapi_sensor_image_iservice_proto_goTypes = []any{
	(*GetImageRequest)(nil),        // 0: sensor.GetImageRequest
	(*GetImageResponse)(nil),       // 1: sensor.GetImageResponse
	(*storage.ContainerImage)(nil), // 2: storage.ContainerImage
	(*storage.Image)(nil),          // 3: storage.Image
}
var file_internalapi_sensor_image_iservice_proto_depIdxs = []int32{
	2, // 0: sensor.GetImageRequest.image:type_name -> storage.ContainerImage
	3, // 1: sensor.GetImageResponse.image:type_name -> storage.Image
	0, // 2: sensor.ImageService.GetImage:input_type -> sensor.GetImageRequest
	1, // 3: sensor.ImageService.GetImage:output_type -> sensor.GetImageResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_image_iservice_proto_init() }
func file_internalapi_sensor_image_iservice_proto_init() {
	if File_internalapi_sensor_image_iservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internalapi_sensor_image_iservice_proto_rawDesc), len(file_internalapi_sensor_image_iservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_sensor_image_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_image_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_image_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_image_iservice_proto = out.File
	file_internalapi_sensor_image_iservice_proto_goTypes = nil
	file_internalapi_sensor_image_iservice_proto_depIdxs = nil
}
