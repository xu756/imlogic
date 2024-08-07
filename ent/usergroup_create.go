// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/ent/usergroup"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserGroupCreate is the builder for creating a UserGroup entity.
type UserGroupCreate struct {
	config
	mutation *UserGroupMutation
	hooks    []Hook
}

// SetJoinAt sets the "join_at" field.
func (ugc *UserGroupCreate) SetJoinAt(t time.Time) *UserGroupCreate {
	ugc.mutation.SetJoinAt(t)
	return ugc
}

// SetNillableJoinAt sets the "join_at" field if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableJoinAt(t *time.Time) *UserGroupCreate {
	if t != nil {
		ugc.SetJoinAt(*t)
	}
	return ugc
}

// SetUserID sets the "user_id" field.
func (ugc *UserGroupCreate) SetUserID(i int64) *UserGroupCreate {
	ugc.mutation.SetUserID(i)
	return ugc
}

// SetGroupID sets the "group_id" field.
func (ugc *UserGroupCreate) SetGroupID(i int64) *UserGroupCreate {
	ugc.mutation.SetGroupID(i)
	return ugc
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugc *UserGroupCreate) Mutation() *UserGroupMutation {
	return ugc.mutation
}

// Save creates the UserGroup in the database.
func (ugc *UserGroupCreate) Save(ctx context.Context) (*UserGroup, error) {
	ugc.defaults()
	return withHooks(ctx, ugc.sqlSave, ugc.mutation, ugc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGroupCreate) SaveX(ctx context.Context) *UserGroup {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugc *UserGroupCreate) Exec(ctx context.Context) error {
	_, err := ugc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugc *UserGroupCreate) ExecX(ctx context.Context) {
	if err := ugc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ugc *UserGroupCreate) defaults() {
	if _, ok := ugc.mutation.JoinAt(); !ok {
		v := usergroup.DefaultJoinAt()
		ugc.mutation.SetJoinAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGroupCreate) check() error {
	if _, ok := ugc.mutation.JoinAt(); !ok {
		return &ValidationError{Name: "join_at", err: errors.New(`ent: missing required field "UserGroup.join_at"`)}
	}
	if _, ok := ugc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserGroup.user_id"`)}
	}
	if _, ok := ugc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "UserGroup.group_id"`)}
	}
	return nil
}

func (ugc *UserGroupCreate) sqlSave(ctx context.Context) (*UserGroup, error) {
	if err := ugc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ugc.mutation.id = &_node.ID
	ugc.mutation.done = true
	return _node, nil
}

func (ugc *UserGroupCreate) createSpec() (*UserGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGroup{config: ugc.config}
		_spec = sqlgraph.NewCreateSpec(usergroup.Table, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	)
	if value, ok := ugc.mutation.JoinAt(); ok {
		_spec.SetField(usergroup.FieldJoinAt, field.TypeTime, value)
		_node.JoinAt = value
	}
	if value, ok := ugc.mutation.UserID(); ok {
		_spec.SetField(usergroup.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := ugc.mutation.GroupID(); ok {
		_spec.SetField(usergroup.FieldGroupID, field.TypeInt64, value)
		_node.GroupID = value
	}
	return _node, _spec
}

// UserGroupCreateBulk is the builder for creating many UserGroup entities in bulk.
type UserGroupCreateBulk struct {
	config
	err      error
	builders []*UserGroupCreate
}

// Save creates the UserGroup entities in the database.
func (ugcb *UserGroupCreateBulk) Save(ctx context.Context) ([]*UserGroup, error) {
	if ugcb.err != nil {
		return nil, ugcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGroup, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) SaveX(ctx context.Context) []*UserGroup {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugcb *UserGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := ugcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) ExecX(ctx context.Context) {
	if err := ugcb.Exec(ctx); err != nil {
		panic(err)
	}
}
