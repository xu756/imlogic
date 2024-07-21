package schema

import (
	"imlogic/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Mixin{},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("parent_id").Default(0).Comment("父角色id"),
		field.Int64("level").Default(0).Comment("权限层级"),
		field.String("role_name").NotEmpty().Comment("角色名"),
		field.String("intro").Comment("角色介绍"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
