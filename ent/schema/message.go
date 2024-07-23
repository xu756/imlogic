package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").Comment("id"),
		field.Int32("chat_type").Comment("类型"),
		field.String("msg_id").Comment("消息id uuid 前端生成"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
