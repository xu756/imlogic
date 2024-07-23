package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").Comment("id"),
		field.String("uuid").DefaultFunc(func() string {
			return "group-" + uuid.New().String()
		}).Comment("群uuid"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.String("name").NotEmpty().Comment("组名"),
		field.String("intro").Comment("组介绍"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
