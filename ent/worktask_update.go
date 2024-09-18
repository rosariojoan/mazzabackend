// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/company"
	"mazza/ent/employee"
	"mazza/ent/predicate"
	"mazza/ent/user"
	"mazza/ent/workshift"
	"mazza/ent/worktag"
	"mazza/ent/worktask"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// WorktaskUpdate is the builder for updating Worktask entities.
type WorktaskUpdate struct {
	config
	hooks    []Hook
	mutation *WorktaskMutation
}

// Where appends a list predicates to the WorktaskUpdate builder.
func (wu *WorktaskUpdate) Where(ps ...predicate.Worktask) *WorktaskUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetUpdatedAt sets the "updatedAt" field.
func (wu *WorktaskUpdate) SetUpdatedAt(t time.Time) *WorktaskUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetDeletedAt sets the "deletedAt" field.
func (wu *WorktaskUpdate) SetDeletedAt(t time.Time) *WorktaskUpdate {
	wu.mutation.SetDeletedAt(t)
	return wu
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableDeletedAt(t *time.Time) *WorktaskUpdate {
	if t != nil {
		wu.SetDeletedAt(*t)
	}
	return wu
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (wu *WorktaskUpdate) ClearDeletedAt() *WorktaskUpdate {
	wu.mutation.ClearDeletedAt()
	return wu
}

// SetDescription sets the "description" field.
func (wu *WorktaskUpdate) SetDescription(s string) *WorktaskUpdate {
	wu.mutation.SetDescription(s)
	return wu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableDescription(s *string) *WorktaskUpdate {
	if s != nil {
		wu.SetDescription(*s)
	}
	return wu
}

// ClearDescription clears the value of the "description" field.
func (wu *WorktaskUpdate) ClearDescription() *WorktaskUpdate {
	wu.mutation.ClearDescription()
	return wu
}

// SetStatus sets the "status" field.
func (wu *WorktaskUpdate) SetStatus(w worktask.Status) *WorktaskUpdate {
	wu.mutation.SetStatus(w)
	return wu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableStatus(w *worktask.Status) *WorktaskUpdate {
	if w != nil {
		wu.SetStatus(*w)
	}
	return wu
}

// SetSubtasks sets the "subtasks" field.
func (wu *WorktaskUpdate) SetSubtasks(s []string) *WorktaskUpdate {
	wu.mutation.SetSubtasks(s)
	return wu
}

// AppendSubtasks appends s to the "subtasks" field.
func (wu *WorktaskUpdate) AppendSubtasks(s []string) *WorktaskUpdate {
	wu.mutation.AppendSubtasks(s)
	return wu
}

// ClearSubtasks clears the value of the "subtasks" field.
func (wu *WorktaskUpdate) ClearSubtasks() *WorktaskUpdate {
	wu.mutation.ClearSubtasks()
	return wu
}

// SetTitle sets the "title" field.
func (wu *WorktaskUpdate) SetTitle(s string) *WorktaskUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableTitle(s *string) *WorktaskUpdate {
	if s != nil {
		wu.SetTitle(*s)
	}
	return wu
}

// SetStartTime sets the "startTime" field.
func (wu *WorktaskUpdate) SetStartTime(t time.Time) *WorktaskUpdate {
	wu.mutation.SetStartTime(t)
	return wu
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableStartTime(t *time.Time) *WorktaskUpdate {
	if t != nil {
		wu.SetStartTime(*t)
	}
	return wu
}

// SetEndTime sets the "endTime" field.
func (wu *WorktaskUpdate) SetEndTime(t time.Time) *WorktaskUpdate {
	wu.mutation.SetEndTime(t)
	return wu
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableEndTime(t *time.Time) *WorktaskUpdate {
	if t != nil {
		wu.SetEndTime(*t)
	}
	return wu
}

// ClearEndTime clears the value of the "endTime" field.
func (wu *WorktaskUpdate) ClearEndTime() *WorktaskUpdate {
	wu.mutation.ClearEndTime()
	return wu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (wu *WorktaskUpdate) SetCompanyID(id int) *WorktaskUpdate {
	wu.mutation.SetCompanyID(id)
	return wu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableCompanyID(id *int) *WorktaskUpdate {
	if id != nil {
		wu = wu.SetCompanyID(*id)
	}
	return wu
}

// SetCompany sets the "company" edge to the Company entity.
func (wu *WorktaskUpdate) SetCompany(c *Company) *WorktaskUpdate {
	return wu.SetCompanyID(c.ID)
}

// SetCreatedByID sets the "createdBy" edge to the User entity by ID.
func (wu *WorktaskUpdate) SetCreatedByID(id int) *WorktaskUpdate {
	wu.mutation.SetCreatedByID(id)
	return wu
}

// SetNillableCreatedByID sets the "createdBy" edge to the User entity by ID if the given value is not nil.
func (wu *WorktaskUpdate) SetNillableCreatedByID(id *int) *WorktaskUpdate {
	if id != nil {
		wu = wu.SetCreatedByID(*id)
	}
	return wu
}

// SetCreatedBy sets the "createdBy" edge to the User entity.
func (wu *WorktaskUpdate) SetCreatedBy(u *User) *WorktaskUpdate {
	return wu.SetCreatedByID(u.ID)
}

// AddAssignedToIDs adds the "assignedTo" edge to the Employee entity by IDs.
func (wu *WorktaskUpdate) AddAssignedToIDs(ids ...int) *WorktaskUpdate {
	wu.mutation.AddAssignedToIDs(ids...)
	return wu
}

// AddAssignedTo adds the "assignedTo" edges to the Employee entity.
func (wu *WorktaskUpdate) AddAssignedTo(e ...*Employee) *WorktaskUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wu.AddAssignedToIDs(ids...)
}

// AddWorkShiftIDs adds the "workShifts" edge to the Workshift entity by IDs.
func (wu *WorktaskUpdate) AddWorkShiftIDs(ids ...int) *WorktaskUpdate {
	wu.mutation.AddWorkShiftIDs(ids...)
	return wu
}

// AddWorkShifts adds the "workShifts" edges to the Workshift entity.
func (wu *WorktaskUpdate) AddWorkShifts(w ...*Workshift) *WorktaskUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWorkShiftIDs(ids...)
}

// AddWorkTagIDs adds the "workTags" edge to the Worktag entity by IDs.
func (wu *WorktaskUpdate) AddWorkTagIDs(ids ...int) *WorktaskUpdate {
	wu.mutation.AddWorkTagIDs(ids...)
	return wu
}

// AddWorkTags adds the "workTags" edges to the Worktag entity.
func (wu *WorktaskUpdate) AddWorkTags(w ...*Worktag) *WorktaskUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWorkTagIDs(ids...)
}

// Mutation returns the WorktaskMutation object of the builder.
func (wu *WorktaskUpdate) Mutation() *WorktaskMutation {
	return wu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (wu *WorktaskUpdate) ClearCompany() *WorktaskUpdate {
	wu.mutation.ClearCompany()
	return wu
}

// ClearCreatedBy clears the "createdBy" edge to the User entity.
func (wu *WorktaskUpdate) ClearCreatedBy() *WorktaskUpdate {
	wu.mutation.ClearCreatedBy()
	return wu
}

// ClearAssignedTo clears all "assignedTo" edges to the Employee entity.
func (wu *WorktaskUpdate) ClearAssignedTo() *WorktaskUpdate {
	wu.mutation.ClearAssignedTo()
	return wu
}

// RemoveAssignedToIDs removes the "assignedTo" edge to Employee entities by IDs.
func (wu *WorktaskUpdate) RemoveAssignedToIDs(ids ...int) *WorktaskUpdate {
	wu.mutation.RemoveAssignedToIDs(ids...)
	return wu
}

// RemoveAssignedTo removes "assignedTo" edges to Employee entities.
func (wu *WorktaskUpdate) RemoveAssignedTo(e ...*Employee) *WorktaskUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wu.RemoveAssignedToIDs(ids...)
}

// ClearWorkShifts clears all "workShifts" edges to the Workshift entity.
func (wu *WorktaskUpdate) ClearWorkShifts() *WorktaskUpdate {
	wu.mutation.ClearWorkShifts()
	return wu
}

// RemoveWorkShiftIDs removes the "workShifts" edge to Workshift entities by IDs.
func (wu *WorktaskUpdate) RemoveWorkShiftIDs(ids ...int) *WorktaskUpdate {
	wu.mutation.RemoveWorkShiftIDs(ids...)
	return wu
}

// RemoveWorkShifts removes "workShifts" edges to Workshift entities.
func (wu *WorktaskUpdate) RemoveWorkShifts(w ...*Workshift) *WorktaskUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWorkShiftIDs(ids...)
}

// ClearWorkTags clears all "workTags" edges to the Worktag entity.
func (wu *WorktaskUpdate) ClearWorkTags() *WorktaskUpdate {
	wu.mutation.ClearWorkTags()
	return wu
}

// RemoveWorkTagIDs removes the "workTags" edge to Worktag entities by IDs.
func (wu *WorktaskUpdate) RemoveWorkTagIDs(ids ...int) *WorktaskUpdate {
	wu.mutation.RemoveWorkTagIDs(ids...)
	return wu
}

// RemoveWorkTags removes "workTags" edges to Worktag entities.
func (wu *WorktaskUpdate) RemoveWorkTags(w ...*Worktag) *WorktaskUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWorkTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorktaskUpdate) Save(ctx context.Context) (int, error) {
	wu.defaults()
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorktaskUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorktaskUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorktaskUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WorktaskUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := worktask.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorktaskUpdate) check() error {
	if v, ok := wu.mutation.Status(); ok {
		if err := worktask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Worktask.status": %w`, err)}
		}
	}
	return nil
}

func (wu *WorktaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(worktask.Table, worktask.Columns, sqlgraph.NewFieldSpec(worktask.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.SetField(worktask.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wu.mutation.DeletedAt(); ok {
		_spec.SetField(worktask.FieldDeletedAt, field.TypeTime, value)
	}
	if wu.mutation.DeletedAtCleared() {
		_spec.ClearField(worktask.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := wu.mutation.Description(); ok {
		_spec.SetField(worktask.FieldDescription, field.TypeString, value)
	}
	if wu.mutation.DescriptionCleared() {
		_spec.ClearField(worktask.FieldDescription, field.TypeString)
	}
	if value, ok := wu.mutation.Status(); ok {
		_spec.SetField(worktask.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := wu.mutation.Subtasks(); ok {
		_spec.SetField(worktask.FieldSubtasks, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedSubtasks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, worktask.FieldSubtasks, value)
		})
	}
	if wu.mutation.SubtasksCleared() {
		_spec.ClearField(worktask.FieldSubtasks, field.TypeJSON)
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.SetField(worktask.FieldTitle, field.TypeString, value)
	}
	if value, ok := wu.mutation.StartTime(); ok {
		_spec.SetField(worktask.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := wu.mutation.EndTime(); ok {
		_spec.SetField(worktask.FieldEndTime, field.TypeTime, value)
	}
	if wu.mutation.EndTimeCleared() {
		_spec.ClearField(worktask.FieldEndTime, field.TypeTime)
	}
	if wu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CompanyTable,
			Columns: []string{worktask.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CompanyTable,
			Columns: []string{worktask.CompanyColumn},
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
	if wu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CreatedByTable,
			Columns: []string{worktask.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CreatedByTable,
			Columns: []string{worktask.CreatedByColumn},
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
	if wu.mutation.AssignedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   worktask.AssignedToTable,
			Columns: worktask.AssignedToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedAssignedToIDs(); len(nodes) > 0 && !wu.mutation.AssignedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   worktask.AssignedToTable,
			Columns: worktask.AssignedToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.AssignedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   worktask.AssignedToTable,
			Columns: worktask.AssignedToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.WorkShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   worktask.WorkShiftsTable,
			Columns: []string{worktask.WorkShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWorkShiftsIDs(); len(nodes) > 0 && !wu.mutation.WorkShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   worktask.WorkShiftsTable,
			Columns: []string{worktask.WorkShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WorkShiftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   worktask.WorkShiftsTable,
			Columns: []string{worktask.WorkShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.WorkTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   worktask.WorkTagsTable,
			Columns: worktask.WorkTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWorkTagsIDs(); len(nodes) > 0 && !wu.mutation.WorkTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   worktask.WorkTagsTable,
			Columns: worktask.WorkTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WorkTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   worktask.WorkTagsTable,
			Columns: worktask.WorkTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{worktask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WorktaskUpdateOne is the builder for updating a single Worktask entity.
type WorktaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorktaskMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (wuo *WorktaskUpdateOne) SetUpdatedAt(t time.Time) *WorktaskUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetDeletedAt sets the "deletedAt" field.
func (wuo *WorktaskUpdateOne) SetDeletedAt(t time.Time) *WorktaskUpdateOne {
	wuo.mutation.SetDeletedAt(t)
	return wuo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableDeletedAt(t *time.Time) *WorktaskUpdateOne {
	if t != nil {
		wuo.SetDeletedAt(*t)
	}
	return wuo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (wuo *WorktaskUpdateOne) ClearDeletedAt() *WorktaskUpdateOne {
	wuo.mutation.ClearDeletedAt()
	return wuo
}

// SetDescription sets the "description" field.
func (wuo *WorktaskUpdateOne) SetDescription(s string) *WorktaskUpdateOne {
	wuo.mutation.SetDescription(s)
	return wuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableDescription(s *string) *WorktaskUpdateOne {
	if s != nil {
		wuo.SetDescription(*s)
	}
	return wuo
}

// ClearDescription clears the value of the "description" field.
func (wuo *WorktaskUpdateOne) ClearDescription() *WorktaskUpdateOne {
	wuo.mutation.ClearDescription()
	return wuo
}

// SetStatus sets the "status" field.
func (wuo *WorktaskUpdateOne) SetStatus(w worktask.Status) *WorktaskUpdateOne {
	wuo.mutation.SetStatus(w)
	return wuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableStatus(w *worktask.Status) *WorktaskUpdateOne {
	if w != nil {
		wuo.SetStatus(*w)
	}
	return wuo
}

// SetSubtasks sets the "subtasks" field.
func (wuo *WorktaskUpdateOne) SetSubtasks(s []string) *WorktaskUpdateOne {
	wuo.mutation.SetSubtasks(s)
	return wuo
}

// AppendSubtasks appends s to the "subtasks" field.
func (wuo *WorktaskUpdateOne) AppendSubtasks(s []string) *WorktaskUpdateOne {
	wuo.mutation.AppendSubtasks(s)
	return wuo
}

// ClearSubtasks clears the value of the "subtasks" field.
func (wuo *WorktaskUpdateOne) ClearSubtasks() *WorktaskUpdateOne {
	wuo.mutation.ClearSubtasks()
	return wuo
}

// SetTitle sets the "title" field.
func (wuo *WorktaskUpdateOne) SetTitle(s string) *WorktaskUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableTitle(s *string) *WorktaskUpdateOne {
	if s != nil {
		wuo.SetTitle(*s)
	}
	return wuo
}

// SetStartTime sets the "startTime" field.
func (wuo *WorktaskUpdateOne) SetStartTime(t time.Time) *WorktaskUpdateOne {
	wuo.mutation.SetStartTime(t)
	return wuo
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableStartTime(t *time.Time) *WorktaskUpdateOne {
	if t != nil {
		wuo.SetStartTime(*t)
	}
	return wuo
}

// SetEndTime sets the "endTime" field.
func (wuo *WorktaskUpdateOne) SetEndTime(t time.Time) *WorktaskUpdateOne {
	wuo.mutation.SetEndTime(t)
	return wuo
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableEndTime(t *time.Time) *WorktaskUpdateOne {
	if t != nil {
		wuo.SetEndTime(*t)
	}
	return wuo
}

// ClearEndTime clears the value of the "endTime" field.
func (wuo *WorktaskUpdateOne) ClearEndTime() *WorktaskUpdateOne {
	wuo.mutation.ClearEndTime()
	return wuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (wuo *WorktaskUpdateOne) SetCompanyID(id int) *WorktaskUpdateOne {
	wuo.mutation.SetCompanyID(id)
	return wuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableCompanyID(id *int) *WorktaskUpdateOne {
	if id != nil {
		wuo = wuo.SetCompanyID(*id)
	}
	return wuo
}

// SetCompany sets the "company" edge to the Company entity.
func (wuo *WorktaskUpdateOne) SetCompany(c *Company) *WorktaskUpdateOne {
	return wuo.SetCompanyID(c.ID)
}

// SetCreatedByID sets the "createdBy" edge to the User entity by ID.
func (wuo *WorktaskUpdateOne) SetCreatedByID(id int) *WorktaskUpdateOne {
	wuo.mutation.SetCreatedByID(id)
	return wuo
}

// SetNillableCreatedByID sets the "createdBy" edge to the User entity by ID if the given value is not nil.
func (wuo *WorktaskUpdateOne) SetNillableCreatedByID(id *int) *WorktaskUpdateOne {
	if id != nil {
		wuo = wuo.SetCreatedByID(*id)
	}
	return wuo
}

// SetCreatedBy sets the "createdBy" edge to the User entity.
func (wuo *WorktaskUpdateOne) SetCreatedBy(u *User) *WorktaskUpdateOne {
	return wuo.SetCreatedByID(u.ID)
}

// AddAssignedToIDs adds the "assignedTo" edge to the Employee entity by IDs.
func (wuo *WorktaskUpdateOne) AddAssignedToIDs(ids ...int) *WorktaskUpdateOne {
	wuo.mutation.AddAssignedToIDs(ids...)
	return wuo
}

// AddAssignedTo adds the "assignedTo" edges to the Employee entity.
func (wuo *WorktaskUpdateOne) AddAssignedTo(e ...*Employee) *WorktaskUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wuo.AddAssignedToIDs(ids...)
}

// AddWorkShiftIDs adds the "workShifts" edge to the Workshift entity by IDs.
func (wuo *WorktaskUpdateOne) AddWorkShiftIDs(ids ...int) *WorktaskUpdateOne {
	wuo.mutation.AddWorkShiftIDs(ids...)
	return wuo
}

// AddWorkShifts adds the "workShifts" edges to the Workshift entity.
func (wuo *WorktaskUpdateOne) AddWorkShifts(w ...*Workshift) *WorktaskUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWorkShiftIDs(ids...)
}

// AddWorkTagIDs adds the "workTags" edge to the Worktag entity by IDs.
func (wuo *WorktaskUpdateOne) AddWorkTagIDs(ids ...int) *WorktaskUpdateOne {
	wuo.mutation.AddWorkTagIDs(ids...)
	return wuo
}

// AddWorkTags adds the "workTags" edges to the Worktag entity.
func (wuo *WorktaskUpdateOne) AddWorkTags(w ...*Worktag) *WorktaskUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWorkTagIDs(ids...)
}

// Mutation returns the WorktaskMutation object of the builder.
func (wuo *WorktaskUpdateOne) Mutation() *WorktaskMutation {
	return wuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (wuo *WorktaskUpdateOne) ClearCompany() *WorktaskUpdateOne {
	wuo.mutation.ClearCompany()
	return wuo
}

// ClearCreatedBy clears the "createdBy" edge to the User entity.
func (wuo *WorktaskUpdateOne) ClearCreatedBy() *WorktaskUpdateOne {
	wuo.mutation.ClearCreatedBy()
	return wuo
}

// ClearAssignedTo clears all "assignedTo" edges to the Employee entity.
func (wuo *WorktaskUpdateOne) ClearAssignedTo() *WorktaskUpdateOne {
	wuo.mutation.ClearAssignedTo()
	return wuo
}

// RemoveAssignedToIDs removes the "assignedTo" edge to Employee entities by IDs.
func (wuo *WorktaskUpdateOne) RemoveAssignedToIDs(ids ...int) *WorktaskUpdateOne {
	wuo.mutation.RemoveAssignedToIDs(ids...)
	return wuo
}

// RemoveAssignedTo removes "assignedTo" edges to Employee entities.
func (wuo *WorktaskUpdateOne) RemoveAssignedTo(e ...*Employee) *WorktaskUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wuo.RemoveAssignedToIDs(ids...)
}

// ClearWorkShifts clears all "workShifts" edges to the Workshift entity.
func (wuo *WorktaskUpdateOne) ClearWorkShifts() *WorktaskUpdateOne {
	wuo.mutation.ClearWorkShifts()
	return wuo
}

// RemoveWorkShiftIDs removes the "workShifts" edge to Workshift entities by IDs.
func (wuo *WorktaskUpdateOne) RemoveWorkShiftIDs(ids ...int) *WorktaskUpdateOne {
	wuo.mutation.RemoveWorkShiftIDs(ids...)
	return wuo
}

// RemoveWorkShifts removes "workShifts" edges to Workshift entities.
func (wuo *WorktaskUpdateOne) RemoveWorkShifts(w ...*Workshift) *WorktaskUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWorkShiftIDs(ids...)
}

// ClearWorkTags clears all "workTags" edges to the Worktag entity.
func (wuo *WorktaskUpdateOne) ClearWorkTags() *WorktaskUpdateOne {
	wuo.mutation.ClearWorkTags()
	return wuo
}

// RemoveWorkTagIDs removes the "workTags" edge to Worktag entities by IDs.
func (wuo *WorktaskUpdateOne) RemoveWorkTagIDs(ids ...int) *WorktaskUpdateOne {
	wuo.mutation.RemoveWorkTagIDs(ids...)
	return wuo
}

// RemoveWorkTags removes "workTags" edges to Worktag entities.
func (wuo *WorktaskUpdateOne) RemoveWorkTags(w ...*Worktag) *WorktaskUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWorkTagIDs(ids...)
}

// Where appends a list predicates to the WorktaskUpdate builder.
func (wuo *WorktaskUpdateOne) Where(ps ...predicate.Worktask) *WorktaskUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorktaskUpdateOne) Select(field string, fields ...string) *WorktaskUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Worktask entity.
func (wuo *WorktaskUpdateOne) Save(ctx context.Context) (*Worktask, error) {
	wuo.defaults()
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorktaskUpdateOne) SaveX(ctx context.Context) *Worktask {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorktaskUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorktaskUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WorktaskUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := worktask.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorktaskUpdateOne) check() error {
	if v, ok := wuo.mutation.Status(); ok {
		if err := worktask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Worktask.status": %w`, err)}
		}
	}
	return nil
}

func (wuo *WorktaskUpdateOne) sqlSave(ctx context.Context) (_node *Worktask, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(worktask.Table, worktask.Columns, sqlgraph.NewFieldSpec(worktask.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Worktask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, worktask.FieldID)
		for _, f := range fields {
			if !worktask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != worktask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.SetField(worktask.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wuo.mutation.DeletedAt(); ok {
		_spec.SetField(worktask.FieldDeletedAt, field.TypeTime, value)
	}
	if wuo.mutation.DeletedAtCleared() {
		_spec.ClearField(worktask.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := wuo.mutation.Description(); ok {
		_spec.SetField(worktask.FieldDescription, field.TypeString, value)
	}
	if wuo.mutation.DescriptionCleared() {
		_spec.ClearField(worktask.FieldDescription, field.TypeString)
	}
	if value, ok := wuo.mutation.Status(); ok {
		_spec.SetField(worktask.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := wuo.mutation.Subtasks(); ok {
		_spec.SetField(worktask.FieldSubtasks, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedSubtasks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, worktask.FieldSubtasks, value)
		})
	}
	if wuo.mutation.SubtasksCleared() {
		_spec.ClearField(worktask.FieldSubtasks, field.TypeJSON)
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.SetField(worktask.FieldTitle, field.TypeString, value)
	}
	if value, ok := wuo.mutation.StartTime(); ok {
		_spec.SetField(worktask.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := wuo.mutation.EndTime(); ok {
		_spec.SetField(worktask.FieldEndTime, field.TypeTime, value)
	}
	if wuo.mutation.EndTimeCleared() {
		_spec.ClearField(worktask.FieldEndTime, field.TypeTime)
	}
	if wuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CompanyTable,
			Columns: []string{worktask.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CompanyTable,
			Columns: []string{worktask.CompanyColumn},
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
	if wuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CreatedByTable,
			Columns: []string{worktask.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   worktask.CreatedByTable,
			Columns: []string{worktask.CreatedByColumn},
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
	if wuo.mutation.AssignedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   worktask.AssignedToTable,
			Columns: worktask.AssignedToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedAssignedToIDs(); len(nodes) > 0 && !wuo.mutation.AssignedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   worktask.AssignedToTable,
			Columns: worktask.AssignedToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.AssignedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   worktask.AssignedToTable,
			Columns: worktask.AssignedToPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.WorkShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   worktask.WorkShiftsTable,
			Columns: []string{worktask.WorkShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWorkShiftsIDs(); len(nodes) > 0 && !wuo.mutation.WorkShiftsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   worktask.WorkShiftsTable,
			Columns: []string{worktask.WorkShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WorkShiftsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   worktask.WorkShiftsTable,
			Columns: []string{worktask.WorkShiftsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workshift.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.WorkTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   worktask.WorkTagsTable,
			Columns: worktask.WorkTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWorkTagsIDs(); len(nodes) > 0 && !wuo.mutation.WorkTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   worktask.WorkTagsTable,
			Columns: worktask.WorkTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WorkTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   worktask.WorkTagsTable,
			Columns: worktask.WorkTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worktag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Worktask{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{worktask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
