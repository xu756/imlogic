package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/xu756/imlogic/ent/schema/mixin"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Mixin of the UserRole.
func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Mixin{},
	}
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户id"),
		field.Int64("role_id").Comment("角色id"),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return nil
}
