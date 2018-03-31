using System.Collections;
using System.Collections.Generic;
using System;
using UnityEngine;
using System.Security.Cryptography;
using System.Text;

namespace YFUtils
{
	public class GameUtility  {
		//获取子对象
		public static Transform GetChild(GameObject root, string path){
			Transform tra = root.transform.Find (path);
			if (tra == null) {
				Debug.LogError ("Can not find path:" + path + " in " + root.name);
			}
			return tra;
		}

		//获取子组件.
		public static T GetComponent<T>(GameObject root, string path){
			Transform tra = GetChild (root, path);
			T t = tra.GetComponent<T> ();
			if (t == null) {
				Debug.LogError ("Can not find Component, path:" + path + " component " + typeof(T).Name);
			}
			return t;
		}

		protected static string s_localDownloadPath = "";
		public static string GetLocalDownloadPath(){
			if (s_localWritePath.Length == 0) {
				#if UNITY_ANDROID
				s_localDownloadPath = "jar:file//"+Application.dataPath+"!/assets/";
				#elif UNITY_IPHONE
				s_localDownloadPath = Application.dataPath+"/Raw/";
				#else
				s_localDownloadPath = "file://" + Application.dataPath + "AssetBundles/";
				#endif
			}
			return s_localDownloadPath;
		}

		//获取本地储存datastore的路径
		protected static string s_localWritePath = "";
		public static string GetLocalDataStorePath(){
			if (s_localWritePath.Length == 0) {
				s_localWritePath = Application.dataPath + "/datastore/";
			}
			return s_localWritePath;
		}


        /// <summary>
        /// 唯一ID生成
        /// </summary>
        /// <returns></returns>
        public static string GetUniqueID()
        {
            string strDateTimeNumber = DateTime.Now.ToString("yyyyMMddHHmmssms");
            string strRandomResult = NextRandom(1000, 1).ToString();
            return strDateTimeNumber + strRandomResult;
        }

        public static string MD5Encrypt(string strText)
        {   
			/*

            MD5 md5 = new MD5CryptoServiceProvider();
            byte[] result = md5.ComputeHash(System.Text.Encoding.Default.GetBytes(strText));

            return System.Text.Encoding.Default.GetString(result);
	*/
			if (strText == null)
			{
				return null;
			}

			MD5 md5Hash = MD5.Create();
			// 将输入字符串转换为字节数组并计算哈希数据 
			byte[] data = md5Hash.ComputeHash(System.Text.Encoding.UTF8.GetBytes(strText));

			// 创建一个 Stringbuilder 来收集字节并创建字符串 
			StringBuilder sBuilder = new StringBuilder();

			// 循环遍历哈希数据的每一个字节并格式化为十六进制字符串 
			for (int i = 0; i < data.Length; i++)
			{
				sBuilder.Append(data[i].ToString("x2"));
			}

			// 返回十六进制字符串 
			return sBuilder.ToString();

        }

        private static int NextRandom(int numSeeds, int length)
        {
            // Create a byte array to hold the random value.  
            byte[] randomNumber = new byte[length];
            // Create a new instance of the RNGCryptoServiceProvider.  
            System.Security.Cryptography.RNGCryptoServiceProvider rng = new System.Security.Cryptography.RNGCryptoServiceProvider();
            // Fill the array with a random value.  
            rng.GetBytes(randomNumber);
            // Convert the byte to an uint value to make the modulus operation easier.  
            uint randomResult = 0x0;
            for (int i = 0; i < length; i++)
            {
                randomResult |= ((uint)randomNumber[i] << ((length - 1 - i) * 8));
            }
            return (int)(randomResult % numSeeds) + 1;
        }
	}
}
