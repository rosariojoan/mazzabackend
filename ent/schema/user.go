package schema

import (
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
		field.String("fcmToken").Nillable().Optional(),
		field.String("email").Nillable().Optional().Unique(),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.String("password").Sensitive(),
		field.String("username").Unique().Annotations(entgql.OrderField("USERNAME")),
		field.Bool("disabled").Nillable().Optional().Default(false),
		field.Bool("notVerified").Nillable().Optional().Default(false),
	}
}

// Edges of the User: a user can belong to many companies
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accountingEntries", AccountingEntry.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.From("company", Company.Type).Ref("users").Required(),
		edge.To("assignedRoles", UserRole.Type).Annotations(entsql.OnDelete(entsql.Cascade)).Comment("a user should be assigned to at least one role in the company"),
		edge.To("createdTasks", Worktask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("Tasks created by this user"),
		edge.To("employee", Employee.Type).Annotations(entsql.OnDelete(entsql.Cascade)).Unique(),
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a user can have more than one token. E.g. of token: pwd reset token
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
