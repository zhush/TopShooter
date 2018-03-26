using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using PureMVC.Patterns;

public class GlobalDataInit : MonoBehaviour {

	public Transform m_uiRoot;

	void Awake(){
		DontDestroyOnLoad (this.gameObject);
		AppFacade.Instance.StartUp (m_uiRoot.gameObject);
		AppFacade.Instance.SendNotification (NotifycationConstant.InitialData);
	}

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
		
	}
}
