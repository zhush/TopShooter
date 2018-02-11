package app

import (
	"libs/log"
	"libs/util"
	"libs/yrpc"
	"net/http"
	"net/rpc"
	"servers/gameserver/config"
	"sync"
	"time"
)

var App *Application

func init() {
	App = &Application{}
}

type Application struct {
	rpcService   *yrpc.YService
	managerMutex sync.Mutex
	managerConn  *rpc.Client
	serverName   string
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

//服务器关闭
func (app *Application) Close() {
	log.Debug("GameServer is closed")
}

//连接管理服务器
func (app *Application) tryConnectManager() {
	log.Debug("Try to connect Manager")
	go func() {
		manageAddr := config.Conf["ManagerAddr"].(string)
		for {
			var err error
			app.managerMutex.Lock()
			app.managerConn, err = rpc.DialHTTP("tcp", manageAddr)
			app.managerMutex.Unlock()
			if err != nil {
				log.Debug("cannot connect Manager(%s), error:%s, try after 1 second!!", manageAddr, err.Error())
				time.Sleep(1 * time.Second)
			} else {
				log.Debug("connect manager(%s), succeed!!", manageAddr)
				app.RegisterToManager()
				break
			}
		}
	}()
}

//向rpc服务器发送消息
func (app *Application) sendMsgToServer(server *rpc.Client, method string, reqJson string) (bool, bool, string) {
	log.Debug("Enter  sendMsgToServer")
	arg := &yrpc.ReqParam{}
	arg.MethodName = method
	arg.JsonContent = reqJson

	reply := &yrpc.RespParam{}
	log.Debug("Ready Call Manager RemoteCall")
	err := server.Call("YService.RomoteCall", arg, reply)
	if err != nil {
		log.Error("YService.RomoteCall rpc error:", err.Error())
		return false, false, ""
	}
	return reply.Result, reply.HasResponse, reply.JsonContent
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
	isOk, _, _ := app.sendMsgToServer(app.managerConn, "RegisterGameServer", serverInfo.GetJson())
	if isOk == false {
		log.Error("Register GameServer(%s) to ManagerServer(%s) Failed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])
		return
	}
	log.Debug("Register GameServer(%s) to ManagerServer(%s) Succeed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])

	isOk1, _, _ := app.sendMsgToServer(app.managerConn, "BroadMessageToAllGame", "Hello, World, This is from GameServer:"+app.serverName)
	if isOk1 == false {
		log.Error("BroadToManager GameServer(%s) to ManagerServer(%s) Failed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])
		return
	}
	log.Debug("BroadToManager GameServer(%s) to ManagerServer(%s) Succeed!!", config.Conf["Addr"], config.Conf["ManagerAddr"])
}
