package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Time("birthdate").Nillable().Optional(),
		field.Enum("gender").Values("male", "female").Annotations(entgql.OrderField("GENDER")),
		field.String("position"),
		field.String("department").Optional().Default("geral"),
		field.String("email").Nillable().Optional(),
		field.String("phone").Nillable().Optional(),
		field.String("avatar").Nillable().Optional(),
		field.Time("hire_date"),
		field.Int("monthly_salary").Min(0).Default(0),
		field.Enum("status").Values("ACTIVE", "ON_LEAVE").Default("ACTIVE").Optional().Annotations(entgql.OrderField("STATUS")),
		field.Float("performace_score").Min(0).Max(5).Default(0).Optional(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("employees").Unique(),
		edge.From("user", User.Type).Ref("employee").Unique(),
		edge.To("subordinates", Employee.Type).Annotations(entsql.OnDelete(entsql.SetNull)), // an employee can be a leader of many employees
		edge.From("leader", Employee.Type).Ref("subordinates").Unique(),                     // an employee can have only one leader
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
