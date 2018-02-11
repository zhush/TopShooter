package net

import (
	"fmt"
	"testing"
	"time"
)

func TestClientServer(t *testing.T) {
	tcpServer := new(TCPServer)
	tcpServer.Addr = "127.0.0.1:15051"
	tcpServer.MaxConnNum = 20
	tcpServer.PendingWriteNum = 10
	tcpServer.LenMsgLen = 2
	tcpServer.MaxMsgLen = 4096
	tcpServer.LittleEndian = true
	tcpServer.NewAgent = func(conn *TCPConn) Agent {
		a := &agent{conn: conn}
		return a
	}
	tcpServer.Start()
	fmt.Println("tcpserver is start")
	go func() {
		tcpClient := new(TCPClient)
		tcpClient.Addr = "127.0.0.1:15051"
		tcpClient.ConnNum = 1
		tcpClient.ConnectInterval = 5 * time.Second
		tcpClient.AutoReconnect = true
		tcpClient.NewAgent = func(conn *TCPConn) Agent {
			a := &cagent{conn: conn}
			t.Error("error")
			return a
		}

		tcpClient.init()
		tcpClient.connect()
	}()

}

type cagent struct {
	conn *TCPConn
}

func (a *cagent) Run() {
	fmt.Println("client is connected")
	a.conn.WriteMsg([]byte("This is come from client"))
}

func (a *cagent) OnClose() {
}

type agent struct {
	conn *TCPConn
}

func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("agent readmsg error:", err)
		}
		a.RecvMsg(string(data))
	}
}

func (a *agent) OnClose() {

}

func (a *agent) WriteMsg(msg string) {
	a.conn.WriteMsg([]byte(msg))
}

func (a *agent) RecvMsg(msg string) {
	fmt.Println("Recv Msg:", msg)
}
