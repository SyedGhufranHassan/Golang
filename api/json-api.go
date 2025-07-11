package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	type Todo struct {
		ID int `json:"id"`
		Note string `json:"note"`
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
	
	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	})

	http.ListenAndServe(":8080", r)
}