package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type DeleteTask struct {
	TaskRepository repository.TaskRepository
}

func (d *DeleteTask) Execute(id string) error {
	pid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return d.TaskRepository.Delete(pid)
}
