package server

import (
	"github.com/dbacilio88/poc-aws-s3-golang/internal/server/routes"
	"go.uber.org/zap"
)

/**
*
* http
* <p>
* http file
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

type HttpConfig struct {
	*zap.Logger
	factory routes.IServerFactory
	port    routes.Port
	name    routes.Name
}

func NewHttpConfig(log *zap.Logger) *HttpConfig {
	return &HttpConfig{
		Logger: log,
	}
}

func (h *HttpConfig) NewHttpServer(instance int) *HttpConfig {
	factory, err := routes.NewServerFactory(h.Logger, instance, h.port, h.name)
	if err != nil {
		return nil
	}
	h.factory = factory
	return h
}

func (h *HttpConfig) Port(port int) *HttpConfig {
	h.port = routes.Port(port)
	return h
}

func (h *HttpConfig) Name(name string) *HttpConfig {
	h.name = routes.Name(name)
	return h
}

func (h *HttpConfig) Start() {
	h.factory.Run()
}
