package main
import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users"))
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All Products"))
}

func getuserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("User ID: " + userID))
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", getAllUsers)
			r.Get("/{userID}", getuserByID)
		})
		r.Route("/products", func(r chi.Router) {
			r.Get("/", getAllProducts)
		})
	})
	
	fmt.Println("server running on http://localhost:8080/")
	http.ListenAndServe(":8080", r)
}

