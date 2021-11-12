// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: cs_msgid.proto

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

type C2SLogicMsgId int32

const (
	C2SLogicMsgId_c2s_base_min_msg_id              C2SLogicMsgId = 0
	C2SLogicMsgId_c2s_base_max_msg_id              C2SLogicMsgId = 100
	C2SLogicMsgId_c2s_verify_conn_req_id           C2SLogicMsgId = 201
	C2SLogicMsgId_s2c_verify_conn_fail_res_id      C2SLogicMsgId = 202
	C2SLogicMsgId_s2c_get_role_list_ntf_id         C2SLogicMsgId = 203
	C2SLogicMsgId_c2s_create_role_req_id           C2SLogicMsgId = 204
	C2SLogicMsgId_s2c_create_role_res_id           C2SLogicMsgId = 205
	C2SLogicMsgId_c2s_game_login_req_id            C2SLogicMsgId = 206
	C2SLogicMsgId_s2c_game_login_res_id            C2SLogicMsgId = 207
	C2SLogicMsgId_c2s_heartbeat_req_id             C2SLogicMsgId = 208
	C2SLogicMsgId_s2c_heartbeat_res_id             C2SLogicMsgId = 209
	C2SLogicMsgId_c2s_exit_game_req_id             C2SLogicMsgId = 210
	C2SLogicMsgId_s2c_exit_game_res_id             C2SLogicMsgId = 211
	C2SLogicMsgId_s2c_servertime_ntf_id            C2SLogicMsgId = 212
	C2SLogicMsgId_s2c_tip_ntf_id                   C2SLogicMsgId = 213
	C2SLogicMsgId_c2s_create_team_req_id           C2SLogicMsgId = 300
	C2SLogicMsgId_s2c_create_team_res_id           C2SLogicMsgId = 301
	C2SLogicMsgId_c2s_invite_team_req_id           C2SLogicMsgId = 302
	C2SLogicMsgId_s2c_invite_team_confirm_req_id   C2SLogicMsgId = 303
	C2SLogicMsgId_c2s_invite_team_confirm_res_id   C2SLogicMsgId = 304
	C2SLogicMsgId_s2c_team_ntf_id                  C2SLogicMsgId = 305
	C2SLogicMsgId_c2s_leave_team_req_id            C2SLogicMsgId = 306
	C2SLogicMsgId_c2s_start_match_req_id           C2SLogicMsgId = 307
	C2SLogicMsgId_s2c_room_confirm_req_id          C2SLogicMsgId = 308
	C2SLogicMsgId_c2s_room_confirm_res_id          C2SLogicMsgId = 309
	C2SLogicMsgId_s2c_restart_match_ntf_id         C2SLogicMsgId = 310
	C2SLogicMsgId_c2s_stop_match_req_id            C2SLogicMsgId = 311
	C2SLogicMsgId_s2c_room_member_ntf_id           C2SLogicMsgId = 312
	C2SLogicMsgId_s2c_room_ok_ntf_id               C2SLogicMsgId = 313
	C2SLogicMsgId_c2s_select_master_in_room_req_id C2SLogicMsgId = 314
	C2SLogicMsgId_c2s_lock_master_in_room_req_id   C2SLogicMsgId = 315
	C2SLogicMsgId_s2c_scene_ok_ntf_id              C2SLogicMsgId = 330
	C2SLogicMsgId_c2s_battle_load_process_req_id   C2SLogicMsgId = 316
	C2SLogicMsgId_s2c_battle_load_process_ntf_id   C2SLogicMsgId = 317
	C2SLogicMsgId_c2s_refresh_chess_req_id         C2SLogicMsgId = 331
	C2SLogicMsgId_s2c_refresh_chess_res_id         C2SLogicMsgId = 332
	C2SLogicMsgId_c2s_buy_chess_req_id             C2SLogicMsgId = 333
	C2SLogicMsgId_s2c_wait_region_chess_ntf_id     C2SLogicMsgId = 334
	C2SLogicMsgId_c2s_select_battle_chess_req_id   C2SLogicMsgId = 335
	C2SLogicMsgId_c2s_sell_chess_req_id            C2SLogicMsgId = 336
)

// Enum value maps for C2SLogicMsgId.
var (
	C2SLogicMsgId_name = map[int32]string{
		0:   "c2s_base_min_msg_id",
		100: "c2s_base_max_msg_id",
		201: "c2s_verify_conn_req_id",
		202: "s2c_verify_conn_fail_res_id",
		203: "s2c_get_role_list_ntf_id",
		204: "c2s_create_role_req_id",
		205: "s2c_create_role_res_id",
		206: "c2s_game_login_req_id",
		207: "s2c_game_login_res_id",
		208: "c2s_heartbeat_req_id",
		209: "s2c_heartbeat_res_id",
		210: "c2s_exit_game_req_id",
		211: "s2c_exit_game_res_id",
		212: "s2c_servertime_ntf_id",
		213: "s2c_tip_ntf_id",
		300: "c2s_create_team_req_id",
		301: "s2c_create_team_res_id",
		302: "c2s_invite_team_req_id",
		303: "s2c_invite_team_confirm_req_id",
		304: "c2s_invite_team_confirm_res_id",
		305: "s2c_team_ntf_id",
		306: "c2s_leave_team_req_id",
		307: "c2s_start_match_req_id",
		308: "s2c_room_confirm_req_id",
		309: "c2s_room_confirm_res_id",
		310: "s2c_restart_match_ntf_id",
		311: "c2s_stop_match_req_id",
		312: "s2c_room_member_ntf_id",
		313: "s2c_room_ok_ntf_id",
		314: "c2s_select_master_in_room_req_id",
		315: "c2s_lock_master_in_room_req_id",
		330: "s2c_scene_ok_ntf_id",
		316: "c2s_battle_load_process_req_id",
		317: "s2c_battle_load_process_ntf_id",
		331: "c2s_refresh_chess_req_id",
		332: "s2c_refresh_chess_res_id",
		333: "c2s_buy_chess_req_id",
		334: "s2c_wait_region_chess_ntf_id",
		335: "c2s_select_battle_chess_req_id",
		336: "c2s_sell_chess_req_id",
	}
	C2SLogicMsgId_value = map[string]int32{
		"c2s_base_min_msg_id":              0,
		"c2s_base_max_msg_id":              100,
		"c2s_verify_conn_req_id":           201,
		"s2c_verify_conn_fail_res_id":      202,
		"s2c_get_role_list_ntf_id":         203,
		"c2s_create_role_req_id":           204,
		"s2c_create_role_res_id":           205,
		"c2s_game_login_req_id":            206,
		"s2c_game_login_res_id":            207,
		"c2s_heartbeat_req_id":             208,
		"s2c_heartbeat_res_id":             209,
		"c2s_exit_game_req_id":             210,
		"s2c_exit_game_res_id":             211,
		"s2c_servertime_ntf_id":            212,
		"s2c_tip_ntf_id":                   213,
		"c2s_create_team_req_id":           300,
		"s2c_create_team_res_id":           301,
		"c2s_invite_team_req_id":           302,
		"s2c_invite_team_confirm_req_id":   303,
		"c2s_invite_team_confirm_res_id":   304,
		"s2c_team_ntf_id":                  305,
		"c2s_leave_team_req_id":            306,
		"c2s_start_match_req_id":           307,
		"s2c_room_confirm_req_id":          308,
		"c2s_room_confirm_res_id":          309,
		"s2c_restart_match_ntf_id":         310,
		"c2s_stop_match_req_id":            311,
		"s2c_room_member_ntf_id":           312,
		"s2c_room_ok_ntf_id":               313,
		"c2s_select_master_in_room_req_id": 314,
		"c2s_lock_master_in_room_req_id":   315,
		"s2c_scene_ok_ntf_id":              330,
		"c2s_battle_load_process_req_id":   316,
		"s2c_battle_load_process_ntf_id":   317,
		"c2s_refresh_chess_req_id":         331,
		"s2c_refresh_chess_res_id":         332,
		"c2s_buy_chess_req_id":             333,
		"s2c_wait_region_chess_ntf_id":     334,
		"c2s_select_battle_chess_req_id":   335,
		"c2s_sell_chess_req_id":            336,
	}
)

func (x C2SLogicMsgId) Enum() *C2SLogicMsgId {
	p := new(C2SLogicMsgId)
	*p = x
	return p
}

func (x C2SLogicMsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (C2SLogicMsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_cs_msgid_proto_enumTypes[0].Descriptor()
}

func (C2SLogicMsgId) Type() protoreflect.EnumType {
	return &file_cs_msgid_proto_enumTypes[0]
}

func (x C2SLogicMsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use C2SLogicMsgId.Descriptor instead.
func (C2SLogicMsgId) EnumDescriptor() ([]byte, []int) {
	return file_cs_msgid_proto_rawDescGZIP(), []int{0}
}

var File_cs_msgid_proto protoreflect.FileDescriptor

var file_cs_msgid_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x2a, 0xb8, 0x09, 0x0a, 0x0d, 0x43, 0x32, 0x53, 0x4c, 0x6f, 0x67, 0x69,
	0x63, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x63, 0x32, 0x73, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x63, 0x32, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f,
	0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x10, 0x64, 0x12, 0x1b, 0x0a, 0x16, 0x63, 0x32, 0x73, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x5f,
	0x69, 0x64, 0x10, 0xc9, 0x01, 0x12, 0x20, 0x0a, 0x1b, 0x73, 0x32, 0x63, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x5f, 0x69, 0x64, 0x10, 0xca, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x73, 0x32, 0x63, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x74, 0x66,
	0x5f, 0x69, 0x64, 0x10, 0xcb, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x63, 0x32, 0x73, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64,
	0x10, 0xcc, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x73, 0x32, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x10, 0xcd, 0x01,
	0x12, 0x1a, 0x0a, 0x15, 0x63, 0x32, 0x73, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xce, 0x01, 0x12, 0x1a, 0x0a, 0x15,
	0x73, 0x32, 0x63, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x73, 0x5f, 0x69, 0x64, 0x10, 0xcf, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x63, 0x32, 0x73, 0x5f,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64,
	0x10, 0xd0, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x73, 0x32, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x10, 0xd1, 0x01, 0x12, 0x19,
	0x0a, 0x14, 0x63, 0x32, 0x73, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xd2, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x73, 0x32, 0x63,
	0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x69,
	0x64, 0x10, 0xd3, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x73, 0x32, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10, 0xd4, 0x01,
	0x12, 0x13, 0x0a, 0x0e, 0x73, 0x32, 0x63, 0x5f, 0x74, 0x69, 0x70, 0x5f, 0x6e, 0x74, 0x66, 0x5f,
	0x69, 0x64, 0x10, 0xd5, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x63, 0x32, 0x73, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10,
	0xac, 0x02, 0x12, 0x1b, 0x0a, 0x16, 0x73, 0x32, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x10, 0xad, 0x02, 0x12,
	0x1b, 0x0a, 0x16, 0x63, 0x32, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xae, 0x02, 0x12, 0x23, 0x0a, 0x1e,
	0x73, 0x32, 0x63, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xaf,
	0x02, 0x12, 0x23, 0x0a, 0x1e, 0x63, 0x32, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x73,
	0x5f, 0x69, 0x64, 0x10, 0xb0, 0x02, 0x12, 0x14, 0x0a, 0x0f, 0x73, 0x32, 0x63, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10, 0xb1, 0x02, 0x12, 0x1a, 0x0a, 0x15,
	0x63, 0x32, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72,
	0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xb2, 0x02, 0x12, 0x1b, 0x0a, 0x16, 0x63, 0x32, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x5f,
	0x69, 0x64, 0x10, 0xb3, 0x02, 0x12, 0x1c, 0x0a, 0x17, 0x73, 0x32, 0x63, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64,
	0x10, 0xb4, 0x02, 0x12, 0x1c, 0x0a, 0x17, 0x63, 0x32, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x10, 0xb5,
	0x02, 0x12, 0x1d, 0x0a, 0x18, 0x73, 0x32, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10, 0xb6, 0x02,
	0x12, 0x1a, 0x0a, 0x15, 0x63, 0x32, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xb7, 0x02, 0x12, 0x1b, 0x0a, 0x16,
	0x73, 0x32, 0x63, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10, 0xb8, 0x02, 0x12, 0x17, 0x0a, 0x12, 0x73, 0x32, 0x63,
	0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6f, 0x6b, 0x5f, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10,
	0xb9, 0x02, 0x12, 0x25, 0x0a, 0x20, 0x63, 0x32, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xba, 0x02, 0x12, 0x23, 0x0a, 0x1e, 0x63, 0x32, 0x73,
	0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x5f,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xbb, 0x02, 0x12, 0x18,
	0x0a, 0x13, 0x73, 0x32, 0x63, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x6f, 0x6b, 0x5f, 0x6e,
	0x74, 0x66, 0x5f, 0x69, 0x64, 0x10, 0xca, 0x02, 0x12, 0x23, 0x0a, 0x1e, 0x63, 0x32, 0x73, 0x5f,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xbc, 0x02, 0x12, 0x23, 0x0a,
	0x1e, 0x73, 0x32, 0x63, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10,
	0xbd, 0x02, 0x12, 0x1d, 0x0a, 0x18, 0x63, 0x32, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xcb,
	0x02, 0x12, 0x1d, 0x0a, 0x18, 0x73, 0x32, 0x63, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x10, 0xcc, 0x02,
	0x12, 0x19, 0x0a, 0x14, 0x63, 0x32, 0x73, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x68, 0x65, 0x73,
	0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xcd, 0x02, 0x12, 0x21, 0x0a, 0x1c, 0x73,
	0x32, 0x63, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x68, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x74, 0x66, 0x5f, 0x69, 0x64, 0x10, 0xce, 0x02, 0x12, 0x23,
	0x0a, 0x1e, 0x63, 0x32, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64,
	0x10, 0xcf, 0x02, 0x12, 0x1a, 0x0a, 0x15, 0x63, 0x32, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f,
	0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x10, 0xd0, 0x02, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs_msgid_proto_rawDescOnce sync.Once
	file_cs_msgid_proto_rawDescData = file_cs_msgid_proto_rawDesc
)

func file_cs_msgid_proto_rawDescGZIP() []byte {
	file_cs_msgid_proto_rawDescOnce.Do(func() {
		file_cs_msgid_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs_msgid_proto_rawDescData)
	})
	return file_cs_msgid_proto_rawDescData
}

var file_cs_msgid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cs_msgid_proto_goTypes = []interface{}{
	(C2SLogicMsgId)(0), // 0: pb.C2SLogicMsgId
}
var file_cs_msgid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cs_msgid_proto_init() }
func file_cs_msgid_proto_init() {
	if File_cs_msgid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cs_msgid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs_msgid_proto_goTypes,
		DependencyIndexes: file_cs_msgid_proto_depIdxs,
		EnumInfos:         file_cs_msgid_proto_enumTypes,
	}.Build()
	File_cs_msgid_proto = out.File
	file_cs_msgid_proto_rawDesc = nil
	file_cs_msgid_proto_goTypes = nil
	file_cs_msgid_proto_depIdxs = nil
}