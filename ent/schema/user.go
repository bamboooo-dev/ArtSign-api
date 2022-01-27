package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("username").
			NotEmpty().
			Annotations(
				entgql.OrderField("USERNAME"),
			),
		field.Text("profile"),
		field.String("avatar_url"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("works", Work.Type),
		edge.To("likes", Work.Type),
		edge.To("treasures", Treasure.Type),
		edge.To("portfolios", Portfolio.Type),
		edge.To("comments", Comment.Type),
		edge.To("like_comments", Comment.Type),
		edge.To("followees", User.Type).From("followers"),
	}
}
