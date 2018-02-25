package main

import (
	"fmt"
	log "libs/log"
	slog "log"
	"os"
	"os/signal"
	"servers/dbserver/app"
)

func init() {
	fmt.Println("Call dbserver init function")
	mlog, _ := log.New("debug", "./ylog", slog.Ltime|slog.Ldate, "dbmgr")
	log.Export(mlog)
}

func main() {
	fmt.Println("Start Call app.App.Run")
	app.App.Run()
	//close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	app.App.Close()
}
