package artsign

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"artsign/ent/work"
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
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

func (r *mutationResolver) CreateUserLike(ctx context.Context, input CreateUserLikeInput) (*CreateUserLikePayload, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	_, err = user.Update().AddLikeIDs(input.WorkID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &CreateUserLikePayload{}, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input ent.CreateCommentInput) (*ent.Comment, error) {
	return ent.FromContext(ctx).Comment.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) CreateUserTreasure(ctx context.Context, input CreateUserTreasureInput) (*CreateUserTreasurePayload, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	_, err = user.Update().AddTreasureIDs(input.WorkID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &CreateUserTreasurePayload{}, nil
}

func (r *mutationResolver) CreateUserLikeComment(ctx context.Context, input CreateUserLikeCommentInput) (*CreateUserLikeCommentPayload, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	_, err = user.Update().AddLikeCommentIDs(input.CommentID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &CreateUserLikeCommentPayload{}, nil
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

func (r *userResolver) TreasureConnection(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.WorkOrder) (*ent.WorkConnection, error) {
	return obj.QueryTreasures().
		Paginate(ctx, after, first, before, last,
			ent.WithWorkOrder(orderBy),
		)
}

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

func (r *workResolver) CommentConnection(ctx context.Context, obj *ent.Work, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	return obj.QueryComments().
		Paginate(ctx, after, first, before, last,
			ent.WithCommentOrder(orderBy),
			ent.WithCommentFilter(where.Filter),
		)
}

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// Work returns WorkResolver implementation.
func (r *Resolver) Work() WorkResolver { return &workResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type workResolver struct{ *Resolver }
