package service

import (
	"context"
	"ecom-go/internal/models"
	"errors"
	"testing"
	"time"
)

type mockStorage struct {
	tasks         map[int]*models.ToDo
	nextID        int
	delay         time.Duration
	shouldError   bool
	errorToReturn error
}

func newMockStorage() *mockStorage {
	return &mockStorage{
		tasks:  make(map[int]*models.ToDo),
		nextID: 1,
	}
}

func (m *mockStorage) Create(task *models.ToDoRequest) *models.ToDo {
	time.Sleep(m.delay)
	todo := &models.ToDo{
		ID:     m.nextID,
		Title:  task.Title,
		Body:   task.Body,
		IsDone: task.IsDone,
	}
	m.tasks[m.nextID] = todo
	m.nextID++
	return todo
}

func (m *mockStorage) GetAll() []*models.ToDo {
	time.Sleep(m.delay)
	result := make([]*models.ToDo, 0, len(m.tasks))
	for _, task := range m.tasks {
		result = append(result, task)
	}
	return result
}

func (m *mockStorage) GetByID(id int) (*models.ToDo, error) {
	time.Sleep(m.delay)
	if m.shouldError {
		return nil, m.errorToReturn
	}
	task, exists := m.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (m *mockStorage) Update(id int, task *models.ToDoRequest) error {
	time.Sleep(m.delay)
	if m.shouldError {
		return m.errorToReturn
	}
	existing, exists := m.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	existing.Title = task.Title
	existing.Body = task.Body
	existing.IsDone = task.IsDone
	return nil
}

func (m *mockStorage) Delete(id int) error {
	time.Sleep(m.delay)
	if m.shouldError {
		return m.errorToReturn
	}
	if _, exists := m.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(m.tasks, id)
	return nil
}

func TestCreateService_Timeout(t *testing.T) {
	mock := newMockStorage()
	mock.delay = 6 * time.Second
	svc := NewService(mock)

	req := &models.ToDoRequest{
		Title: "Test Task",
		Body:  "Test Body",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := svc.Create(ctx, req)
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestGetAllService_Timeout(t *testing.T) {
	mock := newMockStorage()
	mock.delay = 6 * time.Second
	svc := NewService(mock)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := svc.GetAll(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestGetByIDService_NotFound(t *testing.T) {
	mock := newMockStorage()
	svc := NewService(mock)

	_, err := svc.GetByID(context.Background(), 999)
	if err == nil {
		t.Error("expected error for non-existent task, got nil")
	}

	if err.Error() != "task not found" {
		t.Errorf("expected 'task not found' error, got %v", err)
	}
}

func TestGetByIDService_Timeout(t *testing.T) {
	mock := newMockStorage()
	mock.delay = 6 * time.Second
	svc := NewService(mock)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := svc.GetByID(ctx, 1)
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestUpdateService_Timeout(t *testing.T) {
	mock := newMockStorage()
	created := mock.Create(&models.ToDoRequest{Title: "Test"})
	mock.delay = 6 * time.Second
	svc := NewService(mock)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := svc.Update(ctx, created.ID, &models.ToDoRequest{Title: "Updated"})
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestDeleteService_Timeout(t *testing.T) {
	mock := newMockStorage()
	created := mock.Create(&models.ToDoRequest{Title: "Test"})
	mock.delay = 6 * time.Second
	svc := NewService(mock)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := svc.Delete(ctx, created.ID)
	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestDeleteService_NotFound(t *testing.T) {
	mock := newMockStorage()
	svc := NewService(mock)

	err := svc.Delete(context.Background(), 999)
	if err == nil {
		t.Error("expected error for non-existent task, got nil")
	}
	if err.Error() != "task not found" {
		t.Errorf("expected 'task not found' error, got %v", err)
	}
}
