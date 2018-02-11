package app

import (
	"container/list"
	"errors"
	"libs/log"
	"libs/util"
	"libs/yrpc"
	"net/rpc"
	"sync"
	"time"
)

type GameServer struct {
	Proxy        *rpc.Client
	Name         string
	isReady      bool
	connectMutex sync.Mutex
	ServerAddr   string
	ServerId     int32
	ServerType   int32
}

var ServerList *list.List

func init() {
	ServerList = list.New()
}

func (server *GameServer) Init() {
	go func() {
		for {
			server.connectMutex.Lock()
			defer server.connectMutex.Unlock()
			var err error
			server.Proxy, err = rpc.DialHTTP("tcp", server.ServerAddr)
			if err != nil {
				log.Debug("cannot connect GameServer(%s), error:%s, try after 1 second!!", server.ServerAddr, err.Error())
				time.Sleep(1 * time.Second)
				continue
			}
			server.isReady = true
			break
		}
	}()
}

//判断是否准备好
func (server *GameServer) IsReady() bool {
	server.connectMutex.Lock()
	defer server.connectMutex.Unlock()
	return server.isReady
}

//向rpc服务器发送消息
func (server *GameServer) sendMsgToServer(method string, reqJson string) (bool, bool, string) {

	if server.IsReady() == false {
		log.Debug("Call sendMsgToServer Failed, server.IsReady == false")
		return false, false, ""
	}

	log.Debug("Enter sendMsgToServer")
	arg := &yrpc.ReqParam{}
	arg.MethodName = method
	arg.JsonContent = reqJson

	reply := &yrpc.RespParam{}
	log.Debug("Ready Call Manager RemoteCall")
	err := server.Proxy.Call("YService.RomoteCall", arg, reply)
	if err != nil {
		log.Error("YService.RomoteCall rpc error:", err.Error())
		return false, false, ""
	}
	return reply.Result, reply.HasResponse, reply.JsonContent
}

//开启一个新的服务器
func NewGameServer(serverInfo *util.ServerInfo) *GameServer {
	server := &GameServer{}
	server.ServerAddr = serverInfo.ServerAddr
	server.isReady = false
	server.Name = serverInfo.ServerName
	server.ServerId = serverInfo.ServerId
	server.ServerType = serverInfo.ServerType
	server.Init()
	ServerList.PushBack(server)
	return server
}

//删除一个服务器
func DelGameServer(server *GameServer) error {
	for e := ServerList.Front(); e != nil; e = e.Next() {
		if e.Value == server {
			ServerList.Remove(e)
			return nil
		}
	}
	log.Debug("Cannot Del GameServer:%s", server.Name)
	return errors.New("Cannot Del GameServer" + server.Name)
}
