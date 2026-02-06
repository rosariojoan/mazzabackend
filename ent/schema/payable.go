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
		field.Int("entry_group").Positive().
			Annotations(entgql.Skip(entgql.SkipMutationUpdateInput)),
		field.Time("date").Annotations(
			entgql.OrderField("date"),
			entgql.Skip(entgql.SkipMutationUpdateInput)),
		field.String("name").Default("Diversos"),
		field.Float("amount_in_default").Min(0).Default(0).Annotations(
			entgql.OrderField("amountInDefault"),
			entgql.Skip(entgql.SkipMutationUpdateInput)),
		field.Float("outstanding_balance").Min(0).Annotations(
			entgql.OrderField("outstandingBalance"),
			entgql.Skip(entgql.SkipMutationUpdateInput)),
		field.Float("total_transaction").Min(0).Annotations(
			entgql.OrderField("totalTransaction"),
			entgql.Skip(entgql.SkipMutationUpdateInput)),
		field.Time("due_date").Annotations(entgql.OrderField("dueDate")),
		field.Enum("status").Values("paid", "pending", "overdue", "default").Annotations(
			entgql.OrderField("status"),
			entgql.Skip(entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the Payable.
func (Payable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("payables").Unique().Immutable(), // a payable can belong to only one company
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
