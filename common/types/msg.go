package types

import (
	"imlogic/kitex_gen/im"
)

// MsgMeta 消息元数据
type MsgMeta struct {
	Detail  string `json:"detail"`  //  connect | disconnect | heartbeat
	Version string `json:"version"` // 1.0
}

// MsgContent 消息
type MsgContent struct {
	Content string      `json:"content"` // 文本内容
	Img     []ImageType `json:"img"`     // 图片
	File    FileType    `json:"file"`    // 文件
	Video   VideoType   `json:"video"`   // 视频
	Audio   AudioType   `json:"audio"`   // 音频
}

type Message struct {
	MsgMeta    MsgMeta    `json:"msg_meta"`    // 消息元数据
	LinkId     string     `json:"link_id"`     // 连接id
	MsgId      string     `json:"msg_id"`      // 消息id uuid 前端生成
	Timestamp  int64      `json:"timestamp"`   // 消息时间戳
	ChatType   ChatType   `json:"chat_type"`   // 聊天类型	单聊 | 群聊 ｜ 系统消息｜ 通知
	Sender     int64      `json:"sender"`      // 发送者
	ChatId     int64      `json:"chat_id"`     // 聊天id
	GroupId    int64      `json:"group_id"`    // 群id
	MsgType    MsgType    `json:"msg_type"`    // 消息类型
	MsgContent MsgContent `json:"msg_content"` // 消息内容
}

func RpcMsgToMsg(rpcMsg *im.Message) (msg *Message) {
	msg = &Message{}
	msg.LinkId = rpcMsg.Common.LinkId
	msg.MsgId = rpcMsg.Common.MsgId
	msg.Timestamp = rpcMsg.Common.Timestamp
	msg.ChatType = ChatType(rpcMsg.Common.ChatType)
	msg.Sender = rpcMsg.Common.Sender
	msg.ChatId = rpcMsg.Common.ChatId
	msg.GroupId = rpcMsg.Common.GroupId
	msg.MsgType = MsgType(rpcMsg.Common.MsgType)
	msg.MsgContent.Content = rpcMsg.Content
	imgs := make([]ImageType, 0)
	for _, img := range rpcMsg.Img {
		imgs = append(imgs, ImageType{
			UID: img.Uid,
			URL: img.Url,
		})
	}
	msg.MsgContent.Img = imgs
	// msg.MsgContent.File = FileType{
	// 	URL: rpcMsg.File.Url,
	// 	UID: rpcMsg.File.Uid,
	// }
	// msg.MsgContent.Video = VideoType{
	// 	URL: rpcMsg.Video.Url,
	// 	UID: rpcMsg.Video.Uid,
	// }
	// msg.MsgContent.Audio = AudioType{
	// 	URL: rpcMsg.Audio.Url,
	// 	UID: rpcMsg.Audio.Uid,
	// }
	// msg.MsgContent.Audio = AudioType{
	// 	URL: rpcMsg.Audio.Url,
	// 	UID: rpcMsg.Audio.Uid,
	// }
	return msg

}
