// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/v1/resources/group.proto

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

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupname string `protobuf:"bytes,1,opt,name=groupname,proto3" json:"groupname" csv:"groupname" yaml:"groupname"`
	Gid       uint64 `protobuf:"varint,2,opt,name=gid,proto3" json:"gid" csv:"gid" yaml:"gid"`
	// Signed GID for Mac OS X.
	GidSigned int64    `protobuf:"varint,3,opt,name=gid_signed,json=gidSigned,proto3" json:"gid_signed" csv:"gid_signed" yaml:"gid_signed"`
	Members   []string `protobuf:"bytes,4,rep,name=members,proto3" json:"members" csv:"members" yaml:"members"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_resources_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_resources_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_pb_v1_resources_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetGroupname() string {
	if x != nil {
		return x.Groupname
	}
	return ""
}

func (x *Group) GetGid() uint64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *Group) GetGidSigned() int64 {
	if x != nil {
		return x.GidSigned
	}
	return 0
}

func (x *Group) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_pb_v1_resources_group_proto protoreflect.FileDescriptor

var file_pb_v1_resources_group_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70,
	0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x0a,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x69, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x67, 0x69, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2d, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_v1_resources_group_proto_rawDescOnce sync.Once
	file_pb_v1_resources_group_proto_rawDescData = file_pb_v1_resources_group_proto_rawDesc
)

func file_pb_v1_resources_group_proto_rawDescGZIP() []byte {
	file_pb_v1_resources_group_proto_rawDescOnce.Do(func() {
		file_pb_v1_resources_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_v1_resources_group_proto_rawDescData)
	})
	return file_pb_v1_resources_group_proto_rawDescData
}

var file_pb_v1_resources_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_v1_resources_group_proto_goTypes = []interface{}{
	(*Group)(nil), // 0: peekaboo.v1.resources.Group
}
var file_pb_v1_resources_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_v1_resources_group_proto_init() }
func file_pb_v1_resources_group_proto_init() {
	if File_pb_v1_resources_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_v1_resources_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_pb_v1_resources_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_v1_resources_group_proto_goTypes,
		DependencyIndexes: file_pb_v1_resources_group_proto_depIdxs,
		MessageInfos:      file_pb_v1_resources_group_proto_msgTypes,
	}.Build()
	File_pb_v1_resources_group_proto = out.File
	file_pb_v1_resources_group_proto_rawDesc = nil
	file_pb_v1_resources_group_proto_goTypes = nil
	file_pb_v1_resources_group_proto_depIdxs = nil
}
