package app

import (
	"encoding/json"
	"libs/log"
	"libs/net"
	"libs/util"
	"libs/yrpc"
	"reflect"
	"servers/gateserver/config"

	"github.com/golang/protobuf/proto"
)

var App *Application

func init() {
	App = &Application{netServer: new(net.TCPServer), players: make(map[net.Conn]*ClientPlayer)}
}

type Application struct {
	netServer     *net.TCPServer //监听客户端的tcp服务器;
	managerServer *yrpc.YClient  //连接管理服务器
	loginServer   *yrpc.YClient  //连接的登录服务器
	serverName    string
	players       map[net.Conn]*ClientPlayer
}

func (app *Application) Run() {
	log.Debug("GateServer is starting...")
	//启动自己的rpc服务
	app.serverName = config.Conf["Name"].(string)

	log.Debug("GateServer is Running!!")
	//连接管理服务器
	app.tryConnectManager()
	//连接登录服务器
	app.tryConnectLoginServer()
	//开始监听客户端
	app.bindAndListenClient()
	//等待关闭信号
	util.WaitAppCloseSignal()
	//关闭清理服务器
	app.Close()
}

//服务器关闭
func (app *Application) Close() {
	log.Debug("GameServer is closed")
}

//连接管理服务器
func (app *Application) tryConnectManager() {
	app.managerServer = yrpc.NewYClient(config.Conf["ManagerAddr"].(string), "ManagerServer")
	<-app.managerServer.Connected
	log.Debug("Connected Manager Server Succeed")
}

func (app *Application) tryConnectLoginServer() {
	app.managerServer = yrpc.NewYClient(config.Conf["LoginServerAddr"].(string), "LoginServer")
	<-app.managerServer.Connected
	log.Debug("Connected Login Server Succeed")
}

func (app *Application) bindAndListenClient() {
	app.netServer.Addr = config.Conf["Addr"].(string)
	maxConnNum := config.Conf["MaxConnNum"].(float64)
	app.netServer.MaxConnNum = int(maxConnNum)
	app.netServer.PendingWriteNum = 10
	app.netServer.LenMsgLen = 2
	app.netServer.MaxMsgLen = 4096
	app.netServer.LittleEndian = true
	app.netServer.NewAgent = func(conn *net.TCPConn) net.Agent {
		return NewClientPlayer(conn)
	}
	app.netServer.Start()
	log.Debug("Start Listen: %v", app.netServer.Addr)
}

//处理登录的协议;
func (app *Application) HandleLoginMsg(client *ClientPlayer, msgId uint16, msgData []byte) {

	msgType, ok := MsgTypeMaps[msgId]
	if !ok {
		log.Error("Invalid msgId:%v, not register in gate..", msgId)
		return
	}
	m := reflect.New(msgType.Elem())

	err := proto.Unmarshal(msgData, m.Interface().(proto.Message))
	if err != nil {
		log.Error("Invalid LoginReq,error:%s, byte:%s", err.Error(), msgData)
		client.Close()
	}

	msgJson, err := json.Marshal(m)
	if err != nil {
		log.Error("call Json.Marshal failed,error:%s, byte:%s", err.Error(), msgData)
		client.Close()
		return
	}

	sendMsg := &yrpc.MsgS2SParam{MsgId: msgId, MsgBody: string(msgJson)}

	var reqJson []byte
	reqJson, err = json.Marshal(sendMsg)
	if err != nil {
		log.Error("call Json.Marshal failed,error:%s, byte:%s", err.Error(), msgData)
		client.Close()
		return
	}

	ret, hasResponse, respJson := app.loginServer.SendMsg("HandleClientMsg", string(reqJson))
	if ret == false {
		log.Error("call loginServer msg failed, msg:%s", msgJson)
		client.Close()
		return
	}
	//有结果返回
	if hasResponse == true {
		respMsg := &yrpc.MsgS2SParam{}
		err = json.Unmarshal([]byte(respJson), &respMsg)
		if err != nil {
			log.Error("call loginServer msg failed, msg:%s", msgJson)
			client.Close()
			return
		}
		recvMsgId := respMsg.MsgId
		respMsgType, ok2 := MsgTypeMaps[recvMsgId]
		if !ok2 {
			log.Error("Invalid recv msgId:%v from LoginServer, not register in gate..", recvMsgId)
			client.Close()
			return
		}

		realMsgType := reflect.New(respMsgType.Elem())
		err3 := proto.Unmarshal([]byte(respJson), realMsgType.Interface().(proto.Message))
		if err3 != nil {
			log.Error("call loginServer msg failed, msg:%s", msgJson)
			client.Close()
			return
		}
		client.SendMsg(recvMsgId, realMsgType.Interface().(proto.Message))
	}

}

//处理有游戏服的协议;
func (app *Application) HandleGameMsg(client *ClientPlayer, msgId uint16, msgData []byte) {

}
