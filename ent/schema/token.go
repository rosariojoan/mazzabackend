package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expiry"),
		field.Enum("category").Values("accountDelete"),
		field.String("token").NotEmpty(),
	}
}

// Edges of the Token: a token can belong to only one company and one user
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("tokens").Unique(),
		edge.From("user", User.Type).Ref("tokens").Unique(),
	}
}

// Enable query and mutation for the Token schema
func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
