package net

type Agent interface {
	Run()
	OnClose()
}
