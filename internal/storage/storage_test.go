package storage

import (
	"ecom-go/internal/models"
	"testing"
)

func TestMemoryStorage_Create(t *testing.T) {
	store := NewMemoryStorage()

	taskReq := &models.ToDoRequest{
		Title: "test",
		Body:  "test",
	}

	taskRes := store.Create(taskReq)

	if taskRes.ID != 1 {
		t.Errorf("Expected task ID to be 1, got %d", taskRes.ID)
	}

	if taskRes.Title != "test" {
		t.Errorf("Expected task Title to be test, got %s", taskRes.Title)
	}

	if taskRes.Body != "test" {
		t.Errorf("Expected task Body to be test, got %s", taskRes.Body)
	}

	taskRes2 := store.Create(taskReq)
	if taskRes2.ID != 2 {
		t.Errorf("Expected task ID to be 2, got %d", taskRes2.ID)
	}

	if len(store.data) != 2 {
		t.Errorf("Expected data to be 2, got %d", len(store.data))
	}
}

func TestMemoryStorage_GetAll(t *testing.T) {
	store := NewMemoryStorage()

	tasks := store.GetAll()
	if len(tasks) != 0 {
		t.Errorf("Expected no tasks, got %d", len(tasks))
	}

	task := &models.ToDoRequest{
		Title: "test",
		Body:  "test",
	}

	store.Create(task)
	store.Create(task)

	tasks = store.GetAll()
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestMemoryStorage_GetByID(t *testing.T) {
	store := NewMemoryStorage()

	task := &models.ToDoRequest{
		Title: "test",
		Body:  "test",
	}

	_, err := store.GetByID(1)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	store.Create(task)
	_, err = store.GetByID(1)
	if err != nil {
		t.Errorf("Unexpected error getting task by ID")
	}
}

func TestMemoryStorage_Update(t *testing.T) {
	store := NewMemoryStorage()

	task := &models.ToDoRequest{
		Title: "test",
		Body:  "test",
	}

	err := store.Update(1, task)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	store.Create(task)
	err = store.Update(1, &models.ToDoRequest{
		Title:  "updated",
		Body:   "updated",
		IsDone: true,
	})

	if err != nil {
		t.Errorf("Unexpected error updating task")
	}

	task1, _ := store.GetByID(1)

	if task1.Title != "updated" {
		t.Errorf("Expected updated task, got %s", task1.Title)
	}
	if task1.Body != "updated" {
		t.Errorf("Expected updated task, got %s", task1.Body)
	}
	if task1.IsDone != true {
		t.Errorf("Expected updated task to be done true, got %t", task1.IsDone)
	}

}

func TestMemoryStorage_Delete(t *testing.T) {
	store := NewMemoryStorage()
	task := &models.ToDoRequest{
		Title: "test",
		Body:  "test",
	}

	err := store.Delete(1)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	store.Create(task)

	err = store.Delete(1)
	if err != nil {
		t.Errorf("Unexpected error deleting task")
	}
	if len(store.data) != 0 {
		t.Errorf("Expected data to be empty, got %d", len(store.data))
	}
}
