package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/xu756/imlogic/ent/schema/mixin"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

// Mixin of the UserGroup.
func (UserGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Mixin{},
	}
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户id"),
		field.Int64("group_id").Comment("组id"),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return nil
}
