package app

import (
	"bytes"
	"encoding/binary"
	"libs/log"
	"libs/net"

	"github.com/golang/protobuf/proto"
)

type ClientPlayer struct {
	conn net.Conn
}

func (player *ClientPlayer) Run() {
	log.Debug("A new clientPlayer is created!!")
	for {
		data, err := player.conn.ReadMsg()
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
		player.HandleMsg(msgId, msgData)
	}
}

func (player *ClientPlayer) OnClose() {

}

func (player *ClientPlayer) Close() {
	player.conn.Close()
}

//处理客户端发送的消息
func (player *ClientPlayer) HandleMsg(msgId uint16, msgData []byte) {
	if handler, ok := handlerMsgMaps[msgId]; ok {
		handler(player, msgData)
	} else {
		log.Fatal("ClientPlayer HandleMsg, Invalid msgId:%d", msgId)
	}

}

//发送消息到客户端
func (player *ClientPlayer) SendMsg(msgId uint16, msg proto.Message) {
	data, err := proto.Marshal(msg)
	headBuf := new(bytes.Buffer)
	binary.Write(headBuf, binary.LittleEndian, msgId)

	log.Debug("headBuf len:%d", headBuf.Len())
	log.Debug("msgBuf len:%d", len(data))
	buf := headBuf.Bytes()
	buf = append(buf, data...)
	log.Debug("sendClient len:%d", len(buf))
	err = player.conn.WriteMsg(buf)
	if err != nil {
		log.Fatal("ClientPlayer SendMsg Failed!" + err.Error())
	}
}

var handlerMsgMaps map[uint16]func(*ClientPlayer, []byte)

func RegisterMsgHandler(msgId uint16, handler func(*ClientPlayer, []byte)) {
	handlerMsgMaps[msgId] = handler
}

func init() {
	handlerMsgMaps = make(map[uint16]func(*ClientPlayer, []byte))
}
