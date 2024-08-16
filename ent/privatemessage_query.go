// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"imlogic/ent/predicate"
	"imlogic/ent/privatemessage"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrivateMessageQuery is the builder for querying PrivateMessage entities.
type PrivateMessageQuery struct {
	config
	ctx        *QueryContext
	order      []privatemessage.OrderOption
	inters     []Interceptor
	predicates []predicate.PrivateMessage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PrivateMessageQuery builder.
func (pmq *PrivateMessageQuery) Where(ps ...predicate.PrivateMessage) *PrivateMessageQuery {
	pmq.predicates = append(pmq.predicates, ps...)
	return pmq
}

// Limit the number of records to be returned by this query.
func (pmq *PrivateMessageQuery) Limit(limit int) *PrivateMessageQuery {
	pmq.ctx.Limit = &limit
	return pmq
}

// Offset to start from.
func (pmq *PrivateMessageQuery) Offset(offset int) *PrivateMessageQuery {
	pmq.ctx.Offset = &offset
	return pmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pmq *PrivateMessageQuery) Unique(unique bool) *PrivateMessageQuery {
	pmq.ctx.Unique = &unique
	return pmq
}

// Order specifies how the records should be ordered.
func (pmq *PrivateMessageQuery) Order(o ...privatemessage.OrderOption) *PrivateMessageQuery {
	pmq.order = append(pmq.order, o...)
	return pmq
}

// First returns the first PrivateMessage entity from the query.
// Returns a *NotFoundError when no PrivateMessage was found.
func (pmq *PrivateMessageQuery) First(ctx context.Context) (*PrivateMessage, error) {
	nodes, err := pmq.Limit(1).All(setContextOp(ctx, pmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{privatemessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pmq *PrivateMessageQuery) FirstX(ctx context.Context) *PrivateMessage {
	node, err := pmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PrivateMessage ID from the query.
// Returns a *NotFoundError when no PrivateMessage ID was found.
func (pmq *PrivateMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmq.Limit(1).IDs(setContextOp(ctx, pmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{privatemessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pmq *PrivateMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := pmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PrivateMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PrivateMessage entity is found.
// Returns a *NotFoundError when no PrivateMessage entities are found.
func (pmq *PrivateMessageQuery) Only(ctx context.Context) (*PrivateMessage, error) {
	nodes, err := pmq.Limit(2).All(setContextOp(ctx, pmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{privatemessage.Label}
	default:
		return nil, &NotSingularError{privatemessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pmq *PrivateMessageQuery) OnlyX(ctx context.Context) *PrivateMessage {
	node, err := pmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PrivateMessage ID in the query.
// Returns a *NotSingularError when more than one PrivateMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (pmq *PrivateMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmq.Limit(2).IDs(setContextOp(ctx, pmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{privatemessage.Label}
	default:
		err = &NotSingularError{privatemessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pmq *PrivateMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := pmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PrivateMessages.
func (pmq *PrivateMessageQuery) All(ctx context.Context) ([]*PrivateMessage, error) {
	ctx = setContextOp(ctx, pmq.ctx, "All")
	if err := pmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PrivateMessage, *PrivateMessageQuery]()
	return withInterceptors[[]*PrivateMessage](ctx, pmq, qr, pmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pmq *PrivateMessageQuery) AllX(ctx context.Context) []*PrivateMessage {
	nodes, err := pmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PrivateMessage IDs.
func (pmq *PrivateMessageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pmq.ctx.Unique == nil && pmq.path != nil {
		pmq.Unique(true)
	}
	ctx = setContextOp(ctx, pmq.ctx, "IDs")
	if err = pmq.Select(privatemessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pmq *PrivateMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := pmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pmq *PrivateMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pmq.ctx, "Count")
	if err := pmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pmq, querierCount[*PrivateMessageQuery](), pmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pmq *PrivateMessageQuery) CountX(ctx context.Context) int {
	count, err := pmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pmq *PrivateMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pmq.ctx, "Exist")
	switch _, err := pmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pmq *PrivateMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := pmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PrivateMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pmq *PrivateMessageQuery) Clone() *PrivateMessageQuery {
	if pmq == nil {
		return nil
	}
	return &PrivateMessageQuery{
		config:     pmq.config,
		ctx:        pmq.ctx.Clone(),
		order:      append([]privatemessage.OrderOption{}, pmq.order...),
		inters:     append([]Interceptor{}, pmq.inters...),
		predicates: append([]predicate.PrivateMessage{}, pmq.predicates...),
		// clone intermediate query.
		sql:  pmq.sql.Clone(),
		path: pmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MsgType int32 `json:"msg_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PrivateMessage.Query().
//		GroupBy(privatemessage.FieldMsgType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pmq *PrivateMessageQuery) GroupBy(field string, fields ...string) *PrivateMessageGroupBy {
	pmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PrivateMessageGroupBy{build: pmq}
	grbuild.flds = &pmq.ctx.Fields
	grbuild.label = privatemessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MsgType int32 `json:"msg_type,omitempty"`
//	}
//
//	client.PrivateMessage.Query().
//		Select(privatemessage.FieldMsgType).
//		Scan(ctx, &v)
func (pmq *PrivateMessageQuery) Select(fields ...string) *PrivateMessageSelect {
	pmq.ctx.Fields = append(pmq.ctx.Fields, fields...)
	sbuild := &PrivateMessageSelect{PrivateMessageQuery: pmq}
	sbuild.label = privatemessage.Label
	sbuild.flds, sbuild.scan = &pmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PrivateMessageSelect configured with the given aggregations.
func (pmq *PrivateMessageQuery) Aggregate(fns ...AggregateFunc) *PrivateMessageSelect {
	return pmq.Select().Aggregate(fns...)
}

func (pmq *PrivateMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pmq); err != nil {
				return err
			}
		}
	}
	for _, f := range pmq.ctx.Fields {
		if !privatemessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pmq.path != nil {
		prev, err := pmq.path(ctx)
		if err != nil {
			return err
		}
		pmq.sql = prev
	}
	return nil
}

func (pmq *PrivateMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PrivateMessage, error) {
	var (
		nodes = []*PrivateMessage{}
		_spec = pmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PrivateMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PrivateMessage{config: pmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pmq *PrivateMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pmq.querySpec()
	_spec.Node.Columns = pmq.ctx.Fields
	if len(pmq.ctx.Fields) > 0 {
		_spec.Unique = pmq.ctx.Unique != nil && *pmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pmq.driver, _spec)
}

func (pmq *PrivateMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(privatemessage.Table, privatemessage.Columns, sqlgraph.NewFieldSpec(privatemessage.FieldID, field.TypeInt))
	_spec.From = pmq.sql
	if unique := pmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pmq.path != nil {
		_spec.Unique = true
	}
	if fields := pmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, privatemessage.FieldID)
		for i := range fields {
			if fields[i] != privatemessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pmq *PrivateMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pmq.driver.Dialect())
	t1 := builder.Table(privatemessage.Table)
	columns := pmq.ctx.Fields
	if len(columns) == 0 {
		columns = privatemessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pmq.sql != nil {
		selector = pmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pmq.ctx.Unique != nil && *pmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pmq.predicates {
		p(selector)
	}
	for _, p := range pmq.order {
		p(selector)
	}
	if offset := pmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PrivateMessageGroupBy is the group-by builder for PrivateMessage entities.
type PrivateMessageGroupBy struct {
	selector
	build *PrivateMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pmgb *PrivateMessageGroupBy) Aggregate(fns ...AggregateFunc) *PrivateMessageGroupBy {
	pmgb.fns = append(pmgb.fns, fns...)
	return pmgb
}

// Scan applies the selector query and scans the result into the given value.
func (pmgb *PrivateMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pmgb.build.ctx, "GroupBy")
	if err := pmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PrivateMessageQuery, *PrivateMessageGroupBy](ctx, pmgb.build, pmgb, pmgb.build.inters, v)
}

func (pmgb *PrivateMessageGroupBy) sqlScan(ctx context.Context, root *PrivateMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pmgb.fns))
	for _, fn := range pmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pmgb.flds)+len(pmgb.fns))
		for _, f := range *pmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PrivateMessageSelect is the builder for selecting fields of PrivateMessage entities.
type PrivateMessageSelect struct {
	*PrivateMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pms *PrivateMessageSelect) Aggregate(fns ...AggregateFunc) *PrivateMessageSelect {
	pms.fns = append(pms.fns, fns...)
	return pms
}

// Scan applies the selector query and scans the result into the given value.
func (pms *PrivateMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pms.ctx, "Select")
	if err := pms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PrivateMessageQuery, *PrivateMessageSelect](ctx, pms.PrivateMessageQuery, pms, pms.inters, v)
}

func (pms *PrivateMessageSelect) sqlScan(ctx context.Context, root *PrivateMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pms.fns))
	for _, fn := range pms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
