// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/ent/predicate"
	"imlogic/ent/userconn"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserConnUpdate is the builder for updating UserConn entities.
type UserConnUpdate struct {
	config
	hooks    []Hook
	mutation *UserConnMutation
}

// Where appends a list predicates to the UserConnUpdate builder.
func (ucu *UserConnUpdate) Where(ps ...predicate.UserConn) *UserConnUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetUserID sets the "user_id" field.
func (ucu *UserConnUpdate) SetUserID(i int64) *UserConnUpdate {
	ucu.mutation.ResetUserID()
	ucu.mutation.SetUserID(i)
	return ucu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucu *UserConnUpdate) SetNillableUserID(i *int64) *UserConnUpdate {
	if i != nil {
		ucu.SetUserID(*i)
	}
	return ucu
}

// AddUserID adds i to the "user_id" field.
func (ucu *UserConnUpdate) AddUserID(i int64) *UserConnUpdate {
	ucu.mutation.AddUserID(i)
	return ucu
}

// SetHostIP sets the "host_ip" field.
func (ucu *UserConnUpdate) SetHostIP(s string) *UserConnUpdate {
	ucu.mutation.SetHostIP(s)
	return ucu
}

// SetNillableHostIP sets the "host_ip" field if the given value is not nil.
func (ucu *UserConnUpdate) SetNillableHostIP(s *string) *UserConnUpdate {
	if s != nil {
		ucu.SetHostIP(*s)
	}
	return ucu
}

// SetDevice sets the "device" field.
func (ucu *UserConnUpdate) SetDevice(s string) *UserConnUpdate {
	ucu.mutation.SetDevice(s)
	return ucu
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (ucu *UserConnUpdate) SetNillableDevice(s *string) *UserConnUpdate {
	if s != nil {
		ucu.SetDevice(*s)
	}
	return ucu
}

// SetLastHeartbeatTime sets the "last_heartbeat_time" field.
func (ucu *UserConnUpdate) SetLastHeartbeatTime(t time.Time) *UserConnUpdate {
	ucu.mutation.SetLastHeartbeatTime(t)
	return ucu
}

// SetNillableLastHeartbeatTime sets the "last_heartbeat_time" field if the given value is not nil.
func (ucu *UserConnUpdate) SetNillableLastHeartbeatTime(t *time.Time) *UserConnUpdate {
	if t != nil {
		ucu.SetLastHeartbeatTime(*t)
	}
	return ucu
}

// Mutation returns the UserConnMutation object of the builder.
func (ucu *UserConnUpdate) Mutation() *UserConnMutation {
	return ucu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UserConnUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ucu.sqlSave, ucu.mutation, ucu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UserConnUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UserConnUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UserConnUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucu *UserConnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userconn.Table, userconn.Columns, sqlgraph.NewFieldSpec(userconn.FieldID, field.TypeInt))
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucu.mutation.UserID(); ok {
		_spec.SetField(userconn.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ucu.mutation.AddedUserID(); ok {
		_spec.AddField(userconn.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ucu.mutation.HostIP(); ok {
		_spec.SetField(userconn.FieldHostIP, field.TypeString, value)
	}
	if value, ok := ucu.mutation.Device(); ok {
		_spec.SetField(userconn.FieldDevice, field.TypeString, value)
	}
	if value, ok := ucu.mutation.LastHeartbeatTime(); ok {
		_spec.SetField(userconn.FieldLastHeartbeatTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userconn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucu.mutation.done = true
	return n, nil
}

// UserConnUpdateOne is the builder for updating a single UserConn entity.
type UserConnUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserConnMutation
}

// SetUserID sets the "user_id" field.
func (ucuo *UserConnUpdateOne) SetUserID(i int64) *UserConnUpdateOne {
	ucuo.mutation.ResetUserID()
	ucuo.mutation.SetUserID(i)
	return ucuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucuo *UserConnUpdateOne) SetNillableUserID(i *int64) *UserConnUpdateOne {
	if i != nil {
		ucuo.SetUserID(*i)
	}
	return ucuo
}

// AddUserID adds i to the "user_id" field.
func (ucuo *UserConnUpdateOne) AddUserID(i int64) *UserConnUpdateOne {
	ucuo.mutation.AddUserID(i)
	return ucuo
}

// SetHostIP sets the "host_ip" field.
func (ucuo *UserConnUpdateOne) SetHostIP(s string) *UserConnUpdateOne {
	ucuo.mutation.SetHostIP(s)
	return ucuo
}

// SetNillableHostIP sets the "host_ip" field if the given value is not nil.
func (ucuo *UserConnUpdateOne) SetNillableHostIP(s *string) *UserConnUpdateOne {
	if s != nil {
		ucuo.SetHostIP(*s)
	}
	return ucuo
}

// SetDevice sets the "device" field.
func (ucuo *UserConnUpdateOne) SetDevice(s string) *UserConnUpdateOne {
	ucuo.mutation.SetDevice(s)
	return ucuo
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (ucuo *UserConnUpdateOne) SetNillableDevice(s *string) *UserConnUpdateOne {
	if s != nil {
		ucuo.SetDevice(*s)
	}
	return ucuo
}

// SetLastHeartbeatTime sets the "last_heartbeat_time" field.
func (ucuo *UserConnUpdateOne) SetLastHeartbeatTime(t time.Time) *UserConnUpdateOne {
	ucuo.mutation.SetLastHeartbeatTime(t)
	return ucuo
}

// SetNillableLastHeartbeatTime sets the "last_heartbeat_time" field if the given value is not nil.
func (ucuo *UserConnUpdateOne) SetNillableLastHeartbeatTime(t *time.Time) *UserConnUpdateOne {
	if t != nil {
		ucuo.SetLastHeartbeatTime(*t)
	}
	return ucuo
}

// Mutation returns the UserConnMutation object of the builder.
func (ucuo *UserConnUpdateOne) Mutation() *UserConnMutation {
	return ucuo.mutation
}

// Where appends a list predicates to the UserConnUpdate builder.
func (ucuo *UserConnUpdateOne) Where(ps ...predicate.UserConn) *UserConnUpdateOne {
	ucuo.mutation.Where(ps...)
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UserConnUpdateOne) Select(field string, fields ...string) *UserConnUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UserConn entity.
func (ucuo *UserConnUpdateOne) Save(ctx context.Context) (*UserConn, error) {
	return withHooks(ctx, ucuo.sqlSave, ucuo.mutation, ucuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UserConnUpdateOne) SaveX(ctx context.Context) *UserConn {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UserConnUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UserConnUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucuo *UserConnUpdateOne) sqlSave(ctx context.Context) (_node *UserConn, err error) {
	_spec := sqlgraph.NewUpdateSpec(userconn.Table, userconn.Columns, sqlgraph.NewFieldSpec(userconn.FieldID, field.TypeInt))
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserConn.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userconn.FieldID)
		for _, f := range fields {
			if !userconn.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userconn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucuo.mutation.UserID(); ok {
		_spec.SetField(userconn.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ucuo.mutation.AddedUserID(); ok {
		_spec.AddField(userconn.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ucuo.mutation.HostIP(); ok {
		_spec.SetField(userconn.FieldHostIP, field.TypeString, value)
	}
	if value, ok := ucuo.mutation.Device(); ok {
		_spec.SetField(userconn.FieldDevice, field.TypeString, value)
	}
	if value, ok := ucuo.mutation.LastHeartbeatTime(); ok {
		_spec.SetField(userconn.FieldLastHeartbeatTime, field.TypeTime, value)
	}
	_node = &UserConn{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userconn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucuo.mutation.done = true
	return _node, nil
}
