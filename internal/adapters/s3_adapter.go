package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/utils"
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
	helper *utils.Helper
}

type IS3Adapter interface {
	ListBuckets() ([]types.Bucket, error)
}

func NewS3Adapter(log *zap.Logger, region, profile string) IS3Adapter {
	log.Info("Initializing S3 Adapter with: ", zap.String("region", region), zap.String("profile", profile))
	uh := utils.NewHelper()
	credentials := uh.CredentialsAws()

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigFiles([]string{credentials.AwsConfiguration}),
		config.WithSharedCredentialsFiles([]string{credentials.AwsCredentials}),
		config.WithSharedConfigProfile("default"),
		config.WithLogConfigurationWarnings(true),
	)

	log.Info("Successfully loaded AWS config", zap.String("region", region))

	if err != nil {
		log.Error("Error loading AWS configuration", zap.Error(err))
		return nil
	}

	return &S3Adapter{
		Logger: log,

		client: s3.NewFromConfig(cfg),
		helper: utils.NewHelper(),
	}
}

func (a *S3Adapter) ListBuckets() ([]types.Bucket, error) {
	buckets, err := a.client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	return buckets.Buckets, err
}
