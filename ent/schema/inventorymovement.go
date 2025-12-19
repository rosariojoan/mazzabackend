package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InventoryMovement holds the schema definition for the InventoryMovement entity.
type InventoryMovement struct {
	ent.Schema
}

func (InventoryMovement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

var inventoryMovementTypes = []string{"IN", "OUT", "ADJUSTMENT"}

// Fields of the InventoryMovement.
func (InventoryMovement) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("category").Values(inventoryMovementTypes...).Annotations(entgql.OrderField("CATEGORY")),
		field.Float("quantity"),
		field.Float("value"),
		field.Time("date").Default(time.Now),
		field.String("source").Nillable().Optional(),
		field.String("destination").Nillable().Optional(),
		field.String("notes"),
	}
}

// Edges of the InventoryMovement.
func (InventoryMovement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("inventory_movements").Unique().Required().Annotations(
			entsql.OnDelete(entsql.Cascade),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.From("inventory", Inventory.Type).Ref("movements").Unique().Required().Annotations(
			entsql.OnDelete(entsql.SetNull),
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
	}
}

// Enable query and mutation for the Invoice schema
func (InventoryMovement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
