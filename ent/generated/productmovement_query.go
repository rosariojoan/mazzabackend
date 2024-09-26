// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/product"
	"mazza/ent/generated/productmovement"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductMovementQuery is the builder for querying ProductMovement entities.
type ProductMovementQuery struct {
	config
	ctx         *QueryContext
	order       []productmovement.OrderOption
	inters      []Interceptor
	predicates  []predicate.ProductMovement
	withProduct *ProductQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*ProductMovement) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductMovementQuery builder.
func (pmq *ProductMovementQuery) Where(ps ...predicate.ProductMovement) *ProductMovementQuery {
	pmq.predicates = append(pmq.predicates, ps...)
	return pmq
}

// Limit the number of records to be returned by this query.
func (pmq *ProductMovementQuery) Limit(limit int) *ProductMovementQuery {
	pmq.ctx.Limit = &limit
	return pmq
}

// Offset to start from.
func (pmq *ProductMovementQuery) Offset(offset int) *ProductMovementQuery {
	pmq.ctx.Offset = &offset
	return pmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pmq *ProductMovementQuery) Unique(unique bool) *ProductMovementQuery {
	pmq.ctx.Unique = &unique
	return pmq
}

// Order specifies how the records should be ordered.
func (pmq *ProductMovementQuery) Order(o ...productmovement.OrderOption) *ProductMovementQuery {
	pmq.order = append(pmq.order, o...)
	return pmq
}

// QueryProduct chains the current query on the "product" edge.
func (pmq *ProductMovementQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: pmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productmovement.Table, productmovement.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productmovement.ProductTable, productmovement.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(pmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductMovement entity from the query.
// Returns a *NotFoundError when no ProductMovement was found.
func (pmq *ProductMovementQuery) First(ctx context.Context) (*ProductMovement, error) {
	nodes, err := pmq.Limit(1).All(setContextOp(ctx, pmq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productmovement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pmq *ProductMovementQuery) FirstX(ctx context.Context) *ProductMovement {
	node, err := pmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductMovement ID from the query.
// Returns a *NotFoundError when no ProductMovement ID was found.
func (pmq *ProductMovementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmq.Limit(1).IDs(setContextOp(ctx, pmq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productmovement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pmq *ProductMovementQuery) FirstIDX(ctx context.Context) int {
	id, err := pmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductMovement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductMovement entity is found.
// Returns a *NotFoundError when no ProductMovement entities are found.
func (pmq *ProductMovementQuery) Only(ctx context.Context) (*ProductMovement, error) {
	nodes, err := pmq.Limit(2).All(setContextOp(ctx, pmq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productmovement.Label}
	default:
		return nil, &NotSingularError{productmovement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pmq *ProductMovementQuery) OnlyX(ctx context.Context) *ProductMovement {
	node, err := pmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductMovement ID in the query.
// Returns a *NotSingularError when more than one ProductMovement ID is found.
// Returns a *NotFoundError when no entities are found.
func (pmq *ProductMovementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmq.Limit(2).IDs(setContextOp(ctx, pmq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productmovement.Label}
	default:
		err = &NotSingularError{productmovement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pmq *ProductMovementQuery) OnlyIDX(ctx context.Context) int {
	id, err := pmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductMovements.
func (pmq *ProductMovementQuery) All(ctx context.Context) ([]*ProductMovement, error) {
	ctx = setContextOp(ctx, pmq.ctx, ent.OpQueryAll)
	if err := pmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductMovement, *ProductMovementQuery]()
	return withInterceptors[[]*ProductMovement](ctx, pmq, qr, pmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pmq *ProductMovementQuery) AllX(ctx context.Context) []*ProductMovement {
	nodes, err := pmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductMovement IDs.
func (pmq *ProductMovementQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pmq.ctx.Unique == nil && pmq.path != nil {
		pmq.Unique(true)
	}
	ctx = setContextOp(ctx, pmq.ctx, ent.OpQueryIDs)
	if err = pmq.Select(productmovement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pmq *ProductMovementQuery) IDsX(ctx context.Context) []int {
	ids, err := pmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pmq *ProductMovementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pmq.ctx, ent.OpQueryCount)
	if err := pmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pmq, querierCount[*ProductMovementQuery](), pmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pmq *ProductMovementQuery) CountX(ctx context.Context) int {
	count, err := pmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pmq *ProductMovementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pmq.ctx, ent.OpQueryExist)
	switch _, err := pmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pmq *ProductMovementQuery) ExistX(ctx context.Context) bool {
	exist, err := pmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductMovementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pmq *ProductMovementQuery) Clone() *ProductMovementQuery {
	if pmq == nil {
		return nil
	}
	return &ProductMovementQuery{
		config:      pmq.config,
		ctx:         pmq.ctx.Clone(),
		order:       append([]productmovement.OrderOption{}, pmq.order...),
		inters:      append([]Interceptor{}, pmq.inters...),
		predicates:  append([]predicate.ProductMovement{}, pmq.predicates...),
		withProduct: pmq.withProduct.Clone(),
		// clone intermediate query.
		sql:       pmq.sql.Clone(),
		path:      pmq.path,
		modifiers: append([]func(*sql.Selector){}, pmq.modifiers...),
	}
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (pmq *ProductMovementQuery) WithProduct(opts ...func(*ProductQuery)) *ProductMovementQuery {
	query := (&ProductClient{config: pmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pmq.withProduct = query
	return pmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProductMovement.Query().
//		GroupBy(productmovement.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (pmq *ProductMovementQuery) GroupBy(field string, fields ...string) *ProductMovementGroupBy {
	pmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductMovementGroupBy{build: pmq}
	grbuild.flds = &pmq.ctx.Fields
	grbuild.label = productmovement.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.ProductMovement.Query().
//		Select(productmovement.FieldCreatedAt).
//		Scan(ctx, &v)
func (pmq *ProductMovementQuery) Select(fields ...string) *ProductMovementSelect {
	pmq.ctx.Fields = append(pmq.ctx.Fields, fields...)
	sbuild := &ProductMovementSelect{ProductMovementQuery: pmq}
	sbuild.label = productmovement.Label
	sbuild.flds, sbuild.scan = &pmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductMovementSelect configured with the given aggregations.
func (pmq *ProductMovementQuery) Aggregate(fns ...AggregateFunc) *ProductMovementSelect {
	return pmq.Select().Aggregate(fns...)
}

func (pmq *ProductMovementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pmq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pmq); err != nil {
				return err
			}
		}
	}
	for _, f := range pmq.ctx.Fields {
		if !productmovement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
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

func (pmq *ProductMovementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductMovement, error) {
	var (
		nodes       = []*ProductMovement{}
		withFKs     = pmq.withFKs
		_spec       = pmq.querySpec()
		loadedTypes = [1]bool{
			pmq.withProduct != nil,
		}
	)
	if pmq.withProduct != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, productmovement.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductMovement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductMovement{config: pmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pmq.modifiers) > 0 {
		_spec.Modifiers = pmq.modifiers
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
	if query := pmq.withProduct; query != nil {
		if err := pmq.loadProduct(ctx, query, nodes, nil,
			func(n *ProductMovement, e *Product) { n.Edges.Product = e }); err != nil {
			return nil, err
		}
	}
	for i := range pmq.loadTotal {
		if err := pmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pmq *ProductMovementQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*ProductMovement, init func(*ProductMovement), assign func(*ProductMovement, *Product)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProductMovement)
	for i := range nodes {
		if nodes[i].product_product_movements == nil {
			continue
		}
		fk := *nodes[i].product_product_movements
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(product.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_product_movements" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pmq *ProductMovementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pmq.querySpec()
	if len(pmq.modifiers) > 0 {
		_spec.Modifiers = pmq.modifiers
	}
	_spec.Node.Columns = pmq.ctx.Fields
	if len(pmq.ctx.Fields) > 0 {
		_spec.Unique = pmq.ctx.Unique != nil && *pmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pmq.driver, _spec)
}

func (pmq *ProductMovementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productmovement.Table, productmovement.Columns, sqlgraph.NewFieldSpec(productmovement.FieldID, field.TypeInt))
	_spec.From = pmq.sql
	if unique := pmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pmq.path != nil {
		_spec.Unique = true
	}
	if fields := pmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productmovement.FieldID)
		for i := range fields {
			if fields[i] != productmovement.FieldID {
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

func (pmq *ProductMovementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pmq.driver.Dialect())
	t1 := builder.Table(productmovement.Table)
	columns := pmq.ctx.Fields
	if len(columns) == 0 {
		columns = productmovement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pmq.sql != nil {
		selector = pmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pmq.ctx.Unique != nil && *pmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pmq.modifiers {
		m(selector)
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

// Modify adds a query modifier for attaching custom logic to queries.
func (pmq *ProductMovementQuery) Modify(modifiers ...func(s *sql.Selector)) *ProductMovementSelect {
	pmq.modifiers = append(pmq.modifiers, modifiers...)
	return pmq.Select()
}

// ProductMovementGroupBy is the group-by builder for ProductMovement entities.
type ProductMovementGroupBy struct {
	selector
	build *ProductMovementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pmgb *ProductMovementGroupBy) Aggregate(fns ...AggregateFunc) *ProductMovementGroupBy {
	pmgb.fns = append(pmgb.fns, fns...)
	return pmgb
}

// Scan applies the selector query and scans the result into the given value.
func (pmgb *ProductMovementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pmgb.build.ctx, ent.OpQueryGroupBy)
	if err := pmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductMovementQuery, *ProductMovementGroupBy](ctx, pmgb.build, pmgb, pmgb.build.inters, v)
}

func (pmgb *ProductMovementGroupBy) sqlScan(ctx context.Context, root *ProductMovementQuery, v any) error {
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

// ProductMovementSelect is the builder for selecting fields of ProductMovement entities.
type ProductMovementSelect struct {
	*ProductMovementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pms *ProductMovementSelect) Aggregate(fns ...AggregateFunc) *ProductMovementSelect {
	pms.fns = append(pms.fns, fns...)
	return pms
}

// Scan applies the selector query and scans the result into the given value.
func (pms *ProductMovementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pms.ctx, ent.OpQuerySelect)
	if err := pms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductMovementQuery, *ProductMovementSelect](ctx, pms.ProductMovementQuery, pms, pms.inters, v)
}

func (pms *ProductMovementSelect) sqlScan(ctx context.Context, root *ProductMovementQuery, v any) error {
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

// Modify adds a query modifier for attaching custom logic to queries.
func (pms *ProductMovementSelect) Modify(modifiers ...func(s *sql.Selector)) *ProductMovementSelect {
	pms.modifiers = append(pms.modifiers, modifiers...)
	return pms
}
