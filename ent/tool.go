// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/tool"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Tool is the model entity for the Tool schema.
type Tool struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ToolQuery when eager-loading is set.
	Edges ToolEdges `json:"edges"`
}

// ToolEdges holds the relations/edges for other nodes in the graph.
type ToolEdges struct {
	// Works holds the value of the works edge.
	Works []*Work `json:"works,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorksOrErr returns the Works value or an error if the edge
// was not loaded in eager-loading.
func (e ToolEdges) WorksOrErr() ([]*Work, error) {
	if e.loadedTypes[0] {
		return e.Works, nil
	}
	return nil, &NotLoadedError{edge: "works"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tool) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			values[i] = new(sql.NullInt64)
		case tool.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tool", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tool fields.
func (t *Tool) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		}
	}
	return nil
}

// QueryWorks queries the "works" edge of the Tool entity.
func (t *Tool) QueryWorks() *WorkQuery {
	return (&ToolClient{config: t.config}).QueryWorks(t)
}

// Update returns a builder for updating this Tool.
// Note that you need to call Tool.Unwrap() before calling this method if this Tool
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tool) Update() *ToolUpdateOne {
	return (&ToolClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tool) Unwrap() *Tool {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tool is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tool) String() string {
	var builder strings.Builder
	builder.WriteString("Tool(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Tools is a parsable slice of Tool.
type Tools []*Tool

func (t Tools) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}