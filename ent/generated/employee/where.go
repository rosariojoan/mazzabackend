// Code generated by ent, DO NOT EDIT.

package employee

import (
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldName, v))
}

// Birthdate applies equality check predicate on the "birthdate" field. It's identical to BirthdateEQ.
func Birthdate(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldBirthdate, v))
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPosition, v))
}

// Department applies equality check predicate on the "department" field. It's identical to DepartmentEQ.
func Department(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldDepartment, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPhone, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldAvatar, v))
}

// HireDate applies equality check predicate on the "hireDate" field. It's identical to HireDateEQ.
func HireDate(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldHireDate, v))
}

// MonthlySalary applies equality check predicate on the "monthlySalary" field. It's identical to MonthlySalaryEQ.
func MonthlySalary(v int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldMonthlySalary, v))
}

// PerformaceScore applies equality check predicate on the "performaceScore" field. It's identical to PerformaceScoreEQ.
func PerformaceScore(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPerformaceScore, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldName, v))
}

// BirthdateEQ applies the EQ predicate on the "birthdate" field.
func BirthdateEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldBirthdate, v))
}

// BirthdateNEQ applies the NEQ predicate on the "birthdate" field.
func BirthdateNEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldBirthdate, v))
}

// BirthdateIn applies the In predicate on the "birthdate" field.
func BirthdateIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldBirthdate, vs...))
}

// BirthdateNotIn applies the NotIn predicate on the "birthdate" field.
func BirthdateNotIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldBirthdate, vs...))
}

// BirthdateGT applies the GT predicate on the "birthdate" field.
func BirthdateGT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldBirthdate, v))
}

// BirthdateGTE applies the GTE predicate on the "birthdate" field.
func BirthdateGTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldBirthdate, v))
}

// BirthdateLT applies the LT predicate on the "birthdate" field.
func BirthdateLT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldBirthdate, v))
}

// BirthdateLTE applies the LTE predicate on the "birthdate" field.
func BirthdateLTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldBirthdate, v))
}

// BirthdateIsNil applies the IsNil predicate on the "birthdate" field.
func BirthdateIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldBirthdate))
}

// BirthdateNotNil applies the NotNil predicate on the "birthdate" field.
func BirthdateNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldBirthdate))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v Gender) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v Gender) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...Gender) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...Gender) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldGender, vs...))
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPosition, v))
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldPosition, v))
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldPosition, vs...))
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldPosition, vs...))
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldPosition, v))
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldPosition, v))
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldPosition, v))
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldPosition, v))
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldPosition, v))
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldPosition, v))
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldPosition, v))
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldPosition, v))
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldPosition, v))
}

// DepartmentEQ applies the EQ predicate on the "department" field.
func DepartmentEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldDepartment, v))
}

// DepartmentNEQ applies the NEQ predicate on the "department" field.
func DepartmentNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldDepartment, v))
}

// DepartmentIn applies the In predicate on the "department" field.
func DepartmentIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldDepartment, vs...))
}

// DepartmentNotIn applies the NotIn predicate on the "department" field.
func DepartmentNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldDepartment, vs...))
}

// DepartmentGT applies the GT predicate on the "department" field.
func DepartmentGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldDepartment, v))
}

// DepartmentGTE applies the GTE predicate on the "department" field.
func DepartmentGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldDepartment, v))
}

// DepartmentLT applies the LT predicate on the "department" field.
func DepartmentLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldDepartment, v))
}

// DepartmentLTE applies the LTE predicate on the "department" field.
func DepartmentLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldDepartment, v))
}

// DepartmentContains applies the Contains predicate on the "department" field.
func DepartmentContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldDepartment, v))
}

// DepartmentHasPrefix applies the HasPrefix predicate on the "department" field.
func DepartmentHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldDepartment, v))
}

// DepartmentHasSuffix applies the HasSuffix predicate on the "department" field.
func DepartmentHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldDepartment, v))
}

// DepartmentIsNil applies the IsNil predicate on the "department" field.
func DepartmentIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldDepartment))
}

// DepartmentNotNil applies the NotNil predicate on the "department" field.
func DepartmentNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldDepartment))
}

// DepartmentEqualFold applies the EqualFold predicate on the "department" field.
func DepartmentEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldDepartment, v))
}

// DepartmentContainsFold applies the ContainsFold predicate on the "department" field.
func DepartmentContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldDepartment, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldPhone, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldAvatar, v))
}

// HireDateEQ applies the EQ predicate on the "hireDate" field.
func HireDateEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldHireDate, v))
}

// HireDateNEQ applies the NEQ predicate on the "hireDate" field.
func HireDateNEQ(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldHireDate, v))
}

// HireDateIn applies the In predicate on the "hireDate" field.
func HireDateIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldHireDate, vs...))
}

// HireDateNotIn applies the NotIn predicate on the "hireDate" field.
func HireDateNotIn(vs ...time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldHireDate, vs...))
}

// HireDateGT applies the GT predicate on the "hireDate" field.
func HireDateGT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldHireDate, v))
}

// HireDateGTE applies the GTE predicate on the "hireDate" field.
func HireDateGTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldHireDate, v))
}

// HireDateLT applies the LT predicate on the "hireDate" field.
func HireDateLT(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldHireDate, v))
}

// HireDateLTE applies the LTE predicate on the "hireDate" field.
func HireDateLTE(v time.Time) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldHireDate, v))
}

// MonthlySalaryEQ applies the EQ predicate on the "monthlySalary" field.
func MonthlySalaryEQ(v int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldMonthlySalary, v))
}

// MonthlySalaryNEQ applies the NEQ predicate on the "monthlySalary" field.
func MonthlySalaryNEQ(v int) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldMonthlySalary, v))
}

// MonthlySalaryIn applies the In predicate on the "monthlySalary" field.
func MonthlySalaryIn(vs ...int) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldMonthlySalary, vs...))
}

// MonthlySalaryNotIn applies the NotIn predicate on the "monthlySalary" field.
func MonthlySalaryNotIn(vs ...int) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldMonthlySalary, vs...))
}

// MonthlySalaryGT applies the GT predicate on the "monthlySalary" field.
func MonthlySalaryGT(v int) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldMonthlySalary, v))
}

// MonthlySalaryGTE applies the GTE predicate on the "monthlySalary" field.
func MonthlySalaryGTE(v int) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldMonthlySalary, v))
}

// MonthlySalaryLT applies the LT predicate on the "monthlySalary" field.
func MonthlySalaryLT(v int) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldMonthlySalary, v))
}

// MonthlySalaryLTE applies the LTE predicate on the "monthlySalary" field.
func MonthlySalaryLTE(v int) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldMonthlySalary, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldStatus))
}

// PerformaceScoreEQ applies the EQ predicate on the "performaceScore" field.
func PerformaceScoreEQ(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPerformaceScore, v))
}

// PerformaceScoreNEQ applies the NEQ predicate on the "performaceScore" field.
func PerformaceScoreNEQ(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldPerformaceScore, v))
}

// PerformaceScoreIn applies the In predicate on the "performaceScore" field.
func PerformaceScoreIn(vs ...float64) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldPerformaceScore, vs...))
}

// PerformaceScoreNotIn applies the NotIn predicate on the "performaceScore" field.
func PerformaceScoreNotIn(vs ...float64) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldPerformaceScore, vs...))
}

// PerformaceScoreGT applies the GT predicate on the "performaceScore" field.
func PerformaceScoreGT(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldPerformaceScore, v))
}

// PerformaceScoreGTE applies the GTE predicate on the "performaceScore" field.
func PerformaceScoreGTE(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldPerformaceScore, v))
}

// PerformaceScoreLT applies the LT predicate on the "performaceScore" field.
func PerformaceScoreLT(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldPerformaceScore, v))
}

// PerformaceScoreLTE applies the LTE predicate on the "performaceScore" field.
func PerformaceScoreLTE(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldPerformaceScore, v))
}

// PerformaceScoreIsNil applies the IsNil predicate on the "performaceScore" field.
func PerformaceScoreIsNil() predicate.Employee {
	return predicate.Employee(sql.FieldIsNull(FieldPerformaceScore))
}

// PerformaceScoreNotNil applies the NotNil predicate on the "performaceScore" field.
func PerformaceScoreNotNil() predicate.Employee {
	return predicate.Employee(sql.FieldNotNull(FieldPerformaceScore))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubordinates applies the HasEdge predicate on the "subordinates" edge.
func HasSubordinates() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubordinatesTable, SubordinatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubordinatesWith applies the HasEdge predicate on the "subordinates" edge with a given conditions (other predicates).
func HasSubordinatesWith(preds ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := newSubordinatesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLeader applies the HasEdge predicate on the "leader" edge.
func HasLeader() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LeaderTable, LeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLeaderWith applies the HasEdge predicate on the "leader" edge with a given conditions (other predicates).
func HasLeaderWith(preds ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := newLeaderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Employee) predicate.Employee {
	return predicate.Employee(sql.NotPredicates(p))
}
