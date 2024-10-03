// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projecttask"
	"mazza/ent/generated/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskCreate is the builder for creating a ProjectTask entity.
type ProjectTaskCreate struct {
	config
	mutation *ProjectTaskMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ptc *ProjectTaskCreate) SetName(s string) *ProjectTaskCreate {
	ptc.mutation.SetName(s)
	return ptc
}

// SetAssigneeName sets the "assigneeName" field.
func (ptc *ProjectTaskCreate) SetAssigneeName(s string) *ProjectTaskCreate {
	ptc.mutation.SetAssigneeName(s)
	return ptc
}

// SetLocation sets the "location" field.
func (ptc *ProjectTaskCreate) SetLocation(s string) *ProjectTaskCreate {
	ptc.mutation.SetLocation(s)
	return ptc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (ptc *ProjectTaskCreate) SetNillableLocation(s *string) *ProjectTaskCreate {
	if s != nil {
		ptc.SetLocation(*s)
	}
	return ptc
}

// SetDueDate sets the "dueDate" field.
func (ptc *ProjectTaskCreate) SetDueDate(t time.Time) *ProjectTaskCreate {
	ptc.mutation.SetDueDate(t)
	return ptc
}

// SetStartDate sets the "startDate" field.
func (ptc *ProjectTaskCreate) SetStartDate(t time.Time) *ProjectTaskCreate {
	ptc.mutation.SetStartDate(t)
	return ptc
}

// SetEndDate sets the "endDate" field.
func (ptc *ProjectTaskCreate) SetEndDate(t time.Time) *ProjectTaskCreate {
	ptc.mutation.SetEndDate(t)
	return ptc
}

// SetDescription sets the "description" field.
func (ptc *ProjectTaskCreate) SetDescription(s string) *ProjectTaskCreate {
	ptc.mutation.SetDescription(s)
	return ptc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ptc *ProjectTaskCreate) SetNillableDescription(s *string) *ProjectTaskCreate {
	if s != nil {
		ptc.SetDescription(*s)
	}
	return ptc
}

// SetStatus sets the "status" field.
func (ptc *ProjectTaskCreate) SetStatus(pr projecttask.Status) *ProjectTaskCreate {
	ptc.mutation.SetStatus(pr)
	return ptc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ptc *ProjectTaskCreate) SetProjectID(id int) *ProjectTaskCreate {
	ptc.mutation.SetProjectID(id)
	return ptc
}

// SetProject sets the "project" edge to the Project entity.
func (ptc *ProjectTaskCreate) SetProject(p *Project) *ProjectTaskCreate {
	return ptc.SetProjectID(p.ID)
}

// SetAssigneeID sets the "assignee" edge to the User entity by ID.
func (ptc *ProjectTaskCreate) SetAssigneeID(id int) *ProjectTaskCreate {
	ptc.mutation.SetAssigneeID(id)
	return ptc
}

// SetNillableAssigneeID sets the "assignee" edge to the User entity by ID if the given value is not nil.
func (ptc *ProjectTaskCreate) SetNillableAssigneeID(id *int) *ProjectTaskCreate {
	if id != nil {
		ptc = ptc.SetAssigneeID(*id)
	}
	return ptc
}

// SetAssignee sets the "assignee" edge to the User entity.
func (ptc *ProjectTaskCreate) SetAssignee(u *User) *ProjectTaskCreate {
	return ptc.SetAssigneeID(u.ID)
}

// AddParticipantIDs adds the "participants" edge to the User entity by IDs.
func (ptc *ProjectTaskCreate) AddParticipantIDs(ids ...int) *ProjectTaskCreate {
	ptc.mutation.AddParticipantIDs(ids...)
	return ptc
}

// AddParticipants adds the "participants" edges to the User entity.
func (ptc *ProjectTaskCreate) AddParticipants(u ...*User) *ProjectTaskCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ptc.AddParticipantIDs(ids...)
}

// Mutation returns the ProjectTaskMutation object of the builder.
func (ptc *ProjectTaskCreate) Mutation() *ProjectTaskMutation {
	return ptc.mutation
}

// Save creates the ProjectTask in the database.
func (ptc *ProjectTaskCreate) Save(ctx context.Context) (*ProjectTask, error) {
	return withHooks(ctx, ptc.sqlSave, ptc.mutation, ptc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *ProjectTaskCreate) SaveX(ctx context.Context) *ProjectTask {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *ProjectTaskCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *ProjectTaskCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *ProjectTaskCreate) check() error {
	if _, ok := ptc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "ProjectTask.name"`)}
	}
	if v, ok := ptc.mutation.Name(); ok {
		if err := projecttask.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "ProjectTask.name": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.AssigneeName(); !ok {
		return &ValidationError{Name: "assigneeName", err: errors.New(`generated: missing required field "ProjectTask.assigneeName"`)}
	}
	if v, ok := ptc.mutation.AssigneeName(); ok {
		if err := projecttask.AssigneeNameValidator(v); err != nil {
			return &ValidationError{Name: "assigneeName", err: fmt.Errorf(`generated: validator failed for field "ProjectTask.assigneeName": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.DueDate(); !ok {
		return &ValidationError{Name: "dueDate", err: errors.New(`generated: missing required field "ProjectTask.dueDate"`)}
	}
	if _, ok := ptc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "startDate", err: errors.New(`generated: missing required field "ProjectTask.startDate"`)}
	}
	if _, ok := ptc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "endDate", err: errors.New(`generated: missing required field "ProjectTask.endDate"`)}
	}
	if _, ok := ptc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "ProjectTask.status"`)}
	}
	if v, ok := ptc.mutation.Status(); ok {
		if err := projecttask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "ProjectTask.status": %w`, err)}
		}
	}
	if len(ptc.mutation.ProjectIDs()) == 0 {
		return &ValidationError{Name: "project", err: errors.New(`generated: missing required edge "ProjectTask.project"`)}
	}
	return nil
}

func (ptc *ProjectTaskCreate) sqlSave(ctx context.Context) (*ProjectTask, error) {
	if err := ptc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ptc.mutation.id = &_node.ID
	ptc.mutation.done = true
	return _node, nil
}

func (ptc *ProjectTaskCreate) createSpec() (*ProjectTask, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectTask{config: ptc.config}
		_spec = sqlgraph.NewCreateSpec(projecttask.Table, sqlgraph.NewFieldSpec(projecttask.FieldID, field.TypeInt))
	)
	if value, ok := ptc.mutation.Name(); ok {
		_spec.SetField(projecttask.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ptc.mutation.AssigneeName(); ok {
		_spec.SetField(projecttask.FieldAssigneeName, field.TypeString, value)
		_node.AssigneeName = value
	}
	if value, ok := ptc.mutation.Location(); ok {
		_spec.SetField(projecttask.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := ptc.mutation.DueDate(); ok {
		_spec.SetField(projecttask.FieldDueDate, field.TypeTime, value)
		_node.DueDate = value
	}
	if value, ok := ptc.mutation.StartDate(); ok {
		_spec.SetField(projecttask.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := ptc.mutation.EndDate(); ok {
		_spec.SetField(projecttask.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := ptc.mutation.Description(); ok {
		_spec.SetField(projecttask.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ptc.mutation.Status(); ok {
		_spec.SetField(projecttask.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := ptc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttask.ProjectTable,
			Columns: []string{projecttask.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.project_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ptc.mutation.AssigneeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttask.AssigneeTable,
			Columns: []string{projecttask.AssigneeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_assigned_project_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ptc.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   projecttask.ParticipantsTable,
			Columns: projecttask.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectTaskCreateBulk is the builder for creating many ProjectTask entities in bulk.
type ProjectTaskCreateBulk struct {
	config
	err      error
	builders []*ProjectTaskCreate
}

// Save creates the ProjectTask entities in the database.
func (ptcb *ProjectTaskCreateBulk) Save(ctx context.Context) ([]*ProjectTask, error) {
	if ptcb.err != nil {
		return nil, ptcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*ProjectTask, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *ProjectTaskCreateBulk) SaveX(ctx context.Context) []*ProjectTask {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *ProjectTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *ProjectTaskCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
