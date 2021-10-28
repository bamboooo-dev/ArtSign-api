// Code generated by entc, DO NOT EDIT.

package ent

import (
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

// SetText sets the "text" field.
func (wc *WorkCreate) SetText(s string) *WorkCreate {
	wc.mutation.SetText(s)
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

// SetStatus sets the "status" field.
func (wc *WorkCreate) SetStatus(w work.Status) *WorkCreate {
	wc.mutation.SetStatus(w)
	return wc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wc *WorkCreate) SetNillableStatus(w *work.Status) *WorkCreate {
	if w != nil {
		wc.SetStatus(*w)
	}
	return wc
}

// SetPriority sets the "priority" field.
func (wc *WorkCreate) SetPriority(i int) *WorkCreate {
	wc.mutation.SetPriority(i)
	return wc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (wc *WorkCreate) SetNillablePriority(i *int) *WorkCreate {
	if i != nil {
		wc.SetPriority(*i)
	}
	return wc
}

// AddChildIDs adds the "children" edge to the Work entity by IDs.
func (wc *WorkCreate) AddChildIDs(ids ...int) *WorkCreate {
	wc.mutation.AddChildIDs(ids...)
	return wc
}

// AddChildren adds the "children" edges to the Work entity.
func (wc *WorkCreate) AddChildren(w ...*Work) *WorkCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Work entity by ID.
func (wc *WorkCreate) SetParentID(id int) *WorkCreate {
	wc.mutation.SetParentID(id)
	return wc
}

// SetNillableParentID sets the "parent" edge to the Work entity by ID if the given value is not nil.
func (wc *WorkCreate) SetNillableParentID(id *int) *WorkCreate {
	if id != nil {
		wc = wc.SetParentID(*id)
	}
	return wc
}

// SetParent sets the "parent" edge to the Work entity.
func (wc *WorkCreate) SetParent(w *Work) *WorkCreate {
	return wc.SetParentID(w.ID)
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
	if _, ok := wc.mutation.Status(); !ok {
		v := work.DefaultStatus
		wc.mutation.SetStatus(v)
	}
	if _, ok := wc.mutation.Priority(); !ok {
		v := work.DefaultPriority
		wc.mutation.SetPriority(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkCreate) check() error {
	if _, ok := wc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "text"`)}
	}
	if v, ok := wc.mutation.Text(); ok {
		if err := work.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "text": %w`, err)}
		}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := wc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := wc.mutation.Status(); ok {
		if err := work.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "priority"`)}
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
	if value, ok := wc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldText,
		})
		_node.Text = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: work.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := wc.mutation.Priority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
		_node.Priority = value
	}
	if nodes := wc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   work.ChildrenTable,
			Columns: []string{work.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: work.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   work.ParentTable,
			Columns: []string{work.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: work.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.work_parent = &nodes[0]
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
