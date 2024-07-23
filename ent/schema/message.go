package schema

import (
	"imlogic/common/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("chat_id").Comment("聊天id"),
		field.Enum("msg_type").GoType(types.MsgType(0)).Comment("消息类型"),
		field.String("msg_id").Comment("消息id"),
		field.Int64("timestamp").Comment("消息时间戳"),
		field.Int64("sender_id").Comment("发送者id"),
		field.JSON("content", types.Message{}).Comment("消息内容"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}

// Indexes of the Message.
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("msg_id", "sender_id", "chat_id"),
	}
}
