// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: chesses.proto

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

type ChessesStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                  uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NextID              uint32 `protobuf:"varint,2,opt,name=NextID,proto3" json:"NextID,omitempty"`
	Star                uint32 `protobuf:"varint,3,opt,name=Star,proto3" json:"Star,omitempty"`
	Cvlzn               uint32 `protobuf:"varint,4,opt,name=Cvlzn,proto3" json:"Cvlzn,omitempty"`
	Searchable          uint32 `protobuf:"varint,5,opt,name=Searchable,proto3" json:"Searchable,omitempty"`
	Quality             uint32 `protobuf:"varint,6,opt,name=Quality,proto3" json:"Quality,omitempty"`
	Price               uint32 `protobuf:"varint,7,opt,name=Price,proto3" json:"Price,omitempty"`
	Model               uint32 `protobuf:"varint,8,opt,name=Model,proto3" json:"Model,omitempty"`
	Attack              uint32 `protobuf:"varint,9,opt,name=Attack,proto3" json:"Attack,omitempty"`
	StatusHealth        uint32 `protobuf:"varint,10,opt,name=StatusHealth,proto3" json:"StatusHealth,omitempty"`
	ArmorPhysical       uint32 `protobuf:"varint,11,opt,name=ArmorPhysical,proto3" json:"ArmorPhysical,omitempty"`
	MagicalResistance   uint32 `protobuf:"varint,12,opt,name=MagicalResistance,proto3" json:"MagicalResistance,omitempty"`
	StatusHealthRegen   uint32 `protobuf:"varint,13,opt,name=StatusHealthRegen,proto3" json:"StatusHealthRegen,omitempty"`
	StatusMana          uint32 `protobuf:"varint,14,opt,name=StatusMana,proto3" json:"StatusMana,omitempty"`
	StatusManaRegen     uint32 `protobuf:"varint,15,opt,name=StatusManaRegen,proto3" json:"StatusManaRegen,omitempty"`
	Speed               uint32 `protobuf:"varint,16,opt,name=Speed,proto3" json:"Speed,omitempty"`
	AttackRate          uint32 `protobuf:"varint,17,opt,name=AttackRate,proto3" json:"AttackRate,omitempty"`
	AttackRange         uint32 `protobuf:"varint,18,opt,name=AttackRange,proto3" json:"AttackRange,omitempty"`
	Spellpower          uint32 `protobuf:"varint,19,opt,name=Spellpower,proto3" json:"Spellpower,omitempty"`
	SpellPenetration    uint32 `protobuf:"varint,20,opt,name=SpellPenetration,proto3" json:"SpellPenetration,omitempty"`
	Physicalpenetration uint32 `protobuf:"varint,21,opt,name=Physicalpenetration,proto3" json:"Physicalpenetration,omitempty"`
	Critrate            uint32 `protobuf:"varint,22,opt,name=Critrate,proto3" json:"Critrate,omitempty"`
	Skill_1             uint32 `protobuf:"varint,23,opt,name=Skill_1,json=Skill1,proto3" json:"Skill_1,omitempty"`
	Skill_2_1           uint32 `protobuf:"varint,24,opt,name=Skill_2_1,json=Skill21,proto3" json:"Skill_2_1,omitempty"`
	Skill_2_2           uint32 `protobuf:"varint,25,opt,name=Skill_2_2,json=Skill22,proto3" json:"Skill_2_2,omitempty"`
	Skill_2_3           uint32 `protobuf:"varint,26,opt,name=Skill_2_3,json=Skill23,proto3" json:"Skill_2_3,omitempty"`
	Skill_3_1           uint32 `protobuf:"varint,27,opt,name=Skill_3_1,json=Skill31,proto3" json:"Skill_3_1,omitempty"`
	Skill_3_2           uint32 `protobuf:"varint,28,opt,name=Skill_3_2,json=Skill32,proto3" json:"Skill_3_2,omitempty"`
	Skill_3_3           uint32 `protobuf:"varint,29,opt,name=Skill_3_3,json=Skill33,proto3" json:"Skill_3_3,omitempty"`
	Fetters_1           uint32 `protobuf:"varint,30,opt,name=Fetters_1,json=Fetters1,proto3" json:"Fetters_1,omitempty"`
	Fetters_2           uint32 `protobuf:"varint,31,opt,name=Fetters_2,json=Fetters2,proto3" json:"Fetters_2,omitempty"`
	Model_1             uint32 `protobuf:"varint,32,opt,name=Model_1,json=Model1,proto3" json:"Model_1,omitempty"`
	Model_2             uint32 `protobuf:"varint,33,opt,name=Model_2,json=Model2,proto3" json:"Model_2,omitempty"`
	Model_3             uint32 `protobuf:"varint,34,opt,name=Model_3,json=Model3,proto3" json:"Model_3,omitempty"`
	Model_4             uint32 `protobuf:"varint,35,opt,name=Model_4,json=Model4,proto3" json:"Model_4,omitempty"`
	Voice_1             uint32 `protobuf:"varint,36,opt,name=Voice_1,json=Voice1,proto3" json:"Voice_1,omitempty"`
	Voice_2             uint32 `protobuf:"varint,37,opt,name=Voice_2,json=Voice2,proto3" json:"Voice_2,omitempty"`
	Voice_3             uint32 `protobuf:"varint,38,opt,name=Voice_3,json=Voice3,proto3" json:"Voice_3,omitempty"`
	Voice_4             uint32 `protobuf:"varint,39,opt,name=Voice_4,json=Voice4,proto3" json:"Voice_4,omitempty"`
	Skin_1              uint32 `protobuf:"varint,40,opt,name=Skin_1,json=Skin1,proto3" json:"Skin_1,omitempty"`
	Skin_2              uint32 `protobuf:"varint,41,opt,name=Skin_2,json=Skin2,proto3" json:"Skin_2,omitempty"`
	Skin_3              uint32 `protobuf:"varint,42,opt,name=Skin_3,json=Skin3,proto3" json:"Skin_3,omitempty"`
	Skin_4              uint32 `protobuf:"varint,43,opt,name=Skin_4,json=Skin4,proto3" json:"Skin_4,omitempty"`
}

func (x *ChessesStatistics) Reset() {
	*x = ChessesStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chesses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessesStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessesStatistics) ProtoMessage() {}

func (x *ChessesStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_chesses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessesStatistics.ProtoReflect.Descriptor instead.
func (*ChessesStatistics) Descriptor() ([]byte, []int) {
	return file_chesses_proto_rawDescGZIP(), []int{0}
}

func (x *ChessesStatistics) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ChessesStatistics) GetNextID() uint32 {
	if x != nil {
		return x.NextID
	}
	return 0
}

func (x *ChessesStatistics) GetStar() uint32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *ChessesStatistics) GetCvlzn() uint32 {
	if x != nil {
		return x.Cvlzn
	}
	return 0
}

func (x *ChessesStatistics) GetSearchable() uint32 {
	if x != nil {
		return x.Searchable
	}
	return 0
}

func (x *ChessesStatistics) GetQuality() uint32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *ChessesStatistics) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ChessesStatistics) GetModel() uint32 {
	if x != nil {
		return x.Model
	}
	return 0
}

func (x *ChessesStatistics) GetAttack() uint32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *ChessesStatistics) GetStatusHealth() uint32 {
	if x != nil {
		return x.StatusHealth
	}
	return 0
}

func (x *ChessesStatistics) GetArmorPhysical() uint32 {
	if x != nil {
		return x.ArmorPhysical
	}
	return 0
}

func (x *ChessesStatistics) GetMagicalResistance() uint32 {
	if x != nil {
		return x.MagicalResistance
	}
	return 0
}

func (x *ChessesStatistics) GetStatusHealthRegen() uint32 {
	if x != nil {
		return x.StatusHealthRegen
	}
	return 0
}

func (x *ChessesStatistics) GetStatusMana() uint32 {
	if x != nil {
		return x.StatusMana
	}
	return 0
}

func (x *ChessesStatistics) GetStatusManaRegen() uint32 {
	if x != nil {
		return x.StatusManaRegen
	}
	return 0
}

func (x *ChessesStatistics) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *ChessesStatistics) GetAttackRate() uint32 {
	if x != nil {
		return x.AttackRate
	}
	return 0
}

func (x *ChessesStatistics) GetAttackRange() uint32 {
	if x != nil {
		return x.AttackRange
	}
	return 0
}

func (x *ChessesStatistics) GetSpellpower() uint32 {
	if x != nil {
		return x.Spellpower
	}
	return 0
}

func (x *ChessesStatistics) GetSpellPenetration() uint32 {
	if x != nil {
		return x.SpellPenetration
	}
	return 0
}

func (x *ChessesStatistics) GetPhysicalpenetration() uint32 {
	if x != nil {
		return x.Physicalpenetration
	}
	return 0
}

func (x *ChessesStatistics) GetCritrate() uint32 {
	if x != nil {
		return x.Critrate
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_1() uint32 {
	if x != nil {
		return x.Skill_1
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_2_1() uint32 {
	if x != nil {
		return x.Skill_2_1
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_2_2() uint32 {
	if x != nil {
		return x.Skill_2_2
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_2_3() uint32 {
	if x != nil {
		return x.Skill_2_3
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_3_1() uint32 {
	if x != nil {
		return x.Skill_3_1
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_3_2() uint32 {
	if x != nil {
		return x.Skill_3_2
	}
	return 0
}

func (x *ChessesStatistics) GetSkill_3_3() uint32 {
	if x != nil {
		return x.Skill_3_3
	}
	return 0
}

func (x *ChessesStatistics) GetFetters_1() uint32 {
	if x != nil {
		return x.Fetters_1
	}
	return 0
}

func (x *ChessesStatistics) GetFetters_2() uint32 {
	if x != nil {
		return x.Fetters_2
	}
	return 0
}

func (x *ChessesStatistics) GetModel_1() uint32 {
	if x != nil {
		return x.Model_1
	}
	return 0
}

func (x *ChessesStatistics) GetModel_2() uint32 {
	if x != nil {
		return x.Model_2
	}
	return 0
}

func (x *ChessesStatistics) GetModel_3() uint32 {
	if x != nil {
		return x.Model_3
	}
	return 0
}

func (x *ChessesStatistics) GetModel_4() uint32 {
	if x != nil {
		return x.Model_4
	}
	return 0
}

func (x *ChessesStatistics) GetVoice_1() uint32 {
	if x != nil {
		return x.Voice_1
	}
	return 0
}

func (x *ChessesStatistics) GetVoice_2() uint32 {
	if x != nil {
		return x.Voice_2
	}
	return 0
}

func (x *ChessesStatistics) GetVoice_3() uint32 {
	if x != nil {
		return x.Voice_3
	}
	return 0
}

func (x *ChessesStatistics) GetVoice_4() uint32 {
	if x != nil {
		return x.Voice_4
	}
	return 0
}

func (x *ChessesStatistics) GetSkin_1() uint32 {
	if x != nil {
		return x.Skin_1
	}
	return 0
}

func (x *ChessesStatistics) GetSkin_2() uint32 {
	if x != nil {
		return x.Skin_2
	}
	return 0
}

func (x *ChessesStatistics) GetSkin_3() uint32 {
	if x != nil {
		return x.Skin_3
	}
	return 0
}

func (x *ChessesStatistics) GetSkin_4() uint32 {
	if x != nil {
		return x.Skin_4
	}
	return 0
}

type ChessesStatisticsCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[uint32]*ChessesStatistics `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChessesStatisticsCfg) Reset() {
	*x = ChessesStatisticsCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chesses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessesStatisticsCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessesStatisticsCfg) ProtoMessage() {}

func (x *ChessesStatisticsCfg) ProtoReflect() protoreflect.Message {
	mi := &file_chesses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessesStatisticsCfg.ProtoReflect.Descriptor instead.
func (*ChessesStatisticsCfg) Descriptor() ([]byte, []int) {
	return file_chesses_proto_rawDescGZIP(), []int{1}
}

func (x *ChessesStatisticsCfg) GetDatas() map[uint32]*ChessesStatistics {
	if x != nil {
		return x.Datas
	}
	return nil
}

type ChessesSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Price  uint32 `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	Blue   uint32 `protobuf:"varint,3,opt,name=Blue,proto3" json:"Blue,omitempty"`
	Purple uint32 `protobuf:"varint,4,opt,name=Purple,proto3" json:"Purple,omitempty"`
	Orange uint32 `protobuf:"varint,5,opt,name=Orange,proto3" json:"Orange,omitempty"`
}

func (x *ChessesSearch) Reset() {
	*x = ChessesSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chesses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessesSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessesSearch) ProtoMessage() {}

func (x *ChessesSearch) ProtoReflect() protoreflect.Message {
	mi := &file_chesses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessesSearch.ProtoReflect.Descriptor instead.
func (*ChessesSearch) Descriptor() ([]byte, []int) {
	return file_chesses_proto_rawDescGZIP(), []int{2}
}

func (x *ChessesSearch) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ChessesSearch) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ChessesSearch) GetBlue() uint32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

func (x *ChessesSearch) GetPurple() uint32 {
	if x != nil {
		return x.Purple
	}
	return 0
}

func (x *ChessesSearch) GetOrange() uint32 {
	if x != nil {
		return x.Orange
	}
	return 0
}

type ChessesSearchCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[uint32]*ChessesSearch `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChessesSearchCfg) Reset() {
	*x = ChessesSearchCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chesses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessesSearchCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessesSearchCfg) ProtoMessage() {}

func (x *ChessesSearchCfg) ProtoReflect() protoreflect.Message {
	mi := &file_chesses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessesSearchCfg.ProtoReflect.Descriptor instead.
func (*ChessesSearchCfg) Descriptor() ([]byte, []int) {
	return file_chesses_proto_rawDescGZIP(), []int{3}
}

func (x *ChessesSearchCfg) GetDatas() map[uint32]*ChessesSearch {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_chesses_proto protoreflect.FileDescriptor

var file_chesses_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xe4, 0x09, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4e,
	0x65, 0x78, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x76, 0x6c,
	0x7a, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x43, 0x76, 0x6c, 0x7a, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x22, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x50,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x61, 0x67, 0x69, 0x63,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x4d, 0x61, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x6e,
	0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x61, 0x6e, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x6e,
	0x61, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x70, 0x65, 0x6c, 0x6c,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x50, 0x65,
	0x6e, 0x65, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x50, 0x65, 0x6e, 0x65, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x30, 0x0a, 0x13, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x70, 0x65, 0x6e,
	0x65, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13,
	0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x70, 0x65, 0x6e, 0x65, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x72, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x31, 0x12, 0x1a, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x32, 0x5f, 0x31, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x32, 0x31, 0x12, 0x1a, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x32, 0x5f,
	0x32, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x32, 0x32,
	0x12, 0x1a, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x32, 0x5f, 0x33, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x32, 0x33, 0x12, 0x1a, 0x0a, 0x09,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x33, 0x5f, 0x31, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x33, 0x31, 0x12, 0x1a, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x33, 0x5f, 0x32, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x33, 0x5f,
	0x33, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x33, 0x33,
	0x12, 0x1b, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x31, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x31, 0x12, 0x1b, 0x0a,
	0x09, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x32, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x31, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x31, 0x12, 0x17, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x32, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x32, 0x12, 0x17, 0x0a, 0x07,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x33, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x33, 0x12, 0x17, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x34,
	0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x34, 0x12, 0x17,
	0x0a, 0x07, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x31, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x12, 0x17, 0x0a, 0x07, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x32, 0x18, 0x25, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x32,
	0x12, 0x17, 0x0a, 0x07, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x33, 0x18, 0x26, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x12, 0x17, 0x0a, 0x07, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x34, 0x18, 0x27, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x34, 0x12, 0x15, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6e, 0x5f, 0x31, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x53, 0x6b, 0x69, 0x6e, 0x31, 0x12, 0x15, 0x0a, 0x06, 0x53, 0x6b, 0x69,
	0x6e, 0x5f, 0x32, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x6b, 0x69, 0x6e, 0x32,
	0x12, 0x15, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6e, 0x5f, 0x33, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x53, 0x6b, 0x69, 0x6e, 0x33, 0x12, 0x15, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6e, 0x5f,
	0x34, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x6b, 0x69, 0x6e, 0x34, 0x22, 0xaa,
	0x01, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x43, 0x66, 0x67, 0x12, 0x3d, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x43, 0x66, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x53, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x79, 0x0a, 0x0d, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x42, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x75, 0x72, 0x70, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x75, 0x72, 0x70, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x4f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x66, 0x67, 0x12, 0x39, 0x0a, 0x05, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x43, 0x66, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x4f, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x65, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chesses_proto_rawDescOnce sync.Once
	file_chesses_proto_rawDescData = file_chesses_proto_rawDesc
)

func file_chesses_proto_rawDescGZIP() []byte {
	file_chesses_proto_rawDescOnce.Do(func() {
		file_chesses_proto_rawDescData = protoimpl.X.CompressGZIP(file_chesses_proto_rawDescData)
	})
	return file_chesses_proto_rawDescData
}

var file_chesses_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chesses_proto_goTypes = []interface{}{
	(*ChessesStatistics)(nil),    // 0: config.ChessesStatistics
	(*ChessesStatisticsCfg)(nil), // 1: config.ChessesStatisticsCfg
	(*ChessesSearch)(nil),        // 2: config.ChessesSearch
	(*ChessesSearchCfg)(nil),     // 3: config.ChessesSearchCfg
	nil,                          // 4: config.ChessesStatisticsCfg.DatasEntry
	nil,                          // 5: config.ChessesSearchCfg.DatasEntry
}
var file_chesses_proto_depIdxs = []int32{
	4, // 0: config.ChessesStatisticsCfg.datas:type_name -> config.ChessesStatisticsCfg.DatasEntry
	5, // 1: config.ChessesSearchCfg.datas:type_name -> config.ChessesSearchCfg.DatasEntry
	0, // 2: config.ChessesStatisticsCfg.DatasEntry.value:type_name -> config.ChessesStatistics
	2, // 3: config.ChessesSearchCfg.DatasEntry.value:type_name -> config.ChessesSearch
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chesses_proto_init() }
func file_chesses_proto_init() {
	if File_chesses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chesses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessesStatistics); i {
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
		file_chesses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessesStatisticsCfg); i {
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
		file_chesses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessesSearch); i {
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
		file_chesses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessesSearchCfg); i {
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
			RawDescriptor: file_chesses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chesses_proto_goTypes,
		DependencyIndexes: file_chesses_proto_depIdxs,
		MessageInfos:      file_chesses_proto_msgTypes,
	}.Build()
	File_chesses_proto = out.File
	file_chesses_proto_rawDesc = nil
	file_chesses_proto_goTypes = nil
	file_chesses_proto_depIdxs = nil
}
