using System;
using System.Collections;
using System.Collections.Generic;

//存储所有玩家的数据
public class GameData
{
	protected GamePlayer m_Player;
	protected List<GameRole> m_OtherPlayers;

	public GameData (){
		
	}

	public GameRole FindRoleById(long id){
		if (id == m_Player.Id) {
			return m_Player;
		}
		for (int i = 0; i < m_OtherPlayers.Count; i++) {
			if (m_OtherPlayers[i].Id == id){
				return m_OtherPlayers [i];
			}
		}
		return null;
	}
}


