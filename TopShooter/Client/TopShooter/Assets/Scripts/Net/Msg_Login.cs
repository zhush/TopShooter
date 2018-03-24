using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using msg;
using PureMVC.Patterns;

namespace YFNet {
    public class Msg_Login
    {
        public static void handle_SC_LoginResponse(object msgData)
        {
            msg.SC_LoginResponse rcvmsg = msgData as msg.SC_LoginResponse;
			Debug.Log("Receive ,Result:" + rcvmsg.LoginResult + " RoleInfo:" + rcvmsg.PlayerBaseInfo);

			LoginResult loginResult = new LoginResult ();
			loginResult.LoginAccName = rcvmsg.AccName;
			loginResult.Result = 0;
			Facade.SendNotification (NotifycationConstant.LoginGameResponse, loginResult);

        }
    }
}
