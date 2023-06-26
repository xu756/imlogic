package xtype

type ConnData Message

type Sender struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Extra  string `json:"extra"`
}

type Message struct {
	MessageId        string      `json:"messageId"`        // MessageId 消息id 由服务端插入时生成
	Uuid             string      `json:"uuid"`             // UUID 客户端生成的id 由客户端生成 在客户端保证唯一性
	ConversationId   string      `json:"conversationId"`   // 发送到哪个会话
	ConversationType int32       `json:"conversationType"` // 会话类型
	Sender           Sender      `json:"sender"`           // 发送者
	Content          interface{} `json:"content"`          // 消息内容
	ContentType      int32       `json:"contentType"`      // 消息类型
	SendTime         int64       `json:"sendTime"`         // 发送时间 由客户端生成
	InsertTime       int64       `json:"insertTime"`       // 插入时间 由服务端生成
	Seq              int64       `json:"seq"`              // 在会话中的消息顺序
	NeedDecrypt      bool        `json:"needDecrypt"`      // 是否需要解密 （端对端加密技术，服务端无法解密）
	CountUnread      bool        `json:"countUnread"`      // 消息是否需要计入未读数
}
