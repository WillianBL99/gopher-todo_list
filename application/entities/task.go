package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Base
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Done     bool      `json:"done"`
	DueDate  time.Time `json:"due_date"`
}

func NewTask(id uuid.UUID, title string, describe string, dueDate time.Time) *Task {
	if dueDate.IsZero() {
		dueDate = time.Now().AddDate(0, 0, 1)
	}

	b := Base{}
	b.New(id)

	t := Task{
		Title:    title,
		Describe: describe,
		Done:     false,
		DueDate:  dueDate,
		Base:     b,
	}

	return &t
}
