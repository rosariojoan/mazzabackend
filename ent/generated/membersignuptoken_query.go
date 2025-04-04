// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"mazza/ent/generated/company"
	"mazza/ent/generated/membersignuptoken"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberSignupTokenQuery is the builder for querying MemberSignupToken entities.
type MemberSignupTokenQuery struct {
	config
	ctx           *QueryContext
	order         []membersignuptoken.OrderOption
	inters        []Interceptor
	predicates    []predicate.MemberSignupToken
	withCompany   *CompanyQuery
	withCreatedBy *UserQuery
	withFKs       bool
	loadTotal     []func(context.Context, []*MemberSignupToken) error
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberSignupTokenQuery builder.
func (mstq *MemberSignupTokenQuery) Where(ps ...predicate.MemberSignupToken) *MemberSignupTokenQuery {
	mstq.predicates = append(mstq.predicates, ps...)
	return mstq
}

// Limit the number of records to be returned by this query.
func (mstq *MemberSignupTokenQuery) Limit(limit int) *MemberSignupTokenQuery {
	mstq.ctx.Limit = &limit
	return mstq
}

// Offset to start from.
func (mstq *MemberSignupTokenQuery) Offset(offset int) *MemberSignupTokenQuery {
	mstq.ctx.Offset = &offset
	return mstq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mstq *MemberSignupTokenQuery) Unique(unique bool) *MemberSignupTokenQuery {
	mstq.ctx.Unique = &unique
	return mstq
}

// Order specifies how the records should be ordered.
func (mstq *MemberSignupTokenQuery) Order(o ...membersignuptoken.OrderOption) *MemberSignupTokenQuery {
	mstq.order = append(mstq.order, o...)
	return mstq
}

// QueryCompany chains the current query on the "company" edge.
func (mstq *MemberSignupTokenQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: mstq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mstq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mstq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membersignuptoken.Table, membersignuptoken.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membersignuptoken.CompanyTable, membersignuptoken.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(mstq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedBy chains the current query on the "createdBy" edge.
func (mstq *MemberSignupTokenQuery) QueryCreatedBy() *UserQuery {
	query := (&UserClient{config: mstq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mstq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mstq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membersignuptoken.Table, membersignuptoken.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membersignuptoken.CreatedByTable, membersignuptoken.CreatedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(mstq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MemberSignupToken entity from the query.
// Returns a *NotFoundError when no MemberSignupToken was found.
func (mstq *MemberSignupTokenQuery) First(ctx context.Context) (*MemberSignupToken, error) {
	nodes, err := mstq.Limit(1).All(setContextOp(ctx, mstq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{membersignuptoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) FirstX(ctx context.Context) *MemberSignupToken {
	node, err := mstq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MemberSignupToken ID from the query.
// Returns a *NotFoundError when no MemberSignupToken ID was found.
func (mstq *MemberSignupTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mstq.Limit(1).IDs(setContextOp(ctx, mstq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{membersignuptoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := mstq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MemberSignupToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MemberSignupToken entity is found.
// Returns a *NotFoundError when no MemberSignupToken entities are found.
func (mstq *MemberSignupTokenQuery) Only(ctx context.Context) (*MemberSignupToken, error) {
	nodes, err := mstq.Limit(2).All(setContextOp(ctx, mstq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{membersignuptoken.Label}
	default:
		return nil, &NotSingularError{membersignuptoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) OnlyX(ctx context.Context) *MemberSignupToken {
	node, err := mstq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MemberSignupToken ID in the query.
// Returns a *NotSingularError when more than one MemberSignupToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (mstq *MemberSignupTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mstq.Limit(2).IDs(setContextOp(ctx, mstq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{membersignuptoken.Label}
	default:
		err = &NotSingularError{membersignuptoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := mstq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MemberSignupTokens.
func (mstq *MemberSignupTokenQuery) All(ctx context.Context) ([]*MemberSignupToken, error) {
	ctx = setContextOp(ctx, mstq.ctx, ent.OpQueryAll)
	if err := mstq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MemberSignupToken, *MemberSignupTokenQuery]()
	return withInterceptors[[]*MemberSignupToken](ctx, mstq, qr, mstq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) AllX(ctx context.Context) []*MemberSignupToken {
	nodes, err := mstq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MemberSignupToken IDs.
func (mstq *MemberSignupTokenQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mstq.ctx.Unique == nil && mstq.path != nil {
		mstq.Unique(true)
	}
	ctx = setContextOp(ctx, mstq.ctx, ent.OpQueryIDs)
	if err = mstq.Select(membersignuptoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := mstq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mstq *MemberSignupTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mstq.ctx, ent.OpQueryCount)
	if err := mstq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mstq, querierCount[*MemberSignupTokenQuery](), mstq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) CountX(ctx context.Context) int {
	count, err := mstq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mstq *MemberSignupTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mstq.ctx, ent.OpQueryExist)
	switch _, err := mstq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mstq *MemberSignupTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := mstq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberSignupTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mstq *MemberSignupTokenQuery) Clone() *MemberSignupTokenQuery {
	if mstq == nil {
		return nil
	}
	return &MemberSignupTokenQuery{
		config:        mstq.config,
		ctx:           mstq.ctx.Clone(),
		order:         append([]membersignuptoken.OrderOption{}, mstq.order...),
		inters:        append([]Interceptor{}, mstq.inters...),
		predicates:    append([]predicate.MemberSignupToken{}, mstq.predicates...),
		withCompany:   mstq.withCompany.Clone(),
		withCreatedBy: mstq.withCreatedBy.Clone(),
		// clone intermediate query.
		sql:       mstq.sql.Clone(),
		path:      mstq.path,
		modifiers: append([]func(*sql.Selector){}, mstq.modifiers...),
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (mstq *MemberSignupTokenQuery) WithCompany(opts ...func(*CompanyQuery)) *MemberSignupTokenQuery {
	query := (&CompanyClient{config: mstq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mstq.withCompany = query
	return mstq
}

// WithCreatedBy tells the query-builder to eager-load the nodes that are connected to
// the "createdBy" edge. The optional arguments are used to configure the query builder of the edge.
func (mstq *MemberSignupTokenQuery) WithCreatedBy(opts ...func(*UserQuery)) *MemberSignupTokenQuery {
	query := (&UserClient{config: mstq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mstq.withCreatedBy = query
	return mstq
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
//	client.MemberSignupToken.Query().
//		GroupBy(membersignuptoken.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (mstq *MemberSignupTokenQuery) GroupBy(field string, fields ...string) *MemberSignupTokenGroupBy {
	mstq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberSignupTokenGroupBy{build: mstq}
	grbuild.flds = &mstq.ctx.Fields
	grbuild.label = membersignuptoken.Label
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
//	client.MemberSignupToken.Query().
//		Select(membersignuptoken.FieldCreatedAt).
//		Scan(ctx, &v)
func (mstq *MemberSignupTokenQuery) Select(fields ...string) *MemberSignupTokenSelect {
	mstq.ctx.Fields = append(mstq.ctx.Fields, fields...)
	sbuild := &MemberSignupTokenSelect{MemberSignupTokenQuery: mstq}
	sbuild.label = membersignuptoken.Label
	sbuild.flds, sbuild.scan = &mstq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberSignupTokenSelect configured with the given aggregations.
func (mstq *MemberSignupTokenQuery) Aggregate(fns ...AggregateFunc) *MemberSignupTokenSelect {
	return mstq.Select().Aggregate(fns...)
}

func (mstq *MemberSignupTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mstq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mstq); err != nil {
				return err
			}
		}
	}
	for _, f := range mstq.ctx.Fields {
		if !membersignuptoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if mstq.path != nil {
		prev, err := mstq.path(ctx)
		if err != nil {
			return err
		}
		mstq.sql = prev
	}
	return nil
}

func (mstq *MemberSignupTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MemberSignupToken, error) {
	var (
		nodes       = []*MemberSignupToken{}
		withFKs     = mstq.withFKs
		_spec       = mstq.querySpec()
		loadedTypes = [2]bool{
			mstq.withCompany != nil,
			mstq.withCreatedBy != nil,
		}
	)
	if mstq.withCompany != nil || mstq.withCreatedBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, membersignuptoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MemberSignupToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MemberSignupToken{config: mstq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mstq.modifiers) > 0 {
		_spec.Modifiers = mstq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mstq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mstq.withCompany; query != nil {
		if err := mstq.loadCompany(ctx, query, nodes, nil,
			func(n *MemberSignupToken, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := mstq.withCreatedBy; query != nil {
		if err := mstq.loadCreatedBy(ctx, query, nodes, nil,
			func(n *MemberSignupToken, e *User) { n.Edges.CreatedBy = e }); err != nil {
			return nil, err
		}
	}
	for i := range mstq.loadTotal {
		if err := mstq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mstq *MemberSignupTokenQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*MemberSignupToken, init func(*MemberSignupToken), assign func(*MemberSignupToken, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MemberSignupToken)
	for i := range nodes {
		if nodes[i].company_member_signup_tokens == nil {
			continue
		}
		fk := *nodes[i].company_member_signup_tokens
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
			return fmt.Errorf(`unexpected foreign-key "company_member_signup_tokens" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mstq *MemberSignupTokenQuery) loadCreatedBy(ctx context.Context, query *UserQuery, nodes []*MemberSignupToken, init func(*MemberSignupToken), assign func(*MemberSignupToken, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MemberSignupToken)
	for i := range nodes {
		if nodes[i].user_created_member_signup_tokens == nil {
			continue
		}
		fk := *nodes[i].user_created_member_signup_tokens
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
			return fmt.Errorf(`unexpected foreign-key "user_created_member_signup_tokens" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mstq *MemberSignupTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mstq.querySpec()
	if len(mstq.modifiers) > 0 {
		_spec.Modifiers = mstq.modifiers
	}
	_spec.Node.Columns = mstq.ctx.Fields
	if len(mstq.ctx.Fields) > 0 {
		_spec.Unique = mstq.ctx.Unique != nil && *mstq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mstq.driver, _spec)
}

func (mstq *MemberSignupTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(membersignuptoken.Table, membersignuptoken.Columns, sqlgraph.NewFieldSpec(membersignuptoken.FieldID, field.TypeInt))
	_spec.From = mstq.sql
	if unique := mstq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mstq.path != nil {
		_spec.Unique = true
	}
	if fields := mstq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membersignuptoken.FieldID)
		for i := range fields {
			if fields[i] != membersignuptoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mstq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mstq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mstq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mstq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mstq *MemberSignupTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mstq.driver.Dialect())
	t1 := builder.Table(membersignuptoken.Table)
	columns := mstq.ctx.Fields
	if len(columns) == 0 {
		columns = membersignuptoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mstq.sql != nil {
		selector = mstq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mstq.ctx.Unique != nil && *mstq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mstq.modifiers {
		m(selector)
	}
	for _, p := range mstq.predicates {
		p(selector)
	}
	for _, p := range mstq.order {
		p(selector)
	}
	if offset := mstq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mstq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mstq *MemberSignupTokenQuery) Modify(modifiers ...func(s *sql.Selector)) *MemberSignupTokenSelect {
	mstq.modifiers = append(mstq.modifiers, modifiers...)
	return mstq.Select()
}

// MemberSignupTokenGroupBy is the group-by builder for MemberSignupToken entities.
type MemberSignupTokenGroupBy struct {
	selector
	build *MemberSignupTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mstgb *MemberSignupTokenGroupBy) Aggregate(fns ...AggregateFunc) *MemberSignupTokenGroupBy {
	mstgb.fns = append(mstgb.fns, fns...)
	return mstgb
}

// Scan applies the selector query and scans the result into the given value.
func (mstgb *MemberSignupTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mstgb.build.ctx, ent.OpQueryGroupBy)
	if err := mstgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberSignupTokenQuery, *MemberSignupTokenGroupBy](ctx, mstgb.build, mstgb, mstgb.build.inters, v)
}

func (mstgb *MemberSignupTokenGroupBy) sqlScan(ctx context.Context, root *MemberSignupTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mstgb.fns))
	for _, fn := range mstgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mstgb.flds)+len(mstgb.fns))
		for _, f := range *mstgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mstgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mstgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberSignupTokenSelect is the builder for selecting fields of MemberSignupToken entities.
type MemberSignupTokenSelect struct {
	*MemberSignupTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (msts *MemberSignupTokenSelect) Aggregate(fns ...AggregateFunc) *MemberSignupTokenSelect {
	msts.fns = append(msts.fns, fns...)
	return msts
}

// Scan applies the selector query and scans the result into the given value.
func (msts *MemberSignupTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, msts.ctx, ent.OpQuerySelect)
	if err := msts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberSignupTokenQuery, *MemberSignupTokenSelect](ctx, msts.MemberSignupTokenQuery, msts, msts.inters, v)
}

func (msts *MemberSignupTokenSelect) sqlScan(ctx context.Context, root *MemberSignupTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(msts.fns))
	for _, fn := range msts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*msts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := msts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (msts *MemberSignupTokenSelect) Modify(modifiers ...func(s *sql.Selector)) *MemberSignupTokenSelect {
	msts.modifiers = append(msts.modifiers, modifiers...)
	return msts
}
