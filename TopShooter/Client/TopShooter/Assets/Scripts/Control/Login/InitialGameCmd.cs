using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using PureMVC.Patterns;

public class InitialGameCmd : SimpleCommand {
	public const string NAME = "InitialGameCmd";

	public override void Execute(PureMVC.Interfaces.INotification notification){
		LoginRole role = Facade.RetrieveProxy (LoginRole.NAME) as LoginRole;
		role.LoadDataFromDataStore ();
	}
}
