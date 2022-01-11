package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"artsign/ent/treasure"
	"artsign/ent/user"
	"artsign/ent/work"
	"artsign/graph/generated"
	"artsign/graph/model"
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateWork(ctx context.Context, input ent.CreateWorkInput, images []*graphql.Upload) (*ent.Work, error) {
	work, err := ent.FromContext(ctx).Work.
		Create().
		SetInput(input).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	bulk := make([]*ent.ImageCreate, len(images))
	for i, image := range images {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		output, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(r.config.S3.BucketName),
			Key:         aws.String(uuid.String()),
			Body:        image.File,
			ContentType: aws.String(image.ContentType),
			ACL:         types.ObjectCannedACLPublicRead,
		})
		if err != nil {
			return nil, err
		}

		bulk[i] = ent.FromContext(ctx).Image.Create().SetURL(output.Location).SetWork(work)
	}

	_, err = ent.FromContext(ctx).Image.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return work, nil
}

func (r *mutationResolver) UpdateWork(ctx context.Context, id int, input ent.UpdateWorkInput) (*ent.Work, error) {
	return ent.FromContext(ctx).Work.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateWorks(ctx context.Context, ids []int, input ent.UpdateWorkInput) ([]*ent.Work, error) {
	client := ent.FromContext(ctx)
	if err := client.Work.Update().Where(work.IDIn(ids...)).SetInput(input).Exec(ctx); err != nil {
		return nil, err
	}
	return client.Work.Query().Where(work.IDIn(ids...)).All(ctx)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) CreateUserLike(ctx context.Context, input model.CreateUserLikeInput) (*model.CreateUserLikePayload, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	_, err = user.Update().AddLikeIDs(input.WorkID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.CreateUserLikePayload{}, nil
}

func (r *mutationResolver) DeleteUserLike(ctx context.Context, input model.DeleteUserLikeInput) (*model.DeleteUserLikePayload, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	_, err = user.Update().RemoveLikeIDs(input.WorkID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.DeleteUserLikePayload{}, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input ent.CreateCommentInput) (*ent.Comment, error) {
	return ent.FromContext(ctx).Comment.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) CreateTreasure(ctx context.Context, input ent.CreateTreasureInput) (*model.CreateTreasurePayload, error) {
	_, err := ent.FromContext(ctx).Treasure.
		Create().
		SetOwnerID(input.OwnerID).
		SetWorkID(input.WorkID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.CreateTreasurePayload{}, nil
}

func (r *mutationResolver) DeleteTreasure(ctx context.Context, input model.DeleteTreasureInput) (*model.DeleteTreasurePayload, error) {
	_, err := ent.FromContext(ctx).Treasure.
		Delete().
		Where(
			treasure.HasOwnerWith(user.IDEQ(input.UserID)),
			treasure.HasWorkWith(work.IDEQ(input.WorkID))).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &model.DeleteTreasurePayload{}, nil
}

func (r *mutationResolver) CreateUserLikeComment(ctx context.Context, input model.CreateUserLikeCommentInput) (*model.CreateUserLikeCommentPayload, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	_, err = user.Update().AddLikeCommentIDs(input.CommentID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.CreateUserLikeCommentPayload{}, nil
}

func (r *queryResolver) Works(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.WorkOrder, where *ent.WorkWhereInput) (*ent.WorkConnection, error) {
	return r.client.Work.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithWorkOrder(orderBy),
			ent.WithWorkFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

func (r *queryResolver) Categories(ctx context.Context) ([]*ent.Category, error) {
	return r.client.Category.Query().All(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
