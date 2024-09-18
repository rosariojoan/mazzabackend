package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CashMovement holds the schema definition for the CashMovement entity.
type CashMovement struct {
	ent.Schema
}

func (CashMovement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the CashMovement.
func (CashMovement) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.Time("date"),
		field.Int("entryGroup").Positive(),
	}
}

// Edges of the CashMovement.
func (CashMovement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("treasury", Treasury.Type).Ref("cashMovements").Unique(), // a cash movement can belong to only one cash & cash equivalent account
	}
}

// Enable query and mutation for the CashMovement schema
func (CashMovement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
