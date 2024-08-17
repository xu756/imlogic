package types

import (
	"imlogic/kitex_gen/base"
)

//
//// MsgMeta 消息元数据
//type MsgMeta struct {
//	Detail  string `json:"detail"`  //  connect | disconnect | heartbeat
//	Version string `json:"version"` // 1.0
//}
//
//type Message struct {
//	LinkId    string      `json:"link_id"`   // 连接id
//	MsgId     string      `json:"msg_id"`    // 消息id uuid 前端生成
//	Timestamp int64       `json:"timestamp"` // 消息时间戳
//	ChatType  ChatType    `json:"chat_type"` // 聊天类型	单聊 | 群聊 ｜ 系统消息｜ 通知
//	Sender    int64       `json:"sender"`    // 发送者
//	Receiver  int64       `json:"receiver"`  // 接收者
//	ChatId    int64       `json:"chat_id"`   // 聊天id
//	GroupId   int64       `json:"group_id"`  // 群id
//	Content   string      `json:"content"`   // 消息内容
//	MsgType   MsgType     `json:"msg_type"`  // 消息类型
//	Media     []MediaType `json:"media"`     // 媒体类型
//}
//
//func RpcMsgToMsg(rpcMsg *base.Message) (msg *Message) {
//	msg = &Message{}
//	msg.LinkId = rpcMsg.LinkId
//	msg.MsgId = rpcMsg.MsgId
//	msg.Timestamp = rpcMsg.Timestamp
//	msg.ChatType = ChatType(rpcMsg.ChatType)
//	msg.Sender = rpcMsg.Sender
//	msg.Receiver = rpcMsg.Receiver
//	msg.ChatId = rpcMsg.ChatId
//	msg.GroupId = rpcMsg.GroupId
//	msg.Content = rpcMsg.Content
//	msg.MsgType = MsgType(rpcMsg.MsgType)
//	medias := make([]MediaType, 0)
//	for _, media := range rpcMsg.Media {
//		medias = append(medias, MediaType{
//			UID: media.Uid,
//			URL: media.Url,
//		})
//	}
//	msg.Media = medias
//	return msg
//}

type MqPrivateMessage struct {
	Msg      *base.Message
	HostName string
	LinkId   string
}

type MqBroadcastMessage *base.Message
