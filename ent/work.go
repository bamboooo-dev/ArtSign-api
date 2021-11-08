// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/category"
	"artsign/ent/user"
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
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkQuery when eager-loading is set.
	Edges          WorkEdges `json:"edges"`
	category_works *int
	user_works     *int
}

// WorkEdges holds the relations/edges for other nodes in the graph.
type WorkEdges struct {
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Likers holds the value of the likers edge.
	Likers []*User `json:"likers,omitempty"`
	// Treasurers holds the value of the treasurers edge.
	Treasurers []*User `json:"treasurers,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// The edge category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// LikersOrErr returns the Likers value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEdges) LikersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Likers, nil
	}
	return nil, &NotLoadedError{edge: "likers"}
}

// TreasurersOrErr returns the Treasurers value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEdges) TreasurersOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Treasurers, nil
	}
	return nil, &NotLoadedError{edge: "treasurers"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[4] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Work) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case work.FieldID:
			values[i] = new(sql.NullInt64)
		case work.FieldTitle, work.FieldDescription, work.FieldImageURL:
			values[i] = new(sql.NullString)
		case work.FieldCreatedAt, work.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case work.ForeignKeys[0]: // category_works
			values[i] = new(sql.NullInt64)
		case work.ForeignKeys[1]: // user_works
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
		case work.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case work.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case work.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_works", value)
			} else if value.Valid {
				w.category_works = new(int)
				*w.category_works = int(value.Int64)
			}
		case work.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_works", value)
			} else if value.Valid {
				w.user_works = new(int)
				*w.user_works = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCategory queries the "category" edge of the Work entity.
func (w *Work) QueryCategory() *CategoryQuery {
	return (&WorkClient{config: w.config}).QueryCategory(w)
}

// QueryOwner queries the "owner" edge of the Work entity.
func (w *Work) QueryOwner() *UserQuery {
	return (&WorkClient{config: w.config}).QueryOwner(w)
}

// QueryLikers queries the "likers" edge of the Work entity.
func (w *Work) QueryLikers() *UserQuery {
	return (&WorkClient{config: w.config}).QueryLikers(w)
}

// QueryTreasurers queries the "treasurers" edge of the Work entity.
func (w *Work) QueryTreasurers() *UserQuery {
	return (&WorkClient{config: w.config}).QueryTreasurers(w)
}

// QueryComments queries the "comments" edge of the Work entity.
func (w *Work) QueryComments() *CommentQuery {
	return (&WorkClient{config: w.config}).QueryComments(w)
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
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
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
