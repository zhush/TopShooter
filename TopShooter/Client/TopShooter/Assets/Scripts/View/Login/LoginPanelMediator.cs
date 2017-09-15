using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;
using UnityEngine.UI;

public class LoginPanelMediator : Mediator {
	public new const string NAME = "LoginPanelMediator";
	private Button m_btnEnterGame;

	public LoginPanelMediator(GameObject panelRoot){
		m_btnEnterGame = YFUtils.GameUtility.GetComponent<Button>(panelRoot, "BtnEnterGame");
		if (m_btnEnterGame == null) {
			Debug.LogError ("Can not find BtnEnterGame!!");
		}
		m_btnEnterGame.onClick.AddListener (OnClickEnterGame);
	}

	//点击按钮;
	private void OnClickEnterGame(){
		SendNotification (NotifycationConstant.LoginEnterGame);
	}

	//设置感兴趣的消息;
	public override IList<string> ListNotificationInterests(){
		IList<string> list = new List<string> ();
		list.Add (NotifycationConstant.LoginGameResponse);
		return list;
	}

	//处理感兴趣的消息;
	public override void HandleNotification(PureMVC.Interfaces.INotification notification){
		switch (notification.Name) {
		case NotifycationConstant.LoginGameResponse:
			{
				LoginResult resultInfo = notification.Body as LoginResult;
				Debug.Log ("Login response!, AccName:" + resultInfo.LoginAccName + " Result:" + resultInfo.Result);
			}
			break;
		default:
			break;
		}
	}

}
