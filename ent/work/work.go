// Code generated by entc, DO NOT EDIT.

package work

const (
	// Label holds the string label denoting the work type in the database.
	Label = "work"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the work in the database.
	Table = "works"
)

// Columns holds all SQL columns for work fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
