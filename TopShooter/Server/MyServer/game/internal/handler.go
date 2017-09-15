package internal

import (
	"MyServer/msg"
	"reflect"

	"github.com/name5566/leaf/gate"
)

var userinfos map[string]msg.UserBaseInfo

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C2S_Login{}, handle_C2SLogin)
	userinfos = make(map[string]msg.UserBaseInfo)
	userinfos["admin"] = msg.UserBaseInfo{
		NickName: "管理员",
		Level:    100,
	}

}

//处理登录协议
func handle_C2SLogin(args []interface{}) {
	m := args[0].(*msg.C2S_Login)
	a := args[1].(gate.Agent)

	accname := m.AccountName
	pwd := m.Password
	if pwd != "admin" {
		a.WriteMsg(&msg.S2C_Login{
			Result:   -1,
			UserInfo: msg.UserBaseInfo{NickName: "", Level: 0},
		})
		return
	}
	a.WriteMsg(&msg.S2C_Login{
		Result:   0,
		UserInfo: userinfos[accname],
	})

}
