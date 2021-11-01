package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Work holds the schema definition for the Work entity.
type Work struct {
	ent.Schema
}

// Fields of the Work.
func (Work) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Annotations(
				entgql.OrderField("TITLE"),
			),
		field.Text("description").
			NotEmpty(),
		field.String("image_url"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// Edges of the Work.
// TODO: Annotations(entgql.Bind()) を加える
func (Work) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("works").
			Unique(),
	}
}
