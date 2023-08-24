package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type UpdateTask struct {
	TaskRepository repository.TaskRepository
}

func (u *UpdateTask) Execute(id, list, title, description string) *e.Error {
	appErr := e.New().SetLayer(e.Application)
	tid, er := uuid.Parse(id)
	if er != nil {
		return appErr.CustomError(e.InvalidId)
	}

	if list == "" || title == "" || description == "" {
		return appErr.CustomError(e.EmptyField)
	}

	tk, err := u.TaskRepository.GetById(tid)
	if err != nil {
		return err
	}

	tk.Title = title
	tk.Description = description

	return u.TaskRepository.Update(&tk)
}
