using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;


public class AppFacade : Facade {


	public new static AppFacade Instance
	{
		get
		{
			if (m_instance == null)
			{
				lock (m_staticSyncRoot)
				{
					if (m_instance == null) m_instance = new AppFacade();
				}
			}

			return m_instance as AppFacade;
		}
	}


	// Use this for initialization
	protected AppFacade(){

	}

	public void StartUp(GameObject uiRoot){
		RegisterCommand (NotifycationConstant.InitialData, typeof(InitialGameCmd));
		RegisterCommand (NotifycationConstant.LoginGame, typeof(LoginGameCmd));
		RegisterCommand (NotifycationConstant.EnterGame, typeof(EnterGameCmd));
		RegisterCommand (NotifycationConstant.CreateRole, typeof(CreateRoleCmd));
		RegisterMediator (new LoginPanelMediator (uiRoot));
		RegisterProxy (new LoginRole ());
	}
}
