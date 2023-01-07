package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type UpdateTask struct {
	Repository repository.TaskRepository
}

func (u *UpdateTask) Execute(id string, title string, description string) error {
	var err error

	tid, err := uuid.Parse(id)
	if err != nil {
		return herr.NewApp().InvalidTaskId
	}

	tk, err := u.Repository.GetById(tid)
	if err != nil {
		return err
	}

	tk.Title = title
	tk.Description = description
	
	err = u.Repository.Update(&tk)
	if err != nil {
		return err
	}

	return nil
}