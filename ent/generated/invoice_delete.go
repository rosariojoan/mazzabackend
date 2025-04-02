// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvoiceDelete is the builder for deleting a Invoice entity.
type InvoiceDelete struct {
	config
	hooks    []Hook
	mutation *InvoiceMutation
}

// Where appends a list predicates to the InvoiceDelete builder.
func (id *InvoiceDelete) Where(ps ...predicate.Invoice) *InvoiceDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InvoiceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InvoiceDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InvoiceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(invoice.Table, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// InvoiceDeleteOne is the builder for deleting a single Invoice entity.
type InvoiceDeleteOne struct {
	id *InvoiceDelete
}

// Where appends a list predicates to the InvoiceDelete builder.
func (ido *InvoiceDeleteOne) Where(ps ...predicate.Invoice) *InvoiceDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *InvoiceDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invoice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InvoiceDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
