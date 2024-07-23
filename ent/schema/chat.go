package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").Comment("id"),
		field.String("uuid").DefaultFunc(func() string {
			return "chat-" + uuid.New().String()
		}).Comment("聊天uuid"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Int64("user1_id").Comment("用户1 id"),
		field.Int64("user2_id").Comment("用户2 id"),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return nil
}

// Indexes of the Chat.
func (Chat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
	}
}
