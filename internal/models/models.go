package models

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	IsDone bool   `json:"isDone"`
}

type ToDoRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	IsDone bool   `json:"isDone"`
}
