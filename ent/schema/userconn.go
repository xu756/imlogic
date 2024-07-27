package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserConn holds the schema definition for the UserConn entity.
type UserConn struct {
	ent.Schema
}

// Fields of the UserConn.
func (UserConn) Fields() []ent.Field {
	return []ent.Field{
		field.String("link_id").Unique().Immutable().Comment("连接id"),
		field.Time("link_time").Immutable().Default(time.Now).Comment("连接时间"),
		field.Int64("user_id").Comment("用户id"),
		field.String("host_name").Default("").Comment("主机名"),
		field.String("device").Default("").Comment("设备"),
		field.Time("last_heartbeat_time").Default(time.Now).Comment("最后一次心跳时间"),
	}
}

// Edges of the UserConn.
func (UserConn) Edges() []ent.Edge {
	return nil
}

// Indexes of the UserConn.
func (UserConn) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("link_id").Unique(),
		// index.
		index.Fields("user_id", "host_name"),
	}
}
