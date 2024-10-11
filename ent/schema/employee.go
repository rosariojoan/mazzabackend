package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

func (Employee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("gender").Values("male", "female"),
		field.String("position").Nillable().Optional(),
		field.String("email").Nillable().Optional(),
		field.String("phone"),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("employees").Unique(),
		edge.From("user", User.Type).Ref("employee").Unique(),
		// edge.To("subordinates", Employee.Type).Annotations(entsql.OnDelete(entsql.SetNull)), // an employee can be a leader of many employees
		// edge.From("leader", Employee.Type).Ref("subordinates").Unique(),                     // an employee can have only one leader
		// edge.To("workShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		// edge.To("approvedWorkShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		// edge.To("assignedTasks", Worktask.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

// Enable query and mutation for the CompanyClient schema
func (Employee) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
