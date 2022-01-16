package graph

import (
	"artsign/ent"
	"artsign/graph/generated"
	"artsign/graph/model"
	"artsign/pkg/env"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	config   *env.Config
	client   *ent.Client
	uploader *manager.Uploader
	rooms    map[int]*model.Room
	mu       sync.Mutex
}

// NewSchema creates a graphql executable schema.
func NewSchema(config *env.Config, client *ent.Client, uploader *manager.Uploader) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			config:   config,
			client:   client,
			uploader: uploader,
			rooms:    map[int]*model.Room{},
		},
	})
}
