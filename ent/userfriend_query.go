// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"imlogic/ent/predicate"
	"imlogic/ent/userfriend"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFriendQuery is the builder for querying UserFriend entities.
type UserFriendQuery struct {
	config
	ctx        *QueryContext
	order      []userfriend.OrderOption
	inters     []Interceptor
	predicates []predicate.UserFriend
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserFriendQuery builder.
func (ufq *UserFriendQuery) Where(ps ...predicate.UserFriend) *UserFriendQuery {
	ufq.predicates = append(ufq.predicates, ps...)
	return ufq
}

// Limit the number of records to be returned by this query.
func (ufq *UserFriendQuery) Limit(limit int) *UserFriendQuery {
	ufq.ctx.Limit = &limit
	return ufq
}

// Offset to start from.
func (ufq *UserFriendQuery) Offset(offset int) *UserFriendQuery {
	ufq.ctx.Offset = &offset
	return ufq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufq *UserFriendQuery) Unique(unique bool) *UserFriendQuery {
	ufq.ctx.Unique = &unique
	return ufq
}

// Order specifies how the records should be ordered.
func (ufq *UserFriendQuery) Order(o ...userfriend.OrderOption) *UserFriendQuery {
	ufq.order = append(ufq.order, o...)
	return ufq
}

// First returns the first UserFriend entity from the query.
// Returns a *NotFoundError when no UserFriend was found.
func (ufq *UserFriendQuery) First(ctx context.Context) (*UserFriend, error) {
	nodes, err := ufq.Limit(1).All(setContextOp(ctx, ufq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userfriend.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufq *UserFriendQuery) FirstX(ctx context.Context) *UserFriend {
	node, err := ufq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserFriend ID from the query.
// Returns a *NotFoundError when no UserFriend ID was found.
func (ufq *UserFriendQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ufq.Limit(1).IDs(setContextOp(ctx, ufq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userfriend.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufq *UserFriendQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ufq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserFriend entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserFriend entity is found.
// Returns a *NotFoundError when no UserFriend entities are found.
func (ufq *UserFriendQuery) Only(ctx context.Context) (*UserFriend, error) {
	nodes, err := ufq.Limit(2).All(setContextOp(ctx, ufq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userfriend.Label}
	default:
		return nil, &NotSingularError{userfriend.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufq *UserFriendQuery) OnlyX(ctx context.Context) *UserFriend {
	node, err := ufq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserFriend ID in the query.
// Returns a *NotSingularError when more than one UserFriend ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufq *UserFriendQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ufq.Limit(2).IDs(setContextOp(ctx, ufq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userfriend.Label}
	default:
		err = &NotSingularError{userfriend.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufq *UserFriendQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ufq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserFriends.
func (ufq *UserFriendQuery) All(ctx context.Context) ([]*UserFriend, error) {
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryAll)
	if err := ufq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserFriend, *UserFriendQuery]()
	return withInterceptors[[]*UserFriend](ctx, ufq, qr, ufq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufq *UserFriendQuery) AllX(ctx context.Context) []*UserFriend {
	nodes, err := ufq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserFriend IDs.
func (ufq *UserFriendQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ufq.ctx.Unique == nil && ufq.path != nil {
		ufq.Unique(true)
	}
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryIDs)
	if err = ufq.Select(userfriend.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufq *UserFriendQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ufq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufq *UserFriendQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryCount)
	if err := ufq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufq, querierCount[*UserFriendQuery](), ufq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufq *UserFriendQuery) CountX(ctx context.Context) int {
	count, err := ufq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufq *UserFriendQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryExist)
	switch _, err := ufq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufq *UserFriendQuery) ExistX(ctx context.Context) bool {
	exist, err := ufq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserFriendQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufq *UserFriendQuery) Clone() *UserFriendQuery {
	if ufq == nil {
		return nil
	}
	return &UserFriendQuery{
		config:     ufq.config,
		ctx:        ufq.ctx.Clone(),
		order:      append([]userfriend.OrderOption{}, ufq.order...),
		inters:     append([]Interceptor{}, ufq.inters...),
		predicates: append([]predicate.UserFriend{}, ufq.predicates...),
		// clone intermediate query.
		sql:  ufq.sql.Clone(),
		path: ufq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserFriend.Query().
//		GroupBy(userfriend.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufq *UserFriendQuery) GroupBy(field string, fields ...string) *UserFriendGroupBy {
	ufq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserFriendGroupBy{build: ufq}
	grbuild.flds = &ufq.ctx.Fields
	grbuild.label = userfriend.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserFriend.Query().
//		Select(userfriend.FieldCreatedAt).
//		Scan(ctx, &v)
func (ufq *UserFriendQuery) Select(fields ...string) *UserFriendSelect {
	ufq.ctx.Fields = append(ufq.ctx.Fields, fields...)
	sbuild := &UserFriendSelect{UserFriendQuery: ufq}
	sbuild.label = userfriend.Label
	sbuild.flds, sbuild.scan = &ufq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserFriendSelect configured with the given aggregations.
func (ufq *UserFriendQuery) Aggregate(fns ...AggregateFunc) *UserFriendSelect {
	return ufq.Select().Aggregate(fns...)
}

func (ufq *UserFriendQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufq.ctx.Fields {
		if !userfriend.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufq.path != nil {
		prev, err := ufq.path(ctx)
		if err != nil {
			return err
		}
		ufq.sql = prev
	}
	return nil
}

func (ufq *UserFriendQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserFriend, error) {
	var (
		nodes = []*UserFriend{}
		_spec = ufq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserFriend).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserFriend{config: ufq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ufq *UserFriendQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufq.querySpec()
	_spec.Node.Columns = ufq.ctx.Fields
	if len(ufq.ctx.Fields) > 0 {
		_spec.Unique = ufq.ctx.Unique != nil && *ufq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ufq.driver, _spec)
}

func (ufq *UserFriendQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userfriend.Table, userfriend.Columns, sqlgraph.NewFieldSpec(userfriend.FieldID, field.TypeInt64))
	_spec.From = ufq.sql
	if unique := ufq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufq.path != nil {
		_spec.Unique = true
	}
	if fields := ufq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfriend.FieldID)
		for i := range fields {
			if fields[i] != userfriend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ufq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufq *UserFriendQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufq.driver.Dialect())
	t1 := builder.Table(userfriend.Table)
	columns := ufq.ctx.Fields
	if len(columns) == 0 {
		columns = userfriend.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufq.sql != nil {
		selector = ufq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufq.ctx.Unique != nil && *ufq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ufq.predicates {
		p(selector)
	}
	for _, p := range ufq.order {
		p(selector)
	}
	if offset := ufq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserFriendGroupBy is the group-by builder for UserFriend entities.
type UserFriendGroupBy struct {
	selector
	build *UserFriendQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufgb *UserFriendGroupBy) Aggregate(fns ...AggregateFunc) *UserFriendGroupBy {
	ufgb.fns = append(ufgb.fns, fns...)
	return ufgb
}

// Scan applies the selector query and scans the result into the given value.
func (ufgb *UserFriendGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufgb.build.ctx, ent.OpQueryGroupBy)
	if err := ufgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserFriendQuery, *UserFriendGroupBy](ctx, ufgb.build, ufgb, ufgb.build.inters, v)
}

func (ufgb *UserFriendGroupBy) sqlScan(ctx context.Context, root *UserFriendQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufgb.fns))
	for _, fn := range ufgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufgb.flds)+len(ufgb.fns))
		for _, f := range *ufgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserFriendSelect is the builder for selecting fields of UserFriend entities.
type UserFriendSelect struct {
	*UserFriendQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufs *UserFriendSelect) Aggregate(fns ...AggregateFunc) *UserFriendSelect {
	ufs.fns = append(ufs.fns, fns...)
	return ufs
}

// Scan applies the selector query and scans the result into the given value.
func (ufs *UserFriendSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufs.ctx, ent.OpQuerySelect)
	if err := ufs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserFriendQuery, *UserFriendSelect](ctx, ufs.UserFriendQuery, ufs, ufs.inters, v)
}

func (ufs *UserFriendSelect) sqlScan(ctx context.Context, root *UserFriendQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufs.fns))
	for _, fn := range ufs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
