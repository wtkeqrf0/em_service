// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/restService/pkg/ent/enrichedfio"
	"github.com/wtkeqrf0/restService/pkg/ent/predicate"
)

// EnrichedFioQuery is the builder for querying EnrichedFio entities.
type EnrichedFioQuery struct {
	config
	ctx        *QueryContext
	order      []enrichedfio.OrderOption
	inters     []Interceptor
	predicates []predicate.EnrichedFio
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnrichedFioQuery builder.
func (efq *EnrichedFioQuery) Where(ps ...predicate.EnrichedFio) *EnrichedFioQuery {
	efq.predicates = append(efq.predicates, ps...)
	return efq
}

// Limit the number of records to be returned by this query.
func (efq *EnrichedFioQuery) Limit(limit int) *EnrichedFioQuery {
	efq.ctx.Limit = &limit
	return efq
}

// Offset to start from.
func (efq *EnrichedFioQuery) Offset(offset int) *EnrichedFioQuery {
	efq.ctx.Offset = &offset
	return efq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (efq *EnrichedFioQuery) Unique(unique bool) *EnrichedFioQuery {
	efq.ctx.Unique = &unique
	return efq
}

// Order specifies how the records should be ordered.
func (efq *EnrichedFioQuery) Order(o ...enrichedfio.OrderOption) *EnrichedFioQuery {
	efq.order = append(efq.order, o...)
	return efq
}

// First returns the first EnrichedFio entity from the query.
// Returns a *NotFoundError when no EnrichedFio was found.
func (efq *EnrichedFioQuery) First(ctx context.Context) (*EnrichedFio, error) {
	nodes, err := efq.Limit(1).All(setContextOp(ctx, efq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enrichedfio.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (efq *EnrichedFioQuery) FirstX(ctx context.Context) *EnrichedFio {
	node, err := efq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EnrichedFio ID from the query.
// Returns a *NotFoundError when no EnrichedFio ID was found.
func (efq *EnrichedFioQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = efq.Limit(1).IDs(setContextOp(ctx, efq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enrichedfio.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (efq *EnrichedFioQuery) FirstIDX(ctx context.Context) int {
	id, err := efq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EnrichedFio entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EnrichedFio entity is found.
// Returns a *NotFoundError when no EnrichedFio entities are found.
func (efq *EnrichedFioQuery) Only(ctx context.Context) (*EnrichedFio, error) {
	nodes, err := efq.Limit(2).All(setContextOp(ctx, efq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enrichedfio.Label}
	default:
		return nil, &NotSingularError{enrichedfio.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (efq *EnrichedFioQuery) OnlyX(ctx context.Context) *EnrichedFio {
	node, err := efq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EnrichedFio ID in the query.
// Returns a *NotSingularError when more than one EnrichedFio ID is found.
// Returns a *NotFoundError when no entities are found.
func (efq *EnrichedFioQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = efq.Limit(2).IDs(setContextOp(ctx, efq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enrichedfio.Label}
	default:
		err = &NotSingularError{enrichedfio.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (efq *EnrichedFioQuery) OnlyIDX(ctx context.Context) int {
	id, err := efq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EnrichedFios.
func (efq *EnrichedFioQuery) All(ctx context.Context) ([]*EnrichedFio, error) {
	ctx = setContextOp(ctx, efq.ctx, "All")
	if err := efq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EnrichedFio, *EnrichedFioQuery]()
	return withInterceptors[[]*EnrichedFio](ctx, efq, qr, efq.inters)
}

// AllX is like All, but panics if an error occurs.
func (efq *EnrichedFioQuery) AllX(ctx context.Context) []*EnrichedFio {
	nodes, err := efq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EnrichedFio IDs.
func (efq *EnrichedFioQuery) IDs(ctx context.Context) (ids []int, err error) {
	if efq.ctx.Unique == nil && efq.path != nil {
		efq.Unique(true)
	}
	ctx = setContextOp(ctx, efq.ctx, "IDs")
	if err = efq.Select(enrichedfio.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (efq *EnrichedFioQuery) IDsX(ctx context.Context) []int {
	ids, err := efq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (efq *EnrichedFioQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, efq.ctx, "Count")
	if err := efq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, efq, querierCount[*EnrichedFioQuery](), efq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (efq *EnrichedFioQuery) CountX(ctx context.Context) int {
	count, err := efq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (efq *EnrichedFioQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, efq.ctx, "Exist")
	switch _, err := efq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (efq *EnrichedFioQuery) ExistX(ctx context.Context) bool {
	exist, err := efq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnrichedFioQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (efq *EnrichedFioQuery) Clone() *EnrichedFioQuery {
	if efq == nil {
		return nil
	}
	return &EnrichedFioQuery{
		config:     efq.config,
		ctx:        efq.ctx.Clone(),
		order:      append([]enrichedfio.OrderOption{}, efq.order...),
		inters:     append([]Interceptor{}, efq.inters...),
		predicates: append([]predicate.EnrichedFio{}, efq.predicates...),
		// clone intermediate query.
		sql:  efq.sql.Clone(),
		path: efq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EnrichedFio.Query().
//		GroupBy(enrichedfio.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (efq *EnrichedFioQuery) GroupBy(field string, fields ...string) *EnrichedFioGroupBy {
	efq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EnrichedFioGroupBy{build: efq}
	grbuild.flds = &efq.ctx.Fields
	grbuild.label = enrichedfio.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.EnrichedFio.Query().
//		Select(enrichedfio.FieldCreateTime).
//		Scan(ctx, &v)
func (efq *EnrichedFioQuery) Select(fields ...string) *EnrichedFioSelect {
	efq.ctx.Fields = append(efq.ctx.Fields, fields...)
	sbuild := &EnrichedFioSelect{EnrichedFioQuery: efq}
	sbuild.label = enrichedfio.Label
	sbuild.flds, sbuild.scan = &efq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EnrichedFioSelect configured with the given aggregations.
func (efq *EnrichedFioQuery) Aggregate(fns ...AggregateFunc) *EnrichedFioSelect {
	return efq.Select().Aggregate(fns...)
}

func (efq *EnrichedFioQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range efq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, efq); err != nil {
				return err
			}
		}
	}
	for _, f := range efq.ctx.Fields {
		if !enrichedfio.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if efq.path != nil {
		prev, err := efq.path(ctx)
		if err != nil {
			return err
		}
		efq.sql = prev
	}
	return nil
}

func (efq *EnrichedFioQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EnrichedFio, error) {
	var (
		nodes = []*EnrichedFio{}
		_spec = efq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EnrichedFio).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EnrichedFio{config: efq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, efq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (efq *EnrichedFioQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := efq.querySpec()
	_spec.Node.Columns = efq.ctx.Fields
	if len(efq.ctx.Fields) > 0 {
		_spec.Unique = efq.ctx.Unique != nil && *efq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, efq.driver, _spec)
}

func (efq *EnrichedFioQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(enrichedfio.Table, enrichedfio.Columns, sqlgraph.NewFieldSpec(enrichedfio.FieldID, field.TypeInt))
	_spec.From = efq.sql
	if unique := efq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if efq.path != nil {
		_spec.Unique = true
	}
	if fields := efq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enrichedfio.FieldID)
		for i := range fields {
			if fields[i] != enrichedfio.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := efq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := efq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := efq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := efq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (efq *EnrichedFioQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(efq.driver.Dialect())
	t1 := builder.Table(enrichedfio.Table)
	columns := efq.ctx.Fields
	if len(columns) == 0 {
		columns = enrichedfio.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if efq.sql != nil {
		selector = efq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if efq.ctx.Unique != nil && *efq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range efq.predicates {
		p(selector)
	}
	for _, p := range efq.order {
		p(selector)
	}
	if offset := efq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := efq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnrichedFioGroupBy is the group-by builder for EnrichedFio entities.
type EnrichedFioGroupBy struct {
	selector
	build *EnrichedFioQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (efgb *EnrichedFioGroupBy) Aggregate(fns ...AggregateFunc) *EnrichedFioGroupBy {
	efgb.fns = append(efgb.fns, fns...)
	return efgb
}

// Scan applies the selector query and scans the result into the given value.
func (efgb *EnrichedFioGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, efgb.build.ctx, "GroupBy")
	if err := efgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnrichedFioQuery, *EnrichedFioGroupBy](ctx, efgb.build, efgb, efgb.build.inters, v)
}

func (efgb *EnrichedFioGroupBy) sqlScan(ctx context.Context, root *EnrichedFioQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(efgb.fns))
	for _, fn := range efgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*efgb.flds)+len(efgb.fns))
		for _, f := range *efgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*efgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := efgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EnrichedFioSelect is the builder for selecting fields of EnrichedFio entities.
type EnrichedFioSelect struct {
	*EnrichedFioQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (efs *EnrichedFioSelect) Aggregate(fns ...AggregateFunc) *EnrichedFioSelect {
	efs.fns = append(efs.fns, fns...)
	return efs
}

// Scan applies the selector query and scans the result into the given value.
func (efs *EnrichedFioSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, efs.ctx, "Select")
	if err := efs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnrichedFioQuery, *EnrichedFioSelect](ctx, efs.EnrichedFioQuery, efs, efs.inters, v)
}

func (efs *EnrichedFioSelect) sqlScan(ctx context.Context, root *EnrichedFioQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(efs.fns))
	for _, fn := range efs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*efs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := efs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}