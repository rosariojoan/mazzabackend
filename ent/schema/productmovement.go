package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductMovement holds the schema definition for the ProductMovement entity.
type ProductMovement struct {
	ent.Schema
}

func (ProductMovement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the ProductMovement.
func (ProductMovement) Fields() []ent.Field {
	return []ent.Field{
		field.Int("entryGroup").Positive(),
		field.Float("averageCost").Min(0),
		field.Float("unitCost").Min(0),
		field.Float("price").Min(0).Default(0),
		field.Int("quantity"),
	}
}

// Edges of the ProductMovement.
func (ProductMovement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("productMovements").Unique(), // a product movement can belong to only one product
	}
}

// Enable query and mutation for the Payable schema
func (ProductMovement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
