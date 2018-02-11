package util

import (
	"encoding/json"
)

type ServerInfo struct {
	ServerName string
	ServerAddr string
	ServerId   int32
	ServerType int32
}

//获取转换后的json字符串
func (server *ServerInfo) GetJson() string {
	b, _ := json.Marshal(*server)
	return string(b)
}

//解析json到结构体中
func (server *ServerInfo) ParseJson(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), server)
}
