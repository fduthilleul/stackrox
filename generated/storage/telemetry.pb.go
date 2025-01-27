// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: storage/telemetry.proto

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

type TelemetryConfiguration struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Enabled bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Deprecated: Marked as deprecated in storage/telemetry.proto.
	LastSetTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_set_time,json=lastSetTime,proto3" json:"last_set_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TelemetryConfiguration) Reset() {
	*x = TelemetryConfiguration{}
	mi := &file_storage_telemetry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TelemetryConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryConfiguration) ProtoMessage() {}

func (x *TelemetryConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_telemetry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryConfiguration.ProtoReflect.Descriptor instead.
func (*TelemetryConfiguration) Descriptor() ([]byte, []int) {
	return file_storage_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryConfiguration) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Deprecated: Marked as deprecated in storage/telemetry.proto.
func (x *TelemetryConfiguration) GetLastSetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSetTime
	}
	return nil
}

var File_storage_telemetry_proto protoreflect.FileDescriptor

var file_storage_telemetry_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x16, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x31, 0x0a, 0x19, 0x69,
	0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_telemetry_proto_rawDescOnce sync.Once
	file_storage_telemetry_proto_rawDescData []byte
)

func file_storage_telemetry_proto_rawDescGZIP() []byte {
	file_storage_telemetry_proto_rawDescOnce.Do(func() {
		file_storage_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_telemetry_proto_rawDesc), len(file_storage_telemetry_proto_rawDesc)))
	})
	return file_storage_telemetry_proto_rawDescData
}

var file_storage_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_telemetry_proto_goTypes = []any{
	(*TelemetryConfiguration)(nil), // 0: storage.TelemetryConfiguration
	(*timestamppb.Timestamp)(nil),  // 1: google.protobuf.Timestamp
}
var file_storage_telemetry_proto_depIdxs = []int32{
	1, // 0: storage.TelemetryConfiguration.last_set_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_telemetry_proto_init() }
func file_storage_telemetry_proto_init() {
	if File_storage_telemetry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_telemetry_proto_rawDesc), len(file_storage_telemetry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_telemetry_proto_goTypes,
		DependencyIndexes: file_storage_telemetry_proto_depIdxs,
		MessageInfos:      file_storage_telemetry_proto_msgTypes,
	}.Build()
	File_storage_telemetry_proto = out.File
	file_storage_telemetry_proto_goTypes = nil
	file_storage_telemetry_proto_depIdxs = nil
}
