using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LoginInfo  {


}

/*
 *玩家列表的玩家信息;
 */
public class RoleInfo{
	public string NickName;
	public int  TemplateId;
	public int Level;
	public long Uid;
	public int Sex;
	public long Gold;
	public RoleInfo(string NickName, int TemplateId, int Level, long Uid, int Sex, long Gold){
		this.NickName = NickName;
		this.TemplateId = TemplateId;
		this.Level = Level;
		this.Uid = Uid;
		this.Sex = Sex;
		this.Gold = Gold;
	}
}

/*
 *玩家登陆成功后的数据;
 */
public class LoginResult{
	public msg.ELoginResult    Result;
	public List<RoleInfo> RoleList;
}