package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AccountingEntry holds the schema definition for the AccountingEntry entity.
type AccountingEntry struct {
	ent.Schema
}

func (AccountingEntry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

var accountTypes = []string{
	"asset",
	"liability",
	"equity",
	"revenue",
	"expense",
	"taxExpense",
	"income",
	"dividendExpense",
	"contraAsset",
	"contraLiability",
	"contraEquity",
	"contraRevenue",
	"contraExpense",
}

// Fields of the AccountingEntry.
func (AccountingEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").Positive(),
		field.Int("group").Positive().Annotations(entgql.OrderField("group")),
		field.Time("date").Default(time.Now).Annotations(entgql.OrderField("date")),
		field.String("account").NotEmpty().Annotations(entgql.OrderField("account")),
		field.String("label"),
		field.Float("amount").Annotations(entgql.OrderField("amount")),
		field.String("description").Annotations(entgql.OrderField("description")),
		field.Enum("account_type").Values(accountTypes...).Annotations(entgql.OrderField("accountType")),
		field.String("category").Default("").Annotations(entgql.OrderField("category")),
		field.String("main").Default("").Annotations(entgql.OrderField("main")),
		field.Bool("is_debit").Annotations(entgql.OrderField("isDebit")),
		field.Bool("is_reversal").Default(false),
		field.Bool("reversed").Default(false),
	}
}

// Create indexes
func (AccountingEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("number").Edges("company").Unique(), // entry number should not repeat for each company
	}
}

// Edges of the AccountingEntry.
func (AccountingEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("accounting_entries").Unique(), // one company can have many accounting entries
		edge.From("user", User.Type).Ref("accounting_entries").Unique(),
		edge.From("loan", Loan.Type).Ref("transaction_history").Unique(),         // one user can have many accounting entries
		edge.From("loan_schedule", LoanSchedule.Type).Ref("transaction_history"), // many-to-many
	}
}

// Enable query and mutation for the AccountingEntry schema
func (AccountingEntry) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
