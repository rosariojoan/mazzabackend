// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountingEntryQuery is the builder for querying AccountingEntry entities.
type AccountingEntryQuery struct {
	config
	ctx         *QueryContext
	order       []accountingentry.OrderOption
	inters      []Interceptor
	predicates  []predicate.AccountingEntry
	withCompany *CompanyQuery
	withUser    *UserQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*AccountingEntry) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountingEntryQuery builder.
func (aeq *AccountingEntryQuery) Where(ps ...predicate.AccountingEntry) *AccountingEntryQuery {
	aeq.predicates = append(aeq.predicates, ps...)
	return aeq
}

// Limit the number of records to be returned by this query.
func (aeq *AccountingEntryQuery) Limit(limit int) *AccountingEntryQuery {
	aeq.ctx.Limit = &limit
	return aeq
}

// Offset to start from.
func (aeq *AccountingEntryQuery) Offset(offset int) *AccountingEntryQuery {
	aeq.ctx.Offset = &offset
	return aeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aeq *AccountingEntryQuery) Unique(unique bool) *AccountingEntryQuery {
	aeq.ctx.Unique = &unique
	return aeq
}

// Order specifies how the records should be ordered.
func (aeq *AccountingEntryQuery) Order(o ...accountingentry.OrderOption) *AccountingEntryQuery {
	aeq.order = append(aeq.order, o...)
	return aeq
}

// QueryCompany chains the current query on the "company" edge.
func (aeq *AccountingEntryQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: aeq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountingentry.Table, accountingentry.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountingentry.CompanyTable, accountingentry.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(aeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (aeq *AccountingEntryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: aeq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountingentry.Table, accountingentry.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountingentry.UserTable, accountingentry.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccountingEntry entity from the query.
// Returns a *NotFoundError when no AccountingEntry was found.
func (aeq *AccountingEntryQuery) First(ctx context.Context) (*AccountingEntry, error) {
	nodes, err := aeq.Limit(1).All(setContextOp(ctx, aeq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountingentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aeq *AccountingEntryQuery) FirstX(ctx context.Context) *AccountingEntry {
	node, err := aeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountingEntry ID from the query.
// Returns a *NotFoundError when no AccountingEntry ID was found.
func (aeq *AccountingEntryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aeq.Limit(1).IDs(setContextOp(ctx, aeq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountingentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aeq *AccountingEntryQuery) FirstIDX(ctx context.Context) int {
	id, err := aeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountingEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountingEntry entity is found.
// Returns a *NotFoundError when no AccountingEntry entities are found.
func (aeq *AccountingEntryQuery) Only(ctx context.Context) (*AccountingEntry, error) {
	nodes, err := aeq.Limit(2).All(setContextOp(ctx, aeq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountingentry.Label}
	default:
		return nil, &NotSingularError{accountingentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aeq *AccountingEntryQuery) OnlyX(ctx context.Context) *AccountingEntry {
	node, err := aeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountingEntry ID in the query.
// Returns a *NotSingularError when more than one AccountingEntry ID is found.
// Returns a *NotFoundError when no entities are found.
func (aeq *AccountingEntryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aeq.Limit(2).IDs(setContextOp(ctx, aeq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountingentry.Label}
	default:
		err = &NotSingularError{accountingentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aeq *AccountingEntryQuery) OnlyIDX(ctx context.Context) int {
	id, err := aeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountingEntries.
func (aeq *AccountingEntryQuery) All(ctx context.Context) ([]*AccountingEntry, error) {
	ctx = setContextOp(ctx, aeq.ctx, ent.OpQueryAll)
	if err := aeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccountingEntry, *AccountingEntryQuery]()
	return withInterceptors[[]*AccountingEntry](ctx, aeq, qr, aeq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aeq *AccountingEntryQuery) AllX(ctx context.Context) []*AccountingEntry {
	nodes, err := aeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountingEntry IDs.
func (aeq *AccountingEntryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aeq.ctx.Unique == nil && aeq.path != nil {
		aeq.Unique(true)
	}
	ctx = setContextOp(ctx, aeq.ctx, ent.OpQueryIDs)
	if err = aeq.Select(accountingentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aeq *AccountingEntryQuery) IDsX(ctx context.Context) []int {
	ids, err := aeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aeq *AccountingEntryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aeq.ctx, ent.OpQueryCount)
	if err := aeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aeq, querierCount[*AccountingEntryQuery](), aeq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aeq *AccountingEntryQuery) CountX(ctx context.Context) int {
	count, err := aeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aeq *AccountingEntryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aeq.ctx, ent.OpQueryExist)
	switch _, err := aeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aeq *AccountingEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := aeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountingEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aeq *AccountingEntryQuery) Clone() *AccountingEntryQuery {
	if aeq == nil {
		return nil
	}
	return &AccountingEntryQuery{
		config:      aeq.config,
		ctx:         aeq.ctx.Clone(),
		order:       append([]accountingentry.OrderOption{}, aeq.order...),
		inters:      append([]Interceptor{}, aeq.inters...),
		predicates:  append([]predicate.AccountingEntry{}, aeq.predicates...),
		withCompany: aeq.withCompany.Clone(),
		withUser:    aeq.withUser.Clone(),
		// clone intermediate query.
		sql:       aeq.sql.Clone(),
		path:      aeq.path,
		modifiers: append([]func(*sql.Selector){}, aeq.modifiers...),
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (aeq *AccountingEntryQuery) WithCompany(opts ...func(*CompanyQuery)) *AccountingEntryQuery {
	query := (&CompanyClient{config: aeq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aeq.withCompany = query
	return aeq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aeq *AccountingEntryQuery) WithUser(opts ...func(*UserQuery)) *AccountingEntryQuery {
	query := (&UserClient{config: aeq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aeq.withUser = query
	return aeq
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
//	client.AccountingEntry.Query().
//		GroupBy(accountingentry.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (aeq *AccountingEntryQuery) GroupBy(field string, fields ...string) *AccountingEntryGroupBy {
	aeq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountingEntryGroupBy{build: aeq}
	grbuild.flds = &aeq.ctx.Fields
	grbuild.label = accountingentry.Label
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
//	client.AccountingEntry.Query().
//		Select(accountingentry.FieldCreatedAt).
//		Scan(ctx, &v)
func (aeq *AccountingEntryQuery) Select(fields ...string) *AccountingEntrySelect {
	aeq.ctx.Fields = append(aeq.ctx.Fields, fields...)
	sbuild := &AccountingEntrySelect{AccountingEntryQuery: aeq}
	sbuild.label = accountingentry.Label
	sbuild.flds, sbuild.scan = &aeq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountingEntrySelect configured with the given aggregations.
func (aeq *AccountingEntryQuery) Aggregate(fns ...AggregateFunc) *AccountingEntrySelect {
	return aeq.Select().Aggregate(fns...)
}

func (aeq *AccountingEntryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aeq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aeq); err != nil {
				return err
			}
		}
	}
	for _, f := range aeq.ctx.Fields {
		if !accountingentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if aeq.path != nil {
		prev, err := aeq.path(ctx)
		if err != nil {
			return err
		}
		aeq.sql = prev
	}
	return nil
}

func (aeq *AccountingEntryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountingEntry, error) {
	var (
		nodes       = []*AccountingEntry{}
		withFKs     = aeq.withFKs
		_spec       = aeq.querySpec()
		loadedTypes = [2]bool{
			aeq.withCompany != nil,
			aeq.withUser != nil,
		}
	)
	if aeq.withCompany != nil || aeq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, accountingentry.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountingEntry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountingEntry{config: aeq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aeq.modifiers) > 0 {
		_spec.Modifiers = aeq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aeq.withCompany; query != nil {
		if err := aeq.loadCompany(ctx, query, nodes, nil,
			func(n *AccountingEntry, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := aeq.withUser; query != nil {
		if err := aeq.loadUser(ctx, query, nodes, nil,
			func(n *AccountingEntry, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range aeq.loadTotal {
		if err := aeq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aeq *AccountingEntryQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*AccountingEntry, init func(*AccountingEntry), assign func(*AccountingEntry, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AccountingEntry)
	for i := range nodes {
		if nodes[i].company_accounting_entries == nil {
			continue
		}
		fk := *nodes[i].company_accounting_entries
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
			return fmt.Errorf(`unexpected foreign-key "company_accounting_entries" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aeq *AccountingEntryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*AccountingEntry, init func(*AccountingEntry), assign func(*AccountingEntry, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AccountingEntry)
	for i := range nodes {
		if nodes[i].user_accounting_entries == nil {
			continue
		}
		fk := *nodes[i].user_accounting_entries
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_accounting_entries" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aeq *AccountingEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aeq.querySpec()
	if len(aeq.modifiers) > 0 {
		_spec.Modifiers = aeq.modifiers
	}
	_spec.Node.Columns = aeq.ctx.Fields
	if len(aeq.ctx.Fields) > 0 {
		_spec.Unique = aeq.ctx.Unique != nil && *aeq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aeq.driver, _spec)
}

func (aeq *AccountingEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accountingentry.Table, accountingentry.Columns, sqlgraph.NewFieldSpec(accountingentry.FieldID, field.TypeInt))
	_spec.From = aeq.sql
	if unique := aeq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aeq.path != nil {
		_spec.Unique = true
	}
	if fields := aeq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountingentry.FieldID)
		for i := range fields {
			if fields[i] != accountingentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aeq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aeq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aeq *AccountingEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aeq.driver.Dialect())
	t1 := builder.Table(accountingentry.Table)
	columns := aeq.ctx.Fields
	if len(columns) == 0 {
		columns = accountingentry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aeq.sql != nil {
		selector = aeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aeq.ctx.Unique != nil && *aeq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aeq.modifiers {
		m(selector)
	}
	for _, p := range aeq.predicates {
		p(selector)
	}
	for _, p := range aeq.order {
		p(selector)
	}
	if offset := aeq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aeq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aeq *AccountingEntryQuery) Modify(modifiers ...func(s *sql.Selector)) *AccountingEntrySelect {
	aeq.modifiers = append(aeq.modifiers, modifiers...)
	return aeq.Select()
}

// AccountingEntryGroupBy is the group-by builder for AccountingEntry entities.
type AccountingEntryGroupBy struct {
	selector
	build *AccountingEntryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aegb *AccountingEntryGroupBy) Aggregate(fns ...AggregateFunc) *AccountingEntryGroupBy {
	aegb.fns = append(aegb.fns, fns...)
	return aegb
}

// Scan applies the selector query and scans the result into the given value.
func (aegb *AccountingEntryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aegb.build.ctx, ent.OpQueryGroupBy)
	if err := aegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountingEntryQuery, *AccountingEntryGroupBy](ctx, aegb.build, aegb, aegb.build.inters, v)
}

func (aegb *AccountingEntryGroupBy) sqlScan(ctx context.Context, root *AccountingEntryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aegb.fns))
	for _, fn := range aegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aegb.flds)+len(aegb.fns))
		for _, f := range *aegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountingEntrySelect is the builder for selecting fields of AccountingEntry entities.
type AccountingEntrySelect struct {
	*AccountingEntryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aes *AccountingEntrySelect) Aggregate(fns ...AggregateFunc) *AccountingEntrySelect {
	aes.fns = append(aes.fns, fns...)
	return aes
}

// Scan applies the selector query and scans the result into the given value.
func (aes *AccountingEntrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aes.ctx, ent.OpQuerySelect)
	if err := aes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountingEntryQuery, *AccountingEntrySelect](ctx, aes.AccountingEntryQuery, aes, aes.inters, v)
}

func (aes *AccountingEntrySelect) sqlScan(ctx context.Context, root *AccountingEntryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aes.fns))
	for _, fn := range aes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aes *AccountingEntrySelect) Modify(modifiers ...func(s *sql.Selector)) *AccountingEntrySelect {
	aes.modifiers = append(aes.modifiers, modifiers...)
	return aes
}
