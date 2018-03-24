using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using YFNet;
using PureMVC.Patterns;

public class InitialGameCmd : SimpleCommand {
	public const string NAME = "InitialGameCmd";

	public override void Execute(PureMVC.Interfaces.INotification notification){
		LoginRole role = Facade.RetrieveProxy (LoginRole.NAME) as LoginRole;
		role.LoadDataFromDataStore ();
		NetManager.Instance.Init ();
		NetManager.Instance.ConnectServer ();
	}
}
