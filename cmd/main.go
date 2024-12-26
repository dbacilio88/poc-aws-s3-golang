package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	ss3 "github.com/aws/aws-sdk-go/service/s3"
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

	//manejo de sesiones default:
	sessionDefault := LoadSessionDefault()

	// listar buckets con la config default:
	ListBucketsDefault(sessionDefault)

	cfg := &aws.Config{
		Region: aws.String("us-east-1"),
		//Credentials: credentials.NewSharedCredentials("workspace/.aws/credentials", "default"),
	}
	//session
	ns, err := session.NewSession(cfg)
	if err != nil {
		errorPrint(err)
		return
	}

	s3Clint := ss3.New(ns)
	buckets, err := s3Clint.ListBuckets(nil)
	if err != nil {
		return
	}
	for _, bucket := range buckets.Buckets {
		log.Println(*bucket.Name)
	}
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
