package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"artsign/graph/generated"
	"context"
)

func (r *workResolver) ImageConnection(ctx context.Context, obj *ent.Work, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ImageOrder) (*ent.ImageConnection, error) {
	return obj.QueryImages().
		Paginate(ctx, after, first, before, last,
			ent.WithImageOrder(orderBy),
		)
}

func (r *workResolver) LikerConnection(ctx context.Context, obj *ent.Work, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return obj.QueryLikers().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}

func (r *workResolver) TreasurerConnection(ctx context.Context, obj *ent.Work, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return obj.QueryTreasurers().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}

func (r *workResolver) CommentConnection(ctx context.Context, obj *ent.Work, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	return obj.QueryComments().
		Paginate(ctx, after, first, before, last,
			ent.WithCommentOrder(orderBy),
			ent.WithCommentFilter(where.Filter),
		)
}

// Work returns generated.WorkResolver implementation.
func (r *Resolver) Work() generated.WorkResolver { return &workResolver{r} }

type workResolver struct{ *Resolver }
