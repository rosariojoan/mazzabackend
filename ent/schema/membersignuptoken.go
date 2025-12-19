package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MemberSignupToken holds the schema definition for the MemberSignupToken entity.
type MemberSignupToken struct {
	ent.Schema
}

func (MemberSignupToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the MemberSignupToken.
func (MemberSignupToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").Nillable().Optional(),
		field.String("token").MinLen(6).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("avatar").Nillable().Comment("Must be an emoji"),
		field.Enum("role").Values(userRoles...),
		field.String("note").Comment("A description about the user and its role of this user"),
		field.Int("number_accessed").Default(0).NonNegative().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("expires_at").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Bool("already_used").Default(false).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the MemberSignupToken.
func (MemberSignupToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("member_signup_tokens").Unique().Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)), // one company can own multiple tokens
		edge.From("created_by", User.Type).Ref("created_member_signup_tokens").Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)), // one user can create multiple tokens
	}
}

// Enable query and mutation for the MemberSignupToken schema
func (MemberSignupToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
