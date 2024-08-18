package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// UserFriend holds the schema definition for the UserFriend entity.
type UserFriend struct {
	ent.Schema
}

// Fields of the UserFriend.
func (UserFriend) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").Comment("id"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Int64("owner").Comment("用户id"),
		field.Int64("with_id").Comment("聊天对象"),
		field.String("alias").Comment("备注"),
		field.String("description").Comment("描述"),
	}
}

// Edges of the UserFriend.
func (UserFriend) Edges() []ent.Edge {
	return nil
}
