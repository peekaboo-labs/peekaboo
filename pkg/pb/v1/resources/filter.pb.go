// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/v1/resources/filter.proto

package resources

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

type Operator int32

const (
	Operator_EQ   Operator = 0
	Operator_NEQ  Operator = 1
	Operator_GT   Operator = 2
	Operator_GTE  Operator = 3
	Operator_LT   Operator = 4
	Operator_LTE  Operator = 5
	Operator_IN   Operator = 6
	Operator_LIKE Operator = 7
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "EQ",
		1: "NEQ",
		2: "GT",
		3: "GTE",
		4: "LT",
		5: "LTE",
		6: "IN",
		7: "LIKE",
	}
	Operator_value = map[string]int32{
		"EQ":   0,
		"NEQ":  1,
		"GT":   2,
		"GTE":  3,
		"LT":   4,
		"LTE":  5,
		"IN":   6,
		"LIKE": 7,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_v1_resources_filter_proto_enumTypes[0].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_pb_v1_resources_filter_proto_enumTypes[0]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_pb_v1_resources_filter_proto_rawDescGZIP(), []int{0}
}

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field" csv:"field" yaml:"field"`
	Value    string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value" csv:"value" yaml:"value"`
	Values   []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values" csv:"values" yaml:"values"`
	Operator Operator `protobuf:"varint,4,opt,name=operator,proto3,enum=peekaboo.v1.resources.Operator" json:"operator" csv:"operator" yaml:"operator"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_resources_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_resources_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_pb_v1_resources_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Match) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Match) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Match) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Match) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_EQ
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field" csv:"field" yaml:"field"`
	Desc  bool   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc" csv:"desc" yaml:"desc"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_resources_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_resources_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_pb_v1_resources_filter_proto_rawDescGZIP(), []int{1}
}

func (x *Sort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields  []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields" csv:"fields" yaml:"fields"`
	Matches []*Match `protobuf:"bytes,2,rep,name=matches,proto3" json:"matches" csv:"matches" yaml:"matches"`
	Sorting []*Sort  `protobuf:"bytes,3,rep,name=sorting,proto3" json:"sorting" csv:"sorting" yaml:"sorting"`
	Limit   int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit" csv:"limit" yaml:"limit"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_resources_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_resources_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_pb_v1_resources_filter_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Filter) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *Filter) GetSorting() []*Sort {
	if x != nil {
		return x.Sorting
	}
	return nil
}

func (x *Filter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_pb_v1_resources_filter_proto protoreflect.FileDescriptor

var file_pb_v1_resources_filter_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x30, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x22, 0xa5, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x35, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2a, 0x4f, 0x0a, 0x08, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x51, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x45, 0x51, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x54, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x47, 0x54, 0x45, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x54, 0x10, 0x04,
	0x12, 0x07, 0x0a, 0x03, 0x4c, 0x54, 0x45, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10,
	0x06, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x07, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x65, 0x2d, 0x70, 0x65,
	0x72, 0x73, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_v1_resources_filter_proto_rawDescOnce sync.Once
	file_pb_v1_resources_filter_proto_rawDescData = file_pb_v1_resources_filter_proto_rawDesc
)

func file_pb_v1_resources_filter_proto_rawDescGZIP() []byte {
	file_pb_v1_resources_filter_proto_rawDescOnce.Do(func() {
		file_pb_v1_resources_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_v1_resources_filter_proto_rawDescData)
	})
	return file_pb_v1_resources_filter_proto_rawDescData
}

var file_pb_v1_resources_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_v1_resources_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_v1_resources_filter_proto_goTypes = []interface{}{
	(Operator)(0),  // 0: peekaboo.v1.resources.Operator
	(*Match)(nil),  // 1: peekaboo.v1.resources.Match
	(*Sort)(nil),   // 2: peekaboo.v1.resources.Sort
	(*Filter)(nil), // 3: peekaboo.v1.resources.Filter
}
var file_pb_v1_resources_filter_proto_depIdxs = []int32{
	0, // 0: peekaboo.v1.resources.Match.operator:type_name -> peekaboo.v1.resources.Operator
	1, // 1: peekaboo.v1.resources.Filter.matches:type_name -> peekaboo.v1.resources.Match
	2, // 2: peekaboo.v1.resources.Filter.sorting:type_name -> peekaboo.v1.resources.Sort
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_v1_resources_filter_proto_init() }
func file_pb_v1_resources_filter_proto_init() {
	if File_pb_v1_resources_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_v1_resources_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_v1_resources_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_v1_resources_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_v1_resources_filter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_v1_resources_filter_proto_goTypes,
		DependencyIndexes: file_pb_v1_resources_filter_proto_depIdxs,
		EnumInfos:         file_pb_v1_resources_filter_proto_enumTypes,
		MessageInfos:      file_pb_v1_resources_filter_proto_msgTypes,
	}.Build()
	File_pb_v1_resources_filter_proto = out.File
	file_pb_v1_resources_filter_proto_rawDesc = nil
	file_pb_v1_resources_filter_proto_goTypes = nil
	file_pb_v1_resources_filter_proto_depIdxs = nil
}
