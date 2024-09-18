// Code generated by ent, DO NOT EDIT.

package productmovement

import (
	"mazza/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldDeletedAt, v))
}

// EntryGroup applies equality check predicate on the "entryGroup" field. It's identical to EntryGroupEQ.
func EntryGroup(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldEntryGroup, v))
}

// AverageCost applies equality check predicate on the "averageCost" field. It's identical to AverageCostEQ.
func AverageCost(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldAverageCost, v))
}

// UnitCost applies equality check predicate on the "unitCost" field. It's identical to UnitCostEQ.
func UnitCost(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldUnitCost, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldPrice, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldQuantity, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotNull(FieldDeletedAt))
}

// EntryGroupEQ applies the EQ predicate on the "entryGroup" field.
func EntryGroupEQ(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldEntryGroup, v))
}

// EntryGroupNEQ applies the NEQ predicate on the "entryGroup" field.
func EntryGroupNEQ(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldEntryGroup, v))
}

// EntryGroupIn applies the In predicate on the "entryGroup" field.
func EntryGroupIn(vs ...int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldEntryGroup, vs...))
}

// EntryGroupNotIn applies the NotIn predicate on the "entryGroup" field.
func EntryGroupNotIn(vs ...int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldEntryGroup, vs...))
}

// EntryGroupGT applies the GT predicate on the "entryGroup" field.
func EntryGroupGT(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldEntryGroup, v))
}

// EntryGroupGTE applies the GTE predicate on the "entryGroup" field.
func EntryGroupGTE(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldEntryGroup, v))
}

// EntryGroupLT applies the LT predicate on the "entryGroup" field.
func EntryGroupLT(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldEntryGroup, v))
}

// EntryGroupLTE applies the LTE predicate on the "entryGroup" field.
func EntryGroupLTE(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldEntryGroup, v))
}

// AverageCostEQ applies the EQ predicate on the "averageCost" field.
func AverageCostEQ(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldAverageCost, v))
}

// AverageCostNEQ applies the NEQ predicate on the "averageCost" field.
func AverageCostNEQ(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldAverageCost, v))
}

// AverageCostIn applies the In predicate on the "averageCost" field.
func AverageCostIn(vs ...float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldAverageCost, vs...))
}

// AverageCostNotIn applies the NotIn predicate on the "averageCost" field.
func AverageCostNotIn(vs ...float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldAverageCost, vs...))
}

// AverageCostGT applies the GT predicate on the "averageCost" field.
func AverageCostGT(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldAverageCost, v))
}

// AverageCostGTE applies the GTE predicate on the "averageCost" field.
func AverageCostGTE(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldAverageCost, v))
}

// AverageCostLT applies the LT predicate on the "averageCost" field.
func AverageCostLT(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldAverageCost, v))
}

// AverageCostLTE applies the LTE predicate on the "averageCost" field.
func AverageCostLTE(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldAverageCost, v))
}

// UnitCostEQ applies the EQ predicate on the "unitCost" field.
func UnitCostEQ(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldUnitCost, v))
}

// UnitCostNEQ applies the NEQ predicate on the "unitCost" field.
func UnitCostNEQ(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldUnitCost, v))
}

// UnitCostIn applies the In predicate on the "unitCost" field.
func UnitCostIn(vs ...float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldUnitCost, vs...))
}

// UnitCostNotIn applies the NotIn predicate on the "unitCost" field.
func UnitCostNotIn(vs ...float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldUnitCost, vs...))
}

// UnitCostGT applies the GT predicate on the "unitCost" field.
func UnitCostGT(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldUnitCost, v))
}

// UnitCostGTE applies the GTE predicate on the "unitCost" field.
func UnitCostGTE(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldUnitCost, v))
}

// UnitCostLT applies the LT predicate on the "unitCost" field.
func UnitCostLT(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldUnitCost, v))
}

// UnitCostLTE applies the LTE predicate on the "unitCost" field.
func UnitCostLTE(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldUnitCost, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldPrice, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.ProductMovement {
	return predicate.ProductMovement(sql.FieldLTE(FieldQuantity, v))
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.ProductMovement {
	return predicate.ProductMovement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.ProductMovement {
	return predicate.ProductMovement(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductMovement) predicate.ProductMovement {
	return predicate.ProductMovement(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductMovement) predicate.ProductMovement {
	return predicate.ProductMovement(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductMovement) predicate.ProductMovement {
	return predicate.ProductMovement(sql.NotPredicates(p))
}