package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

func (Image) Mixin() []ent.Mixin {
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

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			NotEmpty(),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("work", Work.Type).
			Ref("images").
			Unique().
			Required(),
	}
}
