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
		field.Float("amount").Min(0).Annotations(entgql.OrderField("amount")),
		field.Enum("category").
			Values("auto", "equipment", "expansion", "line_of_credit", "mortgage", "personal", "working_capital").
			Annotations(entgql.OrderField("category")),
		field.String("collateral").Optional(),
		field.String("description").Optional(),
		field.Float("interest_rate").Min(0).Annotations(entgql.OrderField("interestRate")),
		field.Int("installments").Min(0),
		field.Time("maturity_date").Annotations(entgql.OrderField("maturityDate")),
		field.Time("next_payment").Optional().Annotations(entgql.OrderField("nextPayment")),
		field.Float("next_payment_amount").Min(0).Default(0).Optional().Annotations(entgql.OrderField("nextPaymentAmount")),
		field.Float("outstanding_balance").Min(0).Annotations(entgql.OrderField("outstandingBalance")),
		field.Enum("payment_frequency").
			Values("daily", "weekly", "biweekly", "monthly", "quarterly", "semiannual", "annual").
			Default("monthly").
			Annotations(entgql.OrderField("paymentFrequency")),
		field.Int("paid_installments").Default(0),
		field.Enum("paymentType").Values("bullet", "fixedPayment", "fixedPrincipal", "interestOnly").Default("fixedPayment").Comment(
			"Bullet loan: the payment is made in a single shot at the maturity date.\n" +
				"Interest only: interest only paid at the given periods. Principal is paid at the maturity.\n" +
				"FixedPayment: equal payments of interest + principal are made at the given periods until the maturity.\n" +
				"FixedPrincipal: a fixed amount of the principal + variable interest is paid at each period until the maturity.",
		),
		field.String("counterparty_name").NotEmpty(),
		field.Time("start_date").Default(time.Now).Annotations(entgql.OrderField("startDate")),
		field.Enum("status").Values("active", "paid"),
		field.Bool("is_lending").Default(false).Comment("True if the company is the lender"),
	}
}

// Edges of the Loan.
func (Loan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client", Customer.Type).Ref("loan_schedule").Unique().Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
		edge.From("supplier", Supplier.Type).Ref("loan_schedule").Unique().Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
		edge.From("company", Company.Type).Ref("loans").Unique().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		), // a loan can belong to only one company
		edge.To("loan_schedule", LoanSchedule.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("transaction_history", AccountingEntry.Type).
			Annotations(entsql.OnDelete(entsql.SetNull)),
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
