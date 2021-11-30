package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Tool struct {
	ent.Schema
}

// Fields of the Tool.
func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("works", Work.Type),
	}
}
