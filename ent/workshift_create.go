// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/company"
	"mazza/ent/employee"
	"mazza/ent/workshift"
	"mazza/ent/worktask"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkshiftCreate is the builder for creating a Workshift entity.
type WorkshiftCreate struct {
	config
	mutation *WorkshiftMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (wc *WorkshiftCreate) SetCreatedAt(t time.Time) *WorkshiftCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableCreatedAt(t *time.Time) *WorkshiftCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updatedAt" field.
func (wc *WorkshiftCreate) SetUpdatedAt(t time.Time) *WorkshiftCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableUpdatedAt(t *time.Time) *WorkshiftCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetDeletedAt sets the "deletedAt" field.
func (wc *WorkshiftCreate) SetDeletedAt(t time.Time) *WorkshiftCreate {
	wc.mutation.SetDeletedAt(t)
	return wc
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableDeletedAt(t *time.Time) *WorkshiftCreate {
	if t != nil {
		wc.SetDeletedAt(*t)
	}
	return wc
}

// SetApprovedAt sets the "approvedAt" field.
func (wc *WorkshiftCreate) SetApprovedAt(t time.Time) *WorkshiftCreate {
	wc.mutation.SetApprovedAt(t)
	return wc
}

// SetNillableApprovedAt sets the "approvedAt" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableApprovedAt(t *time.Time) *WorkshiftCreate {
	if t != nil {
		wc.SetApprovedAt(*t)
	}
	return wc
}

// SetClockIn sets the "clockIn" field.
func (wc *WorkshiftCreate) SetClockIn(t time.Time) *WorkshiftCreate {
	wc.mutation.SetClockIn(t)
	return wc
}

// SetNillableClockIn sets the "clockIn" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableClockIn(t *time.Time) *WorkshiftCreate {
	if t != nil {
		wc.SetClockIn(*t)
	}
	return wc
}

// SetClockOut sets the "clockOut" field.
func (wc *WorkshiftCreate) SetClockOut(t time.Time) *WorkshiftCreate {
	wc.mutation.SetClockOut(t)
	return wc
}

// SetNillableClockOut sets the "clockOut" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableClockOut(t *time.Time) *WorkshiftCreate {
	if t != nil {
		wc.SetClockOut(*t)
	}
	return wc
}

// SetClockInLocation sets the "clockInLocation" field.
func (wc *WorkshiftCreate) SetClockInLocation(s string) *WorkshiftCreate {
	wc.mutation.SetClockInLocation(s)
	return wc
}

// SetClockOutLocation sets the "clockOutLocation" field.
func (wc *WorkshiftCreate) SetClockOutLocation(s string) *WorkshiftCreate {
	wc.mutation.SetClockOutLocation(s)
	return wc
}

// SetNillableClockOutLocation sets the "clockOutLocation" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableClockOutLocation(s *string) *WorkshiftCreate {
	if s != nil {
		wc.SetClockOutLocation(*s)
	}
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkshiftCreate) SetDescription(s string) *WorkshiftCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableDescription(s *string) *WorkshiftCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetNote sets the "note" field.
func (wc *WorkshiftCreate) SetNote(s string) *WorkshiftCreate {
	wc.mutation.SetNote(s)
	return wc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableNote(s *string) *WorkshiftCreate {
	if s != nil {
		wc.SetNote(*s)
	}
	return wc
}

// SetStatus sets the "status" field.
func (wc *WorkshiftCreate) SetStatus(w workshift.Status) *WorkshiftCreate {
	wc.mutation.SetStatus(w)
	return wc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableStatus(w *workshift.Status) *WorkshiftCreate {
	if w != nil {
		wc.SetStatus(*w)
	}
	return wc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (wc *WorkshiftCreate) SetCompanyID(id int) *WorkshiftCreate {
	wc.mutation.SetCompanyID(id)
	return wc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableCompanyID(id *int) *WorkshiftCreate {
	if id != nil {
		wc = wc.SetCompanyID(*id)
	}
	return wc
}

// SetCompany sets the "company" edge to the Company entity.
func (wc *WorkshiftCreate) SetCompany(c *Company) *WorkshiftCreate {
	return wc.SetCompanyID(c.ID)
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (wc *WorkshiftCreate) SetEmployeeID(id int) *WorkshiftCreate {
	wc.mutation.SetEmployeeID(id)
	return wc
}

// SetNillableEmployeeID sets the "employee" edge to the Employee entity by ID if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableEmployeeID(id *int) *WorkshiftCreate {
	if id != nil {
		wc = wc.SetEmployeeID(*id)
	}
	return wc
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (wc *WorkshiftCreate) SetEmployee(e *Employee) *WorkshiftCreate {
	return wc.SetEmployeeID(e.ID)
}

// SetApprovedByID sets the "approvedBy" edge to the Employee entity by ID.
func (wc *WorkshiftCreate) SetApprovedByID(id int) *WorkshiftCreate {
	wc.mutation.SetApprovedByID(id)
	return wc
}

// SetNillableApprovedByID sets the "approvedBy" edge to the Employee entity by ID if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableApprovedByID(id *int) *WorkshiftCreate {
	if id != nil {
		wc = wc.SetApprovedByID(*id)
	}
	return wc
}

// SetApprovedBy sets the "approvedBy" edge to the Employee entity.
func (wc *WorkshiftCreate) SetApprovedBy(e *Employee) *WorkshiftCreate {
	return wc.SetApprovedByID(e.ID)
}

// SetWorkTaskID sets the "workTask" edge to the Worktask entity by ID.
func (wc *WorkshiftCreate) SetWorkTaskID(id int) *WorkshiftCreate {
	wc.mutation.SetWorkTaskID(id)
	return wc
}

// SetNillableWorkTaskID sets the "workTask" edge to the Worktask entity by ID if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableWorkTaskID(id *int) *WorkshiftCreate {
	if id != nil {
		wc = wc.SetWorkTaskID(*id)
	}
	return wc
}

// SetWorkTask sets the "workTask" edge to the Worktask entity.
func (wc *WorkshiftCreate) SetWorkTask(w *Worktask) *WorkshiftCreate {
	return wc.SetWorkTaskID(w.ID)
}

// SetEditRequestID sets the "editRequest" edge to the Workshift entity by ID.
func (wc *WorkshiftCreate) SetEditRequestID(id int) *WorkshiftCreate {
	wc.mutation.SetEditRequestID(id)
	return wc
}

// SetNillableEditRequestID sets the "editRequest" edge to the Workshift entity by ID if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableEditRequestID(id *int) *WorkshiftCreate {
	if id != nil {
		wc = wc.SetEditRequestID(*id)
	}
	return wc
}

// SetEditRequest sets the "editRequest" edge to the Workshift entity.
func (wc *WorkshiftCreate) SetEditRequest(w *Workshift) *WorkshiftCreate {
	return wc.SetEditRequestID(w.ID)
}

// SetWorkShiftID sets the "workShift" edge to the Workshift entity by ID.
func (wc *WorkshiftCreate) SetWorkShiftID(id int) *WorkshiftCreate {
	wc.mutation.SetWorkShiftID(id)
	return wc
}

// SetNillableWorkShiftID sets the "workShift" edge to the Workshift entity by ID if the given value is not nil.
func (wc *WorkshiftCreate) SetNillableWorkShiftID(id *int) *WorkshiftCreate {
	if id != nil {
		wc = wc.SetWorkShiftID(*id)
	}
	return wc
}

// SetWorkShift sets the "workShift" edge to the Workshift entity.
func (wc *WorkshiftCreate) SetWorkShift(w *Workshift) *WorkshiftCreate {
	return wc.SetWorkShiftID(w.ID)
}

// Mutation returns the WorkshiftMutation object of the builder.
func (wc *WorkshiftCreate) Mutation() *WorkshiftMutation {
	return wc.mutation
}

// Save creates the Workshift in the database.
func (wc *WorkshiftCreate) Save(ctx context.Context) (*Workshift, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkshiftCreate) SaveX(ctx context.Context) *Workshift {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkshiftCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkshiftCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkshiftCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workshift.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := workshift.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.ClockIn(); !ok {
		v := workshift.DefaultClockIn()
		wc.mutation.SetClockIn(v)
	}
	if _, ok := wc.mutation.Status(); !ok {
		v := workshift.DefaultStatus
		wc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkshiftCreate) check() error {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Workshift.createdAt"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Workshift.updatedAt"`)}
	}
	if _, ok := wc.mutation.ClockIn(); !ok {
		return &ValidationError{Name: "clockIn", err: errors.New(`ent: missing required field "Workshift.clockIn"`)}
	}
	if _, ok := wc.mutation.ClockInLocation(); !ok {
		return &ValidationError{Name: "clockInLocation", err: errors.New(`ent: missing required field "Workshift.clockInLocation"`)}
	}
	if _, ok := wc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Workshift.status"`)}
	}
	if v, ok := wc.mutation.Status(); ok {
		if err := workshift.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Workshift.status": %w`, err)}
		}
	}
	return nil
}

func (wc *WorkshiftCreate) sqlSave(ctx context.Context) (*Workshift, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkshiftCreate) createSpec() (*Workshift, *sqlgraph.CreateSpec) {
	var (
		_node = &Workshift{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workshift.Table, sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workshift.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(workshift.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wc.mutation.DeletedAt(); ok {
		_spec.SetField(workshift.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := wc.mutation.ApprovedAt(); ok {
		_spec.SetField(workshift.FieldApprovedAt, field.TypeTime, value)
		_node.ApprovedAt = &value
	}
	if value, ok := wc.mutation.ClockIn(); ok {
		_spec.SetField(workshift.FieldClockIn, field.TypeTime, value)
		_node.ClockIn = value
	}
	if value, ok := wc.mutation.ClockOut(); ok {
		_spec.SetField(workshift.FieldClockOut, field.TypeTime, value)
		_node.ClockOut = &value
	}
	if value, ok := wc.mutation.ClockInLocation(); ok {
		_spec.SetField(workshift.FieldClockInLocation, field.TypeString, value)
		_node.ClockInLocation = value
	}
	if value, ok := wc.mutation.ClockOutLocation(); ok {
		_spec.SetField(workshift.FieldClockOutLocation, field.TypeString, value)
		_node.ClockOutLocation = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(workshift.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.Note(); ok {
		_spec.SetField(workshift.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := wc.mutation.Status(); ok {
		_spec.SetField(workshift.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := wc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workshift.CompanyTable,
			Columns: []string{workshift.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_work_shifts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workshift.EmployeeTable,
			Columns: []string{workshift.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.employee_work_shifts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ApprovedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workshift.ApprovedByTable,
			Columns: []string{workshift.ApprovedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.employee_approved_work_shifts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workshift.WorkTaskTable,
			Columns: []string{workshift.WorkTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.worktask_work_shifts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.EditRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workshift.EditRequestTable,
			Columns: []string{workshift.EditRequestColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workshift_edit_request = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkShiftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workshift.WorkShiftTable,
			Columns: []string{workshift.WorkShiftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workshift_edit_request = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkshiftCreateBulk is the builder for creating many Workshift entities in bulk.
type WorkshiftCreateBulk struct {
	config
	err      error
	builders []*WorkshiftCreate
}

// Save creates the Workshift entities in the database.
func (wcb *WorkshiftCreateBulk) Save(ctx context.Context) ([]*Workshift, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workshift, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkshiftMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkshiftCreateBulk) SaveX(ctx context.Context) []*Workshift {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkshiftCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkshiftCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
