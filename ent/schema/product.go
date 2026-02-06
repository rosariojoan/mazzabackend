package schema

// import (
// 	"entgo.io/contrib/entgql"
// 	"entgo.io/ent"
// 	"entgo.io/ent/schema"
// 	"entgo.io/ent/schema/edge"
// 	"entgo.io/ent/schema/field"
// )

// // Product holds the schema definition for the Product entity.
// type Product struct {
// 	ent.Schema
// }

// func (Product) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		BaseModelMixin{},
// 	}
// }

// // Fields of the Product.
// func (Product) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Int("stock").Min(0).Default(0),
// 	}
// }

// // Edges of the Product.
// func (Product) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("company", Company.Type).Ref("products").Unique(), // a product can belong to only one company
// 	}
// }

// // Enable query and mutation for the AccountingEntry schema
// func (Product) Annotations() []schema.Annotation {
// 	return []schema.Annotation{
// 		// entgql.RelayConnection(),
// 		// entgql.QueryField(),
// 		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
// 	}
// }
