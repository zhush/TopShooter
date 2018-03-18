package yrpc

import (
	"errors"
	"fmt"
	"libs/log"
	"net/http"
	"net/rpc"
	"sync"
)

type YService struct {
	msgHandlers map[string]func(string) (bool, bool, string)
	isRunning   bool
	mutex       sync.Mutex
	bindAddr    string
}

func NewYService(addr string) *YService {
	service := &YService{}
	service.bindAddr = addr
	service.msgHandlers = make(map[string]func(string) (bool, bool, string))
	service.isRunning = false
	return service
}

func RegisterMsgHandler(service *YService, cmd string, handler func(string) (bool, bool, string)) {
	service.msgHandlers[cmd] = handler
}

func ServiceIsRunning(service *YService) bool {
	return service.isRunning
}

func ServiceStartRun(service *YService) {
	rpc.Register(service)
	rpc.HandleHTTP()
	go func() {
		log.Debug("Listening:%s", service.bindAddr)
		err := http.ListenAndServe(service.bindAddr, nil)
		if err != nil {
			log.Fatal("ListenAndServer Failed:%s", err.Error())
		}
		service.isRunning = true
	}()
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
