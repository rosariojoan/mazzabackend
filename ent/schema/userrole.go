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

var userRoles = []string{
	"SUPERUSER",
	"ADMIN",
	"ACCOUNTANT",
	"AUDITOR",
	"STAFF",
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").Values(userRoles...).Annotations(entgql.OrderField("ROLES")),
		field.String("notes").MaxLen(255).Nillable().Comment("Description about this role"),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("available_roles").Unique().Comment("each role must belong to only only company"),
		edge.From("user", User.Type).Ref("assigned_roles").Unique().Comment("a role must belong to only one user"),
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
