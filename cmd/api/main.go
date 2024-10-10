package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: ":4000",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler: routes(),
	}

	log.Println("listening on :4000")
	err = server.ListenAndServe(); if err != nil {
		log.Fatal(err)
	}
}