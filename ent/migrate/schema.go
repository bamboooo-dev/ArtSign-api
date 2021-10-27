// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// WorksColumns holds the columns for the "works" table.
	WorksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// WorksTable holds the schema information for the "works" table.
	WorksTable = &schema.Table{
		Name:       "works",
		Columns:    WorksColumns,
		PrimaryKey: []*schema.Column{WorksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		WorksTable,
	}
)

func init() {
}
