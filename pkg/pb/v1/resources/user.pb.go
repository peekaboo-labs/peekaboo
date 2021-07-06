// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/v1/resources/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" csv:"username" yaml:"username"`
	Comment  string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment" csv:"comment" yaml:"comment"`
	Uid      uint64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid" csv:"uid" yaml:"uid"`
	Gid      uint64 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid" csv:"gid" yaml:"gid"`
	// Signed UID for Mac OS X.
	UidSigned int64 `protobuf:"varint,5,opt,name=uid_signed,json=uidSigned,proto3" json:"uid_signed" csv:"uid_signed" yaml:"uid_signed"`
	// Signed GID for Mac OS X.
	GidSigned int64  `protobuf:"varint,6,opt,name=gid_signed,json=gidSigned,proto3" json:"gid_signed" csv:"gid_signed" yaml:"gid_signed"`
	Directory string `protobuf:"bytes,7,opt,name=directory,proto3" json:"directory" csv:"directory" yaml:"directory"`
	Shell     string `protobuf:"bytes,8,opt,name=shell,proto3" json:"shell" csv:"shell" yaml:"shell"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_resources_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_resources_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_v1_resources_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *User) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *User) GetGid() uint64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *User) GetUidSigned() int64 {
	if x != nil {
		return x.UidSigned
	}
	return 0
}

func (x *User) GetGidSigned() int64 {
	if x != nil {
		return x.GidSigned
	}
	return 0
}

func (x *User) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *User) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

var File_pb_v1_resources_user_proto protoreflect.FileDescriptor

var file_pb_v1_resources_user_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x65,
	0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x64, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x69, 0x64, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x64, 0x5f, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x69, 0x64, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2d,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_v1_resources_user_proto_rawDescOnce sync.Once
	file_pb_v1_resources_user_proto_rawDescData = file_pb_v1_resources_user_proto_rawDesc
)

func file_pb_v1_resources_user_proto_rawDescGZIP() []byte {
	file_pb_v1_resources_user_proto_rawDescOnce.Do(func() {
		file_pb_v1_resources_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_v1_resources_user_proto_rawDescData)
	})
	return file_pb_v1_resources_user_proto_rawDescData
}

var file_pb_v1_resources_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_v1_resources_user_proto_goTypes = []interface{}{
	(*User)(nil), // 0: peekaboo.v1.resources.User
}
var file_pb_v1_resources_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_v1_resources_user_proto_init() }
func file_pb_v1_resources_user_proto_init() {
	if File_pb_v1_resources_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_v1_resources_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_pb_v1_resources_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_v1_resources_user_proto_goTypes,
		DependencyIndexes: file_pb_v1_resources_user_proto_depIdxs,
		MessageInfos:      file_pb_v1_resources_user_proto_msgTypes,
	}.Build()
	File_pb_v1_resources_user_proto = out.File
	file_pb_v1_resources_user_proto_rawDesc = nil
	file_pb_v1_resources_user_proto_goTypes = nil
	file_pb_v1_resources_user_proto_depIdxs = nil
}
