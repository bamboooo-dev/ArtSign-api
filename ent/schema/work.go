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
		field.Float("height"),
		field.Float("width"),
		field.String("size_unit").
			NotEmpty(),
		field.Int("year").Positive(),
		field.Int("month").Max(12).Min(1),
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
// TODO: Annotations(entgql.Bind()) を加えて N+1 を解消する
func (Work) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("works").
			Unique(),
		edge.From("tools", Tool.Type).
			Ref("works"),
		edge.From("owner", User.Type).
			Ref("works").
			Unique(),
		edge.From("likers", User.Type).
			Ref("likes"),
		edge.From("treasurers", User.Type).
			Ref("treasures"),
		edge.To("comments", Comment.Type),
		edge.To("images", Image.Type),
	}
}
