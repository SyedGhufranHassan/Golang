package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List users"))
	})
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create user"))
	})
	r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get user by ID"))
	})

	// Walk through the registered routes
	err := chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("Method: %s | Route: %s\n", method, route)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking routes:", err)
	}

	http.ListenAndServe(":8080", r)
}
