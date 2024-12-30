package services

import (
	"fmt"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/adapters"
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

type StorageService struct {
	s3adapter adapters.IS3Adapter
	*zap.Logger
}

type IStorageService interface {
	ListBucketFromS3(search string)
}

func NewStorageService(log *zap.Logger, s3adapter adapters.IS3Adapter) IStorageService {
	return &StorageService{
		Logger:    log,
		s3adapter: s3adapter,
	}
}

func (s *StorageService) ListBucketFromS3(search string) {
	s.Info("Initializing S3 service: ")
	buckets, err := s.s3adapter.ListBuckets()

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
				objects, err := s.s3adapter.ListObjects(*bucket.Name)
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
