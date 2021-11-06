package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AnnotateFields(
			mixin.CreateTime{},
			entgql.OrderField("CREATE_TIME"),
		),
		mixin.AnnotateFields(
			mixin.UpdateTime{},
			entgql.OrderField("UPDATE_TIME"),
		),
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").
			NotEmpty(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("comments").
			Unique().
			Required(),
		edge.From("work", Work.Type).
			Ref("comments").
			Unique().
			Required(),
		edge.To("parent", Comment.Type).
			Unique().
			From("children"),
	}
}
