//go:build example
// +build example

/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"log"

	minio "github.com/asnur/minio-go"
	"github.com/asnur/minio-go/pkg/credentials"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY and my-bucketname are
	// dummy values, please replace them with original values.

	// Requests are always secure (HTTPS) by default. Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	s3Client, err := minio.New("s3.amazonaws.com", &minio.Options{
		Creds:  credentials.NewStaticV4("YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// s3Client.TraceOn(os.Stderr)

	// Get object lock configuration.
	enabled, mode, validity, unit, err := s3Client.GetObjectLockConfig(context.Background(), "tbucket13a")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("object lock is %v for bucket 'my-bucketname'\n", enabled)

	if mode != nil {
		fmt.Printf("%v mode is enabled for %v %v for bucket 'my-bucketname'\n", *mode, *validity, *unit)
	} else {
		fmt.Println("No mode is enabled for bucket 'my-bucketname'")
	}
}