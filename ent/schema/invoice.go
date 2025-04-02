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

var invoiceStatus = []string{"DRAFT", "CANCELED", "PENDING", "PAID", "OVERDUE", "DEFAULTED"}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("companyLogo").Nillable().Optional(),
		field.String("companyName"),
		field.String("companyTaxID").Nillable().Optional(),
		field.String("companyAddress"),
		field.String("companyCity"),
		field.String("companyEmail").Nillable().Optional(),
		field.String("companyPhone").Nillable().Optional(),

		field.String("number").Nillable().Optional(),
		field.Time("issueDate").Annotations(entgql.OrderField("ISSUE_DATE")),
		field.Time("dueDate").Annotations(entgql.OrderField("DUE_DATE")),
		field.Time("paidAt").Nillable().Optional().Annotations(entgql.OrderField("PAID_AT")),
		field.Enum("status").Values(invoiceStatus...).Default("PAID").Annotations(entgql.OrderField("STATUS")),

		field.String("customerName").Default("other").Optional(),
		field.String("customerTaxID").Nillable().Optional(),
		field.String("customerAddress").Nillable().Optional(),
		field.String("customerCity").Nillable().Optional(),
		field.String("customerEmail").Nillable().Optional(),
		field.String("customerPhone").Nillable().Optional(),

		field.String("items").NotEmpty().Comment("stringified JSON of product rows"),
		field.Float("subtotal").Positive(),
		field.Float("tax").Min(0),
		field.Float("total").Positive().Annotations(entgql.OrderField("TOTAL")),

		field.String("notes").Nillable().Optional(),
		field.String("paymentMethod").Nillable().Optional(),
		field.String("bankName").Nillable().Optional(),
		field.String("bankAgency").Nillable().Optional(),
		field.String("bankAccountNumber").Nillable().Optional(),
		field.String("bankAccountName").Nillable().Optional(),
		
		field.String("storageURI").Nillable().Sensitive().Optional().Annotations(
			// entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.String("URL").Nillable().Optional().Annotations(
			// entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.String("filename").Nillable().Optional().Annotations(
			// entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.Float("size").Positive().Nillable().Optional().Annotations(
			// entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		).Comment("File size in KB"),
		field.String("keywords").NotEmpty().MaxLen(255),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("invoices").Unique().Required().Annotations(
			entsql.OnDelete(entsql.Cascade),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.From("issuedBy", User.Type).Ref("issuedInvoices").Unique().Annotations(
			entsql.OnDelete(entsql.SetNull),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.From("client", Customer.Type).Ref("invoices").Unique().Annotations(
			entsql.OnDelete(entsql.SetNull),
			entgql.Skip(entgql.SkipMutationUpdateInput),
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
