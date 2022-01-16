package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"artsign/graph/generated"
	"context"
)

func (r *commentResolver) ChildrenConnection(ctx context.Context, obj *ent.Comment, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder) (*ent.CommentConnection, error) {
	return obj.QueryChildren().Paginate(ctx, after, first, before, last,
		ent.WithCommentOrder(orderBy),
	)
}

func (r *commentResolver) LikerConnection(ctx context.Context, obj *ent.Comment, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return obj.QueryLikers().Paginate(ctx, after, first, before, last,
		ent.WithUserOrder(orderBy),
	)
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
