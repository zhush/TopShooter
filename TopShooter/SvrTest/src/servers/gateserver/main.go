package main

import (
	log "libs/log"
	slog "log"
	"servers/gateserver/app"
)

func init() {
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "gateserver")
	log.Export(mlog)
}

func main() {
	app.App.Run()
}
