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
}

func NewYClient(addr string) *YClient {
	client := &YClient{}
	client.remoteAddr = addr
	client.isRunning = false
	client.init()
	client.Connected = make(chan int, 1)
	return client
}

func (self *YClient) init() {
	go func() {
		for {
			var err error
			self.mutex.Lock()
			self.conn, err = rpc.DialHTTP("tcp", self.remoteAddr)
			if err != nil {
				log.Debug("cannot connect Manager(%s), error:%s, try after 1 second!!", self.remoteAddr, err.Error())
				self.mutex.Unlock()
				time.Sleep(1 * time.Second)
			} else {
				self.isRunning = true
				self.mutex.Unlock()
				self.Connected <- 1
				break
			}
		}
	}()
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
	log.Debug("Enter  sendToServer")
	arg := &ReqParam{}
	arg.MethodName = method
	arg.JsonContent = reqJson

	reply := &RespParam{}
	log.Debug("Ready Call RemoteCall")
	err := self.conn.Call("YService.RomoteCall", arg, reply)
	if err != nil {
		log.Error("YService.RomoteCall rpc error:", err.Error())
		return false, false, ""
	}
	return reply.Result, reply.HasResponse, reply.JsonContent
}
