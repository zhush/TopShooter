using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using msg;

namespace YFNet {
    public class Msg_Login
    {
        public static void handle_SC_LoginResponse(object msgData)
        {
            msg.SC_LoginResponse msg = msgData as msg.SC_LoginResponse;
            Debug.Log("Receive ,Result:" + msg.LoginResult + " RoleInfo:" + msg.PlayerBaseInfo);
			LoginResult loginResult = new LoginResult ();
			loginResult.Result = msg.LoginResult;
			loginResult.RoleList = new List<RoleInfo> ();
			for (int i = 0; i < msg.PlayerBaseInfo.Count; i++) {
				RoleBaseInfo rbi = msg.PlayerBaseInfo [i];
				loginResult.RoleList.Add (new RoleInfo (rbi.NickName, rbi.TemplateId, rbi.Level, 0, rbi.Sex, rbi.Gold));
			}
			AppFacade.Instance.SendNotification (NotifycationConstant.LoginGameResponse, loginResult);
        }
    }
}
