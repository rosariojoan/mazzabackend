// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/accountingentry"
	"mazza/ent/company"
	"mazza/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountingEntryCreate is the builder for creating a AccountingEntry entity.
type AccountingEntryCreate struct {
	config
	mutation *AccountingEntryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (aec *AccountingEntryCreate) SetCreatedAt(t time.Time) *AccountingEntryCreate {
	aec.mutation.SetCreatedAt(t)
	return aec
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableCreatedAt(t *time.Time) *AccountingEntryCreate {
	if t != nil {
		aec.SetCreatedAt(*t)
	}
	return aec
}

// SetUpdatedAt sets the "updatedAt" field.
func (aec *AccountingEntryCreate) SetUpdatedAt(t time.Time) *AccountingEntryCreate {
	aec.mutation.SetUpdatedAt(t)
	return aec
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableUpdatedAt(t *time.Time) *AccountingEntryCreate {
	if t != nil {
		aec.SetUpdatedAt(*t)
	}
	return aec
}

// SetDeletedAt sets the "deletedAt" field.
func (aec *AccountingEntryCreate) SetDeletedAt(t time.Time) *AccountingEntryCreate {
	aec.mutation.SetDeletedAt(t)
	return aec
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableDeletedAt(t *time.Time) *AccountingEntryCreate {
	if t != nil {
		aec.SetDeletedAt(*t)
	}
	return aec
}

// SetNumber sets the "number" field.
func (aec *AccountingEntryCreate) SetNumber(i int) *AccountingEntryCreate {
	aec.mutation.SetNumber(i)
	return aec
}

// SetGroup sets the "group" field.
func (aec *AccountingEntryCreate) SetGroup(i int) *AccountingEntryCreate {
	aec.mutation.SetGroup(i)
	return aec
}

// SetDate sets the "date" field.
func (aec *AccountingEntryCreate) SetDate(t time.Time) *AccountingEntryCreate {
	aec.mutation.SetDate(t)
	return aec
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableDate(t *time.Time) *AccountingEntryCreate {
	if t != nil {
		aec.SetDate(*t)
	}
	return aec
}

// SetAccount sets the "account" field.
func (aec *AccountingEntryCreate) SetAccount(s string) *AccountingEntryCreate {
	aec.mutation.SetAccount(s)
	return aec
}

// SetLabel sets the "label" field.
func (aec *AccountingEntryCreate) SetLabel(s string) *AccountingEntryCreate {
	aec.mutation.SetLabel(s)
	return aec
}

// SetAmount sets the "amount" field.
func (aec *AccountingEntryCreate) SetAmount(f float64) *AccountingEntryCreate {
	aec.mutation.SetAmount(f)
	return aec
}

// SetDescription sets the "description" field.
func (aec *AccountingEntryCreate) SetDescription(s string) *AccountingEntryCreate {
	aec.mutation.SetDescription(s)
	return aec
}

// SetAccountType sets the "accountType" field.
func (aec *AccountingEntryCreate) SetAccountType(at accountingentry.AccountType) *AccountingEntryCreate {
	aec.mutation.SetAccountType(at)
	return aec
}

// SetIsDebit sets the "isDebit" field.
func (aec *AccountingEntryCreate) SetIsDebit(b bool) *AccountingEntryCreate {
	aec.mutation.SetIsDebit(b)
	return aec
}

// SetIsReversal sets the "isReversal" field.
func (aec *AccountingEntryCreate) SetIsReversal(b bool) *AccountingEntryCreate {
	aec.mutation.SetIsReversal(b)
	return aec
}

// SetNillableIsReversal sets the "isReversal" field if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableIsReversal(b *bool) *AccountingEntryCreate {
	if b != nil {
		aec.SetIsReversal(*b)
	}
	return aec
}

// SetReversed sets the "reversed" field.
func (aec *AccountingEntryCreate) SetReversed(b bool) *AccountingEntryCreate {
	aec.mutation.SetReversed(b)
	return aec
}

// SetNillableReversed sets the "reversed" field if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableReversed(b *bool) *AccountingEntryCreate {
	if b != nil {
		aec.SetReversed(*b)
	}
	return aec
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (aec *AccountingEntryCreate) SetCompanyID(id int) *AccountingEntryCreate {
	aec.mutation.SetCompanyID(id)
	return aec
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableCompanyID(id *int) *AccountingEntryCreate {
	if id != nil {
		aec = aec.SetCompanyID(*id)
	}
	return aec
}

// SetCompany sets the "company" edge to the Company entity.
func (aec *AccountingEntryCreate) SetCompany(c *Company) *AccountingEntryCreate {
	return aec.SetCompanyID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (aec *AccountingEntryCreate) SetUserID(id int) *AccountingEntryCreate {
	aec.mutation.SetUserID(id)
	return aec
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (aec *AccountingEntryCreate) SetNillableUserID(id *int) *AccountingEntryCreate {
	if id != nil {
		aec = aec.SetUserID(*id)
	}
	return aec
}

// SetUser sets the "user" edge to the User entity.
func (aec *AccountingEntryCreate) SetUser(u *User) *AccountingEntryCreate {
	return aec.SetUserID(u.ID)
}

// Mutation returns the AccountingEntryMutation object of the builder.
func (aec *AccountingEntryCreate) Mutation() *AccountingEntryMutation {
	return aec.mutation
}

// Save creates the AccountingEntry in the database.
func (aec *AccountingEntryCreate) Save(ctx context.Context) (*AccountingEntry, error) {
	aec.defaults()
	return withHooks(ctx, aec.sqlSave, aec.mutation, aec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aec *AccountingEntryCreate) SaveX(ctx context.Context) *AccountingEntry {
	v, err := aec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aec *AccountingEntryCreate) Exec(ctx context.Context) error {
	_, err := aec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aec *AccountingEntryCreate) ExecX(ctx context.Context) {
	if err := aec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aec *AccountingEntryCreate) defaults() {
	if _, ok := aec.mutation.CreatedAt(); !ok {
		v := accountingentry.DefaultCreatedAt()
		aec.mutation.SetCreatedAt(v)
	}
	if _, ok := aec.mutation.UpdatedAt(); !ok {
		v := accountingentry.DefaultUpdatedAt()
		aec.mutation.SetUpdatedAt(v)
	}
	if _, ok := aec.mutation.Date(); !ok {
		v := accountingentry.DefaultDate()
		aec.mutation.SetDate(v)
	}
	if _, ok := aec.mutation.IsReversal(); !ok {
		v := accountingentry.DefaultIsReversal
		aec.mutation.SetIsReversal(v)
	}
	if _, ok := aec.mutation.Reversed(); !ok {
		v := accountingentry.DefaultReversed
		aec.mutation.SetReversed(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aec *AccountingEntryCreate) check() error {
	if _, ok := aec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "AccountingEntry.createdAt"`)}
	}
	if _, ok := aec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "AccountingEntry.updatedAt"`)}
	}
	if _, ok := aec.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "AccountingEntry.number"`)}
	}
	if v, ok := aec.mutation.Number(); ok {
		if err := accountingentry.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "AccountingEntry.number": %w`, err)}
		}
	}
	if _, ok := aec.mutation.Group(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required field "AccountingEntry.group"`)}
	}
	if v, ok := aec.mutation.Group(); ok {
		if err := accountingentry.GroupValidator(v); err != nil {
			return &ValidationError{Name: "group", err: fmt.Errorf(`ent: validator failed for field "AccountingEntry.group": %w`, err)}
		}
	}
	if _, ok := aec.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "AccountingEntry.date"`)}
	}
	if _, ok := aec.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required field "AccountingEntry.account"`)}
	}
	if v, ok := aec.mutation.Account(); ok {
		if err := accountingentry.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf(`ent: validator failed for field "AccountingEntry.account": %w`, err)}
		}
	}
	if _, ok := aec.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "AccountingEntry.label"`)}
	}
	if _, ok := aec.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "AccountingEntry.amount"`)}
	}
	if _, ok := aec.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "AccountingEntry.description"`)}
	}
	if _, ok := aec.mutation.AccountType(); !ok {
		return &ValidationError{Name: "accountType", err: errors.New(`ent: missing required field "AccountingEntry.accountType"`)}
	}
	if v, ok := aec.mutation.AccountType(); ok {
		if err := accountingentry.AccountTypeValidator(v); err != nil {
			return &ValidationError{Name: "accountType", err: fmt.Errorf(`ent: validator failed for field "AccountingEntry.accountType": %w`, err)}
		}
	}
	if _, ok := aec.mutation.IsDebit(); !ok {
		return &ValidationError{Name: "isDebit", err: errors.New(`ent: missing required field "AccountingEntry.isDebit"`)}
	}
	if _, ok := aec.mutation.IsReversal(); !ok {
		return &ValidationError{Name: "isReversal", err: errors.New(`ent: missing required field "AccountingEntry.isReversal"`)}
	}
	if _, ok := aec.mutation.Reversed(); !ok {
		return &ValidationError{Name: "reversed", err: errors.New(`ent: missing required field "AccountingEntry.reversed"`)}
	}
	return nil
}

func (aec *AccountingEntryCreate) sqlSave(ctx context.Context) (*AccountingEntry, error) {
	if err := aec.check(); err != nil {
		return nil, err
	}
	_node, _spec := aec.createSpec()
	if err := sqlgraph.CreateNode(ctx, aec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	aec.mutation.id = &_node.ID
	aec.mutation.done = true
	return _node, nil
}

func (aec *AccountingEntryCreate) createSpec() (*AccountingEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountingEntry{config: aec.config}
		_spec = sqlgraph.NewCreateSpec(accountingentry.Table, sqlgraph.NewFieldSpec(accountingentry.FieldID, field.TypeInt))
	)
	if value, ok := aec.mutation.CreatedAt(); ok {
		_spec.SetField(accountingentry.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aec.mutation.UpdatedAt(); ok {
		_spec.SetField(accountingentry.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := aec.mutation.DeletedAt(); ok {
		_spec.SetField(accountingentry.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := aec.mutation.Number(); ok {
		_spec.SetField(accountingentry.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := aec.mutation.Group(); ok {
		_spec.SetField(accountingentry.FieldGroup, field.TypeInt, value)
		_node.Group = value
	}
	if value, ok := aec.mutation.Date(); ok {
		_spec.SetField(accountingentry.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := aec.mutation.Account(); ok {
		_spec.SetField(accountingentry.FieldAccount, field.TypeString, value)
		_node.Account = value
	}
	if value, ok := aec.mutation.Label(); ok {
		_spec.SetField(accountingentry.FieldLabel, field.TypeString, value)
		_node.Label = value
	}
	if value, ok := aec.mutation.Amount(); ok {
		_spec.SetField(accountingentry.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := aec.mutation.Description(); ok {
		_spec.SetField(accountingentry.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := aec.mutation.AccountType(); ok {
		_spec.SetField(accountingentry.FieldAccountType, field.TypeEnum, value)
		_node.AccountType = value
	}
	if value, ok := aec.mutation.IsDebit(); ok {
		_spec.SetField(accountingentry.FieldIsDebit, field.TypeBool, value)
		_node.IsDebit = value
	}
	if value, ok := aec.mutation.IsReversal(); ok {
		_spec.SetField(accountingentry.FieldIsReversal, field.TypeBool, value)
		_node.IsReversal = value
	}
	if value, ok := aec.mutation.Reversed(); ok {
		_spec.SetField(accountingentry.FieldReversed, field.TypeBool, value)
		_node.Reversed = value
	}
	if nodes := aec.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountingentry.CompanyTable,
			Columns: []string{accountingentry.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_accounting_entries = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := aec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountingentry.UserTable,
			Columns: []string{accountingentry.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_accounting_entries = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountingEntryCreateBulk is the builder for creating many AccountingEntry entities in bulk.
type AccountingEntryCreateBulk struct {
	config
	err      error
	builders []*AccountingEntryCreate
}

// Save creates the AccountingEntry entities in the database.
func (aecb *AccountingEntryCreateBulk) Save(ctx context.Context) ([]*AccountingEntry, error) {
	if aecb.err != nil {
		return nil, aecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aecb.builders))
	nodes := make([]*AccountingEntry, len(aecb.builders))
	mutators := make([]Mutator, len(aecb.builders))
	for i := range aecb.builders {
		func(i int, root context.Context) {
			builder := aecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountingEntryMutation)
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
					_, err = mutators[i+1].Mutate(root, aecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aecb *AccountingEntryCreateBulk) SaveX(ctx context.Context) []*AccountingEntry {
	v, err := aecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aecb *AccountingEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := aecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aecb *AccountingEntryCreateBulk) ExecX(ctx context.Context) {
	if err := aecb.Exec(ctx); err != nil {
		panic(err)
	}
}