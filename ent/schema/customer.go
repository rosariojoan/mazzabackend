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
		field.String("address").Default("").Optional(),
		field.String("city").Default("").Optional().Annotations(entgql.OrderField("CITY")),
		field.String("country").Default("").Optional(),
		field.String("description").Default("").Optional(),
		field.String("email").Default("").Optional(),
		field.Bool("isDefault").Default(false).Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("phone").Default("").Optional(),
		field.String("taxId").Default("").Optional(),
	}
}

// Create indexes
// Company-wise, clients must have unique names and unique tax IDs
func (Customer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "taxId").Edges("company").Unique(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("customers").Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)), // a client can belong to only one company
		edge.To("loans", Loan.Type).Annotations(
			entsql.OnDelete(entsql.SetNull),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.To("receivables", Receivable.Type).Annotations(entsql.OnDelete(entsql.Cascade)).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)), // a client can have many receivables
		edge.To("invoices", Invoice.Type).Annotations(entsql.OnDelete(entsql.SetNull)).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)), // a client can have many invoices
	}
}

// Enable query and mutation for the CompanyClient schema
func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
