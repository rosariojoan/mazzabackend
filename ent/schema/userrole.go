package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").Values("superuser", "admin", "accountant", "auditor", "staff").Annotations(entgql.OrderField("ROLES")),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("availableRoles").Unique().Comment("each role must belong to only only company"),
		edge.From("user", User.Type).Ref("assignedRoles").Comment("a role must have at least one user"),
	}
}

// Enable query and mutation for the UserRole schema
func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		// entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
