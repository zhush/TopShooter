package yrpc

import (
	"libs/log"
	"net/rpc"
	"sync"
	"time"
)

type YClient struct {
	conn       *rpc.Client
	isRunning  bool
	mutex      sync.Mutex
	remoteAddr string
	Connected  chan int
	serverName string
}

func NewYClient(addr string, name string) *YClient {
	client := &YClient{}
	client.remoteAddr = addr
	client.isRunning = false
	client.serverName = name
	client.init()
	client.Connected = make(chan int, 1)
	return client
}

func (self *YClient) init() {
	go func() {
		self.tryConnectServer()
	}()
}

func (self *YClient) onConnectedServer() {

}

func (self *YClient) tryConnectServer() {
	for {
		log.Debug("Enter tryConnectServer loop!!")
		var err error
		self.mutex.Lock()
		log.Debug("Start dialHttp:%s!!", self.remoteAddr)
		self.conn, err = rpc.DialHTTP("tcp", self.remoteAddr)
		self.mutex.Unlock()
		if err != nil {
			log.Debug("cannot connect %s(%s), error:%s, try after 1 second!!", self.serverName, self.remoteAddr, err.Error())
			time.Sleep(1 * time.Second)
		} else {
			log.Debug("connect succeed!!!")
			self.isRunning = true
			self.Connected <- 1
			break
		}
	}
	log.Debug("connect %s(%s) succeed!", self.serverName, self.remoteAddr)
	self.onConnectedServer()
}

func (self *YClient) IsRunning() bool {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	return self.isRunning
}

func (self *YClient) SendMsg(method string, reqJson string) (bool, bool, string) {
	if self.IsRunning() == false {
		log.Debug("The YClient is not running, addr:%s", self.remoteAddr)
		return false, false, ""
	}
	log.Debug("SendMsg, req:" + reqJson)
	arg := &ReqParam{}
	arg.MethodName = method
	arg.JsonContent = reqJson

	reply := &RespParam{}
	err := self.conn.Call("YService.RomoteCall", arg, reply)
	if err != nil {
		log.Error("YService.RomoteCall rpc error:", err.Error())
		self.conn.Close()
		self.conn = nil
		self.isRunning = false
		log.Debug("try connected %s again", self.serverName)
		self.tryConnectServer()
		<-self.Connected
		log.Debug("connected %s succeed!", self.serverName)

		err = self.conn.Call("YService.RomoteCall", arg, reply)
		if err != nil {
			return false, false, ""
		}
	}
	log.Debug("SendMsg, resp:%v", reply)
	return reply.Result, reply.HasResponse, reply.JsonContent
}
