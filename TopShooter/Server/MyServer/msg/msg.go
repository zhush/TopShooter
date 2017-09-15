package msg

import (
	"github.com/name5566/leaf/network/json"
)

var (
	JsonProcessor = json.NewProcessor()
)

func init() {
	JsonProcessor.Register(&C2S_Login{})
	JsonProcessor.Register(&S2C_Login{})
	JsonProcessor.Register(&S2C_ChatNotify{})
}

type C2S_Login struct {
	AccountName string //用户账户
	Password    string //用户密码
}

type UserBaseInfo struct {
	NickName string //玩家昵称
	Level    int    //玩家等级
}

type S2C_Login struct {
	Result   int
	UserInfo UserBaseInfo
}

type S2C_ChatNotify struct {
	NickName string //玩家昵称
	Message  string //聊天内容
}
