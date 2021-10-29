// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/predicate"
	"artsign/ent/work"
	"context"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeWork = "Work"
)

// WorkMutation represents an operation that mutates the Work nodes in the graph.
type WorkMutation struct {
	config
	op              Op
	typ             string
	id              *int
	title           *string
	description     *string
	image_url       *string
	updated_at      *time.Time
	text            *string
	created_at      *time.Time
	status          *work.Status
	priority        *int
	addpriority     *int
	clearedFields   map[string]struct{}
	children        map[int]struct{}
	removedchildren map[int]struct{}
	clearedchildren bool
	parent          *int
	clearedparent   bool
	done            bool
	oldValue        func(context.Context) (*Work, error)
	predicates      []predicate.Work
}

var _ ent.Mutation = (*WorkMutation)(nil)

// workOption allows management of the mutation configuration using functional options.
type workOption func(*WorkMutation)

// newWorkMutation creates new mutation for the Work entity.
func newWorkMutation(c config, op Op, opts ...workOption) *WorkMutation {
	m := &WorkMutation{
		config:        c,
		op:            op,
		typ:           TypeWork,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWorkID sets the ID field of the mutation.
func withWorkID(id int) workOption {
	return func(m *WorkMutation) {
		var (
			err   error
			once  sync.Once
			value *Work
		)
		m.oldValue = func(ctx context.Context) (*Work, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Work.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWork sets the old Work of the mutation.
func withWork(node *Work) workOption {
	return func(m *WorkMutation) {
		m.oldValue = func(context.Context) (*Work, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WorkMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WorkMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WorkMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTitle sets the "title" field.
func (m *WorkMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *WorkMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *WorkMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *WorkMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *WorkMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *WorkMutation) ResetDescription() {
	m.description = nil
}

// SetImageURL sets the "image_url" field.
func (m *WorkMutation) SetImageURL(s string) {
	m.image_url = &s
}

// ImageURL returns the value of the "image_url" field in the mutation.
func (m *WorkMutation) ImageURL() (r string, exists bool) {
	v := m.image_url
	if v == nil {
		return
	}
	return *v, true
}

// OldImageURL returns the old "image_url" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldImageURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldImageURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldImageURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldImageURL: %w", err)
	}
	return oldValue.ImageURL, nil
}

// ResetImageURL resets all changes to the "image_url" field.
func (m *WorkMutation) ResetImageURL() {
	m.image_url = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *WorkMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *WorkMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *WorkMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetText sets the "text" field.
func (m *WorkMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *WorkMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *WorkMutation) ResetText() {
	m.text = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *WorkMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *WorkMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *WorkMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetStatus sets the "status" field.
func (m *WorkMutation) SetStatus(w work.Status) {
	m.status = &w
}

// Status returns the value of the "status" field in the mutation.
func (m *WorkMutation) Status() (r work.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldStatus(ctx context.Context) (v work.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *WorkMutation) ResetStatus() {
	m.status = nil
}

// SetPriority sets the "priority" field.
func (m *WorkMutation) SetPriority(i int) {
	m.priority = &i
	m.addpriority = nil
}

// Priority returns the value of the "priority" field in the mutation.
func (m *WorkMutation) Priority() (r int, exists bool) {
	v := m.priority
	if v == nil {
		return
	}
	return *v, true
}

// OldPriority returns the old "priority" field's value of the Work entity.
// If the Work object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkMutation) OldPriority(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPriority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPriority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriority: %w", err)
	}
	return oldValue.Priority, nil
}

// AddPriority adds i to the "priority" field.
func (m *WorkMutation) AddPriority(i int) {
	if m.addpriority != nil {
		*m.addpriority += i
	} else {
		m.addpriority = &i
	}
}

// AddedPriority returns the value that was added to the "priority" field in this mutation.
func (m *WorkMutation) AddedPriority() (r int, exists bool) {
	v := m.addpriority
	if v == nil {
		return
	}
	return *v, true
}

// ResetPriority resets all changes to the "priority" field.
func (m *WorkMutation) ResetPriority() {
	m.priority = nil
	m.addpriority = nil
}

// AddChildIDs adds the "children" edge to the Work entity by ids.
func (m *WorkMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the Work entity.
func (m *WorkMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the Work entity was cleared.
func (m *WorkMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the Work entity by IDs.
func (m *WorkMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.children, ids[i])
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the Work entity.
func (m *WorkMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *WorkMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *WorkMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// SetParentID sets the "parent" edge to the Work entity by id.
func (m *WorkMutation) SetParentID(id int) {
	m.parent = &id
}

// ClearParent clears the "parent" edge to the Work entity.
func (m *WorkMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared reports if the "parent" edge to the Work entity was cleared.
func (m *WorkMutation) ParentCleared() bool {
	return m.clearedparent
}

// ParentID returns the "parent" edge ID in the mutation.
func (m *WorkMutation) ParentID() (id int, exists bool) {
	if m.parent != nil {
		return *m.parent, true
	}
	return
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *WorkMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *WorkMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// Where appends a list predicates to the WorkMutation builder.
func (m *WorkMutation) Where(ps ...predicate.Work) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *WorkMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Work).
func (m *WorkMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WorkMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.title != nil {
		fields = append(fields, work.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, work.FieldDescription)
	}
	if m.image_url != nil {
		fields = append(fields, work.FieldImageURL)
	}
	if m.updated_at != nil {
		fields = append(fields, work.FieldUpdatedAt)
	}
	if m.text != nil {
		fields = append(fields, work.FieldText)
	}
	if m.created_at != nil {
		fields = append(fields, work.FieldCreatedAt)
	}
	if m.status != nil {
		fields = append(fields, work.FieldStatus)
	}
	if m.priority != nil {
		fields = append(fields, work.FieldPriority)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WorkMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case work.FieldTitle:
		return m.Title()
	case work.FieldDescription:
		return m.Description()
	case work.FieldImageURL:
		return m.ImageURL()
	case work.FieldUpdatedAt:
		return m.UpdatedAt()
	case work.FieldText:
		return m.Text()
	case work.FieldCreatedAt:
		return m.CreatedAt()
	case work.FieldStatus:
		return m.Status()
	case work.FieldPriority:
		return m.Priority()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WorkMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case work.FieldTitle:
		return m.OldTitle(ctx)
	case work.FieldDescription:
		return m.OldDescription(ctx)
	case work.FieldImageURL:
		return m.OldImageURL(ctx)
	case work.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case work.FieldText:
		return m.OldText(ctx)
	case work.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case work.FieldStatus:
		return m.OldStatus(ctx)
	case work.FieldPriority:
		return m.OldPriority(ctx)
	}
	return nil, fmt.Errorf("unknown Work field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkMutation) SetField(name string, value ent.Value) error {
	switch name {
	case work.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case work.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case work.FieldImageURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetImageURL(v)
		return nil
	case work.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case work.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case work.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case work.FieldStatus:
		v, ok := value.(work.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case work.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriority(v)
		return nil
	}
	return fmt.Errorf("unknown Work field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WorkMutation) AddedFields() []string {
	var fields []string
	if m.addpriority != nil {
		fields = append(fields, work.FieldPriority)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WorkMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case work.FieldPriority:
		return m.AddedPriority()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkMutation) AddField(name string, value ent.Value) error {
	switch name {
	case work.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPriority(v)
		return nil
	}
	return fmt.Errorf("unknown Work numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WorkMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WorkMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WorkMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Work nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WorkMutation) ResetField(name string) error {
	switch name {
	case work.FieldTitle:
		m.ResetTitle()
		return nil
	case work.FieldDescription:
		m.ResetDescription()
		return nil
	case work.FieldImageURL:
		m.ResetImageURL()
		return nil
	case work.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case work.FieldText:
		m.ResetText()
		return nil
	case work.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case work.FieldStatus:
		m.ResetStatus()
		return nil
	case work.FieldPriority:
		m.ResetPriority()
		return nil
	}
	return fmt.Errorf("unknown Work field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WorkMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.children != nil {
		edges = append(edges, work.EdgeChildren)
	}
	if m.parent != nil {
		edges = append(edges, work.EdgeParent)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WorkMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case work.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	case work.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WorkMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchildren != nil {
		edges = append(edges, work.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WorkMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case work.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WorkMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedchildren {
		edges = append(edges, work.EdgeChildren)
	}
	if m.clearedparent {
		edges = append(edges, work.EdgeParent)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WorkMutation) EdgeCleared(name string) bool {
	switch name {
	case work.EdgeChildren:
		return m.clearedchildren
	case work.EdgeParent:
		return m.clearedparent
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WorkMutation) ClearEdge(name string) error {
	switch name {
	case work.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown Work unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WorkMutation) ResetEdge(name string) error {
	switch name {
	case work.EdgeChildren:
		m.ResetChildren()
		return nil
	case work.EdgeParent:
		m.ResetParent()
		return nil
	}
	return fmt.Errorf("unknown Work edge %s", name)
}
