package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Int("entry_group").Positive().Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
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
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
	}
}

// Edges of the Receivable
func (Receivable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("receivables").Unique().Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		), // a receivable can belong to only one company
		edge.From("invoice", Invoice.Type).Ref("receivable").Unique().Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
			entsql.OnDelete(entsql.SetNull),
		),
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
