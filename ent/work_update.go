// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/predicate"
	"artsign/ent/work"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkUpdate is the builder for updating Work entities.
type WorkUpdate struct {
	config
	hooks    []Hook
	mutation *WorkMutation
}

// Where appends a list predicates to the WorkUpdate builder.
func (wu *WorkUpdate) Where(ps ...predicate.Work) *WorkUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetText sets the "text" field.
func (wu *WorkUpdate) SetText(s string) *WorkUpdate {
	wu.mutation.SetText(s)
	return wu
}

// SetStatus sets the "status" field.
func (wu *WorkUpdate) SetStatus(w work.Status) *WorkUpdate {
	wu.mutation.SetStatus(w)
	return wu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wu *WorkUpdate) SetNillableStatus(w *work.Status) *WorkUpdate {
	if w != nil {
		wu.SetStatus(*w)
	}
	return wu
}

// SetPriority sets the "priority" field.
func (wu *WorkUpdate) SetPriority(i int) *WorkUpdate {
	wu.mutation.ResetPriority()
	wu.mutation.SetPriority(i)
	return wu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (wu *WorkUpdate) SetNillablePriority(i *int) *WorkUpdate {
	if i != nil {
		wu.SetPriority(*i)
	}
	return wu
}

// AddPriority adds i to the "priority" field.
func (wu *WorkUpdate) AddPriority(i int) *WorkUpdate {
	wu.mutation.AddPriority(i)
	return wu
}

// AddChildIDs adds the "children" edge to the Work entity by IDs.
func (wu *WorkUpdate) AddChildIDs(ids ...int) *WorkUpdate {
	wu.mutation.AddChildIDs(ids...)
	return wu
}

// AddChildren adds the "children" edges to the Work entity.
func (wu *WorkUpdate) AddChildren(w ...*Work) *WorkUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Work entity by ID.
func (wu *WorkUpdate) SetParentID(id int) *WorkUpdate {
	wu.mutation.SetParentID(id)
	return wu
}

// SetNillableParentID sets the "parent" edge to the Work entity by ID if the given value is not nil.
func (wu *WorkUpdate) SetNillableParentID(id *int) *WorkUpdate {
	if id != nil {
		wu = wu.SetParentID(*id)
	}
	return wu
}

// SetParent sets the "parent" edge to the Work entity.
func (wu *WorkUpdate) SetParent(w *Work) *WorkUpdate {
	return wu.SetParentID(w.ID)
}

// Mutation returns the WorkMutation object of the builder.
func (wu *WorkUpdate) Mutation() *WorkMutation {
	return wu.mutation
}

// ClearChildren clears all "children" edges to the Work entity.
func (wu *WorkUpdate) ClearChildren() *WorkUpdate {
	wu.mutation.ClearChildren()
	return wu
}

// RemoveChildIDs removes the "children" edge to Work entities by IDs.
func (wu *WorkUpdate) RemoveChildIDs(ids ...int) *WorkUpdate {
	wu.mutation.RemoveChildIDs(ids...)
	return wu
}

// RemoveChildren removes "children" edges to Work entities.
func (wu *WorkUpdate) RemoveChildren(w ...*Work) *WorkUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveChildIDs(ids...)
}

// ClearParent clears the "parent" edge to the Work entity.
func (wu *WorkUpdate) ClearParent() *WorkUpdate {
	wu.mutation.ClearParent()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkUpdate) check() error {
	if v, ok := wu.mutation.Text(); ok {
		if err := work.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf("ent: validator failed for field \"text\": %w", err)}
		}
	}
	if v, ok := wu.mutation.Status(); ok {
		if err := work.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (wu *WorkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   work.Table,
			Columns: work.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: work.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldText,
		})
	}
	if value, ok := wu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: work.FieldStatus,
		})
	}
	if value, ok := wu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if value, ok := wu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if wu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !wu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{work.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkUpdateOne is the builder for updating a single Work entity.
type WorkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkMutation
}

// SetText sets the "text" field.
func (wuo *WorkUpdateOne) SetText(s string) *WorkUpdateOne {
	wuo.mutation.SetText(s)
	return wuo
}

// SetStatus sets the "status" field.
func (wuo *WorkUpdateOne) SetStatus(w work.Status) *WorkUpdateOne {
	wuo.mutation.SetStatus(w)
	return wuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wuo *WorkUpdateOne) SetNillableStatus(w *work.Status) *WorkUpdateOne {
	if w != nil {
		wuo.SetStatus(*w)
	}
	return wuo
}

// SetPriority sets the "priority" field.
func (wuo *WorkUpdateOne) SetPriority(i int) *WorkUpdateOne {
	wuo.mutation.ResetPriority()
	wuo.mutation.SetPriority(i)
	return wuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (wuo *WorkUpdateOne) SetNillablePriority(i *int) *WorkUpdateOne {
	if i != nil {
		wuo.SetPriority(*i)
	}
	return wuo
}

// AddPriority adds i to the "priority" field.
func (wuo *WorkUpdateOne) AddPriority(i int) *WorkUpdateOne {
	wuo.mutation.AddPriority(i)
	return wuo
}

// AddChildIDs adds the "children" edge to the Work entity by IDs.
func (wuo *WorkUpdateOne) AddChildIDs(ids ...int) *WorkUpdateOne {
	wuo.mutation.AddChildIDs(ids...)
	return wuo
}

// AddChildren adds the "children" edges to the Work entity.
func (wuo *WorkUpdateOne) AddChildren(w ...*Work) *WorkUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Work entity by ID.
func (wuo *WorkUpdateOne) SetParentID(id int) *WorkUpdateOne {
	wuo.mutation.SetParentID(id)
	return wuo
}

// SetNillableParentID sets the "parent" edge to the Work entity by ID if the given value is not nil.
func (wuo *WorkUpdateOne) SetNillableParentID(id *int) *WorkUpdateOne {
	if id != nil {
		wuo = wuo.SetParentID(*id)
	}
	return wuo
}

// SetParent sets the "parent" edge to the Work entity.
func (wuo *WorkUpdateOne) SetParent(w *Work) *WorkUpdateOne {
	return wuo.SetParentID(w.ID)
}

// Mutation returns the WorkMutation object of the builder.
func (wuo *WorkUpdateOne) Mutation() *WorkMutation {
	return wuo.mutation
}

// ClearChildren clears all "children" edges to the Work entity.
func (wuo *WorkUpdateOne) ClearChildren() *WorkUpdateOne {
	wuo.mutation.ClearChildren()
	return wuo
}

// RemoveChildIDs removes the "children" edge to Work entities by IDs.
func (wuo *WorkUpdateOne) RemoveChildIDs(ids ...int) *WorkUpdateOne {
	wuo.mutation.RemoveChildIDs(ids...)
	return wuo
}

// RemoveChildren removes "children" edges to Work entities.
func (wuo *WorkUpdateOne) RemoveChildren(w ...*Work) *WorkUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveChildIDs(ids...)
}

// ClearParent clears the "parent" edge to the Work entity.
func (wuo *WorkUpdateOne) ClearParent() *WorkUpdateOne {
	wuo.mutation.ClearParent()
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkUpdateOne) Select(field string, fields ...string) *WorkUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Work entity.
func (wuo *WorkUpdateOne) Save(ctx context.Context) (*Work, error) {
	var (
		err  error
		node *Work
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkUpdateOne) SaveX(ctx context.Context) *Work {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkUpdateOne) check() error {
	if v, ok := wuo.mutation.Text(); ok {
		if err := work.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf("ent: validator failed for field \"text\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.Status(); ok {
		if err := work.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (wuo *WorkUpdateOne) sqlSave(ctx context.Context) (_node *Work, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   work.Table,
			Columns: work.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: work.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Work.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, work.FieldID)
		for _, f := range fields {
			if !work.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != work.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldText,
		})
	}
	if value, ok := wuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: work.FieldStatus,
		})
	}
	if value, ok := wuo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if value, ok := wuo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if wuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !wuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Work{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{work.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
