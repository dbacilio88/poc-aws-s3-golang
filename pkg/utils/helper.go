package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

/**
*
* helper
* <p>
* helper file
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

//const rootProject = "./"

type Helper struct {
	RootDir          string
	AwsCredentials   string
	AwsConfiguration string
}

func NewHelper(rootProject string) *Helper {
	root, _ := filepath.Abs(path.Join(rootProject))
	return &Helper{RootDir: root}
}

func (h *Helper) CredentialsAws() *Helper {
	credentials, _ := filepath.Abs(filepath.Join(h.RootDir, "deploy", "cloud", ".aws"))
	fmt.Println("credentials: ", credentials)
	h.AwsCredentials = strings.Join([]string{credentials, "credentials"}, "/")
	h.AwsConfiguration = strings.Join([]string{credentials, "config"}, "/")
	return h
}
