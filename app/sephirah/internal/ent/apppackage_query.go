// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
)

// AppPackageQuery is the builder for querying AppPackage entities.
type AppPackageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppPackage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppPackageQuery builder.
func (apq *AppPackageQuery) Where(ps ...predicate.AppPackage) *AppPackageQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit adds a limit step to the query.
func (apq *AppPackageQuery) Limit(limit int) *AppPackageQuery {
	apq.limit = &limit
	return apq
}

// Offset adds an offset step to the query.
func (apq *AppPackageQuery) Offset(offset int) *AppPackageQuery {
	apq.offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *AppPackageQuery) Unique(unique bool) *AppPackageQuery {
	apq.unique = &unique
	return apq
}

// Order adds an order step to the query.
func (apq *AppPackageQuery) Order(o ...OrderFunc) *AppPackageQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// First returns the first AppPackage entity from the query.
// Returns a *NotFoundError when no AppPackage was found.
func (apq *AppPackageQuery) First(ctx context.Context) (*AppPackage, error) {
	nodes, err := apq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apppackage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AppPackageQuery) FirstX(ctx context.Context) *AppPackage {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppPackage ID from the query.
// Returns a *NotFoundError when no AppPackage ID was found.
func (apq *AppPackageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apppackage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AppPackageQuery) FirstIDX(ctx context.Context) int {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppPackage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppPackage entity is found.
// Returns a *NotFoundError when no AppPackage entities are found.
func (apq *AppPackageQuery) Only(ctx context.Context) (*AppPackage, error) {
	nodes, err := apq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apppackage.Label}
	default:
		return nil, &NotSingularError{apppackage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AppPackageQuery) OnlyX(ctx context.Context) *AppPackage {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppPackage ID in the query.
// Returns a *NotSingularError when more than one AppPackage ID is found.
// Returns a *NotFoundError when no entities are found.
func (apq *AppPackageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apppackage.Label}
	default:
		err = &NotSingularError{apppackage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AppPackageQuery) OnlyIDX(ctx context.Context) int {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppPackages.
func (apq *AppPackageQuery) All(ctx context.Context) ([]*AppPackage, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return apq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (apq *AppPackageQuery) AllX(ctx context.Context) []*AppPackage {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppPackage IDs.
func (apq *AppPackageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := apq.Select(apppackage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AppPackageQuery) IDsX(ctx context.Context) []int {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AppPackageQuery) Count(ctx context.Context) (int, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return apq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AppPackageQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AppPackageQuery) Exist(ctx context.Context) (bool, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return apq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AppPackageQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppPackageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AppPackageQuery) Clone() *AppPackageQuery {
	if apq == nil {
		return nil
	}
	return &AppPackageQuery{
		config:     apq.config,
		limit:      apq.limit,
		offset:     apq.offset,
		order:      append([]OrderFunc{}, apq.order...),
		predicates: append([]predicate.AppPackage{}, apq.predicates...),
		// clone intermediate query.
		sql:    apq.sql.Clone(),
		path:   apq.path,
		unique: apq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InternalID int64 `json:"internal_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppPackage.Query().
//		GroupBy(apppackage.FieldInternalID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (apq *AppPackageQuery) GroupBy(field string, fields ...string) *AppPackageGroupBy {
	grbuild := &AppPackageGroupBy{config: apq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return apq.sqlQuery(ctx), nil
	}
	grbuild.label = apppackage.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InternalID int64 `json:"internal_id,omitempty"`
//	}
//
//	client.AppPackage.Query().
//		Select(apppackage.FieldInternalID).
//		Scan(ctx, &v)
//
func (apq *AppPackageQuery) Select(fields ...string) *AppPackageSelect {
	apq.fields = append(apq.fields, fields...)
	selbuild := &AppPackageSelect{AppPackageQuery: apq}
	selbuild.label = apppackage.Label
	selbuild.flds, selbuild.scan = &apq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AppPackageSelect configured with the given aggregations.
func (apq *AppPackageQuery) Aggregate(fns ...AggregateFunc) *AppPackageSelect {
	return apq.Select().Aggregate(fns...)
}

func (apq *AppPackageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range apq.fields {
		if !apppackage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *AppPackageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppPackage, error) {
	var (
		nodes = []*AppPackage{}
		_spec = apq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppPackage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppPackage{config: apq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (apq *AppPackageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	_spec.Node.Columns = apq.fields
	if len(apq.fields) > 0 {
		_spec.Unique = apq.unique != nil && *apq.unique
	}
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *AppPackageQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := apq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (apq *AppPackageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apppackage.Table,
			Columns: apppackage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: apppackage.FieldID,
			},
		},
		From:   apq.sql,
		Unique: true,
	}
	if unique := apq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := apq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apppackage.FieldID)
		for i := range fields {
			if fields[i] != apppackage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *AppPackageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(apppackage.Table)
	columns := apq.fields
	if len(columns) == 0 {
		columns = apppackage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.unique != nil && *apq.unique {
		selector.Distinct()
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppPackageGroupBy is the group-by builder for AppPackage entities.
type AppPackageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AppPackageGroupBy) Aggregate(fns ...AggregateFunc) *AppPackageGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the group-by query and scans the result into the given value.
func (apgb *AppPackageGroupBy) Scan(ctx context.Context, v any) error {
	query, err := apgb.path(ctx)
	if err != nil {
		return err
	}
	apgb.sql = query
	return apgb.sqlScan(ctx, v)
}

func (apgb *AppPackageGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range apgb.fields {
		if !apppackage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := apgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (apgb *AppPackageGroupBy) sqlQuery() *sql.Selector {
	selector := apgb.sql.Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(apgb.fields)+len(apgb.fns))
		for _, f := range apgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(apgb.fields...)...)
}

// AppPackageSelect is the builder for selecting fields of AppPackage entities.
type AppPackageSelect struct {
	*AppPackageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aps *AppPackageSelect) Aggregate(fns ...AggregateFunc) *AppPackageSelect {
	aps.fns = append(aps.fns, fns...)
	return aps
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AppPackageSelect) Scan(ctx context.Context, v any) error {
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	aps.sql = aps.AppPackageQuery.sqlQuery(ctx)
	return aps.sqlScan(ctx, v)
}

func (aps *AppPackageSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(aps.fns))
	for _, fn := range aps.fns {
		aggregation = append(aggregation, fn(aps.sql))
	}
	switch n := len(*aps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		aps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		aps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := aps.sql.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
