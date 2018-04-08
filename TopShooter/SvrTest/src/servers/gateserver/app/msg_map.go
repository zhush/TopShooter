package app

import (
	"msg"
	"reflect"

	proto "github.com/golang/protobuf/proto"
)

//这个文件是msg的注册解析类，主要是protobuf中消息id和结构对应关系.

var MsgTypeMaps map[uint16]reflect.Type

func RegisterMsgType(msgId uint16, x proto.Message) {
	t := reflect.TypeOf(x)
	handlerMsgMaps[msgId] = t
}

func init() {
	MsgTypeMaps = make(map[uint16]reflect.Type)
	RegisterAllMsgType()
}

func RegisterAllMsgType() {
	//注册登录协议;
	RegisterMsgType(uint16(msg.MSG_ID_ELogin_Req), (*msg.CS_LoginReq)(nil))
	RegisterMsgType(uint16(msg.MSG_ID_ELogin_Ack), (*msg.SC_LoginResponse)(nil))
	//注册创建角色协议;
	RegisterMsgType(uint16(msg.MSG_ID_ECreateRole_Req), (*msg.CS_CreateRoleReq)(nil))
	RegisterMsgType(uint16(msg.MSG_ID_ECreateRole_Ack), (*msg.SC_CreateRoleAck)(nil))
}
