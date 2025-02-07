// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: storage/schedule.proto

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

type Schedule_IntervalType int32

const (
	Schedule_UNSET   Schedule_IntervalType = 0
	Schedule_DAILY   Schedule_IntervalType = 1
	Schedule_WEEKLY  Schedule_IntervalType = 2
	Schedule_MONTHLY Schedule_IntervalType = 3
)

// Enum value maps for Schedule_IntervalType.
var (
	Schedule_IntervalType_name = map[int32]string{
		0: "UNSET",
		1: "DAILY",
		2: "WEEKLY",
		3: "MONTHLY",
	}
	Schedule_IntervalType_value = map[string]int32{
		"UNSET":   0,
		"DAILY":   1,
		"WEEKLY":  2,
		"MONTHLY": 3,
	}
)

func (x Schedule_IntervalType) Enum() *Schedule_IntervalType {
	p := new(Schedule_IntervalType)
	*p = x
	return p
}

func (x Schedule_IntervalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule_IntervalType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_schedule_proto_enumTypes[0].Descriptor()
}

func (Schedule_IntervalType) Type() protoreflect.EnumType {
	return &file_storage_schedule_proto_enumTypes[0]
}

func (x Schedule_IntervalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule_IntervalType.Descriptor instead.
func (Schedule_IntervalType) EnumDescriptor() ([]byte, []int) {
	return file_storage_schedule_proto_rawDescGZIP(), []int{0, 0}
}

type Schedule struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	IntervalType Schedule_IntervalType  `protobuf:"varint,1,opt,name=interval_type,json=intervalType,proto3,enum=storage.Schedule_IntervalType" json:"interval_type,omitempty"`
	Hour         int32                  `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute       int32                  `protobuf:"varint,3,opt,name=minute,proto3" json:"minute,omitempty"`
	// Types that are valid to be assigned to Interval:
	//
	//	*Schedule_Weekly
	//	*Schedule_DaysOfWeek_
	//	*Schedule_DaysOfMonth_
	Interval      isSchedule_Interval `protobuf_oneof:"Interval"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	mi := &file_storage_schedule_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_storage_schedule_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_storage_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *Schedule) GetIntervalType() Schedule_IntervalType {
	if x != nil {
		return x.IntervalType
	}
	return Schedule_UNSET
}

func (x *Schedule) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

func (x *Schedule) GetMinute() int32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *Schedule) GetInterval() isSchedule_Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *Schedule) GetWeekly() *Schedule_WeeklyInterval {
	if x != nil {
		if x, ok := x.Interval.(*Schedule_Weekly); ok {
			return x.Weekly
		}
	}
	return nil
}

func (x *Schedule) GetDaysOfWeek() *Schedule_DaysOfWeek {
	if x != nil {
		if x, ok := x.Interval.(*Schedule_DaysOfWeek_); ok {
			return x.DaysOfWeek
		}
	}
	return nil
}

func (x *Schedule) GetDaysOfMonth() *Schedule_DaysOfMonth {
	if x != nil {
		if x, ok := x.Interval.(*Schedule_DaysOfMonth_); ok {
			return x.DaysOfMonth
		}
	}
	return nil
}

type isSchedule_Interval interface {
	isSchedule_Interval()
}

type Schedule_Weekly struct {
	Weekly *Schedule_WeeklyInterval `protobuf:"bytes,4,opt,name=weekly,proto3,oneof"`
}

type Schedule_DaysOfWeek_ struct {
	DaysOfWeek *Schedule_DaysOfWeek `protobuf:"bytes,5,opt,name=days_of_week,json=daysOfWeek,proto3,oneof"`
}

type Schedule_DaysOfMonth_ struct {
	DaysOfMonth *Schedule_DaysOfMonth `protobuf:"bytes,6,opt,name=days_of_month,json=daysOfMonth,proto3,oneof"`
}

func (*Schedule_Weekly) isSchedule_Interval() {}

func (*Schedule_DaysOfWeek_) isSchedule_Interval() {}

func (*Schedule_DaysOfMonth_) isSchedule_Interval() {}

type Schedule_WeeklyInterval struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Day           int32                  `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule_WeeklyInterval) Reset() {
	*x = Schedule_WeeklyInterval{}
	mi := &file_storage_schedule_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule_WeeklyInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule_WeeklyInterval) ProtoMessage() {}

func (x *Schedule_WeeklyInterval) ProtoReflect() protoreflect.Message {
	mi := &file_storage_schedule_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule_WeeklyInterval.ProtoReflect.Descriptor instead.
func (*Schedule_WeeklyInterval) Descriptor() ([]byte, []int) {
	return file_storage_schedule_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Schedule_WeeklyInterval) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

// Sunday = 0, Monday = 1, .... Saturday =  6
type Schedule_DaysOfWeek struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Days          []int32                `protobuf:"varint,1,rep,packed,name=days,proto3" json:"days,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule_DaysOfWeek) Reset() {
	*x = Schedule_DaysOfWeek{}
	mi := &file_storage_schedule_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule_DaysOfWeek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule_DaysOfWeek) ProtoMessage() {}

func (x *Schedule_DaysOfWeek) ProtoReflect() protoreflect.Message {
	mi := &file_storage_schedule_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule_DaysOfWeek.ProtoReflect.Descriptor instead.
func (*Schedule_DaysOfWeek) Descriptor() ([]byte, []int) {
	return file_storage_schedule_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Schedule_DaysOfWeek) GetDays() []int32 {
	if x != nil {
		return x.Days
	}
	return nil
}

// 1 for 1st, 2 for 2nd .... 31 for 31st
type Schedule_DaysOfMonth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Days          []int32                `protobuf:"varint,1,rep,packed,name=days,proto3" json:"days,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule_DaysOfMonth) Reset() {
	*x = Schedule_DaysOfMonth{}
	mi := &file_storage_schedule_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule_DaysOfMonth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule_DaysOfMonth) ProtoMessage() {}

func (x *Schedule_DaysOfMonth) ProtoReflect() protoreflect.Message {
	mi := &file_storage_schedule_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule_DaysOfMonth.ProtoReflect.Descriptor instead.
func (*Schedule_DaysOfMonth) Descriptor() ([]byte, []int) {
	return file_storage_schedule_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Schedule_DaysOfMonth) GetDays() []int32 {
	if x != nil {
		return x.Days
	}
	return nil
}

var File_storage_schedule_proto protoreflect.FileDescriptor

var file_storage_schedule_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x22, 0xf2, 0x03, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x43,
	0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12,
	0x3a, 0x0a, 0x06, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x06, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0c, 0x64,
	0x61, 0x79, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x48,
	0x00, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x43, 0x0a,
	0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x1a, 0x22, 0x0a, 0x0e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x1a, 0x20, 0x0a, 0x0a, 0x44, 0x61, 0x79, 0x73, 0x4f, 0x66,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x1a, 0x21, 0x0a, 0x0b, 0x44, 0x61, 0x79, 0x73,
	0x4f, 0x66, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x22, 0x3d, 0x0a, 0x0c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x45, 0x4b, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x03, 0x42, 0x0a, 0x0a, 0x08, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_storage_schedule_proto_rawDescOnce sync.Once
	file_storage_schedule_proto_rawDescData []byte
)

func file_storage_schedule_proto_rawDescGZIP() []byte {
	file_storage_schedule_proto_rawDescOnce.Do(func() {
		file_storage_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_schedule_proto_rawDesc), len(file_storage_schedule_proto_rawDesc)))
	})
	return file_storage_schedule_proto_rawDescData
}

var file_storage_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_schedule_proto_goTypes = []any{
	(Schedule_IntervalType)(0),      // 0: storage.Schedule.IntervalType
	(*Schedule)(nil),                // 1: storage.Schedule
	(*Schedule_WeeklyInterval)(nil), // 2: storage.Schedule.WeeklyInterval
	(*Schedule_DaysOfWeek)(nil),     // 3: storage.Schedule.DaysOfWeek
	(*Schedule_DaysOfMonth)(nil),    // 4: storage.Schedule.DaysOfMonth
}
var file_storage_schedule_proto_depIdxs = []int32{
	0, // 0: storage.Schedule.interval_type:type_name -> storage.Schedule.IntervalType
	2, // 1: storage.Schedule.weekly:type_name -> storage.Schedule.WeeklyInterval
	3, // 2: storage.Schedule.days_of_week:type_name -> storage.Schedule.DaysOfWeek
	4, // 3: storage.Schedule.days_of_month:type_name -> storage.Schedule.DaysOfMonth
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_schedule_proto_init() }
func file_storage_schedule_proto_init() {
	if File_storage_schedule_proto != nil {
		return
	}
	file_storage_schedule_proto_msgTypes[0].OneofWrappers = []any{
		(*Schedule_Weekly)(nil),
		(*Schedule_DaysOfWeek_)(nil),
		(*Schedule_DaysOfMonth_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_schedule_proto_rawDesc), len(file_storage_schedule_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_schedule_proto_goTypes,
		DependencyIndexes: file_storage_schedule_proto_depIdxs,
		EnumInfos:         file_storage_schedule_proto_enumTypes,
		MessageInfos:      file_storage_schedule_proto_msgTypes,
	}.Build()
	File_storage_schedule_proto = out.File
	file_storage_schedule_proto_goTypes = nil
	file_storage_schedule_proto_depIdxs = nil
}
