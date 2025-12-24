package service

import (
	"ecom-go/internal/models"
	"ecom-go/internal/storage"
)

type ToDoService struct {
	storge *storage.MemoryStorage
}

func NewService(storge *storage.MemoryStorage) *ToDoService {
	return &ToDoService{storge: storge}
}

func (s *ToDoService) Create(task *models.ToDoRequest) error {
	return s.storge.Create(task)
}

func (s *ToDoService) GetAll() []*models.ToDo {
	return s.storge.GetAll()
}

func (s *ToDoService) GetByID(id int) (*models.ToDo, error) {
	return s.storge.GetByID(id)
}

func (s *ToDoService) Update(id int, task *models.ToDoRequest) error {
	return s.storge.Update(id, task)
}

func (s *ToDoService) Delete(id int) error {
	return s.storge.Delete(id)
}
