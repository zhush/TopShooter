package app

import (
	"fmt"
	"libs/log"
	"libs/yrpc"
	"servers/dbserver/config"

	"github.com/go-redis/redis"
)

var App *Application

func init() {
	fmt.Println("application init")
	App = &Application{}
}

type Application struct {
	rpcService *yrpc.YService
}

func (app *Application) Run() {
	log.Debug("Database Server is starting...")
	app.startRpcService()
	log.Debug("Database Server is Running!!")

}

func redisOptions() *redis.Options {
	return &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
}

func sqlOptions() string {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.Conf["DatabaseUser"].(string),
		config.Conf["DatabasePwd"].(string),
		config.Conf["DatabaseAddr"].(string),
		config.Conf["Database"].(string))
	return dns
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
	registerAllHandlers(service)
}

func (app *Application) Close() {
	log.Debug("DBServer is closed")
}
