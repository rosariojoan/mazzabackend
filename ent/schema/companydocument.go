package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

var companyDocumentCategories = []string{
	"LEGAL",
	"CONTRACT",
	"LICENSE",
	"TAX",
	"HR",
}

var companyDocumentStatus = []string{
	"PENDING",
	"APPROVED",
	"REJECTED",
	"EXPIRED",
}

// CompanyDocument holds the schema definition for the CompanyDocument entity.
type CompanyDocument struct {
	ent.Schema
}

func (CompanyDocument) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the CompanyDocument.
func (CompanyDocument) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("keywords").NotEmpty(),
		field.Enum("category").Values(companyDocumentCategories...).Annotations(entgql.OrderField("CATEGORY")),
		field.Int("size").Positive().Comment("File size in kilobyte"),
		field.String("fileType").Nillable().Comment("mimetype e.g. application/pdf"),
		field.Enum("status").Values(companyDocumentStatus...).
			Annotations(entgql.OrderField("STATUS"), entgql.Skip(entgql.SkipMutationCreateInput)),
		field.String("url").NotEmpty(),
		field.String("storageURI").NotEmpty().Sensitive().Comment("Firebase cloud storage URI. Not exposed to the client"),
		field.String("thumbnail").Nillable().Optional(),
		field.Time("expiryDate").Nillable().Annotations(entgql.OrderField("EXPIRY_DATE")),
	}
}

// Edges of the CompanyDocument.
func (CompanyDocument) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("documents").Unique().Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)), // one company can upload multiple documents
		edge.From("uploadedBy", User.Type).Ref("uploadedDocuments").Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)), // one user can upload multiple documents
		edge.From("approvedBy", User.Type).Ref("approvedDocuments").Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)), // one user approve multiple documents
	}
}

// Enable query and mutation for the CompanyDocument schema
func (CompanyDocument) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
