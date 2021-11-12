// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: common.proto

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

type MatchMode int32

const (
	MatchMode_normal_match MatchMode = 0
)

// Enum value maps for MatchMode.
var (
	MatchMode_name = map[int32]string{
		0: "normal_match",
	}
	MatchMode_value = map[string]int32{
		"normal_match": 0,
	}
)

func (x MatchMode) Enum() *MatchMode {
	p := new(MatchMode)
	*p = x
	return p
}

func (x MatchMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchMode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (MatchMode) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x MatchMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchMode.Descriptor instead.
func (MatchMode) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type RoleBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid    uint64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
	Rolename  string `protobuf:"bytes,2,opt,name=rolename,proto3" json:"rolename,omitempty"`
	Rolelevel uint32 `protobuf:"varint,3,opt,name=rolelevel,proto3" json:"rolelevel,omitempty"`
}

func (x *RoleBaseInfo) Reset() {
	*x = RoleBaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBaseInfo) ProtoMessage() {}

func (x *RoleBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBaseInfo.ProtoReflect.Descriptor instead.
func (*RoleBaseInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *RoleBaseInfo) GetRoleid() uint64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

func (x *RoleBaseInfo) GetRolename() string {
	if x != nil {
		return x.Rolename
	}
	return ""
}

func (x *RoleBaseInfo) GetRolelevel() uint32 {
	if x != nil {
		return x.Rolelevel
	}
	return 0
}

type MasterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MasterInfo) Reset() {
	*x = MasterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterInfo) ProtoMessage() {}

func (x *MasterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterInfo.ProtoReflect.Descriptor instead.
func (*MasterInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *MasterInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MasterModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Masters        []*MasterInfo `protobuf:"bytes,1,rep,name=masters,proto3" json:"masters,omitempty"`
	Battlemasterid uint32        `protobuf:"varint,2,opt,name=battlemasterid,proto3" json:"battlemasterid,omitempty"`
}

func (x *MasterModule) Reset() {
	*x = MasterModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterModule) ProtoMessage() {}

func (x *MasterModule) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterModule.ProtoReflect.Descriptor instead.
func (*MasterModule) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *MasterModule) GetMasters() []*MasterInfo {
	if x != nil {
		return x.Masters
	}
	return nil
}

func (x *MasterModule) GetBattlemasterid() uint32 {
	if x != nil {
		return x.Battlemasterid
	}
	return 0
}

type BattleLoadProcess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid  uint64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
	Percent uint32 `protobuf:"varint,2,opt,name=percent,proto3" json:"percent,omitempty"` //百分比
}

func (x *BattleLoadProcess) Reset() {
	*x = BattleLoadProcess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleLoadProcess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleLoadProcess) ProtoMessage() {}

func (x *BattleLoadProcess) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleLoadProcess.ProtoReflect.Descriptor instead.
func (*BattleLoadProcess) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *BattleLoadProcess) GetRoleid() uint64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

func (x *BattleLoadProcess) GetPercent() uint32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

type TeamMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId   uint64 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *TeamMember) Reset() {
	*x = TeamMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamMember) ProtoMessage() {}

func (x *TeamMember) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamMember.ProtoReflect.Descriptor instead.
func (*TeamMember) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *TeamMember) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *TeamMember) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type TeamInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode          MatchMode     `protobuf:"varint,1,opt,name=mode,proto3,enum=pb.MatchMode" json:"mode,omitempty"`
	Teamid        uint64        `protobuf:"varint,2,opt,name=teamid,proto3" json:"teamid,omitempty"`
	CaptainRoleId uint64        `protobuf:"varint,3,opt,name=captain_role_id,json=captainRoleId,proto3" json:"captain_role_id,omitempty"`
	Members       []*TeamMember `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *TeamInfo) Reset() {
	*x = TeamInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamInfo) ProtoMessage() {}

func (x *TeamInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamInfo.ProtoReflect.Descriptor instead.
func (*TeamInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *TeamInfo) GetMode() MatchMode {
	if x != nil {
		return x.Mode
	}
	return MatchMode_normal_match
}

func (x *TeamInfo) GetTeamid() uint64 {
	if x != nil {
		return x.Teamid
	}
	return 0
}

func (x *TeamInfo) GetCaptainRoleId() uint64 {
	if x != nil {
		return x.CaptainRoleId
	}
	return 0
}

func (x *TeamInfo) GetMembers() []*TeamMember {
	if x != nil {
		return x.Members
	}
	return nil
}

type RoomMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId   uint64 `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	MasterId uint32 `protobuf:"varint,3,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
}

func (x *RoomMember) Reset() {
	*x = RoomMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMember) ProtoMessage() {}

func (x *RoomMember) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMember.ProtoReflect.Descriptor instead.
func (*RoomMember) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *RoomMember) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *RoomMember) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *RoomMember) GetMasterId() uint32 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

type ChessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Tid uint32 `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *ChessInfo) Reset() {
	*x = ChessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessInfo) ProtoMessage() {}

func (x *ChessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessInfo.ProtoReflect.Descriptor instead.
func (*ChessInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *ChessInfo) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ChessInfo) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x62, 0x0a, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x6f, 0x6c,
	0x65, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x1d, 0x0a, 0x0b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x62, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61,
	0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x22, 0x60, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x2a, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_proto_goTypes = []interface{}{
	(MatchMode)(0),            // 0: pb.match_mode
	(*RoleBaseInfo)(nil),      // 1: pb.role_base_info
	(*MasterInfo)(nil),        // 2: pb.master_info
	(*MasterModule)(nil),      // 3: pb.master_module
	(*BattleLoadProcess)(nil), // 4: pb.battle_load_process
	(*TeamMember)(nil),        // 5: pb.team_member
	(*TeamInfo)(nil),          // 6: pb.team_info
	(*RoomMember)(nil),        // 7: pb.room_member
	(*ChessInfo)(nil),         // 8: pb.chess_info
}
var file_common_proto_depIdxs = []int32{
	2, // 0: pb.master_module.masters:type_name -> pb.master_info
	0, // 1: pb.team_info.mode:type_name -> pb.match_mode
	5, // 2: pb.team_info.members:type_name -> pb.team_member
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBaseInfo); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterInfo); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterModule); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleLoadProcess); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamMember); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamInfo); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMember); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessInfo); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
