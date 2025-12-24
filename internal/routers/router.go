package routers

import (
	"ecom-go/internal/handlers"
	"net/http"
)

func Setup(h *handlers.ToDoHandlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /todos/{id}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
