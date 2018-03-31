package app

import (
	"encoding/json"
	"fmt"
	"libs/log"
	"msg"
	"servers/dbserver/common"

	"github.com/golang/protobuf/proto"
)

func init() {
	RegisterMsgHandler(uint16(msg.MSG_ID_ELogin_Req), handle_CS_LoginReq)
}

func handle_CS_LoginReq(player *ClientPlayer, msgData []byte) {

	loginReq := &msg.CS_LoginReq{}
	err := proto.Unmarshal(msgData, loginReq)
	if err != nil {
		log.Error("Invalid LoginReq,error:%s, byte:%s", err.Error(), msgData)
		player.Close()
	}

	response := &msg.SC_LoginResponse{}

	log.Debug("Recv LoginReq, UserName:%s Password:%s Platform:%d", *(loginReq.AccName), *(loginReq.AccPassword), *(loginReq.PlatForm))
	if App.DBIsReady() == false {
		log.Debug("DB is not ready. login failed")
		errCode := msg.ELoginResult_ServerClosed
		response.LoginResult = &errCode
		player.SendMsg(uint16(msg.MSG_ID_ELogin_Ack), response)
		player.Close()
	}

	args := &Params.ReadParam{}
	args.TableName = "t_account"
	args.Keys = []string{"password"}
	args.Conditions = []Params.KeyValue{Params.KeyValue{"accountName", *(loginReq.AccName)}}

	//reply := &Params.ReadResult{}

	reqJson := fmt.Sprintf(`{"TableName":"%s",
							 "Key":"%s",
							 "Val":"%s"}`, "t_account", "accountName", *(loginReq.AccName))

	ret, _, respJson := App.dbServer.SendMsg("ReadTable", reqJson)

	log.Debug("ret:%v respJson:%v", ret, respJson)

	if ret == false {
		log.Error("DBService.Read rpc error:")
		errCode := msg.ELoginResult_InvalidAccOrPwd
		response.LoginResult = &errCode
		return
	}

	log.Debug("Recv LoginResult is:%s", respJson)

	var loginResult map[string]interface{}
	json.Unmarshal([]byte(respJson), &loginResult)
	log.Debug("loginResult:%v", loginResult)

	password := loginResult["password"]

	if password.(string) != *(loginReq.AccPassword) {
		log.Debug("Server Password is:%v client Password is:%v", password, *(loginReq.AccPassword))
		errCode := msg.ELoginResult_InvalidAccOrPwd
		response.LoginResult = &errCode
		player.SendMsg(uint16(msg.MSG_ID_ELogin_Ack), response)
		player.Close()
	}

	retCode := msg.ELoginResult_Succeed
	response.LoginResult = &retCode
	//填写玩家基本信息
	player.SendMsg(uint16(msg.MSG_ID_ELogin_Ack), response)

}
