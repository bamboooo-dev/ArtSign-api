// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/portfolio"
	"artsign/ent/user"
	"artsign/ent/work"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Portfolio is the model entity for the Portfolio schema.
type Portfolio struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PortfolioQuery when eager-loading is set.
	Edges           PortfolioEdges `json:"edges"`
	user_portfolios *int
	work_portfolios *int
}

// PortfolioEdges holds the relations/edges for other nodes in the graph.
type PortfolioEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Work holds the value of the work edge.
	Work *Work `json:"work,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PortfolioEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// WorkOrErr returns the Work value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PortfolioEdges) WorkOrErr() (*Work, error) {
	if e.loadedTypes[1] {
		if e.Work == nil {
			// The edge work was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: work.Label}
		}
		return e.Work, nil
	}
	return nil, &NotLoadedError{edge: "work"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Portfolio) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case portfolio.FieldID:
			values[i] = new(sql.NullInt64)
		case portfolio.FieldCreateTime, portfolio.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case portfolio.ForeignKeys[0]: // user_portfolios
			values[i] = new(sql.NullInt64)
		case portfolio.ForeignKeys[1]: // work_portfolios
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Portfolio", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Portfolio fields.
func (po *Portfolio) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case portfolio.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case portfolio.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				po.CreateTime = value.Time
			}
		case portfolio.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				po.UpdateTime = value.Time
			}
		case portfolio.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_portfolios", value)
			} else if value.Valid {
				po.user_portfolios = new(int)
				*po.user_portfolios = int(value.Int64)
			}
		case portfolio.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field work_portfolios", value)
			} else if value.Valid {
				po.work_portfolios = new(int)
				*po.work_portfolios = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Portfolio entity.
func (po *Portfolio) QueryOwner() *UserQuery {
	return (&PortfolioClient{config: po.config}).QueryOwner(po)
}

// QueryWork queries the "work" edge of the Portfolio entity.
func (po *Portfolio) QueryWork() *WorkQuery {
	return (&PortfolioClient{config: po.config}).QueryWork(po)
}

// Update returns a builder for updating this Portfolio.
// Note that you need to call Portfolio.Unwrap() before calling this method if this Portfolio
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Portfolio) Update() *PortfolioUpdateOne {
	return (&PortfolioClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Portfolio entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Portfolio) Unwrap() *Portfolio {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Portfolio is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Portfolio) String() string {
	var builder strings.Builder
	builder.WriteString("Portfolio(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(po.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(po.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Portfolios is a parsable slice of Portfolio.
type Portfolios []*Portfolio

func (po Portfolios) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
