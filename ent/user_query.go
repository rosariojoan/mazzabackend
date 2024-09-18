// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"mazza/ent/accountingentry"
	"mazza/ent/company"
	"mazza/ent/employee"
	"mazza/ent/predicate"
	"mazza/ent/token"
	"mazza/ent/user"
	"mazza/ent/userrole"
	"mazza/ent/worktask"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx                        *QueryContext
	order                      []user.OrderOption
	inters                     []Interceptor
	predicates                 []predicate.User
	withAccountingEntries      *AccountingEntryQuery
	withCompany                *CompanyQuery
	withAssignedRoles          *UserRoleQuery
	withCreatedTasks           *WorktaskQuery
	withEmployee               *EmployeeQuery
	withTokens                 *TokenQuery
	modifiers                  []func(*sql.Selector)
	loadTotal                  []func(context.Context, []*User) error
	withNamedAccountingEntries map[string]*AccountingEntryQuery
	withNamedCompany           map[string]*CompanyQuery
	withNamedAssignedRoles     map[string]*UserRoleQuery
	withNamedCreatedTasks      map[string]*WorktaskQuery
	withNamedTokens            map[string]*TokenQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryAccountingEntries chains the current query on the "accountingEntries" edge.
func (uq *UserQuery) QueryAccountingEntries() *AccountingEntryQuery {
	query := (&AccountingEntryClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(accountingentry.Table, accountingentry.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AccountingEntriesTable, user.AccountingEntriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompany chains the current query on the "company" edge.
func (uq *UserQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.CompanyTable, user.CompanyPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssignedRoles chains the current query on the "assignedRoles" edge.
func (uq *UserQuery) QueryAssignedRoles() *UserRoleQuery {
	query := (&UserRoleClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(userrole.Table, userrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.AssignedRolesTable, user.AssignedRolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedTasks chains the current query on the "createdTasks" edge.
func (uq *UserQuery) QueryCreatedTasks() *WorktaskQuery {
	query := (&WorktaskClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(worktask.Table, worktask.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CreatedTasksTable, user.CreatedTasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmployee chains the current query on the "employee" edge.
func (uq *UserQuery) QueryEmployee() *EmployeeQuery {
	query := (&EmployeeClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.EmployeeTable, user.EmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTokens chains the current query on the "tokens" edge.
func (uq *UserQuery) QueryTokens() *TokenQuery {
	query := (&TokenClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TokensTable, user.TokensColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryAll)
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryIDs)
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryCount)
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryExist)
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:                uq.config,
		ctx:                   uq.ctx.Clone(),
		order:                 append([]user.OrderOption{}, uq.order...),
		inters:                append([]Interceptor{}, uq.inters...),
		predicates:            append([]predicate.User{}, uq.predicates...),
		withAccountingEntries: uq.withAccountingEntries.Clone(),
		withCompany:           uq.withCompany.Clone(),
		withAssignedRoles:     uq.withAssignedRoles.Clone(),
		withCreatedTasks:      uq.withCreatedTasks.Clone(),
		withEmployee:          uq.withEmployee.Clone(),
		withTokens:            uq.withTokens.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithAccountingEntries tells the query-builder to eager-load the nodes that are connected to
// the "accountingEntries" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithAccountingEntries(opts ...func(*AccountingEntryQuery)) *UserQuery {
	query := (&AccountingEntryClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withAccountingEntries = query
	return uq
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCompany(opts ...func(*CompanyQuery)) *UserQuery {
	query := (&CompanyClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCompany = query
	return uq
}

// WithAssignedRoles tells the query-builder to eager-load the nodes that are connected to
// the "assignedRoles" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithAssignedRoles(opts ...func(*UserRoleQuery)) *UserQuery {
	query := (&UserRoleClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withAssignedRoles = query
	return uq
}

// WithCreatedTasks tells the query-builder to eager-load the nodes that are connected to
// the "createdTasks" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCreatedTasks(opts ...func(*WorktaskQuery)) *UserQuery {
	query := (&WorktaskClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCreatedTasks = query
	return uq
}

// WithEmployee tells the query-builder to eager-load the nodes that are connected to
// the "employee" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithEmployee(opts ...func(*EmployeeQuery)) *UserQuery {
	query := (&EmployeeClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withEmployee = query
	return uq
}

// WithTokens tells the query-builder to eager-load the nodes that are connected to
// the "tokens" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithTokens(opts ...func(*TokenQuery)) *UserQuery {
	query := (&TokenClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withTokens = query
	return uq
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
//	client.User.Query().
//		GroupBy(user.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
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
//	client.User.Query().
//		Select(user.FieldCreatedAt).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = uq.querySpec()
		loadedTypes = [6]bool{
			uq.withAccountingEntries != nil,
			uq.withCompany != nil,
			uq.withAssignedRoles != nil,
			uq.withCreatedTasks != nil,
			uq.withEmployee != nil,
			uq.withTokens != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withAccountingEntries; query != nil {
		if err := uq.loadAccountingEntries(ctx, query, nodes,
			func(n *User) { n.Edges.AccountingEntries = []*AccountingEntry{} },
			func(n *User, e *AccountingEntry) { n.Edges.AccountingEntries = append(n.Edges.AccountingEntries, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withCompany; query != nil {
		if err := uq.loadCompany(ctx, query, nodes,
			func(n *User) { n.Edges.Company = []*Company{} },
			func(n *User, e *Company) { n.Edges.Company = append(n.Edges.Company, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withAssignedRoles; query != nil {
		if err := uq.loadAssignedRoles(ctx, query, nodes,
			func(n *User) { n.Edges.AssignedRoles = []*UserRole{} },
			func(n *User, e *UserRole) { n.Edges.AssignedRoles = append(n.Edges.AssignedRoles, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withCreatedTasks; query != nil {
		if err := uq.loadCreatedTasks(ctx, query, nodes,
			func(n *User) { n.Edges.CreatedTasks = []*Worktask{} },
			func(n *User, e *Worktask) { n.Edges.CreatedTasks = append(n.Edges.CreatedTasks, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withEmployee; query != nil {
		if err := uq.loadEmployee(ctx, query, nodes, nil,
			func(n *User, e *Employee) { n.Edges.Employee = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withTokens; query != nil {
		if err := uq.loadTokens(ctx, query, nodes,
			func(n *User) { n.Edges.Tokens = []*Token{} },
			func(n *User, e *Token) { n.Edges.Tokens = append(n.Edges.Tokens, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedAccountingEntries {
		if err := uq.loadAccountingEntries(ctx, query, nodes,
			func(n *User) { n.appendNamedAccountingEntries(name) },
			func(n *User, e *AccountingEntry) { n.appendNamedAccountingEntries(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedCompany {
		if err := uq.loadCompany(ctx, query, nodes,
			func(n *User) { n.appendNamedCompany(name) },
			func(n *User, e *Company) { n.appendNamedCompany(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedAssignedRoles {
		if err := uq.loadAssignedRoles(ctx, query, nodes,
			func(n *User) { n.appendNamedAssignedRoles(name) },
			func(n *User, e *UserRole) { n.appendNamedAssignedRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedCreatedTasks {
		if err := uq.loadCreatedTasks(ctx, query, nodes,
			func(n *User) { n.appendNamedCreatedTasks(name) },
			func(n *User, e *Worktask) { n.appendNamedCreatedTasks(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedTokens {
		if err := uq.loadTokens(ctx, query, nodes,
			func(n *User) { n.appendNamedTokens(name) },
			func(n *User, e *Token) { n.appendNamedTokens(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range uq.loadTotal {
		if err := uq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadAccountingEntries(ctx context.Context, query *AccountingEntryQuery, nodes []*User, init func(*User), assign func(*User, *AccountingEntry)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AccountingEntry(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.AccountingEntriesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_accounting_entries
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_accounting_entries" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_accounting_entries" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*User, init func(*User), assign func(*User, *Company)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*User)
	nids := make(map[int]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.CompanyTable)
		s.Join(joinT).On(s.C(company.FieldID), joinT.C(user.CompanyPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(user.CompanyPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.CompanyPrimaryKey[1]))
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
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Company](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "company" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (uq *UserQuery) loadAssignedRoles(ctx context.Context, query *UserRoleQuery, nodes []*User, init func(*User), assign func(*User, *UserRole)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*User)
	nids := make(map[int]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.AssignedRolesTable)
		s.Join(joinT).On(s.C(userrole.FieldID), joinT.C(user.AssignedRolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(user.AssignedRolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.AssignedRolesPrimaryKey[0]))
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
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*UserRole](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "assignedRoles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (uq *UserQuery) loadCreatedTasks(ctx context.Context, query *WorktaskQuery, nodes []*User, init func(*User), assign func(*User, *Worktask)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Worktask(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.CreatedTasksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_created_tasks
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_created_tasks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_created_tasks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadEmployee(ctx context.Context, query *EmployeeQuery, nodes []*User, init func(*User), assign func(*User, *Employee)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.EmployeeColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_employee
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_employee" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_employee" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadTokens(ctx context.Context, query *TokenQuery, nodes []*User, init func(*User), assign func(*User, *Token)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Token(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.TokensColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_tokens
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_tokens" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_tokens" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAccountingEntries tells the query-builder to eager-load the nodes that are connected to the "accountingEntries"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedAccountingEntries(name string, opts ...func(*AccountingEntryQuery)) *UserQuery {
	query := (&AccountingEntryClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedAccountingEntries == nil {
		uq.withNamedAccountingEntries = make(map[string]*AccountingEntryQuery)
	}
	uq.withNamedAccountingEntries[name] = query
	return uq
}

// WithNamedCompany tells the query-builder to eager-load the nodes that are connected to the "company"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedCompany(name string, opts ...func(*CompanyQuery)) *UserQuery {
	query := (&CompanyClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedCompany == nil {
		uq.withNamedCompany = make(map[string]*CompanyQuery)
	}
	uq.withNamedCompany[name] = query
	return uq
}

// WithNamedAssignedRoles tells the query-builder to eager-load the nodes that are connected to the "assignedRoles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedAssignedRoles(name string, opts ...func(*UserRoleQuery)) *UserQuery {
	query := (&UserRoleClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedAssignedRoles == nil {
		uq.withNamedAssignedRoles = make(map[string]*UserRoleQuery)
	}
	uq.withNamedAssignedRoles[name] = query
	return uq
}

// WithNamedCreatedTasks tells the query-builder to eager-load the nodes that are connected to the "createdTasks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedCreatedTasks(name string, opts ...func(*WorktaskQuery)) *UserQuery {
	query := (&WorktaskClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedCreatedTasks == nil {
		uq.withNamedCreatedTasks = make(map[string]*WorktaskQuery)
	}
	uq.withNamedCreatedTasks[name] = query
	return uq
}

// WithNamedTokens tells the query-builder to eager-load the nodes that are connected to the "tokens"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedTokens(name string, opts ...func(*TokenQuery)) *UserQuery {
	query := (&TokenClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedTokens == nil {
		uq.withNamedTokens = make(map[string]*TokenQuery)
	}
	uq.withNamedTokens[name] = query
	return uq
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}