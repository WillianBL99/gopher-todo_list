package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type UpdateTask struct {
	Repository repository.TaskRepository
}

func (u *UpdateTask) Execute(id string, title string, description string) error {
	var err error

	tid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("Error to parse taskId: %v", err.Error())
	}

	tk, err := u.Repository.GetById(tid)
	if err != nil {
		return fmt.Errorf("Task not found: %v", err.Error())
	}

	tk.Title = title
	tk.Description = description
	
	err = u.Repository.Update(&tk)
	if err != nil {
		return fmt.Errorf("Error to update task: %v", err.Error())
	}

	return nil
}