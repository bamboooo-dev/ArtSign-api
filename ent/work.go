// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/work"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Work is the model entity for the Work schema.
type Work struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status work.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkQuery when eager-loading is set.
	Edges       WorkEdges `json:"edges"`
	work_parent *int
}

// WorkEdges holds the relations/edges for other nodes in the graph.
type WorkEdges struct {
	// Children holds the value of the children edge.
	Children []*Work `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Work `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEdges) ChildrenOrErr() ([]*Work, error) {
	if e.loadedTypes[0] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEdges) ParentOrErr() (*Work, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: work.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Work) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case work.FieldID, work.FieldPriority:
			values[i] = new(sql.NullInt64)
		case work.FieldTitle, work.FieldDescription, work.FieldImageURL, work.FieldText, work.FieldStatus:
			values[i] = new(sql.NullString)
		case work.FieldUpdatedAt, work.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case work.ForeignKeys[0]: // work_parent
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Work", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Work fields.
func (w *Work) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case work.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case work.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				w.Title = value.String
			}
		case work.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case work.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				w.ImageURL = value.String
			}
		case work.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case work.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				w.Text = value.String
			}
		case work.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case work.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				w.Status = work.Status(value.String)
			}
		case work.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				w.Priority = int(value.Int64)
			}
		case work.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field work_parent", value)
			} else if value.Valid {
				w.work_parent = new(int)
				*w.work_parent = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryChildren queries the "children" edge of the Work entity.
func (w *Work) QueryChildren() *WorkQuery {
	return (&WorkClient{config: w.config}).QueryChildren(w)
}

// QueryParent queries the "parent" edge of the Work entity.
func (w *Work) QueryParent() *WorkQuery {
	return (&WorkClient{config: w.config}).QueryParent(w)
}

// Update returns a builder for updating this Work.
// Note that you need to call Work.Unwrap() before calling this method if this Work
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Work) Update() *WorkUpdateOne {
	return (&WorkClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Work entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Work) Unwrap() *Work {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Work is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Work) String() string {
	var builder strings.Builder
	builder.WriteString("Work(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", title=")
	builder.WriteString(w.Title)
	builder.WriteString(", description=")
	builder.WriteString(w.Description)
	builder.WriteString(", image_url=")
	builder.WriteString(w.ImageURL)
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", text=")
	builder.WriteString(w.Text)
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", w.Status))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", w.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// Works is a parsable slice of Work.
type Works []*Work

func (w Works) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
