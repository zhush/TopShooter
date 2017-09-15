using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;

public class GlobalDataInit : MonoBehaviour {

	public Transform m_uiRoot;

	void Awake(){
		new AppFacade (m_uiRoot.gameObject);
		Facade.Instance.SendNotification (NotifycationConstant.InitialData);
	}

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
		
	}
}
