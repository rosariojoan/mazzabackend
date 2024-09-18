// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/company"
	"mazza/ent/user"
	"mazza/ent/userrole"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRoleCreate is the builder for creating a UserRole entity.
type UserRoleCreate struct {
	config
	mutation *UserRoleMutation
	hooks    []Hook
}

// SetRole sets the "role" field.
func (urc *UserRoleCreate) SetRole(u userrole.Role) *UserRoleCreate {
	urc.mutation.SetRole(u)
	return urc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (urc *UserRoleCreate) SetCompanyID(id int) *UserRoleCreate {
	urc.mutation.SetCompanyID(id)
	return urc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (urc *UserRoleCreate) SetNillableCompanyID(id *int) *UserRoleCreate {
	if id != nil {
		urc = urc.SetCompanyID(*id)
	}
	return urc
}

// SetCompany sets the "company" edge to the Company entity.
func (urc *UserRoleCreate) SetCompany(c *Company) *UserRoleCreate {
	return urc.SetCompanyID(c.ID)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (urc *UserRoleCreate) AddUserIDs(ids ...int) *UserRoleCreate {
	urc.mutation.AddUserIDs(ids...)
	return urc
}

// AddUser adds the "user" edges to the User entity.
func (urc *UserRoleCreate) AddUser(u ...*User) *UserRoleCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return urc.AddUserIDs(ids...)
}

// Mutation returns the UserRoleMutation object of the builder.
func (urc *UserRoleCreate) Mutation() *UserRoleMutation {
	return urc.mutation
}

// Save creates the UserRole in the database.
func (urc *UserRoleCreate) Save(ctx context.Context) (*UserRole, error) {
	return withHooks(ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRoleCreate) SaveX(ctx context.Context) *UserRole {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRoleCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRoleCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRoleCreate) check() error {
	if _, ok := urc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "UserRole.role"`)}
	}
	if v, ok := urc.mutation.Role(); ok {
		if err := userrole.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "UserRole.role": %w`, err)}
		}
	}
	return nil
}

func (urc *UserRoleCreate) sqlSave(ctx context.Context) (*UserRole, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UserRoleCreate) createSpec() (*UserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRole{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(userrole.Table, sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeInt))
	)
	if value, ok := urc.mutation.Role(); ok {
		_spec.SetField(userrole.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := urc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.CompanyTable,
			Columns: []string{userrole.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_available_roles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := urc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   userrole.UserTable,
			Columns: userrole.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserRoleCreateBulk is the builder for creating many UserRole entities in bulk.
type UserRoleCreateBulk struct {
	config
	err      error
	builders []*UserRoleCreate
}

// Save creates the UserRole entities in the database.
func (urcb *UserRoleCreateBulk) Save(ctx context.Context) ([]*UserRole, error) {
	if urcb.err != nil {
		return nil, urcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRole, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) SaveX(ctx context.Context) []*UserRole {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
