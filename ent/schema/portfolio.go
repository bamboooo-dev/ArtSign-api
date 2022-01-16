package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/mixin"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

func (Portfolio) Mixin() []ent.Mixin {
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

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return nil
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("portfolios").
			Unique().
			Required(),
		edge.From("work", Work.Type).
			Ref("portfolios").
			Unique().
			Required(),
	}
}
