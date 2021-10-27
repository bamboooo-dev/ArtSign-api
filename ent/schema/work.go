package schema

import "entgo.io/ent"

// Work holds the schema definition for the Work entity.
type Work struct {
	ent.Schema
}

// Fields of the Work.
func (Work) Fields() []ent.Field {
	return nil
}

// Edges of the Work.
func (Work) Edges() []ent.Edge {
	return nil
}
