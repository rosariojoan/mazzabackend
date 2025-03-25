// Code generated by ent, DO NOT EDIT.

package companydocument

import (
	"mazza/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldDeletedAt, v))
}

// Filename applies equality check predicate on the "filename" field. It's identical to FilenameEQ.
func Filename(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldFilename, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldTitle, v))
}

// Keywords applies equality check predicate on the "keywords" field. It's identical to KeywordsEQ.
func Keywords(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldKeywords, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldSize, v))
}

// FileType applies equality check predicate on the "fileType" field. It's identical to FileTypeEQ.
func FileType(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldFileType, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldURL, v))
}

// StorageURI applies equality check predicate on the "storageURI" field. It's identical to StorageURIEQ.
func StorageURI(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldStorageURI, v))
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldThumbnail, v))
}

// ExpiryDate applies equality check predicate on the "expiryDate" field. It's identical to ExpiryDateEQ.
func ExpiryDate(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldExpiryDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotNull(FieldDeletedAt))
}

// FilenameEQ applies the EQ predicate on the "filename" field.
func FilenameEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldFilename, v))
}

// FilenameNEQ applies the NEQ predicate on the "filename" field.
func FilenameNEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldFilename, v))
}

// FilenameIn applies the In predicate on the "filename" field.
func FilenameIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldFilename, vs...))
}

// FilenameNotIn applies the NotIn predicate on the "filename" field.
func FilenameNotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldFilename, vs...))
}

// FilenameGT applies the GT predicate on the "filename" field.
func FilenameGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldFilename, v))
}

// FilenameGTE applies the GTE predicate on the "filename" field.
func FilenameGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldFilename, v))
}

// FilenameLT applies the LT predicate on the "filename" field.
func FilenameLT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldFilename, v))
}

// FilenameLTE applies the LTE predicate on the "filename" field.
func FilenameLTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldFilename, v))
}

// FilenameContains applies the Contains predicate on the "filename" field.
func FilenameContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldFilename, v))
}

// FilenameHasPrefix applies the HasPrefix predicate on the "filename" field.
func FilenameHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldFilename, v))
}

// FilenameHasSuffix applies the HasSuffix predicate on the "filename" field.
func FilenameHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldFilename, v))
}

// FilenameEqualFold applies the EqualFold predicate on the "filename" field.
func FilenameEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldFilename, v))
}

// FilenameContainsFold applies the ContainsFold predicate on the "filename" field.
func FilenameContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldFilename, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldTitle, v))
}

// KeywordsEQ applies the EQ predicate on the "keywords" field.
func KeywordsEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldKeywords, v))
}

// KeywordsNEQ applies the NEQ predicate on the "keywords" field.
func KeywordsNEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldKeywords, v))
}

// KeywordsIn applies the In predicate on the "keywords" field.
func KeywordsIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldKeywords, vs...))
}

// KeywordsNotIn applies the NotIn predicate on the "keywords" field.
func KeywordsNotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldKeywords, vs...))
}

// KeywordsGT applies the GT predicate on the "keywords" field.
func KeywordsGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldKeywords, v))
}

// KeywordsGTE applies the GTE predicate on the "keywords" field.
func KeywordsGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldKeywords, v))
}

// KeywordsLT applies the LT predicate on the "keywords" field.
func KeywordsLT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldKeywords, v))
}

// KeywordsLTE applies the LTE predicate on the "keywords" field.
func KeywordsLTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldKeywords, v))
}

// KeywordsContains applies the Contains predicate on the "keywords" field.
func KeywordsContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldKeywords, v))
}

// KeywordsHasPrefix applies the HasPrefix predicate on the "keywords" field.
func KeywordsHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldKeywords, v))
}

// KeywordsHasSuffix applies the HasSuffix predicate on the "keywords" field.
func KeywordsHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldKeywords, v))
}

// KeywordsEqualFold applies the EqualFold predicate on the "keywords" field.
func KeywordsEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldKeywords, v))
}

// KeywordsContainsFold applies the ContainsFold predicate on the "keywords" field.
func KeywordsContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldKeywords, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldCategory, vs...))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldSize, v))
}

// FileTypeEQ applies the EQ predicate on the "fileType" field.
func FileTypeEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldFileType, v))
}

// FileTypeNEQ applies the NEQ predicate on the "fileType" field.
func FileTypeNEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldFileType, v))
}

// FileTypeIn applies the In predicate on the "fileType" field.
func FileTypeIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldFileType, vs...))
}

// FileTypeNotIn applies the NotIn predicate on the "fileType" field.
func FileTypeNotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldFileType, vs...))
}

// FileTypeGT applies the GT predicate on the "fileType" field.
func FileTypeGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldFileType, v))
}

// FileTypeGTE applies the GTE predicate on the "fileType" field.
func FileTypeGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldFileType, v))
}

// FileTypeLT applies the LT predicate on the "fileType" field.
func FileTypeLT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldFileType, v))
}

// FileTypeLTE applies the LTE predicate on the "fileType" field.
func FileTypeLTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldFileType, v))
}

// FileTypeContains applies the Contains predicate on the "fileType" field.
func FileTypeContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldFileType, v))
}

// FileTypeHasPrefix applies the HasPrefix predicate on the "fileType" field.
func FileTypeHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldFileType, v))
}

// FileTypeHasSuffix applies the HasSuffix predicate on the "fileType" field.
func FileTypeHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldFileType, v))
}

// FileTypeEqualFold applies the EqualFold predicate on the "fileType" field.
func FileTypeEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldFileType, v))
}

// FileTypeContainsFold applies the ContainsFold predicate on the "fileType" field.
func FileTypeContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldFileType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldStatus, vs...))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldURL, v))
}

// StorageURIEQ applies the EQ predicate on the "storageURI" field.
func StorageURIEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldStorageURI, v))
}

// StorageURINEQ applies the NEQ predicate on the "storageURI" field.
func StorageURINEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldStorageURI, v))
}

// StorageURIIn applies the In predicate on the "storageURI" field.
func StorageURIIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldStorageURI, vs...))
}

// StorageURINotIn applies the NotIn predicate on the "storageURI" field.
func StorageURINotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldStorageURI, vs...))
}

// StorageURIGT applies the GT predicate on the "storageURI" field.
func StorageURIGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldStorageURI, v))
}

// StorageURIGTE applies the GTE predicate on the "storageURI" field.
func StorageURIGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldStorageURI, v))
}

// StorageURILT applies the LT predicate on the "storageURI" field.
func StorageURILT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldStorageURI, v))
}

// StorageURILTE applies the LTE predicate on the "storageURI" field.
func StorageURILTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldStorageURI, v))
}

// StorageURIContains applies the Contains predicate on the "storageURI" field.
func StorageURIContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldStorageURI, v))
}

// StorageURIHasPrefix applies the HasPrefix predicate on the "storageURI" field.
func StorageURIHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldStorageURI, v))
}

// StorageURIHasSuffix applies the HasSuffix predicate on the "storageURI" field.
func StorageURIHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldStorageURI, v))
}

// StorageURIEqualFold applies the EqualFold predicate on the "storageURI" field.
func StorageURIEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldStorageURI, v))
}

// StorageURIContainsFold applies the ContainsFold predicate on the "storageURI" field.
func StorageURIContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldStorageURI, v))
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldThumbnail, v))
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldThumbnail, v))
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldThumbnail, vs...))
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldThumbnail, vs...))
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldThumbnail, v))
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldThumbnail, v))
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldThumbnail, v))
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldThumbnail, v))
}

// ThumbnailContains applies the Contains predicate on the "thumbnail" field.
func ThumbnailContains(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContains(FieldThumbnail, v))
}

// ThumbnailHasPrefix applies the HasPrefix predicate on the "thumbnail" field.
func ThumbnailHasPrefix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasPrefix(FieldThumbnail, v))
}

// ThumbnailHasSuffix applies the HasSuffix predicate on the "thumbnail" field.
func ThumbnailHasSuffix(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldHasSuffix(FieldThumbnail, v))
}

// ThumbnailIsNil applies the IsNil predicate on the "thumbnail" field.
func ThumbnailIsNil() predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIsNull(FieldThumbnail))
}

// ThumbnailNotNil applies the NotNil predicate on the "thumbnail" field.
func ThumbnailNotNil() predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotNull(FieldThumbnail))
}

// ThumbnailEqualFold applies the EqualFold predicate on the "thumbnail" field.
func ThumbnailEqualFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEqualFold(FieldThumbnail, v))
}

// ThumbnailContainsFold applies the ContainsFold predicate on the "thumbnail" field.
func ThumbnailContainsFold(v string) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldContainsFold(FieldThumbnail, v))
}

// ExpiryDateEQ applies the EQ predicate on the "expiryDate" field.
func ExpiryDateEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldEQ(FieldExpiryDate, v))
}

// ExpiryDateNEQ applies the NEQ predicate on the "expiryDate" field.
func ExpiryDateNEQ(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNEQ(FieldExpiryDate, v))
}

// ExpiryDateIn applies the In predicate on the "expiryDate" field.
func ExpiryDateIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldIn(FieldExpiryDate, vs...))
}

// ExpiryDateNotIn applies the NotIn predicate on the "expiryDate" field.
func ExpiryDateNotIn(vs ...time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldNotIn(FieldExpiryDate, vs...))
}

// ExpiryDateGT applies the GT predicate on the "expiryDate" field.
func ExpiryDateGT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGT(FieldExpiryDate, v))
}

// ExpiryDateGTE applies the GTE predicate on the "expiryDate" field.
func ExpiryDateGTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldGTE(FieldExpiryDate, v))
}

// ExpiryDateLT applies the LT predicate on the "expiryDate" field.
func ExpiryDateLT(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLT(FieldExpiryDate, v))
}

// ExpiryDateLTE applies the LTE predicate on the "expiryDate" field.
func ExpiryDateLTE(v time.Time) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.FieldLTE(FieldExpiryDate, v))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.CompanyDocument {
	return predicate.CompanyDocument(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.CompanyDocument {
	return predicate.CompanyDocument(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUploadedBy applies the HasEdge predicate on the "uploadedBy" edge.
func HasUploadedBy() predicate.CompanyDocument {
	return predicate.CompanyDocument(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UploadedByTable, UploadedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUploadedByWith applies the HasEdge predicate on the "uploadedBy" edge with a given conditions (other predicates).
func HasUploadedByWith(preds ...predicate.User) predicate.CompanyDocument {
	return predicate.CompanyDocument(func(s *sql.Selector) {
		step := newUploadedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApprovedBy applies the HasEdge predicate on the "approvedBy" edge.
func HasApprovedBy() predicate.CompanyDocument {
	return predicate.CompanyDocument(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApprovedByTable, ApprovedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApprovedByWith applies the HasEdge predicate on the "approvedBy" edge with a given conditions (other predicates).
func HasApprovedByWith(preds ...predicate.User) predicate.CompanyDocument {
	return predicate.CompanyDocument(func(s *sql.Selector) {
		step := newApprovedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CompanyDocument) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CompanyDocument) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CompanyDocument) predicate.CompanyDocument {
	return predicate.CompanyDocument(sql.NotPredicates(p))
}
