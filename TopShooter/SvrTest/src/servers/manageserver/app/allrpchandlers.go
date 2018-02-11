package app

import (
	"fmt"
	"libs/log"
	"libs/util"
	"libs/yrpc"
)

func registerAllHandlers(service *yrpc.YService) {
	yrpc.RegisterMsgHandler(service, "BroadMessageToAllGame", BroadMessageToAllGame)
	yrpc.RegisterMsgHandler(service, "RegisterGameServer", RegisterGameServer)
}

func BroadMessageToAllGame(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {

	fmt.Println("Call BroadMessageToAllGame")
	fmt.Println("ServerCount:", ServerList.Len())
	for e := ServerList.Front(); e != nil; e = e.Next() {
		server := e.Value.(*GameServer)
		fmt.Println("Serv:", server.Name)
		server.sendMsgToServer("BroadMessageFromManager", reqJsonContent)
	}
	isOk = true
	hasResponse = false
	return
}

func RegisterGameServer(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	var serverInfo util.ServerInfo
	err := serverInfo.ParseJson(reqJsonContent)
	if err != nil {
		log.Error("RegisterGameServer, ParseJson Failed! reqJson:%s, err:%s", reqJsonContent, err.Error())
		return false, false, ""
	}
	log.Debug("RegisterGameServer ParseJson Succeed:%s", reqJsonContent)
	server := NewGameServer(&serverInfo)
	log.Debug("RegisterGameServer Succeed:%s", server.Name)
	return true, false, ""
}
