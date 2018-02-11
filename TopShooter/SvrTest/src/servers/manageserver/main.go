package main

import (
	log "libs/log"
	"libs/util"
	slog "log"
	"servers/manageserver/app"
)

func init() {
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "managemgr")
	log.Export(mlog)
}

func main() {
	app.App.Run()
	//等待关闭信号
	util.WaitAppCloseSignal()
	app.App.Close()
}
