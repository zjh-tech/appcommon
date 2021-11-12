// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: master.proto

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

type MasterBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Cvlzn uint32 `protobuf:"varint,2,opt,name=Cvlzn,proto3" json:"Cvlzn,omitempty"`
	Skill uint32 `protobuf:"varint,3,opt,name=Skill,proto3" json:"Skill,omitempty"`
}

func (x *MasterBase) Reset() {
	*x = MasterBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterBase) ProtoMessage() {}

func (x *MasterBase) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterBase.ProtoReflect.Descriptor instead.
func (*MasterBase) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

func (x *MasterBase) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MasterBase) GetCvlzn() uint32 {
	if x != nil {
		return x.Cvlzn
	}
	return 0
}

func (x *MasterBase) GetSkill() uint32 {
	if x != nil {
		return x.Skill
	}
	return 0
}

type MasterBaseCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[uint32]*MasterBase `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MasterBaseCfg) Reset() {
	*x = MasterBaseCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterBaseCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterBaseCfg) ProtoMessage() {}

func (x *MasterBaseCfg) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterBaseCfg.ProtoReflect.Descriptor instead.
func (*MasterBaseCfg) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{1}
}

func (x *MasterBaseCfg) GetDatas() map[uint32]*MasterBase {
	if x != nil {
		return x.Datas
	}
	return nil
}

type MasterLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Score uint32 `protobuf:"varint,2,opt,name=Score,proto3" json:"Score,omitempty"`
}

func (x *MasterLevel) Reset() {
	*x = MasterLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterLevel) ProtoMessage() {}

func (x *MasterLevel) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterLevel.ProtoReflect.Descriptor instead.
func (*MasterLevel) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{2}
}

func (x *MasterLevel) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MasterLevel) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type MasterLevelCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[uint32]*MasterLevel `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MasterLevelCfg) Reset() {
	*x = MasterLevelCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterLevelCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterLevelCfg) ProtoMessage() {}

func (x *MasterLevelCfg) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterLevelCfg.ProtoReflect.Descriptor instead.
func (*MasterLevelCfg) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{3}
}

func (x *MasterLevelCfg) GetDatas() map[uint32]*MasterLevel {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_master_proto protoreflect.FileDescriptor

var file_master_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x48, 0x0a, 0x0a, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x76, 0x6c, 0x7a, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x43, 0x76, 0x6c, 0x7a, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x22, 0x95, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x43,
	0x66, 0x67, 0x12, 0x36, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x42, 0x61, 0x73, 0x65, 0x43, 0x66, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x4c, 0x0a, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x33, 0x0a, 0x0b, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x98, 0x01,
	0x0a, 0x0e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x66, 0x67,
	0x12, 0x37, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x43, 0x66, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x4d, 0x0a, 0x0a, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_master_proto_rawDescOnce sync.Once
	file_master_proto_rawDescData = file_master_proto_rawDesc
)

func file_master_proto_rawDescGZIP() []byte {
	file_master_proto_rawDescOnce.Do(func() {
		file_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_proto_rawDescData)
	})
	return file_master_proto_rawDescData
}

var file_master_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_master_proto_goTypes = []interface{}{
	(*MasterBase)(nil),     // 0: config.MasterBase
	(*MasterBaseCfg)(nil),  // 1: config.MasterBaseCfg
	(*MasterLevel)(nil),    // 2: config.MasterLevel
	(*MasterLevelCfg)(nil), // 3: config.MasterLevelCfg
	nil,                    // 4: config.MasterBaseCfg.DatasEntry
	nil,                    // 5: config.MasterLevelCfg.DatasEntry
}
var file_master_proto_depIdxs = []int32{
	4, // 0: config.MasterBaseCfg.datas:type_name -> config.MasterBaseCfg.DatasEntry
	5, // 1: config.MasterLevelCfg.datas:type_name -> config.MasterLevelCfg.DatasEntry
	0, // 2: config.MasterBaseCfg.DatasEntry.value:type_name -> config.MasterBase
	2, // 3: config.MasterLevelCfg.DatasEntry.value:type_name -> config.MasterLevel
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_master_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterBase); i {
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
		file_master_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterBaseCfg); i {
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
		file_master_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterLevel); i {
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
		file_master_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterLevelCfg); i {
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
			RawDescriptor: file_master_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
		MessageInfos:      file_master_proto_msgTypes,
	}.Build()
	File_master_proto = out.File
	file_master_proto_rawDesc = nil
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}