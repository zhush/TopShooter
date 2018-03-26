using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;


//玩家自己，继承玩家类.
public class GamePlayer : GameRole
{
	//玩家的武器列表;
	protected List<int>  m_weaponList;

	public GamePlayer ()
	{
	
	}

	public List<int>WeaponList{
		get{ 
			return m_weaponList;
		}
	}

}

