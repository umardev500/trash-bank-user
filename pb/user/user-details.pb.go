// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pb/user/user-details.proto

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

type UserPhone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	IsWa   bool   `protobuf:"varint,2,opt,name=is_wa,json=isWa,proto3" json:"is_wa,omitempty"`
}

func (x *UserPhone) Reset() {
	*x = UserPhone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPhone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPhone) ProtoMessage() {}

func (x *UserPhone) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPhone.ProtoReflect.Descriptor instead.
func (*UserPhone) Descriptor() ([]byte, []int) {
	return file_pb_user_user_details_proto_rawDescGZIP(), []int{0}
}

func (x *UserPhone) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *UserPhone) GetIsWa() bool {
	if x != nil {
		return x.IsWa
	}
	return false
}

type UserAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Province   string `protobuf:"bytes,1,opt,name=province,proto3" json:"province,omitempty"`
	City       string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	District   string `protobuf:"bytes,3,opt,name=district,proto3" json:"district,omitempty"`
	Village    string `protobuf:"bytes,4,opt,name=village,proto3" json:"village,omitempty"`
	PostalCode string `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Detail     string `protobuf:"bytes,6,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *UserAddress) Reset() {
	*x = UserAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAddress) ProtoMessage() {}

func (x *UserAddress) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAddress.ProtoReflect.Descriptor instead.
func (*UserAddress) Descriptor() ([]byte, []int) {
	return file_pb_user_user_details_proto_rawDescGZIP(), []int{1}
}

func (x *UserAddress) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UserAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserAddress) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *UserAddress) GetVillage() string {
	if x != nil {
		return x.Village
	}
	return ""
}

func (x *UserAddress) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *UserAddress) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type UserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email   string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone   *UserPhone   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address *UserAddress `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UserDetails) Reset() {
	*x = UserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetails) ProtoMessage() {}

func (x *UserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetails.ProtoReflect.Descriptor instead.
func (*UserDetails) Descriptor() ([]byte, []int) {
	return file_pb_user_user_details_proto_rawDescGZIP(), []int{2}
}

func (x *UserDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDetails) GetPhone() *UserPhone {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *UserDetails) GetAddress() *UserAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_pb_user_user_details_proto protoreflect.FileDescriptor

var file_pb_user_user_details_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x77, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x69, 0x73, 0x57, 0x61, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x62, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_user_details_proto_rawDescOnce sync.Once
	file_pb_user_user_details_proto_rawDescData = file_pb_user_user_details_proto_rawDesc
)

func file_pb_user_user_details_proto_rawDescGZIP() []byte {
	file_pb_user_user_details_proto_rawDescOnce.Do(func() {
		file_pb_user_user_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_user_details_proto_rawDescData)
	})
	return file_pb_user_user_details_proto_rawDescData
}

var file_pb_user_user_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_user_user_details_proto_goTypes = []interface{}{
	(*UserPhone)(nil),   // 0: UserPhone
	(*UserAddress)(nil), // 1: UserAddress
	(*UserDetails)(nil), // 2: UserDetails
}
var file_pb_user_user_details_proto_depIdxs = []int32{
	0, // 0: UserDetails.phone:type_name -> UserPhone
	1, // 1: UserDetails.address:type_name -> UserAddress
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_user_user_details_proto_init() }
func file_pb_user_user_details_proto_init() {
	if File_pb_user_user_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_user_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPhone); i {
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
		file_pb_user_user_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAddress); i {
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
		file_pb_user_user_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetails); i {
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
			RawDescriptor: file_pb_user_user_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_user_user_details_proto_goTypes,
		DependencyIndexes: file_pb_user_user_details_proto_depIdxs,
		MessageInfos:      file_pb_user_user_details_proto_msgTypes,
	}.Build()
	File_pb_user_user_details_proto = out.File
	file_pb_user_user_details_proto_rawDesc = nil
	file_pb_user_user_details_proto_goTypes = nil
	file_pb_user_user_details_proto_depIdxs = nil
}