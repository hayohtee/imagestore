package main

import "net/http"

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/posts", getAllPost)
	mux.HandleFunc("POST /v1/posts", createNewPost)
	mux.HandleFunc("DELETE /v1/posts/{id}", deletePost)

	return mux
}