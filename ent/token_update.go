// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/company"
	"mazza/ent/predicate"
	"mazza/ent/token"
	"mazza/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetExpiry sets the "expiry" field.
func (tu *TokenUpdate) SetExpiry(t time.Time) *TokenUpdate {
	tu.mutation.SetExpiry(t)
	return tu
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableExpiry(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetExpiry(*t)
	}
	return tu
}

// SetCategory sets the "category" field.
func (tu *TokenUpdate) SetCategory(t token.Category) *TokenUpdate {
	tu.mutation.SetCategory(t)
	return tu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableCategory(t *token.Category) *TokenUpdate {
	if t != nil {
		tu.SetCategory(*t)
	}
	return tu
}

// SetToken sets the "token" field.
func (tu *TokenUpdate) SetToken(s string) *TokenUpdate {
	tu.mutation.SetToken(s)
	return tu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableToken(s *string) *TokenUpdate {
	if s != nil {
		tu.SetToken(*s)
	}
	return tu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (tu *TokenUpdate) SetCompanyID(id int) *TokenUpdate {
	tu.mutation.SetCompanyID(id)
	return tu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (tu *TokenUpdate) SetNillableCompanyID(id *int) *TokenUpdate {
	if id != nil {
		tu = tu.SetCompanyID(*id)
	}
	return tu
}

// SetCompany sets the "company" edge to the Company entity.
func (tu *TokenUpdate) SetCompany(c *Company) *TokenUpdate {
	return tu.SetCompanyID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tu *TokenUpdate) SetUserID(id int) *TokenUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tu *TokenUpdate) SetNillableUserID(id *int) *TokenUpdate {
	if id != nil {
		tu = tu.SetUserID(*id)
	}
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TokenUpdate) SetUser(u *User) *TokenUpdate {
	return tu.SetUserID(u.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (tu *TokenUpdate) ClearCompany() *TokenUpdate {
	tu.mutation.ClearCompany()
	return tu
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TokenUpdate) ClearUser() *TokenUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokenUpdate) check() error {
	if v, ok := tu.mutation.Category(); ok {
		if err := token.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Token.category": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Token(); ok {
		if err := token.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Token.token": %w`, err)}
		}
	}
	return nil
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Expiry(); ok {
		_spec.SetField(token.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Category(); ok {
		_spec.SetField(token.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
	}
	if tu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.CompanyTable,
			Columns: []string{token.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.CompanyTable,
			Columns: []string{token.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.UserTable,
			Columns: []string{token.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.UserTable,
			Columns: []string{token.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenMutation
}

// SetExpiry sets the "expiry" field.
func (tuo *TokenUpdateOne) SetExpiry(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetExpiry(t)
	return tuo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableExpiry(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetExpiry(*t)
	}
	return tuo
}

// SetCategory sets the "category" field.
func (tuo *TokenUpdateOne) SetCategory(t token.Category) *TokenUpdateOne {
	tuo.mutation.SetCategory(t)
	return tuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableCategory(t *token.Category) *TokenUpdateOne {
	if t != nil {
		tuo.SetCategory(*t)
	}
	return tuo
}

// SetToken sets the "token" field.
func (tuo *TokenUpdateOne) SetToken(s string) *TokenUpdateOne {
	tuo.mutation.SetToken(s)
	return tuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableToken(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetToken(*s)
	}
	return tuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (tuo *TokenUpdateOne) SetCompanyID(id int) *TokenUpdateOne {
	tuo.mutation.SetCompanyID(id)
	return tuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableCompanyID(id *int) *TokenUpdateOne {
	if id != nil {
		tuo = tuo.SetCompanyID(*id)
	}
	return tuo
}

// SetCompany sets the "company" edge to the Company entity.
func (tuo *TokenUpdateOne) SetCompany(c *Company) *TokenUpdateOne {
	return tuo.SetCompanyID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tuo *TokenUpdateOne) SetUserID(id int) *TokenUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUserID(id *int) *TokenUpdateOne {
	if id != nil {
		tuo = tuo.SetUserID(*id)
	}
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TokenUpdateOne) SetUser(u *User) *TokenUpdateOne {
	return tuo.SetUserID(u.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (tuo *TokenUpdateOne) ClearCompany() *TokenUpdateOne {
	tuo.mutation.ClearCompany()
	return tuo
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TokenUpdateOne) ClearUser() *TokenUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Where appends a list predicates to the TokenUpdate builder.
func (tuo *TokenUpdateOne) Where(ps ...predicate.Token) *TokenUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokenUpdateOne) check() error {
	if v, ok := tuo.mutation.Category(); ok {
		if err := token.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Token.category": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Token(); ok {
		if err := token.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Token.token": %w`, err)}
		}
	}
	return nil
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Expiry(); ok {
		_spec.SetField(token.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Category(); ok {
		_spec.SetField(token.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
	}
	if tuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.CompanyTable,
			Columns: []string{token.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.CompanyTable,
			Columns: []string{token.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.UserTable,
			Columns: []string{token.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.UserTable,
			Columns: []string{token.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}