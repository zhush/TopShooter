package internal

import (
	"MyServer/conf"
	"MyServer/game"
	"MyServer/msg"
	"fmt"

	"github.com/name5566/leaf/gate"
)

type Module struct {
	*gate.Gate
}

func (self *Module) OnInit() {
	self.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSocketAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Server.TcpAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		AgentChanRPC:    game.ChanRpc,
	}

	switch conf.Encoding {
	case "json":
		self.Gate.Processor = msg.JsonProcessor
	default:
		fmt.Println("Invalid encoding")
	}

}

func (self *Module) OnDestroy() {

}
