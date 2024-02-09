package models

import (
	"encoding/json"
	"io"
	"time"
)

// Task structure
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Status string

var (
	StatusDone    Status = "done"
	StatusPending Status = "pending"
)

func WriteJSONResponse(w io.Writer, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func FindTaskByID(tasks []Task, id string) *Task {
	for _, task := range tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}
