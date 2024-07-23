// Code generated by ent, DO NOT EDIT.

package message

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChatID holds the string denoting the chat_id field in the database.
	FieldChatID = "chat_id"
	// FieldMsgType holds the string denoting the msg_type field in the database.
	FieldMsgType = "msg_type"
	// FieldMsgID holds the string denoting the msg_id field in the database.
	FieldMsgID = "msg_id"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldSenderID holds the string denoting the sender_id field in the database.
	FieldSenderID = "sender_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// Table holds the table name of the message in the database.
	Table = "messages"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldChatID,
	FieldMsgType,
	FieldMsgID,
	FieldTimestamp,
	FieldSenderID,
	FieldContent,
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

// OrderOption defines the ordering options for the Message queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChatID orders the results by the chat_id field.
func ByChatID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatID, opts...).ToFunc()
}

// ByMsgType orders the results by the msg_type field.
func ByMsgType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMsgType, opts...).ToFunc()
}

// ByMsgID orders the results by the msg_id field.
func ByMsgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMsgID, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// BySenderID orders the results by the sender_id field.
func BySenderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSenderID, opts...).ToFunc()
}