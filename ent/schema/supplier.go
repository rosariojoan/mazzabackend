package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Supplier holds the schema definition for the Supplier entity.
type Supplier struct {
	ent.Schema
}

func (Supplier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Supplier.
func (Supplier) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.String("city"),
		field.String("country"),
		field.String("description"),
		field.String("email"),
		field.Bool("isDefault").Default(false).Optional(),
		field.String("name").NotEmpty(),
		field.String("phone"),
		field.String("taxId").NotEmpty(),
	}
}

// Create indexes
// Company-wise, supplier must have unique names and unique tax IDs
func (Supplier) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(),
		index.Fields("taxId").Edges("company").Unique(),
	}
}

// Edges of the Supplier.
func (Supplier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("suppliers").Unique(),                   // a supplier can belong to only one company
		edge.To("payables", Payable.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a supplier can have multiple accounts payable
	}
}

// Enable query and mutation for the AccountingEntry schema
func (Supplier) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
