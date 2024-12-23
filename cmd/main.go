package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

func main() {

	defaultConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("default"))

	if err != nil {
		return
	}

	defaultConfig.Region = "us-east-1"

	client := s3.NewFromConfig(defaultConfig)
	fmt.Println(client)
}
