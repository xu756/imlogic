// Code generated by ent, DO NOT EDIT.

package role

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// Table holds the table name of the role in the database.
	Table = "roles"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldGroupID,
	FieldName,
	FieldIntro,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultGroupID holds the default value on creation for the "group_id" field.
	DefaultGroupID int64
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultIntro holds the default value on creation for the "intro" field.
	DefaultIntro string
)

// OrderOption defines the ordering options for the Role queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByGroupID orders the results by the group_id field.
func ByGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIntro orders the results by the intro field.
func ByIntro(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntro, opts...).ToFunc()
}
