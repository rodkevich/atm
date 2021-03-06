// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: persons.proto

package persons

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

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     int32  `protobuf:"varint,1,opt,name=text,proto3" json:"text,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_persons_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetText() int32 {
	if x != nil {
		return x.Text
	}
	return 0
}

func (x *TestRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_persons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_persons_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_persons_proto protoreflect.FileDescriptor

var file_persons_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22, 0x3d, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x32, 0x42, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x06,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x64, 0x6b, 0x65, 0x76, 0x69, 0x63, 0x68, 0x2f, 0x61, 0x74,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_persons_proto_rawDescOnce sync.Once
	file_persons_proto_rawDescData = file_persons_proto_rawDesc
)

func file_persons_proto_rawDescGZIP() []byte {
	file_persons_proto_rawDescOnce.Do(func() {
		file_persons_proto_rawDescData = protoimpl.X.CompressGZIP(file_persons_proto_rawDescData)
	})
	return file_persons_proto_rawDescData
}

var file_persons_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_persons_proto_goTypes = []interface{}{
	(*TestRequest)(nil),  // 0: backend.TestRequest
	(*TestResponse)(nil), // 1: backend.TestResponse
}
var file_persons_proto_depIdxs = []int32{
	0, // 0: backend.Persons.Person:input_type -> backend.TestRequest
	1, // 1: backend.Persons.Person:output_type -> backend.TestResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_persons_proto_init() }
func file_persons_proto_init() {
	if File_persons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_persons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_persons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
			RawDescriptor: file_persons_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_persons_proto_goTypes,
		DependencyIndexes: file_persons_proto_depIdxs,
		MessageInfos:      file_persons_proto_msgTypes,
	}.Build()
	File_persons_proto = out.File
	file_persons_proto_rawDesc = nil
	file_persons_proto_goTypes = nil
	file_persons_proto_depIdxs = nil
}
