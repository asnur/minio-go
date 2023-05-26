//go:build example
// +build example

/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2018-2019 MinIO, Inc.
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
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY, my-bucketname, my-objectname and
	// my-testfile are dummy values, please replace them with original values.

	// Requests are always secure (HTTPS) by default. Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	s3Client, err := minio.New("s3.amazonaws.com", &minio.Options{
		Creds:  credentials.NewStaticV4("YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", ""),
		Secure: true,
	})

	objectInfo, err := s3Client.GetObjectACL(context.Background(), "my-bucketname", "my-objectname")
	if err != nil {
		log.Fatalln(err)
	}

	// print object owner informations
	fmt.Printf(`Object owner:
Display name: %q
ID: %q
`, objectInfo.Owner.DisplayName, objectInfo.Owner.ID)

	// print object grant informations
	for _, g := range objectInfo.Grant {
		fmt.Printf(`Object grant:
 - Display name: %q
 - ID: %q
 - URI: %q
 - Permission: %q
`, g.Grantee.DisplayName, g.Grantee.ID, g.Grantee.URI, g.Permission)
	}

	// print all value header (acl, metadata, standard header value...)
	for k, v := range objectInfo.Metadata {
		fmt.Println("key:", k)
		fmt.Printf(" - value: %v\n", v)
	}
}
