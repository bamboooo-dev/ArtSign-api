// Code generated by entc, DO NOT EDIT.

package tool

const (
	// Label holds the string label denoting the tool type in the database.
	Label = "tool"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeWorks holds the string denoting the works edge name in mutations.
	EdgeWorks = "works"
	// Table holds the table name of the tool in the database.
	Table = "tools"
	// WorksTable is the table that holds the works relation/edge. The primary key declared below.
	WorksTable = "tool_works"
	// WorksInverseTable is the table name for the Work entity.
	// It exists in this package in order to avoid circular dependency with the "work" package.
	WorksInverseTable = "works"
)

// Columns holds all SQL columns for tool fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// WorksPrimaryKey and WorksColumn2 are the table columns denoting the
	// primary key for the works relation (M2M).
	WorksPrimaryKey = []string{"tool_id", "work_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
