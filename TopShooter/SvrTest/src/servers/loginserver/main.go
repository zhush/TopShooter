package main

import (
	log "libs/log"
	slog "log"
	"os"
	"os/signal"
	"servers/loginserver/app"
)

func init() {
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "loginmgr")
	log.Export(mlog)
}

func main() {
	app.App.Run()
	//close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	app.App.Close()
}
