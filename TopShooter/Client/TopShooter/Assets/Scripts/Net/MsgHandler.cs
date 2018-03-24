using System.Collections;
using System.Collections.Generic;
using System;
using System.IO;
using UnityEngine;
using YFUtils;

namespace YFNet
{
	public delegate void HandlerType(object data);
	public class MsgHandler : Singleton<MsgHandler> 
	{

		protected Dictionary <msg.MSG_ID, HandlerType> mMsgEvent = new Dictionary<msg.MSG_ID, HandlerType>();
        protected Dictionary<msg.MSG_ID, Type> mMsgType = new Dictionary<msg.MSG_ID, Type>();
        public void Init(){
			this.RegisterAllMsgHandlers ();
		}

		public void RegisterMsgHandler(msg.MSG_ID msgId, Type msgType, HandlerType handler){
			mMsgEvent.Add (msgId, handler);
            mMsgType.Add(msgId, msgType);
        }

		protected void RegisterAllMsgHandlers(){
            this.RegisterMsgHandler(msg.MSG_ID.ELogin_Ack, typeof(msg.SC_LoginResponse), Msg_Login.handle_SC_LoginResponse);
		}

        public void Process(int msgId, byte[] buff) {
            if (mMsgType.ContainsKey((msg.MSG_ID)msgId) == false) {
                Debug.Log("Invalid msgId:" + msgId);
                return;
            }
            if (mMsgEvent.ContainsKey((msg.MSG_ID)msgId) == false) {
                Debug.Log(" msgId:" + msgId+" hasn't handler method");
                return;
            }

            Type protoType = mMsgType[(msg.MSG_ID)msgId];
            object toc = ProtoBuf.Serializer.Deserialize(protoType, new MemoryStream(buff));
            //object toc = ProtoBuf.Serializer.Deserialize<typeof(protoType)>(new MemoryStream(buff));
            mMsgEvent[(msg.MSG_ID)msgId](toc);
        }

	}
}
