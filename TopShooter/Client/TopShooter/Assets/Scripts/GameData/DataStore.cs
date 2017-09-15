using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.IO;
using System.Xml;


//记录本地的数据信息类

public class DataStore {
	public const string DataStoreFileName = "datastore.xml";
	static volatile protected DataStore _instance = null;
	static public DataStore Instance{
		get
		{
			if (_instance == null) {
				_instance = new DataStore ();
			}
			return _instance;
		}
	}

	public string XmlFullPath{
		get{ 
			return YFUtils.GameUtility.GetLocalDataStorePath () + DataStoreFileName;
		}
	}

	protected DataStore(){
		checkXmlFileIsCreated ();
	}

	//检查本地的xml文件是否存在
	protected void checkXmlFileIsCreated(){
		string DataStorePath = YFUtils.GameUtility.GetLocalDataStorePath ();
		if (!Directory.Exists (DataStorePath)) {
			Directory.CreateDirectory (DataStorePath);
		}
		if (!File.Exists (this.XmlFullPath)) {
			XmlDocument doc = new XmlDocument ();
			XmlNode headNode = doc.CreateXmlDeclaration ("1.0", "utf-8", "yes");
			doc.AppendChild (headNode);
			XmlNode rootNode = doc.CreateElement ("root");
			doc.AppendChild (rootNode);
			doc.Save (this.XmlFullPath);
		}
	}
		
	//设置本地要保存的值
	public void SetKey(string key, string value){
		XmlDocument doc = new XmlDocument ();
		doc.Load (this.XmlFullPath);
		XmlNode rootNode= doc.SelectSingleNode ("root");
		XmlNode dstNode = rootNode.SelectSingleNode (key);
		if (dstNode == null) {
			dstNode = doc.CreateElement (key);
			rootNode.AppendChild (dstNode);
		}
		dstNode.InnerText = value;
		doc.Save (this.XmlFullPath);
	}

	//获取本地存储的值
	public string GetKey(string key){

		XmlDocument doc = new XmlDocument ();
		//doc.Load (reader);
		doc.Load (this.XmlFullPath);
		XmlNode rootNode= doc.SelectSingleNode ("root");
		XmlNode dstNode = rootNode.SelectSingleNode (key);
		if (dstNode == null) {
			return "";
		}
		return dstNode.InnerText;		
	}

}
