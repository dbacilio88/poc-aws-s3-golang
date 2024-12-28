package services

import (
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
	ListBucketFromS3()
}

func NewStorageService(log *zap.Logger, s3adapter adapters.IS3Adapter) IStorageService {
	return &StorageService{
		Logger:    log,
		s3adapter: s3adapter,
	}
}

func (s *StorageService) ListBucketFromS3() {

	s.Info("Initializing S3 service")
	buckets, err := s.s3adapter.ListBuckets()

	if err != nil {
		s.Error("Failed to list buckets", zap.Error(err))
		return
	}

	if len(buckets) == 0 {
		s.Info("No buckets found or You don't have any buckets!")
	} else {
		for _, bucket := range buckets {
			s.Info("Name bucket", zap.String("value", *bucket.Name))
			objects, err := s.s3adapter.ListObjects("DATA")
			if err != nil {
				s.Error("Failed to list objects")
				//s.Error("Failed to list objects", zap.Error(err))
				return
			}
			for _, object := range objects {
				s.Info("Name object", zap.String("value", *object.Key))
			}
		}
	}
}
