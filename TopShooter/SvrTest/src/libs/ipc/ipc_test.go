package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server EchoServer) Name() string {
	return "EchoServer"
}

func (server EchoServer) Handle(method, params string) *Response {
	var resp = Response{"200", "Echo:" + params}
	return &resp
}

var _ Server = &EchoServer{}

func TestIpc(t *testing.T) {
	server := NewIpcServer(EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	ret1, err2 := client1.Call("Hello", "This is from client1")
	if err2 != nil {
		t.Error("Client1.Call failed")
	}
	if ret1.Body != "Echo:This is from client1" {
		t.Error("Client1.Call respect return:", "This is from client1", "Act return:", ret1)
	}
	ret2, err2 := client2.Call("Hello", "This is from client2")
	if err2 != nil {
		t.Error("Client2.Call failed")
	}
	if ret2.Body != "Echo:This is from client2" {
		t.Error("Client2.Call respect return:", "This is from client2", "Act return:", ret2)
	}
}
