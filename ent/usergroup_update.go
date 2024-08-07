// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/ent/predicate"
	"imlogic/ent/usergroup"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserGroupUpdate is the builder for updating UserGroup entities.
type UserGroupUpdate struct {
	config
	hooks    []Hook
	mutation *UserGroupMutation
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (ugu *UserGroupUpdate) Where(ps ...predicate.UserGroup) *UserGroupUpdate {
	ugu.mutation.Where(ps...)
	return ugu
}

// SetUserID sets the "user_id" field.
func (ugu *UserGroupUpdate) SetUserID(i int64) *UserGroupUpdate {
	ugu.mutation.ResetUserID()
	ugu.mutation.SetUserID(i)
	return ugu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ugu *UserGroupUpdate) SetNillableUserID(i *int64) *UserGroupUpdate {
	if i != nil {
		ugu.SetUserID(*i)
	}
	return ugu
}

// AddUserID adds i to the "user_id" field.
func (ugu *UserGroupUpdate) AddUserID(i int64) *UserGroupUpdate {
	ugu.mutation.AddUserID(i)
	return ugu
}

// SetGroupID sets the "group_id" field.
func (ugu *UserGroupUpdate) SetGroupID(i int64) *UserGroupUpdate {
	ugu.mutation.ResetGroupID()
	ugu.mutation.SetGroupID(i)
	return ugu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (ugu *UserGroupUpdate) SetNillableGroupID(i *int64) *UserGroupUpdate {
	if i != nil {
		ugu.SetGroupID(*i)
	}
	return ugu
}

// AddGroupID adds i to the "group_id" field.
func (ugu *UserGroupUpdate) AddGroupID(i int64) *UserGroupUpdate {
	ugu.mutation.AddGroupID(i)
	return ugu
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugu *UserGroupUpdate) Mutation() *UserGroupMutation {
	return ugu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ugu *UserGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ugu.sqlSave, ugu.mutation, ugu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ugu *UserGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ugu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ugu *UserGroupUpdate) Exec(ctx context.Context) error {
	_, err := ugu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugu *UserGroupUpdate) ExecX(ctx context.Context) {
	if err := ugu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ugu *UserGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(usergroup.Table, usergroup.Columns, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	if ps := ugu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ugu.mutation.UserID(); ok {
		_spec.SetField(usergroup.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ugu.mutation.AddedUserID(); ok {
		_spec.AddField(usergroup.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ugu.mutation.GroupID(); ok {
		_spec.SetField(usergroup.FieldGroupID, field.TypeInt64, value)
	}
	if value, ok := ugu.mutation.AddedGroupID(); ok {
		_spec.AddField(usergroup.FieldGroupID, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ugu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ugu.mutation.done = true
	return n, nil
}

// UserGroupUpdateOne is the builder for updating a single UserGroup entity.
type UserGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserGroupMutation
}

// SetUserID sets the "user_id" field.
func (uguo *UserGroupUpdateOne) SetUserID(i int64) *UserGroupUpdateOne {
	uguo.mutation.ResetUserID()
	uguo.mutation.SetUserID(i)
	return uguo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uguo *UserGroupUpdateOne) SetNillableUserID(i *int64) *UserGroupUpdateOne {
	if i != nil {
		uguo.SetUserID(*i)
	}
	return uguo
}

// AddUserID adds i to the "user_id" field.
func (uguo *UserGroupUpdateOne) AddUserID(i int64) *UserGroupUpdateOne {
	uguo.mutation.AddUserID(i)
	return uguo
}

// SetGroupID sets the "group_id" field.
func (uguo *UserGroupUpdateOne) SetGroupID(i int64) *UserGroupUpdateOne {
	uguo.mutation.ResetGroupID()
	uguo.mutation.SetGroupID(i)
	return uguo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (uguo *UserGroupUpdateOne) SetNillableGroupID(i *int64) *UserGroupUpdateOne {
	if i != nil {
		uguo.SetGroupID(*i)
	}
	return uguo
}

// AddGroupID adds i to the "group_id" field.
func (uguo *UserGroupUpdateOne) AddGroupID(i int64) *UserGroupUpdateOne {
	uguo.mutation.AddGroupID(i)
	return uguo
}

// Mutation returns the UserGroupMutation object of the builder.
func (uguo *UserGroupUpdateOne) Mutation() *UserGroupMutation {
	return uguo.mutation
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (uguo *UserGroupUpdateOne) Where(ps ...predicate.UserGroup) *UserGroupUpdateOne {
	uguo.mutation.Where(ps...)
	return uguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uguo *UserGroupUpdateOne) Select(field string, fields ...string) *UserGroupUpdateOne {
	uguo.fields = append([]string{field}, fields...)
	return uguo
}

// Save executes the query and returns the updated UserGroup entity.
func (uguo *UserGroupUpdateOne) Save(ctx context.Context) (*UserGroup, error) {
	return withHooks(ctx, uguo.sqlSave, uguo.mutation, uguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) SaveX(ctx context.Context) *UserGroup {
	node, err := uguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uguo *UserGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := uguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) ExecX(ctx context.Context) {
	if err := uguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uguo *UserGroupUpdateOne) sqlSave(ctx context.Context) (_node *UserGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(usergroup.Table, usergroup.Columns, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	id, ok := uguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usergroup.FieldID)
		for _, f := range fields {
			if !usergroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uguo.mutation.UserID(); ok {
		_spec.SetField(usergroup.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := uguo.mutation.AddedUserID(); ok {
		_spec.AddField(usergroup.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := uguo.mutation.GroupID(); ok {
		_spec.SetField(usergroup.FieldGroupID, field.TypeInt64, value)
	}
	if value, ok := uguo.mutation.AddedGroupID(); ok {
		_spec.AddField(usergroup.FieldGroupID, field.TypeInt64, value)
	}
	_node = &UserGroup{config: uguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uguo.mutation.done = true
	return _node, nil
}
