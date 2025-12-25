package handlers

import (
	"ecom-go/internal/models"
	"ecom-go/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ToDoHandlers struct {
	service *service.ToDoService
}

func NewHandlers(service *service.ToDoService) *ToDoHandlers {
	return &ToDoHandlers{service: service}
}

func (h *ToDoHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var task models.ToDoRequest
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if task.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	t, err := h.service.Create(r.Context(), &task)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "context deadline exceeded" {
			status = http.StatusGatewayTimeout
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func (h *ToDoHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll(r.Context())
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "context deadline exceeded" {
			status = http.StatusGatewayTimeout
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *ToDoHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "context deadline exceeded" {
			status = http.StatusGatewayTimeout
		} else if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *ToDoHandlers) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var task models.ToDoRequest
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(r.Context(), id, &task); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "context deadline exceeded" {
			status = http.StatusGatewayTimeout
		} else if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func (h *ToDoHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "context deadline exceeded" {
			status = http.StatusGatewayTimeout
		} else if err.Error() == "not found" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
