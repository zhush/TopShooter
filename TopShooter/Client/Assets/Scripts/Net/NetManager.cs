using System.Collections;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using UnityEngine;
using YFUtils;
using System;
using System.Text;
using System.IO;

namespace YFNet
{
	public class NetManager : Singleton<NetManager> 
	{
		protected NetworkStream mConnStream;
		protected TcpClient mClient;
		protected const int MAX_RECV_SIZE = 8096;
		protected byte[] mRecvBuf = new byte[MAX_RECV_SIZE];
		protected MemoryStream mMemStream;
		protected BinaryReader mBinaryReader;

		// Use this for initialization
		public void Init () {
			mClient = new TcpClient ();
			mClient.SendTimeout = 2000;
			mClient.ReceiveTimeout = 2000;
			mClient.NoDelay = true;
			mMemStream = new MemoryStream ();
			mBinaryReader = new BinaryReader (mMemStream);
			MsgHandler.Instance.Init ();
		}

		public void ConnectServer(){
			if (mClient.Connected) {
				Debug.Log ("Client Close Socket 1");
				mClient.Close ();
			}

			try{
				mClient.BeginConnect (Constant.ServerIp, Constant.Port, new System.AsyncCallback(OnConnected), mClient);
			}catch(Exception ex){
				Debug.Log ("Connected Server Failed!:"+ex.Message);
			}
				
		}

		//连接回调，可能成功也可能失败;
		protected void OnConnected(IAsyncResult ar){
			if (mClient.Connected) {
				Debug.Log ("Establish Succeed");

				mConnStream = mClient.GetStream ();
				mConnStream.BeginRead (mRecvBuf, 0, mRecvBuf.Length, new System.AsyncCallback (OnRead), null);

			} else {
				Debug.Log ("Establish Failed");
			}
			mClient.EndConnect (ar);
		}

		protected void Close(){
			if (mClient.Connected) {
				mClient.Close ();
			}
		}

		//读取到消息
		protected void OnRead(IAsyncResult ar){
			int nReadSize = 0;
			try{
				lock(mConnStream)
				{
					nReadSize = mConnStream.EndRead(ar);
				}
				if(nReadSize < 1){
					Debug.Log("Server Disconnected..");
					Close();
				}
				OnReceiveBytes(mRecvBuf, nReadSize);
				mConnStream.BeginRead (mRecvBuf, 0, mRecvBuf.Length, new System.AsyncCallback (OnRead), null);
			}catch(ObjectDisposedException ex){
				Debug.Log ("Server Disconnected..2" + ex.Message);
			}catch(Exception ex){
				Debug.Log ("Read Error:" + ex.Message);
				Close ();
			}
		}

		protected void OnReceiveBytes(byte[] bytes, int nLen){
			Debug.Log ("Call OnReceiveBytes, nLen:" + nLen);
			mMemStream.Seek (0, SeekOrigin.End);
			mMemStream.Write (bytes, 0, nLen);

			mMemStream.Seek (0, SeekOrigin.Begin);

			while (RemainBytesLength () > 2) {
				ushort msgLen = mBinaryReader.ReadUInt16 ();
				Debug.Log ("Call OnReceiveBytes, msgLen:" + msgLen);
				if (RemainBytesLength () >= msgLen) {
					MemoryStream memStream = new MemoryStream ();
					BinaryWriter writer = new BinaryWriter (memStream);
					writer.Write (mBinaryReader.ReadBytes (msgLen));
					OnRecvMessage (memStream);
				} else {
					mMemStream.Position = mMemStream.Position - 2;
					break;
				}
			}


		}

		protected int RemainBytesLength(){
			int nLeftLen = (int)(mMemStream.Length - mMemStream.Position);
			return nLeftLen;
		}

		protected void OnRecvMessage(MemoryStream ms){
			Debug.Log ("Call OnRecvMessage");
			string msg = Encoding.ASCII.GetString (ms.GetBuffer ());
			Debug.Log ("Recv Msg:" + msg);

            ByteBuffer buff = new ByteBuffer(ms.GetBuffer());

            int msgId = buff.ReadShort();

            Debug.Log("MsgId:" + msgId);

            byte[] msgData = buff.ReadBytes((int)ms.Length - 2);
            MsgHandler.Instance.Process(msgId, msgData);
        }

		//写消息
		public void WriteMessage(byte[] message){
			MemoryStream ms;
			using (ms = new MemoryStream ()) 
			{
				BinaryWriter writer = new BinaryWriter (ms);
				ushort msgLen = (ushort)message.Length;
				writer.Write (msgLen);
				writer.Write (message);
				writer.Flush ();
				if (mClient != null && mClient.Connected) {
					byte[] payload = ms.ToArray ();
					lock (mConnStream) {
						mConnStream.BeginWrite (payload, 0, payload.Length, new System.AsyncCallback (OnWrite), null);
					}
				}
			}
		}

		public void OnWrite(IAsyncResult ar){
			try{
				mConnStream.EndWrite(ar);
			}catch(Exception ex){
				Debug.Log ("Client Close Socket 3");
				Close();
				Debug.Log("Write Buffer Failed!!:"+ex.Message);
			}
		}

        //写网络消息(消息ID+消息体);
        public void SendMessage(uint msgId, object msg) {
            ByteBuffer buff = new ByteBuffer();
            buff.WriteShort((ushort)msgId);
            MemoryStream ms = new MemoryStream();
            ProtoBuf.Serializer.Serialize(ms, msg);
            byte[] result = ms.ToArray();
			byte[] headBuf = buff.ToBytes ();
			List<byte> tmp = new List<byte> (headBuf.Length + result.Length);
			tmp.AddRange (headBuf);
			tmp.AddRange (result);
			WriteMessage(tmp.ToArray());
        }
	}
}
