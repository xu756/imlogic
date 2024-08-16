// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "uuid", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user1_id", Type: field.TypeInt64},
		{Name: "user2_id", Type: field.TypeInt64},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "chat_uuid",
				Unique:  true,
				Columns: []*schema.Column{ChatsColumns[1]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "uuid", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// GroupMessagesColumns holds the columns for the "group_messages" table.
	GroupMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "msg_type", Type: field.TypeInt32},
		{Name: "msg_id", Type: field.TypeString},
		{Name: "group_id", Type: field.TypeInt64},
		{Name: "timestamp", Type: field.TypeInt64},
		{Name: "sender_id", Type: field.TypeInt64},
		{Name: "content", Type: field.TypeJSON},
	}
	// GroupMessagesTable holds the schema information for the "group_messages" table.
	GroupMessagesTable = &schema.Table{
		Name:       "group_messages",
		Columns:    GroupMessagesColumns,
		PrimaryKey: []*schema.Column{GroupMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "groupmessage_msg_id_sender_id_group_id",
				Unique:  false,
				Columns: []*schema.Column{GroupMessagesColumns[2], GroupMessagesColumns[5], GroupMessagesColumns[3]},
			},
		},
	}
	// PrivateMessagesColumns holds the columns for the "private_messages" table.
	PrivateMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "msg_type", Type: field.TypeInt32},
		{Name: "msg_id", Type: field.TypeString},
		{Name: "chat_id", Type: field.TypeInt64},
		{Name: "sender_id", Type: field.TypeInt64},
		{Name: "timestamp", Type: field.TypeInt64},
		{Name: "content", Type: field.TypeJSON},
	}
	// PrivateMessagesTable holds the schema information for the "private_messages" table.
	PrivateMessagesTable = &schema.Table{
		Name:       "private_messages",
		Columns:    PrivateMessagesColumns,
		PrimaryKey: []*schema.Column{PrivateMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "privatemessage_msg_id_sender_id_chat_id",
				Unique:  false,
				Columns: []*schema.Column{PrivateMessagesColumns[2], PrivateMessagesColumns[4], PrivateMessagesColumns[3]},
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "group_id", Type: field.TypeInt64, Default: 0},
		{Name: "name", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString, Default: ""},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "uuid", Type: field.TypeString},
		{Name: "editor", Type: field.TypeInt64, Default: 0},
		{Name: "username", Type: field.TypeString, Default: ""},
		{Name: "password", Type: field.TypeString, Default: ""},
		{Name: "mobile", Type: field.TypeString, Default: ""},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "desc", Type: field.TypeString, Default: "此用户没有任何记录"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_uuid",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[4]},
			},
			{
				Name:    "user_username_mobile",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[6], UsersColumns[8]},
			},
		},
	}
	// UserConnsColumns holds the columns for the "user_conns" table.
	UserConnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "link_id", Type: field.TypeString, Unique: true},
		{Name: "link_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "host_name", Type: field.TypeString, Default: ""},
		{Name: "device", Type: field.TypeString, Default: ""},
		{Name: "last_heartbeat_time", Type: field.TypeTime},
	}
	// UserConnsTable holds the schema information for the "user_conns" table.
	UserConnsTable = &schema.Table{
		Name:       "user_conns",
		Columns:    UserConnsColumns,
		PrimaryKey: []*schema.Column{UserConnsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userconn_link_id",
				Unique:  true,
				Columns: []*schema.Column{UserConnsColumns[1]},
			},
			{
				Name:    "userconn_user_id_host_name",
				Unique:  false,
				Columns: []*schema.Column{UserConnsColumns[3], UserConnsColumns[4]},
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "join_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "group_id", Type: field.TypeInt64},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "usergroup_user_id_group_id",
				Unique:  false,
				Columns: []*schema.Column{UserGroupsColumns[2], UserGroupsColumns[3]},
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "role_id", Type: field.TypeInt64},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatsTable,
		GroupsTable,
		GroupMessagesTable,
		PrivateMessagesTable,
		RolesTable,
		UsersTable,
		UserConnsTable,
		UserGroupsTable,
		UserRolesTable,
	}
)

func init() {
}
