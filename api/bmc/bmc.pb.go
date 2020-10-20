// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: bmc.proto

package bmc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BMCAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionString string `protobuf:"bytes,1,opt,name=connectionString,proto3" json:"connectionString,omitempty"`
}

func (x *BMCAccess) Reset() {
	*x = BMCAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bmc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BMCAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BMCAccess) ProtoMessage() {}

func (x *BMCAccess) ProtoReflect() protoreflect.Message {
	mi := &file_bmc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BMCAccess.ProtoReflect.Descriptor instead.
func (*BMCAccess) Descriptor() ([]byte, []int) {
	return file_bmc_proto_rawDescGZIP(), []int{0}
}

func (x *BMCAccess) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

type ServerPowerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ServerPowerStatus) Reset() {
	*x = ServerPowerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bmc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerPowerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerPowerStatus) ProtoMessage() {}

func (x *ServerPowerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_bmc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerPowerStatus.ProtoReflect.Descriptor instead.
func (*ServerPowerStatus) Descriptor() ([]byte, []int) {
	return file_bmc_proto_rawDescGZIP(), []int{1}
}

func (x *ServerPowerStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_bmc_proto protoreflect.FileDescriptor

var file_bmc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x6d, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x6d, 0x63,
	0x22, 0x37, 0x0a, 0x09, 0x42, 0x4d, 0x43, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x47, 0x0a, 0x03, 0x42, 0x4d, 0x43, 0x12, 0x40, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e, 0x62, 0x6d, 0x63, 0x2e, 0x42, 0x4d, 0x43, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x16, 0x2e, 0x62, 0x6d, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x6f, 0x70, 0x61, 0x33, 0x32, 0x33, 0x2f, 0x6b, 0x69, 0x6d, 0x62, 0x61, 0x70, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x6d, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bmc_proto_rawDescOnce sync.Once
	file_bmc_proto_rawDescData = file_bmc_proto_rawDesc
)

func file_bmc_proto_rawDescGZIP() []byte {
	file_bmc_proto_rawDescOnce.Do(func() {
		file_bmc_proto_rawDescData = protoimpl.X.CompressGZIP(file_bmc_proto_rawDescData)
	})
	return file_bmc_proto_rawDescData
}

var file_bmc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bmc_proto_goTypes = []interface{}{
	(*BMCAccess)(nil),         // 0: bmc.BMCAccess
	(*ServerPowerStatus)(nil), // 1: bmc.ServerPowerStatus
}
var file_bmc_proto_depIdxs = []int32{
	0, // 0: bmc.BMC.GetServerPowerStatus:input_type -> bmc.BMCAccess
	1, // 1: bmc.BMC.GetServerPowerStatus:output_type -> bmc.ServerPowerStatus
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bmc_proto_init() }
func file_bmc_proto_init() {
	if File_bmc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bmc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BMCAccess); i {
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
		file_bmc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerPowerStatus); i {
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
			RawDescriptor: file_bmc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bmc_proto_goTypes,
		DependencyIndexes: file_bmc_proto_depIdxs,
		MessageInfos:      file_bmc_proto_msgTypes,
	}.Build()
	File_bmc_proto = out.File
	file_bmc_proto_rawDesc = nil
	file_bmc_proto_goTypes = nil
	file_bmc_proto_depIdxs = nil
}
