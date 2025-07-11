package main

import (

	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)
}