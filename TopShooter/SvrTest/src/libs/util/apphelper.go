package util

import (
	"os"
	"os/signal"
)

func WaitAppCloseSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
