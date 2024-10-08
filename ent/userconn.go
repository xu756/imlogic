// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"imlogic/ent/userconn"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserConn is the model entity for the UserConn schema.
type UserConn struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 连接id
	LinkID string `json:"link_id,omitempty"`
	// 连接时间
	LinkTime time.Time `json:"link_time,omitempty"`
	// 用户id
	UserID int64 `json:"user_id,omitempty"`
	// 主机ip
	HostIP string `json:"host_ip,omitempty"`
	// 设备
	Device string `json:"device,omitempty"`
	// 最后一次心跳时间
	LastHeartbeatTime time.Time `json:"last_heartbeat_time,omitempty"`
	selectValues      sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserConn) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userconn.FieldID, userconn.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userconn.FieldLinkID, userconn.FieldHostIP, userconn.FieldDevice:
			values[i] = new(sql.NullString)
		case userconn.FieldLinkTime, userconn.FieldLastHeartbeatTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserConn fields.
func (uc *UserConn) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userconn.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int(value.Int64)
		case userconn.FieldLinkID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link_id", values[i])
			} else if value.Valid {
				uc.LinkID = value.String
			}
		case userconn.FieldLinkTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field link_time", values[i])
			} else if value.Valid {
				uc.LinkTime = value.Time
			}
		case userconn.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uc.UserID = value.Int64
			}
		case userconn.FieldHostIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_ip", values[i])
			} else if value.Valid {
				uc.HostIP = value.String
			}
		case userconn.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				uc.Device = value.String
			}
		case userconn.FieldLastHeartbeatTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_heartbeat_time", values[i])
			} else if value.Valid {
				uc.LastHeartbeatTime = value.Time
			}
		default:
			uc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserConn.
// This includes values selected through modifiers, order, etc.
func (uc *UserConn) Value(name string) (ent.Value, error) {
	return uc.selectValues.Get(name)
}

// Update returns a builder for updating this UserConn.
// Note that you need to call UserConn.Unwrap() before calling this method if this UserConn
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserConn) Update() *UserConnUpdateOne {
	return NewUserConnClient(uc.config).UpdateOne(uc)
}

// Unwrap unwraps the UserConn entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserConn) Unwrap() *UserConn {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserConn is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserConn) String() string {
	var builder strings.Builder
	builder.WriteString("UserConn(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("link_id=")
	builder.WriteString(uc.LinkID)
	builder.WriteString(", ")
	builder.WriteString("link_time=")
	builder.WriteString(uc.LinkTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uc.UserID))
	builder.WriteString(", ")
	builder.WriteString("host_ip=")
	builder.WriteString(uc.HostIP)
	builder.WriteString(", ")
	builder.WriteString("device=")
	builder.WriteString(uc.Device)
	builder.WriteString(", ")
	builder.WriteString("last_heartbeat_time=")
	builder.WriteString(uc.LastHeartbeatTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserConns is a parsable slice of UserConn.
type UserConns []*UserConn
