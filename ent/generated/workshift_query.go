// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"mazza/ent/generated/company"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/workshift"
	"mazza/ent/generated/worktask"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkshiftQuery is the builder for querying Workshift entities.
type WorkshiftQuery struct {
	config
	ctx             *QueryContext
	order           []workshift.OrderOption
	inters          []Interceptor
	predicates      []predicate.Workshift
	withCompany     *CompanyQuery
	withEmployee    *EmployeeQuery
	withApprovedBy  *EmployeeQuery
	withWorkTask    *WorktaskQuery
	withEditRequest *WorkshiftQuery
	withWorkShift   *WorkshiftQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*Workshift) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkshiftQuery builder.
func (wq *WorkshiftQuery) Where(ps ...predicate.Workshift) *WorkshiftQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WorkshiftQuery) Limit(limit int) *WorkshiftQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WorkshiftQuery) Offset(offset int) *WorkshiftQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WorkshiftQuery) Unique(unique bool) *WorkshiftQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WorkshiftQuery) Order(o ...workshift.OrderOption) *WorkshiftQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryCompany chains the current query on the "company" edge.
func (wq *WorkshiftQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workshift.Table, workshift.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workshift.CompanyTable, workshift.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmployee chains the current query on the "employee" edge.
func (wq *WorkshiftQuery) QueryEmployee() *EmployeeQuery {
	query := (&EmployeeClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workshift.Table, workshift.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workshift.EmployeeTable, workshift.EmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApprovedBy chains the current query on the "approvedBy" edge.
func (wq *WorkshiftQuery) QueryApprovedBy() *EmployeeQuery {
	query := (&EmployeeClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workshift.Table, workshift.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workshift.ApprovedByTable, workshift.ApprovedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkTask chains the current query on the "workTask" edge.
func (wq *WorkshiftQuery) QueryWorkTask() *WorktaskQuery {
	query := (&WorktaskClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workshift.Table, workshift.FieldID, selector),
			sqlgraph.To(worktask.Table, worktask.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workshift.WorkTaskTable, workshift.WorkTaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEditRequest chains the current query on the "editRequest" edge.
func (wq *WorkshiftQuery) QueryEditRequest() *WorkshiftQuery {
	query := (&WorkshiftClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workshift.Table, workshift.FieldID, selector),
			sqlgraph.To(workshift.Table, workshift.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, workshift.EditRequestTable, workshift.EditRequestColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkShift chains the current query on the "workShift" edge.
func (wq *WorkshiftQuery) QueryWorkShift() *WorkshiftQuery {
	query := (&WorkshiftClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workshift.Table, workshift.FieldID, selector),
			sqlgraph.To(workshift.Table, workshift.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, workshift.WorkShiftTable, workshift.WorkShiftColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Workshift entity from the query.
// Returns a *NotFoundError when no Workshift was found.
func (wq *WorkshiftQuery) First(ctx context.Context) (*Workshift, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workshift.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkshiftQuery) FirstX(ctx context.Context) *Workshift {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Workshift ID from the query.
// Returns a *NotFoundError when no Workshift ID was found.
func (wq *WorkshiftQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workshift.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WorkshiftQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Workshift entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Workshift entity is found.
// Returns a *NotFoundError when no Workshift entities are found.
func (wq *WorkshiftQuery) Only(ctx context.Context) (*Workshift, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workshift.Label}
	default:
		return nil, &NotSingularError{workshift.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkshiftQuery) OnlyX(ctx context.Context) *Workshift {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Workshift ID in the query.
// Returns a *NotSingularError when more than one Workshift ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WorkshiftQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workshift.Label}
	default:
		err = &NotSingularError{workshift.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkshiftQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workshifts.
func (wq *WorkshiftQuery) All(ctx context.Context) ([]*Workshift, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryAll)
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Workshift, *WorkshiftQuery]()
	return withInterceptors[[]*Workshift](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkshiftQuery) AllX(ctx context.Context) []*Workshift {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Workshift IDs.
func (wq *WorkshiftQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryIDs)
	if err = wq.Select(workshift.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkshiftQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkshiftQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryCount)
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WorkshiftQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkshiftQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkshiftQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryExist)
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkshiftQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkshiftQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkshiftQuery) Clone() *WorkshiftQuery {
	if wq == nil {
		return nil
	}
	return &WorkshiftQuery{
		config:          wq.config,
		ctx:             wq.ctx.Clone(),
		order:           append([]workshift.OrderOption{}, wq.order...),
		inters:          append([]Interceptor{}, wq.inters...),
		predicates:      append([]predicate.Workshift{}, wq.predicates...),
		withCompany:     wq.withCompany.Clone(),
		withEmployee:    wq.withEmployee.Clone(),
		withApprovedBy:  wq.withApprovedBy.Clone(),
		withWorkTask:    wq.withWorkTask.Clone(),
		withEditRequest: wq.withEditRequest.Clone(),
		withWorkShift:   wq.withWorkShift.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkshiftQuery) WithCompany(opts ...func(*CompanyQuery)) *WorkshiftQuery {
	query := (&CompanyClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withCompany = query
	return wq
}

// WithEmployee tells the query-builder to eager-load the nodes that are connected to
// the "employee" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkshiftQuery) WithEmployee(opts ...func(*EmployeeQuery)) *WorkshiftQuery {
	query := (&EmployeeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withEmployee = query
	return wq
}

// WithApprovedBy tells the query-builder to eager-load the nodes that are connected to
// the "approvedBy" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkshiftQuery) WithApprovedBy(opts ...func(*EmployeeQuery)) *WorkshiftQuery {
	query := (&EmployeeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withApprovedBy = query
	return wq
}

// WithWorkTask tells the query-builder to eager-load the nodes that are connected to
// the "workTask" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkshiftQuery) WithWorkTask(opts ...func(*WorktaskQuery)) *WorkshiftQuery {
	query := (&WorktaskClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkTask = query
	return wq
}

// WithEditRequest tells the query-builder to eager-load the nodes that are connected to
// the "editRequest" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkshiftQuery) WithEditRequest(opts ...func(*WorkshiftQuery)) *WorkshiftQuery {
	query := (&WorkshiftClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withEditRequest = query
	return wq
}

// WithWorkShift tells the query-builder to eager-load the nodes that are connected to
// the "workShift" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkshiftQuery) WithWorkShift(opts ...func(*WorkshiftQuery)) *WorkshiftQuery {
	query := (&WorkshiftClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkShift = query
	return wq
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
//	client.Workshift.Query().
//		GroupBy(workshift.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (wq *WorkshiftQuery) GroupBy(field string, fields ...string) *WorkshiftGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkshiftGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = workshift.Label
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
//	client.Workshift.Query().
//		Select(workshift.FieldCreatedAt).
//		Scan(ctx, &v)
func (wq *WorkshiftQuery) Select(fields ...string) *WorkshiftSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WorkshiftSelect{WorkshiftQuery: wq}
	sbuild.label = workshift.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkshiftSelect configured with the given aggregations.
func (wq *WorkshiftQuery) Aggregate(fns ...AggregateFunc) *WorkshiftSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WorkshiftQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !workshift.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WorkshiftQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Workshift, error) {
	var (
		nodes       = []*Workshift{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [6]bool{
			wq.withCompany != nil,
			wq.withEmployee != nil,
			wq.withApprovedBy != nil,
			wq.withWorkTask != nil,
			wq.withEditRequest != nil,
			wq.withWorkShift != nil,
		}
	)
	if wq.withCompany != nil || wq.withEmployee != nil || wq.withApprovedBy != nil || wq.withWorkTask != nil || wq.withEditRequest != nil || wq.withWorkShift != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workshift.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Workshift).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Workshift{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withCompany; query != nil {
		if err := wq.loadCompany(ctx, query, nodes, nil,
			func(n *Workshift, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withEmployee; query != nil {
		if err := wq.loadEmployee(ctx, query, nodes, nil,
			func(n *Workshift, e *Employee) { n.Edges.Employee = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withApprovedBy; query != nil {
		if err := wq.loadApprovedBy(ctx, query, nodes, nil,
			func(n *Workshift, e *Employee) { n.Edges.ApprovedBy = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkTask; query != nil {
		if err := wq.loadWorkTask(ctx, query, nodes, nil,
			func(n *Workshift, e *Worktask) { n.Edges.WorkTask = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withEditRequest; query != nil {
		if err := wq.loadEditRequest(ctx, query, nodes, nil,
			func(n *Workshift, e *Workshift) { n.Edges.EditRequest = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkShift; query != nil {
		if err := wq.loadWorkShift(ctx, query, nodes, nil,
			func(n *Workshift, e *Workshift) { n.Edges.WorkShift = e }); err != nil {
			return nil, err
		}
	}
	for i := range wq.loadTotal {
		if err := wq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WorkshiftQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Workshift, init func(*Workshift), assign func(*Workshift, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Workshift)
	for i := range nodes {
		if nodes[i].company_work_shifts == nil {
			continue
		}
		fk := *nodes[i].company_work_shifts
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
			return fmt.Errorf(`unexpected foreign-key "company_work_shifts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkshiftQuery) loadEmployee(ctx context.Context, query *EmployeeQuery, nodes []*Workshift, init func(*Workshift), assign func(*Workshift, *Employee)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Workshift)
	for i := range nodes {
		if nodes[i].employee_work_shifts == nil {
			continue
		}
		fk := *nodes[i].employee_work_shifts
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
			return fmt.Errorf(`unexpected foreign-key "employee_work_shifts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkshiftQuery) loadApprovedBy(ctx context.Context, query *EmployeeQuery, nodes []*Workshift, init func(*Workshift), assign func(*Workshift, *Employee)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Workshift)
	for i := range nodes {
		if nodes[i].employee_approved_work_shifts == nil {
			continue
		}
		fk := *nodes[i].employee_approved_work_shifts
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
			return fmt.Errorf(`unexpected foreign-key "employee_approved_work_shifts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkshiftQuery) loadWorkTask(ctx context.Context, query *WorktaskQuery, nodes []*Workshift, init func(*Workshift), assign func(*Workshift, *Worktask)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Workshift)
	for i := range nodes {
		if nodes[i].worktask_work_shifts == nil {
			continue
		}
		fk := *nodes[i].worktask_work_shifts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(worktask.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "worktask_work_shifts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkshiftQuery) loadEditRequest(ctx context.Context, query *WorkshiftQuery, nodes []*Workshift, init func(*Workshift), assign func(*Workshift, *Workshift)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Workshift)
	for i := range nodes {
		if nodes[i].workshift_edit_request == nil {
			continue
		}
		fk := *nodes[i].workshift_edit_request
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workshift.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workshift_edit_request" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkshiftQuery) loadWorkShift(ctx context.Context, query *WorkshiftQuery, nodes []*Workshift, init func(*Workshift), assign func(*Workshift, *Workshift)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Workshift)
	for i := range nodes {
		if nodes[i].workshift_edit_request == nil {
			continue
		}
		fk := *nodes[i].workshift_edit_request
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workshift.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workshift_edit_request" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wq *WorkshiftQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkshiftQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workshift.Table, workshift.Columns, sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workshift.FieldID)
		for i := range fields {
			if fields[i] != workshift.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkshiftQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(workshift.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = workshift.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkshiftGroupBy is the group-by builder for Workshift entities.
type WorkshiftGroupBy struct {
	selector
	build *WorkshiftQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkshiftGroupBy) Aggregate(fns ...AggregateFunc) *WorkshiftGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WorkshiftGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, ent.OpQueryGroupBy)
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkshiftQuery, *WorkshiftGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WorkshiftGroupBy) sqlScan(ctx context.Context, root *WorkshiftQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkshiftSelect is the builder for selecting fields of Workshift entities.
type WorkshiftSelect struct {
	*WorkshiftQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WorkshiftSelect) Aggregate(fns ...AggregateFunc) *WorkshiftSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WorkshiftSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, ent.OpQuerySelect)
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkshiftQuery, *WorkshiftSelect](ctx, ws.WorkshiftQuery, ws, ws.inters, v)
}

func (ws *WorkshiftSelect) sqlScan(ctx context.Context, root *WorkshiftQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}