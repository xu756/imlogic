// Code generated by ent, DO NOT EDIT.

package userconn

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the userconn type in the database.
	Label = "user_conn"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLinkID holds the string denoting the link_id field in the database.
	FieldLinkID = "link_id"
	// FieldLinkTime holds the string denoting the link_time field in the database.
	FieldLinkTime = "link_time"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldHostName holds the string denoting the host_name field in the database.
	FieldHostName = "host_name"
	// FieldDevice holds the string denoting the device field in the database.
	FieldDevice = "device"
	// Table holds the table name of the userconn in the database.
	Table = "user_conns"
)

// Columns holds all SQL columns for userconn fields.
var Columns = []string{
	FieldID,
	FieldLinkID,
	FieldLinkTime,
	FieldUserID,
	FieldHostName,
	FieldDevice,
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
	// DefaultLinkTime holds the default value on creation for the "link_time" field.
	DefaultLinkTime func() time.Time
	// DefaultHostName holds the default value on creation for the "host_name" field.
	DefaultHostName string
	// DefaultDevice holds the default value on creation for the "device" field.
	DefaultDevice string
)

// OrderOption defines the ordering options for the UserConn queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLinkID orders the results by the link_id field.
func ByLinkID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLinkID, opts...).ToFunc()
}

// ByLinkTime orders the results by the link_time field.
func ByLinkTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLinkTime, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByHostName orders the results by the host_name field.
func ByHostName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostName, opts...).ToFunc()
}

// ByDevice orders the results by the device field.
func ByDevice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDevice, opts...).ToFunc()
}
