package app

import (
	"fmt"
	"libs/yrpc"
)

//注册供其他服务器调用的数据库操作接口
func registerAllHandlers(service *yrpc.YService) {
	yrpc.RegisterMsgHandler(service, "AddRecord", AddRecord)
	yrpc.RegisterMsgHandler(service, "DelRecord", DelRecord)
	yrpc.RegisterMsgHandler(service, "UpdateRecord", UpdateRecord)
	yrpc.RegisterMsgHandler(service, "SelectRecord", SelectRecord)
}

//增加记录
func AddRecord(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	isOk = true
	hasResponse = false
	return
}

func DelRecord(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	fmt.Println("Recv Msg From Manager:", reqJsonContent)
	isOk = true
	hasResponse = false
	return
}

func UpdateRecord(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	fmt.Println("Recv Msg From Manager:", reqJsonContent)
	isOk = true
	hasResponse = false
	return
}

func SelectRecord(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	fmt.Println("Recv Msg From Manager:", reqJsonContent)
	isOk = true
	hasResponse = false
	return
}
