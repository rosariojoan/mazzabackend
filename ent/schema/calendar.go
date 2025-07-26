package schema

import "entgo.io/ent"

// Calendar holds the schema definition for the Calendar entity.
type Calendar struct {
	ent.Schema
}

// Fields of the Calendar.
func (Calendar) Fields() []ent.Field {
	return nil
}

// Edges of the Calendar.
func (Calendar) Edges() []ent.Edge {
	return nil
}
