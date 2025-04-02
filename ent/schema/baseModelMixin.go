package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// A mixin can be shared among package schemas
type BaseModelMixin struct {
	mixin.Schema
}

func (BaseModelMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(
			entgql.OrderField("CREATED_AT"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(
			entgql.OrderField("UPDATED_AT"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("deletedAt").Nillable().Optional().Annotations(entgql.Skip(
			entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput,
		)),
	}
}
