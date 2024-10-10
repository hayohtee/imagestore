package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/posts", app.getAllPost)
	mux.HandleFunc("POST /v1/posts", app.createNewPost)
	mux.HandleFunc("DELETE /v1/posts/{id}", app.deletePost)

	return mux
}
