package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Int64("user_id").Comment("用户id"),
		field.Int64("group_id").Comment("组id"),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (UserGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "group_id"),
	}
}
