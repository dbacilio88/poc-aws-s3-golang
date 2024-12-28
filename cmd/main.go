package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	config2 "github.com/dbacilio88/poc-aws-s3-golang/config"
	"github.com/dbacilio88/poc-aws-s3-golang/pkg/logs"
	"go.uber.org/zap"
	"log"
)

/**
*
* main
* <p>
* main file
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
* @since 20/12/2024
*
 */

const REGION = "us-west-2"

func LoadSessionDefault() *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		//config.WithRegion(REGION),
		//config.WithSharedConfigProfile("default"),
	)

	log.Println(cfg)

	if err != nil {
		errorPrint(err)
		return nil
	}

	client := s3.NewFromConfig(cfg)

	return client
}

func ListBucketsDefault(client *s3.Client) {
	buckets, err := client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		errorPrint(err)
		return
	}
	for _, bucket := range buckets.Buckets {
		log.Println(*bucket.Name)
	}

}

func errorPrint(err error) {
	log.Fatal("Error:", err)
}

func main() {

	//configuration properties:
	if err := config2.LoadProperties(); err != nil {
		errorPrint(err)
	}

	//configuration logs:
	l, err := logs.LoggerConfiguration(config2.YAML.Server.Environment)

	std := zap.RedirectStdLog(l)

	defer std()

	if err != nil {
		return
	}

	fmt.Println(l)

	//manejo de sesiones default:
	sessionDefault := LoadSessionDefault()

	// listar buckets con la config default:
	ListBucketsDefault(sessionDefault)

	//m

	/*



		newSession, err := session.NewSession(cfg)

		if err != nil {
			return
		}

		bucketName := "s3-golang-bucket-test"

		svc := s3.New(newSession)

		buckets, err := svc.ListBuckets(nil)
		if err != nil {
			return
		}

		for _, bucket := range buckets.Buckets {
			fmt.Println(*bucket.Name)
		}

		objects, err := svc.ListObjects(&s3.ListObjectsInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			return
		}
		for _, o := range objects.Contents {
			fmt.Println(*o.Key)
		}

	*/
}
