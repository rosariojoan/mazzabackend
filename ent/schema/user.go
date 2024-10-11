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
		field.String("fcmToken").Nillable().Optional().Sensitive(),
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
		edge.To("subordinates", User.Type).Annotations(entsql.OnDelete(entsql.SetNull)), // an employee can be a leader of many employees
		edge.From("leader", User.Type).Ref("subordinates").Unique(),                     // an employee can have only one leader
		// edge.To("createdTasks", Worktask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("Tasks created by this user"),
		edge.To("employee", Employee.Type).Annotations(entsql.OnDelete(entsql.Cascade)).Unique(),
		edge.To("createdProjects", Project.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("Represent the projects created by the user"),
		edge.To("leaderedProjects", Project.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("Represent the projects leadered or supervised by the user"),
		edge.To("assignedProjectTasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("These are the project tasks assigned to the user and he is responsible for them"),
		edge.To("participatedProjectTasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Comment("These are the project tasks in which the user is a member. E.g. a meeting"),
		edge.To("createdTasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.SetNull)).Immutable().Comment("Represents the tasks created by a user"),
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a user can have more than one token. E.g. of token: pwd reset token
		edge.To("approvedWorkShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("workShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
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
