// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: api/v1/signal.proto

package v1

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

// Signal message tracks process and network activity.
// Specifically, process launches and network connects/accepts.
type Signal struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Signal:
	//
	//	*Signal_ProcessSignal
	Signal        isSignal_Signal `protobuf_oneof:"signal"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Signal) Reset() {
	*x = Signal{}
	mi := &file_api_v1_signal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_signal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_api_v1_signal_proto_rawDescGZIP(), []int{0}
}

func (x *Signal) GetSignal() isSignal_Signal {
	if x != nil {
		return x.Signal
	}
	return nil
}

func (x *Signal) GetProcessSignal() *storage.ProcessSignal {
	if x != nil {
		if x, ok := x.Signal.(*Signal_ProcessSignal); ok {
			return x.ProcessSignal
		}
	}
	return nil
}

type isSignal_Signal interface {
	isSignal_Signal()
}

type Signal_ProcessSignal struct {
	ProcessSignal *storage.ProcessSignal `protobuf:"bytes,1,opt,name=process_signal,json=processSignal,proto3,oneof"`
}

func (*Signal_ProcessSignal) isSignal_Signal() {}

var File_api_v1_signal_proto protoreflect.FileDescriptor

var file_api_v1_signal_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x42,
	0x2a, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_api_v1_signal_proto_rawDescOnce sync.Once
	file_api_v1_signal_proto_rawDescData []byte
)

func file_api_v1_signal_proto_rawDescGZIP() []byte {
	file_api_v1_signal_proto_rawDescOnce.Do(func() {
		file_api_v1_signal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_signal_proto_rawDesc), len(file_api_v1_signal_proto_rawDesc)))
	})
	return file_api_v1_signal_proto_rawDescData
}

var file_api_v1_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_signal_proto_goTypes = []any{
	(*Signal)(nil),                // 0: v1.Signal
	(*storage.ProcessSignal)(nil), // 1: storage.ProcessSignal
}
var file_api_v1_signal_proto_depIdxs = []int32{
	1, // 0: v1.Signal.process_signal:type_name -> storage.ProcessSignal
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_signal_proto_init() }
func file_api_v1_signal_proto_init() {
	if File_api_v1_signal_proto != nil {
		return
	}
	file_api_v1_signal_proto_msgTypes[0].OneofWrappers = []any{
		(*Signal_ProcessSignal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_signal_proto_rawDesc), len(file_api_v1_signal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_signal_proto_goTypes,
		DependencyIndexes: file_api_v1_signal_proto_depIdxs,
		MessageInfos:      file_api_v1_signal_proto_msgTypes,
	}.Build()
	File_api_v1_signal_proto = out.File
	file_api_v1_signal_proto_goTypes = nil
	file_api_v1_signal_proto_depIdxs = nil
}
