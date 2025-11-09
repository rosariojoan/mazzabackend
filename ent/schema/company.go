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
		field.String("baseCurrency").Default("mzn"),
		field.String("ceoName").Nillable().Optional(),
		field.String("city").Annotations(entgql.OrderField("CITY")),
		field.String("country").Annotations(entgql.OrderField("COUNTRY")),
		field.Time("establishedAt").Annotations(entgql.OrderField("ESTABLISHEDAT")),
		field.String("description").Nillable().Optional(),
		field.String("email").Nillable().Optional(),
		field.String("industry").Nillable().Optional(),
		field.Time("lastEntryDate").Nillable(),
		field.Int32("lastInvoiceNumber").Optional().Default(0).NonNegative(),
		field.String("logoURL").Nillable().Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("logoStorageURI").Nillable().Optional().Nillable().Sensitive().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.Int32("numberOfEmployees").NonNegative().Default(0),
		field.String("phone").Nillable().Optional(),
		field.String("taxId").Unique().Nillable(),
		field.Float("vatRate").Default(0.16),
		field.String("website").Nillable().Optional(),
		field.Bool("incompleteSetup").Optional().Default(true),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("availableRoles", UserRole.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a company must have at least one role to which its users are assigned
		edge.To("accountingEntries", AccountingEntry.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("customers", Customer.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("documents", CompanyDocument.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("employees", Employee.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("files", File.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("inventory", Inventory.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("inventoryMovements", InventoryMovement.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("invoices", Invoice.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("loans", Loan.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("loan_schedule", LoanSchedule.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("memberSignupTokens", MemberSignupToken.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("products", Product.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("projects", Project.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("payables", Payable.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("receivables", Receivable.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("suppliers", Supplier.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a company can have many tokens
		edge.To("treasuries", Treasury.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("workShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("users", User.Type), // a company can have many users

		// One-to-many relationship: each company can have many daughter companies, but only one parent company
		edge.To("daughterCompanies", Company.Type),
		edge.From("parentCompany", Company.Type).Ref("daughterCompanies").Unique(),
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
