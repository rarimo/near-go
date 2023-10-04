package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	defaultRequestPayer = "requester"
	Delimiter           = "/"
)

type Config struct {
	Bucket          string `fig:"bucket,required"`
	Region          string `fig:"region,required"`
	AccessKey       string `fig:"access_key,required"`
	SecretAccessKey string `fig:"secret_access_key,required"`
}

type Connector interface {
	GetObject(ctx context.Context, key string) (*s3.GetObjectOutput, error)
	ListObjects(ctx context.Context, limit int64, startAfter string) (*s3.ListObjectsV2Output, error)
}

type connector struct {
	client *s3.S3
	Config
}

func New(cfg Config) (Connector, error) {
	awsCfg := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretAccessKey, "")).
		WithRegion(cfg.Region)

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}

	res := &connector{client: s3.New(sess), Config: cfg}

	return res, nil
}

func (c *connector) GetObject(ctx context.Context, key string) (*s3.GetObjectOutput, error) {
	input := &s3.GetObjectInput{
		Key:          aws.String(key),
		Bucket:       aws.String(c.Config.Bucket),
		RequestPayer: aws.String(defaultRequestPayer),
	}

	res, err := c.client.GetObjectWithContext(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get object from s3")
	}

	return res, nil
}

func (c *connector) ListObjects(ctx context.Context, limit int64, startAfter string) (*s3.ListObjectsV2Output, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:       aws.String(c.Config.Bucket),
		MaxKeys:      aws.Int64(limit),
		StartAfter:   aws.String(startAfter),
		Delimiter:    aws.String(Delimiter),
		RequestPayer: aws.String(defaultRequestPayer),
	}

	res, err := c.client.ListObjectsV2WithContext(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list objects from s3")
	}

	return res, nil
}
