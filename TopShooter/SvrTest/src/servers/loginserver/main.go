package main

import (
	log "libs/log"
	slog "log"
	"servers/loginserver/app"
)

func init() {
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "loginmgr")
	log.Export(mlog)
}

func main() {
	app.App.Run()
}
