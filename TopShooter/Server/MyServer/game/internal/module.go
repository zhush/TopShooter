package internal

import (
	"MyServer/base"
	//"MyServer/msg"
	"fmt"
	"time"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRpc  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (self *Module) OnInit() {
	log.Debug("Server OnInit!!")
	self.Skeleton = skeleton
	fmt.Println("Game.Module onInit")

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			<-ticker.C
			//broadcastMessageAll(&msg.S2C_ChatNotify{NickName: "Random", Message: "Hello,World"})
			//fmt.Println("Call BroadMessage")
		}
	}()
}

func (self *Module) OnDestroy() {
	log.Debug("Server OnDestroy!!")
	fmt.Println("Game.Module OnDestroy")
}
