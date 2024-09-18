// Code generated by ent, DO NOT EDIT.

package payable

import (
	"mazza/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldDeletedAt, v))
}

// EntryGroup applies equality check predicate on the "entryGroup" field. It's identical to EntryGroupEQ.
func EntryGroup(v int) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldEntryGroup, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldDate, v))
}

// OutstandingBalance applies equality check predicate on the "outstandingBalance" field. It's identical to OutstandingBalanceEQ.
func OutstandingBalance(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldOutstandingBalance, v))
}

// TotalTransaction applies equality check predicate on the "totalTransaction" field. It's identical to TotalTransactionEQ.
func TotalTransaction(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldTotalTransaction, v))
}

// DaysDue applies equality check predicate on the "daysDue" field. It's identical to DaysDueEQ.
func DaysDue(v int) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldDaysDue, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.Payable {
	return predicate.Payable(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.Payable {
	return predicate.Payable(sql.FieldNotNull(FieldDeletedAt))
}

// EntryGroupEQ applies the EQ predicate on the "entryGroup" field.
func EntryGroupEQ(v int) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldEntryGroup, v))
}

// EntryGroupNEQ applies the NEQ predicate on the "entryGroup" field.
func EntryGroupNEQ(v int) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldEntryGroup, v))
}

// EntryGroupIn applies the In predicate on the "entryGroup" field.
func EntryGroupIn(vs ...int) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldEntryGroup, vs...))
}

// EntryGroupNotIn applies the NotIn predicate on the "entryGroup" field.
func EntryGroupNotIn(vs ...int) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldEntryGroup, vs...))
}

// EntryGroupGT applies the GT predicate on the "entryGroup" field.
func EntryGroupGT(v int) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldEntryGroup, v))
}

// EntryGroupGTE applies the GTE predicate on the "entryGroup" field.
func EntryGroupGTE(v int) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldEntryGroup, v))
}

// EntryGroupLT applies the LT predicate on the "entryGroup" field.
func EntryGroupLT(v int) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldEntryGroup, v))
}

// EntryGroupLTE applies the LTE predicate on the "entryGroup" field.
func EntryGroupLTE(v int) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldEntryGroup, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldDate, v))
}

// OutstandingBalanceEQ applies the EQ predicate on the "outstandingBalance" field.
func OutstandingBalanceEQ(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldOutstandingBalance, v))
}

// OutstandingBalanceNEQ applies the NEQ predicate on the "outstandingBalance" field.
func OutstandingBalanceNEQ(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldOutstandingBalance, v))
}

// OutstandingBalanceIn applies the In predicate on the "outstandingBalance" field.
func OutstandingBalanceIn(vs ...float64) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldOutstandingBalance, vs...))
}

// OutstandingBalanceNotIn applies the NotIn predicate on the "outstandingBalance" field.
func OutstandingBalanceNotIn(vs ...float64) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldOutstandingBalance, vs...))
}

// OutstandingBalanceGT applies the GT predicate on the "outstandingBalance" field.
func OutstandingBalanceGT(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldOutstandingBalance, v))
}

// OutstandingBalanceGTE applies the GTE predicate on the "outstandingBalance" field.
func OutstandingBalanceGTE(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldOutstandingBalance, v))
}

// OutstandingBalanceLT applies the LT predicate on the "outstandingBalance" field.
func OutstandingBalanceLT(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldOutstandingBalance, v))
}

// OutstandingBalanceLTE applies the LTE predicate on the "outstandingBalance" field.
func OutstandingBalanceLTE(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldOutstandingBalance, v))
}

// TotalTransactionEQ applies the EQ predicate on the "totalTransaction" field.
func TotalTransactionEQ(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldTotalTransaction, v))
}

// TotalTransactionNEQ applies the NEQ predicate on the "totalTransaction" field.
func TotalTransactionNEQ(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldTotalTransaction, v))
}

// TotalTransactionIn applies the In predicate on the "totalTransaction" field.
func TotalTransactionIn(vs ...float64) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldTotalTransaction, vs...))
}

// TotalTransactionNotIn applies the NotIn predicate on the "totalTransaction" field.
func TotalTransactionNotIn(vs ...float64) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldTotalTransaction, vs...))
}

// TotalTransactionGT applies the GT predicate on the "totalTransaction" field.
func TotalTransactionGT(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldTotalTransaction, v))
}

// TotalTransactionGTE applies the GTE predicate on the "totalTransaction" field.
func TotalTransactionGTE(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldTotalTransaction, v))
}

// TotalTransactionLT applies the LT predicate on the "totalTransaction" field.
func TotalTransactionLT(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldTotalTransaction, v))
}

// TotalTransactionLTE applies the LTE predicate on the "totalTransaction" field.
func TotalTransactionLTE(v float64) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldTotalTransaction, v))
}

// DaysDueEQ applies the EQ predicate on the "daysDue" field.
func DaysDueEQ(v int) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldDaysDue, v))
}

// DaysDueNEQ applies the NEQ predicate on the "daysDue" field.
func DaysDueNEQ(v int) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldDaysDue, v))
}

// DaysDueIn applies the In predicate on the "daysDue" field.
func DaysDueIn(vs ...int) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldDaysDue, vs...))
}

// DaysDueNotIn applies the NotIn predicate on the "daysDue" field.
func DaysDueNotIn(vs ...int) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldDaysDue, vs...))
}

// DaysDueGT applies the GT predicate on the "daysDue" field.
func DaysDueGT(v int) predicate.Payable {
	return predicate.Payable(sql.FieldGT(FieldDaysDue, v))
}

// DaysDueGTE applies the GTE predicate on the "daysDue" field.
func DaysDueGTE(v int) predicate.Payable {
	return predicate.Payable(sql.FieldGTE(FieldDaysDue, v))
}

// DaysDueLT applies the LT predicate on the "daysDue" field.
func DaysDueLT(v int) predicate.Payable {
	return predicate.Payable(sql.FieldLT(FieldDaysDue, v))
}

// DaysDueLTE applies the LTE predicate on the "daysDue" field.
func DaysDueLTE(v int) predicate.Payable {
	return predicate.Payable(sql.FieldLTE(FieldDaysDue, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Payable {
	return predicate.Payable(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Payable {
	return predicate.Payable(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Payable {
	return predicate.Payable(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Payable {
	return predicate.Payable(sql.FieldNotIn(FieldStatus, vs...))
}

// HasSupplier applies the HasEdge predicate on the "supplier" edge.
func HasSupplier() predicate.Payable {
	return predicate.Payable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SupplierTable, SupplierColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSupplierWith applies the HasEdge predicate on the "supplier" edge with a given conditions (other predicates).
func HasSupplierWith(preds ...predicate.Supplier) predicate.Payable {
	return predicate.Payable(func(s *sql.Selector) {
		step := newSupplierStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payable) predicate.Payable {
	return predicate.Payable(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payable) predicate.Payable {
	return predicate.Payable(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payable) predicate.Payable {
	return predicate.Payable(sql.NotPredicates(p))
}
