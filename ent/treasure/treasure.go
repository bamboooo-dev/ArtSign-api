// Code generated by entc, DO NOT EDIT.

package treasure

import (
	"time"
)

const (
	// Label holds the string label denoting the treasure type in the database.
	Label = "treasure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeWork holds the string denoting the work edge name in mutations.
	EdgeWork = "work"
	// Table holds the table name of the treasure in the database.
	Table = "treasures"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "treasures"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_treasures"
	// WorkTable is the table that holds the work relation/edge.
	WorkTable = "treasures"
	// WorkInverseTable is the table name for the Work entity.
	// It exists in this package in order to avoid circular dependency with the "work" package.
	WorkInverseTable = "works"
	// WorkColumn is the table column denoting the work relation/edge.
	WorkColumn = "work_treasures"
)

// Columns holds all SQL columns for treasure fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "treasures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_treasures",
	"work_treasures",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
