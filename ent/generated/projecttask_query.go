// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projecttask"
	"mazza/ent/generated/user"
	"mazza/ent/generated/workshift"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskQuery is the builder for querying ProjectTask entities.
type ProjectTaskQuery struct {
	config
	ctx                   *QueryContext
	order                 []projecttask.OrderOption
	inters                []Interceptor
	predicates            []predicate.ProjectTask
	withProject           *ProjectQuery
	withAssignee          *UserQuery
	withParticipants      *UserQuery
	withCreatedBy         *UserQuery
	withWorkShifts        *WorkshiftQuery
	withFKs               bool
	loadTotal             []func(context.Context, []*ProjectTask) error
	modifiers             []func(*sql.Selector)
	withNamedParticipants map[string]*UserQuery
	withNamedWorkShifts   map[string]*WorkshiftQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectTaskQuery builder.
func (ptq *ProjectTaskQuery) Where(ps ...predicate.ProjectTask) *ProjectTaskQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *ProjectTaskQuery) Limit(limit int) *ProjectTaskQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *ProjectTaskQuery) Offset(offset int) *ProjectTaskQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *ProjectTaskQuery) Unique(unique bool) *ProjectTaskQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *ProjectTaskQuery) Order(o ...projecttask.OrderOption) *ProjectTaskQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryProject chains the current query on the "project" edge.
func (ptq *ProjectTaskQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttask.Table, projecttask.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttask.ProjectTable, projecttask.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssignee chains the current query on the "assignee" edge.
func (ptq *ProjectTaskQuery) QueryAssignee() *UserQuery {
	query := (&UserClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttask.Table, projecttask.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttask.AssigneeTable, projecttask.AssigneeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParticipants chains the current query on the "participants" edge.
func (ptq *ProjectTaskQuery) QueryParticipants() *UserQuery {
	query := (&UserClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttask.Table, projecttask.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, projecttask.ParticipantsTable, projecttask.ParticipantsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedBy chains the current query on the "createdBy" edge.
func (ptq *ProjectTaskQuery) QueryCreatedBy() *UserQuery {
	query := (&UserClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttask.Table, projecttask.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttask.CreatedByTable, projecttask.CreatedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkShifts chains the current query on the "workShifts" edge.
func (ptq *ProjectTaskQuery) QueryWorkShifts() *WorkshiftQuery {
	query := (&WorkshiftClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttask.Table, projecttask.FieldID, selector),
			sqlgraph.To(workshift.Table, workshift.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, projecttask.WorkShiftsTable, projecttask.WorkShiftsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectTask entity from the query.
// Returns a *NotFoundError when no ProjectTask was found.
func (ptq *ProjectTaskQuery) First(ctx context.Context) (*ProjectTask, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projecttask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *ProjectTaskQuery) FirstX(ctx context.Context) *ProjectTask {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectTask ID from the query.
// Returns a *NotFoundError when no ProjectTask ID was found.
func (ptq *ProjectTaskQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projecttask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *ProjectTaskQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProjectTask entity is found.
// Returns a *NotFoundError when no ProjectTask entities are found.
func (ptq *ProjectTaskQuery) Only(ctx context.Context) (*ProjectTask, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projecttask.Label}
	default:
		return nil, &NotSingularError{projecttask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *ProjectTaskQuery) OnlyX(ctx context.Context) *ProjectTask {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectTask ID in the query.
// Returns a *NotSingularError when more than one ProjectTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *ProjectTaskQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projecttask.Label}
	default:
		err = &NotSingularError{projecttask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *ProjectTaskQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectTasks.
func (ptq *ProjectTaskQuery) All(ctx context.Context) ([]*ProjectTask, error) {
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryAll)
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProjectTask, *ProjectTaskQuery]()
	return withInterceptors[[]*ProjectTask](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *ProjectTaskQuery) AllX(ctx context.Context) []*ProjectTask {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectTask IDs.
func (ptq *ProjectTaskQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryIDs)
	if err = ptq.Select(projecttask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *ProjectTaskQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *ProjectTaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryCount)
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*ProjectTaskQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *ProjectTaskQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *ProjectTaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryExist)
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *ProjectTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *ProjectTaskQuery) Clone() *ProjectTaskQuery {
	if ptq == nil {
		return nil
	}
	return &ProjectTaskQuery{
		config:           ptq.config,
		ctx:              ptq.ctx.Clone(),
		order:            append([]projecttask.OrderOption{}, ptq.order...),
		inters:           append([]Interceptor{}, ptq.inters...),
		predicates:       append([]predicate.ProjectTask{}, ptq.predicates...),
		withProject:      ptq.withProject.Clone(),
		withAssignee:     ptq.withAssignee.Clone(),
		withParticipants: ptq.withParticipants.Clone(),
		withCreatedBy:    ptq.withCreatedBy.Clone(),
		withWorkShifts:   ptq.withWorkShifts.Clone(),
		// clone intermediate query.
		sql:       ptq.sql.Clone(),
		path:      ptq.path,
		modifiers: append([]func(*sql.Selector){}, ptq.modifiers...),
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectTaskQuery {
	query := (&ProjectClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withProject = query
	return ptq
}

// WithAssignee tells the query-builder to eager-load the nodes that are connected to
// the "assignee" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithAssignee(opts ...func(*UserQuery)) *ProjectTaskQuery {
	query := (&UserClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withAssignee = query
	return ptq
}

// WithParticipants tells the query-builder to eager-load the nodes that are connected to
// the "participants" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithParticipants(opts ...func(*UserQuery)) *ProjectTaskQuery {
	query := (&UserClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withParticipants = query
	return ptq
}

// WithCreatedBy tells the query-builder to eager-load the nodes that are connected to
// the "createdBy" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithCreatedBy(opts ...func(*UserQuery)) *ProjectTaskQuery {
	query := (&UserClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withCreatedBy = query
	return ptq
}

// WithWorkShifts tells the query-builder to eager-load the nodes that are connected to
// the "workShifts" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithWorkShifts(opts ...func(*WorkshiftQuery)) *ProjectTaskQuery {
	query := (&WorkshiftClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withWorkShifts = query
	return ptq
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
//	client.ProjectTask.Query().
//		GroupBy(projecttask.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ptq *ProjectTaskQuery) GroupBy(field string, fields ...string) *ProjectTaskGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProjectTaskGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = projecttask.Label
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
//	client.ProjectTask.Query().
//		Select(projecttask.FieldCreatedAt).
//		Scan(ctx, &v)
func (ptq *ProjectTaskQuery) Select(fields ...string) *ProjectTaskSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &ProjectTaskSelect{ProjectTaskQuery: ptq}
	sbuild.label = projecttask.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProjectTaskSelect configured with the given aggregations.
func (ptq *ProjectTaskQuery) Aggregate(fns ...AggregateFunc) *ProjectTaskSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *ProjectTaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !projecttask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *ProjectTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProjectTask, error) {
	var (
		nodes       = []*ProjectTask{}
		withFKs     = ptq.withFKs
		_spec       = ptq.querySpec()
		loadedTypes = [5]bool{
			ptq.withProject != nil,
			ptq.withAssignee != nil,
			ptq.withParticipants != nil,
			ptq.withCreatedBy != nil,
			ptq.withWorkShifts != nil,
		}
	)
	if ptq.withProject != nil || ptq.withAssignee != nil || ptq.withCreatedBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, projecttask.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProjectTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProjectTask{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ptq.modifiers) > 0 {
		_spec.Modifiers = ptq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withProject; query != nil {
		if err := ptq.loadProject(ctx, query, nodes, nil,
			func(n *ProjectTask, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := ptq.withAssignee; query != nil {
		if err := ptq.loadAssignee(ctx, query, nodes, nil,
			func(n *ProjectTask, e *User) { n.Edges.Assignee = e }); err != nil {
			return nil, err
		}
	}
	if query := ptq.withParticipants; query != nil {
		if err := ptq.loadParticipants(ctx, query, nodes,
			func(n *ProjectTask) { n.Edges.Participants = []*User{} },
			func(n *ProjectTask, e *User) { n.Edges.Participants = append(n.Edges.Participants, e) }); err != nil {
			return nil, err
		}
	}
	if query := ptq.withCreatedBy; query != nil {
		if err := ptq.loadCreatedBy(ctx, query, nodes, nil,
			func(n *ProjectTask, e *User) { n.Edges.CreatedBy = e }); err != nil {
			return nil, err
		}
	}
	if query := ptq.withWorkShifts; query != nil {
		if err := ptq.loadWorkShifts(ctx, query, nodes,
			func(n *ProjectTask) { n.Edges.WorkShifts = []*Workshift{} },
			func(n *ProjectTask, e *Workshift) { n.Edges.WorkShifts = append(n.Edges.WorkShifts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ptq.withNamedParticipants {
		if err := ptq.loadParticipants(ctx, query, nodes,
			func(n *ProjectTask) { n.appendNamedParticipants(name) },
			func(n *ProjectTask, e *User) { n.appendNamedParticipants(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ptq.withNamedWorkShifts {
		if err := ptq.loadWorkShifts(ctx, query, nodes,
			func(n *ProjectTask) { n.appendNamedWorkShifts(name) },
			func(n *ProjectTask, e *Workshift) { n.appendNamedWorkShifts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ptq.loadTotal {
		if err := ptq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *ProjectTaskQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*ProjectTask, init func(*ProjectTask), assign func(*ProjectTask, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectTask)
	for i := range nodes {
		if nodes[i].project_tasks == nil {
			continue
		}
		fk := *nodes[i].project_tasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ptq *ProjectTaskQuery) loadAssignee(ctx context.Context, query *UserQuery, nodes []*ProjectTask, init func(*ProjectTask), assign func(*ProjectTask, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectTask)
	for i := range nodes {
		if nodes[i].user_assigned_project_tasks == nil {
			continue
		}
		fk := *nodes[i].user_assigned_project_tasks
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
			return fmt.Errorf(`unexpected foreign-key "user_assigned_project_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ptq *ProjectTaskQuery) loadParticipants(ctx context.Context, query *UserQuery, nodes []*ProjectTask, init func(*ProjectTask), assign func(*ProjectTask, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ProjectTask)
	nids := make(map[int]map[*ProjectTask]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(projecttask.ParticipantsTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(projecttask.ParticipantsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(projecttask.ParticipantsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(projecttask.ParticipantsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ProjectTask]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "participants" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ptq *ProjectTaskQuery) loadCreatedBy(ctx context.Context, query *UserQuery, nodes []*ProjectTask, init func(*ProjectTask), assign func(*ProjectTask, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectTask)
	for i := range nodes {
		if nodes[i].user_created_tasks == nil {
			continue
		}
		fk := *nodes[i].user_created_tasks
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
			return fmt.Errorf(`unexpected foreign-key "user_created_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ptq *ProjectTaskQuery) loadWorkShifts(ctx context.Context, query *WorkshiftQuery, nodes []*ProjectTask, init func(*ProjectTask), assign func(*ProjectTask, *Workshift)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ProjectTask)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Workshift(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(projecttask.WorkShiftsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.project_task_work_shifts
		if fk == nil {
			return fmt.Errorf(`foreign-key "project_task_work_shifts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "project_task_work_shifts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ptq *ProjectTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	if len(ptq.modifiers) > 0 {
		_spec.Modifiers = ptq.modifiers
	}
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *ProjectTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(projecttask.Table, projecttask.Columns, sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttask.FieldID)
		for i := range fields {
			if fields[i] != projecttask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *ProjectTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(projecttask.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = projecttask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ptq.modifiers {
		m(selector)
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptq *ProjectTaskQuery) Modify(modifiers ...func(s *sql.Selector)) *ProjectTaskSelect {
	ptq.modifiers = append(ptq.modifiers, modifiers...)
	return ptq.Select()
}

// WithNamedParticipants tells the query-builder to eager-load the nodes that are connected to the "participants"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithNamedParticipants(name string, opts ...func(*UserQuery)) *ProjectTaskQuery {
	query := (&UserClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ptq.withNamedParticipants == nil {
		ptq.withNamedParticipants = make(map[string]*UserQuery)
	}
	ptq.withNamedParticipants[name] = query
	return ptq
}

// WithNamedWorkShifts tells the query-builder to eager-load the nodes that are connected to the "workShifts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTaskQuery) WithNamedWorkShifts(name string, opts ...func(*WorkshiftQuery)) *ProjectTaskQuery {
	query := (&WorkshiftClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ptq.withNamedWorkShifts == nil {
		ptq.withNamedWorkShifts = make(map[string]*WorkshiftQuery)
	}
	ptq.withNamedWorkShifts[name] = query
	return ptq
}

// ProjectTaskGroupBy is the group-by builder for ProjectTask entities.
type ProjectTaskGroupBy struct {
	selector
	build *ProjectTaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *ProjectTaskGroupBy) Aggregate(fns ...AggregateFunc) *ProjectTaskGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *ProjectTaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, ent.OpQueryGroupBy)
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectTaskQuery, *ProjectTaskGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *ProjectTaskGroupBy) sqlScan(ctx context.Context, root *ProjectTaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProjectTaskSelect is the builder for selecting fields of ProjectTask entities.
type ProjectTaskSelect struct {
	*ProjectTaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *ProjectTaskSelect) Aggregate(fns ...AggregateFunc) *ProjectTaskSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *ProjectTaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, ent.OpQuerySelect)
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectTaskQuery, *ProjectTaskSelect](ctx, pts.ProjectTaskQuery, pts, pts.inters, v)
}

func (pts *ProjectTaskSelect) sqlScan(ctx context.Context, root *ProjectTaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pts *ProjectTaskSelect) Modify(modifiers ...func(s *sql.Selector)) *ProjectTaskSelect {
	pts.modifiers = append(pts.modifiers, modifiers...)
	return pts
}
