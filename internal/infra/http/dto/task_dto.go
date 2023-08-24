package dto

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type TaskValue struct {
	m map[string]string
}

func NewTaskValue() *TaskValue {
	return &TaskValue{
		m: make(map[string]string),
	}
}

const (
	Id              = "id"
	TaskId          = "taskId"
	TaskCtx         = "taskCtx"
	TaskTitle       = "title"
	TaskDescription = "description"
)

func (t *TaskValue) Set(key string, value string) {
	t.m[key] = value
}
func (t *TaskValue) Get(key string) string {
	return t.m[key]
}

type SaveTaskBodyRequest struct {
	List        string `json:"list"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *SaveTaskBodyRequest) Validate() error {
	reqFields := requiredFields(*s)
	if len(reqFields) > 0 {
		return fmt.Errorf("Missing required fields: %v", reqFields)
	}

	if len(s.Title) > 50 {
		return errors.New("Title must be less than 50 characters")
	}

	if len(s.Description) > 255 {
		return errors.New("Description must be less than 255 characters")
	}

	return nil
}

type UpdateTaskBodyRequest struct {
	Id          string `json:"id"`
	List        string `json:"list"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (u *UpdateTaskBodyRequest) Validate() error {
	reqFields := requiredFields(*u)
	if len(reqFields) > 0 {
		return fmt.Errorf("Missing required fields: %v", reqFields)
	}

	if _, err := uuid.Parse(u.Id); err != nil {
		return errors.New("Invalid Id")
	}

	if len(u.List) > 20 {
		return errors.New("List must be less than 20 characters")
	}

	if len(u.Title) > 50 {
		return errors.New("Title must be less than 50 characters")
	}

	if len(u.Description) > 255 {
		return errors.New("Description must be less than 255 characters")
	}

	return nil
}

type TaskContext struct {
	taskId string
}

func (t *TaskContext) Validate() error {
	reqFields := requiredFields(*t)
	if len(reqFields) > 0 {
		return fmt.Errorf("Missing required fields: %v", reqFields)
	}

	return nil
}
