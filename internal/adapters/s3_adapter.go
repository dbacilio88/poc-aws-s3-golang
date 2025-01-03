package adapters

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/dbacilio88/poc-aws-s3-golang/pkg/utils"
	"go.uber.org/zap"
)

/**
*
* s3_adapter
* <p>
* s3_adapter file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 27/12/2024
*
 */

type S3Adapter struct {
	*zap.Logger
	client *s3.Client
	*utils.AwsError
}

type IS3Adapter interface {
	ListBuckets(ctx context.Context) ([]types.Bucket, error)
	ListObjects(ctx context.Context, bucketName string) ([]types.Object, error)
	ExistBucket(bucketName string) (bool, error)
	CreateBucket(bucketName, region string) (bool, error)
	ListBucketByRegion(bucketName string) (string, error)
}

func NewS3Adapter(log *zap.Logger, region, profile string) IS3Adapter {
	log.Info("Initializing S3 Adapter with: ", zap.String("region", region), zap.String("profile", profile))
	uh := utils.NewHelper(".")
	credentials := uh.CredentialsAws()

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithSharedConfigFiles([]string{credentials.AwsConfiguration}),
		config.WithSharedCredentialsFiles([]string{credentials.AwsCredentials}),
		config.WithSharedConfigProfile(profile),
		config.WithLogConfigurationWarnings(true),
	)

	log.Info("Successfully loaded AWS config: ", zap.String("region", region))

	if err != nil {
		log.Warn("Error loading AWS configuration: ", zap.Error(err))
		return nil
	}

	return &S3Adapter{
		Logger:   log,
		client:   s3.NewFromConfig(cfg),
		AwsError: &utils.AwsError{},
	}
}

func (a *S3Adapter) ListBuckets(ctx context.Context) ([]types.Bucket, error) {
	a.Info("Listing buckets from AWS: ")
	buckets, err := a.client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, a.ValidateError(err)
	}
	a.Info("Successfully listed buckets: ", zap.Int("buckets", len(buckets.Buckets)))
	return buckets.Buckets, err
}

func (a *S3Adapter) ListBucketByRegion(bucketName string) (string, error) {
	a.Info("Listing buckets by region: ", zap.String("region", bucketName))
	location, err := a.client.GetBucketLocation(context.Background(), &s3.GetBucketLocationInput{Bucket: &bucketName})
	if err != nil {
		return "", err
	}

	constraint := string(location.LocationConstraint)

	if constraint == "" {
		constraint = "us-east-1"
	}

	a.Info("Listed buckets by region: ", zap.String("bucket", bucketName), zap.String("region", constraint))
	return constraint, nil
}

func (a *S3Adapter) ListObjects(ctx context.Context, bucketName string) ([]types.Object, error) {
	a.Info("Listing objects by bucket from AWS: ", zap.String("bucket", bucketName))
	var err error
	var objects []types.Object
	output := &s3.ListObjectsV2Output{}

	_, err = a.ExistBucket(bucketName)
	if err != nil {
		return nil, err
	}

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	op := s3.NewListObjectsV2Paginator(a.client, input)

	for op.HasMorePages() {
		output, err = op.NextPage(context.Background())
		if err != nil {
			return nil, a.ValidateError(err)
		}

		objects = append(objects, output.Contents...)
	}
	a.Info("Successfully listed objects: ", zap.Int("objects", len(objects)))

	if len(objects) == 0 {
		return []types.Object{}, nil
	}

	return objects, nil
}

func (a *S3Adapter) ExistBucket(bucketName string) (bool, error) {

	a.Info("Checking if bucket exists: ", zap.String("bucket", bucketName))

	_, err := a.client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		return false, a.ValidateError(err)
	}

	a.Info("Bucket exists: ", zap.String("bucket", bucketName))
	return true, nil
}

func (a *S3Adapter) CreateBucket(bucketName, region string) (bool, error) {

	a.Info("Creating bucket: ", zap.String("bucket", bucketName), zap.String("region", region))

	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	if region != "us-east-1" {
		cnf := &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		}
		input.CreateBucketConfiguration = cnf
	}

	bucket, err := a.client.CreateBucket(context.Background(), input)

	if err != nil {
		return false, a.ValidateError(err)
	}
	fmt.Println(bucket)

	a.Info("Successfully created bucket: ", zap.String("bucket", bucketName))

	return true, nil
}
