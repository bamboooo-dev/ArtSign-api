// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProfile holds the string denoting the profile field in the database.
	FieldProfile = "profile"
	// EdgeWorks holds the string denoting the works edge name in mutations.
	EdgeWorks = "works"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// WorksTable is the table that holds the works relation/edge.
	WorksTable = "works"
	// WorksInverseTable is the table name for the Work entity.
	// It exists in this package in order to avoid circular dependency with the "work" package.
	WorksInverseTable = "works"
	// WorksColumn is the table column denoting the works relation/edge.
	WorksColumn = "user_works"
	// LikesTable is the table that holds the likes relation/edge. The primary key declared below.
	LikesTable = "user_likes"
	// LikesInverseTable is the table name for the Work entity.
	// It exists in this package in order to avoid circular dependency with the "work" package.
	LikesInverseTable = "works"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldProfile,
}

var (
	// LikesPrimaryKey and LikesColumn2 are the table columns denoting the
	// primary key for the likes relation (M2M).
	LikesPrimaryKey = []string{"user_id", "work_id"}
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
