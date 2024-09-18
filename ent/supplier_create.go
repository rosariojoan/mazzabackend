// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/company"
	"mazza/ent/payable"
	"mazza/ent/supplier"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SupplierCreate is the builder for creating a Supplier entity.
type SupplierCreate struct {
	config
	mutation *SupplierMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (sc *SupplierCreate) SetCreatedAt(t time.Time) *SupplierCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (sc *SupplierCreate) SetNillableCreatedAt(t *time.Time) *SupplierCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updatedAt" field.
func (sc *SupplierCreate) SetUpdatedAt(t time.Time) *SupplierCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (sc *SupplierCreate) SetNillableUpdatedAt(t *time.Time) *SupplierCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deletedAt" field.
func (sc *SupplierCreate) SetDeletedAt(t time.Time) *SupplierCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (sc *SupplierCreate) SetNillableDeletedAt(t *time.Time) *SupplierCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetAddress sets the "address" field.
func (sc *SupplierCreate) SetAddress(s string) *SupplierCreate {
	sc.mutation.SetAddress(s)
	return sc
}

// SetCity sets the "city" field.
func (sc *SupplierCreate) SetCity(s string) *SupplierCreate {
	sc.mutation.SetCity(s)
	return sc
}

// SetCountry sets the "country" field.
func (sc *SupplierCreate) SetCountry(s string) *SupplierCreate {
	sc.mutation.SetCountry(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SupplierCreate) SetDescription(s string) *SupplierCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetEmail sets the "email" field.
func (sc *SupplierCreate) SetEmail(s string) *SupplierCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetIsDefault sets the "isDefault" field.
func (sc *SupplierCreate) SetIsDefault(b bool) *SupplierCreate {
	sc.mutation.SetIsDefault(b)
	return sc
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (sc *SupplierCreate) SetNillableIsDefault(b *bool) *SupplierCreate {
	if b != nil {
		sc.SetIsDefault(*b)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SupplierCreate) SetName(s string) *SupplierCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetPhone sets the "phone" field.
func (sc *SupplierCreate) SetPhone(s string) *SupplierCreate {
	sc.mutation.SetPhone(s)
	return sc
}

// SetTaxId sets the "taxId" field.
func (sc *SupplierCreate) SetTaxId(s string) *SupplierCreate {
	sc.mutation.SetTaxId(s)
	return sc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (sc *SupplierCreate) SetCompanyID(id int) *SupplierCreate {
	sc.mutation.SetCompanyID(id)
	return sc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (sc *SupplierCreate) SetNillableCompanyID(id *int) *SupplierCreate {
	if id != nil {
		sc = sc.SetCompanyID(*id)
	}
	return sc
}

// SetCompany sets the "company" edge to the Company entity.
func (sc *SupplierCreate) SetCompany(c *Company) *SupplierCreate {
	return sc.SetCompanyID(c.ID)
}

// AddPayableIDs adds the "payables" edge to the Payable entity by IDs.
func (sc *SupplierCreate) AddPayableIDs(ids ...int) *SupplierCreate {
	sc.mutation.AddPayableIDs(ids...)
	return sc
}

// AddPayables adds the "payables" edges to the Payable entity.
func (sc *SupplierCreate) AddPayables(p ...*Payable) *SupplierCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPayableIDs(ids...)
}

// Mutation returns the SupplierMutation object of the builder.
func (sc *SupplierCreate) Mutation() *SupplierMutation {
	return sc.mutation
}

// Save creates the Supplier in the database.
func (sc *SupplierCreate) Save(ctx context.Context) (*Supplier, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SupplierCreate) SaveX(ctx context.Context) *Supplier {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SupplierCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SupplierCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SupplierCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := supplier.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := supplier.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.IsDefault(); !ok {
		v := supplier.DefaultIsDefault
		sc.mutation.SetIsDefault(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SupplierCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Supplier.createdAt"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Supplier.updatedAt"`)}
	}
	if _, ok := sc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Supplier.address"`)}
	}
	if _, ok := sc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "Supplier.city"`)}
	}
	if _, ok := sc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "Supplier.country"`)}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Supplier.description"`)}
	}
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Supplier.email"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Supplier.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := supplier.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Supplier.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Supplier.phone"`)}
	}
	if _, ok := sc.mutation.TaxId(); !ok {
		return &ValidationError{Name: "taxId", err: errors.New(`ent: missing required field "Supplier.taxId"`)}
	}
	if v, ok := sc.mutation.TaxId(); ok {
		if err := supplier.TaxIdValidator(v); err != nil {
			return &ValidationError{Name: "taxId", err: fmt.Errorf(`ent: validator failed for field "Supplier.taxId": %w`, err)}
		}
	}
	return nil
}

func (sc *SupplierCreate) sqlSave(ctx context.Context) (*Supplier, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SupplierCreate) createSpec() (*Supplier, *sqlgraph.CreateSpec) {
	var (
		_node = &Supplier{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(supplier.Table, sqlgraph.NewFieldSpec(supplier.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(supplier.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(supplier.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(supplier.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := sc.mutation.Address(); ok {
		_spec.SetField(supplier.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := sc.mutation.City(); ok {
		_spec.SetField(supplier.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := sc.mutation.Country(); ok {
		_spec.SetField(supplier.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(supplier.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.SetField(supplier.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := sc.mutation.IsDefault(); ok {
		_spec.SetField(supplier.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(supplier.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Phone(); ok {
		_spec.SetField(supplier.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := sc.mutation.TaxId(); ok {
		_spec.SetField(supplier.FieldTaxId, field.TypeString, value)
		_node.TaxId = value
	}
	if nodes := sc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   supplier.CompanyTable,
			Columns: []string{supplier.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_suppliers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PayablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   supplier.PayablesTable,
			Columns: []string{supplier.PayablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SupplierCreateBulk is the builder for creating many Supplier entities in bulk.
type SupplierCreateBulk struct {
	config
	err      error
	builders []*SupplierCreate
}

// Save creates the Supplier entities in the database.
func (scb *SupplierCreateBulk) Save(ctx context.Context) ([]*Supplier, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Supplier, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SupplierMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SupplierCreateBulk) SaveX(ctx context.Context) []*Supplier {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SupplierCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SupplierCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}