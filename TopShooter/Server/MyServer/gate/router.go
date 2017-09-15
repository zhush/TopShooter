package gate

import (
	"MyServer/game"
	"MyServer/msg"
)

func init() {
	msg.JsonProcessor.SetRouter(&msg.C2S_Login{}, game.ChanRpc)
}
