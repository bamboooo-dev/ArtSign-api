package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"artsign/ent/user"
	"artsign/graph/generated"
	"context"
)

func (r *userResolver) ViewerIsFollowing(ctx context.Context, obj *ent.User) (bool, error) {
	return obj.
		QueryFollowers().
		Where(user.ID(17179869184)).
		Exist(ctx)
}

func (r *userResolver) WorkConnection(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.WorkOrder) (*ent.WorkConnection, error) {
	return obj.QueryWorks().
		Paginate(ctx, after, first, before, last,
			ent.WithWorkOrder(orderBy),
		)
}

func (r *userResolver) LikeConnection(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.WorkOrder) (*ent.WorkConnection, error) {
	return obj.QueryLikes().
		Paginate(ctx, after, first, before, last,
			ent.WithWorkOrder(orderBy),
		)
}

func (r *userResolver) TreasureConnection(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TreasureOrder) (*ent.TreasureConnection, error) {
	return obj.QueryTreasures().
		Paginate(ctx, after, first, before, last,
			ent.WithTreasureOrder(orderBy),
		)
}

func (r *userResolver) FolloweeConnection(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return obj.QueryFollowees().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
