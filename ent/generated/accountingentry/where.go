// Code generated by ent, DO NOT EDIT.

package accountingentry

import (
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldDeletedAt, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldNumber, v))
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldGroup, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldDate, v))
}

// Account applies equality check predicate on the "account" field. It's identical to AccountEQ.
func Account(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldAccount, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldAmount, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldDescription, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldCategory, v))
}

// IsDebit applies equality check predicate on the "isDebit" field. It's identical to IsDebitEQ.
func IsDebit(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldIsDebit, v))
}

// IsReversal applies equality check predicate on the "isReversal" field. It's identical to IsReversalEQ.
func IsReversal(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldIsReversal, v))
}

// Reversed applies equality check predicate on the "reversed" field. It's identical to ReversedEQ.
func Reversed(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldReversed, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotNull(FieldDeletedAt))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldNumber, v))
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v int) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldGroup, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldDate, v))
}

// AccountEQ applies the EQ predicate on the "account" field.
func AccountEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldAccount, v))
}

// AccountNEQ applies the NEQ predicate on the "account" field.
func AccountNEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldAccount, v))
}

// AccountIn applies the In predicate on the "account" field.
func AccountIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldAccount, vs...))
}

// AccountNotIn applies the NotIn predicate on the "account" field.
func AccountNotIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldAccount, vs...))
}

// AccountGT applies the GT predicate on the "account" field.
func AccountGT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldAccount, v))
}

// AccountGTE applies the GTE predicate on the "account" field.
func AccountGTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldAccount, v))
}

// AccountLT applies the LT predicate on the "account" field.
func AccountLT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldAccount, v))
}

// AccountLTE applies the LTE predicate on the "account" field.
func AccountLTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldAccount, v))
}

// AccountContains applies the Contains predicate on the "account" field.
func AccountContains(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContains(FieldAccount, v))
}

// AccountHasPrefix applies the HasPrefix predicate on the "account" field.
func AccountHasPrefix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasPrefix(FieldAccount, v))
}

// AccountHasSuffix applies the HasSuffix predicate on the "account" field.
func AccountHasSuffix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasSuffix(FieldAccount, v))
}

// AccountEqualFold applies the EqualFold predicate on the "account" field.
func AccountEqualFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEqualFold(FieldAccount, v))
}

// AccountContainsFold applies the ContainsFold predicate on the "account" field.
func AccountContainsFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContainsFold(FieldAccount, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContainsFold(FieldLabel, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldAmount, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContainsFold(FieldDescription, v))
}

// AccountTypeEQ applies the EQ predicate on the "accountType" field.
func AccountTypeEQ(v AccountType) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldAccountType, v))
}

// AccountTypeNEQ applies the NEQ predicate on the "accountType" field.
func AccountTypeNEQ(v AccountType) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldAccountType, v))
}

// AccountTypeIn applies the In predicate on the "accountType" field.
func AccountTypeIn(vs ...AccountType) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldAccountType, vs...))
}

// AccountTypeNotIn applies the NotIn predicate on the "accountType" field.
func AccountTypeNotIn(vs ...AccountType) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldAccountType, vs...))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldContainsFold(FieldCategory, v))
}

// IsDebitEQ applies the EQ predicate on the "isDebit" field.
func IsDebitEQ(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldIsDebit, v))
}

// IsDebitNEQ applies the NEQ predicate on the "isDebit" field.
func IsDebitNEQ(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldIsDebit, v))
}

// IsReversalEQ applies the EQ predicate on the "isReversal" field.
func IsReversalEQ(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldIsReversal, v))
}

// IsReversalNEQ applies the NEQ predicate on the "isReversal" field.
func IsReversalNEQ(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldIsReversal, v))
}

// ReversedEQ applies the EQ predicate on the "reversed" field.
func ReversedEQ(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldEQ(FieldReversed, v))
}

// ReversedNEQ applies the NEQ predicate on the "reversed" field.
func ReversedNEQ(v bool) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.FieldNEQ(FieldReversed, v))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.AccountingEntry {
	return predicate.AccountingEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.AccountingEntry {
	return predicate.AccountingEntry(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.AccountingEntry {
	return predicate.AccountingEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.AccountingEntry {
	return predicate.AccountingEntry(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoan applies the HasEdge predicate on the "loan" edge.
func HasLoan() predicate.AccountingEntry {
	return predicate.AccountingEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LoanTable, LoanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoanWith applies the HasEdge predicate on the "loan" edge with a given conditions (other predicates).
func HasLoanWith(preds ...predicate.Loan) predicate.AccountingEntry {
	return predicate.AccountingEntry(func(s *sql.Selector) {
		step := newLoanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountingEntry) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountingEntry) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccountingEntry) predicate.AccountingEntry {
	return predicate.AccountingEntry(sql.NotPredicates(p))
}
