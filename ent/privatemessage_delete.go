// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"imlogic/ent/predicate"
	"imlogic/ent/privatemessage"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrivateMessageDelete is the builder for deleting a PrivateMessage entity.
type PrivateMessageDelete struct {
	config
	hooks    []Hook
	mutation *PrivateMessageMutation
}

// Where appends a list predicates to the PrivateMessageDelete builder.
func (pmd *PrivateMessageDelete) Where(ps ...predicate.PrivateMessage) *PrivateMessageDelete {
	pmd.mutation.Where(ps...)
	return pmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pmd *PrivateMessageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pmd.sqlExec, pmd.mutation, pmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pmd *PrivateMessageDelete) ExecX(ctx context.Context) int {
	n, err := pmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pmd *PrivateMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(privatemessage.Table, sqlgraph.NewFieldSpec(privatemessage.FieldID, field.TypeInt))
	if ps := pmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pmd.mutation.done = true
	return affected, err
}

// PrivateMessageDeleteOne is the builder for deleting a single PrivateMessage entity.
type PrivateMessageDeleteOne struct {
	pmd *PrivateMessageDelete
}

// Where appends a list predicates to the PrivateMessageDelete builder.
func (pmdo *PrivateMessageDeleteOne) Where(ps ...predicate.PrivateMessage) *PrivateMessageDeleteOne {
	pmdo.pmd.mutation.Where(ps...)
	return pmdo
}

// Exec executes the deletion query.
func (pmdo *PrivateMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := pmdo.pmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{privatemessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pmdo *PrivateMessageDeleteOne) ExecX(ctx context.Context) {
	if err := pmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
