// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"imlogic/ent/privatemessage"
	"imlogic/kitex_gen/im"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PrivateMessage is the model entity for the PrivateMessage schema.
type PrivateMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 消息类型
	MsgType int32 `json:"msg_type,omitempty"`
	// 消息id
	MsgID string `json:"msg_id,omitempty"`
	// 聊天id
	ChatID int64 `json:"chat_id,omitempty"`
	// 发送者id
	SenderID int64 `json:"sender_id,omitempty"`
	// 消息时间戳
	Timestamp int64 `json:"timestamp,omitempty"`
	// 消息内容
	Content      *im.Message `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PrivateMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case privatemessage.FieldContent:
			values[i] = new([]byte)
		case privatemessage.FieldID, privatemessage.FieldMsgType, privatemessage.FieldChatID, privatemessage.FieldSenderID, privatemessage.FieldTimestamp:
			values[i] = new(sql.NullInt64)
		case privatemessage.FieldMsgID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PrivateMessage fields.
func (pm *PrivateMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case privatemessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pm.ID = int(value.Int64)
		case privatemessage.FieldMsgType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field msg_type", values[i])
			} else if value.Valid {
				pm.MsgType = int32(value.Int64)
			}
		case privatemessage.FieldMsgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field msg_id", values[i])
			} else if value.Valid {
				pm.MsgID = value.String
			}
		case privatemessage.FieldChatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value.Valid {
				pm.ChatID = value.Int64
			}
		case privatemessage.FieldSenderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sender_id", values[i])
			} else if value.Valid {
				pm.SenderID = value.Int64
			}
		case privatemessage.FieldTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				pm.Timestamp = value.Int64
			}
		case privatemessage.FieldContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pm.Content); err != nil {
					return fmt.Errorf("unmarshal field content: %w", err)
				}
			}
		default:
			pm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PrivateMessage.
// This includes values selected through modifiers, order, etc.
func (pm *PrivateMessage) Value(name string) (ent.Value, error) {
	return pm.selectValues.Get(name)
}

// Update returns a builder for updating this PrivateMessage.
// Note that you need to call PrivateMessage.Unwrap() before calling this method if this PrivateMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *PrivateMessage) Update() *PrivateMessageUpdateOne {
	return NewPrivateMessageClient(pm.config).UpdateOne(pm)
}

// Unwrap unwraps the PrivateMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *PrivateMessage) Unwrap() *PrivateMessage {
	_tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PrivateMessage is not a transactional entity")
	}
	pm.config.driver = _tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *PrivateMessage) String() string {
	var builder strings.Builder
	builder.WriteString("PrivateMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pm.ID))
	builder.WriteString("msg_type=")
	builder.WriteString(fmt.Sprintf("%v", pm.MsgType))
	builder.WriteString(", ")
	builder.WriteString("msg_id=")
	builder.WriteString(pm.MsgID)
	builder.WriteString(", ")
	builder.WriteString("chat_id=")
	builder.WriteString(fmt.Sprintf("%v", pm.ChatID))
	builder.WriteString(", ")
	builder.WriteString("sender_id=")
	builder.WriteString(fmt.Sprintf("%v", pm.SenderID))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(fmt.Sprintf("%v", pm.Timestamp))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(fmt.Sprintf("%v", pm.Content))
	builder.WriteByte(')')
	return builder.String()
}

// PrivateMessages is a parsable slice of PrivateMessage.
type PrivateMessages []*PrivateMessage