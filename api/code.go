// ✅ Step 1: Basic Go HTTP Server using net/http
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

// ✅ Step 2: First Chi Application
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Chi Router"))
    })

    http.ListenAndServe(":8080", r)
}

// ✅ Step 3: URL Parameters Example
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

// ✅ Step 4: Route Grouping (Nested Routes)
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

// ✅ Step 5: Middleware Usage
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Home with Middleware"))
    })

    r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
        panic("Intentional Panic")
    })

    http.ListenAndServe(":8080", r)
}

// ✅ Step 6: Serving Static Files
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

// ✅ Step 7: JSON API Example
package main

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
)

type Todo struct {
    ID   int    `json:"id"`
    Note string `json:"note"`
}

func main() {
    r := chi.NewRouter()

    r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
        var todo Todo
        json.NewDecoder(r.Body).Decode(&todo)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todo)
    })

    http.ListenAndServe(":8080", r)
}

// ✅ Step 8: Testing with httptest
package main

import (
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!"))
}

func TestHello(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    Hello(w, req)

    resp := w.Result()
    body, _ := io.ReadAll(resp.Body)

    if string(body) != "Hello!" {
        t.Errorf("expected Hello!, got %s", body)
    }
}
