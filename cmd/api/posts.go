package main

import (
	"log"
	"net/http"
)

func (app *application) createNewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10485760)
	log.Println(r.FormValue("name"))
	log.Println(r.FormValue("email"))

	_, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "error parsing getting image from the form-data", http.StatusUnprocessableEntity)
		return
	}
	log.Println(header.Filename)
	log.Println(r.FormValue("caption"))
}

func (app *application) deletePost(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getAllPost(w http.ResponseWriter, r *http.Request) {

}
