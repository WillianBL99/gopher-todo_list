package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type GetAllTasks struct {
	TasksRepository repository.TaskRepository
}

func (g *GetAllTasks) Execute(uid string) ([]entity.Task, error) {
	puid, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}

	tasks, err := g.TasksRepository.GetAll(puid)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
