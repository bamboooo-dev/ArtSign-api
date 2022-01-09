// Code generated by entc, DO NOT EDIT.

package work

import (
	"time"
)

const (
	// Label holds the string label denoting the work type in the database.
	Label = "work"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldSizeUnit holds the string denoting the size_unit field in the database.
	FieldSizeUnit = "size_unit"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldMonth holds the string denoting the month field in the database.
	FieldMonth = "month"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeTools holds the string denoting the tools edge name in mutations.
	EdgeTools = "tools"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeLikers holds the string denoting the likers edge name in mutations.
	EdgeLikers = "likers"
	// EdgeTreasures holds the string denoting the treasures edge name in mutations.
	EdgeTreasures = "treasures"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// Table holds the table name of the work in the database.
	Table = "works"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "works"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_works"
	// ToolsTable is the table that holds the tools relation/edge. The primary key declared below.
	ToolsTable = "tool_works"
	// ToolsInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolsInverseTable = "tools"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "works"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_works"
	// LikersTable is the table that holds the likers relation/edge. The primary key declared below.
	LikersTable = "user_likes"
	// LikersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	LikersInverseTable = "users"
	// TreasuresTable is the table that holds the treasures relation/edge.
	TreasuresTable = "treasures"
	// TreasuresInverseTable is the table name for the Treasure entity.
	// It exists in this package in order to avoid circular dependency with the "treasure" package.
	TreasuresInverseTable = "treasures"
	// TreasuresColumn is the table column denoting the treasures relation/edge.
	TreasuresColumn = "work_treasures"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "work_comments"
	// ImagesTable is the table that holds the images relation/edge.
	ImagesTable = "images"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "work_images"
)

// Columns holds all SQL columns for work fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldHeight,
	FieldWidth,
	FieldSizeUnit,
	FieldYear,
	FieldMonth,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "works"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"category_works",
	"user_works",
}

var (
	// ToolsPrimaryKey and ToolsColumn2 are the table columns denoting the
	// primary key for the tools relation (M2M).
	ToolsPrimaryKey = []string{"tool_id", "work_id"}
	// LikersPrimaryKey and LikersColumn2 are the table columns denoting the
	// primary key for the likers relation (M2M).
	LikersPrimaryKey = []string{"user_id", "work_id"}
)

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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// SizeUnitValidator is a validator for the "size_unit" field. It is called by the builders before save.
	SizeUnitValidator func(string) error
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
	// MonthValidator is a validator for the "month" field. It is called by the builders before save.
	MonthValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
