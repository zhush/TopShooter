package app

import (
	"msg"
)

//这个文件是msg的注册解析类，主要是protobuf中消息id和结构对应关系.

var MsgTypeMaps map[uint16]string

func RegisterMsgType(msgId uint16, msgName string) {
	MsgTypeMaps[msgId] = msgName
}

func init() {
	MsgTypeMaps = make(map[uint16]string)
	RegisterAllMsgType()
}

func RegisterAllMsgType() {
	//注册登录协议;
	RegisterMsgType(uint16(msg.MSG_ID_ELogin_Req), "msg.CS_LoginReq")
	RegisterMsgType(uint16(msg.MSG_ID_ELogin_Ack), "msg.SC_LoginResponse")
	//注册创建角色协议;
	RegisterMsgType(uint16(msg.MSG_ID_ECreateRole_Req), "msg.CS_CreateRoleReq")
	RegisterMsgType(uint16(msg.MSG_ID_ECreateRole_Ack), "msg.SC_CreateRoleAck")
}
