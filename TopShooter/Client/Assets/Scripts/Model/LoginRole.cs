using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;



public class LoginRole : Proxy {
	//登录的账户名称;
	protected string m_accName = "";
	//登录的密码的md5码
	protected string m_accPwdMD5 = "";
	//是否游客登录
	protected bool   m_isTourist = true;
	//最后一次登录的服务器ID.
	protected int	 m_lastServerId = 0;

	public string AccName{
		get{ 
			return m_accName;
		}
	}

	public string AccPwdMD5{
		get{ 
			return m_accPwdMD5;
		}		
	}

	public bool IsTourist{
		get{ 
			return m_isTourist;
		}		
	}

	public int LastServerId{
		get{ 
			return m_lastServerId;
		}		
	}


	public void LoadDataFromDataStore(){
		string loginInfo = DataStore.Instance.GetKey ("UserLoginInfo");
		//没有保存登录信息, 则是第一次登陆.
		if (loginInfo.Length == 0) {
			this.SetLoginInfo ("test", 
				YFUtils.GameUtility.MD5Encrypt ("test"),
				false,
				0);
		} else {
			SimpleJson.JsonObject jsonObject = SimpleJson.SimpleJson.DeserializeObject(loginInfo) as SimpleJson.JsonObject;
			if (jsonObject == null) {
				Debug.LogError ("Save Login Data is valid json:" + loginInfo);
			}

			m_accName = jsonObject ["AccName"].ToString ();
			m_accPwdMD5 = jsonObject ["AccPwd"].ToString ();
			m_isTourist = bool.Parse(jsonObject ["IsTourist"].ToString ());
			m_lastServerId = int.Parse (jsonObject ["LastServerId"].ToString ());
		}
	}

	public void SetLoginInfo(string accName, string accPwd, bool isTourist, int lastLoginServerId){
		m_accName = accName;
		m_accPwdMD5 = accPwd;
		m_isTourist = isTourist;
		m_lastServerId = lastLoginServerId;		

		SimpleJson.JsonObject saveJson = new SimpleJson.JsonObject ();
		saveJson ["AccName"] = m_accName;
		saveJson ["AccPwd"] = m_accPwdMD5;
		saveJson ["IsTourist"] = m_isTourist.ToString ();
		saveJson ["LastServerId"] = m_lastServerId.ToString ();
		DataStore.Instance.SetKey ("UserLoginInfo", saveJson.ToString());
	}
}
