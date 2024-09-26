package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("accountNumber").Optional(),
		field.Float("balance"),
		field.String("bankName").Optional(),
		field.Enum("currency").Values("mzn"),
		field.String("description").Optional(),
		field.String("iban").Optional(),
		field.Bool("isDefault").Optional().Default(false),
		field.Bool("isMainAccount").Optional().Default(false),
		field.String("name").NotEmpty(),
		field.Enum("category").Values("DEPOSIT", "CASH"),
		field.String("swiftCode").Optional(),
	}
}

// Create indexes
// Company-wise, cash & cash equivalent accounts must have unique names
func (Treasury) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("company").Unique(),
	}
}

// Edges of the Treasury.
func (Treasury) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("treasuries").Unique(),                            // a treasury can belong to only one company
		edge.To("cashMovements", CashMovement.Type).Annotations(entsql.OnDelete(entsql.Cascade)), // a cash & cash equivalent account can have many cash movements
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
