package util

type ServerType int

const (
	SType_LoginServer   ServerType = iota + 1 // value = 1
	SType_GameServer                          //value = 2
	SType_ManagerServer                       //value = 3
	SType_GateServer                          //value = 4
	SType_DBServer                            //value = 5
)
