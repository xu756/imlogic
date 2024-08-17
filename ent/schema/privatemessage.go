package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"imlogic/kitex_gen/base"
)

// PrivateMessage holds the schema definition for the PrivateMessage entity.
type PrivateMessage struct {
	ent.Schema
}

// Fields of the PrivateMessage.
func (PrivateMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("msg_id").Comment("消息id"),
		field.Int64("msg_type").Comment("消息类型"),
		field.Int64("sender_id").Comment("发送者id"),
		field.Int64("receiver_id").Comment("接收者id"),
		field.Int64("timestamp").Comment("消息时间戳"),
		field.JSON("content", &base.Message{}).Comment("消息内容"),
	}
}

// Edges of the PrivateMessage.
func (PrivateMessage) Edges() []ent.Edge {
	return nil
}

// Indexes of the PrivateMessage.
func (PrivateMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("msg_id", "sender_id", "receiver_id"),
	}
}
