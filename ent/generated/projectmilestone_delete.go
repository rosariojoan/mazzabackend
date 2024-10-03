// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"mazza/ent/generated/predicate"
	"mazza/ent/generated/projectmilestone"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectMilestoneDelete is the builder for deleting a ProjectMilestone entity.
type ProjectMilestoneDelete struct {
	config
	hooks    []Hook
	mutation *ProjectMilestoneMutation
}

// Where appends a list predicates to the ProjectMilestoneDelete builder.
func (pmd *ProjectMilestoneDelete) Where(ps ...predicate.ProjectMilestone) *ProjectMilestoneDelete {
	pmd.mutation.Where(ps...)
	return pmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pmd *ProjectMilestoneDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pmd.sqlExec, pmd.mutation, pmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pmd *ProjectMilestoneDelete) ExecX(ctx context.Context) int {
	n, err := pmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pmd *ProjectMilestoneDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(projectmilestone.Table, sqlgraph.NewFieldSpec(projectmilestone.FieldID, field.TypeInt))
	if ps := pmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pmd.mutation.done = true
	return affected, err
}

// ProjectMilestoneDeleteOne is the builder for deleting a single ProjectMilestone entity.
type ProjectMilestoneDeleteOne struct {
	pmd *ProjectMilestoneDelete
}

// Where appends a list predicates to the ProjectMilestoneDelete builder.
func (pmdo *ProjectMilestoneDeleteOne) Where(ps ...predicate.ProjectMilestone) *ProjectMilestoneDeleteOne {
	pmdo.pmd.mutation.Where(ps...)
	return pmdo
}

// Exec executes the deletion query.
func (pmdo *ProjectMilestoneDeleteOne) Exec(ctx context.Context) error {
	n, err := pmdo.pmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projectmilestone.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pmdo *ProjectMilestoneDeleteOne) ExecX(ctx context.Context) {
	if err := pmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
