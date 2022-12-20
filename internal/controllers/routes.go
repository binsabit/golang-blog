package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", app.Home).Methods("GET")
	r.HandleFunc("/post/{id:[0-9]+}", app.ShowPost).Methods("GET")
	r.HandleFunc("/post/create", app.CreatePost).Methods("POST")
	r.HandleFunc("/post/create", app.CreatePostForm).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(app.config.dir))))

	return r
}
