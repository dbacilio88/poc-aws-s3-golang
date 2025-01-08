package aws

/**
*
* s3
* <p>
* s3 file
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

// S3Adapter struct, for s3 adapter
type S3Adapter struct{}

// NewS3Adapter func, create instance adapter
func NewS3Adapter() *S3Adapter {
	return &S3Adapter{}
}

// Connection func, create connection to S3 AWS
func (a *S3Adapter) Connection() error {
	return nil
}

// Download func, Get object from S3 AWS
func (a *S3Adapter) Download() (interface{}, error) {
	return nil, nil
}

// Upload func, Put object to S3 AWS
func (a *S3Adapter) Upload() error {
	return nil
}

// Disconnection func, close connection to S3 AWS
func (a *S3Adapter) Disconnection() error {
	return nil
}
