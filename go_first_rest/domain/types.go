package domain

// main type
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// main var
var TASKS []Task = make([]Task, 0)
