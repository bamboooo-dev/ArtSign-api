package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Work holds the schema definition for the Work entity.
type Work struct {
	ent.Schema
}

// Fields of the Work.
func (Work) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			Values("in_progress", "completed").
			Default("in_progress"),
		field.Int("priority").
			Default(0),
	}
}

// Edges of the Work.
func (Work) Edges() []ent.Edge {
	return nil
}
