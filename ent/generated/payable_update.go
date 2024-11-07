// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PayableUpdate is the builder for updating Payable entities.
type PayableUpdate struct {
	config
	hooks     []Hook
	mutation  *PayableMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PayableUpdate builder.
func (pu *PayableUpdate) Where(ps ...predicate.Payable) *PayableUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updatedAt" field.
func (pu *PayableUpdate) SetUpdatedAt(t time.Time) *PayableUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDeletedAt sets the "deletedAt" field.
func (pu *PayableUpdate) SetDeletedAt(t time.Time) *PayableUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableDeletedAt(t *time.Time) *PayableUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (pu *PayableUpdate) ClearDeletedAt() *PayableUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetEntryGroup sets the "entryGroup" field.
func (pu *PayableUpdate) SetEntryGroup(i int) *PayableUpdate {
	pu.mutation.ResetEntryGroup()
	pu.mutation.SetEntryGroup(i)
	return pu
}

// SetNillableEntryGroup sets the "entryGroup" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableEntryGroup(i *int) *PayableUpdate {
	if i != nil {
		pu.SetEntryGroup(*i)
	}
	return pu
}

// AddEntryGroup adds i to the "entryGroup" field.
func (pu *PayableUpdate) AddEntryGroup(i int) *PayableUpdate {
	pu.mutation.AddEntryGroup(i)
	return pu
}

// SetDate sets the "date" field.
func (pu *PayableUpdate) SetDate(t time.Time) *PayableUpdate {
	pu.mutation.SetDate(t)
	return pu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableDate(t *time.Time) *PayableUpdate {
	if t != nil {
		pu.SetDate(*t)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PayableUpdate) SetName(s string) *PayableUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableName(s *string) *PayableUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetOutstandingBalance sets the "outstandingBalance" field.
func (pu *PayableUpdate) SetOutstandingBalance(f float64) *PayableUpdate {
	pu.mutation.ResetOutstandingBalance()
	pu.mutation.SetOutstandingBalance(f)
	return pu
}

// SetNillableOutstandingBalance sets the "outstandingBalance" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableOutstandingBalance(f *float64) *PayableUpdate {
	if f != nil {
		pu.SetOutstandingBalance(*f)
	}
	return pu
}

// AddOutstandingBalance adds f to the "outstandingBalance" field.
func (pu *PayableUpdate) AddOutstandingBalance(f float64) *PayableUpdate {
	pu.mutation.AddOutstandingBalance(f)
	return pu
}

// SetTotalTransaction sets the "totalTransaction" field.
func (pu *PayableUpdate) SetTotalTransaction(f float64) *PayableUpdate {
	pu.mutation.ResetTotalTransaction()
	pu.mutation.SetTotalTransaction(f)
	return pu
}

// SetNillableTotalTransaction sets the "totalTransaction" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableTotalTransaction(f *float64) *PayableUpdate {
	if f != nil {
		pu.SetTotalTransaction(*f)
	}
	return pu
}

// AddTotalTransaction adds f to the "totalTransaction" field.
func (pu *PayableUpdate) AddTotalTransaction(f float64) *PayableUpdate {
	pu.mutation.AddTotalTransaction(f)
	return pu
}

// SetDueDate sets the "dueDate" field.
func (pu *PayableUpdate) SetDueDate(t time.Time) *PayableUpdate {
	pu.mutation.SetDueDate(t)
	return pu
}

// SetNillableDueDate sets the "dueDate" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableDueDate(t *time.Time) *PayableUpdate {
	if t != nil {
		pu.SetDueDate(*t)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PayableUpdate) SetStatus(pa payable.Status) *PayableUpdate {
	pu.mutation.SetStatus(pa)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PayableUpdate) SetNillableStatus(pa *payable.Status) *PayableUpdate {
	if pa != nil {
		pu.SetStatus(*pa)
	}
	return pu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pu *PayableUpdate) SetCompanyID(id int) *PayableUpdate {
	pu.mutation.SetCompanyID(id)
	return pu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pu *PayableUpdate) SetNillableCompanyID(id *int) *PayableUpdate {
	if id != nil {
		pu = pu.SetCompanyID(*id)
	}
	return pu
}

// SetCompany sets the "company" edge to the Company entity.
func (pu *PayableUpdate) SetCompany(c *Company) *PayableUpdate {
	return pu.SetCompanyID(c.ID)
}

// Mutation returns the PayableMutation object of the builder.
func (pu *PayableUpdate) Mutation() *PayableMutation {
	return pu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (pu *PayableUpdate) ClearCompany() *PayableUpdate {
	pu.mutation.ClearCompany()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PayableUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PayableUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PayableUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PayableUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PayableUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := payable.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PayableUpdate) check() error {
	if v, ok := pu.mutation.EntryGroup(); ok {
		if err := payable.EntryGroupValidator(v); err != nil {
			return &ValidationError{Name: "entryGroup", err: fmt.Errorf(`generated: validator failed for field "Payable.entryGroup": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := payable.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Payable.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PayableUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PayableUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PayableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payable.Table, payable.Columns, sqlgraph.NewFieldSpec(payable.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(payable.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(payable.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(payable.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.EntryGroup(); ok {
		_spec.SetField(payable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedEntryGroup(); ok {
		_spec.AddField(payable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Date(); ok {
		_spec.SetField(payable.FieldDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(payable.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.OutstandingBalance(); ok {
		_spec.SetField(payable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedOutstandingBalance(); ok {
		_spec.AddField(payable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.TotalTransaction(); ok {
		_spec.SetField(payable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedTotalTransaction(); ok {
		_spec.AddField(payable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.DueDate(); ok {
		_spec.SetField(payable.FieldDueDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(payable.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payable.CompanyTable,
			Columns: []string{payable.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payable.CompanyTable,
			Columns: []string{payable.CompanyColumn},
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
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PayableUpdateOne is the builder for updating a single Payable entity.
type PayableUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PayableMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updatedAt" field.
func (puo *PayableUpdateOne) SetUpdatedAt(t time.Time) *PayableUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDeletedAt sets the "deletedAt" field.
func (puo *PayableUpdateOne) SetDeletedAt(t time.Time) *PayableUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableDeletedAt(t *time.Time) *PayableUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (puo *PayableUpdateOne) ClearDeletedAt() *PayableUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetEntryGroup sets the "entryGroup" field.
func (puo *PayableUpdateOne) SetEntryGroup(i int) *PayableUpdateOne {
	puo.mutation.ResetEntryGroup()
	puo.mutation.SetEntryGroup(i)
	return puo
}

// SetNillableEntryGroup sets the "entryGroup" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableEntryGroup(i *int) *PayableUpdateOne {
	if i != nil {
		puo.SetEntryGroup(*i)
	}
	return puo
}

// AddEntryGroup adds i to the "entryGroup" field.
func (puo *PayableUpdateOne) AddEntryGroup(i int) *PayableUpdateOne {
	puo.mutation.AddEntryGroup(i)
	return puo
}

// SetDate sets the "date" field.
func (puo *PayableUpdateOne) SetDate(t time.Time) *PayableUpdateOne {
	puo.mutation.SetDate(t)
	return puo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableDate(t *time.Time) *PayableUpdateOne {
	if t != nil {
		puo.SetDate(*t)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PayableUpdateOne) SetName(s string) *PayableUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableName(s *string) *PayableUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetOutstandingBalance sets the "outstandingBalance" field.
func (puo *PayableUpdateOne) SetOutstandingBalance(f float64) *PayableUpdateOne {
	puo.mutation.ResetOutstandingBalance()
	puo.mutation.SetOutstandingBalance(f)
	return puo
}

// SetNillableOutstandingBalance sets the "outstandingBalance" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableOutstandingBalance(f *float64) *PayableUpdateOne {
	if f != nil {
		puo.SetOutstandingBalance(*f)
	}
	return puo
}

// AddOutstandingBalance adds f to the "outstandingBalance" field.
func (puo *PayableUpdateOne) AddOutstandingBalance(f float64) *PayableUpdateOne {
	puo.mutation.AddOutstandingBalance(f)
	return puo
}

// SetTotalTransaction sets the "totalTransaction" field.
func (puo *PayableUpdateOne) SetTotalTransaction(f float64) *PayableUpdateOne {
	puo.mutation.ResetTotalTransaction()
	puo.mutation.SetTotalTransaction(f)
	return puo
}

// SetNillableTotalTransaction sets the "totalTransaction" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableTotalTransaction(f *float64) *PayableUpdateOne {
	if f != nil {
		puo.SetTotalTransaction(*f)
	}
	return puo
}

// AddTotalTransaction adds f to the "totalTransaction" field.
func (puo *PayableUpdateOne) AddTotalTransaction(f float64) *PayableUpdateOne {
	puo.mutation.AddTotalTransaction(f)
	return puo
}

// SetDueDate sets the "dueDate" field.
func (puo *PayableUpdateOne) SetDueDate(t time.Time) *PayableUpdateOne {
	puo.mutation.SetDueDate(t)
	return puo
}

// SetNillableDueDate sets the "dueDate" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableDueDate(t *time.Time) *PayableUpdateOne {
	if t != nil {
		puo.SetDueDate(*t)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PayableUpdateOne) SetStatus(pa payable.Status) *PayableUpdateOne {
	puo.mutation.SetStatus(pa)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableStatus(pa *payable.Status) *PayableUpdateOne {
	if pa != nil {
		puo.SetStatus(*pa)
	}
	return puo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (puo *PayableUpdateOne) SetCompanyID(id int) *PayableUpdateOne {
	puo.mutation.SetCompanyID(id)
	return puo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (puo *PayableUpdateOne) SetNillableCompanyID(id *int) *PayableUpdateOne {
	if id != nil {
		puo = puo.SetCompanyID(*id)
	}
	return puo
}

// SetCompany sets the "company" edge to the Company entity.
func (puo *PayableUpdateOne) SetCompany(c *Company) *PayableUpdateOne {
	return puo.SetCompanyID(c.ID)
}

// Mutation returns the PayableMutation object of the builder.
func (puo *PayableUpdateOne) Mutation() *PayableMutation {
	return puo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (puo *PayableUpdateOne) ClearCompany() *PayableUpdateOne {
	puo.mutation.ClearCompany()
	return puo
}

// Where appends a list predicates to the PayableUpdate builder.
func (puo *PayableUpdateOne) Where(ps ...predicate.Payable) *PayableUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PayableUpdateOne) Select(field string, fields ...string) *PayableUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payable entity.
func (puo *PayableUpdateOne) Save(ctx context.Context) (*Payable, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PayableUpdateOne) SaveX(ctx context.Context) *Payable {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PayableUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PayableUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PayableUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := payable.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PayableUpdateOne) check() error {
	if v, ok := puo.mutation.EntryGroup(); ok {
		if err := payable.EntryGroupValidator(v); err != nil {
			return &ValidationError{Name: "entryGroup", err: fmt.Errorf(`generated: validator failed for field "Payable.entryGroup": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := payable.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Payable.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PayableUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PayableUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PayableUpdateOne) sqlSave(ctx context.Context) (_node *Payable, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payable.Table, payable.Columns, sqlgraph.NewFieldSpec(payable.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Payable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payable.FieldID)
		for _, f := range fields {
			if !payable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != payable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(payable.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(payable.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(payable.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.EntryGroup(); ok {
		_spec.SetField(payable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedEntryGroup(); ok {
		_spec.AddField(payable.FieldEntryGroup, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Date(); ok {
		_spec.SetField(payable.FieldDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(payable.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.OutstandingBalance(); ok {
		_spec.SetField(payable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedOutstandingBalance(); ok {
		_spec.AddField(payable.FieldOutstandingBalance, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.TotalTransaction(); ok {
		_spec.SetField(payable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedTotalTransaction(); ok {
		_spec.AddField(payable.FieldTotalTransaction, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.DueDate(); ok {
		_spec.SetField(payable.FieldDueDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(payable.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payable.CompanyTable,
			Columns: []string{payable.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payable.CompanyTable,
			Columns: []string{payable.CompanyColumn},
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
	_spec.AddModifiers(puo.modifiers...)
	_node = &Payable{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
