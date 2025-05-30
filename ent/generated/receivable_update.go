// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/receivable"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReceivableUpdate is the builder for updating Receivable entities.
type ReceivableUpdate struct {
	config
	hooks     []Hook
	mutation  *ReceivableMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ReceivableUpdate builder.
func (ru *ReceivableUpdate) Where(ps ...predicate.Receivable) *ReceivableUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updatedAt" field.
func (ru *ReceivableUpdate) SetUpdatedAt(t time.Time) *ReceivableUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeletedAt sets the "deletedAt" field.
func (ru *ReceivableUpdate) SetDeletedAt(t time.Time) *ReceivableUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableDeletedAt(t *time.Time) *ReceivableUpdate {
	if t != nil {
		ru.SetDeletedAt(*t)
	}
	return ru
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (ru *ReceivableUpdate) ClearDeletedAt() *ReceivableUpdate {
	ru.mutation.ClearDeletedAt()
	return ru
}

// SetEntryGroup sets the "entryGroup" field.
func (ru *ReceivableUpdate) SetEntryGroup(i int) *ReceivableUpdate {
	ru.mutation.ResetEntryGroup()
	ru.mutation.SetEntryGroup(i)
	return ru
}

// SetNillableEntryGroup sets the "entryGroup" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableEntryGroup(i *int) *ReceivableUpdate {
	if i != nil {
		ru.SetEntryGroup(*i)
	}
	return ru
}

// AddEntryGroup adds i to the "entryGroup" field.
func (ru *ReceivableUpdate) AddEntryGroup(i int) *ReceivableUpdate {
	ru.mutation.AddEntryGroup(i)
	return ru
}

// SetDate sets the "date" field.
func (ru *ReceivableUpdate) SetDate(t time.Time) *ReceivableUpdate {
	ru.mutation.SetDate(t)
	return ru
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableDate(t *time.Time) *ReceivableUpdate {
	if t != nil {
		ru.SetDate(*t)
	}
	return ru
}

// SetName sets the "name" field.
func (ru *ReceivableUpdate) SetName(s string) *ReceivableUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableName(s *string) *ReceivableUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetAmountInDefault sets the "amountInDefault" field.
func (ru *ReceivableUpdate) SetAmountInDefault(f float64) *ReceivableUpdate {
	ru.mutation.ResetAmountInDefault()
	ru.mutation.SetAmountInDefault(f)
	return ru
}

// SetNillableAmountInDefault sets the "amountInDefault" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableAmountInDefault(f *float64) *ReceivableUpdate {
	if f != nil {
		ru.SetAmountInDefault(*f)
	}
	return ru
}

// AddAmountInDefault adds f to the "amountInDefault" field.
func (ru *ReceivableUpdate) AddAmountInDefault(f float64) *ReceivableUpdate {
	ru.mutation.AddAmountInDefault(f)
	return ru
}

// SetOutstandingBalance sets the "outstandingBalance" field.
func (ru *ReceivableUpdate) SetOutstandingBalance(f float64) *ReceivableUpdate {
	ru.mutation.ResetOutstandingBalance()
	ru.mutation.SetOutstandingBalance(f)
	return ru
}

// SetNillableOutstandingBalance sets the "outstandingBalance" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableOutstandingBalance(f *float64) *ReceivableUpdate {
	if f != nil {
		ru.SetOutstandingBalance(*f)
	}
	return ru
}

// AddOutstandingBalance adds f to the "outstandingBalance" field.
func (ru *ReceivableUpdate) AddOutstandingBalance(f float64) *ReceivableUpdate {
	ru.mutation.AddOutstandingBalance(f)
	return ru
}

// SetTotalTransaction sets the "totalTransaction" field.
func (ru *ReceivableUpdate) SetTotalTransaction(f float64) *ReceivableUpdate {
	ru.mutation.ResetTotalTransaction()
	ru.mutation.SetTotalTransaction(f)
	return ru
}

// SetNillableTotalTransaction sets the "totalTransaction" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableTotalTransaction(f *float64) *ReceivableUpdate {
	if f != nil {
		ru.SetTotalTransaction(*f)
	}
	return ru
}

// AddTotalTransaction adds f to the "totalTransaction" field.
func (ru *ReceivableUpdate) AddTotalTransaction(f float64) *ReceivableUpdate {
	ru.mutation.AddTotalTransaction(f)
	return ru
}

// SetDueDate sets the "dueDate" field.
func (ru *ReceivableUpdate) SetDueDate(t time.Time) *ReceivableUpdate {
	ru.mutation.SetDueDate(t)
	return ru
}

// SetNillableDueDate sets the "dueDate" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableDueDate(t *time.Time) *ReceivableUpdate {
	if t != nil {
		ru.SetDueDate(*t)
	}
	return ru
}

// SetStatus sets the "status" field.
func (ru *ReceivableUpdate) SetStatus(r receivable.Status) *ReceivableUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableStatus(r *receivable.Status) *ReceivableUpdate {
	if r != nil {
		ru.SetStatus(*r)
	}
	return ru
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ru *ReceivableUpdate) SetCompanyID(id int) *ReceivableUpdate {
	ru.mutation.SetCompanyID(id)
	return ru
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableCompanyID(id *int) *ReceivableUpdate {
	if id != nil {
		ru = ru.SetCompanyID(*id)
	}
	return ru
}

// SetCompany sets the "company" edge to the Company entity.
func (ru *ReceivableUpdate) SetCompany(c *Company) *ReceivableUpdate {
	return ru.SetCompanyID(c.ID)
}

// SetInvoiceID sets the "invoice" edge to the Invoice entity by ID.
func (ru *ReceivableUpdate) SetInvoiceID(id int) *ReceivableUpdate {
	ru.mutation.SetInvoiceID(id)
	return ru
}

// SetNillableInvoiceID sets the "invoice" edge to the Invoice entity by ID if the given value is not nil.
func (ru *ReceivableUpdate) SetNillableInvoiceID(id *int) *ReceivableUpdate {
	if id != nil {
		ru = ru.SetInvoiceID(*id)
	}
	return ru
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (ru *ReceivableUpdate) SetInvoice(i *Invoice) *ReceivableUpdate {
	return ru.SetInvoiceID(i.ID)
}

// Mutation returns the ReceivableMutation object of the builder.
func (ru *ReceivableUpdate) Mutation() *ReceivableMutation {
	return ru.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ru *ReceivableUpdate) ClearCompany() *ReceivableUpdate {
	ru.mutation.ClearCompany()
	return ru
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (ru *ReceivableUpdate) ClearInvoice() *ReceivableUpdate {
	ru.mutation.ClearInvoice()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReceivableUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReceivableUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReceivableUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReceivableUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReceivableUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := receivable.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReceivableUpdate) check() error {
	if v, ok := ru.mutation.EntryGroup(); ok {
		if err := receivable.EntryGroupValidator(v); err != nil {
			return &ValidationError{Name: "entryGroup", err: fmt.Errorf(`generated: validator failed for field "Receivable.entryGroup": %w`, err)}
		}
	}
	if v, ok := ru.mutation.AmountInDefault(); ok {
		if err := receivable.AmountInDefaultValidator(v); err != nil {
			return &ValidationError{Name: "amountInDefault", err: fmt.Errorf(`generated: validator failed for field "Receivable.amountInDefault": %w`, err)}
		}
	}
	if v, ok := ru.mutation.OutstandingBalance(); ok {
		if err := receivable.OutstandingBalanceValidator(v); err != nil {
			return &ValidationError{Name: "outstandingBalance", err: fmt.Errorf(`generated: validator failed for field "Receivable.outstandingBalance": %w`, err)}
		}
	}
	if v, ok := ru.mutation.TotalTransaction(); ok {
		if err := receivable.TotalTransactionValidator(v); err != nil {
			return &ValidationError{Name: "totalTransaction", err: fmt.Errorf(`generated: validator failed for field "Receivable.totalTransaction": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Status(); ok {
		if err := receivable.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Receivable.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *ReceivableUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReceivableUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *ReceivableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(receivable.Table, receivable.Columns, sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(receivable.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(receivable.FieldDeletedAt, field.TypeTime, value)
	}
	if ru.mutation.DeletedAtCleared() {
		_spec.ClearField(receivable.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ru.mutation.EntryGroup(); ok {
		_spec.SetField(receivable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedEntryGroup(); ok {
		_spec.AddField(receivable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Date(); ok {
		_spec.SetField(receivable.FieldDate, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(receivable.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.AmountInDefault(); ok {
		_spec.SetField(receivable.FieldAmountInDefault, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedAmountInDefault(); ok {
		_spec.AddField(receivable.FieldAmountInDefault, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.OutstandingBalance(); ok {
		_spec.SetField(receivable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedOutstandingBalance(); ok {
		_spec.AddField(receivable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.TotalTransaction(); ok {
		_spec.SetField(receivable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedTotalTransaction(); ok {
		_spec.AddField(receivable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.DueDate(); ok {
		_spec.SetField(receivable.FieldDueDate, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(receivable.FieldStatus, field.TypeEnum, value)
	}
	if ru.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receivable.CompanyTable,
			Columns: []string{receivable.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receivable.CompanyTable,
			Columns: []string{receivable.CompanyColumn},
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
	if ru.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   receivable.InvoiceTable,
			Columns: []string{receivable.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   receivable.InvoiceTable,
			Columns: []string{receivable.InvoiceColumn},
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
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{receivable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReceivableUpdateOne is the builder for updating a single Receivable entity.
type ReceivableUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ReceivableMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updatedAt" field.
func (ruo *ReceivableUpdateOne) SetUpdatedAt(t time.Time) *ReceivableUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeletedAt sets the "deletedAt" field.
func (ruo *ReceivableUpdateOne) SetDeletedAt(t time.Time) *ReceivableUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableDeletedAt(t *time.Time) *ReceivableUpdateOne {
	if t != nil {
		ruo.SetDeletedAt(*t)
	}
	return ruo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (ruo *ReceivableUpdateOne) ClearDeletedAt() *ReceivableUpdateOne {
	ruo.mutation.ClearDeletedAt()
	return ruo
}

// SetEntryGroup sets the "entryGroup" field.
func (ruo *ReceivableUpdateOne) SetEntryGroup(i int) *ReceivableUpdateOne {
	ruo.mutation.ResetEntryGroup()
	ruo.mutation.SetEntryGroup(i)
	return ruo
}

// SetNillableEntryGroup sets the "entryGroup" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableEntryGroup(i *int) *ReceivableUpdateOne {
	if i != nil {
		ruo.SetEntryGroup(*i)
	}
	return ruo
}

// AddEntryGroup adds i to the "entryGroup" field.
func (ruo *ReceivableUpdateOne) AddEntryGroup(i int) *ReceivableUpdateOne {
	ruo.mutation.AddEntryGroup(i)
	return ruo
}

// SetDate sets the "date" field.
func (ruo *ReceivableUpdateOne) SetDate(t time.Time) *ReceivableUpdateOne {
	ruo.mutation.SetDate(t)
	return ruo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableDate(t *time.Time) *ReceivableUpdateOne {
	if t != nil {
		ruo.SetDate(*t)
	}
	return ruo
}

// SetName sets the "name" field.
func (ruo *ReceivableUpdateOne) SetName(s string) *ReceivableUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableName(s *string) *ReceivableUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetAmountInDefault sets the "amountInDefault" field.
func (ruo *ReceivableUpdateOne) SetAmountInDefault(f float64) *ReceivableUpdateOne {
	ruo.mutation.ResetAmountInDefault()
	ruo.mutation.SetAmountInDefault(f)
	return ruo
}

// SetNillableAmountInDefault sets the "amountInDefault" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableAmountInDefault(f *float64) *ReceivableUpdateOne {
	if f != nil {
		ruo.SetAmountInDefault(*f)
	}
	return ruo
}

// AddAmountInDefault adds f to the "amountInDefault" field.
func (ruo *ReceivableUpdateOne) AddAmountInDefault(f float64) *ReceivableUpdateOne {
	ruo.mutation.AddAmountInDefault(f)
	return ruo
}

// SetOutstandingBalance sets the "outstandingBalance" field.
func (ruo *ReceivableUpdateOne) SetOutstandingBalance(f float64) *ReceivableUpdateOne {
	ruo.mutation.ResetOutstandingBalance()
	ruo.mutation.SetOutstandingBalance(f)
	return ruo
}

// SetNillableOutstandingBalance sets the "outstandingBalance" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableOutstandingBalance(f *float64) *ReceivableUpdateOne {
	if f != nil {
		ruo.SetOutstandingBalance(*f)
	}
	return ruo
}

// AddOutstandingBalance adds f to the "outstandingBalance" field.
func (ruo *ReceivableUpdateOne) AddOutstandingBalance(f float64) *ReceivableUpdateOne {
	ruo.mutation.AddOutstandingBalance(f)
	return ruo
}

// SetTotalTransaction sets the "totalTransaction" field.
func (ruo *ReceivableUpdateOne) SetTotalTransaction(f float64) *ReceivableUpdateOne {
	ruo.mutation.ResetTotalTransaction()
	ruo.mutation.SetTotalTransaction(f)
	return ruo
}

// SetNillableTotalTransaction sets the "totalTransaction" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableTotalTransaction(f *float64) *ReceivableUpdateOne {
	if f != nil {
		ruo.SetTotalTransaction(*f)
	}
	return ruo
}

// AddTotalTransaction adds f to the "totalTransaction" field.
func (ruo *ReceivableUpdateOne) AddTotalTransaction(f float64) *ReceivableUpdateOne {
	ruo.mutation.AddTotalTransaction(f)
	return ruo
}

// SetDueDate sets the "dueDate" field.
func (ruo *ReceivableUpdateOne) SetDueDate(t time.Time) *ReceivableUpdateOne {
	ruo.mutation.SetDueDate(t)
	return ruo
}

// SetNillableDueDate sets the "dueDate" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableDueDate(t *time.Time) *ReceivableUpdateOne {
	if t != nil {
		ruo.SetDueDate(*t)
	}
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *ReceivableUpdateOne) SetStatus(r receivable.Status) *ReceivableUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableStatus(r *receivable.Status) *ReceivableUpdateOne {
	if r != nil {
		ruo.SetStatus(*r)
	}
	return ruo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ruo *ReceivableUpdateOne) SetCompanyID(id int) *ReceivableUpdateOne {
	ruo.mutation.SetCompanyID(id)
	return ruo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableCompanyID(id *int) *ReceivableUpdateOne {
	if id != nil {
		ruo = ruo.SetCompanyID(*id)
	}
	return ruo
}

// SetCompany sets the "company" edge to the Company entity.
func (ruo *ReceivableUpdateOne) SetCompany(c *Company) *ReceivableUpdateOne {
	return ruo.SetCompanyID(c.ID)
}

// SetInvoiceID sets the "invoice" edge to the Invoice entity by ID.
func (ruo *ReceivableUpdateOne) SetInvoiceID(id int) *ReceivableUpdateOne {
	ruo.mutation.SetInvoiceID(id)
	return ruo
}

// SetNillableInvoiceID sets the "invoice" edge to the Invoice entity by ID if the given value is not nil.
func (ruo *ReceivableUpdateOne) SetNillableInvoiceID(id *int) *ReceivableUpdateOne {
	if id != nil {
		ruo = ruo.SetInvoiceID(*id)
	}
	return ruo
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (ruo *ReceivableUpdateOne) SetInvoice(i *Invoice) *ReceivableUpdateOne {
	return ruo.SetInvoiceID(i.ID)
}

// Mutation returns the ReceivableMutation object of the builder.
func (ruo *ReceivableUpdateOne) Mutation() *ReceivableMutation {
	return ruo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ruo *ReceivableUpdateOne) ClearCompany() *ReceivableUpdateOne {
	ruo.mutation.ClearCompany()
	return ruo
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (ruo *ReceivableUpdateOne) ClearInvoice() *ReceivableUpdateOne {
	ruo.mutation.ClearInvoice()
	return ruo
}

// Where appends a list predicates to the ReceivableUpdate builder.
func (ruo *ReceivableUpdateOne) Where(ps ...predicate.Receivable) *ReceivableUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReceivableUpdateOne) Select(field string, fields ...string) *ReceivableUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Receivable entity.
func (ruo *ReceivableUpdateOne) Save(ctx context.Context) (*Receivable, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReceivableUpdateOne) SaveX(ctx context.Context) *Receivable {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReceivableUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReceivableUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReceivableUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := receivable.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReceivableUpdateOne) check() error {
	if v, ok := ruo.mutation.EntryGroup(); ok {
		if err := receivable.EntryGroupValidator(v); err != nil {
			return &ValidationError{Name: "entryGroup", err: fmt.Errorf(`generated: validator failed for field "Receivable.entryGroup": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.AmountInDefault(); ok {
		if err := receivable.AmountInDefaultValidator(v); err != nil {
			return &ValidationError{Name: "amountInDefault", err: fmt.Errorf(`generated: validator failed for field "Receivable.amountInDefault": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.OutstandingBalance(); ok {
		if err := receivable.OutstandingBalanceValidator(v); err != nil {
			return &ValidationError{Name: "outstandingBalance", err: fmt.Errorf(`generated: validator failed for field "Receivable.outstandingBalance": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.TotalTransaction(); ok {
		if err := receivable.TotalTransactionValidator(v); err != nil {
			return &ValidationError{Name: "totalTransaction", err: fmt.Errorf(`generated: validator failed for field "Receivable.totalTransaction": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Status(); ok {
		if err := receivable.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Receivable.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *ReceivableUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReceivableUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *ReceivableUpdateOne) sqlSave(ctx context.Context) (_node *Receivable, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(receivable.Table, receivable.Columns, sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Receivable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, receivable.FieldID)
		for _, f := range fields {
			if !receivable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != receivable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(receivable.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(receivable.FieldDeletedAt, field.TypeTime, value)
	}
	if ruo.mutation.DeletedAtCleared() {
		_spec.ClearField(receivable.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ruo.mutation.EntryGroup(); ok {
		_spec.SetField(receivable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedEntryGroup(); ok {
		_spec.AddField(receivable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Date(); ok {
		_spec.SetField(receivable.FieldDate, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(receivable.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.AmountInDefault(); ok {
		_spec.SetField(receivable.FieldAmountInDefault, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedAmountInDefault(); ok {
		_spec.AddField(receivable.FieldAmountInDefault, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.OutstandingBalance(); ok {
		_spec.SetField(receivable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedOutstandingBalance(); ok {
		_spec.AddField(receivable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.TotalTransaction(); ok {
		_spec.SetField(receivable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedTotalTransaction(); ok {
		_spec.AddField(receivable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.DueDate(); ok {
		_spec.SetField(receivable.FieldDueDate, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(receivable.FieldStatus, field.TypeEnum, value)
	}
	if ruo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receivable.CompanyTable,
			Columns: []string{receivable.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receivable.CompanyTable,
			Columns: []string{receivable.CompanyColumn},
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
	if ruo.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   receivable.InvoiceTable,
			Columns: []string{receivable.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   receivable.InvoiceTable,
			Columns: []string{receivable.InvoiceColumn},
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
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Receivable{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{receivable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
