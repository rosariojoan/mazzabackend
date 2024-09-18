package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Worktask holds the schema definition for the Worktask entity.
type Worktask struct {
	ent.Schema
}

func (Worktask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Worktask.
func (Worktask) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional(),
		field.Enum("status").Values("DRAFT", "ASSIGNED", "DONE"),
		field.JSON("subtasks", []string{}).Optional().Comment(
			"The content of this field should be a serialized slice of JSONs.\n" +
				"Each entry should hve the following fields: " +
				"description: string, order: number, completed: bool, completedBy: {id: number, name: string}, completedAt: time",
		),
		field.String("title"),
		field.Time("startTime"),
		field.Time("endTime").Optional(),
	}
}

// Edges of the Worktask.
func (Worktask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("workTasks").Unique(),
		edge.From("createdBy", User.Type).Ref("createdTasks").Unique(),
		edge.From("assignedTo", Employee.Type).Ref("assignedTasks"),
		// edge.To("subtasks", Worktask.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a task can have many subtasks
		// edge.From("workTask", Worktask.Type).Ref("subtasks").Unique(),                   // a subtask can only belong to one task
		edge.To("workShifts", Workshift.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("workTags", Worktag.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

// Enable query and mutation for the CompanyClient schema
func (Worktask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
