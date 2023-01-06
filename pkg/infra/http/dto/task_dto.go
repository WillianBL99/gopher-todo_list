package dto

import "errors"

type TaskValue struct {
	m map[string]string
}

func NewTaskValue() *TaskValue {
	return &TaskValue{
		m: make(map[string]string),
	}
}

const (
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
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *SaveTaskBodyRequest) Validate() error {
	if s.Title == "" {
		return errors.New("Title is required")
	}

	if len(s.Title) > 50 {
		return errors.New("Title must be less than 50 characters")
	}

	if s.Description == "" {
		return errors.New("Description is required")
	}

	if len(s.Description) > 255 {
		return errors.New("Description must be less than 255 characters")
	}

	return nil
}

type UpdateTaskBodyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (u *UpdateTaskBodyRequest) Validate() error {
	if u.Title == "" && u.Description == "" {
		return errors.New("Title or Description is required")
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
	if t.taskId == "" {
		return errors.New("TaskId is required")
	}

	return nil
}
