package artsign

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"artsign/ent"
	"artsign/ent/work"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateWork(ctx context.Context, input ent.CreateWorkInput, image *string, fileExtension *string) (*ent.Work, error) {
	data, err := base64.StdEncoding.DecodeString(*image)
	if err != nil {
		return nil, err
	}
	wb := new(bytes.Buffer)
	_, err = wb.Write(data)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	output, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.config.S3.BucketName),
		Key:         aws.String(fmt.Sprintf("%s.%s", uuid.String(), *fileExtension)),
		Body:        wb,
		ContentType: aws.String(fmt.Sprintf("image/%s)", *fileExtension)),
	})
	if err != nil {
		return nil, err
	}

	return ent.FromContext(ctx).Work.
		Create().
		SetInput(input).
		SetImageURL(output.Location).
		Save(ctx)
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
