package handlers

import "ecom-go/internal/service"

type ToDoHandlers struct {
	service *service.ToDoService
}

func NewHandlers(service *service.ToDoService) *ToDoHandlers {
	return &ToDoHandlers{service: service}
}
