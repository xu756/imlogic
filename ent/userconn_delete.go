// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"imlogic/ent/predicate"
	"imlogic/ent/userconn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserConnDelete is the builder for deleting a UserConn entity.
type UserConnDelete struct {
	config
	hooks    []Hook
	mutation *UserConnMutation
}

// Where appends a list predicates to the UserConnDelete builder.
func (ucd *UserConnDelete) Where(ps ...predicate.UserConn) *UserConnDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserConnDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ucd.sqlExec, ucd.mutation, ucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserConnDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserConnDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userconn.Table, sqlgraph.NewFieldSpec(userconn.FieldID, field.TypeInt))
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucd.mutation.done = true
	return affected, err
}

// UserConnDeleteOne is the builder for deleting a single UserConn entity.
type UserConnDeleteOne struct {
	ucd *UserConnDelete
}

// Where appends a list predicates to the UserConnDelete builder.
func (ucdo *UserConnDeleteOne) Where(ps ...predicate.UserConn) *UserConnDeleteOne {
	ucdo.ucd.mutation.Where(ps...)
	return ucdo
}

// Exec executes the deletion query.
func (ucdo *UserConnDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userconn.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserConnDeleteOne) ExecX(ctx context.Context) {
	if err := ucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
