package artsign

import (
	"artsign/ent"
	"artsign/pkg/env"

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
}

// NewSchema creates a graphql executable schema.
func NewSchema(config *env.Config, client *ent.Client, uploader *manager.Uploader) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{config, client, uploader},
	})
}
