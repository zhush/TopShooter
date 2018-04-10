package app

import (
	"encoding/json"
	"fmt"
	"libs/log"
	"libs/yrpc"
)

func registerAllHandlers(service *yrpc.YService) {
	yrpc.RegisterMsgHandler(service, "HandleClientMsg", HandleClientMsg)
}

func HandleClientMsg(reqJsonContent string) (isOk bool, hasResponse bool, responseJson string) {
	//解析协议;
	reqMsg := &yrpc.MsgS2SParam{}
	err := json.Unmarshal([]byte(reqJsonContent), &reqMsg)
	if err != nil {
		isOk = false
		hasResponse = false
		log.Error("excute HandleClientMsg Failed!, inalid reqJsonContent:%v", HandleClientMsg)
		return
	}

	//找对应的处理逻辑函数
	handler, ok := handlerClientMsgMaps[reqMsg.MsgId]
	if !ok {
		isOk = false
		hasResponse = false
		log.Error("excute HandleClientMsg Failed!, inalid reqJsonContent:%v", HandleClientMsg)
		return
	}

	//处理相应的协议逻辑;
	respMsgId, respJson := handler(reqMsg.MsgBody)

	//发送给远程调用的服务器;
	respMsg := &yrpc.MsgS2SParam{MsgId: respMsgId, MsgBody: respJson}
	isOk = true
	hasResponse = true
	respJsonBytes, err2 := json.Marshal(respMsg)
	if err2 != nil {
		isOk = false
		hasResponse = false
		log.Error("excute HandleClientMsg Failed!, inalid reqJsonContent:%v, err:%v", HandleClientMsg, err2.Error())
		return
	}
	responseJson = string(respJsonBytes)
	fmt.Println("reqJson:", reqJsonContent, " respJson:", responseJson)
	return
}
