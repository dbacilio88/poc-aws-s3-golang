package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
*
* healthcheck
* <p>
* healthcheck file
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

func HealthCheckHandlerGin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"alive": true,
	})
}
