// Code generated by ent, DO NOT EDIT.

package file

import (
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.File {
	return predicate.File(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.File {
	return predicate.File(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDeletedAt, v))
}

// Extension applies equality check predicate on the "extension" field. It's identical to ExtensionEQ.
func Extension(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldExtension, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldURI, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldURL, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldDeletedAt))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.File {
	return predicate.File(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCategory, vs...))
}

// ExtensionEQ applies the EQ predicate on the "extension" field.
func ExtensionEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldExtension, v))
}

// ExtensionNEQ applies the NEQ predicate on the "extension" field.
func ExtensionNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldExtension, v))
}

// ExtensionIn applies the In predicate on the "extension" field.
func ExtensionIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldExtension, vs...))
}

// ExtensionNotIn applies the NotIn predicate on the "extension" field.
func ExtensionNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldExtension, vs...))
}

// ExtensionGT applies the GT predicate on the "extension" field.
func ExtensionGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldExtension, v))
}

// ExtensionGTE applies the GTE predicate on the "extension" field.
func ExtensionGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldExtension, v))
}

// ExtensionLT applies the LT predicate on the "extension" field.
func ExtensionLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldExtension, v))
}

// ExtensionLTE applies the LTE predicate on the "extension" field.
func ExtensionLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldExtension, v))
}

// ExtensionContains applies the Contains predicate on the "extension" field.
func ExtensionContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldExtension, v))
}

// ExtensionHasPrefix applies the HasPrefix predicate on the "extension" field.
func ExtensionHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldExtension, v))
}

// ExtensionHasSuffix applies the HasSuffix predicate on the "extension" field.
func ExtensionHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldExtension, v))
}

// ExtensionEqualFold applies the EqualFold predicate on the "extension" field.
func ExtensionEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldExtension, v))
}

// ExtensionContainsFold applies the ContainsFold predicate on the "extension" field.
func ExtensionContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldExtension, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSize, v))
}

// SizeContains applies the Contains predicate on the "size" field.
func SizeContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldSize, v))
}

// SizeHasPrefix applies the HasPrefix predicate on the "size" field.
func SizeHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldSize, v))
}

// SizeHasSuffix applies the HasSuffix predicate on the "size" field.
func SizeHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldSize, v))
}

// SizeEqualFold applies the EqualFold predicate on the "size" field.
func SizeEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldSize, v))
}

// SizeContainsFold applies the ContainsFold predicate on the "size" field.
func SizeContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldSize, v))
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldURI, v))
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldURI, v))
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldURI, vs...))
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldURI, vs...))
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldURI, v))
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldURI, v))
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldURI, v))
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldURI, v))
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldURI, v))
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldURI, v))
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldURI, v))
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldURI, v))
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldURI, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldURL, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldDescription, v))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(sql.NotPredicates(p))
}
