using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;

public class GameRole : Proxy {
	//玩家昵称;
	protected string m_nickName = "";
	//玩家性别;
	protected int m_sex = 0;
	//玩家模板id;
	protected int   m_templateId = 0;
	//玩家唯一id;
	protected long   m_id = 0;
	//最后一次登录的服务器ID.
	protected int	 m_lastServerId = 0;	
	//是否初始化;
	protected bool   m_bInit = false;

	//玩家手上的武器;
	protected int   m_handleWeapon;
	//玩家等级;
	protected int   m_level;
	//玩家金币
	protected long  m_gold;

	public string NickName{
		get{ 
			return m_nickName;
		}
	}

	public int Sex{
		get{ 
			return m_sex;
		}		
	}

	public int TemplateId{
		get{ 
			return m_templateId;
		}		
	}

	public long Id{
		get{ 
			return m_id;
		}		
	}

	public int LastServerId{
		get{ 
			return m_lastServerId;
		}
	}

	public bool IsInit{
		get{
			return m_bInit;
		}
	}

	public int HandleWeapon{
		get{ 
			return m_handleWeapon;
		}
	}

	public int Level{
		get{ 
			return m_level;
		}
	}

	public long Gold{
		get{ 
			return m_gold;
		}
	}

}
	

