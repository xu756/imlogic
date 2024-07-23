package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").Comment("id"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.Bool("deleted").Default(false).Comment("删除状态"),
		field.String("uuid").DefaultFunc(func() string {
			return "user-" + uuid.New().String()
		}).Comment("群uuid"),
		field.Int64("editor").Default(0).Comment("修改人"),
		field.String("username").Default("").Comment("姓名"),
		field.String("password").Default("").Sensitive().Comment("密码"),
		field.String("mobile").Default("").Comment("手机号"),
		field.String("avatar").Default("").Comment("头像"),
		field.String("desc").Default("此用户没有任何记录").Comment("描述"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
		index.Fields("username", "mobile"),
	}
}
