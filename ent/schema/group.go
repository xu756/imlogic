package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/xu756/imlogic/ent/schema/mixin"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Mixin of the Group.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Mixin{},
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").DefaultFunc(func() string {
			return uuid.New().String()
		}).
			Unique().Immutable().Comment("组uuid"),
		field.Int64("parent_id").Default(0).Comment("父组id"),
		field.Int64("level").Default(0).Comment("组层级"),
		field.String("name").NotEmpty().Comment("组名"),
		field.String("intro").Comment("组介绍"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
