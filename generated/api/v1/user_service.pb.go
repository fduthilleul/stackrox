// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.3
// source: api/v1/user_service.proto

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

// Next Tag: 2
type GetUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*storage.User        `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	mi := &file_api_v1_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUsersResponse) GetUsers() []*storage.User {
	if x != nil {
		return x.Users
	}
	return nil
}

// UserAttributeTuple descript the auth:key:value tuple that decides group membership.
// Next Tag: 4
type UserAttributeTuple struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	AuthProviderId string                 `protobuf:"bytes,1,opt,name=auth_provider_id,json=authProviderId,proto3" json:"auth_provider_id,omitempty"`
	Key            string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value          string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UserAttributeTuple) Reset() {
	*x = UserAttributeTuple{}
	mi := &file_api_v1_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAttributeTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttributeTuple) ProtoMessage() {}

func (x *UserAttributeTuple) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttributeTuple.ProtoReflect.Descriptor instead.
func (*UserAttributeTuple) Descriptor() ([]byte, []int) {
	return file_api_v1_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserAttributeTuple) GetAuthProviderId() string {
	if x != nil {
		return x.AuthProviderId
	}
	return ""
}

func (x *UserAttributeTuple) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UserAttributeTuple) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Next Tag: 2
type GetUsersAttributesResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UsersAttributes []*UserAttributeTuple  `protobuf:"bytes,1,rep,name=users_attributes,json=usersAttributes,proto3" json:"users_attributes,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetUsersAttributesResponse) Reset() {
	*x = GetUsersAttributesResponse{}
	mi := &file_api_v1_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersAttributesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersAttributesResponse) ProtoMessage() {}

func (x *GetUsersAttributesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersAttributesResponse.ProtoReflect.Descriptor instead.
func (*GetUsersAttributesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersAttributesResponse) GetUsersAttributes() []*UserAttributeTuple {
	if x != nil {
		return x.UsersAttributes
	}
	return nil
}

var File_api_v1_user_service_proto protoreflect.FileDescriptor

var file_api_v1_user_service_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5f, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x32, 0xef, 0x01, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x5c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x27,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_api_v1_user_service_proto_rawDescOnce sync.Once
	file_api_v1_user_service_proto_rawDescData []byte
)

func file_api_v1_user_service_proto_rawDescGZIP() []byte {
	file_api_v1_user_service_proto_rawDescOnce.Do(func() {
		file_api_v1_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_user_service_proto_rawDesc), len(file_api_v1_user_service_proto_rawDesc)))
	})
	return file_api_v1_user_service_proto_rawDescData
}

var file_api_v1_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_user_service_proto_goTypes = []any{
	(*GetUsersResponse)(nil),           // 0: v1.GetUsersResponse
	(*UserAttributeTuple)(nil),         // 1: v1.UserAttributeTuple
	(*GetUsersAttributesResponse)(nil), // 2: v1.GetUsersAttributesResponse
	(*storage.User)(nil),               // 3: storage.User
	(*Empty)(nil),                      // 4: v1.Empty
	(*ResourceByID)(nil),               // 5: v1.ResourceByID
}
var file_api_v1_user_service_proto_depIdxs = []int32{
	3, // 0: v1.GetUsersResponse.users:type_name -> storage.User
	1, // 1: v1.GetUsersAttributesResponse.users_attributes:type_name -> v1.UserAttributeTuple
	4, // 2: v1.UserService.GetUsers:input_type -> v1.Empty
	5, // 3: v1.UserService.GetUser:input_type -> v1.ResourceByID
	4, // 4: v1.UserService.GetUsersAttributes:input_type -> v1.Empty
	0, // 5: v1.UserService.GetUsers:output_type -> v1.GetUsersResponse
	3, // 6: v1.UserService.GetUser:output_type -> storage.User
	2, // 7: v1.UserService.GetUsersAttributes:output_type -> v1.GetUsersAttributesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_user_service_proto_init() }
func file_api_v1_user_service_proto_init() {
	if File_api_v1_user_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_user_service_proto_rawDesc), len(file_api_v1_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_user_service_proto_goTypes,
		DependencyIndexes: file_api_v1_user_service_proto_depIdxs,
		MessageInfos:      file_api_v1_user_service_proto_msgTypes,
	}.Build()
	File_api_v1_user_service_proto = out.File
	file_api_v1_user_service_proto_goTypes = nil
	file_api_v1_user_service_proto_depIdxs = nil
}
