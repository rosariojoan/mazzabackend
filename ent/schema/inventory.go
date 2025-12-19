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

// Inventory holds the schema definition for the Inventory entity.
type Inventory struct {
	ent.Schema
}

func (Inventory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

var inventoryCategories = []string{
	"RAW_MATERIALS",
	"FINISHED_GOODS",
	"MERCHANDISE",
	"SUPPLIES",
	"EQUIPMENT",
	"OTHER",
}

// Fields of the Inventory.
func (Inventory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("category").Values(inventoryCategories...).Annotations(entgql.OrderField("CATEGORY")),
		field.Float("quantity").Min(0).Annotations(entsql.Check("quantity >= 0")),
		field.String("unit").Comment("Unit of measurement. E.g. litre, unit, gram"),
		field.Float("minimum_level").Min(0),
		field.Float("current_value").Min(0).Annotations(entgql.OrderField("CURRENT_VALUE")),
		field.String("notes").NotEmpty(),
	}
}

// Edges of the Inventory.
func (Inventory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("inventory").Unique().Required().Annotations(
			entsql.OnDelete(entsql.Cascade),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.To("movements", InventoryMovement.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
	}
}

// Company-wise, there should not be two items of inventory with the same name
func (Inventory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(),
	}
}

// Enable query and mutation for the Invoice schema
func (Inventory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
