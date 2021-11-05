package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Config struct {
	EndpointURL  string `envconfig:"endpoint_url"`
	UsePathStyle bool   `envconfig:"use_path_style"`
	BucketName   string `envconfig:"bucket_name"`
}

func NewUploader(ctx context.Context, cfg Config) (*manager.Uploader, error) {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           cfg.EndpointURL,
			SigningRegion: "ap-northeast-1",
		}, nil
	})
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"), config.WithEndpointResolver(customResolver))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = cfg.UsePathStyle
	})
	return manager.NewUploader(s3Client), nil
}
