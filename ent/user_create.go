// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/accountingentry"
	"mazza/ent/company"
	"mazza/ent/employee"
	"mazza/ent/token"
	"mazza/ent/user"
	"mazza/ent/userrole"
	"mazza/ent/worktask"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updatedAt" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeletedAt sets the "deletedAt" field.
func (uc *UserCreate) SetDeletedAt(t time.Time) *UserCreate {
	uc.mutation.SetDeletedAt(t)
	return uc
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeletedAt(*t)
	}
	return uc
}

// SetFcmToken sets the "fcmToken" field.
func (uc *UserCreate) SetFcmToken(s string) *UserCreate {
	uc.mutation.SetFcmToken(s)
	return uc
}

// SetNillableFcmToken sets the "fcmToken" field if the given value is not nil.
func (uc *UserCreate) SetNillableFcmToken(s *string) *UserCreate {
	if s != nil {
		uc.SetFcmToken(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetDisabled sets the "disabled" field.
func (uc *UserCreate) SetDisabled(b bool) *UserCreate {
	uc.mutation.SetDisabled(b)
	return uc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisabled(b *bool) *UserCreate {
	if b != nil {
		uc.SetDisabled(*b)
	}
	return uc
}

// SetNotVerified sets the "notVerified" field.
func (uc *UserCreate) SetNotVerified(b bool) *UserCreate {
	uc.mutation.SetNotVerified(b)
	return uc
}

// SetNillableNotVerified sets the "notVerified" field if the given value is not nil.
func (uc *UserCreate) SetNillableNotVerified(b *bool) *UserCreate {
	if b != nil {
		uc.SetNotVerified(*b)
	}
	return uc
}

// AddAccountingEntryIDs adds the "accountingEntries" edge to the AccountingEntry entity by IDs.
func (uc *UserCreate) AddAccountingEntryIDs(ids ...int) *UserCreate {
	uc.mutation.AddAccountingEntryIDs(ids...)
	return uc
}

// AddAccountingEntries adds the "accountingEntries" edges to the AccountingEntry entity.
func (uc *UserCreate) AddAccountingEntries(a ...*AccountingEntry) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountingEntryIDs(ids...)
}

// AddCompanyIDs adds the "company" edge to the Company entity by IDs.
func (uc *UserCreate) AddCompanyIDs(ids ...int) *UserCreate {
	uc.mutation.AddCompanyIDs(ids...)
	return uc
}

// AddCompany adds the "company" edges to the Company entity.
func (uc *UserCreate) AddCompany(c ...*Company) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCompanyIDs(ids...)
}

// AddAssignedRoleIDs adds the "assignedRoles" edge to the UserRole entity by IDs.
func (uc *UserCreate) AddAssignedRoleIDs(ids ...int) *UserCreate {
	uc.mutation.AddAssignedRoleIDs(ids...)
	return uc
}

// AddAssignedRoles adds the "assignedRoles" edges to the UserRole entity.
func (uc *UserCreate) AddAssignedRoles(u ...*UserRole) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddAssignedRoleIDs(ids...)
}

// AddCreatedTaskIDs adds the "createdTasks" edge to the Worktask entity by IDs.
func (uc *UserCreate) AddCreatedTaskIDs(ids ...int) *UserCreate {
	uc.mutation.AddCreatedTaskIDs(ids...)
	return uc
}

// AddCreatedTasks adds the "createdTasks" edges to the Worktask entity.
func (uc *UserCreate) AddCreatedTasks(w ...*Worktask) *UserCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddCreatedTaskIDs(ids...)
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (uc *UserCreate) SetEmployeeID(id int) *UserCreate {
	uc.mutation.SetEmployeeID(id)
	return uc
}

// SetNillableEmployeeID sets the "employee" edge to the Employee entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableEmployeeID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetEmployeeID(*id)
	}
	return uc
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (uc *UserCreate) SetEmployee(e *Employee) *UserCreate {
	return uc.SetEmployeeID(e.ID)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uc *UserCreate) AddTokenIDs(ids ...int) *UserCreate {
	uc.mutation.AddTokenIDs(ids...)
	return uc
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uc *UserCreate) AddTokens(t ...*Token) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Disabled(); !ok {
		v := user.DefaultDisabled
		uc.mutation.SetDisabled(v)
	}
	if _, ok := uc.mutation.NotVerified(); !ok {
		v := user.DefaultNotVerified
		uc.mutation.SetNotVerified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "User.createdAt"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "User.updatedAt"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if len(uc.mutation.CompanyIDs()) == 0 {
		return &ValidationError{Name: "company", err: errors.New(`ent: missing required edge "User.company"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := uc.mutation.FcmToken(); ok {
		_spec.SetField(user.FieldFcmToken, field.TypeString, value)
		_node.FcmToken = &value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = &value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Disabled(); ok {
		_spec.SetField(user.FieldDisabled, field.TypeBool, value)
		_node.Disabled = &value
	}
	if value, ok := uc.mutation.NotVerified(); ok {
		_spec.SetField(user.FieldNotVerified, field.TypeBool, value)
		_node.NotVerified = &value
	}
	if nodes := uc.mutation.AccountingEntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountingEntriesTable,
			Columns: []string{user.AccountingEntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountingentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CompanyTable,
			Columns: user.CompanyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AssignedRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AssignedRolesTable,
			Columns: user.AssignedRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CreatedTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreatedTasksTable,
			Columns: []string{user.CreatedTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.EmployeeTable,
			Columns: []string{user.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}