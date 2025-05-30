// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/receivable"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerQuery is the builder for querying Customer entities.
type CustomerQuery struct {
	config
	ctx                  *QueryContext
	order                []customer.OrderOption
	inters               []Interceptor
	predicates           []predicate.Customer
	withCompany          *CompanyQuery
	withReceivables      *ReceivableQuery
	withInvoices         *InvoiceQuery
	withFKs              bool
	loadTotal            []func(context.Context, []*Customer) error
	modifiers            []func(*sql.Selector)
	withNamedReceivables map[string]*ReceivableQuery
	withNamedInvoices    map[string]*InvoiceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerQuery builder.
func (cq *CustomerQuery) Where(ps ...predicate.Customer) *CustomerQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CustomerQuery) Limit(limit int) *CustomerQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CustomerQuery) Offset(offset int) *CustomerQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CustomerQuery) Unique(unique bool) *CustomerQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CustomerQuery) Order(o ...customer.OrderOption) *CustomerQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCompany chains the current query on the "company" edge.
func (cq *CustomerQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, customer.CompanyTable, customer.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReceivables chains the current query on the "receivables" edge.
func (cq *CustomerQuery) QueryReceivables() *ReceivableQuery {
	query := (&ReceivableClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(receivable.Table, receivable.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.ReceivablesTable, customer.ReceivablesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvoices chains the current query on the "invoices" edge.
func (cq *CustomerQuery) QueryInvoices() *InvoiceQuery {
	query := (&InvoiceClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.InvoicesTable, customer.InvoicesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Customer entity from the query.
// Returns a *NotFoundError when no Customer was found.
func (cq *CustomerQuery) First(ctx context.Context) (*Customer, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CustomerQuery) FirstX(ctx context.Context) *Customer {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Customer ID from the query.
// Returns a *NotFoundError when no Customer ID was found.
func (cq *CustomerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CustomerQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Customer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Customer entity is found.
// Returns a *NotFoundError when no Customer entities are found.
func (cq *CustomerQuery) Only(ctx context.Context) (*Customer, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customer.Label}
	default:
		return nil, &NotSingularError{customer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CustomerQuery) OnlyX(ctx context.Context) *Customer {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Customer ID in the query.
// Returns a *NotSingularError when more than one Customer ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CustomerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customer.Label}
	default:
		err = &NotSingularError{customer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CustomerQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Customers.
func (cq *CustomerQuery) All(ctx context.Context) ([]*Customer, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Customer, *CustomerQuery]()
	return withInterceptors[[]*Customer](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CustomerQuery) AllX(ctx context.Context) []*Customer {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Customer IDs.
func (cq *CustomerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(customer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CustomerQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CustomerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CustomerQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CustomerQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CustomerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CustomerQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CustomerQuery) Clone() *CustomerQuery {
	if cq == nil {
		return nil
	}
	return &CustomerQuery{
		config:          cq.config,
		ctx:             cq.ctx.Clone(),
		order:           append([]customer.OrderOption{}, cq.order...),
		inters:          append([]Interceptor{}, cq.inters...),
		predicates:      append([]predicate.Customer{}, cq.predicates...),
		withCompany:     cq.withCompany.Clone(),
		withReceivables: cq.withReceivables.Clone(),
		withInvoices:    cq.withInvoices.Clone(),
		// clone intermediate query.
		sql:       cq.sql.Clone(),
		path:      cq.path,
		modifiers: append([]func(*sql.Selector){}, cq.modifiers...),
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithCompany(opts ...func(*CompanyQuery)) *CustomerQuery {
	query := (&CompanyClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCompany = query
	return cq
}

// WithReceivables tells the query-builder to eager-load the nodes that are connected to
// the "receivables" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithReceivables(opts ...func(*ReceivableQuery)) *CustomerQuery {
	query := (&ReceivableClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withReceivables = query
	return cq
}

// WithInvoices tells the query-builder to eager-load the nodes that are connected to
// the "invoices" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithInvoices(opts ...func(*InvoiceQuery)) *CustomerQuery {
	query := (&InvoiceClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withInvoices = query
	return cq
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
//	client.Customer.Query().
//		GroupBy(customer.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (cq *CustomerQuery) GroupBy(field string, fields ...string) *CustomerGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = customer.Label
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
//	client.Customer.Query().
//		Select(customer.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CustomerQuery) Select(fields ...string) *CustomerSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CustomerSelect{CustomerQuery: cq}
	sbuild.label = customer.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerSelect configured with the given aggregations.
func (cq *CustomerQuery) Aggregate(fns ...AggregateFunc) *CustomerSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CustomerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !customer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CustomerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Customer, error) {
	var (
		nodes       = []*Customer{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withCompany != nil,
			cq.withReceivables != nil,
			cq.withInvoices != nil,
		}
	)
	if cq.withCompany != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, customer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Customer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Customer{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCompany; query != nil {
		if err := cq.loadCompany(ctx, query, nodes, nil,
			func(n *Customer, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withReceivables; query != nil {
		if err := cq.loadReceivables(ctx, query, nodes,
			func(n *Customer) { n.Edges.Receivables = []*Receivable{} },
			func(n *Customer, e *Receivable) { n.Edges.Receivables = append(n.Edges.Receivables, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withInvoices; query != nil {
		if err := cq.loadInvoices(ctx, query, nodes,
			func(n *Customer) { n.Edges.Invoices = []*Invoice{} },
			func(n *Customer, e *Invoice) { n.Edges.Invoices = append(n.Edges.Invoices, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedReceivables {
		if err := cq.loadReceivables(ctx, query, nodes,
			func(n *Customer) { n.appendNamedReceivables(name) },
			func(n *Customer, e *Receivable) { n.appendNamedReceivables(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedInvoices {
		if err := cq.loadInvoices(ctx, query, nodes,
			func(n *Customer) { n.appendNamedInvoices(name) },
			func(n *Customer, e *Invoice) { n.appendNamedInvoices(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CustomerQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Customer)
	for i := range nodes {
		if nodes[i].company_customers == nil {
			continue
		}
		fk := *nodes[i].company_customers
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
			return fmt.Errorf(`unexpected foreign-key "company_customers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CustomerQuery) loadReceivables(ctx context.Context, query *ReceivableQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Receivable)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Receivable(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.ReceivablesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.customer_receivables
		if fk == nil {
			return fmt.Errorf(`foreign-key "customer_receivables" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_receivables" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CustomerQuery) loadInvoices(ctx context.Context, query *InvoiceQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Invoice)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Invoice(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.InvoicesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.customer_invoices
		if fk == nil {
			return fmt.Errorf(`foreign-key "customer_invoices" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_invoices" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CustomerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CustomerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for i := range fields {
			if fields[i] != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CustomerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(customer.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = customer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CustomerQuery) Modify(modifiers ...func(s *sql.Selector)) *CustomerSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// WithNamedReceivables tells the query-builder to eager-load the nodes that are connected to the "receivables"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithNamedReceivables(name string, opts ...func(*ReceivableQuery)) *CustomerQuery {
	query := (&ReceivableClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedReceivables == nil {
		cq.withNamedReceivables = make(map[string]*ReceivableQuery)
	}
	cq.withNamedReceivables[name] = query
	return cq
}

// WithNamedInvoices tells the query-builder to eager-load the nodes that are connected to the "invoices"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithNamedInvoices(name string, opts ...func(*InvoiceQuery)) *CustomerQuery {
	query := (&InvoiceClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedInvoices == nil {
		cq.withNamedInvoices = make(map[string]*InvoiceQuery)
	}
	cq.withNamedInvoices[name] = query
	return cq
}

// CustomerGroupBy is the group-by builder for Customer entities.
type CustomerGroupBy struct {
	selector
	build *CustomerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CustomerGroupBy) Aggregate(fns ...AggregateFunc) *CustomerGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CustomerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerQuery, *CustomerGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CustomerGroupBy) sqlScan(ctx context.Context, root *CustomerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerSelect is the builder for selecting fields of Customer entities.
type CustomerSelect struct {
	*CustomerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CustomerSelect) Aggregate(fns ...AggregateFunc) *CustomerSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CustomerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerQuery, *CustomerSelect](ctx, cs.CustomerQuery, cs, cs.inters, v)
}

func (cs *CustomerSelect) sqlScan(ctx context.Context, root *CustomerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CustomerSelect) Modify(modifiers ...func(s *sql.Selector)) *CustomerSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
