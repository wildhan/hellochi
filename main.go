package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)    // Example middleware usage
	r.Use(middleware.Recoverer) // Example middleware usage

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte("Hello, " + name + "!"))
	})
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
