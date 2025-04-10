// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"mazza/ent/generated/company"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/receivable"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReceivableQuery is the builder for querying Receivable entities.
type ReceivableQuery struct {
	config
	ctx         *QueryContext
	order       []receivable.OrderOption
	inters      []Interceptor
	predicates  []predicate.Receivable
	withCompany *CompanyQuery
	withInvoice *InvoiceQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*Receivable) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReceivableQuery builder.
func (rq *ReceivableQuery) Where(ps ...predicate.Receivable) *ReceivableQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReceivableQuery) Limit(limit int) *ReceivableQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReceivableQuery) Offset(offset int) *ReceivableQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReceivableQuery) Unique(unique bool) *ReceivableQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReceivableQuery) Order(o ...receivable.OrderOption) *ReceivableQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryCompany chains the current query on the "company" edge.
func (rq *ReceivableQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receivable.Table, receivable.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, receivable.CompanyTable, receivable.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvoice chains the current query on the "invoice" edge.
func (rq *ReceivableQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receivable.Table, receivable.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, receivable.InvoiceTable, receivable.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Receivable entity from the query.
// Returns a *NotFoundError when no Receivable was found.
func (rq *ReceivableQuery) First(ctx context.Context) (*Receivable, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{receivable.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReceivableQuery) FirstX(ctx context.Context) *Receivable {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Receivable ID from the query.
// Returns a *NotFoundError when no Receivable ID was found.
func (rq *ReceivableQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{receivable.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReceivableQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Receivable entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Receivable entity is found.
// Returns a *NotFoundError when no Receivable entities are found.
func (rq *ReceivableQuery) Only(ctx context.Context) (*Receivable, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{receivable.Label}
	default:
		return nil, &NotSingularError{receivable.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReceivableQuery) OnlyX(ctx context.Context) *Receivable {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Receivable ID in the query.
// Returns a *NotSingularError when more than one Receivable ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReceivableQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{receivable.Label}
	default:
		err = &NotSingularError{receivable.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReceivableQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Receivables.
func (rq *ReceivableQuery) All(ctx context.Context) ([]*Receivable, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryAll)
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Receivable, *ReceivableQuery]()
	return withInterceptors[[]*Receivable](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReceivableQuery) AllX(ctx context.Context) []*Receivable {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Receivable IDs.
func (rq *ReceivableQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryIDs)
	if err = rq.Select(receivable.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReceivableQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReceivableQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryCount)
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReceivableQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReceivableQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReceivableQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryExist)
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReceivableQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReceivableQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReceivableQuery) Clone() *ReceivableQuery {
	if rq == nil {
		return nil
	}
	return &ReceivableQuery{
		config:      rq.config,
		ctx:         rq.ctx.Clone(),
		order:       append([]receivable.OrderOption{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Receivable{}, rq.predicates...),
		withCompany: rq.withCompany.Clone(),
		withInvoice: rq.withInvoice.Clone(),
		// clone intermediate query.
		sql:       rq.sql.Clone(),
		path:      rq.path,
		modifiers: append([]func(*sql.Selector){}, rq.modifiers...),
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReceivableQuery) WithCompany(opts ...func(*CompanyQuery)) *ReceivableQuery {
	query := (&CompanyClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withCompany = query
	return rq
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReceivableQuery) WithInvoice(opts ...func(*InvoiceQuery)) *ReceivableQuery {
	query := (&InvoiceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withInvoice = query
	return rq
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
//	client.Receivable.Query().
//		GroupBy(receivable.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (rq *ReceivableQuery) GroupBy(field string, fields ...string) *ReceivableGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReceivableGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = receivable.Label
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
//	client.Receivable.Query().
//		Select(receivable.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *ReceivableQuery) Select(fields ...string) *ReceivableSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReceivableSelect{ReceivableQuery: rq}
	sbuild.label = receivable.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReceivableSelect configured with the given aggregations.
func (rq *ReceivableQuery) Aggregate(fns ...AggregateFunc) *ReceivableSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReceivableQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !receivable.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReceivableQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Receivable, error) {
	var (
		nodes       = []*Receivable{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withCompany != nil,
			rq.withInvoice != nil,
		}
	)
	if rq.withCompany != nil || rq.withInvoice != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, receivable.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Receivable).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Receivable{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withCompany; query != nil {
		if err := rq.loadCompany(ctx, query, nodes, nil,
			func(n *Receivable, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withInvoice; query != nil {
		if err := rq.loadInvoice(ctx, query, nodes, nil,
			func(n *Receivable, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	for i := range rq.loadTotal {
		if err := rq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReceivableQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Receivable, init func(*Receivable), assign func(*Receivable, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Receivable)
	for i := range nodes {
		if nodes[i].company_receivables == nil {
			continue
		}
		fk := *nodes[i].company_receivables
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_receivables" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReceivableQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*Receivable, init func(*Receivable), assign func(*Receivable, *Invoice)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Receivable)
	for i := range nodes {
		if nodes[i].invoice_receivable == nil {
			continue
		}
		fk := *nodes[i].invoice_receivable
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_receivable" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *ReceivableQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReceivableQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(receivable.Table, receivable.Columns, sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, receivable.FieldID)
		for i := range fields {
			if fields[i] != receivable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReceivableQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(receivable.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = receivable.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rq.modifiers {
		m(selector)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rq *ReceivableQuery) Modify(modifiers ...func(s *sql.Selector)) *ReceivableSelect {
	rq.modifiers = append(rq.modifiers, modifiers...)
	return rq.Select()
}

// ReceivableGroupBy is the group-by builder for Receivable entities.
type ReceivableGroupBy struct {
	selector
	build *ReceivableQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReceivableGroupBy) Aggregate(fns ...AggregateFunc) *ReceivableGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReceivableGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, ent.OpQueryGroupBy)
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReceivableQuery, *ReceivableGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReceivableGroupBy) sqlScan(ctx context.Context, root *ReceivableQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReceivableSelect is the builder for selecting fields of Receivable entities.
type ReceivableSelect struct {
	*ReceivableQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReceivableSelect) Aggregate(fns ...AggregateFunc) *ReceivableSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReceivableSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, ent.OpQuerySelect)
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReceivableQuery, *ReceivableSelect](ctx, rs.ReceivableQuery, rs, rs.inters, v)
}

func (rs *ReceivableSelect) sqlScan(ctx context.Context, root *ReceivableQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *ReceivableSelect) Modify(modifiers ...func(s *sql.Selector)) *ReceivableSelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
