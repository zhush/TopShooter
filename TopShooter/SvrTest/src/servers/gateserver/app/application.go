package app

import (
	"libs/log"
	"libs/net"
	"libs/util"
	"libs/yrpc"
	"servers/gateserver/agent"
	"servers/gateserver/config"
)

var App *Application

func init() {
	App = &Application{netServer: new(net.TCPServer), players: make(map[net.Conn]*agent.ClientPlayer)}
}

type Application struct {
	netServer     *net.TCPServer //监听客户端的tcp服务器;
	managerServer *yrpc.YClient  //连接管理服务器
	loginServer   *yrpc.YClient  //连接的登录服务器
	serverName    string
	players       map[net.Conn]*agent.ClientPlayer
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
	log.Debug("Connected Manager Server Succeed")
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
		a := &agent.ClientPlayer{conn}
		return a
	}
	app.netServer.Start()
}
