package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Workshift holds the schema definition for the Workshift entity.
type Workshift struct {
	ent.Schema
}

type Location struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}

func (Workshift) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModelMixin{},
	}
}

// Fields of the Workshift.
func (Workshift) Fields() []ent.Field {
	return []ent.Field{
		field.Time("approvedAt").Nillable().Optional().Comment("time that this shift was approved by the supervisor"),
		field.Time("clockIn").Default(time.Now),
		field.Time("clockOut").Nillable().Optional(),
		field.String("clockInLocation").Comment("it expects a serialized json like: {latitude: float, longitude: float, description: string}"),
		field.String("clockOutLocation").Optional().Comment("it expects a serialized json like: {latitude: float, longitude: float, description: string}"),
		field.String("description").Optional(),
		field.String("note").Optional().Comment("this is only used when the current item is a shift edit request"),
		field.Enum("status").Values("APPROVED", "PENDING").Default("PENDING"),
	}
}

// Edges of the Workshift.
func (Workshift) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("workShifts").Unique(),
		edge.From("employee", Employee.Type).Ref("workShifts").Unique(),
		edge.From("approvedBy", Employee.Type).Ref("approvedWorkShifts").Unique(),
		edge.From("workTask", Worktask.Type).Ref("workShifts").Unique(),
		edge.To("editRequest", Workshift.Type).Unique(),
		edge.From("workShift", Workshift.Type).Ref("editRequest").Unique(),
	}
}

// Enable query and mutation for the CompanyClient schema
func (Workshift) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
