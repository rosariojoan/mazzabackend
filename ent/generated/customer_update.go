// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/receivable"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks     []Hook
	mutation  *CustomerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updatedAt" field.
func (cu *CustomerUpdate) SetUpdatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deletedAt" field.
func (cu *CustomerUpdate) SetDeletedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableDeletedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (cu *CustomerUpdate) ClearDeletedAt() *CustomerUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetAddress sets the "address" field.
func (cu *CustomerUpdate) SetAddress(s string) *CustomerUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableAddress(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetAddress(*s)
	}
	return cu
}

// SetCity sets the "city" field.
func (cu *CustomerUpdate) SetCity(s string) *CustomerUpdate {
	cu.mutation.SetCity(s)
	return cu
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCity(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetCity(*s)
	}
	return cu
}

// SetCountry sets the "country" field.
func (cu *CustomerUpdate) SetCountry(s string) *CustomerUpdate {
	cu.mutation.SetCountry(s)
	return cu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCountry(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetCountry(*s)
	}
	return cu
}

// ClearCountry clears the value of the "country" field.
func (cu *CustomerUpdate) ClearCountry() *CustomerUpdate {
	cu.mutation.ClearCountry()
	return cu
}

// SetDescription sets the "description" field.
func (cu *CustomerUpdate) SetDescription(s string) *CustomerUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableDescription(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CustomerUpdate) ClearDescription() *CustomerUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetEmail sets the "email" field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableEmail(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetEmail(*s)
	}
	return cu
}

// ClearEmail clears the value of the "email" field.
func (cu *CustomerUpdate) ClearEmail() *CustomerUpdate {
	cu.mutation.ClearEmail()
	return cu
}

// SetIsDefault sets the "isDefault" field.
func (cu *CustomerUpdate) SetIsDefault(b bool) *CustomerUpdate {
	cu.mutation.SetIsDefault(b)
	return cu
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableIsDefault(b *bool) *CustomerUpdate {
	if b != nil {
		cu.SetIsDefault(*b)
	}
	return cu
}

// ClearIsDefault clears the value of the "isDefault" field.
func (cu *CustomerUpdate) ClearIsDefault() *CustomerUpdate {
	cu.mutation.ClearIsDefault()
	return cu
}

// SetName sets the "name" field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetPhone sets the "phone" field.
func (cu *CustomerUpdate) SetPhone(s string) *CustomerUpdate {
	cu.mutation.SetPhone(s)
	return cu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillablePhone(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetPhone(*s)
	}
	return cu
}

// SetTaxId sets the "taxId" field.
func (cu *CustomerUpdate) SetTaxId(s string) *CustomerUpdate {
	cu.mutation.SetTaxId(s)
	return cu
}

// SetNillableTaxId sets the "taxId" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableTaxId(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetTaxId(*s)
	}
	return cu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cu *CustomerUpdate) SetCompanyID(id int) *CustomerUpdate {
	cu.mutation.SetCompanyID(id)
	return cu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCompanyID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetCompanyID(*id)
	}
	return cu
}

// SetCompany sets the "company" edge to the Company entity.
func (cu *CustomerUpdate) SetCompany(c *Company) *CustomerUpdate {
	return cu.SetCompanyID(c.ID)
}

// AddReceivableIDs adds the "receivables" edge to the Receivable entity by IDs.
func (cu *CustomerUpdate) AddReceivableIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddReceivableIDs(ids...)
	return cu
}

// AddReceivables adds the "receivables" edges to the Receivable entity.
func (cu *CustomerUpdate) AddReceivables(r ...*Receivable) *CustomerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddReceivableIDs(ids...)
}

// AddInvoiceIDs adds the "invoices" edge to the Invoice entity by IDs.
func (cu *CustomerUpdate) AddInvoiceIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddInvoiceIDs(ids...)
	return cu
}

// AddInvoices adds the "invoices" edges to the Invoice entity.
func (cu *CustomerUpdate) AddInvoices(i ...*Invoice) *CustomerUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddInvoiceIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (cu *CustomerUpdate) ClearCompany() *CustomerUpdate {
	cu.mutation.ClearCompany()
	return cu
}

// ClearReceivables clears all "receivables" edges to the Receivable entity.
func (cu *CustomerUpdate) ClearReceivables() *CustomerUpdate {
	cu.mutation.ClearReceivables()
	return cu
}

// RemoveReceivableIDs removes the "receivables" edge to Receivable entities by IDs.
func (cu *CustomerUpdate) RemoveReceivableIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveReceivableIDs(ids...)
	return cu
}

// RemoveReceivables removes "receivables" edges to Receivable entities.
func (cu *CustomerUpdate) RemoveReceivables(r ...*Receivable) *CustomerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveReceivableIDs(ids...)
}

// ClearInvoices clears all "invoices" edges to the Invoice entity.
func (cu *CustomerUpdate) ClearInvoices() *CustomerUpdate {
	cu.mutation.ClearInvoices()
	return cu
}

// RemoveInvoiceIDs removes the "invoices" edge to Invoice entities by IDs.
func (cu *CustomerUpdate) RemoveInvoiceIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveInvoiceIDs(ids...)
	return cu
}

// RemoveInvoices removes "invoices" edges to Invoice entities.
func (cu *CustomerUpdate) RemoveInvoices(i ...*Invoice) *CustomerUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveInvoiceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CustomerUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := customer.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CustomerUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.TaxId(); ok {
		if err := customer.TaxIdValidator(v); err != nil {
			return &ValidationError{Name: "taxId", err: fmt.Errorf(`generated: validator failed for field "Customer.taxId": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CustomerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomerUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(customer.FieldDeletedAt, field.TypeTime, value)
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.ClearField(customer.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
	}
	if value, ok := cu.mutation.City(); ok {
		_spec.SetField(customer.FieldCity, field.TypeString, value)
	}
	if value, ok := cu.mutation.Country(); ok {
		_spec.SetField(customer.FieldCountry, field.TypeString, value)
	}
	if cu.mutation.CountryCleared() {
		_spec.ClearField(customer.FieldCountry, field.TypeString)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(customer.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(customer.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if cu.mutation.EmailCleared() {
		_spec.ClearField(customer.FieldEmail, field.TypeString)
	}
	if value, ok := cu.mutation.IsDefault(); ok {
		_spec.SetField(customer.FieldIsDefault, field.TypeBool, value)
	}
	if cu.mutation.IsDefaultCleared() {
		_spec.ClearField(customer.FieldIsDefault, field.TypeBool)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Phone(); ok {
		_spec.SetField(customer.FieldPhone, field.TypeString, value)
	}
	if value, ok := cu.mutation.TaxId(); ok {
		_spec.SetField(customer.FieldTaxId, field.TypeString, value)
	}
	if cu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
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
	if cu.mutation.ReceivablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceivablesTable,
			Columns: []string{customer.ReceivablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedReceivablesIDs(); len(nodes) > 0 && !cu.mutation.ReceivablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceivablesTable,
			Columns: []string{customer.ReceivablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ReceivablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceivablesTable,
			Columns: []string{customer.ReceivablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.InvoicesTable,
			Columns: []string{customer.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedInvoicesIDs(); len(nodes) > 0 && !cu.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.InvoicesTable,
			Columns: []string{customer.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.InvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.InvoicesTable,
			Columns: []string{customer.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CustomerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updatedAt" field.
func (cuo *CustomerUpdateOne) SetUpdatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deletedAt" field.
func (cuo *CustomerUpdateOne) SetDeletedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableDeletedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (cuo *CustomerUpdateOne) ClearDeletedAt() *CustomerUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CustomerUpdateOne) SetAddress(s string) *CustomerUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableAddress(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetAddress(*s)
	}
	return cuo
}

// SetCity sets the "city" field.
func (cuo *CustomerUpdateOne) SetCity(s string) *CustomerUpdateOne {
	cuo.mutation.SetCity(s)
	return cuo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCity(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetCity(*s)
	}
	return cuo
}

// SetCountry sets the "country" field.
func (cuo *CustomerUpdateOne) SetCountry(s string) *CustomerUpdateOne {
	cuo.mutation.SetCountry(s)
	return cuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCountry(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetCountry(*s)
	}
	return cuo
}

// ClearCountry clears the value of the "country" field.
func (cuo *CustomerUpdateOne) ClearCountry() *CustomerUpdateOne {
	cuo.mutation.ClearCountry()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CustomerUpdateOne) SetDescription(s string) *CustomerUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableDescription(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CustomerUpdateOne) ClearDescription() *CustomerUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetEmail sets the "email" field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableEmail(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetEmail(*s)
	}
	return cuo
}

// ClearEmail clears the value of the "email" field.
func (cuo *CustomerUpdateOne) ClearEmail() *CustomerUpdateOne {
	cuo.mutation.ClearEmail()
	return cuo
}

// SetIsDefault sets the "isDefault" field.
func (cuo *CustomerUpdateOne) SetIsDefault(b bool) *CustomerUpdateOne {
	cuo.mutation.SetIsDefault(b)
	return cuo
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableIsDefault(b *bool) *CustomerUpdateOne {
	if b != nil {
		cuo.SetIsDefault(*b)
	}
	return cuo
}

// ClearIsDefault clears the value of the "isDefault" field.
func (cuo *CustomerUpdateOne) ClearIsDefault() *CustomerUpdateOne {
	cuo.mutation.ClearIsDefault()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetPhone sets the "phone" field.
func (cuo *CustomerUpdateOne) SetPhone(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhone(s)
	return cuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillablePhone(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetPhone(*s)
	}
	return cuo
}

// SetTaxId sets the "taxId" field.
func (cuo *CustomerUpdateOne) SetTaxId(s string) *CustomerUpdateOne {
	cuo.mutation.SetTaxId(s)
	return cuo
}

// SetNillableTaxId sets the "taxId" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableTaxId(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetTaxId(*s)
	}
	return cuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cuo *CustomerUpdateOne) SetCompanyID(id int) *CustomerUpdateOne {
	cuo.mutation.SetCompanyID(id)
	return cuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCompanyID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetCompanyID(*id)
	}
	return cuo
}

// SetCompany sets the "company" edge to the Company entity.
func (cuo *CustomerUpdateOne) SetCompany(c *Company) *CustomerUpdateOne {
	return cuo.SetCompanyID(c.ID)
}

// AddReceivableIDs adds the "receivables" edge to the Receivable entity by IDs.
func (cuo *CustomerUpdateOne) AddReceivableIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddReceivableIDs(ids...)
	return cuo
}

// AddReceivables adds the "receivables" edges to the Receivable entity.
func (cuo *CustomerUpdateOne) AddReceivables(r ...*Receivable) *CustomerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddReceivableIDs(ids...)
}

// AddInvoiceIDs adds the "invoices" edge to the Invoice entity by IDs.
func (cuo *CustomerUpdateOne) AddInvoiceIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddInvoiceIDs(ids...)
	return cuo
}

// AddInvoices adds the "invoices" edges to the Invoice entity.
func (cuo *CustomerUpdateOne) AddInvoices(i ...*Invoice) *CustomerUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddInvoiceIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (cuo *CustomerUpdateOne) ClearCompany() *CustomerUpdateOne {
	cuo.mutation.ClearCompany()
	return cuo
}

// ClearReceivables clears all "receivables" edges to the Receivable entity.
func (cuo *CustomerUpdateOne) ClearReceivables() *CustomerUpdateOne {
	cuo.mutation.ClearReceivables()
	return cuo
}

// RemoveReceivableIDs removes the "receivables" edge to Receivable entities by IDs.
func (cuo *CustomerUpdateOne) RemoveReceivableIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveReceivableIDs(ids...)
	return cuo
}

// RemoveReceivables removes "receivables" edges to Receivable entities.
func (cuo *CustomerUpdateOne) RemoveReceivables(r ...*Receivable) *CustomerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveReceivableIDs(ids...)
}

// ClearInvoices clears all "invoices" edges to the Invoice entity.
func (cuo *CustomerUpdateOne) ClearInvoices() *CustomerUpdateOne {
	cuo.mutation.ClearInvoices()
	return cuo
}

// RemoveInvoiceIDs removes the "invoices" edge to Invoice entities by IDs.
func (cuo *CustomerUpdateOne) RemoveInvoiceIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveInvoiceIDs(ids...)
	return cuo
}

// RemoveInvoices removes "invoices" edges to Invoice entities.
func (cuo *CustomerUpdateOne) RemoveInvoices(i ...*Invoice) *CustomerUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveInvoiceIDs(ids...)
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cuo *CustomerUpdateOne) Where(ps ...predicate.Customer) *CustomerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CustomerUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := customer.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CustomerUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.TaxId(); ok {
		if err := customer.TaxIdValidator(v); err != nil {
			return &ValidationError{Name: "taxId", err: fmt.Errorf(`generated: validator failed for field "Customer.taxId": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CustomerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomerUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(customer.FieldDeletedAt, field.TypeTime, value)
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.ClearField(customer.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
	}
	if value, ok := cuo.mutation.City(); ok {
		_spec.SetField(customer.FieldCity, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Country(); ok {
		_spec.SetField(customer.FieldCountry, field.TypeString, value)
	}
	if cuo.mutation.CountryCleared() {
		_spec.ClearField(customer.FieldCountry, field.TypeString)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(customer.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(customer.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if cuo.mutation.EmailCleared() {
		_spec.ClearField(customer.FieldEmail, field.TypeString)
	}
	if value, ok := cuo.mutation.IsDefault(); ok {
		_spec.SetField(customer.FieldIsDefault, field.TypeBool, value)
	}
	if cuo.mutation.IsDefaultCleared() {
		_spec.ClearField(customer.FieldIsDefault, field.TypeBool)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Phone(); ok {
		_spec.SetField(customer.FieldPhone, field.TypeString, value)
	}
	if value, ok := cuo.mutation.TaxId(); ok {
		_spec.SetField(customer.FieldTaxId, field.TypeString, value)
	}
	if cuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.CompanyTable,
			Columns: []string{customer.CompanyColumn},
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
	if cuo.mutation.ReceivablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceivablesTable,
			Columns: []string{customer.ReceivablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedReceivablesIDs(); len(nodes) > 0 && !cuo.mutation.ReceivablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceivablesTable,
			Columns: []string{customer.ReceivablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ReceivablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceivablesTable,
			Columns: []string{customer.ReceivablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.InvoicesTable,
			Columns: []string{customer.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedInvoicesIDs(); len(nodes) > 0 && !cuo.mutation.InvoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.InvoicesTable,
			Columns: []string{customer.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.InvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.InvoicesTable,
			Columns: []string{customer.InvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
