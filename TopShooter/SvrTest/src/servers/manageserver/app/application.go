package app

import (
	"libs/log"
	"libs/yrpc"
	"net/rpc"
	"servers/manageserver/config"
)

var App *Application

func init() {
	App = &Application{}
}

type Application struct {
	dbconn     *rpc.Client
	rpcService *yrpc.YService
}

func (app *Application) Run() {
	log.Debug("ManagerServer is starting...")
	app.startRpcService()
	log.Debug("Manager Server is Running!!")

}

func (app *Application) startRpcService() {
	log.Debug("ManagerServer is start rpc service...")
	app.rpcService = yrpc.NewYService(config.Conf["Addr"].(string))
	app.RegisterRpcMethod(app.rpcService)
	yrpc.ServiceStartRun(app.rpcService)
}

//注册供其他服调用的method.
func (app *Application) RegisterRpcMethod(service *yrpc.YService) {
	registerAllHandlers(service)
}

func (app *Application) Close() {
	log.Debug("DBServer is closed")
}
