package adapters

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
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
}

type IS3Adapter interface {
	ListBuckets() ([]types.Bucket, error)
	ListObjects(bucketName string) ([]types.Object, error)
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

	log.Info("Successfully loaded AWS config", zap.String("region", region))

	if err != nil {
		log.Error("Error loading AWS configuration", zap.Error(err))
		return nil
	}

	return &S3Adapter{
		Logger: log,
		client: s3.NewFromConfig(cfg),
	}
}

func (a *S3Adapter) ListBuckets() ([]types.Bucket, error) {
	buckets, err := a.client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "AccessDenied" {
			a.Error("AccessDenied - Bucket Not Found")
			return nil, errors.New(ae.ErrorCode())
		} else {
			a.Error("Error listing buckets", zap.Error(err))
			return nil, err
		}
	}
	return buckets.Buckets, err
}

func (a *S3Adapter) ListObjects(bucketName string) ([]types.Object, error) {

	var err error
	var objects []types.Object
	output := &s3.ListObjectsV2Output{}

	err = a.ExistBucket(bucketName)
	fmt.Println("ListObjects - Bucket: ", err)
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
			var nsb *types.NoSuchBucket
			if errors.As(err, &nsb) {
				//fmt.Println("Bucket Not Found", nsb)
				//a.Error("No such bucket", zap.String("bucket", bucketName))
				//fmt.Println("BACILIO ", err)
				//a.Error("Bucket does not exist", zap.String("bucketName", bucketName))
				//return nil, fmt.Errorf("el bucket '%s' no existe", bucketName)
				return nil, err
			}
			//			a.Logger.Error("Error al listar objetos", zap.Error(err))
			//return nil, fmt.Errorf("error al listar objetos en el bucket '%s': %w", bucketName, err)
			return nil, err
			//fmt.Println("Error listing objects", zap.Error(err))
			//return nil, err
		}
		objects = append(objects, output.Contents...)

	}

	return objects, nil
}

func (a *S3Adapter) ExistBucket(bucketName string) error {
	_, err := a.client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		var nf *types.NotFound
		if errors.As(err, &nf) {
			//fmt.Println("Error checking if bucket exists", err)
			return fmt.Errorf("el bucket '%s' no existe o no tienes acceso a él", bucketName)
		}
		return fmt.Errorf("error al verificar el bucket '%s': %w", bucketName, err)
	}

	return nil

}
