package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"imlogic/kitex_gen/im"
)

// PrivateMessage holds the schema definition for the PrivateMessage entity.
type PrivateMessage struct {
	ent.Schema
}

// Fields of the PrivateMessage.
func (PrivateMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("msg_type").Comment("消息类型"),
		field.String("msg_id").Comment("消息id"),
		field.Int64("chat_id").Comment("聊天id"),
		field.Int64("sender_id").Comment("发送者id"),
		field.Int64("timestamp").Comment("消息时间戳"),
		field.JSON("content", &im.Message{}).Comment("消息内容"),
	}
}

// Edges of the PrivateMessage.
func (PrivateMessage) Edges() []ent.Edge {
	return nil
}

// Indexes of the PrivateMessage.
func (PrivateMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("msg_id", "sender_id", "chat_id"),
	}
}