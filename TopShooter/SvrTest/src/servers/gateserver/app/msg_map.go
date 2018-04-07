package app

//这个文件是msg的注册解析类，主要是protobuf中消息id和结构对应关系.

var handlerMsgMaps map[uint16]

func RegisterMsgHandler(msgId uint16, handler func(*ClientPlayer, []byte)) {
	handlerMsgMaps[msgId] = handler
}

func init() {
	handlerMsgMaps = make(map[uint16]func(*ClientPlayer, []byte))
}
