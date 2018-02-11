package app

import (
	"libs/log"
	"libs/yrpc"
	"net/http"
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
	app.rpcService = yrpc.NewYService()
	app.RegisterRpcMethod(app.rpcService)
	rpc.Register(app.rpcService)
	rpc.HandleHTTP()
	go func() {
		log.Debug("Listening:%s", config.Conf["Addr"])
		err := http.ListenAndServe(config.Conf["Addr"].(string), nil)
		if err != nil {
			log.Fatal("ListenAndServer Failed:%s", err.Error())
		}
	}()
}

//注册供其他服调用的method.
func (app *Application) RegisterRpcMethod(service *yrpc.YService) {
	registerAllHandlers(service)
}

func (app *Application) Close() {
	log.Debug("DBServer is closed")
}
