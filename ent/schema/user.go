package schema

import (
	"net/mail"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("device").Nillable().Optional().Comment("Current user device info"),
		field.Bool("is_demo_user").Default(false).Nillable().Optional().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("firebase_UID").NotEmpty().Unique().Sensitive(),
		field.String("fcm_token").Nillable().Optional().Sensitive(),
		field.String("expo_push_token").Nillable().Optional().Sensitive(),
		field.String("email").NotEmpty().Unique().Immutable().Validate(func(s string) error {
			_, err := mail.ParseAddress(s)
			return err
		}),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.String("address").Nillable().Optional(),
		field.String("avatar").Nillable().Optional(),
		field.String("photo_URL").Nillable().Optional(),
		field.String("department").Nillable().Optional(),
		field.String("phone").Nillable().Optional(),
		field.Time("birthdate").Nillable().Optional(),
		field.Time("last_login").Nillable().Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)).
			Comment("It can be the last time the user opened the app and synced with the backend."),
		field.Enum("gender").Values("male", "female"),
		field.Bool("active").Default(false).Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		// field.String("password").Sensitive(),
		// field.String("username").Unique().Annotations(entgql.OrderField("USERNAME")),
	}
}

// Edges of the User: a user can belong to many companies
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounting_entries", AccountingEntry.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.From("company", Company.Type).Ref("users").Required().Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.To("assigned_roles", UserRole.Type).Annotations(entsql.OnDelete(entsql.Cascade)).Comment("a user should be assigned to only one role in the company"),
		edge.To("subordinates", User.Type).Annotations(entsql.OnDelete(entsql.SetNull)), // an employee can be a leader of many employees
		edge.From("leader", User.Type).Ref("subordinates").Unique(),                     // an employee can have only one leader
		edge.To("created_member_signup_tokens", MemberSignupToken.Type).Annotations(entsql.OnDelete(entsql.SetNull)),

		edge.To("employee", Employee.Type).Annotations(entsql.OnDelete(entsql.Cascade)).Unique(),
		edge.To("issued_invoices", Invoice.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("created_projects", Project.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("Represents the projects created by the user"),
		edge.To("leadered_projects", Project.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("Represents the projects leadered or supervised by the user"),
		edge.To("assigned_project_tasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("These are the project tasks assigned to the user and he is responsible for them"),
		edge.To("participated_project_tasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("These are the project tasks in which the user is a member. E.g. a meeting"),
		edge.To("created_tasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Immutable().Comment("Represents the tasks created by a user"),
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a user can have more than one token. E.g. of token: pwd reset token
		edge.To("approved_work_shifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("work_shifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("uploaded_documents", CompanyDocument.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("approved_documents", CompanyDocument.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

// Enable query and mutation for the User schema
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// // Create indexes
// func (User) Indexes() []ent.Index {
// 	return []ent.Index{
// 		index.Edges("company").Edges("employee").Unique(), // a user should be associated with only one employee per company
// 	}
// }
