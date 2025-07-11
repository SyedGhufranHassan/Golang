package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		w.Write([]byte("User ID is: " + userID))
	})

	http.ListenAndServe(":8080", r)
}
