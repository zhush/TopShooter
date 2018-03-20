package app

import (
	"libs/log"
	"libs/yrpc"
)

func registerAllHandlers(service *yrpc.YService) {
	//register the crud method of databases
	yrpc.RegisterMsgHandler(service, "ReadTable", ReadTable)
	yrpc.RegisterMsgHandler(service, "AddTable", AddTable)
	yrpc.RegisterMsgHandler(service, "UpdateTable", UpdateTable)
	yrpc.RegisterMsgHandler(service, "RemoveTable", RemoveTable)
}

func ReadTable(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	log.Debug("Call ReadTable, reqJson:%s", reqJsonContent)
	isOk = true
	hasResponse = false

	return
}

func AddTable(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	isOk = true
	hasResponse = false
	return
}

func UpdateTable(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	isOk = true
	hasResponse = false
	return
}

func RemoveTable(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	isOk = true
	hasResponse = false
	return
}
