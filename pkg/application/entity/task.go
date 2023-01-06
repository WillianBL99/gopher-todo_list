package entity

import (
	"github.com/google/uuid"
)

type Status string

const (
	Done   Status = "done"
	Undone Status = "undone"
)

type Task struct {
	Base
	Title       string    `json:"title"`
	Status      Status    `json:"status"`
	UserId      uuid.UUID `json:"user_id"`
	Description string    `json:"description"`
}

func NewTask(id uuid.UUID, title string, description string, userId uuid.UUID) *Task {
	b := Base{}
	b.New(id)

	t := Task{
		Base:        b,
		Title:       title,
		Status:      Undone,
		UserId:      userId,
		Description: description,
	}

	return &t
}
