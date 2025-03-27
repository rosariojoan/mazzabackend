// Code generated by ent, DO NOT EDIT.

package membersignuptoken

import (
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldEmail, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldToken, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldAvatar, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldNote, v))
}

// NumberAccessed applies equality check predicate on the "numberAccessed" field. It's identical to NumberAccessedEQ.
func NumberAccessed(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldNumberAccessed, v))
}

// ExpiresAt applies equality check predicate on the "expiresAt" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldExpiresAt, v))
}

// AlreadyUsed applies equality check predicate on the "alreadyUsed" field. It's identical to AlreadyUsedEQ.
func AlreadyUsed(v bool) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldAlreadyUsed, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContainsFold(FieldEmail, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContainsFold(FieldToken, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContainsFold(FieldAvatar, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldRole, vs...))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldHasSuffix(FieldNote, v))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldContainsFold(FieldNote, v))
}

// NumberAccessedEQ applies the EQ predicate on the "numberAccessed" field.
func NumberAccessedEQ(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldNumberAccessed, v))
}

// NumberAccessedNEQ applies the NEQ predicate on the "numberAccessed" field.
func NumberAccessedNEQ(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldNumberAccessed, v))
}

// NumberAccessedIn applies the In predicate on the "numberAccessed" field.
func NumberAccessedIn(vs ...int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldNumberAccessed, vs...))
}

// NumberAccessedNotIn applies the NotIn predicate on the "numberAccessed" field.
func NumberAccessedNotIn(vs ...int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldNumberAccessed, vs...))
}

// NumberAccessedGT applies the GT predicate on the "numberAccessed" field.
func NumberAccessedGT(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldNumberAccessed, v))
}

// NumberAccessedGTE applies the GTE predicate on the "numberAccessed" field.
func NumberAccessedGTE(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldNumberAccessed, v))
}

// NumberAccessedLT applies the LT predicate on the "numberAccessed" field.
func NumberAccessedLT(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldNumberAccessed, v))
}

// NumberAccessedLTE applies the LTE predicate on the "numberAccessed" field.
func NumberAccessedLTE(v int) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldNumberAccessed, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expiresAt" field.
func ExpiresAtEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expiresAt" field.
func ExpiresAtNEQ(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expiresAt" field.
func ExpiresAtIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expiresAt" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expiresAt" field.
func ExpiresAtGT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expiresAt" field.
func ExpiresAtGTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expiresAt" field.
func ExpiresAtLT(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expiresAt" field.
func ExpiresAtLTE(v time.Time) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldLTE(FieldExpiresAt, v))
}

// AlreadyUsedEQ applies the EQ predicate on the "alreadyUsed" field.
func AlreadyUsedEQ(v bool) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldEQ(FieldAlreadyUsed, v))
}

// AlreadyUsedNEQ applies the NEQ predicate on the "alreadyUsed" field.
func AlreadyUsedNEQ(v bool) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.FieldNEQ(FieldAlreadyUsed, v))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.MemberSignupToken {
	return predicate.MemberSignupToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedBy applies the HasEdge predicate on the "createdBy" edge.
func HasCreatedBy() predicate.MemberSignupToken {
	return predicate.MemberSignupToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "createdBy" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.User) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(func(s *sql.Selector) {
		step := newCreatedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemberSignupToken) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemberSignupToken) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemberSignupToken) predicate.MemberSignupToken {
	return predicate.MemberSignupToken(sql.NotPredicates(p))
}
