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
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "work_parent", Type: field.TypeInt, Nullable: true},
	}
	// WorksTable holds the schema information for the "works" table.
	WorksTable = &schema.Table{
		Name:       "works",
		Columns:    WorksColumns,
		PrimaryKey: []*schema.Column{WorksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "works_works_parent",
				Columns:    []*schema.Column{WorksColumns[5]},
				RefColumns: []*schema.Column{WorksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		WorksTable,
	}
)

func init() {
	WorksTable.ForeignKeys[0].RefTable = WorksTable
}
