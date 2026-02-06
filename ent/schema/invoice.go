package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

func (Invoice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

var invoiceStatus = []string{"draft", "cancelled", "pending", "paid", "overdue", "defaulted"}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		// field.String("company_logo").Nillable().Optional(),
		field.Bool("is_invoice").Default(true),
		field.String("company_name"),
		field.String("company_tax_id").Nillable().Optional(),
		field.String("company_address"),
		field.String("company_city"),
		field.String("company_email").Nillable().Optional(),
		field.String("company_phone").Nillable().Optional(),

		field.String("currency"),
		field.String("number").Nillable().Optional(),
		field.Time("issue_date").Annotations(entgql.OrderField("issueDate")),
		field.Time("due_date").Annotations(entgql.OrderField("dueDate")),
		field.Time("paid_at").Nillable().Optional().Annotations(entgql.OrderField("paidAt")),
		field.Enum("status").Values(invoiceStatus...).Default("paid").Annotations(entgql.OrderField("status")),

		field.String("customer_name").Default("other").Optional(),
		field.String("customer_tax_id").Nillable().Optional(),
		field.String("customer_address").Nillable().Optional(),
		field.String("customer_city").Nillable().Optional(),
		field.String("customer_email").Nillable().Optional(),
		field.String("customer_phone").Nillable().Optional(),

		field.String("items").NotEmpty().Comment("stringified JSON of product rows"),
		field.Float("subtotal").Positive(),
		field.Float("tax").Min(0),
		field.Float("total").Positive().Annotations(entgql.OrderField("total")),
		field.String("terms").Optional(),
		field.String("keywords").NotEmpty().MaxLen(255),

		// field.String("notes").Nillable().Optional(),
		// field.String("payment_method").Nillable().Optional(),
		// field.String("bank_name").Nillable().Optional(),
		// field.String("bank_agency").Nillable().Optional(),
		// field.String("bank_account_number").Nillable().Optional(),
		// field.String("bank_account_name").Nillable().Optional(),

		// field.String("storage_URI").Nillable().Sensitive().Optional().Annotations(
		// // entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		// ),
		// field.String("URL").Nillable().Optional().Annotations(
		// // entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		// ),
		// field.String("filename").Nillable().Optional().Annotations(
		// // entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		// ),
		// field.Float("size").Positive().Nillable().Optional().Annotations(
		// // entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		// ).Comment("File size in KB"),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("invoices").Unique().Required().Annotations(
			entsql.OnDelete(entsql.Cascade),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.From("issued_by", User.Type).Ref("issued_invoices").Unique().Annotations(
			entsql.OnDelete(entsql.SetNull),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.From("client", Customer.Type).Ref("invoices").Unique().Annotations(
			entsql.OnDelete(entsql.SetNull),
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
		edge.To("receivable", Receivable.Type).Unique().Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
			entsql.OnDelete(entsql.SetNull),
		),
	}
}

// Enable query and mutation for the Invoice schema
func (Invoice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Create indexes
func (Invoice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("number").Edges("company").Unique().Annotations(entsql.IndexWhere("number IS NOT NULL")),
	}
}
