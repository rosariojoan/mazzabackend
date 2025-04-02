// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/inventory"
	"mazza/ent/generated/inventorymovement"
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryUpdate is the builder for updating Inventory entities.
type InventoryUpdate struct {
	config
	hooks     []Hook
	mutation  *InventoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the InventoryUpdate builder.
func (iu *InventoryUpdate) Where(ps ...predicate.Inventory) *InventoryUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updatedAt" field.
func (iu *InventoryUpdate) SetUpdatedAt(t time.Time) *InventoryUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetDeletedAt sets the "deletedAt" field.
func (iu *InventoryUpdate) SetDeletedAt(t time.Time) *InventoryUpdate {
	iu.mutation.SetDeletedAt(t)
	return iu
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableDeletedAt(t *time.Time) *InventoryUpdate {
	if t != nil {
		iu.SetDeletedAt(*t)
	}
	return iu
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (iu *InventoryUpdate) ClearDeletedAt() *InventoryUpdate {
	iu.mutation.ClearDeletedAt()
	return iu
}

// SetName sets the "name" field.
func (iu *InventoryUpdate) SetName(s string) *InventoryUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableName(s *string) *InventoryUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetCategory sets the "category" field.
func (iu *InventoryUpdate) SetCategory(i inventory.Category) *InventoryUpdate {
	iu.mutation.SetCategory(i)
	return iu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableCategory(i *inventory.Category) *InventoryUpdate {
	if i != nil {
		iu.SetCategory(*i)
	}
	return iu
}

// SetQuantity sets the "quantity" field.
func (iu *InventoryUpdate) SetQuantity(f float64) *InventoryUpdate {
	iu.mutation.ResetQuantity()
	iu.mutation.SetQuantity(f)
	return iu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableQuantity(f *float64) *InventoryUpdate {
	if f != nil {
		iu.SetQuantity(*f)
	}
	return iu
}

// AddQuantity adds f to the "quantity" field.
func (iu *InventoryUpdate) AddQuantity(f float64) *InventoryUpdate {
	iu.mutation.AddQuantity(f)
	return iu
}

// SetUnit sets the "unit" field.
func (iu *InventoryUpdate) SetUnit(s string) *InventoryUpdate {
	iu.mutation.SetUnit(s)
	return iu
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableUnit(s *string) *InventoryUpdate {
	if s != nil {
		iu.SetUnit(*s)
	}
	return iu
}

// SetMinimumLevel sets the "minimumLevel" field.
func (iu *InventoryUpdate) SetMinimumLevel(f float64) *InventoryUpdate {
	iu.mutation.ResetMinimumLevel()
	iu.mutation.SetMinimumLevel(f)
	return iu
}

// SetNillableMinimumLevel sets the "minimumLevel" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableMinimumLevel(f *float64) *InventoryUpdate {
	if f != nil {
		iu.SetMinimumLevel(*f)
	}
	return iu
}

// AddMinimumLevel adds f to the "minimumLevel" field.
func (iu *InventoryUpdate) AddMinimumLevel(f float64) *InventoryUpdate {
	iu.mutation.AddMinimumLevel(f)
	return iu
}

// SetCurrentValue sets the "currentValue" field.
func (iu *InventoryUpdate) SetCurrentValue(f float64) *InventoryUpdate {
	iu.mutation.ResetCurrentValue()
	iu.mutation.SetCurrentValue(f)
	return iu
}

// SetNillableCurrentValue sets the "currentValue" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableCurrentValue(f *float64) *InventoryUpdate {
	if f != nil {
		iu.SetCurrentValue(*f)
	}
	return iu
}

// AddCurrentValue adds f to the "currentValue" field.
func (iu *InventoryUpdate) AddCurrentValue(f float64) *InventoryUpdate {
	iu.mutation.AddCurrentValue(f)
	return iu
}

// SetNotes sets the "notes" field.
func (iu *InventoryUpdate) SetNotes(s string) *InventoryUpdate {
	iu.mutation.SetNotes(s)
	return iu
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (iu *InventoryUpdate) SetNillableNotes(s *string) *InventoryUpdate {
	if s != nil {
		iu.SetNotes(*s)
	}
	return iu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (iu *InventoryUpdate) SetCompanyID(id int) *InventoryUpdate {
	iu.mutation.SetCompanyID(id)
	return iu
}

// SetCompany sets the "company" edge to the Company entity.
func (iu *InventoryUpdate) SetCompany(c *Company) *InventoryUpdate {
	return iu.SetCompanyID(c.ID)
}

// AddMovementIDs adds the "movements" edge to the InventoryMovement entity by IDs.
func (iu *InventoryUpdate) AddMovementIDs(ids ...int) *InventoryUpdate {
	iu.mutation.AddMovementIDs(ids...)
	return iu
}

// AddMovements adds the "movements" edges to the InventoryMovement entity.
func (iu *InventoryUpdate) AddMovements(i ...*InventoryMovement) *InventoryUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iu.AddMovementIDs(ids...)
}

// Mutation returns the InventoryMutation object of the builder.
func (iu *InventoryUpdate) Mutation() *InventoryMutation {
	return iu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (iu *InventoryUpdate) ClearCompany() *InventoryUpdate {
	iu.mutation.ClearCompany()
	return iu
}

// ClearMovements clears all "movements" edges to the InventoryMovement entity.
func (iu *InventoryUpdate) ClearMovements() *InventoryUpdate {
	iu.mutation.ClearMovements()
	return iu
}

// RemoveMovementIDs removes the "movements" edge to InventoryMovement entities by IDs.
func (iu *InventoryUpdate) RemoveMovementIDs(ids ...int) *InventoryUpdate {
	iu.mutation.RemoveMovementIDs(ids...)
	return iu
}

// RemoveMovements removes "movements" edges to InventoryMovement entities.
func (iu *InventoryUpdate) RemoveMovements(i ...*InventoryMovement) *InventoryUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iu.RemoveMovementIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InventoryUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InventoryUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InventoryUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InventoryUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InventoryUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := inventory.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InventoryUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := inventory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Inventory.name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Category(); ok {
		if err := inventory.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`generated: validator failed for field "Inventory.category": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Quantity(); ok {
		if err := inventory.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`generated: validator failed for field "Inventory.quantity": %w`, err)}
		}
	}
	if v, ok := iu.mutation.MinimumLevel(); ok {
		if err := inventory.MinimumLevelValidator(v); err != nil {
			return &ValidationError{Name: "minimumLevel", err: fmt.Errorf(`generated: validator failed for field "Inventory.minimumLevel": %w`, err)}
		}
	}
	if v, ok := iu.mutation.CurrentValue(); ok {
		if err := inventory.CurrentValueValidator(v); err != nil {
			return &ValidationError{Name: "currentValue", err: fmt.Errorf(`generated: validator failed for field "Inventory.currentValue": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Notes(); ok {
		if err := inventory.NotesValidator(v); err != nil {
			return &ValidationError{Name: "notes", err: fmt.Errorf(`generated: validator failed for field "Inventory.notes": %w`, err)}
		}
	}
	if iu.mutation.CompanyCleared() && len(iu.mutation.CompanyIDs()) > 0 {
		return errors.New(`generated: clearing a required unique edge "Inventory.company"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iu *InventoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InventoryUpdate {
	iu.modifiers = append(iu.modifiers, modifiers...)
	return iu
}

func (iu *InventoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(inventory.Table, inventory.Columns, sqlgraph.NewFieldSpec(inventory.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(inventory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.DeletedAt(); ok {
		_spec.SetField(inventory.FieldDeletedAt, field.TypeTime, value)
	}
	if iu.mutation.DeletedAtCleared() {
		_spec.ClearField(inventory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(inventory.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Category(); ok {
		_spec.SetField(inventory.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := iu.mutation.Quantity(); ok {
		_spec.SetField(inventory.FieldQuantity, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedQuantity(); ok {
		_spec.AddField(inventory.FieldQuantity, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.Unit(); ok {
		_spec.SetField(inventory.FieldUnit, field.TypeString, value)
	}
	if value, ok := iu.mutation.MinimumLevel(); ok {
		_spec.SetField(inventory.FieldMinimumLevel, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedMinimumLevel(); ok {
		_spec.AddField(inventory.FieldMinimumLevel, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.CurrentValue(); ok {
		_spec.SetField(inventory.FieldCurrentValue, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedCurrentValue(); ok {
		_spec.AddField(inventory.FieldCurrentValue, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.Notes(); ok {
		_spec.SetField(inventory.FieldNotes, field.TypeString, value)
	}
	if iu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventory.CompanyTable,
			Columns: []string{inventory.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventory.CompanyTable,
			Columns: []string{inventory.CompanyColumn},
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
	if iu.mutation.MovementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventory.MovementsTable,
			Columns: []string{inventory.MovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorymovement.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedMovementsIDs(); len(nodes) > 0 && !iu.mutation.MovementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventory.MovementsTable,
			Columns: []string{inventory.MovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorymovement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.MovementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventory.MovementsTable,
			Columns: []string{inventory.MovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorymovement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(iu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inventory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InventoryUpdateOne is the builder for updating a single Inventory entity.
type InventoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *InventoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updatedAt" field.
func (iuo *InventoryUpdateOne) SetUpdatedAt(t time.Time) *InventoryUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetDeletedAt sets the "deletedAt" field.
func (iuo *InventoryUpdateOne) SetDeletedAt(t time.Time) *InventoryUpdateOne {
	iuo.mutation.SetDeletedAt(t)
	return iuo
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableDeletedAt(t *time.Time) *InventoryUpdateOne {
	if t != nil {
		iuo.SetDeletedAt(*t)
	}
	return iuo
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (iuo *InventoryUpdateOne) ClearDeletedAt() *InventoryUpdateOne {
	iuo.mutation.ClearDeletedAt()
	return iuo
}

// SetName sets the "name" field.
func (iuo *InventoryUpdateOne) SetName(s string) *InventoryUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableName(s *string) *InventoryUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetCategory sets the "category" field.
func (iuo *InventoryUpdateOne) SetCategory(i inventory.Category) *InventoryUpdateOne {
	iuo.mutation.SetCategory(i)
	return iuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableCategory(i *inventory.Category) *InventoryUpdateOne {
	if i != nil {
		iuo.SetCategory(*i)
	}
	return iuo
}

// SetQuantity sets the "quantity" field.
func (iuo *InventoryUpdateOne) SetQuantity(f float64) *InventoryUpdateOne {
	iuo.mutation.ResetQuantity()
	iuo.mutation.SetQuantity(f)
	return iuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableQuantity(f *float64) *InventoryUpdateOne {
	if f != nil {
		iuo.SetQuantity(*f)
	}
	return iuo
}

// AddQuantity adds f to the "quantity" field.
func (iuo *InventoryUpdateOne) AddQuantity(f float64) *InventoryUpdateOne {
	iuo.mutation.AddQuantity(f)
	return iuo
}

// SetUnit sets the "unit" field.
func (iuo *InventoryUpdateOne) SetUnit(s string) *InventoryUpdateOne {
	iuo.mutation.SetUnit(s)
	return iuo
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableUnit(s *string) *InventoryUpdateOne {
	if s != nil {
		iuo.SetUnit(*s)
	}
	return iuo
}

// SetMinimumLevel sets the "minimumLevel" field.
func (iuo *InventoryUpdateOne) SetMinimumLevel(f float64) *InventoryUpdateOne {
	iuo.mutation.ResetMinimumLevel()
	iuo.mutation.SetMinimumLevel(f)
	return iuo
}

// SetNillableMinimumLevel sets the "minimumLevel" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableMinimumLevel(f *float64) *InventoryUpdateOne {
	if f != nil {
		iuo.SetMinimumLevel(*f)
	}
	return iuo
}

// AddMinimumLevel adds f to the "minimumLevel" field.
func (iuo *InventoryUpdateOne) AddMinimumLevel(f float64) *InventoryUpdateOne {
	iuo.mutation.AddMinimumLevel(f)
	return iuo
}

// SetCurrentValue sets the "currentValue" field.
func (iuo *InventoryUpdateOne) SetCurrentValue(f float64) *InventoryUpdateOne {
	iuo.mutation.ResetCurrentValue()
	iuo.mutation.SetCurrentValue(f)
	return iuo
}

// SetNillableCurrentValue sets the "currentValue" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableCurrentValue(f *float64) *InventoryUpdateOne {
	if f != nil {
		iuo.SetCurrentValue(*f)
	}
	return iuo
}

// AddCurrentValue adds f to the "currentValue" field.
func (iuo *InventoryUpdateOne) AddCurrentValue(f float64) *InventoryUpdateOne {
	iuo.mutation.AddCurrentValue(f)
	return iuo
}

// SetNotes sets the "notes" field.
func (iuo *InventoryUpdateOne) SetNotes(s string) *InventoryUpdateOne {
	iuo.mutation.SetNotes(s)
	return iuo
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (iuo *InventoryUpdateOne) SetNillableNotes(s *string) *InventoryUpdateOne {
	if s != nil {
		iuo.SetNotes(*s)
	}
	return iuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (iuo *InventoryUpdateOne) SetCompanyID(id int) *InventoryUpdateOne {
	iuo.mutation.SetCompanyID(id)
	return iuo
}

// SetCompany sets the "company" edge to the Company entity.
func (iuo *InventoryUpdateOne) SetCompany(c *Company) *InventoryUpdateOne {
	return iuo.SetCompanyID(c.ID)
}

// AddMovementIDs adds the "movements" edge to the InventoryMovement entity by IDs.
func (iuo *InventoryUpdateOne) AddMovementIDs(ids ...int) *InventoryUpdateOne {
	iuo.mutation.AddMovementIDs(ids...)
	return iuo
}

// AddMovements adds the "movements" edges to the InventoryMovement entity.
func (iuo *InventoryUpdateOne) AddMovements(i ...*InventoryMovement) *InventoryUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iuo.AddMovementIDs(ids...)
}

// Mutation returns the InventoryMutation object of the builder.
func (iuo *InventoryUpdateOne) Mutation() *InventoryMutation {
	return iuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (iuo *InventoryUpdateOne) ClearCompany() *InventoryUpdateOne {
	iuo.mutation.ClearCompany()
	return iuo
}

// ClearMovements clears all "movements" edges to the InventoryMovement entity.
func (iuo *InventoryUpdateOne) ClearMovements() *InventoryUpdateOne {
	iuo.mutation.ClearMovements()
	return iuo
}

// RemoveMovementIDs removes the "movements" edge to InventoryMovement entities by IDs.
func (iuo *InventoryUpdateOne) RemoveMovementIDs(ids ...int) *InventoryUpdateOne {
	iuo.mutation.RemoveMovementIDs(ids...)
	return iuo
}

// RemoveMovements removes "movements" edges to InventoryMovement entities.
func (iuo *InventoryUpdateOne) RemoveMovements(i ...*InventoryMovement) *InventoryUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iuo.RemoveMovementIDs(ids...)
}

// Where appends a list predicates to the InventoryUpdate builder.
func (iuo *InventoryUpdateOne) Where(ps ...predicate.Inventory) *InventoryUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InventoryUpdateOne) Select(field string, fields ...string) *InventoryUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Inventory entity.
func (iuo *InventoryUpdateOne) Save(ctx context.Context) (*Inventory, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InventoryUpdateOne) SaveX(ctx context.Context) *Inventory {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InventoryUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InventoryUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InventoryUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := inventory.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InventoryUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := inventory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Inventory.name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Category(); ok {
		if err := inventory.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`generated: validator failed for field "Inventory.category": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Quantity(); ok {
		if err := inventory.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`generated: validator failed for field "Inventory.quantity": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.MinimumLevel(); ok {
		if err := inventory.MinimumLevelValidator(v); err != nil {
			return &ValidationError{Name: "minimumLevel", err: fmt.Errorf(`generated: validator failed for field "Inventory.minimumLevel": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.CurrentValue(); ok {
		if err := inventory.CurrentValueValidator(v); err != nil {
			return &ValidationError{Name: "currentValue", err: fmt.Errorf(`generated: validator failed for field "Inventory.currentValue": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Notes(); ok {
		if err := inventory.NotesValidator(v); err != nil {
			return &ValidationError{Name: "notes", err: fmt.Errorf(`generated: validator failed for field "Inventory.notes": %w`, err)}
		}
	}
	if iuo.mutation.CompanyCleared() && len(iuo.mutation.CompanyIDs()) > 0 {
		return errors.New(`generated: clearing a required unique edge "Inventory.company"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iuo *InventoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InventoryUpdateOne {
	iuo.modifiers = append(iuo.modifiers, modifiers...)
	return iuo
}

func (iuo *InventoryUpdateOne) sqlSave(ctx context.Context) (_node *Inventory, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(inventory.Table, inventory.Columns, sqlgraph.NewFieldSpec(inventory.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Inventory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inventory.FieldID)
		for _, f := range fields {
			if !inventory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != inventory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(inventory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.DeletedAt(); ok {
		_spec.SetField(inventory.FieldDeletedAt, field.TypeTime, value)
	}
	if iuo.mutation.DeletedAtCleared() {
		_spec.ClearField(inventory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(inventory.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Category(); ok {
		_spec.SetField(inventory.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := iuo.mutation.Quantity(); ok {
		_spec.SetField(inventory.FieldQuantity, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedQuantity(); ok {
		_spec.AddField(inventory.FieldQuantity, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.Unit(); ok {
		_spec.SetField(inventory.FieldUnit, field.TypeString, value)
	}
	if value, ok := iuo.mutation.MinimumLevel(); ok {
		_spec.SetField(inventory.FieldMinimumLevel, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedMinimumLevel(); ok {
		_spec.AddField(inventory.FieldMinimumLevel, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.CurrentValue(); ok {
		_spec.SetField(inventory.FieldCurrentValue, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedCurrentValue(); ok {
		_spec.AddField(inventory.FieldCurrentValue, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.Notes(); ok {
		_spec.SetField(inventory.FieldNotes, field.TypeString, value)
	}
	if iuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventory.CompanyTable,
			Columns: []string{inventory.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventory.CompanyTable,
			Columns: []string{inventory.CompanyColumn},
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
	if iuo.mutation.MovementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventory.MovementsTable,
			Columns: []string{inventory.MovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorymovement.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedMovementsIDs(); len(nodes) > 0 && !iuo.mutation.MovementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventory.MovementsTable,
			Columns: []string{inventory.MovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorymovement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.MovementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventory.MovementsTable,
			Columns: []string{inventory.MovementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorymovement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(iuo.modifiers...)
	_node = &Inventory{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inventory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
