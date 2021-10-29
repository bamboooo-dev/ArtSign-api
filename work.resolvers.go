package artsign

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"context"
)

func (r *mutationResolver) CreateWork(ctx context.Context, work WorkInput) (*ent.Work, error) {
	client := ent.FromContext(ctx)
	return client.Work.Create().
		SetText(work.Text).
		SetStatus(work.Status).
		SetNillablePriority(work.Priority). // Set the "priority" field if provided.
		SetNillableParentID(work.Parent).   // Set the "parent_id" field if provided.
		Save(ctx)
}

func (r *queryResolver) Works(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.WorkOrder) (*ent.WorkConnection, error) {
	return r.client.Work.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithWorkOrder(orderBy),
		)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
