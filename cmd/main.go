package main

import (
	"github.com/dbacilio88/poc-aws-s3-golang/config"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/server"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/server/routes"
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

const REGION = "us-east-1"
const PROFILE = "default"

func errorPrint(err error) {
	log.Fatal("Error:", err)
}

func main() {

	//configuration properties:
	if err := config.LoadProperties(); err != nil {
		errorPrint(err)
	}

	//configuration logs:
	l, err := logs.LoggerConfiguration(config.YAML.Server.Environment)

	std := zap.RedirectStdLog(l)

	defer std()

	if err != nil {
		errorPrint(err)
		return
	}

	//configuration http server:
	srv := server.NewHttpConfig(l).
		Port(config.YAML.Server.Port).
		Name(routes.NameRouterGin).
		NewHttpServer(routes.InstanceRouterGin)

	// start instance server http
	srv.Start()

	//s3Instance := adapters.NewS3Adapter(l, REGION, PROFILE)
	//storageInstance := services.NewStorageService(l, s3Instance)
	//storageInstance.ListBucketFromS3(REGION)
}
