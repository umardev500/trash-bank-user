// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pb/user/user-find-one.proto

package user

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

type UserFindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User    string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pass    string `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
	IsLogin bool   `protobuf:"varint,4,opt,name=is_login,json=isLogin,proto3" json:"is_login,omitempty"`
}

func (x *UserFindOneRequest) Reset() {
	*x = UserFindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_find_one_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindOneRequest) ProtoMessage() {}

func (x *UserFindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_find_one_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindOneRequest.ProtoReflect.Descriptor instead.
func (*UserFindOneRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_user_find_one_proto_rawDescGZIP(), []int{0}
}

func (x *UserFindOneRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserFindOneRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UserFindOneRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *UserFindOneRequest) GetIsLogin() bool {
	if x != nil {
		return x.IsLogin
	}
	return false
}

type UserFindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEmpty bool  `protobuf:"varint,1,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
	Payload *User `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UserFindOneResponse) Reset() {
	*x = UserFindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_find_one_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindOneResponse) ProtoMessage() {}

func (x *UserFindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_find_one_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindOneResponse.ProtoReflect.Descriptor instead.
func (*UserFindOneResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_user_find_one_proto_rawDescGZIP(), []int{1}
}

func (x *UserFindOneResponse) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

func (x *UserFindOneResponse) GetPayload() *User {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_pb_user_user_find_one_proto protoreflect.FileDescriptor

var file_pb_user_user_find_one_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x66,
	0x69, 0x6e, 0x64, 0x2d, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70,
	0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x70, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x22, 0x51, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_user_find_one_proto_rawDescOnce sync.Once
	file_pb_user_user_find_one_proto_rawDescData = file_pb_user_user_find_one_proto_rawDesc
)

func file_pb_user_user_find_one_proto_rawDescGZIP() []byte {
	file_pb_user_user_find_one_proto_rawDescOnce.Do(func() {
		file_pb_user_user_find_one_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_user_find_one_proto_rawDescData)
	})
	return file_pb_user_user_find_one_proto_rawDescData
}

var file_pb_user_user_find_one_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_user_user_find_one_proto_goTypes = []interface{}{
	(*UserFindOneRequest)(nil),  // 0: UserFindOneRequest
	(*UserFindOneResponse)(nil), // 1: UserFindOneResponse
	(*User)(nil),                // 2: User
}
var file_pb_user_user_find_one_proto_depIdxs = []int32{
	2, // 0: UserFindOneResponse.payload:type_name -> User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_user_user_find_one_proto_init() }
func file_pb_user_user_find_one_proto_init() {
	if File_pb_user_user_find_one_proto != nil {
		return
	}
	file_pb_user_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_user_user_find_one_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindOneRequest); i {
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
		file_pb_user_user_find_one_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindOneResponse); i {
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
			RawDescriptor: file_pb_user_user_find_one_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_user_user_find_one_proto_goTypes,
		DependencyIndexes: file_pb_user_user_find_one_proto_depIdxs,
		MessageInfos:      file_pb_user_user_find_one_proto_msgTypes,
	}.Build()
	File_pb_user_user_find_one_proto = out.File
	file_pb_user_user_find_one_proto_rawDesc = nil
	file_pb_user_user_find_one_proto_goTypes = nil
	file_pb_user_user_find_one_proto_depIdxs = nil
}
