package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProjectMilestone holds the schema definition for the ProjectMilestone entity.
type ProjectMilestone struct {
	ent.Schema
}

// Fields of the ProjectMilestone.
func (ProjectMilestone) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Time("due_date"),
	}
}

// Edges of the ProjectMilestone.
func (ProjectMilestone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("milestones").Unique().Required(), // a milestone can belong to only one project
	}
}

// Create indexes
func (ProjectMilestone) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("project").Unique(), // each milestone in a project is unique
	}
}

// Enable query and mutation for the Project schema
func (ProjectMilestone) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
