// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/category"
	"artsign/ent/comment"
	"artsign/ent/image"
	"artsign/ent/user"
	"artsign/ent/work"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkCreate is the builder for creating a Work entity.
type WorkCreate struct {
	config
	mutation *WorkMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (wc *WorkCreate) SetTitle(s string) *WorkCreate {
	wc.mutation.SetTitle(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkCreate) SetDescription(s string) *WorkCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkCreate) SetCreatedAt(t time.Time) *WorkCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkCreate) SetNillableCreatedAt(t *time.Time) *WorkCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorkCreate) SetUpdatedAt(t time.Time) *WorkCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorkCreate) SetNillableUpdatedAt(t *time.Time) *WorkCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (wc *WorkCreate) SetCategoryID(id int) *WorkCreate {
	wc.mutation.SetCategoryID(id)
	return wc
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (wc *WorkCreate) SetNillableCategoryID(id *int) *WorkCreate {
	if id != nil {
		wc = wc.SetCategoryID(*id)
	}
	return wc
}

// SetCategory sets the "category" edge to the Category entity.
func (wc *WorkCreate) SetCategory(c *Category) *WorkCreate {
	return wc.SetCategoryID(c.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (wc *WorkCreate) SetOwnerID(id int) *WorkCreate {
	wc.mutation.SetOwnerID(id)
	return wc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (wc *WorkCreate) SetNillableOwnerID(id *int) *WorkCreate {
	if id != nil {
		wc = wc.SetOwnerID(*id)
	}
	return wc
}

// SetOwner sets the "owner" edge to the User entity.
func (wc *WorkCreate) SetOwner(u *User) *WorkCreate {
	return wc.SetOwnerID(u.ID)
}

// AddLikerIDs adds the "likers" edge to the User entity by IDs.
func (wc *WorkCreate) AddLikerIDs(ids ...int) *WorkCreate {
	wc.mutation.AddLikerIDs(ids...)
	return wc
}

// AddLikers adds the "likers" edges to the User entity.
func (wc *WorkCreate) AddLikers(u ...*User) *WorkCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wc.AddLikerIDs(ids...)
}

// AddTreasurerIDs adds the "treasurers" edge to the User entity by IDs.
func (wc *WorkCreate) AddTreasurerIDs(ids ...int) *WorkCreate {
	wc.mutation.AddTreasurerIDs(ids...)
	return wc
}

// AddTreasurers adds the "treasurers" edges to the User entity.
func (wc *WorkCreate) AddTreasurers(u ...*User) *WorkCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wc.AddTreasurerIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (wc *WorkCreate) AddCommentIDs(ids ...int) *WorkCreate {
	wc.mutation.AddCommentIDs(ids...)
	return wc
}

// AddComments adds the "comments" edges to the Comment entity.
func (wc *WorkCreate) AddComments(c ...*Comment) *WorkCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wc.AddCommentIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (wc *WorkCreate) AddImageIDs(ids ...int) *WorkCreate {
	wc.mutation.AddImageIDs(ids...)
	return wc
}

// AddImages adds the "images" edges to the Image entity.
func (wc *WorkCreate) AddImages(i ...*Image) *WorkCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wc.AddImageIDs(ids...)
}

// Mutation returns the WorkMutation object of the builder.
func (wc *WorkCreate) Mutation() *WorkMutation {
	return wc.mutation
}

// Save creates the Work in the database.
func (wc *WorkCreate) Save(ctx context.Context) (*Work, error) {
	var (
		err  error
		node *Work
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkCreate) SaveX(ctx context.Context) *Work {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := work.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := work.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkCreate) check() error {
	if _, ok := wc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := wc.mutation.Title(); ok {
		if err := work.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if v, ok := wc.mutation.Description(); ok {
		if err := work.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "description": %w`, err)}
		}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (wc *WorkCreate) sqlSave(ctx context.Context) (*Work, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wc *WorkCreate) createSpec() (*Work, *sqlgraph.CreateSpec) {
	var (
		_node = &Work{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: work.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: work.FieldID,
			},
		}
	)
	if value, ok := wc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := wc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   work.CategoryTable,
			Columns: []string{work.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_works = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   work.OwnerTable,
			Columns: []string{work.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_works = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.LikersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   work.LikersTable,
			Columns: work.LikersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TreasurersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   work.TreasurersTable,
			Columns: work.TreasurersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.CommentsTable,
			Columns: []string{work.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.ImagesTable,
			Columns: []string{work.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkCreateBulk is the builder for creating many Work entities in bulk.
type WorkCreateBulk struct {
	config
	builders []*WorkCreate
}

// Save creates the Work entities in the database.
func (wcb *WorkCreateBulk) Save(ctx context.Context) ([]*Work, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Work, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkCreateBulk) SaveX(ctx context.Context) []*Work {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
