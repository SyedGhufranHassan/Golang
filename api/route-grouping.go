package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API Home"))
		})
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Status OK"))
		})
	})

	http.ListenAndServe(":8080", r)
}
