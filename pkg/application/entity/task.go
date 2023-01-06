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
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Status   Status    `json:"status"`
	UserId   uuid.UUID `json:"user_id"`
}

func NewTask(id uuid.UUID, title string, describe string, userId uuid.UUID) *Task {
	b := Base{}
	b.New(id)

	t := Task{
		Base:     b,
		Title:    title,
		Describe: describe,
		Status:   Undone,
		UserId:   userId,
	}

	return &t
}
