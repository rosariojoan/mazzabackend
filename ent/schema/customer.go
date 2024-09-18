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

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
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
// Company-wise, clients must have unique names and unique tax IDs
func (Customer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(),
		index.Fields("taxId").Edges("company").Unique(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("customers").Unique(),                         // a client can belong to only one company
		edge.To("receivables", Receivable.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a client can have many receivables
	}
}

// Enable query and mutation for the CompanyClient schema
func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
