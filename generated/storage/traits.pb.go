// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.3
// source: storage/traits.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EXPERIMENTAL.
// NOTE: Please refer from using MutabilityMode for the time being. It will be replaced in the future (ROX-14276).
// MutabilityMode specifies whether and how an object can be modified. Default
// is ALLOW_MUTATE and means there are no modification restrictions; this is equivalent
// to the absence of MutabilityMode specification. ALLOW_MUTATE_FORCED forbids all
// modifying operations except object removal with force bit on.
//
// Be careful when changing the state of this field. For example, modifying an
// object from ALLOW_MUTATE to ALLOW_MUTATE_FORCED is allowed but will prohibit any further
// changes to it, including modifying it back to ALLOW_MUTATE.
type Traits_MutabilityMode int32

const (
	Traits_ALLOW_MUTATE        Traits_MutabilityMode = 0
	Traits_ALLOW_MUTATE_FORCED Traits_MutabilityMode = 1
)

// Enum value maps for Traits_MutabilityMode.
var (
	Traits_MutabilityMode_name = map[int32]string{
		0: "ALLOW_MUTATE",
		1: "ALLOW_MUTATE_FORCED",
	}
	Traits_MutabilityMode_value = map[string]int32{
		"ALLOW_MUTATE":        0,
		"ALLOW_MUTATE_FORCED": 1,
	}
)

func (x Traits_MutabilityMode) Enum() *Traits_MutabilityMode {
	p := new(Traits_MutabilityMode)
	*p = x
	return p
}

func (x Traits_MutabilityMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Traits_MutabilityMode) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_traits_proto_enumTypes[0].Descriptor()
}

func (Traits_MutabilityMode) Type() protoreflect.EnumType {
	return &file_storage_traits_proto_enumTypes[0]
}

func (x Traits_MutabilityMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Traits_MutabilityMode.Descriptor instead.
func (Traits_MutabilityMode) EnumDescriptor() ([]byte, []int) {
	return file_storage_traits_proto_rawDescGZIP(), []int{0, 0}
}

// EXPERIMENTAL.
// visibility allows to specify whether the object should be visible for certain APIs.
type Traits_Visibility int32

const (
	Traits_VISIBLE Traits_Visibility = 0
	Traits_HIDDEN  Traits_Visibility = 1
)

// Enum value maps for Traits_Visibility.
var (
	Traits_Visibility_name = map[int32]string{
		0: "VISIBLE",
		1: "HIDDEN",
	}
	Traits_Visibility_value = map[string]int32{
		"VISIBLE": 0,
		"HIDDEN":  1,
	}
)

func (x Traits_Visibility) Enum() *Traits_Visibility {
	p := new(Traits_Visibility)
	*p = x
	return p
}

func (x Traits_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Traits_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_traits_proto_enumTypes[1].Descriptor()
}

func (Traits_Visibility) Type() protoreflect.EnumType {
	return &file_storage_traits_proto_enumTypes[1]
}

func (x Traits_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Traits_Visibility.Descriptor instead.
func (Traits_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_storage_traits_proto_rawDescGZIP(), []int{0, 1}
}

// Origin specifies the origin of an object.
// Objects can have four different origins:
// - IMPERATIVE: the object was created via the API. This is assumed by default.
// - DEFAULT: the object is a default object, such as default roles, access scopes etc.
// - DECLARATIVE: the object is created via declarative configuration.
// - DECLARATIVE_ORPHANED: the object is created via declarative configuration and then unsuccessfully deleted(for example, because it is referenced by another object)
// Based on the origin, different rules apply to the objects.
// Objects with the DECLARATIVE origin are not allowed to be modified via API, only via declarative configuration.
// Additionally, they may not reference objects with the IMPERATIVE origin.
// Objects with the DEFAULT origin are not allowed to be modified via either API or declarative configuration.
// They may be referenced by all other objects.
// Objects with the IMPERATIVE origin are allowed to be modified via API, not via declarative configuration.
// They may reference all other objects.
// Objects with the DECLARATIVE_ORPHANED origin are not allowed to be modified via either API or declarative configuration.
// DECLARATIVE_ORPHANED resource can become DECLARATIVE again if it is redefined in declarative configuration.
// Objects with this origin will be cleaned up from the system immediately after they are not referenced by other resources anymore.
// They may be referenced by all other objects.
type Traits_Origin int32

const (
	Traits_IMPERATIVE           Traits_Origin = 0
	Traits_DEFAULT              Traits_Origin = 1
	Traits_DECLARATIVE          Traits_Origin = 2
	Traits_DECLARATIVE_ORPHANED Traits_Origin = 3
)

// Enum value maps for Traits_Origin.
var (
	Traits_Origin_name = map[int32]string{
		0: "IMPERATIVE",
		1: "DEFAULT",
		2: "DECLARATIVE",
		3: "DECLARATIVE_ORPHANED",
	}
	Traits_Origin_value = map[string]int32{
		"IMPERATIVE":           0,
		"DEFAULT":              1,
		"DECLARATIVE":          2,
		"DECLARATIVE_ORPHANED": 3,
	}
)

func (x Traits_Origin) Enum() *Traits_Origin {
	p := new(Traits_Origin)
	*p = x
	return p
}

func (x Traits_Origin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Traits_Origin) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_traits_proto_enumTypes[2].Descriptor()
}

func (Traits_Origin) Type() protoreflect.EnumType {
	return &file_storage_traits_proto_enumTypes[2]
}

func (x Traits_Origin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Traits_Origin.Descriptor instead.
func (Traits_Origin) EnumDescriptor() ([]byte, []int) {
	return file_storage_traits_proto_rawDescGZIP(), []int{0, 2}
}

type Traits struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	MutabilityMode Traits_MutabilityMode  `protobuf:"varint,1,opt,name=mutability_mode,json=mutabilityMode,proto3,enum=storage.Traits_MutabilityMode" json:"mutability_mode,omitempty"`
	Visibility     Traits_Visibility      `protobuf:"varint,2,opt,name=visibility,proto3,enum=storage.Traits_Visibility" json:"visibility,omitempty"`
	Origin         Traits_Origin          `protobuf:"varint,3,opt,name=origin,proto3,enum=storage.Traits_Origin" json:"origin,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Traits) Reset() {
	*x = Traits{}
	mi := &file_storage_traits_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Traits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traits) ProtoMessage() {}

func (x *Traits) ProtoReflect() protoreflect.Message {
	mi := &file_storage_traits_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traits.ProtoReflect.Descriptor instead.
func (*Traits) Descriptor() ([]byte, []int) {
	return file_storage_traits_proto_rawDescGZIP(), []int{0}
}

func (x *Traits) GetMutabilityMode() Traits_MutabilityMode {
	if x != nil {
		return x.MutabilityMode
	}
	return Traits_ALLOW_MUTATE
}

func (x *Traits) GetVisibility() Traits_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Traits_VISIBLE
}

func (x *Traits) GetOrigin() Traits_Origin {
	if x != nil {
		return x.Origin
	}
	return Traits_IMPERATIVE
}

var File_storage_traits_proto protoreflect.FileDescriptor

var file_storage_traits_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22,
	0xf3, 0x02, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0f, 0x6d, 0x75,
	0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x0e, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x2e, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22,
	0x3b, 0x0a, 0x0e, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54,
	0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x4d, 0x55, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x44, 0x10, 0x01, 0x22, 0x25, 0x0a, 0x0a,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49,
	0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x49, 0x44, 0x44, 0x45,
	0x4e, 0x10, 0x01, 0x22, 0x50, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45,
	0x43, 0x4c, 0x41, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x44,
	0x45, 0x43, 0x4c, 0x41, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x4f, 0x52, 0x50, 0x48, 0x41,
	0x4e, 0x45, 0x44, 0x10, 0x03, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_traits_proto_rawDescOnce sync.Once
	file_storage_traits_proto_rawDescData = file_storage_traits_proto_rawDesc
)

func file_storage_traits_proto_rawDescGZIP() []byte {
	file_storage_traits_proto_rawDescOnce.Do(func() {
		file_storage_traits_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_traits_proto_rawDescData)
	})
	return file_storage_traits_proto_rawDescData
}

var file_storage_traits_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_storage_traits_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_traits_proto_goTypes = []any{
	(Traits_MutabilityMode)(0), // 0: storage.Traits.MutabilityMode
	(Traits_Visibility)(0),     // 1: storage.Traits.Visibility
	(Traits_Origin)(0),         // 2: storage.Traits.Origin
	(*Traits)(nil),             // 3: storage.Traits
}
var file_storage_traits_proto_depIdxs = []int32{
	0, // 0: storage.Traits.mutability_mode:type_name -> storage.Traits.MutabilityMode
	1, // 1: storage.Traits.visibility:type_name -> storage.Traits.Visibility
	2, // 2: storage.Traits.origin:type_name -> storage.Traits.Origin
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_traits_proto_init() }
func file_storage_traits_proto_init() {
	if File_storage_traits_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_traits_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_traits_proto_goTypes,
		DependencyIndexes: file_storage_traits_proto_depIdxs,
		EnumInfos:         file_storage_traits_proto_enumTypes,
		MessageInfos:      file_storage_traits_proto_msgTypes,
	}.Build()
	File_storage_traits_proto = out.File
	file_storage_traits_proto_rawDesc = nil
	file_storage_traits_proto_goTypes = nil
	file_storage_traits_proto_depIdxs = nil
}
