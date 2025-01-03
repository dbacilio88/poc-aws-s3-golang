package services

import (
	"context"
	"fmt"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/adapters"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/models/response/bucket"
	"go.uber.org/zap"
)

/**
*
* storage
* <p>
* storage file
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

type BucketService struct {
	s3adapter adapters.IS3Adapter
	*zap.Logger
}

type IBucketService interface {
	ListBucketFromS3(ctx context.Context, search string)
	ListBucket(ctx context.Context, search string) ([]bucket.ListBuckets, error)
	ListObjects(ctx context.Context, bucket string) ([]bucket.ListObjects, error)
}

func NewStorageService(log *zap.Logger) IBucketService {
	return &BucketService{
		Logger:    log,
		s3adapter: adapters.NewS3Adapter(log, "", ""),
	}
}

func (s *BucketService) ListBucket(ctx context.Context, search string) ([]bucket.ListBuckets, error) {
	s.Info("Initializing ListBucket S3 - Bucket service", zap.String("search", search))

	buckets, err := s.s3adapter.ListBuckets(ctx)

	if err != nil {
		return nil, err
	}

	result := make([]bucket.ListBuckets, 0, len(buckets))

	if len(buckets) == 0 {
		s.Info("Successfully ListBucket S3 - Bucket service", zap.String("bucket", search), zap.Int("count", len(result)))
		return result, nil
	}

	for _, bk := range buckets {
		br := bucket.ListBuckets{
			Name: *bk.Name,
			Date: *bk.CreationDate,
			//Region: *bk.BucketRegion,
		}
		result = append(result, br)
	}
	s.Info("Successfully ListBucket S3 - Bucket service", zap.String("bucket", search), zap.Int("count", len(result)))
	return result, nil
}

func (s *BucketService) ListObjects(ctx context.Context, name string) ([]bucket.ListObjects, error) {
	s.Info("Initializing ListObjects S3  - Bucket service", zap.String("bucket", name))

	objects, err := s.s3adapter.ListObjects(ctx, name)
	if err != nil {
		return nil, err
	}

	result := make([]bucket.ListObjects, 0, len(objects))

	if len(objects) == 0 {
		s.Info("Successfully ListObjects S3 - Bucket service", zap.String("bucket", name), zap.Int("count", len(objects)))
		return result, nil
	}

	for _, ob := range objects {
		or := bucket.ListObjects{
			Name: *ob.Key,
		}
		result = append(result, or)
	}

	s.Info("Successfully ListObjects S3 - Bucket service", zap.String("bucket", name), zap.Int("count", len(objects)))

	return result, nil
}
func (s *BucketService) ListBucketFromS3(ctx context.Context, search string) {
	s.Info("Initializing S3 service: ")
	buckets, err := s.s3adapter.ListBuckets(ctx)

	if err != nil {
		s.Error("Failed to list buckets: ", zap.String("error", err.Error()))
		return
	}

	if len(buckets) == 0 {
		s.Info("No buckets found or You don't have any buckets!:")
	} else {
		for _, bucket := range buckets {
			s.Info("List buckets by region: ", zap.String("name", *bucket.Name))
			region, err := s.s3adapter.ListBucketByRegion(*bucket.Name)
			if err != nil {
				s.Error("Failed to list buckets by region: ", zap.String("error", err.Error()))
				return
			}

			if region == search {
				s.Info("List objects by region: ", zap.String("name", *bucket.Name))
				objects, err := s.s3adapter.ListObjects(ctx, *bucket.Name)
				if err != nil {
					s.Error("Error to list objects: ", zap.Error(err))
					return
				}
				for _, object := range objects {
					s.Info("Name object: ", zap.String("value", *object.Key))
				}
			}
		}
	}

	bucket, err := s.s3adapter.CreateBucket("christian-cb", "us-east-2")
	//bucket, err := s.s3adapter.CreateBucket("data", "us-east-1")
	if err != nil {
		s.Error("Error to create bucket: ", zap.String("error", err.Error()))
		return
	}
	fmt.Println(bucket)

}
