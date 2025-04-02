// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"mazza/ent/generated/company"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeQuery is the builder for querying Employee entities.
type EmployeeQuery struct {
	config
	ctx                   *QueryContext
	order                 []employee.OrderOption
	inters                []Interceptor
	predicates            []predicate.Employee
	withCompany           *CompanyQuery
	withUser              *UserQuery
	withSubordinates      *EmployeeQuery
	withLeader            *EmployeeQuery
	withFKs               bool
	loadTotal             []func(context.Context, []*Employee) error
	modifiers             []func(*sql.Selector)
	withNamedSubordinates map[string]*EmployeeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeeQuery builder.
func (eq *EmployeeQuery) Where(ps ...predicate.Employee) *EmployeeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EmployeeQuery) Limit(limit int) *EmployeeQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EmployeeQuery) Offset(offset int) *EmployeeQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EmployeeQuery) Unique(unique bool) *EmployeeQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EmployeeQuery) Order(o ...employee.OrderOption) *EmployeeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryCompany chains the current query on the "company" edge.
func (eq *EmployeeQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, employee.CompanyTable, employee.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (eq *EmployeeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, employee.UserTable, employee.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubordinates chains the current query on the "subordinates" edge.
func (eq *EmployeeQuery) QuerySubordinates() *EmployeeQuery {
	query := (&EmployeeClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employee.SubordinatesTable, employee.SubordinatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLeader chains the current query on the "leader" edge.
func (eq *EmployeeQuery) QueryLeader() *EmployeeQuery {
	query := (&EmployeeClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, employee.LeaderTable, employee.LeaderColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Employee entity from the query.
// Returns a *NotFoundError when no Employee was found.
func (eq *EmployeeQuery) First(ctx context.Context) (*Employee, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employee.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EmployeeQuery) FirstX(ctx context.Context) *Employee {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Employee ID from the query.
// Returns a *NotFoundError when no Employee ID was found.
func (eq *EmployeeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employee.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EmployeeQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Employee entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Employee entity is found.
// Returns a *NotFoundError when no Employee entities are found.
func (eq *EmployeeQuery) Only(ctx context.Context) (*Employee, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employee.Label}
	default:
		return nil, &NotSingularError{employee.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EmployeeQuery) OnlyX(ctx context.Context) *Employee {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Employee ID in the query.
// Returns a *NotSingularError when more than one Employee ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EmployeeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employee.Label}
	default:
		err = &NotSingularError{employee.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EmployeeQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Employees.
func (eq *EmployeeQuery) All(ctx context.Context) ([]*Employee, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryAll)
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Employee, *EmployeeQuery]()
	return withInterceptors[[]*Employee](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EmployeeQuery) AllX(ctx context.Context) []*Employee {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Employee IDs.
func (eq *EmployeeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryIDs)
	if err = eq.Select(employee.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EmployeeQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EmployeeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryCount)
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EmployeeQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EmployeeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EmployeeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryExist)
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EmployeeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EmployeeQuery) Clone() *EmployeeQuery {
	if eq == nil {
		return nil
	}
	return &EmployeeQuery{
		config:           eq.config,
		ctx:              eq.ctx.Clone(),
		order:            append([]employee.OrderOption{}, eq.order...),
		inters:           append([]Interceptor{}, eq.inters...),
		predicates:       append([]predicate.Employee{}, eq.predicates...),
		withCompany:      eq.withCompany.Clone(),
		withUser:         eq.withUser.Clone(),
		withSubordinates: eq.withSubordinates.Clone(),
		withLeader:       eq.withLeader.Clone(),
		// clone intermediate query.
		sql:       eq.sql.Clone(),
		path:      eq.path,
		modifiers: append([]func(*sql.Selector){}, eq.modifiers...),
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EmployeeQuery) WithCompany(opts ...func(*CompanyQuery)) *EmployeeQuery {
	query := (&CompanyClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withCompany = query
	return eq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EmployeeQuery) WithUser(opts ...func(*UserQuery)) *EmployeeQuery {
	query := (&UserClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withUser = query
	return eq
}

// WithSubordinates tells the query-builder to eager-load the nodes that are connected to
// the "subordinates" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EmployeeQuery) WithSubordinates(opts ...func(*EmployeeQuery)) *EmployeeQuery {
	query := (&EmployeeClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withSubordinates = query
	return eq
}

// WithLeader tells the query-builder to eager-load the nodes that are connected to
// the "leader" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EmployeeQuery) WithLeader(opts ...func(*EmployeeQuery)) *EmployeeQuery {
	query := (&EmployeeClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withLeader = query
	return eq
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
//	client.Employee.Query().
//		GroupBy(employee.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (eq *EmployeeQuery) GroupBy(field string, fields ...string) *EmployeeGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmployeeGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = employee.Label
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
//	client.Employee.Query().
//		Select(employee.FieldCreatedAt).
//		Scan(ctx, &v)
func (eq *EmployeeQuery) Select(fields ...string) *EmployeeSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EmployeeSelect{EmployeeQuery: eq}
	sbuild.label = employee.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmployeeSelect configured with the given aggregations.
func (eq *EmployeeQuery) Aggregate(fns ...AggregateFunc) *EmployeeSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EmployeeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !employee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EmployeeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Employee, error) {
	var (
		nodes       = []*Employee{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [4]bool{
			eq.withCompany != nil,
			eq.withUser != nil,
			eq.withSubordinates != nil,
			eq.withLeader != nil,
		}
	)
	if eq.withCompany != nil || eq.withUser != nil || eq.withLeader != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, employee.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Employee).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Employee{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withCompany; query != nil {
		if err := eq.loadCompany(ctx, query, nodes, nil,
			func(n *Employee, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withUser; query != nil {
		if err := eq.loadUser(ctx, query, nodes, nil,
			func(n *Employee, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withSubordinates; query != nil {
		if err := eq.loadSubordinates(ctx, query, nodes,
			func(n *Employee) { n.Edges.Subordinates = []*Employee{} },
			func(n *Employee, e *Employee) { n.Edges.Subordinates = append(n.Edges.Subordinates, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withLeader; query != nil {
		if err := eq.loadLeader(ctx, query, nodes, nil,
			func(n *Employee, e *Employee) { n.Edges.Leader = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range eq.withNamedSubordinates {
		if err := eq.loadSubordinates(ctx, query, nodes,
			func(n *Employee) { n.appendNamedSubordinates(name) },
			func(n *Employee, e *Employee) { n.appendNamedSubordinates(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range eq.loadTotal {
		if err := eq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EmployeeQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Employee, init func(*Employee), assign func(*Employee, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Employee)
	for i := range nodes {
		if nodes[i].company_employees == nil {
			continue
		}
		fk := *nodes[i].company_employees
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
			return fmt.Errorf(`unexpected foreign-key "company_employees" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EmployeeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Employee, init func(*Employee), assign func(*Employee, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Employee)
	for i := range nodes {
		if nodes[i].user_employee == nil {
			continue
		}
		fk := *nodes[i].user_employee
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
			return fmt.Errorf(`unexpected foreign-key "user_employee" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EmployeeQuery) loadSubordinates(ctx context.Context, query *EmployeeQuery, nodes []*Employee, init func(*Employee), assign func(*Employee, *Employee)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Employee)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(employee.SubordinatesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.employee_subordinates
		if fk == nil {
			return fmt.Errorf(`foreign-key "employee_subordinates" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "employee_subordinates" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EmployeeQuery) loadLeader(ctx context.Context, query *EmployeeQuery, nodes []*Employee, init func(*Employee), assign func(*Employee, *Employee)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Employee)
	for i := range nodes {
		if nodes[i].employee_subordinates == nil {
			continue
		}
		fk := *nodes[i].employee_subordinates
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(employee.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "employee_subordinates" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eq *EmployeeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EmployeeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for i := range fields {
			if fields[i] != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EmployeeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(employee.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = employee.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range eq.modifiers {
		m(selector)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (eq *EmployeeQuery) Modify(modifiers ...func(s *sql.Selector)) *EmployeeSelect {
	eq.modifiers = append(eq.modifiers, modifiers...)
	return eq.Select()
}

// WithNamedSubordinates tells the query-builder to eager-load the nodes that are connected to the "subordinates"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (eq *EmployeeQuery) WithNamedSubordinates(name string, opts ...func(*EmployeeQuery)) *EmployeeQuery {
	query := (&EmployeeClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if eq.withNamedSubordinates == nil {
		eq.withNamedSubordinates = make(map[string]*EmployeeQuery)
	}
	eq.withNamedSubordinates[name] = query
	return eq
}

// EmployeeGroupBy is the group-by builder for Employee entities.
type EmployeeGroupBy struct {
	selector
	build *EmployeeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EmployeeGroupBy) Aggregate(fns ...AggregateFunc) *EmployeeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EmployeeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, ent.OpQueryGroupBy)
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeQuery, *EmployeeGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EmployeeGroupBy) sqlScan(ctx context.Context, root *EmployeeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmployeeSelect is the builder for selecting fields of Employee entities.
type EmployeeSelect struct {
	*EmployeeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EmployeeSelect) Aggregate(fns ...AggregateFunc) *EmployeeSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EmployeeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, ent.OpQuerySelect)
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeQuery, *EmployeeSelect](ctx, es.EmployeeQuery, es, es.inters, v)
}

func (es *EmployeeSelect) sqlScan(ctx context.Context, root *EmployeeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (es *EmployeeSelect) Modify(modifiers ...func(s *sql.Selector)) *EmployeeSelect {
	es.modifiers = append(es.modifiers, modifiers...)
	return es
}
