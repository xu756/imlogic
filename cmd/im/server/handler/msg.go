package handler

// MsgMeta 消息元数据
type MsgMeta struct {
	DetailType string `json:"detailType"` //  connect | disconnect | heartbeat
	Version    string `json:"version"`    // 1.0
	Interval   int64  `json:"interval"`   // detailType=heartbeat  到下次心跳的间隔，单位：毫秒
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
	MsgId      string            `json:"msgId"`     // 消息id uuid 前端生成
	Device     string            `json:"device"`    // 设备类型 pc | mobile | pad | web | other
	Timestamp  int64             `json:"timestamp"` // 消息时间戳
	Parms      map[string]string `json:"parms"`     // 消息参数
	Action     string            `json:"action"`    // 消息动作 send | receive | broadcast
	From       string            `json:"from"`      // 发送者 uuid
	To         string            `json:"to"`        // 接收者 uuid
	MsgType    string            `json:"msgType"`   // meta | text | image | file | audio | video | location | custom
	MsgMeta    MsgMeta           `json:"msgMeta"`   // msgType = meta 时，此字段有值
	MsgContent MsgContent        `json:"msgContent"`
}
