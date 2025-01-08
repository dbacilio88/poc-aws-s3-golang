package factory

import (
	"github.com/dbacilio88/poc-aws-s3-golang/config"
	"github.com/dbacilio88/poc-aws-s3-golang/internal/adapters/file-transfer/aws"
	"go.uber.org/zap"
)

/**
*
* registry
* <p>
* registry file
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
* @since 8/01/2025
*
 */

type AdapterRegistry func(log *zap.Logger, cfg config.Properties) (IAdapterFactory, bool)

var AdapterFactories []AdapterRegistry

func RegisterAdapter() {
	AdapterFactories = append(AdapterFactories, func(log *zap.Logger, cfg config.Properties) (IAdapterFactory, bool) {
		if cfg.Server.Environment != "local" {
			return aws.NewS3Adapter(), true
		}
		return nil, false
	})

}
