package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// Mixin 这是公用结构，每一个数据库都会有创建时间、更新时间、删除状态、版本号
type Mixin struct {
	mixin.Schema
}

func (Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().
			StorageKey("id").Comment("id"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.Bool("deleted").Default(false).Comment("删除状态"),
		field.Int64("creator").Default(0).Comment("创建人"),
		field.Int64("editor").Default(1).Comment("修改人"),
		field.Int64("version").Default(0).Comment("版本号"),
	}

}
