package types

import "github.com/xu756/imlogic/kitex_gen/im"

// MsgMeta 消息元数据
type MsgMeta struct {
	DetailType string `json:"detailType"` //  connect | disconnect | heartbeat
	Version    string `json:"version"`    // 1.0
}

// MsgContent 消息
type MsgContent struct {
	DetailType string `json:"detailType"` //  private | group | broadcast
	Text       string `json:"text"`       // 文本消息
	ImgUrl     string `json:"imgUrl"`     // 图片消息
	AudioUrl   string `json:"audioUrl"`   // 音频消息
	VideoUrl   string `json:"videoUrl"`   // 视频消息
}

type Message struct {
	LinkId     string            `json:"linkId"`    // 连接id
	MsgId      string            `json:"msgId"`     // 消息id uuid 前端生成
	Timestamp  int64             `json:"timestamp"` // 消息时间戳
	Device     string            `json:"device"`    // 设备类型 pc | mobile | pad | web | other
	Action     string            `json:"action"`    // 消息动作 send | receive | broadcast
	From       string            `json:"from"`      // 发送者 uuid
	To         string            `json:"to"`        // 接收者 uuid
	MsgType    string            `json:"msgType"`   // meta | text | image | file | audio | video | location | custom
	MsgMeta    MsgMeta           `json:"msgMeta"`   // msgType = meta 时，此字段有值
	MsgContent MsgContent        `json:"msgContent"`
	Params     map[string]string `json:"params"` // 消息参数
}

func RpcMsgToMsg(msg *im.Message) *Message {
	newMsg := Message{
		MsgId:     msg.MsgId,
		Timestamp: msg.Timestamp,
		Params:    msg.Params,
		Device:    msg.Device,
		Action:    msg.Action,
		From:      msg.From,
		To:        msg.To,
		MsgType:   msg.MsgType,
	}
	if msg.MsgMeta != nil {
		newMsg.MsgMeta = MsgMeta{
			DetailType: msg.MsgMeta.DetailType,
			Version:    msg.MsgMeta.Version,
		}
	}
	if msg.MsgContent != nil {
		newMsg.MsgContent = MsgContent{
			DetailType: msg.MsgContent.DetailType,
			Text:       msg.MsgContent.Text,
			ImgUrl:     msg.MsgContent.ImgUrl,
			AudioUrl:   msg.MsgContent.AudioUrl,
			VideoUrl:   msg.MsgContent.VideoUrl,
		}
	}
	return &newMsg
}

func RpcMsgResToMsg(msg *im.MessageRes) *Message {
	return &Message{
		MsgId: msg.MsgId,
		From:  msg.From,
		To:    msg.To,
	}
}
