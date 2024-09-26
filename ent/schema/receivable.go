package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Receivable holds the schema definition for the Receivable entity.
type Receivable struct {
	ent.Schema
}

func (Receivable) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Receivable.
func (Receivable) Fields() []ent.Field {
	return []ent.Field{
		field.Int("entryGroup").Positive(),
		field.Time("date"),
		field.Float("outstandingBalance"),
		field.Float("totalTransaction"),
		field.Int("daysDue").NonNegative().Annotations(entgql.OrderField("DAYSDUE")),
		field.Enum("status").Values("paid", "unpaid", "doubtful", "default").Annotations(entgql.OrderField("STATUS")),
	}
}

// Edges of the Receivable
func (Receivable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("receivables").Unique(), // a receivable can belong to only one client
	}
}

// Enable query and mutation for the Receivable schema
func (Receivable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
