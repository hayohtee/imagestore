package main

import "github.com/aws/aws-sdk-go-v2/service/s3"

type application struct {
	s3Client *s3.Client
}