// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"artsign/ent/migrate"

	"artsign/ent/category"
	"artsign/ent/comment"
	"artsign/ent/image"
	"artsign/ent/tool"
	"artsign/ent/user"
	"artsign/ent/work"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// Tool is the client for interacting with the Tool builders.
	Tool *ToolClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Work is the client for interacting with the Work builders.
	Work *WorkClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Category = NewCategoryClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.Image = NewImageClient(c.config)
	c.Tool = NewToolClient(c.config)
	c.User = NewUserClient(c.config)
	c.Work = NewWorkClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Category: NewCategoryClient(cfg),
		Comment:  NewCommentClient(cfg),
		Image:    NewImageClient(cfg),
		Tool:     NewToolClient(cfg),
		User:     NewUserClient(cfg),
		Work:     NewWorkClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Category: NewCategoryClient(cfg),
		Comment:  NewCommentClient(cfg),
		Image:    NewImageClient(cfg),
		Tool:     NewToolClient(cfg),
		User:     NewUserClient(cfg),
		Work:     NewWorkClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Category.Use(hooks...)
	c.Comment.Use(hooks...)
	c.Image.Use(hooks...)
	c.Tool.Use(hooks...)
	c.User.Use(hooks...)
	c.Work.Use(hooks...)
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a create builder for Category.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id int) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CategoryClient) DeleteOneID(id int) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id int) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id int) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorks queries the works edge of a Category.
func (c *CategoryClient) QueryWorks(ca *Category) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, category.WorksTable, category.WorksColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Create returns a create builder for Comment.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id int) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommentClient) DeleteOneID(id int) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id int) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id int) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Comment.
func (c *CommentClient) QueryOwner(co *Comment) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.OwnerTable, comment.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWork queries the work edge of a Comment.
func (c *CommentClient) QueryWork(co *Comment) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.WorkTable, comment.WorkColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Comment.
func (c *CommentClient) QueryChildren(co *Comment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, comment.ChildrenTable, comment.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Comment.
func (c *CommentClient) QueryParent(co *Comment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, comment.ParentTable, comment.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikers queries the likers edge of a Comment.
func (c *CommentClient) QueryLikers(co *Comment) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, comment.LikersTable, comment.LikersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Create returns a create builder for Image.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id int) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ImageClient) DeleteOneID(id int) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id int) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id int) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWork queries the work edge of a Image.
func (c *ImageClient) QueryWork(i *Image) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, image.WorkTable, image.WorkColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// ToolClient is a client for the Tool schema.
type ToolClient struct {
	config
}

// NewToolClient returns a client for the Tool from the given config.
func NewToolClient(c config) *ToolClient {
	return &ToolClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tool.Hooks(f(g(h())))`.
func (c *ToolClient) Use(hooks ...Hook) {
	c.hooks.Tool = append(c.hooks.Tool, hooks...)
}

// Create returns a create builder for Tool.
func (c *ToolClient) Create() *ToolCreate {
	mutation := newToolMutation(c.config, OpCreate)
	return &ToolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tool entities.
func (c *ToolClient) CreateBulk(builders ...*ToolCreate) *ToolCreateBulk {
	return &ToolCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tool.
func (c *ToolClient) Update() *ToolUpdate {
	mutation := newToolMutation(c.config, OpUpdate)
	return &ToolUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ToolClient) UpdateOne(t *Tool) *ToolUpdateOne {
	mutation := newToolMutation(c.config, OpUpdateOne, withTool(t))
	return &ToolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ToolClient) UpdateOneID(id int) *ToolUpdateOne {
	mutation := newToolMutation(c.config, OpUpdateOne, withToolID(id))
	return &ToolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tool.
func (c *ToolClient) Delete() *ToolDelete {
	mutation := newToolMutation(c.config, OpDelete)
	return &ToolDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ToolClient) DeleteOne(t *Tool) *ToolDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ToolClient) DeleteOneID(id int) *ToolDeleteOne {
	builder := c.Delete().Where(tool.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ToolDeleteOne{builder}
}

// Query returns a query builder for Tool.
func (c *ToolClient) Query() *ToolQuery {
	return &ToolQuery{
		config: c.config,
	}
}

// Get returns a Tool entity by its id.
func (c *ToolClient) Get(ctx context.Context, id int) (*Tool, error) {
	return c.Query().Where(tool.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ToolClient) GetX(ctx context.Context, id int) *Tool {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorks queries the works edge of a Tool.
func (c *ToolClient) QueryWorks(t *Tool) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tool.Table, tool.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, tool.WorksTable, tool.WorksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ToolClient) Hooks() []Hook {
	return c.hooks.Tool
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorks queries the works edge of a User.
func (c *UserClient) QueryWorks(u *User) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.WorksTable, user.WorksColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikes queries the likes edge of a User.
func (c *UserClient) QueryLikes(u *User) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.LikesTable, user.LikesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTreasures queries the treasures edge of a User.
func (c *UserClient) QueryTreasures(u *User) *WorkQuery {
	query := &WorkQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.TreasuresTable, user.TreasuresPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a User.
func (c *UserClient) QueryComments(u *User) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CommentsTable, user.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikeComments queries the like_comments edge of a User.
func (c *UserClient) QueryLikeComments(u *User) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.LikeCommentsTable, user.LikeCommentsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// WorkClient is a client for the Work schema.
type WorkClient struct {
	config
}

// NewWorkClient returns a client for the Work from the given config.
func NewWorkClient(c config) *WorkClient {
	return &WorkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `work.Hooks(f(g(h())))`.
func (c *WorkClient) Use(hooks ...Hook) {
	c.hooks.Work = append(c.hooks.Work, hooks...)
}

// Create returns a create builder for Work.
func (c *WorkClient) Create() *WorkCreate {
	mutation := newWorkMutation(c.config, OpCreate)
	return &WorkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Work entities.
func (c *WorkClient) CreateBulk(builders ...*WorkCreate) *WorkCreateBulk {
	return &WorkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Work.
func (c *WorkClient) Update() *WorkUpdate {
	mutation := newWorkMutation(c.config, OpUpdate)
	return &WorkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkClient) UpdateOne(w *Work) *WorkUpdateOne {
	mutation := newWorkMutation(c.config, OpUpdateOne, withWork(w))
	return &WorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkClient) UpdateOneID(id int) *WorkUpdateOne {
	mutation := newWorkMutation(c.config, OpUpdateOne, withWorkID(id))
	return &WorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Work.
func (c *WorkClient) Delete() *WorkDelete {
	mutation := newWorkMutation(c.config, OpDelete)
	return &WorkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkClient) DeleteOne(w *Work) *WorkDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkClient) DeleteOneID(id int) *WorkDeleteOne {
	builder := c.Delete().Where(work.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkDeleteOne{builder}
}

// Query returns a query builder for Work.
func (c *WorkClient) Query() *WorkQuery {
	return &WorkQuery{
		config: c.config,
	}
}

// Get returns a Work entity by its id.
func (c *WorkClient) Get(ctx context.Context, id int) (*Work, error) {
	return c.Query().Where(work.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkClient) GetX(ctx context.Context, id int) *Work {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCategory queries the category edge of a Work.
func (c *WorkClient) QueryCategory(w *Work) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, work.CategoryTable, work.CategoryColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTools queries the tools edge of a Work.
func (c *WorkClient) QueryTools(w *Work) *ToolQuery {
	query := &ToolQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(tool.Table, tool.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, work.ToolsTable, work.ToolsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Work.
func (c *WorkClient) QueryOwner(w *Work) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, work.OwnerTable, work.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikers queries the likers edge of a Work.
func (c *WorkClient) QueryLikers(w *Work) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, work.LikersTable, work.LikersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTreasurers queries the treasurers edge of a Work.
func (c *WorkClient) QueryTreasurers(w *Work) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, work.TreasurersTable, work.TreasurersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Work.
func (c *WorkClient) QueryComments(w *Work) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, work.CommentsTable, work.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryImages queries the images edge of a Work.
func (c *WorkClient) QueryImages(w *Work) *ImageQuery {
	query := &ImageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(work.Table, work.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, work.ImagesTable, work.ImagesColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkClient) Hooks() []Hook {
	return c.hooks.Work
}
