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
	netServer *net.TCPServer
	players   map[net.Conn]*ClientPlayer
	dbServer  *yrpc.YClient
	dbmutex   sync.Mutex
	dbIsReady bool
}

func (app *Application) Run() {
	app.netServer.Addr = config.Conf["Addr"].(string)
	maxConnNum := config.Conf["MaxConnNum"].(float64)
	app.netServer.MaxConnNum = int(maxConnNum)
	app.netServer.PendingWriteNum = 10
	app.netServer.LenMsgLen = 2
	app.netServer.MaxMsgLen = 4096
	app.netServer.LittleEndian = true
	app.netServer.NewAgent = func(conn *net.TCPConn) net.Agent {
		a := &ClientPlayer{conn}
		return a
	}
	app.netServer.Start()
	log.Debug("LoginServer Starting")
	log.Debug("Listen addr:%s", config.Conf["Addr"].(string))
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
