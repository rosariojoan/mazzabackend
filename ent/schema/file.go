package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

var fileCategory = []string{
	"invoice",
	"salesQuotation",
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("category").Values(fileCategory...).Annotations(entgql.OrderField("CATEGORY")),
		field.String("extension"),
		field.String("size"),
		field.String("url").Unique(),
		field.String("description"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("files").Unique(),    // one company can have many files
		edge.From("product", Product.Type).Ref("pictures").Unique(), // one product can have many pictures
		// edge.From("accountingEntries", AccountingEntry.Type).Ref("attachments"), // one accounting entry can have many files
	}
}

// Enable query and mutation for the AccountingEntry schema
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
