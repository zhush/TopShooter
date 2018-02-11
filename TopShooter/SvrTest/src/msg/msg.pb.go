// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	CS_LoginReq
	RoleBaseInfo
	SC_LoginResponse
	CS_CreateRoleReq
	SC_CreateRoleAck
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSG_ID int32

const (
	MSG_ID_ELogin_Req MSG_ID = 1001
	MSG_ID_ELogin_Ack MSG_ID = 1002
)

var MSG_ID_name = map[int32]string{
	1001: "ELogin_Req",
	1002: "ELogin_Ack",
}
var MSG_ID_value = map[string]int32{
	"ELogin_Req": 1001,
	"ELogin_Ack": 1002,
}

func (x MSG_ID) Enum() *MSG_ID {
	p := new(MSG_ID)
	*p = x
	return p
}
func (x MSG_ID) String() string {
	return proto.EnumName(MSG_ID_name, int32(x))
}
func (x *MSG_ID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MSG_ID_value, data, "MSG_ID")
	if err != nil {
		return err
	}
	*x = MSG_ID(value)
	return nil
}
func (MSG_ID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ELoginResult int32

const (
	ELoginResult_Succeed         ELoginResult = 0
	ELoginResult_InvalidAccOrPwd ELoginResult = 1
	ELoginResult_ServerClosed    ELoginResult = 2
)

var ELoginResult_name = map[int32]string{
	0: "Succeed",
	1: "InvalidAccOrPwd",
	2: "ServerClosed",
}
var ELoginResult_value = map[string]int32{
	"Succeed":         0,
	"InvalidAccOrPwd": 1,
	"ServerClosed":    2,
}

func (x ELoginResult) Enum() *ELoginResult {
	p := new(ELoginResult)
	*p = x
	return p
}
func (x ELoginResult) String() string {
	return proto.EnumName(ELoginResult_name, int32(x))
}
func (x *ELoginResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ELoginResult_value, data, "ELoginResult")
	if err != nil {
		return err
	}
	*x = ELoginResult(value)
	return nil
}
func (ELoginResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EPlatForm int32

const (
	EPlatForm_Android EPlatForm = 0
	EPlatForm_Ios     EPlatForm = 1
	EPlatForm_Windows EPlatForm = 2
)

var EPlatForm_name = map[int32]string{
	0: "Android",
	1: "Ios",
	2: "Windows",
}
var EPlatForm_value = map[string]int32{
	"Android": 0,
	"Ios":     1,
	"Windows": 2,
}

func (x EPlatForm) Enum() *EPlatForm {
	p := new(EPlatForm)
	*p = x
	return p
}
func (x EPlatForm) String() string {
	return proto.EnumName(EPlatForm_name, int32(x))
}
func (x *EPlatForm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EPlatForm_value, data, "EPlatForm")
	if err != nil {
		return err
	}
	*x = EPlatForm(value)
	return nil
}
func (EPlatForm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type EResultCreateRole int32

const (
	EResultCreateRole_eSucceed         EResultCreateRole = 0
	EResultCreateRole_eNicknameExsists EResultCreateRole = 1
	EResultCreateRole_eRoleCountLimit  EResultCreateRole = 2
	EResultCreateRole_eUnknownError    EResultCreateRole = 3
)

var EResultCreateRole_name = map[int32]string{
	0: "eSucceed",
	1: "eNicknameExsists",
	2: "eRoleCountLimit",
	3: "eUnknownError",
}
var EResultCreateRole_value = map[string]int32{
	"eSucceed":         0,
	"eNicknameExsists": 1,
	"eRoleCountLimit":  2,
	"eUnknownError":    3,
}

func (x EResultCreateRole) Enum() *EResultCreateRole {
	p := new(EResultCreateRole)
	*p = x
	return p
}
func (x EResultCreateRole) String() string {
	return proto.EnumName(EResultCreateRole_name, int32(x))
}
func (x *EResultCreateRole) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EResultCreateRole_value, data, "EResultCreateRole")
	if err != nil {
		return err
	}
	*x = EResultCreateRole(value)
	return nil
}
func (EResultCreateRole) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CS_LoginReq struct {
	AccName          *string    `protobuf:"bytes,1,req,name=AccName" json:"AccName,omitempty"`
	AccPassword      *string    `protobuf:"bytes,2,req,name=AccPassword" json:"AccPassword,omitempty"`
	PlatForm         *EPlatForm `protobuf:"varint,3,req,name=PlatForm,enum=msg.EPlatForm" json:"PlatForm,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *CS_LoginReq) Reset()                    { *m = CS_LoginReq{} }
func (m *CS_LoginReq) String() string            { return proto.CompactTextString(m) }
func (*CS_LoginReq) ProtoMessage()               {}
func (*CS_LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CS_LoginReq) GetAccName() string {
	if m != nil && m.AccName != nil {
		return *m.AccName
	}
	return ""
}

func (m *CS_LoginReq) GetAccPassword() string {
	if m != nil && m.AccPassword != nil {
		return *m.AccPassword
	}
	return ""
}

func (m *CS_LoginReq) GetPlatForm() EPlatForm {
	if m != nil && m.PlatForm != nil {
		return *m.PlatForm
	}
	return EPlatForm_Android
}

type RoleBaseInfo struct {
	NickName         *string `protobuf:"bytes,1,req,name=NickName" json:"NickName,omitempty"`
	TemplateId       *int32  `protobuf:"varint,2,req,name=TemplateId" json:"TemplateId,omitempty"`
	Level            *int32  `protobuf:"varint,3,req,name=Level" json:"Level,omitempty"`
	Sex              *int32  `protobuf:"varint,4,req,name=Sex" json:"Sex,omitempty"`
	Gold             *int64  `protobuf:"varint,5,req,name=Gold" json:"Gold,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoleBaseInfo) Reset()                    { *m = RoleBaseInfo{} }
func (m *RoleBaseInfo) String() string            { return proto.CompactTextString(m) }
func (*RoleBaseInfo) ProtoMessage()               {}
func (*RoleBaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RoleBaseInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *RoleBaseInfo) GetTemplateId() int32 {
	if m != nil && m.TemplateId != nil {
		return *m.TemplateId
	}
	return 0
}

func (m *RoleBaseInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *RoleBaseInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *RoleBaseInfo) GetGold() int64 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

type SC_LoginResponse struct {
	LoginResult      *ELoginResult   `protobuf:"varint,1,req,name=LoginResult,enum=msg.ELoginResult" json:"LoginResult,omitempty"`
	PlayerBaseInfo   []*RoleBaseInfo `protobuf:"bytes,2,rep,name=PlayerBaseInfo" json:"PlayerBaseInfo,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SC_LoginResponse) Reset()                    { *m = SC_LoginResponse{} }
func (m *SC_LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*SC_LoginResponse) ProtoMessage()               {}
func (*SC_LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SC_LoginResponse) GetLoginResult() ELoginResult {
	if m != nil && m.LoginResult != nil {
		return *m.LoginResult
	}
	return ELoginResult_Succeed
}

func (m *SC_LoginResponse) GetPlayerBaseInfo() []*RoleBaseInfo {
	if m != nil {
		return m.PlayerBaseInfo
	}
	return nil
}

type CS_CreateRoleReq struct {
	NickName         *string `protobuf:"bytes,1,req,name=NickName" json:"NickName,omitempty"`
	TemplateId       *int32  `protobuf:"varint,2,req,name=TemplateId" json:"TemplateId,omitempty"`
	Sex              *int32  `protobuf:"varint,3,req,name=Sex" json:"Sex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CS_CreateRoleReq) Reset()                    { *m = CS_CreateRoleReq{} }
func (m *CS_CreateRoleReq) String() string            { return proto.CompactTextString(m) }
func (*CS_CreateRoleReq) ProtoMessage()               {}
func (*CS_CreateRoleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CS_CreateRoleReq) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *CS_CreateRoleReq) GetTemplateId() int32 {
	if m != nil && m.TemplateId != nil {
		return *m.TemplateId
	}
	return 0
}

func (m *CS_CreateRoleReq) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

type SC_CreateRoleAck struct {
	Result           *EResultCreateRole `protobuf:"varint,1,req,name=Result,enum=msg.EResultCreateRole" json:"Result,omitempty"`
	RoleInfo         []*RoleBaseInfo    `protobuf:"bytes,2,rep,name=RoleInfo" json:"RoleInfo,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SC_CreateRoleAck) Reset()                    { *m = SC_CreateRoleAck{} }
func (m *SC_CreateRoleAck) String() string            { return proto.CompactTextString(m) }
func (*SC_CreateRoleAck) ProtoMessage()               {}
func (*SC_CreateRoleAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SC_CreateRoleAck) GetResult() EResultCreateRole {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return EResultCreateRole_eSucceed
}

func (m *SC_CreateRoleAck) GetRoleInfo() []*RoleBaseInfo {
	if m != nil {
		return m.RoleInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CS_LoginReq)(nil), "msg.CS_LoginReq")
	proto.RegisterType((*RoleBaseInfo)(nil), "msg.RoleBaseInfo")
	proto.RegisterType((*SC_LoginResponse)(nil), "msg.SC_LoginResponse")
	proto.RegisterType((*CS_CreateRoleReq)(nil), "msg.CS_CreateRoleReq")
	proto.RegisterType((*SC_CreateRoleAck)(nil), "msg.SC_CreateRoleAck")
	proto.RegisterEnum("msg.MSG_ID", MSG_ID_name, MSG_ID_value)
	proto.RegisterEnum("msg.ELoginResult", ELoginResult_name, ELoginResult_value)
	proto.RegisterEnum("msg.EPlatForm", EPlatForm_name, EPlatForm_value)
	proto.RegisterEnum("msg.EResultCreateRole", EResultCreateRole_name, EResultCreateRole_value)
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x69, 0xdc, 0xd4, 0xc9, 0xd8, 0x4d, 0x37, 0x0b, 0x42, 0x3e, 0x56, 0x46, 0x42, 0x25,
	0x87, 0x1c, 0xfa, 0x06, 0xa9, 0x09, 0x95, 0xa5, 0x50, 0xa2, 0x9a, 0x8a, 0x1b, 0x96, 0x65, 0x0f,
	0x95, 0x15, 0x7b, 0x27, 0xec, 0x3a, 0x49, 0x79, 0x64, 0x78, 0x0a, 0xc6, 0x76, 0x43, 0x1c, 0x21,
	0x21, 0x4e, 0xf6, 0xfe, 0x3b, 0xfb, 0xcf, 0x37, 0xff, 0xc0, 0xb0, 0x34, 0x8f, 0xd3, 0xb5, 0xa6,
	0x8a, 0xa4, 0xc5, 0xbf, 0xfe, 0x03, 0x38, 0x41, 0x14, 0x2f, 0xe8, 0x31, 0x57, 0xf7, 0xf8, 0x5d,
	0x5e, 0x80, 0x3d, 0x4b, 0xd3, 0xbb, 0xa4, 0x44, 0xef, 0xe4, 0xb2, 0x77, 0x35, 0x94, 0x2f, 0xc1,
	0x61, 0x61, 0x99, 0x18, 0xb3, 0x23, 0x9d, 0x79, 0xbd, 0x46, 0xbc, 0x84, 0xc1, 0xb2, 0x48, 0xaa,
	0x0f, 0xa4, 0x4b, 0xcf, 0x62, 0x65, 0x74, 0x3d, 0x9a, 0xd6, 0xbe, 0xf3, 0xbd, 0xea, 0x7f, 0x05,
	0xf7, 0x9e, 0x0a, 0xbc, 0x49, 0x0c, 0x86, 0xea, 0x1b, 0x49, 0x01, 0x83, 0xbb, 0x3c, 0x5d, 0x75,
	0x8c, 0x25, 0xc0, 0x67, 0x2c, 0xd7, 0xfc, 0x00, 0xc3, 0xd6, 0xb7, 0x2f, 0xcf, 0xa1, 0xbf, 0xc0,
	0x2d, 0x16, 0x8d, 0x69, 0x5f, 0x3a, 0x60, 0x45, 0xf8, 0xe4, 0x9d, 0x36, 0x07, 0x17, 0x4e, 0x6f,
	0xa9, 0xc8, 0xbc, 0x3e, 0x9f, 0x2c, 0x1f, 0x41, 0x44, 0xc1, 0x1e, 0xdb, 0xac, 0x49, 0x19, 0x94,
	0x6f, 0xc1, 0xd9, 0x0b, 0x9b, 0xa2, 0x6a, 0xda, 0x8c, 0xae, 0xc7, 0x2d, 0x58, 0xe7, 0x42, 0xbe,
	0x83, 0x11, 0x73, 0xfe, 0x40, 0xbd, 0xa7, 0xe3, 0xee, 0xd6, 0x95, 0xf3, 0x5c, 0xda, 0xc5, 0xf6,
	0xe7, 0x20, 0x38, 0x9d, 0x40, 0x23, 0x53, 0xd6, 0x17, 0x75, 0x44, 0xff, 0x37, 0xca, 0x33, 0x7b,
	0x33, 0x88, 0x1f, 0x37, 0xb4, 0x07, 0x9b, 0x59, 0xba, 0x62, 0xda, 0xb3, 0x23, 0xd0, 0xd7, 0x2d,
	0x68, 0xab, 0x1d, 0x4a, 0xe5, 0x1b, 0x18, 0xd4, 0xdf, 0x7f, 0x72, 0x4e, 0x26, 0x70, 0xf6, 0x31,
	0xba, 0x8d, 0xc3, 0xf7, 0xbc, 0x40, 0x68, 0x87, 0x8d, 0x99, 0x55, 0xfc, 0xb4, 0x3b, 0x02, 0x77,
	0x15, 0xbf, 0xec, 0xc9, 0x0d, 0xb8, 0x47, 0x71, 0x38, 0x60, 0x47, 0x9b, 0x34, 0x45, 0xcc, 0xc4,
	0x0b, 0x5e, 0xf7, 0x45, 0xa8, 0xb6, 0x49, 0x91, 0x67, 0xbc, 0xf5, 0x4f, 0x7a, 0xb9, 0xcb, 0xc4,
	0x09, 0x4f, 0xec, 0x46, 0xa8, 0xb7, 0xa8, 0x83, 0x82, 0x0c, 0x97, 0xf5, 0x26, 0x53, 0x18, 0xfe,
	0xd9, 0x75, 0x6d, 0x30, 0x53, 0x99, 0xa6, 0xbc, 0x36, 0xb0, 0xc1, 0x0a, 0xc9, 0xf0, 0x23, 0x56,
	0xbf, 0xe4, 0x2a, 0xa3, 0x9d, 0xe1, 0xfa, 0x18, 0xc6, 0x7f, 0x4f, 0xe6, 0xc2, 0x00, 0x0f, 0x9d,
	0x5f, 0x81, 0xc0, 0x3a, 0x57, 0xc5, 0xb9, 0xce, 0x9f, 0x4c, 0x6e, 0xaa, 0xda, 0x85, 0x79, 0x9a,
	0xe2, 0x80, 0x36, 0xaa, 0x5a, 0xe4, 0x65, 0x5e, 0x89, 0x9e, 0x1c, 0xc3, 0x39, 0x3e, 0xa8, 0x95,
	0xa2, 0x9d, 0x9a, 0x6b, 0x4d, 0x5a, 0x58, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xae, 0x51, 0xdc,
	0x99, 0xd7, 0x02, 0x00, 0x00,
}