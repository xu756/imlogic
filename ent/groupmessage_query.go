// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"imlogic/ent/groupmessage"
	"imlogic/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupMessageQuery is the builder for querying GroupMessage entities.
type GroupMessageQuery struct {
	config
	ctx        *QueryContext
	order      []groupmessage.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupMessage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupMessageQuery builder.
func (gmq *GroupMessageQuery) Where(ps ...predicate.GroupMessage) *GroupMessageQuery {
	gmq.predicates = append(gmq.predicates, ps...)
	return gmq
}

// Limit the number of records to be returned by this query.
func (gmq *GroupMessageQuery) Limit(limit int) *GroupMessageQuery {
	gmq.ctx.Limit = &limit
	return gmq
}

// Offset to start from.
func (gmq *GroupMessageQuery) Offset(offset int) *GroupMessageQuery {
	gmq.ctx.Offset = &offset
	return gmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gmq *GroupMessageQuery) Unique(unique bool) *GroupMessageQuery {
	gmq.ctx.Unique = &unique
	return gmq
}

// Order specifies how the records should be ordered.
func (gmq *GroupMessageQuery) Order(o ...groupmessage.OrderOption) *GroupMessageQuery {
	gmq.order = append(gmq.order, o...)
	return gmq
}

// First returns the first GroupMessage entity from the query.
// Returns a *NotFoundError when no GroupMessage was found.
func (gmq *GroupMessageQuery) First(ctx context.Context) (*GroupMessage, error) {
	nodes, err := gmq.Limit(1).All(setContextOp(ctx, gmq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gmq *GroupMessageQuery) FirstX(ctx context.Context) *GroupMessage {
	node, err := gmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupMessage ID from the query.
// Returns a *NotFoundError when no GroupMessage ID was found.
func (gmq *GroupMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gmq.Limit(1).IDs(setContextOp(ctx, gmq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gmq *GroupMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := gmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupMessage entity is found.
// Returns a *NotFoundError when no GroupMessage entities are found.
func (gmq *GroupMessageQuery) Only(ctx context.Context) (*GroupMessage, error) {
	nodes, err := gmq.Limit(2).All(setContextOp(ctx, gmq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupmessage.Label}
	default:
		return nil, &NotSingularError{groupmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gmq *GroupMessageQuery) OnlyX(ctx context.Context) *GroupMessage {
	node, err := gmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupMessage ID in the query.
// Returns a *NotSingularError when more than one GroupMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (gmq *GroupMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gmq.Limit(2).IDs(setContextOp(ctx, gmq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupmessage.Label}
	default:
		err = &NotSingularError{groupmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gmq *GroupMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := gmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupMessages.
func (gmq *GroupMessageQuery) All(ctx context.Context) ([]*GroupMessage, error) {
	ctx = setContextOp(ctx, gmq.ctx, ent.OpQueryAll)
	if err := gmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupMessage, *GroupMessageQuery]()
	return withInterceptors[[]*GroupMessage](ctx, gmq, qr, gmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gmq *GroupMessageQuery) AllX(ctx context.Context) []*GroupMessage {
	nodes, err := gmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupMessage IDs.
func (gmq *GroupMessageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gmq.ctx.Unique == nil && gmq.path != nil {
		gmq.Unique(true)
	}
	ctx = setContextOp(ctx, gmq.ctx, ent.OpQueryIDs)
	if err = gmq.Select(groupmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gmq *GroupMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := gmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gmq *GroupMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gmq.ctx, ent.OpQueryCount)
	if err := gmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gmq, querierCount[*GroupMessageQuery](), gmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gmq *GroupMessageQuery) CountX(ctx context.Context) int {
	count, err := gmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gmq *GroupMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gmq.ctx, ent.OpQueryExist)
	switch _, err := gmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gmq *GroupMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := gmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gmq *GroupMessageQuery) Clone() *GroupMessageQuery {
	if gmq == nil {
		return nil
	}
	return &GroupMessageQuery{
		config:     gmq.config,
		ctx:        gmq.ctx.Clone(),
		order:      append([]groupmessage.OrderOption{}, gmq.order...),
		inters:     append([]Interceptor{}, gmq.inters...),
		predicates: append([]predicate.GroupMessage{}, gmq.predicates...),
		// clone intermediate query.
		sql:  gmq.sql.Clone(),
		path: gmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MsgID string `json:"msg_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupMessage.Query().
//		GroupBy(groupmessage.FieldMsgID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gmq *GroupMessageQuery) GroupBy(field string, fields ...string) *GroupMessageGroupBy {
	gmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupMessageGroupBy{build: gmq}
	grbuild.flds = &gmq.ctx.Fields
	grbuild.label = groupmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MsgID string `json:"msg_id,omitempty"`
//	}
//
//	client.GroupMessage.Query().
//		Select(groupmessage.FieldMsgID).
//		Scan(ctx, &v)
func (gmq *GroupMessageQuery) Select(fields ...string) *GroupMessageSelect {
	gmq.ctx.Fields = append(gmq.ctx.Fields, fields...)
	sbuild := &GroupMessageSelect{GroupMessageQuery: gmq}
	sbuild.label = groupmessage.Label
	sbuild.flds, sbuild.scan = &gmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupMessageSelect configured with the given aggregations.
func (gmq *GroupMessageQuery) Aggregate(fns ...AggregateFunc) *GroupMessageSelect {
	return gmq.Select().Aggregate(fns...)
}

func (gmq *GroupMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gmq); err != nil {
				return err
			}
		}
	}
	for _, f := range gmq.ctx.Fields {
		if !groupmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gmq.path != nil {
		prev, err := gmq.path(ctx)
		if err != nil {
			return err
		}
		gmq.sql = prev
	}
	return nil
}

func (gmq *GroupMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupMessage, error) {
	var (
		nodes = []*GroupMessage{}
		_spec = gmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupMessage{config: gmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gmq *GroupMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gmq.querySpec()
	_spec.Node.Columns = gmq.ctx.Fields
	if len(gmq.ctx.Fields) > 0 {
		_spec.Unique = gmq.ctx.Unique != nil && *gmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gmq.driver, _spec)
}

func (gmq *GroupMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupmessage.Table, groupmessage.Columns, sqlgraph.NewFieldSpec(groupmessage.FieldID, field.TypeInt))
	_spec.From = gmq.sql
	if unique := gmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gmq.path != nil {
		_spec.Unique = true
	}
	if fields := gmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmessage.FieldID)
		for i := range fields {
			if fields[i] != groupmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gmq *GroupMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gmq.driver.Dialect())
	t1 := builder.Table(groupmessage.Table)
	columns := gmq.ctx.Fields
	if len(columns) == 0 {
		columns = groupmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gmq.sql != nil {
		selector = gmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gmq.ctx.Unique != nil && *gmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gmq.predicates {
		p(selector)
	}
	for _, p := range gmq.order {
		p(selector)
	}
	if offset := gmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupMessageGroupBy is the group-by builder for GroupMessage entities.
type GroupMessageGroupBy struct {
	selector
	build *GroupMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gmgb *GroupMessageGroupBy) Aggregate(fns ...AggregateFunc) *GroupMessageGroupBy {
	gmgb.fns = append(gmgb.fns, fns...)
	return gmgb
}

// Scan applies the selector query and scans the result into the given value.
func (gmgb *GroupMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gmgb.build.ctx, ent.OpQueryGroupBy)
	if err := gmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupMessageQuery, *GroupMessageGroupBy](ctx, gmgb.build, gmgb, gmgb.build.inters, v)
}

func (gmgb *GroupMessageGroupBy) sqlScan(ctx context.Context, root *GroupMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gmgb.fns))
	for _, fn := range gmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gmgb.flds)+len(gmgb.fns))
		for _, f := range *gmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupMessageSelect is the builder for selecting fields of GroupMessage entities.
type GroupMessageSelect struct {
	*GroupMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gms *GroupMessageSelect) Aggregate(fns ...AggregateFunc) *GroupMessageSelect {
	gms.fns = append(gms.fns, fns...)
	return gms
}

// Scan applies the selector query and scans the result into the given value.
func (gms *GroupMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gms.ctx, ent.OpQuerySelect)
	if err := gms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupMessageQuery, *GroupMessageSelect](ctx, gms.GroupMessageQuery, gms, gms.inters, v)
}

func (gms *GroupMessageSelect) sqlScan(ctx context.Context, root *GroupMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gms.fns))
	for _, fn := range gms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
