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
		field.Int("entryGroup").Positive(),
		field.Time("date").Annotations(entgql.OrderField("DATE")),
		field.String("name").Default("Diversos"),
		field.Float("amountInDefault").Min(0).Default(0).Annotations(entgql.OrderField("AMOUNT_IN_DEFAULT")),
		field.Float("outstandingBalance").Min(0).Annotations(entgql.OrderField("OUTSTANDING_BALANCE")),
		field.Float("totalTransaction").Min(0).Annotations(entgql.OrderField("TOTAL_TRANSACTION")),
		field.Time("dueDate").Annotations(entgql.OrderField("DUEDATE")),
		field.Enum("status").Values("paid", "pending", "overdue", "default").Annotations(entgql.OrderField("STATUS")),
	}
}

// Edges of the Receivable
func (Receivable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("receivables").Unique(), // a receivable can belong to only one company
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
