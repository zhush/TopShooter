package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "methods"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("unmarshal json error:", err.Error())
				break
			}
			resp := server.Handle(req.Method, req.Params)
			b, err1 := json.Marshal(resp)
			if err1 != nil {
				fmt.Println("json marshal error:", resp)
			}
			c <- string(b)
		}
		fmt.Println("Session closed")
	}(session)
	fmt.Println("A new Session has created")
	return session
}
