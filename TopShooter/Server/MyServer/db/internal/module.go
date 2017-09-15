package internal

import (
	"MyServer/base"
	"fmt"

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
}

func (self *Module) OnDestroy() {
	log.Debug("Server OnDestroy!!")
	fmt.Println("Game.Module OnDestroy")
}
