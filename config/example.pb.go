// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: example.proto

package config

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

type ExampleBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TestBool   bool    `protobuf:"varint,2,opt,name=TestBool,proto3" json:"TestBool,omitempty"`
	TestInt32  int32   `protobuf:"varint,3,opt,name=TestInt32,proto3" json:"TestInt32,omitempty"`
	TestInt64  int64   `protobuf:"varint,4,opt,name=TestInt64,proto3" json:"TestInt64,omitempty"`
	TestUint32 uint32  `protobuf:"varint,5,opt,name=TestUint32,proto3" json:"TestUint32,omitempty"`
	TestUint64 uint64  `protobuf:"varint,6,opt,name=TestUint64,proto3" json:"TestUint64,omitempty"`
	TestString string  `protobuf:"bytes,7,opt,name=TestString,proto3" json:"TestString,omitempty"`
	TestFloat  float32 `protobuf:"fixed32,8,opt,name=TestFloat,proto3" json:"TestFloat,omitempty"`
	TestDouble float64 `protobuf:"fixed64,9,opt,name=TestDouble,proto3" json:"TestDouble,omitempty"`
}

func (x *ExampleBase) Reset() {
	*x = ExampleBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleBase) ProtoMessage() {}

func (x *ExampleBase) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleBase.ProtoReflect.Descriptor instead.
func (*ExampleBase) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleBase) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ExampleBase) GetTestBool() bool {
	if x != nil {
		return x.TestBool
	}
	return false
}

func (x *ExampleBase) GetTestInt32() int32 {
	if x != nil {
		return x.TestInt32
	}
	return 0
}

func (x *ExampleBase) GetTestInt64() int64 {
	if x != nil {
		return x.TestInt64
	}
	return 0
}

func (x *ExampleBase) GetTestUint32() uint32 {
	if x != nil {
		return x.TestUint32
	}
	return 0
}

func (x *ExampleBase) GetTestUint64() uint64 {
	if x != nil {
		return x.TestUint64
	}
	return 0
}

func (x *ExampleBase) GetTestString() string {
	if x != nil {
		return x.TestString
	}
	return ""
}

func (x *ExampleBase) GetTestFloat() float32 {
	if x != nil {
		return x.TestFloat
	}
	return 0
}

func (x *ExampleBase) GetTestDouble() float64 {
	if x != nil {
		return x.TestDouble
	}
	return 0
}

type ExampleBaseCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[uint32]*ExampleBase `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExampleBaseCfg) Reset() {
	*x = ExampleBaseCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleBaseCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleBaseCfg) ProtoMessage() {}

func (x *ExampleBaseCfg) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleBaseCfg.ProtoReflect.Descriptor instead.
func (*ExampleBaseCfg) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleBaseCfg) GetDatas() map[uint32]*ExampleBase {
	if x != nil {
		return x.Datas
	}
	return nil
}

type ExampleComplex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TestInt32S  []int32           `protobuf:"varint,2,rep,packed,name=TestInt32s,proto3" json:"TestInt32s,omitempty"`
	TestInt64S  []int64           `protobuf:"varint,3,rep,packed,name=TestInt64s,proto3" json:"TestInt64s,omitempty"`
	TestUint32S []uint32          `protobuf:"varint,4,rep,packed,name=TestUint32s,proto3" json:"TestUint32s,omitempty"`
	TestUint64S []uint64          `protobuf:"varint,5,rep,packed,name=TestUint64s,proto3" json:"TestUint64s,omitempty"`
	TestStrings []string          `protobuf:"bytes,6,rep,name=TestStrings,proto3" json:"TestStrings,omitempty"`
	TestMap     map[uint32]uint32 `protobuf:"bytes,7,rep,name=TestMap,proto3" json:"TestMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ExampleComplex) Reset() {
	*x = ExampleComplex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleComplex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleComplex) ProtoMessage() {}

func (x *ExampleComplex) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleComplex.ProtoReflect.Descriptor instead.
func (*ExampleComplex) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *ExampleComplex) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ExampleComplex) GetTestInt32S() []int32 {
	if x != nil {
		return x.TestInt32S
	}
	return nil
}

func (x *ExampleComplex) GetTestInt64S() []int64 {
	if x != nil {
		return x.TestInt64S
	}
	return nil
}

func (x *ExampleComplex) GetTestUint32S() []uint32 {
	if x != nil {
		return x.TestUint32S
	}
	return nil
}

func (x *ExampleComplex) GetTestUint64S() []uint64 {
	if x != nil {
		return x.TestUint64S
	}
	return nil
}

func (x *ExampleComplex) GetTestStrings() []string {
	if x != nil {
		return x.TestStrings
	}
	return nil
}

func (x *ExampleComplex) GetTestMap() map[uint32]uint32 {
	if x != nil {
		return x.TestMap
	}
	return nil
}

type ExampleComplexCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[uint32]*ExampleComplex `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExampleComplexCfg) Reset() {
	*x = ExampleComplexCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleComplexCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleComplexCfg) ProtoMessage() {}

func (x *ExampleComplexCfg) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleComplexCfg.ProtoReflect.Descriptor instead.
func (*ExampleComplexCfg) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{3}
}

func (x *ExampleComplexCfg) GetDatas() map[uint32]*ExampleComplex {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x93, 0x02, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x42,
	0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x54, 0x65, 0x73, 0x74, 0x42,
	0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x54, 0x65, 0x73, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22, 0x98, 0x01,
	0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x43, 0x66, 0x67,
	0x12, 0x37, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x43, 0x66, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x4d, 0x0a, 0x0a, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc1, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x3d, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70,
	0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x01, 0x0a,
	0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x43,
	0x66, 0x67, 0x12, 0x3a, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x66, 0x67, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x50,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_example_proto_goTypes = []interface{}{
	(*ExampleBase)(nil),       // 0: config.ExampleBase
	(*ExampleBaseCfg)(nil),    // 1: config.ExampleBaseCfg
	(*ExampleComplex)(nil),    // 2: config.ExampleComplex
	(*ExampleComplexCfg)(nil), // 3: config.ExampleComplexCfg
	nil,                       // 4: config.ExampleBaseCfg.DatasEntry
	nil,                       // 5: config.ExampleComplex.TestMapEntry
	nil,                       // 6: config.ExampleComplexCfg.DatasEntry
}
var file_example_proto_depIdxs = []int32{
	4, // 0: config.ExampleBaseCfg.datas:type_name -> config.ExampleBaseCfg.DatasEntry
	5, // 1: config.ExampleComplex.TestMap:type_name -> config.ExampleComplex.TestMapEntry
	6, // 2: config.ExampleComplexCfg.datas:type_name -> config.ExampleComplexCfg.DatasEntry
	0, // 3: config.ExampleBaseCfg.DatasEntry.value:type_name -> config.ExampleBase
	2, // 4: config.ExampleComplexCfg.DatasEntry.value:type_name -> config.ExampleComplex
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleBase); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleBaseCfg); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleComplex); i {
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
		file_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleComplexCfg); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}