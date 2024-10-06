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
	// FieldHostIP holds the string denoting the host_ip field in the database.
	FieldHostIP = "host_ip"
	// FieldDevice holds the string denoting the device field in the database.
	FieldDevice = "device"
	// FieldLastHeartbeatTime holds the string denoting the last_heartbeat_time field in the database.
	FieldLastHeartbeatTime = "last_heartbeat_time"
	// Table holds the table name of the userconn in the database.
	Table = "user_conns"
)

// Columns holds all SQL columns for userconn fields.
var Columns = []string{
	FieldID,
	FieldLinkID,
	FieldLinkTime,
	FieldUserID,
	FieldHostIP,
	FieldDevice,
	FieldLastHeartbeatTime,
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
	// DefaultHostIP holds the default value on creation for the "host_ip" field.
	DefaultHostIP string
	// DefaultDevice holds the default value on creation for the "device" field.
	DefaultDevice string
	// DefaultLastHeartbeatTime holds the default value on creation for the "last_heartbeat_time" field.
	DefaultLastHeartbeatTime func() time.Time
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

// ByHostIP orders the results by the host_ip field.
func ByHostIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostIP, opts...).ToFunc()
}

// ByDevice orders the results by the device field.
func ByDevice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDevice, opts...).ToFunc()
}

// ByLastHeartbeatTime orders the results by the last_heartbeat_time field.
func ByLastHeartbeatTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastHeartbeatTime, opts...).ToFunc()
}
