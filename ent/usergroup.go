// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"imlogic/ent/usergroup"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserGroup is the model entity for the UserGroup schema.
type UserGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 加群时间
	JoinAt time.Time `json:"join_at,omitempty"`
	// 用户id
	UserID int64 `json:"user_id,omitempty"`
	// 组id
	GroupID      int64 `json:"group_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usergroup.FieldID, usergroup.FieldUserID, usergroup.FieldGroupID:
			values[i] = new(sql.NullInt64)
		case usergroup.FieldJoinAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserGroup fields.
func (ug *UserGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usergroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ug.ID = int(value.Int64)
		case usergroup.FieldJoinAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field join_at", values[i])
			} else if value.Valid {
				ug.JoinAt = value.Time
			}
		case usergroup.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ug.UserID = value.Int64
			}
		case usergroup.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				ug.GroupID = value.Int64
			}
		default:
			ug.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserGroup.
// This includes values selected through modifiers, order, etc.
func (ug *UserGroup) Value(name string) (ent.Value, error) {
	return ug.selectValues.Get(name)
}

// Update returns a builder for updating this UserGroup.
// Note that you need to call UserGroup.Unwrap() before calling this method if this UserGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (ug *UserGroup) Update() *UserGroupUpdateOne {
	return NewUserGroupClient(ug.config).UpdateOne(ug)
}

// Unwrap unwraps the UserGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ug *UserGroup) Unwrap() *UserGroup {
	_tx, ok := ug.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserGroup is not a transactional entity")
	}
	ug.config.driver = _tx.drv
	return ug
}

// String implements the fmt.Stringer.
func (ug *UserGroup) String() string {
	var builder strings.Builder
	builder.WriteString("UserGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ug.ID))
	builder.WriteString("join_at=")
	builder.WriteString(ug.JoinAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ug.UserID))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", ug.GroupID))
	builder.WriteByte(')')
	return builder.String()
}

// UserGroups is a parsable slice of UserGroup.
type UserGroups []*UserGroup
