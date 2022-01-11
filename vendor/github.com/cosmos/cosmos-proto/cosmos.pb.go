// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos_proto/cosmos.proto

package cosmos_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScalarType int32

const (
	ScalarType_SCALAR_TYPE_UNSPECIFIED ScalarType = 0
	ScalarType_SCALAR_TYPE_STRING      ScalarType = 1
	ScalarType_SCALAR_TYPE_BYTES       ScalarType = 2
)

// Enum value maps for ScalarType.
var (
	ScalarType_name = map[int32]string{
		0: "SCALAR_TYPE_UNSPECIFIED",
		1: "SCALAR_TYPE_STRING",
		2: "SCALAR_TYPE_BYTES",
	}
	ScalarType_value = map[string]int32{
		"SCALAR_TYPE_UNSPECIFIED": 0,
		"SCALAR_TYPE_STRING":      1,
		"SCALAR_TYPE_BYTES":       2,
	}
)

func (x ScalarType) Enum() *ScalarType {
	p := new(ScalarType)
	*p = x
	return p
}

func (x ScalarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScalarType) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_proto_cosmos_proto_enumTypes[0].Descriptor()
}

func (ScalarType) Type() protoreflect.EnumType {
	return &file_cosmos_proto_cosmos_proto_enumTypes[0]
}

func (x ScalarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScalarType.Descriptor instead.
func (ScalarType) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_proto_cosmos_proto_rawDescGZIP(), []int{0}
}

// InterfaceDescriptor describes an interface type to be used with
// accepts_interface and implements_interface and declared by declare_interface.
type InterfaceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the interface. It should be a short-name (without
	// a period) such that the fully qualified name of the interface will be
	// package.name, ex. for the package a.b and interface named C, the
	// fully-qualified name will be a.b.C.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// description is a human-readable description of the interface and its
	// purpose.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InterfaceDescriptor) Reset() {
	*x = InterfaceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_proto_cosmos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceDescriptor) ProtoMessage() {}

func (x *InterfaceDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_proto_cosmos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceDescriptor.ProtoReflect.Descriptor instead.
func (*InterfaceDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_proto_cosmos_proto_rawDescGZIP(), []int{0}
}

func (x *InterfaceDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InterfaceDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// ScalarDescriptor describes an scalar type to be used with
// the scalar field option and declared by declare_scalar.
// Scalars extend simple protobuf built-in types with additional
// syntax and semantics, for instance to represent big integers.
// Scalars should ideally define an encoding such that there is only one
// valid syntactical representation for a given semantic meaning,
// i.e. the encoding should be deterministic.
type ScalarDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the scalar. It should be a short-name (without
	// a period) such that the fully qualified name of the scalar will be
	// package.name, ex. for the package a.b and scalar named C, the
	// fully-qualified name will be a.b.C.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// description is a human-readable description of the scalar and its
	// encoding format. For instance a big integer or decimal scalar should
	// specify precisely the expected encoding format.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// field_type is the type of field with which this scalar can be used.
	// Scalars can be used with one and only one type of field so that
	// encoding standards and simple and clear. Currently only string and
	// bytes fields are supported for scalars.
	FieldType []ScalarType `protobuf:"varint,3,rep,packed,name=field_type,json=fieldType,proto3,enum=cosmos_proto.ScalarType" json:"field_type,omitempty"`
}

func (x *ScalarDescriptor) Reset() {
	*x = ScalarDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_proto_cosmos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalarDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalarDescriptor) ProtoMessage() {}

func (x *ScalarDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_proto_cosmos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalarDescriptor.ProtoReflect.Descriptor instead.
func (*ScalarDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_proto_cosmos_proto_rawDescGZIP(), []int{1}
}

func (x *ScalarDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScalarDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ScalarDescriptor) GetFieldType() []ScalarType {
	if x != nil {
		return x.FieldType
	}
	return nil
}

var file_cosmos_proto_cosmos_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         93001,
		Name:          "cosmos_proto.implements_interface",
		Tag:           "bytes,93001,rep,name=implements_interface",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93001,
		Name:          "cosmos_proto.accepts_interface",
		Tag:           "bytes,93001,opt,name=accepts_interface",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93002,
		Name:          "cosmos_proto.scalar",
		Tag:           "bytes,93002,opt,name=scalar",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: ([]*InterfaceDescriptor)(nil),
		Field:         793021,
		Name:          "cosmos_proto.declare_interface",
		Tag:           "bytes,793021,rep,name=declare_interface",
		Filename:      "cosmos_proto/cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: ([]*ScalarDescriptor)(nil),
		Field:         793022,
		Name:          "cosmos_proto.declare_scalar",
		Tag:           "bytes,793022,rep,name=declare_scalar",
		Filename:      "cosmos_proto/cosmos.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// implements_interface is used to indicate the type name of the interface
	// that a message implements so that it can be used in google.protobuf.Any
	// fields that accept that interface. A message can implement multiple
	// interfaces. Interfaces should be declared using a declare_interface
	// file option.
	//
	// repeated string implements_interface = 93001;
	E_ImplementsInterface = &file_cosmos_proto_cosmos_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// accepts_interface is used to annotate that a google.protobuf.Any
	// field accepts messages that implement the specified interface.
	// Interfaces should be declared using a declare_interface file option.
	//
	// optional string accepts_interface = 93001;
	E_AcceptsInterface = &file_cosmos_proto_cosmos_proto_extTypes[1]
	// scalar is used to indicate that this field follows the formatting defined
	// by the named scalar which should be declared with declare_scalar. Code
	// generators may choose to use this information to map this field to a
	// language-specific type representing the scalar.
	//
	// optional string scalar = 93002;
	E_Scalar = &file_cosmos_proto_cosmos_proto_extTypes[2]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// declare_interface declares an interface type to be used with
	// accepts_interface and implements_interface. Interface names are
	// expected to follow the following convention such that their declaration
	// can be discovered by tools: for a given interface type a.b.C, it is
	// expected that the declaration will be found in a protobuf file named
	// a/b/interfaces.proto in the file descriptor set.
	//
	// repeated cosmos_proto.InterfaceDescriptor declare_interface = 793021;
	E_DeclareInterface = &file_cosmos_proto_cosmos_proto_extTypes[3]
	// declare_scalar declares a scalar type to be used with
	// the scalar field option. Scalar names are
	// expected to follow the following convention such that their declaration
	// can be discovered by tools: for a given scalar type a.b.C, it is
	// expected that the declaration will be found in a protobuf file named
	// a/b/scalars.proto in the file descriptor set.
	//
	// repeated cosmos_proto.ScalarDescriptor declare_scalar = 793022;
	E_DeclareScalar = &file_cosmos_proto_cosmos_proto_extTypes[4]
)

var File_cosmos_proto_cosmos_proto protoreflect.FileDescriptor

var file_cosmos_proto_cosmos_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x13, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x58, 0x0a, 0x0a,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x43,
	0x41, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x41, 0x4c, 0x41,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x59, 0x54, 0x45, 0x53, 0x10, 0x02, 0x3a, 0x54, 0x0a, 0x14, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xc9, 0xd6, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x3a, 0x4c, 0x0a, 0x11,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xc9, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x3a, 0x37, 0x0a, 0x06, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xca, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x3a, 0x6e, 0x0a, 0x11, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xb3, 0x30, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x52, 0x10, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x3a, 0x65, 0x0a, 0x0e, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x5f, 0x73,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xb3, 0x30, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x64, 0x65, 0x63,
	0x6c, 0x61, 0x72, 0x65, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x42, 0x9e, 0x01, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42,
	0x0b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x0b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cosmos_proto_cosmos_proto_rawDescOnce sync.Once
	file_cosmos_proto_cosmos_proto_rawDescData = file_cosmos_proto_cosmos_proto_rawDesc
)

func file_cosmos_proto_cosmos_proto_rawDescGZIP() []byte {
	file_cosmos_proto_cosmos_proto_rawDescOnce.Do(func() {
		file_cosmos_proto_cosmos_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_proto_cosmos_proto_rawDescData)
	})
	return file_cosmos_proto_cosmos_proto_rawDescData
}

var file_cosmos_proto_cosmos_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cosmos_proto_cosmos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_proto_cosmos_proto_goTypes = []interface{}{
	(ScalarType)(0),                     // 0: cosmos_proto.ScalarType
	(*InterfaceDescriptor)(nil),         // 1: cosmos_proto.InterfaceDescriptor
	(*ScalarDescriptor)(nil),            // 2: cosmos_proto.ScalarDescriptor
	(*descriptorpb.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
	(*descriptorpb.FileOptions)(nil),    // 5: google.protobuf.FileOptions
}
var file_cosmos_proto_cosmos_proto_depIdxs = []int32{
	0, // 0: cosmos_proto.ScalarDescriptor.field_type:type_name -> cosmos_proto.ScalarType
	3, // 1: cosmos_proto.implements_interface:extendee -> google.protobuf.MessageOptions
	4, // 2: cosmos_proto.accepts_interface:extendee -> google.protobuf.FieldOptions
	4, // 3: cosmos_proto.scalar:extendee -> google.protobuf.FieldOptions
	5, // 4: cosmos_proto.declare_interface:extendee -> google.protobuf.FileOptions
	5, // 5: cosmos_proto.declare_scalar:extendee -> google.protobuf.FileOptions
	1, // 6: cosmos_proto.declare_interface:type_name -> cosmos_proto.InterfaceDescriptor
	2, // 7: cosmos_proto.declare_scalar:type_name -> cosmos_proto.ScalarDescriptor
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	1, // [1:6] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_proto_cosmos_proto_init() }
func file_cosmos_proto_cosmos_proto_init() {
	if File_cosmos_proto_cosmos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_proto_cosmos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceDescriptor); i {
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
		file_cosmos_proto_cosmos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalarDescriptor); i {
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
			RawDescriptor: file_cosmos_proto_cosmos_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_proto_cosmos_proto_goTypes,
		DependencyIndexes: file_cosmos_proto_cosmos_proto_depIdxs,
		EnumInfos:         file_cosmos_proto_cosmos_proto_enumTypes,
		MessageInfos:      file_cosmos_proto_cosmos_proto_msgTypes,
		ExtensionInfos:    file_cosmos_proto_cosmos_proto_extTypes,
	}.Build()
	File_cosmos_proto_cosmos_proto = out.File
	file_cosmos_proto_cosmos_proto_rawDesc = nil
	file_cosmos_proto_cosmos_proto_goTypes = nil
	file_cosmos_proto_cosmos_proto_depIdxs = nil
}
