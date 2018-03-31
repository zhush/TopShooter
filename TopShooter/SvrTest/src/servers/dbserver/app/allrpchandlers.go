package app

import (
	"encoding/json"
	"libs/log"
	"libs/yrpc"
)

var readTableHandlerMap map[string]func(string, interface{}) map[string]string
var addTableHandlerMap map[string]func(string) (bool, int64)
var updateTableHandlerMap map[string]func(string, string) (bool, int64)
var removeTableHandlerMap map[string]func(string) (bool, int64)

func init() {
	readTableHandlerMap = make(map[string]func(string, interface{}) map[string]string)
	addTableHandlerMap = make(map[string]func(string) (bool, int64))
	updateTableHandlerMap = make(map[string]func(string, string) (bool, int64))
	removeTableHandlerMap = make(map[string]func(string) (bool, int64))
}

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

	var readTableParams map[string]interface{}
	err := json.Unmarshal([]byte(reqJsonContent), &readTableParams)
	if err != nil {
		log.Fatal("Unmarshal! %v", err)
	}
	tableName := readTableParams["TableName"].(string)
	keyName := readTableParams["Key"].(string)
	valName := readTableParams["Val"].(string)
	handler, isOk := readTableHandlerMap[tableName]
	if isOk == false {
		log.Fatal("Invalid Table Name:" + tableName)
	}

	log.Debug("ready to find handler,key:%v val:%v", keyName, valName)
	resultData := handler(keyName, valName)
	log.Debug("after excute handler %v", resultData)
	isOk = true
	hasResponse = true
	responseJsonBytes, err2 := json.Marshal(resultData)
	if err2 != nil {
		isOk = false
		hasResponse = false
		return
	}
	log.Debug("after excute responseJson:" + string(responseJsonBytes))
	responseJson = string(responseJsonBytes)
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

func RegisterReadTableHandler(tableName string, handler func(string, interface{}) map[string]string) {
	readTableHandlerMap[tableName] = handler
}

func RegisterAddTableHandler(tableName string, handler func(string) (bool, int64)) {
	addTableHandlerMap[tableName] = handler
}

func RegisterUpdateTableHandler(tableName string, handler func(string, string) (bool, int64)) {
	updateTableHandlerMap[tableName] = handler
}

func RegisterRemoveTableHandler(tableName string, handler func(string) (bool, int64)) {
	removeTableHandlerMap[tableName] = handler
}
