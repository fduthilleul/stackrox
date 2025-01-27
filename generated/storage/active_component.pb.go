// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: storage/active_component.proto

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

// Next available tag: 3
type ActiveComponent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// base 64 encoded Deployment:ActiveComponent ids.
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Image Sha,hidden" sql:"pk"`                                         // @gotags: search:"Image Sha,hidden" sql:"pk"
	DeploymentId string `protobuf:"bytes,3,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty" search:"Deployment ID,hidden" sql:"fk(Deployment:id),no-fk-constraint,directional,index=hash,type(uuid)"` // @gotags: search:"Deployment ID,hidden" sql:"fk(Deployment:id),no-fk-constraint,directional,index=hash,type(uuid)"
	ComponentId  string `protobuf:"bytes,4,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty" search:"Component ID,hidden" sql:"fk(ImageComponent:id),no-fk-constraint,directional"`    // @gotags: search:"Component ID,hidden" sql:"fk(ImageComponent:id),no-fk-constraint,directional"
	// Map from container name to the active context of an edge.
	DEPRECATEDActiveContexts map[string]*ActiveComponent_ActiveContext `protobuf:"bytes,2,rep,name=DEPRECATED_active_contexts,json=DEPRECATEDActiveContexts,proto3" json:"DEPRECATED_active_contexts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" search:"-"` // @gotags: search:"-"
	ActiveContextsSlice      []*ActiveComponent_ActiveContext          `protobuf:"bytes,5,rep,name=active_contexts_slice,json=activeContextsSlice,proto3" json:"active_contexts_slice,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ActiveComponent) Reset() {
	*x = ActiveComponent{}
	mi := &file_storage_active_component_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveComponent) ProtoMessage() {}

func (x *ActiveComponent) ProtoReflect() protoreflect.Message {
	mi := &file_storage_active_component_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveComponent.ProtoReflect.Descriptor instead.
func (*ActiveComponent) Descriptor() ([]byte, []int) {
	return file_storage_active_component_proto_rawDescGZIP(), []int{0}
}

func (x *ActiveComponent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActiveComponent) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ActiveComponent) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

func (x *ActiveComponent) GetDEPRECATEDActiveContexts() map[string]*ActiveComponent_ActiveContext {
	if x != nil {
		return x.DEPRECATEDActiveContexts
	}
	return nil
}

func (x *ActiveComponent) GetActiveContextsSlice() []*ActiveComponent_ActiveContext {
	if x != nil {
		return x.ActiveContextsSlice
	}
	return nil
}

// Represent a context of the active edge.
type ActiveComponent_ActiveContext struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContainerName string                 `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty" search:"Container Name,hidden"` // @gotags: search:"Container Name,hidden"
	ImageId       string                 `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty" search:"Image Sha,hidden"`                   // @gotags: search:"Image Sha,hidden"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActiveComponent_ActiveContext) Reset() {
	*x = ActiveComponent_ActiveContext{}
	mi := &file_storage_active_component_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveComponent_ActiveContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveComponent_ActiveContext) ProtoMessage() {}

func (x *ActiveComponent_ActiveContext) ProtoReflect() protoreflect.Message {
	mi := &file_storage_active_component_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveComponent_ActiveContext.ProtoReflect.Descriptor instead.
func (*ActiveComponent_ActiveContext) Descriptor() ([]byte, []int) {
	return file_storage_active_component_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ActiveComponent_ActiveContext) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ActiveComponent_ActiveContext) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

var File_storage_active_component_proto protoreflect.FileDescriptor

var file_storage_active_component_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x83, 0x04, 0x0a, 0x0f, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x74, 0x0a, 0x1a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x18, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x5a, 0x0a, 0x15, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x5f, 0x73,
	0x6c, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x1a, 0x51, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x1a, 0x73, 0x0a, 0x1d, 0x44, 0x45,
	0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_active_component_proto_rawDescOnce sync.Once
	file_storage_active_component_proto_rawDescData []byte
)

func file_storage_active_component_proto_rawDescGZIP() []byte {
	file_storage_active_component_proto_rawDescOnce.Do(func() {
		file_storage_active_component_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_active_component_proto_rawDesc), len(file_storage_active_component_proto_rawDesc)))
	})
	return file_storage_active_component_proto_rawDescData
}

var file_storage_active_component_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storage_active_component_proto_goTypes = []any{
	(*ActiveComponent)(nil),               // 0: storage.ActiveComponent
	(*ActiveComponent_ActiveContext)(nil), // 1: storage.ActiveComponent.ActiveContext
	nil,                                   // 2: storage.ActiveComponent.DEPRECATEDActiveContextsEntry
}
var file_storage_active_component_proto_depIdxs = []int32{
	2, // 0: storage.ActiveComponent.DEPRECATED_active_contexts:type_name -> storage.ActiveComponent.DEPRECATEDActiveContextsEntry
	1, // 1: storage.ActiveComponent.active_contexts_slice:type_name -> storage.ActiveComponent.ActiveContext
	1, // 2: storage.ActiveComponent.DEPRECATEDActiveContextsEntry.value:type_name -> storage.ActiveComponent.ActiveContext
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_active_component_proto_init() }
func file_storage_active_component_proto_init() {
	if File_storage_active_component_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_active_component_proto_rawDesc), len(file_storage_active_component_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_active_component_proto_goTypes,
		DependencyIndexes: file_storage_active_component_proto_depIdxs,
		MessageInfos:      file_storage_active_component_proto_msgTypes,
	}.Build()
	File_storage_active_component_proto = out.File
	file_storage_active_component_proto_goTypes = nil
	file_storage_active_component_proto_depIdxs = nil
}
