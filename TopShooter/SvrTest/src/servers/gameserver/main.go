package main

import (
	log "libs/log"
	"libs/util"
	slog "log"
	"servers/gameserver/app"
)

func init() {
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "loginmgr")
	log.Export(mlog)
}

func main() {
	app.App.Run()
	util.WaitAppCloseSignal()
}
