using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;

public class ClickEnterGameCmd : SimpleCommand {

	public const string NAME = "ClickEnterGameCmd";

	public override void Execute(PureMVC.Interfaces.INotification notification){
		LoginRole role = Facade.RetrieveProxy (LoginRole.NAME) as LoginRole;
		//向服务器发送登陆消息
		msg.CS_LoginReq sendmsg = new msg.CS_LoginReq();
		sendmsg.AccName = role.AccName;
		sendmsg.AccPassword = role.AccPwdMD5;
		sendmsg.PlatForm = msg.EPlatForm.Windows;
		YFNet.NetManager.Instance.SendMessage ((uint)msg.MSG_ID.ELogin_Req, sendmsg);
	}
		
}
