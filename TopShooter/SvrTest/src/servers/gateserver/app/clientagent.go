package app

import (
	"bytes"
	"encoding/binary"
	"libs/log"
	"libs/net"

	"github.com/golang/protobuf/proto"
)

type ClientPlayer struct {
	Conn    net.Conn
	State   AgentStatus
	isClose bool
}

func NewClientPlayer(conn net.Conn) *ClientPlayer {
	client := &ClientPlayer{Conn: conn, State: StatusLogin, isClose: false}
	return client
}

func (self *ClientPlayer) Run() {
	log.Debug("A new clientPlayer is created!!")
	for {
		if self.isClose == true {
			break
		}
		data, err := self.Conn.ReadMsg()
		if err != nil {
			log.Debug("read message Error: %v", err)
			break
		}
		msgLenBuf := data[:2]
		msgData := data[2:]
		buf := bytes.NewReader(msgLenBuf)
		var msgId uint16
		binary.Read(buf, binary.LittleEndian, &msgId)
		log.Debug("recv MsgId:%d", msgId)
		self.HandleMsg(msgId, msgData)
	}
}

func (self *ClientPlayer) OnClose() {
	self.isClose = true
}

func (self *ClientPlayer) Close() {
	self.Conn.Close()
	self.OnClose()
}

func (self *ClientPlayer) SetState(state AgentStatus) {
	self.State = state
}

//处理客户端发送的消息
func (self *ClientPlayer) HandleMsg(msgId uint16, msgData []byte) {
	if self.State == StatusLogin {
		App.HandleLoginMsg(self, msgId, msgData)
	} else if self.State == StatusGaming {
		App.HandleLoginMsg(self, msgId, msgData)
	}
}

//发送消息到客户端
func (self *ClientPlayer) SendMsg(msgId uint16, msg proto.Message) {
	data, err := proto.Marshal(msg)
	headBuf := new(bytes.Buffer)
	binary.Write(headBuf, binary.LittleEndian, msgId)

	log.Debug("headBuf len:%d", headBuf.Len())
	log.Debug("msgBuf len:%d", len(data))
	buf := headBuf.Bytes()
	buf = append(buf, data...)
	log.Debug("sendClient len:%d", len(buf))
	err = self.Conn.WriteMsg(buf)
	if err != nil {
		log.Fatal("ClientPlayer SendMsg Failed!" + err.Error())
	}
}
