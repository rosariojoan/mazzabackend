package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payable holds the schema definition for the Payable entity.
type Payable struct {
	ent.Schema
}

func (Payable) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Payable.
func (Payable) Fields() []ent.Field {
	return []ent.Field{
		field.Int("entryGroup").Positive(),
		field.Time("date"),
		field.String("name").Default("Diversos"),
		field.Float("outstandingBalance"),
		field.Float("totalTransaction"),
		field.Time("dueDate").Annotations(entgql.OrderField("DUEDATE")),
		field.Enum("status").Values("paid", "pending", "default").Annotations(entgql.OrderField("STATUS")),
	}
}

// Edges of the Payable.
func (Payable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("payables").Unique(), // a payable can belong to only one company
	}
}

// Enable query and mutation for the Payable schema
func (Payable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
