package app

import (
	"fmt"
	"libs/yrpc"
)

func init() {
}

func registerAllHandlers(service *yrpc.YService) {
	yrpc.RegisterMsgHandler(service, "NoticeAllGameStates", NoticeAllGameStates)
	yrpc.RegisterMsgHandler(service, "BroadMessageFromManager", BroadMessageFromManager)
}

func NoticeAllGameStates(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	isOk = true
	hasResponse = false
	return
}

func BroadMessageFromManager(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	fmt.Println("Recv Msg From Manager:", reqJsonContent)
	isOk = true
	hasResponse = false
	return
}
