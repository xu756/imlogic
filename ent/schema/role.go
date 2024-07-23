package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").Comment("id"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Int64("group_id").Default(0).Comment("组id"),
		field.String("name").NotEmpty().Comment("角色名"),
		field.String("intro").Default("").Comment("角色介绍"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
