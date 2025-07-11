package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"

)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var db *sql.DB

func main() {
	var err error


	db, err = sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	))
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	r.Route("/todos", func(r chi.Router) {
		r.Get("/", getAllTodos)        // GET /todos
		r.Post("/", createTodo)        // POST /todos
		r.Get("/{id}", getTodo)        // GET /todos/{id}
		r.Put("/{id}", updateTodo)     // PUT /todos/{id}
		r.Delete("/{id}", deleteTodo)  // DELETE /todos/{id}
	})

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)

	err := db.QueryRow("INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id", t.Title, t.Done).Scan(&t.ID)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}


func getAllTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, done FROM todos")
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Title, &t.Done)
		todos = append(todos, t)
	}

	json.NewEncoder(w).Encode(todos)
}


func getTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var t Todo
	err := db.QueryRow("SELECT id, title, done FROM todos WHERE id=$1", id).Scan(&t.ID, &t.Title, &t.Done)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t)
}


func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var t Todo
	json.NewDecoder(r.Body).Decode(&t)

	t.ID, _ = strconv.Atoi(id)

	_, err := db.Exec("UPDATE todos SET title=$1, done=$2 WHERE id=$3", t.Title, t.Done, t.ID)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(t)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
