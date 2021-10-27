package artsign

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateWork(ctx context.Context, work WorkInput) (*ent.Work, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Works(ctx context.Context) ([]*ent.Work, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *workResolver) Parent(ctx context.Context, obj *ent.Work) (*ent.Work, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *workResolver) Children(ctx context.Context, obj *ent.Work) ([]*ent.Work, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Work returns WorkResolver implementation.
func (r *Resolver) Work() WorkResolver { return &workResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type workResolver struct{ *Resolver }
