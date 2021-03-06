// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/category"
	"artsign/ent/comment"
	"artsign/ent/image"
	"artsign/ent/portfolio"
	"artsign/ent/schema"
	"artsign/ent/tool"
	"artsign/ent/treasure"
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
	commentMixinFields1 := commentMixin[1].Fields()
	_ = commentMixinFields1
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreateTime is the schema descriptor for create_time field.
	commentDescCreateTime := commentMixinFields0[0].Descriptor()
	// comment.DefaultCreateTime holds the default value on creation for the create_time field.
	comment.DefaultCreateTime = commentDescCreateTime.Default.(func() time.Time)
	// commentDescUpdateTime is the schema descriptor for update_time field.
	commentDescUpdateTime := commentMixinFields1[0].Descriptor()
	// comment.DefaultUpdateTime holds the default value on creation for the update_time field.
	comment.DefaultUpdateTime = commentDescUpdateTime.Default.(func() time.Time)
	// comment.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	comment.UpdateDefaultUpdateTime = commentDescUpdateTime.UpdateDefault.(func() time.Time)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[0].Descriptor()
	// comment.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	comment.ContentValidator = commentDescContent.Validators[0].(func(string) error)
	imageMixin := schema.Image{}.Mixin()
	imageMixinFields0 := imageMixin[0].Fields()
	_ = imageMixinFields0
	imageMixinFields1 := imageMixin[1].Fields()
	_ = imageMixinFields1
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescCreateTime is the schema descriptor for create_time field.
	imageDescCreateTime := imageMixinFields0[0].Descriptor()
	// image.DefaultCreateTime holds the default value on creation for the create_time field.
	image.DefaultCreateTime = imageDescCreateTime.Default.(func() time.Time)
	// imageDescUpdateTime is the schema descriptor for update_time field.
	imageDescUpdateTime := imageMixinFields1[0].Descriptor()
	// image.DefaultUpdateTime holds the default value on creation for the update_time field.
	image.DefaultUpdateTime = imageDescUpdateTime.Default.(func() time.Time)
	// image.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	image.UpdateDefaultUpdateTime = imageDescUpdateTime.UpdateDefault.(func() time.Time)
	// imageDescURL is the schema descriptor for url field.
	imageDescURL := imageFields[0].Descriptor()
	// image.URLValidator is a validator for the "url" field. It is called by the builders before save.
	image.URLValidator = imageDescURL.Validators[0].(func(string) error)
	portfolioMixin := schema.Portfolio{}.Mixin()
	portfolioMixinFields0 := portfolioMixin[0].Fields()
	_ = portfolioMixinFields0
	portfolioMixinFields1 := portfolioMixin[1].Fields()
	_ = portfolioMixinFields1
	portfolioFields := schema.Portfolio{}.Fields()
	_ = portfolioFields
	// portfolioDescCreateTime is the schema descriptor for create_time field.
	portfolioDescCreateTime := portfolioMixinFields0[0].Descriptor()
	// portfolio.DefaultCreateTime holds the default value on creation for the create_time field.
	portfolio.DefaultCreateTime = portfolioDescCreateTime.Default.(func() time.Time)
	// portfolioDescUpdateTime is the schema descriptor for update_time field.
	portfolioDescUpdateTime := portfolioMixinFields1[0].Descriptor()
	// portfolio.DefaultUpdateTime holds the default value on creation for the update_time field.
	portfolio.DefaultUpdateTime = portfolioDescUpdateTime.Default.(func() time.Time)
	// portfolio.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	portfolio.UpdateDefaultUpdateTime = portfolioDescUpdateTime.UpdateDefault.(func() time.Time)
	toolFields := schema.Tool{}.Fields()
	_ = toolFields
	// toolDescName is the schema descriptor for name field.
	toolDescName := toolFields[0].Descriptor()
	// tool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tool.NameValidator = toolDescName.Validators[0].(func(string) error)
	treasureMixin := schema.Treasure{}.Mixin()
	treasureMixinFields0 := treasureMixin[0].Fields()
	_ = treasureMixinFields0
	treasureMixinFields1 := treasureMixin[1].Fields()
	_ = treasureMixinFields1
	treasureFields := schema.Treasure{}.Fields()
	_ = treasureFields
	// treasureDescCreateTime is the schema descriptor for create_time field.
	treasureDescCreateTime := treasureMixinFields0[0].Descriptor()
	// treasure.DefaultCreateTime holds the default value on creation for the create_time field.
	treasure.DefaultCreateTime = treasureDescCreateTime.Default.(func() time.Time)
	// treasureDescUpdateTime is the schema descriptor for update_time field.
	treasureDescUpdateTime := treasureMixinFields1[0].Descriptor()
	// treasure.DefaultUpdateTime holds the default value on creation for the update_time field.
	treasure.DefaultUpdateTime = treasureDescUpdateTime.Default.(func() time.Time)
	// treasure.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	treasure.UpdateDefaultUpdateTime = treasureDescUpdateTime.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
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
	// workDescSizeUnit is the schema descriptor for size_unit field.
	workDescSizeUnit := workFields[4].Descriptor()
	// work.SizeUnitValidator is a validator for the "size_unit" field. It is called by the builders before save.
	work.SizeUnitValidator = workDescSizeUnit.Validators[0].(func(string) error)
	// workDescYear is the schema descriptor for year field.
	workDescYear := workFields[5].Descriptor()
	// work.YearValidator is a validator for the "year" field. It is called by the builders before save.
	work.YearValidator = workDescYear.Validators[0].(func(int) error)
	// workDescMonth is the schema descriptor for month field.
	workDescMonth := workFields[6].Descriptor()
	// work.MonthValidator is a validator for the "month" field. It is called by the builders before save.
	work.MonthValidator = func() func(int) error {
		validators := workDescMonth.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(month int) error {
			for _, fn := range fns {
				if err := fn(month); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// workDescCreatedAt is the schema descriptor for created_at field.
	workDescCreatedAt := workFields[7].Descriptor()
	// work.DefaultCreatedAt holds the default value on creation for the created_at field.
	work.DefaultCreatedAt = workDescCreatedAt.Default.(func() time.Time)
	// workDescUpdatedAt is the schema descriptor for updated_at field.
	workDescUpdatedAt := workFields[8].Descriptor()
	// work.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	work.DefaultUpdatedAt = workDescUpdatedAt.Default.(func() time.Time)
}
