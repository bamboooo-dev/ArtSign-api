package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/mixin"
)

// Treasure holds the schema definition for the Treasure entity.
type Treasure struct {
	ent.Schema
}

func (Treasure) Mixin() []ent.Mixin {
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

// Fields of the Treasure.
func (Treasure) Fields() []ent.Field {
	return nil
}

// Edges of the Treasure.
func (Treasure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("treasures").
			Unique().
			Required(),
		edge.From("work", Work.Type).
			Ref("treasures").
			Unique().
			Required(),
	}
}
