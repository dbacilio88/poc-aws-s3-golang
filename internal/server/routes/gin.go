package routes

import (
	"github.com/dbacilio88/poc-aws-s3-golang/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
)

/**
*
* server
* <p>
* server file
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

type GinFramework struct {
	*zap.Logger
	router     *gin.Engine
	middleware *negroni.Negroni
	port       Port
	name       Name
	handler    handlers.IS3Bucket
}

func newGinFramework(log *zap.Logger, port Port, name Name) *GinFramework {
	return &GinFramework{
		Logger:     log,
		router:     gin.Default(),
		middleware: negroni.New(),
		port:       port,
		name:       name,
		handler:    handlers.NewBucketsHandler(log),
	}
}

func (g *GinFramework) Run() {
	g.router.GET("/health", handlers.HealthCheckHandlerGin)
	g.router.GET("/buckets", g.handler.GetBuckets)
	g.router.GET("/buckets/:name", g.handler.GetObjects)
	g.middleware.UseHandler(g.router)

	listenAndServe(g.port, g.name, g.middleware, g.Logger)
}
