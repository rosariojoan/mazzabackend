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

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
		field.Time("plannedStartDate").StructTag("plannedStartDate").Annotations(entgql.OrderField("PLANNED_START_DATE")),
		field.Time("actualStartDate").StructTag("actualStartDate").Nillable().Optional().Annotations(entgql.OrderField("ACTUAL_START_DATE")),
		field.Time("plannedEndDate").StructTag("plannedEndDate").Annotations(entgql.OrderField("PLANNED_END_DATE")),
		field.Time("actualEndDate").StructTag("actualEndDate").Nillable().Optional().Annotations(entgql.OrderField("ACTUAL_END_DATE")),
		field.Float("progress").Min(0).Max(1).Default(0).Annotations(entgql.OrderField("PROGRESS")),
		field.Enum("status").Values("pending", "inProgress", "completed", "delayed").Default("pending").Annotations(entgql.OrderField("STATUS")),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("projects").Unique(), // a project can belong to only one company
		edge.From("createdBy", User.Type).Ref("createdProjects").Unique(),
		edge.From("leader", User.Type).Ref("leaderedProjects").Unique(),
		edge.To("tasks", ProjectTask.Type).Annotations(entsql.OnDelete(entsql.Cascade)),           // a project can have many tasks
		edge.To("milestones", ProjectMilestone.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a project can have many milestones
	}
}

// Create indexes
func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(),
	}
}

// Enable query and mutation for the Project schema
func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
