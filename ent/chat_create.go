// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/ent/chat"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatCreate is the builder for creating a Chat entity.
type ChatCreate struct {
	config
	mutation *ChatMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (cc *ChatCreate) SetUUID(s string) *ChatCreate {
	cc.mutation.SetUUID(s)
	return cc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (cc *ChatCreate) SetNillableUUID(s *string) *ChatCreate {
	if s != nil {
		cc.SetUUID(*s)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ChatCreate) SetCreatedAt(t time.Time) *ChatCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ChatCreate) SetNillableCreatedAt(t *time.Time) *ChatCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUser1ID sets the "user1_id" field.
func (cc *ChatCreate) SetUser1ID(i int64) *ChatCreate {
	cc.mutation.SetUser1ID(i)
	return cc
}

// SetUser2ID sets the "user2_id" field.
func (cc *ChatCreate) SetUser2ID(i int64) *ChatCreate {
	cc.mutation.SetUser2ID(i)
	return cc
}

// SetID sets the "id" field.
func (cc *ChatCreate) SetID(i int64) *ChatCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the ChatMutation object of the builder.
func (cc *ChatCreate) Mutation() *ChatMutation {
	return cc.mutation
}

// Save creates the Chat in the database.
func (cc *ChatCreate) Save(ctx context.Context) (*Chat, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChatCreate) SaveX(ctx context.Context) *Chat {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChatCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChatCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChatCreate) defaults() {
	if _, ok := cc.mutation.UUID(); !ok {
		v := chat.DefaultUUID()
		cc.mutation.SetUUID(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := chat.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChatCreate) check() error {
	if _, ok := cc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Chat.uuid"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Chat.created_at"`)}
	}
	if _, ok := cc.mutation.User1ID(); !ok {
		return &ValidationError{Name: "user1_id", err: errors.New(`ent: missing required field "Chat.user1_id"`)}
	}
	if _, ok := cc.mutation.User2ID(); !ok {
		return &ValidationError{Name: "user2_id", err: errors.New(`ent: missing required field "Chat.user2_id"`)}
	}
	return nil
}

func (cc *ChatCreate) sqlSave(ctx context.Context) (*Chat, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChatCreate) createSpec() (*Chat, *sqlgraph.CreateSpec) {
	var (
		_node = &Chat{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chat.Table, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.UUID(); ok {
		_spec.SetField(chat.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(chat.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.User1ID(); ok {
		_spec.SetField(chat.FieldUser1ID, field.TypeInt64, value)
		_node.User1ID = value
	}
	if value, ok := cc.mutation.User2ID(); ok {
		_spec.SetField(chat.FieldUser2ID, field.TypeInt64, value)
		_node.User2ID = value
	}
	return _node, _spec
}

// ChatCreateBulk is the builder for creating many Chat entities in bulk.
type ChatCreateBulk struct {
	config
	err      error
	builders []*ChatCreate
}

// Save creates the Chat entities in the database.
func (ccb *ChatCreateBulk) Save(ctx context.Context) ([]*Chat, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chat, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChatCreateBulk) SaveX(ctx context.Context) []*Chat {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChatCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChatCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
