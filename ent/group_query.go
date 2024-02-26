// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xu756/imlogic/ent/group"
	"github.com/xu756/imlogic/ent/predicate"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
	config
	ctx        *QueryContext
	order      []group.OrderOption
	inters     []Interceptor
	predicates []predicate.Group
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupQuery builder.
func (gq *GroupQuery) Where(ps ...predicate.Group) *GroupQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GroupQuery) Limit(limit int) *GroupQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GroupQuery) Offset(offset int) *GroupQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GroupQuery) Unique(unique bool) *GroupQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GroupQuery) Order(o ...group.OrderOption) *GroupQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// First returns the first Group entity from the query.
// Returns a *NotFoundError when no Group was found.
func (gq *GroupQuery) First(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{group.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GroupQuery) FirstX(ctx context.Context) *Group {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Group ID from the query.
// Returns a *NotFoundError when no Group ID was found.
func (gq *GroupQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{group.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GroupQuery) FirstIDX(ctx context.Context) int64 {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Group entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Group entity is found.
// Returns a *NotFoundError when no Group entities are found.
func (gq *GroupQuery) Only(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{group.Label}
	default:
		return nil, &NotSingularError{group.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GroupQuery) OnlyX(ctx context.Context) *Group {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Group ID in the query.
// Returns a *NotSingularError when more than one Group ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GroupQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = &NotSingularError{group.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GroupQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (gq *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Group, *GroupQuery]()
	return withInterceptors[[]*Group](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GroupQuery) AllX(ctx context.Context) []*Group {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Group IDs.
func (gq *GroupQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(group.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GroupQuery) IDsX(ctx context.Context) []int64 {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GroupQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GroupQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GroupQuery) Clone() *GroupQuery {
	if gq == nil {
		return nil
	}
	return &GroupQuery{
		config:     gq.config,
		ctx:        gq.ctx.Clone(),
		order:      append([]group.OrderOption{}, gq.order...),
		inters:     append([]Interceptor{}, gq.inters...),
		predicates: append([]predicate.Group{}, gq.predicates...),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
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
//	client.Group.Query().
//		GroupBy(group.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GroupQuery) GroupBy(field string, fields ...string) *GroupGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = group.Label
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
//	client.Group.Query().
//		Select(group.FieldCreatedAt).
//		Scan(ctx, &v)
func (gq *GroupQuery) Select(fields ...string) *GroupSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GroupSelect{GroupQuery: gq}
	sbuild.label = group.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupSelect configured with the given aggregations.
func (gq *GroupQuery) Aggregate(fns ...AggregateFunc) *GroupSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !group.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Group, error) {
	var (
		nodes = []*Group{}
		_spec = gq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Group).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Group{config: gq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gq *GroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for i := range fields {
			if fields[i] != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(group.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = group.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupGroupBy is the group-by builder for Group entities.
type GroupGroupBy struct {
	selector
	build *GroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GroupGroupBy) Aggregate(fns ...AggregateFunc) *GroupGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupQuery, *GroupGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GroupGroupBy) sqlScan(ctx context.Context, root *GroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupSelect is the builder for selecting fields of Group entities.
type GroupSelect struct {
	*GroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GroupSelect) Aggregate(fns ...AggregateFunc) *GroupSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupQuery, *GroupSelect](ctx, gs.GroupQuery, gs, gs.inters, v)
}

func (gs *GroupSelect) sqlScan(ctx context.Context, root *GroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}