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

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Nillable().Optional(),
		field.String("base_currency").Default("MZN"),
		field.String("ceo_name").Nillable().Optional(),
		field.String("city").Annotations(entgql.OrderField("CITY")),
		field.String("country").Annotations(entgql.OrderField("COUNTRY")),
		field.Time("established_at").Annotations(entgql.OrderField("ESTABLISHEDAT")),
		field.String("description").Nillable().Optional(),
		field.String("email").Nillable().Optional(),
		field.String("industry").Nillable().Optional(),
		field.Time("last_entry_date").Nillable().Optional(),
		field.Int32("last_invoice_number").Optional().Default(0).NonNegative(),
		field.String("logo_URL").Nillable().Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("logo_storage_URI").Nillable().Optional().Nillable().Sensitive().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.Int32("number_employees").NonNegative().Default(0),
		field.String("phone").Nillable().Optional(),
		field.String("tax_id").Nillable().Optional(),
		field.Float("vat_rate").Default(0.16),
		field.String("website").Nillable().Optional(),
		field.Bool("incomplete_setup").Optional().Default(true),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("available_roles", UserRole.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a company must have at least one role to which its users are assigned
		edge.To("accounting_entries", AccountingEntry.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("customers", Customer.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("documents", CompanyDocument.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("employees", Employee.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("files", File.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("inventory", Inventory.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("inventory_movements", InventoryMovement.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("invoices", Invoice.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("loans", Loan.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("loan_schedule", LoanSchedule.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("member_signup_tokens", MemberSignupToken.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("products", Product.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("projects", Project.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("payables", Payable.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("receivables", Receivable.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("suppliers", Supplier.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a company can have many tokens
		edge.To("treasuries", Treasury.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("work_shifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("users", User.Type), // a company can have many users

		// One-to-many relationship: each company can have many daughter companies, but only one parent company
		edge.To("daughter_companies", Company.Type),
		edge.From("parent_company", Company.Type).Ref("daughter_companies").Unique(),
	}
}

// Create indexes
func (Company) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country", "name").Unique(),
	}
}

// Enable query and mutation for the Company schema
func (Company) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		// entgql.MultiOrder(),
	}
}
