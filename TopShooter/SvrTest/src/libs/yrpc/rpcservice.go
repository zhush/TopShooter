package yrpc

import (
	"errors"
	"fmt"
	"libs/log"
)

type YService struct {
	msgHandlers map[string]func(string) (bool, bool, string)
}

func NewYService() *YService {
	service := &YService{}
	service.msgHandlers = make(map[string]func(string) (bool, bool, string))
	return service
}

func RegisterMsgHandler(service *YService, cmd string, handler func(string) (bool, bool, string)) {
	service.msgHandlers[cmd] = handler
}

func (service *YService) RomoteCall(param *ReqParam, result *RespParam) error {
	cmd := param.MethodName
	reqContent := param.JsonContent
	handler, isOk := service.msgHandlers[cmd]
	if isOk == false {
		(*result).Result = false
		(*result).JsonContent = ""

		errmsg := fmt.Sprintf("Unregister Cmd (%s) failed! ", cmd)
		(*result).JsonContent = errmsg
		log.Error(errmsg)
		return errors.New(errmsg)
	}

	isHandleOk, hasResponse, respContent := handler(reqContent)
	(*result).Result = isHandleOk
	(*result).HasResponse = hasResponse
	(*result).JsonContent = respContent
	return nil
}
