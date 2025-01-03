package handlers

import (
	"fmt"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

/**
*
* buckets
* <p>
* buckets file
*
* Copyright (c) 2025 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 2/01/2025
*
 */

type S3Bucket struct {
	*zap.Logger
	serviceBucket services.IBucketService
}

type IS3Bucket interface {
	GetBuckets(ctx *gin.Context)
	GetObjects(ctx *gin.Context)
}

func NewBucketsHandler(log *zap.Logger) IS3Bucket {
	return &S3Bucket{
		Logger:        log,
		serviceBucket: services.NewStorageService(log),
	}
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func (b *S3Bucket) GetBuckets(ctx *gin.Context) {
	b.Info("Handle GET Buckets")
	bucket, err := b.serviceBucket.ListBucket(ctx, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, bucket)

}

func (b *S3Bucket) GetObjects(ctx *gin.Context) {
	b.Info("Handle GET Buckets")

	name := ctx.Param("name")
	fmt.Println(name)

	bucket, err := b.serviceBucket.ListObjects(ctx, name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println(bucket)

	ctx.JSON(http.StatusOK, bucket)

}
