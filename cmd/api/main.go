package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(os.Getenv("AWS_BUCKET_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KE"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)

	app := application{
		s3Client: s3.NewFromConfig(awsConfig),
	}

	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:         ":4000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      app.routes(),
	}

	log.Println("listening on :4000")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
