// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"mazza/ent/generated/company"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projectmilestone"
	"mazza/ent/generated/projecttask"
	"mazza/ent/generated/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectQuery is the builder for querying Project entities.
type ProjectQuery struct {
	config
	ctx                 *QueryContext
	order               []project.OrderOption
	inters              []Interceptor
	predicates          []predicate.Project
	withCompany         *CompanyQuery
	withCreatedBy       *UserQuery
	withLeader          *UserQuery
	withTasks           *ProjectTaskQuery
	withMilestones      *ProjectMilestoneQuery
	withFKs             bool
	loadTotal           []func(context.Context, []*Project) error
	modifiers           []func(*sql.Selector)
	withNamedTasks      map[string]*ProjectTaskQuery
	withNamedMilestones map[string]*ProjectMilestoneQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectQuery builder.
func (pq *ProjectQuery) Where(ps ...predicate.Project) *ProjectQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProjectQuery) Limit(limit int) *ProjectQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProjectQuery) Offset(offset int) *ProjectQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProjectQuery) Unique(unique bool) *ProjectQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProjectQuery) Order(o ...project.OrderOption) *ProjectQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryCompany chains the current query on the "company" edge.
func (pq *ProjectQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, project.CompanyTable, project.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedBy chains the current query on the "createdBy" edge.
func (pq *ProjectQuery) QueryCreatedBy() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, project.CreatedByTable, project.CreatedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLeader chains the current query on the "leader" edge.
func (pq *ProjectQuery) QueryLeader() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, project.LeaderTable, project.LeaderColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTasks chains the current query on the "tasks" edge.
func (pq *ProjectQuery) QueryTasks() *ProjectTaskQuery {
	query := (&ProjectTaskClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, selector),
			sqlgraph.To(projecttask.Table, projecttask.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, project.TasksTable, project.TasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMilestones chains the current query on the "milestones" edge.
func (pq *ProjectQuery) QueryMilestones() *ProjectMilestoneQuery {
	query := (&ProjectMilestoneClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, selector),
			sqlgraph.To(projectmilestone.Table, projectmilestone.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, project.MilestonesTable, project.MilestonesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Project entity from the query.
// Returns a *NotFoundError when no Project was found.
func (pq *ProjectQuery) First(ctx context.Context) (*Project, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{project.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProjectQuery) FirstX(ctx context.Context) *Project {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Project ID from the query.
// Returns a *NotFoundError when no Project ID was found.
func (pq *ProjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{project.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProjectQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Project entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Project entity is found.
// Returns a *NotFoundError when no Project entities are found.
func (pq *ProjectQuery) Only(ctx context.Context) (*Project, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{project.Label}
	default:
		return nil, &NotSingularError{project.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProjectQuery) OnlyX(ctx context.Context) *Project {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Project ID in the query.
// Returns a *NotSingularError when more than one Project ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = &NotSingularError{project.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Projects.
func (pq *ProjectQuery) All(ctx context.Context) ([]*Project, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Project, *ProjectQuery]()
	return withInterceptors[[]*Project](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProjectQuery) AllX(ctx context.Context) []*Project {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Project IDs.
func (pq *ProjectQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(project.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProjectQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProjectQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProjectQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProjectQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProjectQuery) Clone() *ProjectQuery {
	if pq == nil {
		return nil
	}
	return &ProjectQuery{
		config:         pq.config,
		ctx:            pq.ctx.Clone(),
		order:          append([]project.OrderOption{}, pq.order...),
		inters:         append([]Interceptor{}, pq.inters...),
		predicates:     append([]predicate.Project{}, pq.predicates...),
		withCompany:    pq.withCompany.Clone(),
		withCreatedBy:  pq.withCreatedBy.Clone(),
		withLeader:     pq.withLeader.Clone(),
		withTasks:      pq.withTasks.Clone(),
		withMilestones: pq.withMilestones.Clone(),
		// clone intermediate query.
		sql:       pq.sql.Clone(),
		path:      pq.path,
		modifiers: append([]func(*sql.Selector){}, pq.modifiers...),
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithCompany(opts ...func(*CompanyQuery)) *ProjectQuery {
	query := (&CompanyClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCompany = query
	return pq
}

// WithCreatedBy tells the query-builder to eager-load the nodes that are connected to
// the "createdBy" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithCreatedBy(opts ...func(*UserQuery)) *ProjectQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCreatedBy = query
	return pq
}

// WithLeader tells the query-builder to eager-load the nodes that are connected to
// the "leader" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithLeader(opts ...func(*UserQuery)) *ProjectQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withLeader = query
	return pq
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithTasks(opts ...func(*ProjectTaskQuery)) *ProjectQuery {
	query := (&ProjectTaskClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withTasks = query
	return pq
}

// WithMilestones tells the query-builder to eager-load the nodes that are connected to
// the "milestones" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithMilestones(opts ...func(*ProjectMilestoneQuery)) *ProjectQuery {
	query := (&ProjectMilestoneClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withMilestones = query
	return pq
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
//	client.Project.Query().
//		GroupBy(project.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (pq *ProjectQuery) GroupBy(field string, fields ...string) *ProjectGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProjectGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = project.Label
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
//	client.Project.Query().
//		Select(project.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *ProjectQuery) Select(fields ...string) *ProjectSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProjectSelect{ProjectQuery: pq}
	sbuild.label = project.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProjectSelect configured with the given aggregations.
func (pq *ProjectQuery) Aggregate(fns ...AggregateFunc) *ProjectSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !project.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Project, error) {
	var (
		nodes       = []*Project{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [5]bool{
			pq.withCompany != nil,
			pq.withCreatedBy != nil,
			pq.withLeader != nil,
			pq.withTasks != nil,
			pq.withMilestones != nil,
		}
	)
	if pq.withCompany != nil || pq.withCreatedBy != nil || pq.withLeader != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, project.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Project).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Project{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withCompany; query != nil {
		if err := pq.loadCompany(ctx, query, nodes, nil,
			func(n *Project, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withCreatedBy; query != nil {
		if err := pq.loadCreatedBy(ctx, query, nodes, nil,
			func(n *Project, e *User) { n.Edges.CreatedBy = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withLeader; query != nil {
		if err := pq.loadLeader(ctx, query, nodes, nil,
			func(n *Project, e *User) { n.Edges.Leader = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withTasks; query != nil {
		if err := pq.loadTasks(ctx, query, nodes,
			func(n *Project) { n.Edges.Tasks = []*ProjectTask{} },
			func(n *Project, e *ProjectTask) { n.Edges.Tasks = append(n.Edges.Tasks, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withMilestones; query != nil {
		if err := pq.loadMilestones(ctx, query, nodes,
			func(n *Project) { n.Edges.Milestones = []*ProjectMilestone{} },
			func(n *Project, e *ProjectMilestone) { n.Edges.Milestones = append(n.Edges.Milestones, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedTasks {
		if err := pq.loadTasks(ctx, query, nodes,
			func(n *Project) { n.appendNamedTasks(name) },
			func(n *Project, e *ProjectTask) { n.appendNamedTasks(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedMilestones {
		if err := pq.loadMilestones(ctx, query, nodes,
			func(n *Project) { n.appendNamedMilestones(name) },
			func(n *Project, e *ProjectMilestone) { n.appendNamedMilestones(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProjectQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Project, init func(*Project), assign func(*Project, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Project)
	for i := range nodes {
		if nodes[i].company_projects == nil {
			continue
		}
		fk := *nodes[i].company_projects
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
			return fmt.Errorf(`unexpected foreign-key "company_projects" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProjectQuery) loadCreatedBy(ctx context.Context, query *UserQuery, nodes []*Project, init func(*Project), assign func(*Project, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Project)
	for i := range nodes {
		if nodes[i].user_created_projects == nil {
			continue
		}
		fk := *nodes[i].user_created_projects
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
			return fmt.Errorf(`unexpected foreign-key "user_created_projects" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProjectQuery) loadLeader(ctx context.Context, query *UserQuery, nodes []*Project, init func(*Project), assign func(*Project, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Project)
	for i := range nodes {
		if nodes[i].user_leadered_projects == nil {
			continue
		}
		fk := *nodes[i].user_leadered_projects
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
			return fmt.Errorf(`unexpected foreign-key "user_leadered_projects" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProjectQuery) loadTasks(ctx context.Context, query *ProjectTaskQuery, nodes []*Project, init func(*Project), assign func(*Project, *ProjectTask)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Project)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ProjectTask(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(project.TasksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.project_tasks
		if fk == nil {
			return fmt.Errorf(`foreign-key "project_tasks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "project_tasks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProjectQuery) loadMilestones(ctx context.Context, query *ProjectMilestoneQuery, nodes []*Project, init func(*Project), assign func(*Project, *ProjectMilestone)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Project)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ProjectMilestone(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(project.MilestonesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.project_milestones
		if fk == nil {
			return fmt.Errorf(`foreign-key "project_milestones" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "project_milestones" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *ProjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for i := range fields {
			if fields[i] != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(project.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = project.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *ProjectQuery) Modify(modifiers ...func(s *sql.Selector)) *ProjectSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// WithNamedTasks tells the query-builder to eager-load the nodes that are connected to the "tasks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithNamedTasks(name string, opts ...func(*ProjectTaskQuery)) *ProjectQuery {
	query := (&ProjectTaskClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedTasks == nil {
		pq.withNamedTasks = make(map[string]*ProjectTaskQuery)
	}
	pq.withNamedTasks[name] = query
	return pq
}

// WithNamedMilestones tells the query-builder to eager-load the nodes that are connected to the "milestones"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *ProjectQuery) WithNamedMilestones(name string, opts ...func(*ProjectMilestoneQuery)) *ProjectQuery {
	query := (&ProjectMilestoneClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedMilestones == nil {
		pq.withNamedMilestones = make(map[string]*ProjectMilestoneQuery)
	}
	pq.withNamedMilestones[name] = query
	return pq
}

// ProjectGroupBy is the group-by builder for Project entities.
type ProjectGroupBy struct {
	selector
	build *ProjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProjectGroupBy) Aggregate(fns ...AggregateFunc) *ProjectGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectQuery, *ProjectGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProjectGroupBy) sqlScan(ctx context.Context, root *ProjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProjectSelect is the builder for selecting fields of Project entities.
type ProjectSelect struct {
	*ProjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProjectSelect) Aggregate(fns ...AggregateFunc) *ProjectSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectQuery, *ProjectSelect](ctx, ps.ProjectQuery, ps, ps.inters, v)
}

func (ps *ProjectSelect) sqlScan(ctx context.Context, root *ProjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *ProjectSelect) Modify(modifiers ...func(s *sql.Selector)) *ProjectSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
