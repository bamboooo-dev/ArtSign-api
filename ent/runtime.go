// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/category"
	"artsign/ent/comment"
	"artsign/ent/schema"
	"artsign/ent/user"
	"artsign/ent/work"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	commentMixin := schema.Comment{}.Mixin()
	commentMixinFields0 := commentMixin[0].Fields()
	_ = commentMixinFields0
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreateTime is the schema descriptor for create_time field.
	commentDescCreateTime := commentMixinFields0[0].Descriptor()
	// comment.DefaultCreateTime holds the default value on creation for the create_time field.
	comment.DefaultCreateTime = commentDescCreateTime.Default.(func() time.Time)
	// commentDescUpdateTime is the schema descriptor for update_time field.
	commentDescUpdateTime := commentMixinFields0[1].Descriptor()
	// comment.DefaultUpdateTime holds the default value on creation for the update_time field.
	comment.DefaultUpdateTime = commentDescUpdateTime.Default.(func() time.Time)
	// comment.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	comment.UpdateDefaultUpdateTime = commentDescUpdateTime.UpdateDefault.(func() time.Time)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[0].Descriptor()
	// comment.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	comment.ContentValidator = commentDescContent.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	workFields := schema.Work{}.Fields()
	_ = workFields
	// workDescTitle is the schema descriptor for title field.
	workDescTitle := workFields[0].Descriptor()
	// work.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	work.TitleValidator = workDescTitle.Validators[0].(func(string) error)
	// workDescDescription is the schema descriptor for description field.
	workDescDescription := workFields[1].Descriptor()
	// work.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	work.DescriptionValidator = workDescDescription.Validators[0].(func(string) error)
	// workDescCreatedAt is the schema descriptor for created_at field.
	workDescCreatedAt := workFields[3].Descriptor()
	// work.DefaultCreatedAt holds the default value on creation for the created_at field.
	work.DefaultCreatedAt = workDescCreatedAt.Default.(func() time.Time)
	// workDescUpdatedAt is the schema descriptor for updated_at field.
	workDescUpdatedAt := workFields[4].Descriptor()
	// work.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	work.DefaultUpdatedAt = workDescUpdatedAt.Default.(func() time.Time)
}
