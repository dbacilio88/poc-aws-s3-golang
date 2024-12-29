package utils

import (
	"errors"
	"fmt"
	"github.com/aws/smithy-go"
)

/**
*
* s3_util
* <p>
* s3_util file
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
* @since 29/12/2024
*
 */

type AwsError struct {
}

func (e *AwsError) ValidateError(err error) error {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		fmt.Println(ae.ErrorCode())
		switch ae.ErrorCode() {

		case "SignatureDoesNotMatch":
			return errors.New("AWS Signature Does Not Match")
		case "AccessDenied":
			return errors.New("AWS Access Denied")
		case "InvalidCredentials":
			return errors.New("AWS Credentials Invalid")
		case "InvalidAccessKeyId":
			return errors.New("AWS Access KeyId Invalid")
		case "NoSuchBucket":
			return errors.New("AWS Bucket Not Found")
		case "NoSuchKey":
			return errors.New("AWS Key Not Found")
		case "NoSuchUploader":
			return errors.New("AWS Uploader Not Found")
		case "NotFound":
			return errors.New("AWS Bucket Not Found")
		case "InvalidBucketName":
			return errors.New("AWS Bucket Name Invalid")
		default:
			return errors.New(ae.ErrorCode())
		}
	}
	return errors.New(ae.ErrorCode())
}
