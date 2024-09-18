package schema

import (
	"time"

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
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deletedAt").Nillable().Optional(),
	}
}
