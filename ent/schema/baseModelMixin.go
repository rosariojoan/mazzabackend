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
		field.Time("created_at").Default(time.Now).Immutable().Annotations(
			entgql.OrderField("createdAt"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(
			entgql.OrderField("updatedAt"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("deleted_at").Nillable().Optional().Annotations(entgql.Skip(
			entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput,
		)),
	}
}
