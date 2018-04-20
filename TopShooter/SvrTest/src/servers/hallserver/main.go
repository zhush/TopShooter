package main

import (
	log "libs/log"
	"libs/util"
	slog "log"
	"servers/hallserver/app"
)

func init() {
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "hallserver")
	log.Export(mlog)
}

func main() {
	app.App.Run()
	util.WaitAppCloseSignal()
}
