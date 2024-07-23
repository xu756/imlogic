package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
			return "user-" + uuid.New().String()
		}).Comment("uuid"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间,判断顶置"),
		field.Int32("chat_type").Comment("聊天类型"),
		field.Int64("user_id").Comment("聊天属于者id"),
		field.Int64("chat_with").Comment("聊天对象id 对方 用户id 群id"),
		field.Bool("top").Default(false).Comment("是否顶置"),
		field.Bool("show").Default(true).Comment("是否在聊天列表中显示"),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return nil
}
