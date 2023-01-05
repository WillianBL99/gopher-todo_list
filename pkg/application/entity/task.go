package entity

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Done   Status = "done"
	Undone Status = "undone"
)

type Task struct {
	Base
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Status   Status    `json:"status"`
	DueDate  time.Time `json:"due_date"`
	UserId   uuid.UUID `json:"user_id"`
}

func NewTask(id uuid.UUID, title string, describe string, dueDate time.Time, userId uuid.UUID) *Task {
	if dueDate.IsZero() {
		dueDate = time.Now().AddDate(0, 0, 1)
	}

	b := Base{}
	b.New(id)

	t := Task{
		Base:     b,
		Title:    title,
		Describe: describe,
		Status:   Undone,
		DueDate:  dueDate,
		UserId:   userId,
	}

	return &t
}
