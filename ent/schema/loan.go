package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Loan holds the schema definition for the Loan entity.
type Loan struct {
	ent.Schema
}

func (Loan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Loan.
func (Loan) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount").Min(0),
		field.Enum("category").
			Values("auto", "equipment", "expansion", "line_of_credit", "mortgage", "personal", "working_capital"),
		field.String("collateral").Optional(),
		field.String("description").Optional(),
		field.Float("interestRate").Min(0),
		field.Int("installments").Min(0),
		field.Time("maturityDate"),
		field.Time("nextPayment").Optional(),
		field.Float("nextPaymentAmount").Min(0).Default(0).Optional(),
		field.Float("outstandingBalance").Min(0),
		field.Enum("paymentFrequency").
			Values("weekly", "biweekly", "monthly", "quartely", "semiannual", "annual").
			Default("monthly"),
		field.Int("paidInstallments").Default(0),
		field.String("provider").NotEmpty(),
		field.Time("startDate").Default(time.Now),
		field.Enum("status").Values("active", "paid"),
	}
}

// Edges of the Loan.
func (Loan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("loans").Unique(), // a loan can belong to only one company
	}
}

// Enable query and mutation for the Loan schema
func (Loan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
