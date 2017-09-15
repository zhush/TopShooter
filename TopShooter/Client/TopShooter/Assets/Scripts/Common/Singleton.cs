using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace YFUtils
{
	public class Singleton<T> where T: new()
	{
		protected static T _instance = default(T);
		protected static System.Object _objBlock = new System.Object();

		protected Singleton(){
			Debug.Assert (_instance == null);
		}

		public static T Instance
		{
			get
			{
				return instance ();
			}

		}
		public static T instance(){
			if (_instance == null) {
				lock (_objBlock)
				{
					if (_instance == null) {
						_instance = new T ();
					}
				}
			}
			return _instance;
		}
	}
}
