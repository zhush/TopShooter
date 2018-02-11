package app

import (
	"libs/log"
	"libs/net"
	"libs/util"
	"net/rpc"
	"servers/loginserver/config"
	"sync"
	"time"
)

var App *Application

func init() {
	App = &Application{netServer: new(net.TCPServer), players: make(map[net.Conn]*ClientPlayer)}
}

type Application struct {
	netServer *net.TCPServer
	players   map[net.Conn]*ClientPlayer
	dbconn    *rpc.Client
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
	go func() {
		dbAddr := config.Conf["DBAddr"].(string)
		for {
			var err error
			app.dbmutex.Lock()
			app.dbconn, err = rpc.DialHTTP("tcp", dbAddr)
			app.dbmutex.Unlock()
			if err != nil {
				log.Debug("cannot connect db(%s), error:%s, try after 1 second!!", dbAddr, err.Error())
				time.Sleep(1 * time.Second)
			} else {
				log.Debug("connect db(%s), succeed!!", dbAddr)
				break
			}
		}
	}()
}

func (app *Application) DBIsReady() bool {
	if app.dbIsReady == true {
		return true
	}
	isReady := false
	app.dbmutex.Lock()
	isReady = app.dbconn != nil
	app.dbmutex.Unlock()
	if isReady == true {
		app.dbIsReady = true
	}
	return app.dbIsReady
}

//关闭玩家
func (app *Application) ClosePlayer(player *ClientPlayer) {
	player.Close()
	app.players[player.conn] = nil
}

func (app *Application) Close() {
	log.Debug("LoginServer Closing")
}
