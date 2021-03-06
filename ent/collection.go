// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CategoryQuery) CollectFields(ctx context.Context, satisfies ...string) *CategoryQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CategoryQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CategoryQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CommentQuery) CollectFields(ctx context.Context, satisfies ...string) *CommentQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CommentQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CommentQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *ImageQuery) CollectFields(ctx context.Context, satisfies ...string) *ImageQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		i = i.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return i
}

func (i *ImageQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ImageQuery {
	return i
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PortfolioQuery) CollectFields(ctx context.Context, satisfies ...string) *PortfolioQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		po = po.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return po
}

func (po *PortfolioQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PortfolioQuery {
	return po
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *ToolQuery) CollectFields(ctx context.Context, satisfies ...string) *ToolQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *ToolQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ToolQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TreasureQuery) CollectFields(ctx context.Context, satisfies ...string) *TreasureQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TreasureQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TreasureQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (w *WorkQuery) CollectFields(ctx context.Context, satisfies ...string) *WorkQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		w = w.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return w
}

func (w *WorkQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *WorkQuery {
	return w
}
