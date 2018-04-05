package app

import (
	"libs/log"
	"libs/net"
	"libs/util"
	"libs/yrpc"
	"servers/loginserver/config"
	"sync"
)

var App *Application

func init() {
	App = &Application{netServer: new(net.TCPServer), players: make(map[net.Conn]*ClientPlayer)}
}

type Application struct {
	rpcService *yrpc.YService
	netServer  *net.TCPServer
	players    map[net.Conn]*ClientPlayer
	dbServer   *yrpc.YClient
	dbmutex    sync.Mutex
	dbIsReady  bool
}

func (app *Application) Run() {
	log.Debug("LoginServer Starting")
	app.startRpcService()
	app.dbIsReady = false
	app.tryConnectDB()
	util.WaitAppCloseSignal()
	app.Close()
}

//连接db
func (app *Application) tryConnectDB() {
	log.Debug("Start Ready To connect DBServer")
	app.dbServer = yrpc.NewYClient(config.Conf["DBAddr"].(string), "DBManager")
	<-app.dbServer.Connected
	log.Debug("Succeed connect DBServer")
}

func (app *Application) DBIsReady() bool {
	return app.dbServer.IsRunning()
}

//关闭玩家
func (app *Application) ClosePlayer(player *ClientPlayer) {
	player.Close()
	app.players[player.conn] = nil
}

func (app *Application) Close() {
	log.Debug("LoginServer Closing")
}

func (app *Application) startRpcService() {
	log.Debug("Database Server is start rpc service...")
	log.Debug("ManagerServer is start rpc service...")
	app.rpcService = yrpc.NewYService(config.Conf["Addr"].(string))
	app.RegisterRpcMethod(app.rpcService)
	yrpc.ServiceStartRun(app.rpcService)
}

//注册供其他服调用的method.
func (app *Application) RegisterRpcMethod(service *yrpc.YService) {

}
