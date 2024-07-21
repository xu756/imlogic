package schema

import (
	"imlogic/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Mixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").DefaultFunc(func() string {
			return uuid.New().String()
		}).
			Unique().Immutable().Comment("用户uuid"),
		field.String("username").Comment("姓名"),
		field.String("password").Comment("密码"),
		field.String("mobile").Unique().Comment("手机号"),
		field.String("avatar").Comment("头像").Default("https://cos.imlogic.cn/appadmin/images/avatar.jpeg"),
		field.String("device").Default("mini").Immutable().Comment("设备"),
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
