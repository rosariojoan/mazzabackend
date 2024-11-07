package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Treasury holds the schema definition for the Treasury entity.
type Treasury struct {
	ent.Schema
}

func (Treasury) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Treasury.
func (Treasury) Fields() []ent.Field {
	return []ent.Field{
		field.Float("balance"),
	}
}

// Edges of the Treasury.
func (Treasury) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("treasuries").Unique(), // a treasury can belong to only one company
	}
}

// Enable query and mutation for the Treasury schema
func (Treasury) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
