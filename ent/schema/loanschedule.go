package schema

import (
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LoanSchedule holds the schema definition for the LoanSchedule entity.
type LoanSchedule struct {
	ent.Schema
}

func (LoanSchedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the LoanSchedule.
func (LoanSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount").Validate(func(f float64) error {
			if f == 0 {
				return fmt.Errorf("amount should not be equal to zero")
			}
			return nil
		}),
		field.Float("amount_paid").Default(0),
		field.Time("due_date").Comment("Due date"),
		field.Time("date_paid").Nillable().Optional(),
		field.Float("interest"),
		field.Int("installment_number").NonNegative(),
		field.Float("principal"),
		field.Float("remaining_balance").Optional(),
		field.Enum("status").
			Values("pending", "paid", "partial", "overdue").
			Default("pending"),
	}
}

// Edges of the LoanSchedule.
func (LoanSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("loan", Loan.Type).Ref("loan_schedule").Required().Unique().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.To("transaction_history", AccountingEntry.Type).
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

// Enable query and mutation for the Loan schema
func (LoanSchedule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
