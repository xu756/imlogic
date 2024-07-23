package types

import (
	"encoding/json"
	"imlogic/kitex_gen/im"
	"log"
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
	From       int64      `json:"from"`        // 发送者 uuid(用户唯一)
	To         int64      `json:"to"`          // 接收者 uuid(用户唯一/组)
	MsgType    MsgType    `json:"msg_type"`    // 消息类型
	MsgContent MsgContent `json:"msg_content"` // 消息内容
}

func RpcMsgToMsg(rpcMsg *im.Message) (msg *Message) {

	// newMsg := Message{
	// 	LinkId:    msg.LinkId,
	// 	MsgId:     msg.MsgId,
	// 	Timestamp: msg.Timestamp,
	// 	ChatType:  ChatType(msg.ChatType),
	// 	From:      msg.From,
	// 	To:        msg.To,
	// 	MsgType:   MsgType(msg.GetMsgType()),
	// 	MsgContent: MsgContent{
	// 		Content: msg.GetText().Content,
	// 		// Img:     imgs,
	// 		File: FileType{
	// 			UID: msg.GetContent().File.Uid,
	// 			URL: msg.GetContent().File.Url,
	// 		},
	// 		Video: VideoType{
	// 			UID: msg.GetContent().Video.Uid,
	// 			URL: msg.GetContent().Video.Url,
	// 		},
	// 		Audio: AudioType{
	// 			UID: msg.GetContent().Audio.Uid,
	// 			URL: msg.GetContent().Audio.Url,
	// 		},
	// 	},
	// }
	jsonStr, err := json.Marshal(rpcMsg)
	if err != nil {
		log.Print(err)
		return nil
	}
	err = json.Unmarshal(jsonStr, &msg)
	if err != nil {
		log.Print(err)
		return nil
	}
	return msg
}

func RpcMsgResToMsg(msg *im.MessageRes) *Message {
	return &Message{}
}
