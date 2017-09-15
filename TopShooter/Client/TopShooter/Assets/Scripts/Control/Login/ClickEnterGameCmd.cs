using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;

public class ClickEnterGameCmd : SimpleCommand {

	public const string NAME = "ClickEnterGameCmd";

	public override void Execute(PureMVC.Interfaces.INotification notification){
		LoginRole role = Facade.RetrieveProxy (LoginRole.NAME) as LoginRole;

		//模仿服务器消息，设定登录成功!
		LoginResult loginResult = new LoginResult ();
		loginResult.LoginAccName = role.AccName;
		loginResult.Result = 0;
		Facade.SendNotification (NotifycationConstant.LoginGameResponse, loginResult);
	}
		
}
