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

            /*
			LoginResult loginResult = new LoginResult ();
			loginResult.LoginAccName = msg.AccName;
			loginResult.Result = 0;
			Facade.SendNotification (NotifycationConstant.LoginGameResponse, loginResult);
            */
        }
    }
}
