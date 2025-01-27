// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: internalapi/sensor/admission_control_iservice.proto

package sensor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type MsgFromAdmissionControl struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MsgFromAdmissionControl) Reset() {
	*x = MsgFromAdmissionControl{}
	mi := &file_internalapi_sensor_admission_control_iservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgFromAdmissionControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgFromAdmissionControl) ProtoMessage() {}

func (x *MsgFromAdmissionControl) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_admission_control_iservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgFromAdmissionControl.ProtoReflect.Descriptor instead.
func (*MsgFromAdmissionControl) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_admission_control_iservice_proto_rawDescGZIP(), []int{0}
}

type MsgToAdmissionControl struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Msg:
	//
	//	*MsgToAdmissionControl_SettingsPush
	//	*MsgToAdmissionControl_UpdateResourceRequest
	Msg           isMsgToAdmissionControl_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MsgToAdmissionControl) Reset() {
	*x = MsgToAdmissionControl{}
	mi := &file_internalapi_sensor_admission_control_iservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgToAdmissionControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgToAdmissionControl) ProtoMessage() {}

func (x *MsgToAdmissionControl) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_admission_control_iservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgToAdmissionControl.ProtoReflect.Descriptor instead.
func (*MsgToAdmissionControl) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_admission_control_iservice_proto_rawDescGZIP(), []int{1}
}

func (x *MsgToAdmissionControl) GetMsg() isMsgToAdmissionControl_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *MsgToAdmissionControl) GetSettingsPush() *AdmissionControlSettings {
	if x != nil {
		if x, ok := x.Msg.(*MsgToAdmissionControl_SettingsPush); ok {
			return x.SettingsPush
		}
	}
	return nil
}

func (x *MsgToAdmissionControl) GetUpdateResourceRequest() *AdmCtrlUpdateResourceRequest {
	if x != nil {
		if x, ok := x.Msg.(*MsgToAdmissionControl_UpdateResourceRequest); ok {
			return x.UpdateResourceRequest
		}
	}
	return nil
}

type isMsgToAdmissionControl_Msg interface {
	isMsgToAdmissionControl_Msg()
}

type MsgToAdmissionControl_SettingsPush struct {
	SettingsPush *AdmissionControlSettings `protobuf:"bytes,1,opt,name=settings_push,json=settingsPush,proto3,oneof"`
}

type MsgToAdmissionControl_UpdateResourceRequest struct {
	UpdateResourceRequest *AdmCtrlUpdateResourceRequest `protobuf:"bytes,2,opt,name=update_resource_request,json=updateResourceRequest,proto3,oneof"`
}

func (*MsgToAdmissionControl_SettingsPush) isMsgToAdmissionControl_Msg() {}

func (*MsgToAdmissionControl_UpdateResourceRequest) isMsgToAdmissionControl_Msg() {}

var File_internalapi_sensor_admission_control_iservice_proto protoreflect.FileDescriptor

var file_internalapi_sensor_admission_control_iservice_proto_rawDesc = string([]byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x22, 0xc7, 0x01, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x41, 0x64, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x47, 0x0a, 0x0d, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x50, 0x75, 0x73, 0x68, 0x12, 0x5e, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x41,
	0x64, 0x6d, 0x43, 0x74, 0x72, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xbe, 0x01, 0x0a, 0x21,
	0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x6f,
	0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1d, 0x5a, 0x1b,
	0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x3b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_internalapi_sensor_admission_control_iservice_proto_rawDescOnce sync.Once
	file_internalapi_sensor_admission_control_iservice_proto_rawDescData []byte
)

func file_internalapi_sensor_admission_control_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_admission_control_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_admission_control_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internalapi_sensor_admission_control_iservice_proto_rawDesc), len(file_internalapi_sensor_admission_control_iservice_proto_rawDesc)))
	})
	return file_internalapi_sensor_admission_control_iservice_proto_rawDescData
}

var file_internalapi_sensor_admission_control_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internalapi_sensor_admission_control_iservice_proto_goTypes = []any{
	(*MsgFromAdmissionControl)(nil),      // 0: sensor.MsgFromAdmissionControl
	(*MsgToAdmissionControl)(nil),        // 1: sensor.MsgToAdmissionControl
	(*AdmissionControlSettings)(nil),     // 2: sensor.AdmissionControlSettings
	(*AdmCtrlUpdateResourceRequest)(nil), // 3: sensor.AdmCtrlUpdateResourceRequest
	(*AdmissionControlAlerts)(nil),       // 4: sensor.AdmissionControlAlerts
	(*emptypb.Empty)(nil),                // 5: google.protobuf.Empty
}
var file_internalapi_sensor_admission_control_iservice_proto_depIdxs = []int32{
	2, // 0: sensor.MsgToAdmissionControl.settings_push:type_name -> sensor.AdmissionControlSettings
	3, // 1: sensor.MsgToAdmissionControl.update_resource_request:type_name -> sensor.AdmCtrlUpdateResourceRequest
	0, // 2: sensor.AdmissionControlManagementService.Communicate:input_type -> sensor.MsgFromAdmissionControl
	4, // 3: sensor.AdmissionControlManagementService.PolicyAlerts:input_type -> sensor.AdmissionControlAlerts
	1, // 4: sensor.AdmissionControlManagementService.Communicate:output_type -> sensor.MsgToAdmissionControl
	5, // 5: sensor.AdmissionControlManagementService.PolicyAlerts:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_admission_control_iservice_proto_init() }
func file_internalapi_sensor_admission_control_iservice_proto_init() {
	if File_internalapi_sensor_admission_control_iservice_proto != nil {
		return
	}
	file_internalapi_sensor_admission_control_proto_init()
	file_internalapi_sensor_admission_control_iservice_proto_msgTypes[1].OneofWrappers = []any{
		(*MsgToAdmissionControl_SettingsPush)(nil),
		(*MsgToAdmissionControl_UpdateResourceRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internalapi_sensor_admission_control_iservice_proto_rawDesc), len(file_internalapi_sensor_admission_control_iservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_sensor_admission_control_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_admission_control_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_admission_control_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_admission_control_iservice_proto = out.File
	file_internalapi_sensor_admission_control_iservice_proto_goTypes = nil
	file_internalapi_sensor_admission_control_iservice_proto_depIdxs = nil
}
