// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks     []Hook
	mutation  *EmployeeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updatedAt" field.
func (eu *EmployeeUpdate) SetUpdatedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetDeletedAt sets the "deletedAt" field.
func (eu *EmployeeUpdate) SetDeletedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetDeletedAt(t)
	return eu
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDeletedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetDeletedAt(*t)
	}
	return eu
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (eu *EmployeeUpdate) ClearDeletedAt() *EmployeeUpdate {
	eu.mutation.ClearDeletedAt()
	return eu
}

// SetName sets the "name" field.
func (eu *EmployeeUpdate) SetName(s string) *EmployeeUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableName(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetGender sets the "gender" field.
func (eu *EmployeeUpdate) SetGender(e employee.Gender) *EmployeeUpdate {
	eu.mutation.SetGender(e)
	return eu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableGender(e *employee.Gender) *EmployeeUpdate {
	if e != nil {
		eu.SetGender(*e)
	}
	return eu
}

// SetPosition sets the "position" field.
func (eu *EmployeeUpdate) SetPosition(s string) *EmployeeUpdate {
	eu.mutation.SetPosition(s)
	return eu
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillablePosition(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetPosition(*s)
	}
	return eu
}

// ClearPosition clears the value of the "position" field.
func (eu *EmployeeUpdate) ClearPosition() *EmployeeUpdate {
	eu.mutation.ClearPosition()
	return eu
}

// SetEmail sets the "email" field.
func (eu *EmployeeUpdate) SetEmail(s string) *EmployeeUpdate {
	eu.mutation.SetEmail(s)
	return eu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableEmail(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetEmail(*s)
	}
	return eu
}

// ClearEmail clears the value of the "email" field.
func (eu *EmployeeUpdate) ClearEmail() *EmployeeUpdate {
	eu.mutation.ClearEmail()
	return eu
}

// SetPhone sets the "phone" field.
func (eu *EmployeeUpdate) SetPhone(s string) *EmployeeUpdate {
	eu.mutation.SetPhone(s)
	return eu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillablePhone(s *string) *EmployeeUpdate {
	if s != nil {
		eu.SetPhone(*s)
	}
	return eu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (eu *EmployeeUpdate) SetCompanyID(id int) *EmployeeUpdate {
	eu.mutation.SetCompanyID(id)
	return eu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableCompanyID(id *int) *EmployeeUpdate {
	if id != nil {
		eu = eu.SetCompanyID(*id)
	}
	return eu
}

// SetCompany sets the "company" edge to the Company entity.
func (eu *EmployeeUpdate) SetCompany(c *Company) *EmployeeUpdate {
	return eu.SetCompanyID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (eu *EmployeeUpdate) SetUserID(id int) *EmployeeUpdate {
	eu.mutation.SetUserID(id)
	return eu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableUserID(id *int) *EmployeeUpdate {
	if id != nil {
		eu = eu.SetUserID(*id)
	}
	return eu
}

// SetUser sets the "user" edge to the User entity.
func (eu *EmployeeUpdate) SetUser(u *User) *EmployeeUpdate {
	return eu.SetUserID(u.ID)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (eu *EmployeeUpdate) ClearCompany() *EmployeeUpdate {
	eu.mutation.ClearCompany()
	return eu
}

// ClearUser clears the "user" edge to the User entity.
func (eu *EmployeeUpdate) ClearUser() *EmployeeUpdate {
	eu.mutation.ClearUser()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	eu.defaults()
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EmployeeUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := employee.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmployeeUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := employee.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Employee.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Gender(); ok {
		if err := employee.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`generated: validator failed for field "Employee.gender": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eu *EmployeeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EmployeeUpdate {
	eu.modifiers = append(eu.modifiers, modifiers...)
	return eu
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
	}
	if eu.mutation.DeletedAtCleared() {
		_spec.ClearField(employee.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(employee.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeEnum, value)
	}
	if value, ok := eu.mutation.Position(); ok {
		_spec.SetField(employee.FieldPosition, field.TypeString, value)
	}
	if eu.mutation.PositionCleared() {
		_spec.ClearField(employee.FieldPosition, field.TypeString)
	}
	if value, ok := eu.mutation.Email(); ok {
		_spec.SetField(employee.FieldEmail, field.TypeString, value)
	}
	if eu.mutation.EmailCleared() {
		_spec.ClearField(employee.FieldEmail, field.TypeString)
	}
	if value, ok := eu.mutation.Phone(); ok {
		_spec.SetField(employee.FieldPhone, field.TypeString, value)
	}
	if eu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.CompanyTable,
			Columns: []string{employee.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.CompanyTable,
			Columns: []string{employee.CompanyColumn},
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
	if eu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.UserTable,
			Columns: []string{employee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.UserTable,
			Columns: []string{employee.UserColumn},
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
	_spec.AddModifiers(eu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EmployeeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updatedAt" field.
func (euo *EmployeeUpdateOne) SetUpdatedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetDeletedAt sets the "deletedAt" field.
func (euo *EmployeeUpdateOne) SetDeletedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetDeletedAt(t)
	return euo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDeletedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetDeletedAt(*t)
	}
	return euo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (euo *EmployeeUpdateOne) ClearDeletedAt() *EmployeeUpdateOne {
	euo.mutation.ClearDeletedAt()
	return euo
}

// SetName sets the "name" field.
func (euo *EmployeeUpdateOne) SetName(s string) *EmployeeUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableName(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetGender sets the "gender" field.
func (euo *EmployeeUpdateOne) SetGender(e employee.Gender) *EmployeeUpdateOne {
	euo.mutation.SetGender(e)
	return euo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableGender(e *employee.Gender) *EmployeeUpdateOne {
	if e != nil {
		euo.SetGender(*e)
	}
	return euo
}

// SetPosition sets the "position" field.
func (euo *EmployeeUpdateOne) SetPosition(s string) *EmployeeUpdateOne {
	euo.mutation.SetPosition(s)
	return euo
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillablePosition(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetPosition(*s)
	}
	return euo
}

// ClearPosition clears the value of the "position" field.
func (euo *EmployeeUpdateOne) ClearPosition() *EmployeeUpdateOne {
	euo.mutation.ClearPosition()
	return euo
}

// SetEmail sets the "email" field.
func (euo *EmployeeUpdateOne) SetEmail(s string) *EmployeeUpdateOne {
	euo.mutation.SetEmail(s)
	return euo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableEmail(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetEmail(*s)
	}
	return euo
}

// ClearEmail clears the value of the "email" field.
func (euo *EmployeeUpdateOne) ClearEmail() *EmployeeUpdateOne {
	euo.mutation.ClearEmail()
	return euo
}

// SetPhone sets the "phone" field.
func (euo *EmployeeUpdateOne) SetPhone(s string) *EmployeeUpdateOne {
	euo.mutation.SetPhone(s)
	return euo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillablePhone(s *string) *EmployeeUpdateOne {
	if s != nil {
		euo.SetPhone(*s)
	}
	return euo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (euo *EmployeeUpdateOne) SetCompanyID(id int) *EmployeeUpdateOne {
	euo.mutation.SetCompanyID(id)
	return euo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableCompanyID(id *int) *EmployeeUpdateOne {
	if id != nil {
		euo = euo.SetCompanyID(*id)
	}
	return euo
}

// SetCompany sets the "company" edge to the Company entity.
func (euo *EmployeeUpdateOne) SetCompany(c *Company) *EmployeeUpdateOne {
	return euo.SetCompanyID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (euo *EmployeeUpdateOne) SetUserID(id int) *EmployeeUpdateOne {
	euo.mutation.SetUserID(id)
	return euo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableUserID(id *int) *EmployeeUpdateOne {
	if id != nil {
		euo = euo.SetUserID(*id)
	}
	return euo
}

// SetUser sets the "user" edge to the User entity.
func (euo *EmployeeUpdateOne) SetUser(u *User) *EmployeeUpdateOne {
	return euo.SetUserID(u.ID)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (euo *EmployeeUpdateOne) ClearCompany() *EmployeeUpdateOne {
	euo.mutation.ClearCompany()
	return euo
}

// ClearUser clears the "user" edge to the User entity.
func (euo *EmployeeUpdateOne) ClearUser() *EmployeeUpdateOne {
	euo.mutation.ClearUser()
	return euo
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (euo *EmployeeUpdateOne) Where(ps ...predicate.Employee) *EmployeeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmployeeUpdateOne) Select(field string, fields ...string) *EmployeeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Employee entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	euo.defaults()
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EmployeeUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := employee.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmployeeUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := employee.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Employee.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Gender(); ok {
		if err := employee.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`generated: validator failed for field "Employee.gender": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (euo *EmployeeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EmployeeUpdateOne {
	euo.modifiers = append(euo.modifiers, modifiers...)
	return euo
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (_node *Employee, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Employee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for _, f := range fields {
			if !employee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
	}
	if euo.mutation.DeletedAtCleared() {
		_spec.ClearField(employee.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(employee.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeEnum, value)
	}
	if value, ok := euo.mutation.Position(); ok {
		_spec.SetField(employee.FieldPosition, field.TypeString, value)
	}
	if euo.mutation.PositionCleared() {
		_spec.ClearField(employee.FieldPosition, field.TypeString)
	}
	if value, ok := euo.mutation.Email(); ok {
		_spec.SetField(employee.FieldEmail, field.TypeString, value)
	}
	if euo.mutation.EmailCleared() {
		_spec.ClearField(employee.FieldEmail, field.TypeString)
	}
	if value, ok := euo.mutation.Phone(); ok {
		_spec.SetField(employee.FieldPhone, field.TypeString, value)
	}
	if euo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.CompanyTable,
			Columns: []string{employee.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.CompanyTable,
			Columns: []string{employee.CompanyColumn},
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
	if euo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.UserTable,
			Columns: []string{employee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   employee.UserTable,
			Columns: []string{employee.UserColumn},
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
	_spec.AddModifiers(euo.modifiers...)
	_node = &Employee{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
