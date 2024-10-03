// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projectmilestone"
	"mazza/ent/generated/projecttask"
	"mazza/ent/generated/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks     []Hook
	mutation  *ProjectMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updatedAt" field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDeletedAt sets the "deletedAt" field.
func (pu *ProjectUpdate) SetDeletedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDeletedAt(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (pu *ProjectUpdate) ClearDeletedAt() *ProjectUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDescription(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetStartDate sets the "startDate" field.
func (pu *ProjectUpdate) SetStartDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetStartDate(t)
	return pu
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStartDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetStartDate(*t)
	}
	return pu
}

// SetEndDate sets the "endDate" field.
func (pu *ProjectUpdate) SetEndDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetEndDate(t)
	return pu
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableEndDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetEndDate(*t)
	}
	return pu
}

// SetProgress sets the "progress" field.
func (pu *ProjectUpdate) SetProgress(f float64) *ProjectUpdate {
	pu.mutation.ResetProgress()
	pu.mutation.SetProgress(f)
	return pu
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableProgress(f *float64) *ProjectUpdate {
	if f != nil {
		pu.SetProgress(*f)
	}
	return pu
}

// AddProgress adds f to the "progress" field.
func (pu *ProjectUpdate) AddProgress(f float64) *ProjectUpdate {
	pu.mutation.AddProgress(f)
	return pu
}

// SetStatus sets the "status" field.
func (pu *ProjectUpdate) SetStatus(pr project.Status) *ProjectUpdate {
	pu.mutation.SetStatus(pr)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStatus(pr *project.Status) *ProjectUpdate {
	if pr != nil {
		pu.SetStatus(*pr)
	}
	return pu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pu *ProjectUpdate) SetCompanyID(id int) *ProjectUpdate {
	pu.mutation.SetCompanyID(id)
	return pu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCompanyID(id *int) *ProjectUpdate {
	if id != nil {
		pu = pu.SetCompanyID(*id)
	}
	return pu
}

// SetCompany sets the "company" edge to the Company entity.
func (pu *ProjectUpdate) SetCompany(c *Company) *ProjectUpdate {
	return pu.SetCompanyID(c.ID)
}

// SetCreatedByID sets the "createdBy" edge to the User entity by ID.
func (pu *ProjectUpdate) SetCreatedByID(id int) *ProjectUpdate {
	pu.mutation.SetCreatedByID(id)
	return pu
}

// SetNillableCreatedByID sets the "createdBy" edge to the User entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCreatedByID(id *int) *ProjectUpdate {
	if id != nil {
		pu = pu.SetCreatedByID(*id)
	}
	return pu
}

// SetCreatedBy sets the "createdBy" edge to the User entity.
func (pu *ProjectUpdate) SetCreatedBy(u *User) *ProjectUpdate {
	return pu.SetCreatedByID(u.ID)
}

// SetLeaderID sets the "leader" edge to the User entity by ID.
func (pu *ProjectUpdate) SetLeaderID(id int) *ProjectUpdate {
	pu.mutation.SetLeaderID(id)
	return pu
}

// SetNillableLeaderID sets the "leader" edge to the User entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableLeaderID(id *int) *ProjectUpdate {
	if id != nil {
		pu = pu.SetLeaderID(*id)
	}
	return pu
}

// SetLeader sets the "leader" edge to the User entity.
func (pu *ProjectUpdate) SetLeader(u *User) *ProjectUpdate {
	return pu.SetLeaderID(u.ID)
}

// AddTaskIDs adds the "tasks" edge to the ProjectTask entity by IDs.
func (pu *ProjectUpdate) AddTaskIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddTaskIDs(ids...)
	return pu
}

// AddTasks adds the "tasks" edges to the ProjectTask entity.
func (pu *ProjectUpdate) AddTasks(p ...*ProjectTask) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddTaskIDs(ids...)
}

// AddMilestoneIDs adds the "milestones" edge to the ProjectMilestone entity by IDs.
func (pu *ProjectUpdate) AddMilestoneIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddMilestoneIDs(ids...)
	return pu
}

// AddMilestones adds the "milestones" edges to the ProjectMilestone entity.
func (pu *ProjectUpdate) AddMilestones(p ...*ProjectMilestone) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddMilestoneIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (pu *ProjectUpdate) ClearCompany() *ProjectUpdate {
	pu.mutation.ClearCompany()
	return pu
}

// ClearCreatedBy clears the "createdBy" edge to the User entity.
func (pu *ProjectUpdate) ClearCreatedBy() *ProjectUpdate {
	pu.mutation.ClearCreatedBy()
	return pu
}

// ClearLeader clears the "leader" edge to the User entity.
func (pu *ProjectUpdate) ClearLeader() *ProjectUpdate {
	pu.mutation.ClearLeader()
	return pu
}

// ClearTasks clears all "tasks" edges to the ProjectTask entity.
func (pu *ProjectUpdate) ClearTasks() *ProjectUpdate {
	pu.mutation.ClearTasks()
	return pu
}

// RemoveTaskIDs removes the "tasks" edge to ProjectTask entities by IDs.
func (pu *ProjectUpdate) RemoveTaskIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveTaskIDs(ids...)
	return pu
}

// RemoveTasks removes "tasks" edges to ProjectTask entities.
func (pu *ProjectUpdate) RemoveTasks(p ...*ProjectTask) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveTaskIDs(ids...)
}

// ClearMilestones clears all "milestones" edges to the ProjectMilestone entity.
func (pu *ProjectUpdate) ClearMilestones() *ProjectUpdate {
	pu.mutation.ClearMilestones()
	return pu
}

// RemoveMilestoneIDs removes the "milestones" edge to ProjectMilestone entities by IDs.
func (pu *ProjectUpdate) RemoveMilestoneIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveMilestoneIDs(ids...)
	return pu
}

// RemoveMilestones removes "milestones" edges to ProjectMilestone entities.
func (pu *ProjectUpdate) RemoveMilestones(p ...*ProjectMilestone) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveMilestoneIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := project.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`generated: validator failed for field "Project.description": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Progress(); ok {
		if err := project.ProgressValidator(v); err != nil {
			return &ValidationError{Name: "progress", err: fmt.Errorf(`generated: validator failed for field "Project.progress": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Project.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *ProjectUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProjectUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(project.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(project.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Progress(); ok {
		_spec.SetField(project.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedProgress(); ok {
		_spec.AddField(project.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CompanyTable,
			Columns: []string{project.CompanyColumn},
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
			Table:   project.CompanyTable,
			Columns: []string{project.CompanyColumn},
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
	if pu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CreatedByTable,
			Columns: []string{project.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CreatedByTable,
			Columns: []string{project.CreatedByColumn},
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
	if pu.mutation.LeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.LeaderTable,
			Columns: []string{project.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.LeaderTable,
			Columns: []string{project.LeaderColumn},
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
	if pu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MilestonesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.MilestonesTable,
			Columns: []string{project.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedMilestonesIDs(); len(nodes) > 0 && !pu.mutation.MilestonesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.MilestonesTable,
			Columns: []string{project.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MilestonesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.MilestonesTable,
			Columns: []string{project.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt),
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
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProjectMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updatedAt" field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDeletedAt sets the "deletedAt" field.
func (puo *ProjectUpdateOne) SetDeletedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDeletedAt(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (puo *ProjectUpdateOne) ClearDeletedAt() *ProjectUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDescription(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetStartDate sets the "startDate" field.
func (puo *ProjectUpdateOne) SetStartDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetStartDate(t)
	return puo
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStartDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetStartDate(*t)
	}
	return puo
}

// SetEndDate sets the "endDate" field.
func (puo *ProjectUpdateOne) SetEndDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetEndDate(t)
	return puo
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableEndDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetEndDate(*t)
	}
	return puo
}

// SetProgress sets the "progress" field.
func (puo *ProjectUpdateOne) SetProgress(f float64) *ProjectUpdateOne {
	puo.mutation.ResetProgress()
	puo.mutation.SetProgress(f)
	return puo
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableProgress(f *float64) *ProjectUpdateOne {
	if f != nil {
		puo.SetProgress(*f)
	}
	return puo
}

// AddProgress adds f to the "progress" field.
func (puo *ProjectUpdateOne) AddProgress(f float64) *ProjectUpdateOne {
	puo.mutation.AddProgress(f)
	return puo
}

// SetStatus sets the "status" field.
func (puo *ProjectUpdateOne) SetStatus(pr project.Status) *ProjectUpdateOne {
	puo.mutation.SetStatus(pr)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStatus(pr *project.Status) *ProjectUpdateOne {
	if pr != nil {
		puo.SetStatus(*pr)
	}
	return puo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (puo *ProjectUpdateOne) SetCompanyID(id int) *ProjectUpdateOne {
	puo.mutation.SetCompanyID(id)
	return puo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCompanyID(id *int) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetCompanyID(*id)
	}
	return puo
}

// SetCompany sets the "company" edge to the Company entity.
func (puo *ProjectUpdateOne) SetCompany(c *Company) *ProjectUpdateOne {
	return puo.SetCompanyID(c.ID)
}

// SetCreatedByID sets the "createdBy" edge to the User entity by ID.
func (puo *ProjectUpdateOne) SetCreatedByID(id int) *ProjectUpdateOne {
	puo.mutation.SetCreatedByID(id)
	return puo
}

// SetNillableCreatedByID sets the "createdBy" edge to the User entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCreatedByID(id *int) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetCreatedByID(*id)
	}
	return puo
}

// SetCreatedBy sets the "createdBy" edge to the User entity.
func (puo *ProjectUpdateOne) SetCreatedBy(u *User) *ProjectUpdateOne {
	return puo.SetCreatedByID(u.ID)
}

// SetLeaderID sets the "leader" edge to the User entity by ID.
func (puo *ProjectUpdateOne) SetLeaderID(id int) *ProjectUpdateOne {
	puo.mutation.SetLeaderID(id)
	return puo
}

// SetNillableLeaderID sets the "leader" edge to the User entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableLeaderID(id *int) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetLeaderID(*id)
	}
	return puo
}

// SetLeader sets the "leader" edge to the User entity.
func (puo *ProjectUpdateOne) SetLeader(u *User) *ProjectUpdateOne {
	return puo.SetLeaderID(u.ID)
}

// AddTaskIDs adds the "tasks" edge to the ProjectTask entity by IDs.
func (puo *ProjectUpdateOne) AddTaskIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddTaskIDs(ids...)
	return puo
}

// AddTasks adds the "tasks" edges to the ProjectTask entity.
func (puo *ProjectUpdateOne) AddTasks(p ...*ProjectTask) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddTaskIDs(ids...)
}

// AddMilestoneIDs adds the "milestones" edge to the ProjectMilestone entity by IDs.
func (puo *ProjectUpdateOne) AddMilestoneIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddMilestoneIDs(ids...)
	return puo
}

// AddMilestones adds the "milestones" edges to the ProjectMilestone entity.
func (puo *ProjectUpdateOne) AddMilestones(p ...*ProjectMilestone) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddMilestoneIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (puo *ProjectUpdateOne) ClearCompany() *ProjectUpdateOne {
	puo.mutation.ClearCompany()
	return puo
}

// ClearCreatedBy clears the "createdBy" edge to the User entity.
func (puo *ProjectUpdateOne) ClearCreatedBy() *ProjectUpdateOne {
	puo.mutation.ClearCreatedBy()
	return puo
}

// ClearLeader clears the "leader" edge to the User entity.
func (puo *ProjectUpdateOne) ClearLeader() *ProjectUpdateOne {
	puo.mutation.ClearLeader()
	return puo
}

// ClearTasks clears all "tasks" edges to the ProjectTask entity.
func (puo *ProjectUpdateOne) ClearTasks() *ProjectUpdateOne {
	puo.mutation.ClearTasks()
	return puo
}

// RemoveTaskIDs removes the "tasks" edge to ProjectTask entities by IDs.
func (puo *ProjectUpdateOne) RemoveTaskIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveTaskIDs(ids...)
	return puo
}

// RemoveTasks removes "tasks" edges to ProjectTask entities.
func (puo *ProjectUpdateOne) RemoveTasks(p ...*ProjectTask) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveTaskIDs(ids...)
}

// ClearMilestones clears all "milestones" edges to the ProjectMilestone entity.
func (puo *ProjectUpdateOne) ClearMilestones() *ProjectUpdateOne {
	puo.mutation.ClearMilestones()
	return puo
}

// RemoveMilestoneIDs removes the "milestones" edge to ProjectMilestone entities by IDs.
func (puo *ProjectUpdateOne) RemoveMilestoneIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveMilestoneIDs(ids...)
	return puo
}

// RemoveMilestones removes "milestones" edges to ProjectMilestone entities.
func (puo *ProjectUpdateOne) RemoveMilestones(p ...*ProjectMilestone) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveMilestoneIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := project.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`generated: validator failed for field "Project.description": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Progress(); ok {
		if err := project.ProgressValidator(v); err != nil {
			return &ValidationError{Name: "progress", err: fmt.Errorf(`generated: validator failed for field "Project.progress": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Project.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *ProjectUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProjectUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != project.FieldID {
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
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(project.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(project.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Progress(); ok {
		_spec.SetField(project.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedProgress(); ok {
		_spec.AddField(project.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CompanyTable,
			Columns: []string{project.CompanyColumn},
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
			Table:   project.CompanyTable,
			Columns: []string{project.CompanyColumn},
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
	if puo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CreatedByTable,
			Columns: []string{project.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CreatedByTable,
			Columns: []string{project.CreatedByColumn},
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
	if puo.mutation.LeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.LeaderTable,
			Columns: []string{project.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.LeaderTable,
			Columns: []string{project.LeaderColumn},
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
	if puo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !puo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TasksTable,
			Columns: []string{project.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MilestonesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.MilestonesTable,
			Columns: []string{project.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedMilestonesIDs(); len(nodes) > 0 && !puo.mutation.MilestonesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.MilestonesTable,
			Columns: []string{project.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MilestonesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.MilestonesTable,
			Columns: []string{project.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
