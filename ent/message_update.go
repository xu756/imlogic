// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imlogic/common/types"
	"imlogic/ent/message"
	"imlogic/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetChatID sets the "chat_id" field.
func (mu *MessageUpdate) SetChatID(i int64) *MessageUpdate {
	mu.mutation.ResetChatID()
	mu.mutation.SetChatID(i)
	return mu
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableChatID(i *int64) *MessageUpdate {
	if i != nil {
		mu.SetChatID(*i)
	}
	return mu
}

// AddChatID adds i to the "chat_id" field.
func (mu *MessageUpdate) AddChatID(i int64) *MessageUpdate {
	mu.mutation.AddChatID(i)
	return mu
}

// SetMsgType sets the "msg_type" field.
func (mu *MessageUpdate) SetMsgType(i int32) *MessageUpdate {
	mu.mutation.ResetMsgType()
	mu.mutation.SetMsgType(i)
	return mu
}

// SetNillableMsgType sets the "msg_type" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableMsgType(i *int32) *MessageUpdate {
	if i != nil {
		mu.SetMsgType(*i)
	}
	return mu
}

// AddMsgType adds i to the "msg_type" field.
func (mu *MessageUpdate) AddMsgType(i int32) *MessageUpdate {
	mu.mutation.AddMsgType(i)
	return mu
}

// SetMsgID sets the "msg_id" field.
func (mu *MessageUpdate) SetMsgID(s string) *MessageUpdate {
	mu.mutation.SetMsgID(s)
	return mu
}

// SetNillableMsgID sets the "msg_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableMsgID(s *string) *MessageUpdate {
	if s != nil {
		mu.SetMsgID(*s)
	}
	return mu
}

// SetTimestamp sets the "timestamp" field.
func (mu *MessageUpdate) SetTimestamp(i int64) *MessageUpdate {
	mu.mutation.ResetTimestamp()
	mu.mutation.SetTimestamp(i)
	return mu
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableTimestamp(i *int64) *MessageUpdate {
	if i != nil {
		mu.SetTimestamp(*i)
	}
	return mu
}

// AddTimestamp adds i to the "timestamp" field.
func (mu *MessageUpdate) AddTimestamp(i int64) *MessageUpdate {
	mu.mutation.AddTimestamp(i)
	return mu
}

// SetSenderID sets the "sender_id" field.
func (mu *MessageUpdate) SetSenderID(i int64) *MessageUpdate {
	mu.mutation.ResetSenderID()
	mu.mutation.SetSenderID(i)
	return mu
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableSenderID(i *int64) *MessageUpdate {
	if i != nil {
		mu.SetSenderID(*i)
	}
	return mu
}

// AddSenderID adds i to the "sender_id" field.
func (mu *MessageUpdate) AddSenderID(i int64) *MessageUpdate {
	mu.mutation.AddSenderID(i)
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(t types.Message) *MessageUpdate {
	mu.mutation.SetContent(t)
	return mu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableContent(t *types.Message) *MessageUpdate {
	if t != nil {
		mu.SetContent(*t)
	}
	return mu
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.ChatID(); ok {
		_spec.SetField(message.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedChatID(); ok {
		_spec.AddField(message.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.MsgType(); ok {
		_spec.SetField(message.FieldMsgType, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedMsgType(); ok {
		_spec.AddField(message.FieldMsgType, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.MsgID(); ok {
		_spec.SetField(message.FieldMsgID, field.TypeString, value)
	}
	if value, ok := mu.mutation.Timestamp(); ok {
		_spec.SetField(message.FieldTimestamp, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedTimestamp(); ok {
		_spec.AddField(message.FieldTimestamp, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.SenderID(); ok {
		_spec.SetField(message.FieldSenderID, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedSenderID(); ok {
		_spec.AddField(message.FieldSenderID, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeJSON, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetChatID sets the "chat_id" field.
func (muo *MessageUpdateOne) SetChatID(i int64) *MessageUpdateOne {
	muo.mutation.ResetChatID()
	muo.mutation.SetChatID(i)
	return muo
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableChatID(i *int64) *MessageUpdateOne {
	if i != nil {
		muo.SetChatID(*i)
	}
	return muo
}

// AddChatID adds i to the "chat_id" field.
func (muo *MessageUpdateOne) AddChatID(i int64) *MessageUpdateOne {
	muo.mutation.AddChatID(i)
	return muo
}

// SetMsgType sets the "msg_type" field.
func (muo *MessageUpdateOne) SetMsgType(i int32) *MessageUpdateOne {
	muo.mutation.ResetMsgType()
	muo.mutation.SetMsgType(i)
	return muo
}

// SetNillableMsgType sets the "msg_type" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableMsgType(i *int32) *MessageUpdateOne {
	if i != nil {
		muo.SetMsgType(*i)
	}
	return muo
}

// AddMsgType adds i to the "msg_type" field.
func (muo *MessageUpdateOne) AddMsgType(i int32) *MessageUpdateOne {
	muo.mutation.AddMsgType(i)
	return muo
}

// SetMsgID sets the "msg_id" field.
func (muo *MessageUpdateOne) SetMsgID(s string) *MessageUpdateOne {
	muo.mutation.SetMsgID(s)
	return muo
}

// SetNillableMsgID sets the "msg_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableMsgID(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetMsgID(*s)
	}
	return muo
}

// SetTimestamp sets the "timestamp" field.
func (muo *MessageUpdateOne) SetTimestamp(i int64) *MessageUpdateOne {
	muo.mutation.ResetTimestamp()
	muo.mutation.SetTimestamp(i)
	return muo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableTimestamp(i *int64) *MessageUpdateOne {
	if i != nil {
		muo.SetTimestamp(*i)
	}
	return muo
}

// AddTimestamp adds i to the "timestamp" field.
func (muo *MessageUpdateOne) AddTimestamp(i int64) *MessageUpdateOne {
	muo.mutation.AddTimestamp(i)
	return muo
}

// SetSenderID sets the "sender_id" field.
func (muo *MessageUpdateOne) SetSenderID(i int64) *MessageUpdateOne {
	muo.mutation.ResetSenderID()
	muo.mutation.SetSenderID(i)
	return muo
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableSenderID(i *int64) *MessageUpdateOne {
	if i != nil {
		muo.SetSenderID(*i)
	}
	return muo
}

// AddSenderID adds i to the "sender_id" field.
func (muo *MessageUpdateOne) AddSenderID(i int64) *MessageUpdateOne {
	muo.mutation.AddSenderID(i)
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(t types.Message) *MessageUpdateOne {
	muo.mutation.SetContent(t)
	return muo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableContent(t *types.Message) *MessageUpdateOne {
	if t != nil {
		muo.SetContent(*t)
	}
	return muo
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.ChatID(); ok {
		_spec.SetField(message.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedChatID(); ok {
		_spec.AddField(message.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.MsgType(); ok {
		_spec.SetField(message.FieldMsgType, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedMsgType(); ok {
		_spec.AddField(message.FieldMsgType, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.MsgID(); ok {
		_spec.SetField(message.FieldMsgID, field.TypeString, value)
	}
	if value, ok := muo.mutation.Timestamp(); ok {
		_spec.SetField(message.FieldTimestamp, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedTimestamp(); ok {
		_spec.AddField(message.FieldTimestamp, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.SenderID(); ok {
		_spec.SetField(message.FieldSenderID, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedSenderID(); ok {
		_spec.AddField(message.FieldSenderID, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeJSON, value)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
