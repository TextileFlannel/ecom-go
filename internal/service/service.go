package service

import "ecom-go/internal/storage"

type ToDoService struct {
	storge *storage.MemoryStorage
}

func NewService(storge *storage.MemoryStorage) *ToDoService {
	return &ToDoService{storge: storge}
}
