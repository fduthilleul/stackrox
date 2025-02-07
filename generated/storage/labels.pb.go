// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: storage/labels.proto

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

type LabelSelector_Operator int32

const (
	LabelSelector_UNKNOWN    LabelSelector_Operator = 0
	LabelSelector_IN         LabelSelector_Operator = 1
	LabelSelector_NOT_IN     LabelSelector_Operator = 2
	LabelSelector_EXISTS     LabelSelector_Operator = 3
	LabelSelector_NOT_EXISTS LabelSelector_Operator = 4
)

// Enum value maps for LabelSelector_Operator.
var (
	LabelSelector_Operator_name = map[int32]string{
		0: "UNKNOWN",
		1: "IN",
		2: "NOT_IN",
		3: "EXISTS",
		4: "NOT_EXISTS",
	}
	LabelSelector_Operator_value = map[string]int32{
		"UNKNOWN":    0,
		"IN":         1,
		"NOT_IN":     2,
		"EXISTS":     3,
		"NOT_EXISTS": 4,
	}
)

func (x LabelSelector_Operator) Enum() *LabelSelector_Operator {
	p := new(LabelSelector_Operator)
	*p = x
	return p
}

func (x LabelSelector_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabelSelector_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_labels_proto_enumTypes[0].Descriptor()
}

func (LabelSelector_Operator) Type() protoreflect.EnumType {
	return &file_storage_labels_proto_enumTypes[0]
}

func (x LabelSelector_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabelSelector_Operator.Descriptor instead.
func (LabelSelector_Operator) EnumDescriptor() ([]byte, []int) {
	return file_storage_labels_proto_rawDescGZIP(), []int{0, 0}
}

type SetBasedLabelSelector_Operator int32

const (
	SetBasedLabelSelector_UNKNOWN    SetBasedLabelSelector_Operator = 0
	SetBasedLabelSelector_IN         SetBasedLabelSelector_Operator = 1
	SetBasedLabelSelector_NOT_IN     SetBasedLabelSelector_Operator = 2
	SetBasedLabelSelector_EXISTS     SetBasedLabelSelector_Operator = 3
	SetBasedLabelSelector_NOT_EXISTS SetBasedLabelSelector_Operator = 4
)

// Enum value maps for SetBasedLabelSelector_Operator.
var (
	SetBasedLabelSelector_Operator_name = map[int32]string{
		0: "UNKNOWN",
		1: "IN",
		2: "NOT_IN",
		3: "EXISTS",
		4: "NOT_EXISTS",
	}
	SetBasedLabelSelector_Operator_value = map[string]int32{
		"UNKNOWN":    0,
		"IN":         1,
		"NOT_IN":     2,
		"EXISTS":     3,
		"NOT_EXISTS": 4,
	}
)

func (x SetBasedLabelSelector_Operator) Enum() *SetBasedLabelSelector_Operator {
	p := new(SetBasedLabelSelector_Operator)
	*p = x
	return p
}

func (x SetBasedLabelSelector_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetBasedLabelSelector_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_labels_proto_enumTypes[1].Descriptor()
}

func (SetBasedLabelSelector_Operator) Type() protoreflect.EnumType {
	return &file_storage_labels_proto_enumTypes[1]
}

func (x SetBasedLabelSelector_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetBasedLabelSelector_Operator.Descriptor instead.
func (SetBasedLabelSelector_Operator) EnumDescriptor() ([]byte, []int) {
	return file_storage_labels_proto_rawDescGZIP(), []int{1, 0}
}

// Label selector components are joined with logical AND, see
//
//	https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
//
// Next available tag: 3
type LabelSelector struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This is actually a oneof, but we can't make it one due to backwards
	// compatibility constraints.
	MatchLabels   map[string]string            `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Requirements  []*LabelSelector_Requirement `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelSelector) Reset() {
	*x = LabelSelector{}
	mi := &file_storage_labels_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSelector) ProtoMessage() {}

func (x *LabelSelector) ProtoReflect() protoreflect.Message {
	mi := &file_storage_labels_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSelector.ProtoReflect.Descriptor instead.
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return file_storage_labels_proto_rawDescGZIP(), []int{0}
}

func (x *LabelSelector) GetMatchLabels() map[string]string {
	if x != nil {
		return x.MatchLabels
	}
	return nil
}

func (x *LabelSelector) GetRequirements() []*LabelSelector_Requirement {
	if x != nil {
		return x.Requirements
	}
	return nil
}

// SetBasedLabelSelector only allows set-based label requirements.
//
// Next available tag: 3
type SetBasedLabelSelector struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	Requirements  []*SetBasedLabelSelector_Requirement `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetBasedLabelSelector) Reset() {
	*x = SetBasedLabelSelector{}
	mi := &file_storage_labels_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetBasedLabelSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBasedLabelSelector) ProtoMessage() {}

func (x *SetBasedLabelSelector) ProtoReflect() protoreflect.Message {
	mi := &file_storage_labels_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBasedLabelSelector.ProtoReflect.Descriptor instead.
func (*SetBasedLabelSelector) Descriptor() ([]byte, []int) {
	return file_storage_labels_proto_rawDescGZIP(), []int{1}
}

func (x *SetBasedLabelSelector) GetRequirements() []*SetBasedLabelSelector_Requirement {
	if x != nil {
		return x.Requirements
	}
	return nil
}

// Next available tag: 4
type LabelSelector_Requirement struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op            LabelSelector_Operator `protobuf:"varint,2,opt,name=op,proto3,enum=storage.LabelSelector_Operator" json:"op,omitempty"`
	Values        []string               `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelSelector_Requirement) Reset() {
	*x = LabelSelector_Requirement{}
	mi := &file_storage_labels_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelSelector_Requirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSelector_Requirement) ProtoMessage() {}

func (x *LabelSelector_Requirement) ProtoReflect() protoreflect.Message {
	mi := &file_storage_labels_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSelector_Requirement.ProtoReflect.Descriptor instead.
func (*LabelSelector_Requirement) Descriptor() ([]byte, []int) {
	return file_storage_labels_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LabelSelector_Requirement) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LabelSelector_Requirement) GetOp() LabelSelector_Operator {
	if x != nil {
		return x.Op
	}
	return LabelSelector_UNKNOWN
}

func (x *LabelSelector_Requirement) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Next available tag: 4
type SetBasedLabelSelector_Requirement struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Key           string                         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op            SetBasedLabelSelector_Operator `protobuf:"varint,2,opt,name=op,proto3,enum=storage.SetBasedLabelSelector_Operator" json:"op,omitempty"`
	Values        []string                       `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetBasedLabelSelector_Requirement) Reset() {
	*x = SetBasedLabelSelector_Requirement{}
	mi := &file_storage_labels_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetBasedLabelSelector_Requirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBasedLabelSelector_Requirement) ProtoMessage() {}

func (x *SetBasedLabelSelector_Requirement) ProtoReflect() protoreflect.Message {
	mi := &file_storage_labels_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBasedLabelSelector_Requirement.ProtoReflect.Descriptor instead.
func (*SetBasedLabelSelector_Requirement) Descriptor() ([]byte, []int) {
	return file_storage_labels_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SetBasedLabelSelector_Requirement) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetBasedLabelSelector_Requirement) GetOp() SetBasedLabelSelector_Operator {
	if x != nil {
		return x.Op
	}
	return SetBasedLabelSelector_UNKNOWN
}

func (x *SetBasedLabelSelector_Requirement) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_storage_labels_proto protoreflect.FileDescriptor

var file_storage_labels_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22,
	0x96, 0x03, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x4a, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x46, 0x0a,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x68, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a,
	0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x47, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x04, 0x22, 0xa8, 0x02, 0x0a, 0x15, 0x53, 0x65, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x4e, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x70, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x61, 0x73,
	0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x04, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_labels_proto_rawDescOnce sync.Once
	file_storage_labels_proto_rawDescData []byte
)

func file_storage_labels_proto_rawDescGZIP() []byte {
	file_storage_labels_proto_rawDescOnce.Do(func() {
		file_storage_labels_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_labels_proto_rawDesc), len(file_storage_labels_proto_rawDesc)))
	})
	return file_storage_labels_proto_rawDescData
}

var file_storage_labels_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_storage_labels_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storage_labels_proto_goTypes = []any{
	(LabelSelector_Operator)(0),               // 0: storage.LabelSelector.Operator
	(SetBasedLabelSelector_Operator)(0),       // 1: storage.SetBasedLabelSelector.Operator
	(*LabelSelector)(nil),                     // 2: storage.LabelSelector
	(*SetBasedLabelSelector)(nil),             // 3: storage.SetBasedLabelSelector
	(*LabelSelector_Requirement)(nil),         // 4: storage.LabelSelector.Requirement
	nil,                                       // 5: storage.LabelSelector.MatchLabelsEntry
	(*SetBasedLabelSelector_Requirement)(nil), // 6: storage.SetBasedLabelSelector.Requirement
}
var file_storage_labels_proto_depIdxs = []int32{
	5, // 0: storage.LabelSelector.match_labels:type_name -> storage.LabelSelector.MatchLabelsEntry
	4, // 1: storage.LabelSelector.requirements:type_name -> storage.LabelSelector.Requirement
	6, // 2: storage.SetBasedLabelSelector.requirements:type_name -> storage.SetBasedLabelSelector.Requirement
	0, // 3: storage.LabelSelector.Requirement.op:type_name -> storage.LabelSelector.Operator
	1, // 4: storage.SetBasedLabelSelector.Requirement.op:type_name -> storage.SetBasedLabelSelector.Operator
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_storage_labels_proto_init() }
func file_storage_labels_proto_init() {
	if File_storage_labels_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_labels_proto_rawDesc), len(file_storage_labels_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_labels_proto_goTypes,
		DependencyIndexes: file_storage_labels_proto_depIdxs,
		EnumInfos:         file_storage_labels_proto_enumTypes,
		MessageInfos:      file_storage_labels_proto_msgTypes,
	}.Build()
	File_storage_labels_proto = out.File
	file_storage_labels_proto_goTypes = nil
	file_storage_labels_proto_depIdxs = nil
}
