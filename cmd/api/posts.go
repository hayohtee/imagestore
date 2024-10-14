package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (app *application) createNewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10485760)
	log.Println(r.FormValue("name"))
	log.Println(r.FormValue("email"))

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "error parsing getting image from the form-data", http.StatusUnprocessableEntity)
		return
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	
	randBytes, err := generateRandomBytes(32)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = app.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key: aws.String(string(randBytes)),
		Body: file,
		ContentType: aws.String("images/*"),
	})

	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("could not upload %s", header.Filename), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("image uploaded successfully"))
}

func (app *application) deletePost(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getAllPost(w http.ResponseWriter, r *http.Request) {

}
