package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Worktag holds the schema definition for the Worktag entity.
type Worktag struct {
	ent.Schema
}

func (Worktag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Worktag.
func (Worktag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("color"),
	}
}

// Edges of the Worktag.
func (Worktag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("workTags").Unique(),
		edge.From("workTasks", Worktask.Type).Ref("workTags"),
	}
}

// Enable query and mutation for the Worktag schema
func (Worktag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Create indexes
func (Worktag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(), // work tags should be unique accross companies
	}
}
