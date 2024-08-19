// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/ent/userfriend"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFriendCreate is the builder for creating a UserFriend entity.
type UserFriendCreate struct {
	config
	mutation *UserFriendMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ufc *UserFriendCreate) SetCreatedAt(t time.Time) *UserFriendCreate {
	ufc.mutation.SetCreatedAt(t)
	return ufc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ufc *UserFriendCreate) SetNillableCreatedAt(t *time.Time) *UserFriendCreate {
	if t != nil {
		ufc.SetCreatedAt(*t)
	}
	return ufc
}

// SetOwner sets the "owner" field.
func (ufc *UserFriendCreate) SetOwner(i int64) *UserFriendCreate {
	ufc.mutation.SetOwner(i)
	return ufc
}

// SetWithID sets the "with_id" field.
func (ufc *UserFriendCreate) SetWithID(i int64) *UserFriendCreate {
	ufc.mutation.SetWithID(i)
	return ufc
}

// SetAlias sets the "alias" field.
func (ufc *UserFriendCreate) SetAlias(s string) *UserFriendCreate {
	ufc.mutation.SetAlias(s)
	return ufc
}

// SetOwnerDesc sets the "owner_desc" field.
func (ufc *UserFriendCreate) SetOwnerDesc(s string) *UserFriendCreate {
	ufc.mutation.SetOwnerDesc(s)
	return ufc
}

// SetID sets the "id" field.
func (ufc *UserFriendCreate) SetID(i int64) *UserFriendCreate {
	ufc.mutation.SetID(i)
	return ufc
}

// Mutation returns the UserFriendMutation object of the builder.
func (ufc *UserFriendCreate) Mutation() *UserFriendMutation {
	return ufc.mutation
}

// Save creates the UserFriend in the database.
func (ufc *UserFriendCreate) Save(ctx context.Context) (*UserFriend, error) {
	ufc.defaults()
	return withHooks(ctx, ufc.sqlSave, ufc.mutation, ufc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ufc *UserFriendCreate) SaveX(ctx context.Context) *UserFriend {
	v, err := ufc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufc *UserFriendCreate) Exec(ctx context.Context) error {
	_, err := ufc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufc *UserFriendCreate) ExecX(ctx context.Context) {
	if err := ufc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufc *UserFriendCreate) defaults() {
	if _, ok := ufc.mutation.CreatedAt(); !ok {
		v := userfriend.DefaultCreatedAt()
		ufc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufc *UserFriendCreate) check() error {
	if _, ok := ufc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserFriend.created_at"`)}
	}
	if _, ok := ufc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "UserFriend.owner"`)}
	}
	if _, ok := ufc.mutation.WithID(); !ok {
		return &ValidationError{Name: "with_id", err: errors.New(`ent: missing required field "UserFriend.with_id"`)}
	}
	if _, ok := ufc.mutation.Alias(); !ok {
		return &ValidationError{Name: "alias", err: errors.New(`ent: missing required field "UserFriend.alias"`)}
	}
	if _, ok := ufc.mutation.OwnerDesc(); !ok {
		return &ValidationError{Name: "owner_desc", err: errors.New(`ent: missing required field "UserFriend.owner_desc"`)}
	}
	return nil
}

func (ufc *UserFriendCreate) sqlSave(ctx context.Context) (*UserFriend, error) {
	if err := ufc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ufc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ufc.mutation.id = &_node.ID
	ufc.mutation.done = true
	return _node, nil
}

func (ufc *UserFriendCreate) createSpec() (*UserFriend, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFriend{config: ufc.config}
		_spec = sqlgraph.NewCreateSpec(userfriend.Table, sqlgraph.NewFieldSpec(userfriend.FieldID, field.TypeInt64))
	)
	if id, ok := ufc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ufc.mutation.CreatedAt(); ok {
		_spec.SetField(userfriend.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ufc.mutation.Owner(); ok {
		_spec.SetField(userfriend.FieldOwner, field.TypeInt64, value)
		_node.Owner = value
	}
	if value, ok := ufc.mutation.WithID(); ok {
		_spec.SetField(userfriend.FieldWithID, field.TypeInt64, value)
		_node.WithID = value
	}
	if value, ok := ufc.mutation.Alias(); ok {
		_spec.SetField(userfriend.FieldAlias, field.TypeString, value)
		_node.Alias = value
	}
	if value, ok := ufc.mutation.OwnerDesc(); ok {
		_spec.SetField(userfriend.FieldOwnerDesc, field.TypeString, value)
		_node.OwnerDesc = value
	}
	return _node, _spec
}

// UserFriendCreateBulk is the builder for creating many UserFriend entities in bulk.
type UserFriendCreateBulk struct {
	config
	err      error
	builders []*UserFriendCreate
}

// Save creates the UserFriend entities in the database.
func (ufcb *UserFriendCreateBulk) Save(ctx context.Context) ([]*UserFriend, error) {
	if ufcb.err != nil {
		return nil, ufcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ufcb.builders))
	nodes := make([]*UserFriend, len(ufcb.builders))
	mutators := make([]Mutator, len(ufcb.builders))
	for i := range ufcb.builders {
		func(i int, root context.Context) {
			builder := ufcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFriendMutation)
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
					_, err = mutators[i+1].Mutate(root, ufcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ufcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, ufcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ufcb *UserFriendCreateBulk) SaveX(ctx context.Context) []*UserFriend {
	v, err := ufcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufcb *UserFriendCreateBulk) Exec(ctx context.Context) error {
	_, err := ufcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufcb *UserFriendCreateBulk) ExecX(ctx context.Context) {
	if err := ufcb.Exec(ctx); err != nil {
		panic(err)
	}
}
