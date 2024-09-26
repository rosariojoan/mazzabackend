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

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
		field.Bool("isDefault").Default(false),
		field.Int("minimumStock").NonNegative().Default(0),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Int("price").NonNegative().Default(0),
		field.String("sku").NotEmpty(),
		field.Float("stock").Min(0).Default(0),
		field.Enum("category").Values("MERCHANDISE", "OTHER", "SERVICE").Annotations(entgql.OrderField("CATEGORY")),
		field.Float("unitCost").Min(0),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("products").Unique(),                                    // a product can belong to only one company
		edge.To("pictures", File.Type).Annotations(entsql.OnDelete(entsql.Cascade)),                    // a product can have many pictures
		edge.To("productMovements", ProductMovement.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a product can have many product movements
	}
}

// Enable query and mutation for the AccountingEntry schema
func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Create indexes
func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(),
		index.Fields("sku").Edges("company").Unique(),
	}
}
