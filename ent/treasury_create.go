// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/cashmovement"
	"mazza/ent/company"
	"mazza/ent/treasury"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TreasuryCreate is the builder for creating a Treasury entity.
type TreasuryCreate struct {
	config
	mutation *TreasuryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (tc *TreasuryCreate) SetCreatedAt(t time.Time) *TreasuryCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableCreatedAt(t *time.Time) *TreasuryCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updatedAt" field.
func (tc *TreasuryCreate) SetUpdatedAt(t time.Time) *TreasuryCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableUpdatedAt(t *time.Time) *TreasuryCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeletedAt sets the "deletedAt" field.
func (tc *TreasuryCreate) SetDeletedAt(t time.Time) *TreasuryCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableDeletedAt(t *time.Time) *TreasuryCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetAccountNumber sets the "accountNumber" field.
func (tc *TreasuryCreate) SetAccountNumber(s string) *TreasuryCreate {
	tc.mutation.SetAccountNumber(s)
	return tc
}

// SetNillableAccountNumber sets the "accountNumber" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableAccountNumber(s *string) *TreasuryCreate {
	if s != nil {
		tc.SetAccountNumber(*s)
	}
	return tc
}

// SetBalance sets the "balance" field.
func (tc *TreasuryCreate) SetBalance(f float64) *TreasuryCreate {
	tc.mutation.SetBalance(f)
	return tc
}

// SetBankName sets the "bankName" field.
func (tc *TreasuryCreate) SetBankName(s string) *TreasuryCreate {
	tc.mutation.SetBankName(s)
	return tc
}

// SetNillableBankName sets the "bankName" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableBankName(s *string) *TreasuryCreate {
	if s != nil {
		tc.SetBankName(*s)
	}
	return tc
}

// SetCurrency sets the "currency" field.
func (tc *TreasuryCreate) SetCurrency(t treasury.Currency) *TreasuryCreate {
	tc.mutation.SetCurrency(t)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TreasuryCreate) SetDescription(s string) *TreasuryCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableDescription(s *string) *TreasuryCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetIban sets the "iban" field.
func (tc *TreasuryCreate) SetIban(s string) *TreasuryCreate {
	tc.mutation.SetIban(s)
	return tc
}

// SetNillableIban sets the "iban" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableIban(s *string) *TreasuryCreate {
	if s != nil {
		tc.SetIban(*s)
	}
	return tc
}

// SetIsDefault sets the "isDefault" field.
func (tc *TreasuryCreate) SetIsDefault(b bool) *TreasuryCreate {
	tc.mutation.SetIsDefault(b)
	return tc
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableIsDefault(b *bool) *TreasuryCreate {
	if b != nil {
		tc.SetIsDefault(*b)
	}
	return tc
}

// SetIsMainAccount sets the "isMainAccount" field.
func (tc *TreasuryCreate) SetIsMainAccount(b bool) *TreasuryCreate {
	tc.mutation.SetIsMainAccount(b)
	return tc
}

// SetNillableIsMainAccount sets the "isMainAccount" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableIsMainAccount(b *bool) *TreasuryCreate {
	if b != nil {
		tc.SetIsMainAccount(*b)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TreasuryCreate) SetName(s string) *TreasuryCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetCategory sets the "category" field.
func (tc *TreasuryCreate) SetCategory(t treasury.Category) *TreasuryCreate {
	tc.mutation.SetCategory(t)
	return tc
}

// SetSwiftCode sets the "swiftCode" field.
func (tc *TreasuryCreate) SetSwiftCode(s string) *TreasuryCreate {
	tc.mutation.SetSwiftCode(s)
	return tc
}

// SetNillableSwiftCode sets the "swiftCode" field if the given value is not nil.
func (tc *TreasuryCreate) SetNillableSwiftCode(s *string) *TreasuryCreate {
	if s != nil {
		tc.SetSwiftCode(*s)
	}
	return tc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (tc *TreasuryCreate) SetCompanyID(id int) *TreasuryCreate {
	tc.mutation.SetCompanyID(id)
	return tc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (tc *TreasuryCreate) SetNillableCompanyID(id *int) *TreasuryCreate {
	if id != nil {
		tc = tc.SetCompanyID(*id)
	}
	return tc
}

// SetCompany sets the "company" edge to the Company entity.
func (tc *TreasuryCreate) SetCompany(c *Company) *TreasuryCreate {
	return tc.SetCompanyID(c.ID)
}

// AddCashMovementIDs adds the "cashMovements" edge to the CashMovement entity by IDs.
func (tc *TreasuryCreate) AddCashMovementIDs(ids ...int) *TreasuryCreate {
	tc.mutation.AddCashMovementIDs(ids...)
	return tc
}

// AddCashMovements adds the "cashMovements" edges to the CashMovement entity.
func (tc *TreasuryCreate) AddCashMovements(c ...*CashMovement) *TreasuryCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCashMovementIDs(ids...)
}

// Mutation returns the TreasuryMutation object of the builder.
func (tc *TreasuryCreate) Mutation() *TreasuryMutation {
	return tc.mutation
}

// Save creates the Treasury in the database.
func (tc *TreasuryCreate) Save(ctx context.Context) (*Treasury, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TreasuryCreate) SaveX(ctx context.Context) *Treasury {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TreasuryCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TreasuryCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TreasuryCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := treasury.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := treasury.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.IsDefault(); !ok {
		v := treasury.DefaultIsDefault
		tc.mutation.SetIsDefault(v)
	}
	if _, ok := tc.mutation.IsMainAccount(); !ok {
		v := treasury.DefaultIsMainAccount
		tc.mutation.SetIsMainAccount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TreasuryCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Treasury.createdAt"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Treasury.updatedAt"`)}
	}
	if _, ok := tc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "Treasury.balance"`)}
	}
	if _, ok := tc.mutation.Currency(); !ok {
		return &ValidationError{Name: "currency", err: errors.New(`ent: missing required field "Treasury.currency"`)}
	}
	if v, ok := tc.mutation.Currency(); ok {
		if err := treasury.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "Treasury.currency": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Treasury.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := treasury.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Treasury.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Treasury.category"`)}
	}
	if v, ok := tc.mutation.Category(); ok {
		if err := treasury.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Treasury.category": %w`, err)}
		}
	}
	return nil
}

func (tc *TreasuryCreate) sqlSave(ctx context.Context) (*Treasury, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TreasuryCreate) createSpec() (*Treasury, *sqlgraph.CreateSpec) {
	var (
		_node = &Treasury{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(treasury.Table, sqlgraph.NewFieldSpec(treasury.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(treasury.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(treasury.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(treasury.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := tc.mutation.AccountNumber(); ok {
		_spec.SetField(treasury.FieldAccountNumber, field.TypeString, value)
		_node.AccountNumber = value
	}
	if value, ok := tc.mutation.Balance(); ok {
		_spec.SetField(treasury.FieldBalance, field.TypeFloat64, value)
		_node.Balance = value
	}
	if value, ok := tc.mutation.BankName(); ok {
		_spec.SetField(treasury.FieldBankName, field.TypeString, value)
		_node.BankName = value
	}
	if value, ok := tc.mutation.Currency(); ok {
		_spec.SetField(treasury.FieldCurrency, field.TypeEnum, value)
		_node.Currency = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(treasury.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Iban(); ok {
		_spec.SetField(treasury.FieldIban, field.TypeString, value)
		_node.Iban = value
	}
	if value, ok := tc.mutation.IsDefault(); ok {
		_spec.SetField(treasury.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := tc.mutation.IsMainAccount(); ok {
		_spec.SetField(treasury.FieldIsMainAccount, field.TypeBool, value)
		_node.IsMainAccount = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(treasury.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Category(); ok {
		_spec.SetField(treasury.FieldCategory, field.TypeEnum, value)
		_node.Category = value
	}
	if value, ok := tc.mutation.SwiftCode(); ok {
		_spec.SetField(treasury.FieldSwiftCode, field.TypeString, value)
		_node.SwiftCode = value
	}
	if nodes := tc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treasury.CompanyTable,
			Columns: []string{treasury.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_treasuries = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CashMovementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   treasury.CashMovementsTable,
			Columns: []string{treasury.CashMovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cashmovement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TreasuryCreateBulk is the builder for creating many Treasury entities in bulk.
type TreasuryCreateBulk struct {
	config
	err      error
	builders []*TreasuryCreate
}

// Save creates the Treasury entities in the database.
func (tcb *TreasuryCreateBulk) Save(ctx context.Context) ([]*Treasury, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Treasury, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TreasuryMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TreasuryCreateBulk) SaveX(ctx context.Context) []*Treasury {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TreasuryCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TreasuryCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
