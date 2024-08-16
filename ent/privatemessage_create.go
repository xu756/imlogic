// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/ent/privatemessage"
	"imlogic/kitex_gen/im"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrivateMessageCreate is the builder for creating a PrivateMessage entity.
type PrivateMessageCreate struct {
	config
	mutation *PrivateMessageMutation
	hooks    []Hook
}

// SetMsgType sets the "msg_type" field.
func (pmc *PrivateMessageCreate) SetMsgType(i int32) *PrivateMessageCreate {
	pmc.mutation.SetMsgType(i)
	return pmc
}

// SetMsgID sets the "msg_id" field.
func (pmc *PrivateMessageCreate) SetMsgID(s string) *PrivateMessageCreate {
	pmc.mutation.SetMsgID(s)
	return pmc
}

// SetChatID sets the "chat_id" field.
func (pmc *PrivateMessageCreate) SetChatID(i int64) *PrivateMessageCreate {
	pmc.mutation.SetChatID(i)
	return pmc
}

// SetSenderID sets the "sender_id" field.
func (pmc *PrivateMessageCreate) SetSenderID(i int64) *PrivateMessageCreate {
	pmc.mutation.SetSenderID(i)
	return pmc
}

// SetTimestamp sets the "timestamp" field.
func (pmc *PrivateMessageCreate) SetTimestamp(i int64) *PrivateMessageCreate {
	pmc.mutation.SetTimestamp(i)
	return pmc
}

// SetContent sets the "content" field.
func (pmc *PrivateMessageCreate) SetContent(i *im.Message) *PrivateMessageCreate {
	pmc.mutation.SetContent(i)
	return pmc
}

// Mutation returns the PrivateMessageMutation object of the builder.
func (pmc *PrivateMessageCreate) Mutation() *PrivateMessageMutation {
	return pmc.mutation
}

// Save creates the PrivateMessage in the database.
func (pmc *PrivateMessageCreate) Save(ctx context.Context) (*PrivateMessage, error) {
	return withHooks(ctx, pmc.sqlSave, pmc.mutation, pmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pmc *PrivateMessageCreate) SaveX(ctx context.Context) *PrivateMessage {
	v, err := pmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmc *PrivateMessageCreate) Exec(ctx context.Context) error {
	_, err := pmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmc *PrivateMessageCreate) ExecX(ctx context.Context) {
	if err := pmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmc *PrivateMessageCreate) check() error {
	if _, ok := pmc.mutation.MsgType(); !ok {
		return &ValidationError{Name: "msg_type", err: errors.New(`ent: missing required field "PrivateMessage.msg_type"`)}
	}
	if _, ok := pmc.mutation.MsgID(); !ok {
		return &ValidationError{Name: "msg_id", err: errors.New(`ent: missing required field "PrivateMessage.msg_id"`)}
	}
	if _, ok := pmc.mutation.ChatID(); !ok {
		return &ValidationError{Name: "chat_id", err: errors.New(`ent: missing required field "PrivateMessage.chat_id"`)}
	}
	if _, ok := pmc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender_id", err: errors.New(`ent: missing required field "PrivateMessage.sender_id"`)}
	}
	if _, ok := pmc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "PrivateMessage.timestamp"`)}
	}
	if _, ok := pmc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "PrivateMessage.content"`)}
	}
	return nil
}

func (pmc *PrivateMessageCreate) sqlSave(ctx context.Context) (*PrivateMessage, error) {
	if err := pmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pmc.mutation.id = &_node.ID
	pmc.mutation.done = true
	return _node, nil
}

func (pmc *PrivateMessageCreate) createSpec() (*PrivateMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &PrivateMessage{config: pmc.config}
		_spec = sqlgraph.NewCreateSpec(privatemessage.Table, sqlgraph.NewFieldSpec(privatemessage.FieldID, field.TypeInt))
	)
	if value, ok := pmc.mutation.MsgType(); ok {
		_spec.SetField(privatemessage.FieldMsgType, field.TypeInt32, value)
		_node.MsgType = value
	}
	if value, ok := pmc.mutation.MsgID(); ok {
		_spec.SetField(privatemessage.FieldMsgID, field.TypeString, value)
		_node.MsgID = value
	}
	if value, ok := pmc.mutation.ChatID(); ok {
		_spec.SetField(privatemessage.FieldChatID, field.TypeInt64, value)
		_node.ChatID = value
	}
	if value, ok := pmc.mutation.SenderID(); ok {
		_spec.SetField(privatemessage.FieldSenderID, field.TypeInt64, value)
		_node.SenderID = value
	}
	if value, ok := pmc.mutation.Timestamp(); ok {
		_spec.SetField(privatemessage.FieldTimestamp, field.TypeInt64, value)
		_node.Timestamp = value
	}
	if value, ok := pmc.mutation.Content(); ok {
		_spec.SetField(privatemessage.FieldContent, field.TypeJSON, value)
		_node.Content = value
	}
	return _node, _spec
}

// PrivateMessageCreateBulk is the builder for creating many PrivateMessage entities in bulk.
type PrivateMessageCreateBulk struct {
	config
	err      error
	builders []*PrivateMessageCreate
}

// Save creates the PrivateMessage entities in the database.
func (pmcb *PrivateMessageCreateBulk) Save(ctx context.Context) ([]*PrivateMessage, error) {
	if pmcb.err != nil {
		return nil, pmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pmcb.builders))
	nodes := make([]*PrivateMessage, len(pmcb.builders))
	mutators := make([]Mutator, len(pmcb.builders))
	for i := range pmcb.builders {
		func(i int, root context.Context) {
			builder := pmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrivateMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, pmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pmcb *PrivateMessageCreateBulk) SaveX(ctx context.Context) []*PrivateMessage {
	v, err := pmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmcb *PrivateMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := pmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmcb *PrivateMessageCreateBulk) ExecX(ctx context.Context) {
	if err := pmcb.Exec(ctx); err != nil {
		panic(err)
	}
}