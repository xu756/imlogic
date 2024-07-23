package schema

import (
	"imlogic/common/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// GroupMessage holds the schema definition for the GroupMessage entity.
type GroupMessage struct {
	ent.Schema
}

// Fields of the GroupMessage.
func (GroupMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("group_id").Comment("群id"),
		field.Int32("msg_type").Comment("消息类型"),
		field.String("msg_id").Comment("消息id"),
		field.Int64("timestamp").Comment("消息时间戳"),
		field.Int64("sender_id").Comment("发送者id"),
		field.JSON("content", types.Message{}).Comment("消息内容"),
	}
}

// Edges of the GroupMessage.
func (GroupMessage) Edges() []ent.Edge {
	return nil
}

// Indexes of the GroupMessage.
func (GroupMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("msg_id", "sender_id", "group_id"),
	}
}
