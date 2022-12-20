package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing Home"))
}

func (app *application) ShowPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusNotFound)
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, "Showing post with id %d", id)
}
func (app *application) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Post form page"))
}
func (app *application) CreatePostForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create post"))
}
