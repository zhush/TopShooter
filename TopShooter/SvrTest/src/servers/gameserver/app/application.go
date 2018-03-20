package app

import (
	"libs/log"
	"libs/util"
	"libs/yrpc"
	"servers/gameserver/config"
)

var App *Application

func init() {
	App = &Application{}
}

type Application struct {
	rpcService    *yrpc.YService
	managerServer *yrpc.YClient
	serverName    string
}

func (app *Application) Run() {
	log.Debug("GameServer is starting...")
	//启动自己的rpc服务
	app.serverName = config.Conf["Name"].(string)
	app.startRpcService()
	log.Debug("GameServer is Running!!")
	//连接管理器
	app.tryConnectManager()

}

//启动自己的rpc服务
func (app *Application) startRpcService() {
	log.Debug("GameServer is start rpc service...")
	app.rpcService = yrpc.NewYService(config.Conf["Addr"].(string))
	app.RegisterRpcMethod(app.rpcService)
	yrpc.ServiceStartRun(app.rpcService)
}

//注册供其他服调用的method.
func (app *Application) RegisterRpcMethod(service *yrpc.YService) {
	registerAllHandlers(service)
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
	log.Debug("Start Register To Manager Server")
	app.RegisterToManager()
}

//向管理服务器进行注册
func (app *Application) RegisterToManager() {

	log.Debug("Enter RegisterToManager")
	serverInfo := &util.ServerInfo{
		config.Conf["Name"].(string),
		config.Conf["Addr"].(string),
		int32(config.Conf["ServerId"].(float64)),
		int32(config.Conf["ServerType"].(float64))}
	log.Debug("Construct serverInfo")
	isOk, _, _ := app.managerServer.SendMsg("RegisterGameServer", serverInfo.GetJson())
	if isOk == false {
		log.Error("Register GameServer(%s) to ManagerServer(%s) Failed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])
		return
	}
	log.Debug("Register GameServer(%s) to ManagerServer(%s) Succeed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])

	isOk1, _, _ := app.managerServer.SendMsg("BroadMessageToAllGame", "Hello, World, This is from GameServer:"+app.serverName)
	if isOk1 == false {
		log.Error("BroadToManager GameServer(%s) to ManagerServer(%s) Failed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])
		return
	}
	log.Debug("BroadToManager GameServer(%s) to ManagerServer(%s) Succeed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])

}
