// Code generated by entc, DO NOT EDIT.

package ent

import (
	"artsign/ent/category"
	"artsign/ent/user"
	"artsign/ent/work"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// CategoryEdge is the edge representation of Category.
type CategoryEdge struct {
	Node   *Category `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CategoryConnection is the connection containing edges to Category.
type CategoryConnection struct {
	Edges      []*CategoryEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

// CategoryPaginateOption enables pagination customization.
type CategoryPaginateOption func(*categoryPager) error

// WithCategoryOrder configures pagination ordering.
func WithCategoryOrder(order *CategoryOrder) CategoryPaginateOption {
	if order == nil {
		order = DefaultCategoryOrder
	}
	o := *order
	return func(pager *categoryPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCategoryOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCategoryFilter configures pagination filter.
func WithCategoryFilter(filter func(*CategoryQuery) (*CategoryQuery, error)) CategoryPaginateOption {
	return func(pager *categoryPager) error {
		if filter == nil {
			return errors.New("CategoryQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type categoryPager struct {
	order  *CategoryOrder
	filter func(*CategoryQuery) (*CategoryQuery, error)
}

func newCategoryPager(opts []CategoryPaginateOption) (*categoryPager, error) {
	pager := &categoryPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCategoryOrder
	}
	return pager, nil
}

func (p *categoryPager) applyFilter(query *CategoryQuery) (*CategoryQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *categoryPager) toCursor(c *Category) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *categoryPager) applyCursors(query *CategoryQuery, after, before *Cursor) *CategoryQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCategoryOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *categoryPager) applyOrder(query *CategoryQuery, reverse bool) *CategoryQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCategoryOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCategoryOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Category.
func (c *CategoryQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CategoryPaginateOption,
) (*CategoryConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCategoryPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}

	conn := &CategoryConnection{Edges: []*CategoryEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := c.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := c.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	c = pager.applyCursors(c, after, before)
	c = pager.applyOrder(c, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		c = c.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := c.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Category
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Category {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Category {
			return nodes[i]
		}
	}

	conn.Edges = make([]*CategoryEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &CategoryEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// CategoryOrderFieldName orders Category by name.
	CategoryOrderFieldName = &CategoryOrderField{
		field: category.FieldName,
		toCursor: func(c *Category) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f CategoryOrderField) String() string {
	var str string
	switch f.field {
	case category.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f CategoryOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *CategoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("CategoryOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *CategoryOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid CategoryOrderField", str)
	}
	return nil
}

// CategoryOrderField defines the ordering field of Category.
type CategoryOrderField struct {
	field    string
	toCursor func(*Category) Cursor
}

// CategoryOrder defines the ordering of Category.
type CategoryOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CategoryOrderField `json:"field"`
}

// DefaultCategoryOrder is the default ordering of Category.
var DefaultCategoryOrder = &CategoryOrder{
	Direction: OrderDirectionAsc,
	Field: &CategoryOrderField{
		field: category.FieldID,
		toCursor: func(c *Category) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Category into CategoryEdge.
func (c *Category) ToEdge(order *CategoryOrder) *CategoryEdge {
	if order == nil {
		order = DefaultCategoryOrder
	}
	return &CategoryEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := u.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		u = u.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}

	conn.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}

// WorkEdge is the edge representation of Work.
type WorkEdge struct {
	Node   *Work  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// WorkConnection is the connection containing edges to Work.
type WorkConnection struct {
	Edges      []*WorkEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// WorkPaginateOption enables pagination customization.
type WorkPaginateOption func(*workPager) error

// WithWorkOrder configures pagination ordering.
func WithWorkOrder(order *WorkOrder) WorkPaginateOption {
	if order == nil {
		order = DefaultWorkOrder
	}
	o := *order
	return func(pager *workPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultWorkOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithWorkFilter configures pagination filter.
func WithWorkFilter(filter func(*WorkQuery) (*WorkQuery, error)) WorkPaginateOption {
	return func(pager *workPager) error {
		if filter == nil {
			return errors.New("WorkQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type workPager struct {
	order  *WorkOrder
	filter func(*WorkQuery) (*WorkQuery, error)
}

func newWorkPager(opts []WorkPaginateOption) (*workPager, error) {
	pager := &workPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultWorkOrder
	}
	return pager, nil
}

func (p *workPager) applyFilter(query *WorkQuery) (*WorkQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *workPager) toCursor(w *Work) Cursor {
	return p.order.Field.toCursor(w)
}

func (p *workPager) applyCursors(query *WorkQuery, after, before *Cursor) *WorkQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultWorkOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *workPager) applyOrder(query *WorkQuery, reverse bool) *WorkQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultWorkOrder.Field {
		query = query.Order(direction.orderFunc(DefaultWorkOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Work.
func (w *WorkQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...WorkPaginateOption,
) (*WorkConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newWorkPager(opts)
	if err != nil {
		return nil, err
	}

	if w, err = pager.applyFilter(w); err != nil {
		return nil, err
	}

	conn := &WorkConnection{Edges: []*WorkEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := w.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := w.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	w = pager.applyCursors(w, after, before)
	w = pager.applyOrder(w, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		w = w.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		w = w.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := w.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Work
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Work {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Work {
			return nodes[i]
		}
	}

	conn.Edges = make([]*WorkEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &WorkEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// WorkOrderFieldTitle orders Work by title.
	WorkOrderFieldTitle = &WorkOrderField{
		field: work.FieldTitle,
		toCursor: func(w *Work) Cursor {
			return Cursor{
				ID:    w.ID,
				Value: w.Title,
			}
		},
	}
	// WorkOrderFieldCreatedAt orders Work by created_at.
	WorkOrderFieldCreatedAt = &WorkOrderField{
		field: work.FieldCreatedAt,
		toCursor: func(w *Work) Cursor {
			return Cursor{
				ID:    w.ID,
				Value: w.CreatedAt,
			}
		},
	}
	// WorkOrderFieldUpdatedAt orders Work by updated_at.
	WorkOrderFieldUpdatedAt = &WorkOrderField{
		field: work.FieldUpdatedAt,
		toCursor: func(w *Work) Cursor {
			return Cursor{
				ID:    w.ID,
				Value: w.UpdatedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f WorkOrderField) String() string {
	var str string
	switch f.field {
	case work.FieldTitle:
		str = "TITLE"
	case work.FieldCreatedAt:
		str = "CREATED_AT"
	case work.FieldUpdatedAt:
		str = "UPDATED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f WorkOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *WorkOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("WorkOrderField %T must be a string", v)
	}
	switch str {
	case "TITLE":
		*f = *WorkOrderFieldTitle
	case "CREATED_AT":
		*f = *WorkOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *WorkOrderFieldUpdatedAt
	default:
		return fmt.Errorf("%s is not a valid WorkOrderField", str)
	}
	return nil
}

// WorkOrderField defines the ordering field of Work.
type WorkOrderField struct {
	field    string
	toCursor func(*Work) Cursor
}

// WorkOrder defines the ordering of Work.
type WorkOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *WorkOrderField `json:"field"`
}

// DefaultWorkOrder is the default ordering of Work.
var DefaultWorkOrder = &WorkOrder{
	Direction: OrderDirectionAsc,
	Field: &WorkOrderField{
		field: work.FieldID,
		toCursor: func(w *Work) Cursor {
			return Cursor{ID: w.ID}
		},
	},
}

// ToEdge converts Work into WorkEdge.
func (w *Work) ToEdge(order *WorkOrder) *WorkEdge {
	if order == nil {
		order = DefaultWorkOrder
	}
	return &WorkEdge{
		Node:   w,
		Cursor: order.Field.toCursor(w),
	}
}
