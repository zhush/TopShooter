using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;


public class AppFacade : Facade {

	// Use this for initialization
	public AppFacade(GameObject uiRoot){
		RegisterCommand (NotifycationConstant.InitialData, typeof(InitialGameCmd));
		RegisterCommand (NotifycationConstant.LoginEnterGame, typeof(ClickEnterGameCmd));
		RegisterMediator (new LoginPanelMediator (uiRoot));
		RegisterProxy (new LoginRole ());
	}
}
