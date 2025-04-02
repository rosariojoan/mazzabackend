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

// var accountTypes = []string{
// 	"asset",
// 	"liability",
// 	"equity",
// 	"revenue",
// 	"expense",
// 	"taxExpense",
// 	"income",
// 	"dividendExpense",
// 	"contraAsset",
// 	"contraLiability",
// 	"contraRevenue",
// 	"contraExpense",
// }

var accountTypes = []string{
	"ASSET",
	"LIABILITY",
	"EQUITY",
	"REVENUE",
	"EXPENSE",
	"TAX_EXPENSE",
	"INCOME",
	"DIVIDEND_EXPENSE",
	"CONTRA_ASSET",
	"CONTRA_LIABILITY",
	"CONTRA_REVENUE",
	"CONTRA_EXPENSE",
}

// Fields of the AccountingEntry.
func (AccountingEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").Positive(),
		field.Int("group").Positive().Annotations(entgql.OrderField("GROUP")),
		field.Time("date").Default(time.Now).Annotations(entgql.OrderField("DATE")),
		field.String("account").NotEmpty().Annotations(entgql.OrderField("ACCOUNT")),
		field.String("label"),
		field.Float("amount").Annotations(entgql.OrderField("AMOUNT")),
		field.String("description").Annotations(entgql.OrderField("DESCRIPTION")),
		field.Enum("accountType").Values(accountTypes...).Annotations(entgql.OrderField("ACCOUNTTYPE")),
		field.String("category").Default("").Annotations(entgql.OrderField("CATEGORY")),
		field.Bool("isDebit").Annotations(entgql.OrderField("ISDEBIT")),
		field.Bool("isReversal").Default(false),
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
		edge.From("company", Company.Type).Ref("accountingEntries").Unique(), // one company can have many accounting entries
		edge.From("user", User.Type).Ref("accountingEntries").Unique(),       // one user can have many accounting entries
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
