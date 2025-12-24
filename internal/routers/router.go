package routers

import (
	"ecom-go/internal/handlers"
	"net/http"
)

func Setup(h *handlers.ToDoHandlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /todos", h.Create)
	mux.HandleFunc("GET /todos", h.GetAll)
	mux.HandleFunc("GET /todos/{id}", h.GetByID)
	mux.HandleFunc("PUT /todos/{id}", h.Update)
	mux.HandleFunc("DELETE /todos/{id}", h.Delete)

	return mux
}
