package app

import (
	"encoding/json"
	"fmt"
	"libs/log"
	"msg"
	"servers/dbserver/common"
)

var handlerClientMsgMaps map[uint16]func(string) (uint16, string)

func init() {
	handlerClientMsgMaps = make(map[uint16]func(string) (uint16, string))
	RegisterAllClientMsgHandler()
}

func RegisterAllClientMsgHandler() {
	RegisterClientMsgHandler(uint16(msg.MSG_ID_ELogin_Req), handle_CS_LoginReq)
}

func RegisterClientMsgHandler(msgId uint16, handler func(string) (uint16, string)) {
	handlerClientMsgMaps[msgId] = handler
}

//处理玩家登录
func handle_CS_LoginReq(reqJson string) (respMsgId uint16, respJson string) {

	response := &msg.SC_LoginResponse{}
	defer func() {
		respJsonBytes, _ := json.Marshal(response)
		respJson = string(respJsonBytes)
	}()

	loginReq := &msg.CS_LoginReq{}
	err := json.Unmarshal([]byte(respJson), &loginReq)
	respMsgId = uint16(msg.MSG_ID_ELogin_Ack)
	if err != nil {
		errCode := msg.ELoginResult_ServerClosed
		response.LoginResult = &errCode
		log.Error("Invalid LoginReq,error:%s, %s", err.Error(), respJson)
		return
	}

	log.Debug("Recv LoginReq, UserName:%s Password:%s Platform:%d", *(loginReq.AccName), *(loginReq.AccPassword), *(loginReq.PlatForm))
	if App.DBIsReady() == false {
		log.Debug("DB is not ready. login failed")
		errCode := msg.ELoginResult_ServerClosed
		response.LoginResult = &errCode
		return
		//player.SendMsg(uint16(msg.MSG_ID_ELogin_Ack), response)
		//player.Close()
	}

	args := &Params.ReadParam{}
	args.TableName = "t_account"
	args.Keys = []string{"password"}
	args.Conditions = []Params.KeyValue{Params.KeyValue{"accountName", *(loginReq.AccName)}}

	//reply := &Params.ReadResult{}

	reqDBJson := fmt.Sprintf(`{"TableName":"%s",
							 "Key":"%s",
							 "Val":"%s"}`, "t_account", "accountName", *(loginReq.AccName))

	ret, _, respDBJson := App.dbServer.SendMsg("ReadTable", reqDBJson)

	log.Debug("ret:%v respJson:%v", ret, respDBJson)

	if ret == false {
		log.Error("DBService.Read rpc error:")
		errCode := msg.ELoginResult_InvalidAccOrPwd
		response.LoginResult = &errCode
		return
	}

	log.Debug("Recv LoginResult is:%s", respDBJson)

	var loginResult map[string]interface{}
	json.Unmarshal([]byte(respDBJson), &loginResult)
	log.Debug("loginResult:%v", loginResult)

	password := loginResult["password"]

	if password.(string) != *(loginReq.AccPassword) {
		log.Debug("Server Password is:%v client Password is:%v", password, *(loginReq.AccPassword))
		errCode := msg.ELoginResult_InvalidAccOrPwd
		response.LoginResult = &errCode
		return
	}

	retCode := msg.ELoginResult_Succeed
	response.LoginResult = &retCode
	//填写玩家基本信息
	return
}
